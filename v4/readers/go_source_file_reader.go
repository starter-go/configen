package readers

import (
	"strings"

	"github.com/starter-go/afs"
	v4 "github.com/starter-go/configen/v4"
)

// 定义几个特殊的模式 token ...
const (
	PatternTokenAnyChar = "<?>"
	PatternTokenHex     = "<*>"
	PatternTokenMore    = "<...>"
)

// codeFragment 用来区分同一个代码实体的不同部分
type codeFragment string

// GoSourceFileReader ...
type GoSourceFileReader struct {
	context *v4.Context
	sf      *v4.SourceFolder
	pack    *v4.Package
	source  *v4.Source

	hGroupForStarterToken goCodeRowHandlerGroup
	hGroupForGoCode       goCodeRowHandlerGroup
	hGroupForGoCodeWithST goCodeRowHandlerGroup // go_code + starter_token
	hGroupForNone         goCodeRowHandlerGroup

	// status
	inImportBlock     bool
	inTypeStructBlock bool
	inFuncBlock       bool

	importItems map[string]*v4.CodeImport // map[alias/fullname] item
}

// NewGoSourceFileReader ...
func NewGoSourceFileReader() *GoSourceFileReader {
	reader := &GoSourceFileReader{}
	return reader
}

// Init ...
func (inst *GoSourceFileReader) Init(ctx *v4.Context, src *v4.SourceFolder) {
	inst.context = ctx
	inst.sf = src
	inst.pack = &v4.Package{}
	inst.source = &v4.Source{}
	inst.importItems = make(map[string]*v4.CodeImport)

	inst.initGroupStarterToken()
	inst.initGroupGoCode()
	inst.initGroupGoCodeWithST()
}

func (inst *GoSourceFileReader) initGroupStarterToken() {
	group := &inst.hGroupForStarterToken
	group.add(&rowHandlerForStarterComponent{})
	group.add(&rowHandlerForStarterInject{})
	group.add(&rowHandlerForStarterInterface{})
}

func (inst *GoSourceFileReader) initGroupGoCode() {
	group := &inst.hGroupForGoCode
	group.add(&rowHandlerForPackage{})
	group.add(&rowHandlerForImport{})
	group.add(&rowHandlerForTypeStruct{})
	group.add(&rowHandlerForFuncOfStruct{})
}

func (inst *GoSourceFileReader) initGroupGoCodeWithST() {
	group := &inst.hGroupForGoCodeWithST
	group.add(&rowHandlerForStarterInject{})
}

func (inst *GoSourceFileReader) Read(file afs.Path) (*v4.Source, error) {

	pack := &v4.Package{}
	result := &v4.Source{}
	inst.pack = pack
	inst.source = result

	result.Name = file.GetName()
	result.Path = file // .GetPath()
	result.OwnerPackage = pack

	rows, err := ReadRows(file)
	if err != nil {
		return nil, err
	}

	err = inst.parseRows(rows)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (inst *GoSourceFileReader) parseRows(rows []string) error {
	for i, row := range rows {
		err := inst.parseRow(row, i+1)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *GoSourceFileReader) getHandler0() goCodeRowHandler {
	return &inst.hGroupForNone
}

func (inst *GoSourceFileReader) parseRow(row string, rowNum int) error {
	const (
		prefixStarterToken = "//starter:"
	)
	h := inst.getHandler0()
	if row == "" {
		return nil
	} else if strings.HasPrefix(row, "//") {
		if strings.HasPrefix(row, prefixStarterToken) {
			h = &inst.hGroupForStarterToken
		}
	} else {
		if !strings.Contains(row, prefixStarterToken) {
			h = &inst.hGroupForGoCode
		} else {
			h = &inst.hGroupForGoCodeWithST
		}
	}
	gcRow := &goCodeRow{
		context: inst.context,
		pack:    inst.pack,
		sf:      inst.sf,
		source:  inst.source,
		module:  inst.context.Module,
		reader:  inst,
	}
	gcRow.init(rowNum, row)
	return h.Handle(gcRow)
}

// func (inst *GoSourceFileReader) parseRowAsStarterConfigenToken(row string, rowNum int) error {

// words := inst.parseWords(row)
// vlog.Info("", words)
// return nil
// }

// func (inst *GoSourceFileReader) parseRowAsGoCode(row string, rowNum int) error {

// pre-del

// words := inst.parseWords(row)

// if inst.hasPatternInWords(words, "package", "*") {
// 	vlog.Info("package_name_row: %s", row)

// } else if inst.hasPatternInWords(words, "import", "(") {
// 	vlog.Info("begin_of_import_list:", row)

// } else if inst.hasPatternInWords(words, "import", "*") {
// 	vlog.Info("single_import_row:", row)

// } else if inst.hasPatternInWords(words, "type", "*", "struct", "{") {
// 	vlog.Info("begin_of_type_struct_block:", row)

// } else if inst.hasPatternInWords(words, "func", "(", "*", PatternTokenHex, "*", ")", "*", "(", PatternTokenMore) {
// 	vlog.Info("begin_of_func_for_struct_block:", row)

// } else if inst.hasPatternInWords(words, ")") {
// 	vlog.Info("end_of_import_list:", row)

// } else if inst.hasPatternInWords(words, "}") {
// 	vlog.Info("end_of_block:", row)
// }

// 	return nil
// }

// parseRowAsGoCodeWithSCT like parseRowAsGoCode but with Starter-Configen-Token
// func (inst *GoSourceFileReader) parseRowAsGoCodeWithSCT(row string, rowNum int) error {
// 	words := inst.parseWords(row)

// 	if inst.hasPatternInWords(words, "*", "*", ".", "*", "/", "/", "starter", ":") {
// 		vlog.Info("inject_one_interface_row:", row)
// 	} else if inst.hasPatternInWords(words, "*", "[", "]", "*", ".", "*", "/", "/", "starter", ":") {
// 		vlog.Info("inject_interface_array_row:", row)
// 	}

// 	return nil
// }

////////////////////////////////////////////////////////////////////////////////

type goCodeRow struct {
	rowNumber int
	text      string
	words     []string

	context *v4.Context
	module  *v4.Module
	sf      *v4.SourceFolder
	pack    *v4.Package
	source  *v4.Source

	reader   *GoSourceFileReader
	fragment codeFragment // 用来区分同一个代码实体的不同部分
}

func (inst *goCodeRow) init(rowNum int, text string) {
	reader := &GoCodeWordsReader{}
	inst.words = reader.Read(text)
	inst.text = text
	inst.rowNumber = rowNum
}

func (inst *goCodeRow) hasPattern(pattern ...string) bool {
	words := inst.words
	sizeWords := len(words)
	sizePattern := len(pattern)
	if sizeWords < sizePattern {
		return false
	}
	for i := 0; i < sizePattern; i++ {
		str1 := words[i]
		str2 := pattern[i]
		if str2 == "*" {
			// continue
		} else if str1 == "*" && str2 == PatternTokenHex {
			// continue
		} else if str1 == str2 {
			// continue
		} else if str2 == PatternTokenMore {
			return true
		} else {
			return false
		}
	}
	return true
}

////////////////////////////////////////////////////////////////////////////////

type goCodeRowHandler interface {
	Accept(row *goCodeRow) bool
	Handle(row *goCodeRow) error
}

////////////////////////////////////////////////////////////////////////////////

type goCodeRowHandlerGroup struct {
	handlers []goCodeRowHandler
}

func (inst *goCodeRowHandlerGroup) _Impl() goCodeRowHandler {
	return inst
}

func (inst *goCodeRowHandlerGroup) add(h goCodeRowHandler) {
	if h == nil {
		return
	}
	inst.handlers = append(inst.handlers, h)
}

func (inst *goCodeRowHandlerGroup) Accept(row *goCodeRow) bool {
	return true
}

func (inst *goCodeRowHandlerGroup) Handle(row *goCodeRow) error {
	all := inst.handlers
	for _, h := range all {
		if h.Accept(row) {
			return h.Handle(row)
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
// 以下是各种具体类型的行处理器 //////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

type rowHandlerForPackage struct{}

func (inst *rowHandlerForPackage) Accept(row *goCodeRow) bool {
	return row.hasPattern("package", "*")
}

func (inst *rowHandlerForPackage) Handle(row *goCodeRow) error {

	p := row.pack
	simpleName := row.words[1]
	module := row.module
	sourceFile := row.source.Path

	p.SimpleName = simpleName
	p.Alias = simpleName
	p.Path = sourceFile.GetParent()
	p.OwnerModule = module
	p.FullName = inst.computePackageFullName(row)

	return nil
}

func (inst *rowHandlerForPackage) computePackageFullName(row *goCodeRow) string {

	path1 := row.module.Path
	path2 := row.pack.Path

	if path1.GetName() == "go.mod" {
		path1 = path1.GetParent()
	}

	str1 := path1.GetPath()
	elements := make([]string, 0)

	for p := path2; p != nil; p = p.GetParent() {
		str := p.GetPath()
		if str == str1 || len(str) < len(str1) {
			break
		}
		elements = append(elements, p.GetName())
	}

	builder := strings.Builder{}
	builder.WriteString(row.module.Name)

	for i := len(elements) - 1; i > 0; i-- {
		builder.WriteString("/")
		builder.WriteString(elements[i])
	}

	builder.WriteString("/")
	builder.WriteString(row.pack.SimpleName)
	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////

type rowHandlerForImport struct {
	theBlockBegin       codeFragment
	theBlockEnd         codeFragment
	theBlockInnerItem1  codeFragment // fullname only
	theBlockInnerItem2  codeFragment // alias + fullname
	theIndependentItem1 codeFragment // fullname only
	theIndependentItem2 codeFragment // alias + fullname
}

func (inst *rowHandlerForImport) init() {
	inst.theBlockBegin = "import("
	inst.theBlockInnerItem1 = "fullname"
	inst.theBlockInnerItem2 = "alias_fullname"
	inst.theBlockEnd = ")"
	inst.theIndependentItem1 = "import_fullname"
	inst.theIndependentItem2 = "import_alias_fullname"
}

func (inst *rowHandlerForImport) Accept(row *goCodeRow) bool {
	inst.init()

	inBlock := row.reader.inImportBlock

	if inBlock {
		if row.hasPattern(")") {
			row.fragment = inst.theBlockEnd
			row.reader.inImportBlock = false
		} else if row.hasPattern("*") {
			row.fragment = inst.theBlockInnerItem1
		} else if row.hasPattern("*", "*") {
			row.fragment = inst.theBlockInnerItem2
		} else {
			return false
		}
	} else {
		if row.hasPattern("import", "(") {
			row.fragment = inst.theBlockBegin
			row.reader.inImportBlock = true
		} else if row.hasPattern("import", "*", "*") {
			row.fragment = inst.theIndependentItem2
		} else if row.hasPattern("import", "*") {
			row.fragment = inst.theIndependentItem1
		} else {
			return false
		}
	}

	return true
}

func (inst *rowHandlerForImport) Handle(row *goCodeRow) error {

	fragment := row.fragment
	item := &v4.CodeImport{}
	words := row.words

	if fragment == inst.theBlockBegin {
		return nil

	} else if fragment == inst.theBlockEnd {
		return nil

	} else if fragment == inst.theBlockInnerItem1 {
		item.FullName = words[0]

	} else if fragment == inst.theBlockInnerItem2 {
		item.Alias = words[0]
		item.FullName = words[1]

	} else if fragment == inst.theIndependentItem1 {
		item.FullName = words[1]

	} else if fragment == inst.theIndependentItem2 {
		item.Alias = words[1]
		item.FullName = words[2]
	}

	if item.FullName != "" {
		if item.Alias == "" {
			item.Alias = inst.getAliasByFullname(item.FullName)
		}
		table := row.reader.importItems
		table[item.Alias] = item
		table[item.FullName] = item
	}

	return nil // todo ...
}

func (inst *rowHandlerForImport) getAliasByFullname(fullname string) string {
	i := strings.LastIndex(fullname, "/")
	if i < 0 {
		return fullname
	}
	alias := fullname[i+1:]
	return strings.TrimSpace(alias)
}

////////////////////////////////////////////////////////////////////////////////

type rowHandlerForTypeStruct struct{}

func (inst *rowHandlerForTypeStruct) Accept(row *goCodeRow) bool {
	return false // todo ...
}

func (inst *rowHandlerForTypeStruct) Handle(row *goCodeRow) error {
	return nil // todo ...
}

////////////////////////////////////////////////////////////////////////////////

type rowHandlerForFuncOfStruct struct{}

func (inst *rowHandlerForFuncOfStruct) Accept(row *goCodeRow) bool {
	return false // todo ...
}

func (inst *rowHandlerForFuncOfStruct) Handle(row *goCodeRow) error {
	return nil // todo ...
}

////////////////////////////////////////////////////////////////////////////////

type rowHandlerForStarterComponent struct{}

func (inst *rowHandlerForStarterComponent) Accept(row *goCodeRow) bool {
	return false // todo ...
}

func (inst *rowHandlerForStarterComponent) Handle(row *goCodeRow) error {
	return nil // todo ...
}

////////////////////////////////////////////////////////////////////////////////

type rowHandlerForStarterInject struct{}

func (inst *rowHandlerForStarterInject) Accept(row *goCodeRow) bool {
	return false // todo ...
}

func (inst *rowHandlerForStarterInject) Handle(row *goCodeRow) error {
	return nil // todo ...
}

////////////////////////////////////////////////////////////////////////////////

type rowHandlerForStarterInterface struct{}

func (inst *rowHandlerForStarterInterface) Accept(row *goCodeRow) bool {
	return false // todo ...
}

func (inst *rowHandlerForStarterInterface) Handle(row *goCodeRow) error {
	return nil // todo ...
}

////////////////////////////////////////////////////////////////////////////////
