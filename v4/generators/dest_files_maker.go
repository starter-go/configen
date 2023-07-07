package generators

import (
	"crypto/md5"
	"strings"

	"github.com/starter-go/afs"
	"github.com/starter-go/base/lang"
	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/gocode"
	"github.com/starter-go/vlog"
)

type destFilesMaker struct{}

func (inst *destFilesMaker) run(c *v4.Context) error {
	all := c.Destinations
	ctx := &makingContext{}
	for _, dest := range all {
		ctx.context = c
		ctx.folderDest = dest
		err := inst.makeDestFolder(ctx, dest)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *destFilesMaker) makeDestFolder(ctx *makingContext, dest *gocode.DestinationFolder) error {
	ctx = ctx.clone()
	path := dest.Path.GetPath()
	srclist := ctx.folderDest.Sources

	builder := &complexGoFileBuilder{}
	builder.packageSimpleName = dest.PackageSimpleName
	builder.target = dest.Path.GetChild("configen-main-gen.go")
	ctx.complexBuilder = builder

	vlog.Info("make destination:%s, path=%s", dest.ID, path)

	for _, src := range srclist {
		ctx.folderSrc = src
		err := inst.makeFromSrcFolder(ctx, src)
		if err != nil {
			return err
		}
	}

	return builder.WriteToFile()
}

func (inst *destFilesMaker) makeFromSrcFolder(ctx *makingContext, srcdir *gocode.SourceFolder) error {
	ctx = ctx.clone()

	builder := &simpleGoFileBuilder{}
	builder.packageSimpleName = ctx.folderDest.PackageSimpleName
	builder.target = inst.getTargetFile(ctx, srcdir)
	builder.hub = ctx.complexBuilder

	ctx.simpleBuilder = builder

	vlog.Info("  make from source: %s", srcdir.ID)
	list := ctx.listCurrentSourceFiles().List()
	for _, srcfile := range list {
		err := inst.makeFromSrcFile(ctx, srcfile)
		if err != nil {
			return err
		}
	}

	// text := ctx.simpleBuilder.buffer.String()
	// vlog.Warn("configen-xxxx-gen.go:\n %s", text)

	return builder.WriteToFile()
}

func (inst *destFilesMaker) getTargetFile(ctx *makingContext, srcdir *gocode.SourceFolder) afs.Path {
	srcid := srcdir.ID.String()
	dir := ctx.folderDest.Path
	name := "configen-src-" + srcid + "-gen.go"
	return dir.GetChild(name)
}

func (inst *destFilesMaker) makeFromSrcFile(ctx *makingContext, src *gocode.Source) error {
	ctx = ctx.clone()
	path := src.Path.GetPath()
	vlog.Info("  make from %s, path=%s", src.Name, path)
	list := src.TypeStructSet.List()
	for _, ts := range list {
		err := inst.makeComponent(ctx, ts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *destFilesMaker) makeComponent(ctx *makingContext, ts *gocode.TypeStruct) error {
	ctx = ctx.clone()

	if ts.ComScope == "" {
		ts.ComScope = "singleton"
	}

	if ts.ComID == "" {
		ts.ComID = inst.makeDefaultComponentID(ts)
	}

	if ts.IsComponent {
		vlog.Info("  make component, type=%s, id=%s", ts.Name, ts.ComID)
		err := ctx.simpleBuilder.WriteComponent(ts)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *destFilesMaker) makeDefaultComponentID(ts *gocode.TypeStruct) string {

	name := ts.Name
	alias := ts.OwnerPackage.Alias
	pack := ts.OwnerPackage.FullName

	sum := md5.Sum([]byte(pack))
	hex := lang.HexFromBytes(sum[0:8])

	return "com-" + hex.String() + "-" + alias + "-" + name
}

////////////////////////////////////////////////////////////////////////////////

type makingContext struct {
	context        *v4.Context
	folderDest     *gocode.DestinationFolder
	folderSrc      *gocode.SourceFolder
	complexBuilder *complexGoFileBuilder
	simpleBuilder  *simpleGoFileBuilder
}

func (inst *makingContext) clone() *makingContext {
	to := &makingContext{}
	to.context = inst.context
	to.folderDest = inst.folderDest
	to.folderSrc = inst.folderSrc
	to.complexBuilder = inst.complexBuilder
	to.simpleBuilder = inst.simpleBuilder
	return to
}

func (inst *makingContext) listCurrentSourceFiles() *gocode.SourceList {
	dst := &gocode.SourceList{}
	src := inst.context.GoFiles.List()
	path1 := inst.folderSrc.Path.GetPath()
	for _, item := range src {
		path2 := item.Path.GetPath()
		if strings.HasPrefix(path2, path1) {
			dst.Add(item)
		}
	}
	return dst
}
