package readers

import (
	"fmt"
	"strings"

	"github.com/starter-go/afs"
	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/gocode"
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
	// context
	context *v4.Context
	sf      *v4.SourceFolder
	pack    *gocode.Package

	// handler-groups
	hGroupForStarterToken goCodeRowHandlerGroup
	hGroupForGoCode       goCodeRowHandlerGroup
	hGroupForGoCodeWithST goCodeRowHandlerGroup // go_code + starter_token
	hGroupForNone         goCodeRowHandlerGroup

	// status
	currentImportBlock *gocode.ImportSet
	currentTypeStruct  *gocode.TypeStruct

	// result
	result *gocode.Source
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
	inst.pack = nil
	inst.result = nil

	inst.initGroupStarterToken()
	inst.initGroupGoCode()
	inst.initGroupGoCodeWithST()
}

func (inst *GoSourceFileReader) initGroupStarterToken() {
	group := &inst.hGroupForStarterToken
	group.add(&rowHandlerForStarterComponent{})
	group.add(&rowHandlerForStarterInject{})
	group.Init()
}

func (inst *GoSourceFileReader) initGroupGoCode() {
	group := &inst.hGroupForGoCode
	group.add(&rowHandlerForPackage{})
	group.add(&rowHandlerForImport{})
	group.add(&rowHandlerForTypeStruct{})
	// group.add(&rowHandlerForFuncOfStruct{})
	group.Init()
}

func (inst *GoSourceFileReader) initGroupGoCodeWithST() {
	group := &inst.hGroupForGoCodeWithST
	group.add(&rowHandlerForStarterInject{})
	group.add(&rowHandlerForStarterAs{})
	group.Init()
}

func (inst *GoSourceFileReader) Read(file afs.Path) (*gocode.Source, error) {

	pack := &gocode.Package{}
	result := &gocode.Source{}

	inst.pack = pack
	inst.result = result

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
		source:  inst.result,
		module:  inst.context.Module,
		reader:  inst,
	}
	gcRow.init(rowNum, row)
	return h.Handle(gcRow)
}

////////////////////////////////////////////////////////////////////////////////

type goCodeRow struct {
	rowNumber int
	text      string
	words     gocode.Words

	context *v4.Context
	sf      *v4.SourceFolder

	module *gocode.Module
	pack   *gocode.Package
	source *gocode.Source

	reader   *GoSourceFileReader
	fragment codeFragment // 用来区分同一个代码实体的不同部分
}

func (inst *goCodeRow) init(rowNum int, text string) {

	words := gocode.ParseWords(text)

	inst.words = *words
	inst.text = text
	inst.rowNumber = rowNum
}

func (inst *goCodeRow) wordAt(index int, def string) string {
	return inst.words.WordAt(index, def)
}

func (inst *goCodeRow) hasPattern(pattern ...string) bool {
	return inst.words.HasPattern(pattern...)
}

////////////////////////////////////////////////////////////////////////////////

type goCodeRowHandler interface {
	Init()
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

func (inst *goCodeRowHandlerGroup) Init() {
	list := inst.handlers
	for _, h := range list {
		h.Init()
	}
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

func (inst *rowHandlerForPackage) Init() {
}

func (inst *rowHandlerForPackage) Accept(row *goCodeRow) bool {
	return row.hasPattern("package", "*")
}

func (inst *rowHandlerForPackage) Handle(row *goCodeRow) error {

	p := row.pack
	simpleName := row.words.WordAt(1, "")
	module := row.module
	sourceFile := row.source.Path

	p.SimpleName = simpleName
	p.Alias = simpleName
	p.Path = sourceFile.GetParent()
	p.OwnerModule = module
	p.FullName = inst.computePackageFullName(row)

	inst.registerSelf(row)

	return nil
}

func (inst *rowHandlerForPackage) registerSelf(row *goCodeRow) {
	i := &row.source.ImportSet
	p := row.source.OwnerPackage
	fullname := p.FullName
	i.Add(&gocode.Import{
		Alias:    "",
		FullName: fullname,
	})
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
	theBlockInnerItem1  codeFragment // "fullname"
	theBlockInnerItem2  codeFragment // alias  "fullname"
	theIndependentItem1 codeFragment // import "fullname"
	theIndependentItem2 codeFragment // import alias "fullname"
}

func (inst *rowHandlerForImport) Init() {
	inst.theBlockBegin = "import("
	inst.theBlockInnerItem1 = "fullname"
	inst.theBlockInnerItem2 = "alias_fullname"
	inst.theBlockEnd = ")"
	inst.theIndependentItem1 = "import_fullname"
	inst.theIndependentItem2 = "import_alias_fullname"
}

func (inst *rowHandlerForImport) Accept(row *goCodeRow) bool {

	block := row.reader.currentImportBlock

	if block != nil {
		if row.hasPattern(")") {
			row.fragment = inst.theBlockEnd
			return true
		} else if row.hasPattern("*", "*") {
			row.fragment = inst.theBlockInnerItem2
			return true
		} else if row.hasPattern("*") {
			row.fragment = inst.theBlockInnerItem1
			return true
		}
	} else {
		if row.hasPattern("import", "(") {
			row.fragment = inst.theBlockBegin
			return true
		} else if row.hasPattern("import", "*", "*") {
			row.fragment = inst.theIndependentItem2
			return true
		} else if row.hasPattern("import", "*") {
			row.fragment = inst.theIndependentItem1
			return true
		}
	}
	return false
}

func (inst *rowHandlerForImport) Handle(row *goCodeRow) error {

	fragment := row.fragment
	item := &gocode.Import{}
	words := row.words

	if fragment == inst.theBlockBegin {
		row.reader.currentImportBlock = &gocode.ImportSet{}
		return nil

	} else if fragment == inst.theBlockEnd {
		row.reader.currentImportBlock = nil
		return nil

	} else if fragment == inst.theBlockInnerItem1 {
		item.FullName = words.WordAt(0, "")

	} else if fragment == inst.theBlockInnerItem2 {
		item.Alias = words.WordAt(0, "")
		item.FullName = words.WordAt(1, "")

	} else if fragment == inst.theIndependentItem1 {
		item.FullName = words.WordAt(1, "")

	} else if fragment == inst.theIndependentItem2 {
		item.Alias = words.WordAt(1, "")
		item.FullName = words.WordAt(2, "")
	} else {
		return nil
	}

	if item.FullName != "" {
		if item.Alias == "" {
			item.Alias = inst.getAliasByFullname(item.FullName)
		}
		row.reader.result.ImportSet.Add(item)
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

type rowHandlerForTypeStruct struct {
	theBlockBegin codeFragment
	theBlockEnd   codeFragment
}

func (inst *rowHandlerForTypeStruct) Init() {
	inst.theBlockBegin = "type name struct{"
	inst.theBlockEnd = "}"
}

func (inst *rowHandlerForTypeStruct) Accept(row *goCodeRow) bool {

	ts := row.reader.currentTypeStruct
	if ts == nil {
		if row.hasPattern("type", "*", "struct", "{") {
			row.fragment = inst.theBlockBegin
			return true
		}
	} else {
		if row.wordAt(0, "") == "}" {
			row.fragment = inst.theBlockEnd
			return true
		}
	}
	return false
}

func (inst *rowHandlerForTypeStruct) Handle(row *goCodeRow) error {

	fragment := row.fragment

	if fragment == inst.theBlockBegin {
		return inst.handleBlockBegin(row)

	} else if fragment == inst.theBlockEnd {
		return inst.handleBlockEnd(row)
	}

	return nil // todo ...
}

func (inst *rowHandlerForTypeStruct) handleBlockBegin(row *goCodeRow) error {
	block := &gocode.TypeStruct{}
	block.Name = row.wordAt(1, "")
	row.reader.currentTypeStruct = block
	return nil
}

func (inst *rowHandlerForTypeStruct) handleBlockEnd(row *goCodeRow) error {
	item := row.reader.currentTypeStruct
	row.reader.currentTypeStruct = nil
	row.source.TypeStructSet.Add(item)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// type rowHandlerForFuncOfStruct struct{}

// func (inst *rowHandlerForFuncOfStruct) Init() {
// }

// func (inst *rowHandlerForFuncOfStruct) Accept(row *goCodeRow) bool {
// 	return false // todo ...
// }

// func (inst *rowHandlerForFuncOfStruct) Handle(row *goCodeRow) error {
// 	return nil // todo ...
// }

////////////////////////////////////////////////////////////////////////////////

type rowHandlerForStarterComponent struct {
	prefix string
}

func (inst *rowHandlerForStarterComponent) Init() {
	inst.prefix = "//starter:component"
}

func (inst *rowHandlerForStarterComponent) Accept(row *goCodeRow) bool {
	return strings.HasPrefix(row.text, inst.prefix)
}

func (inst *rowHandlerForStarterComponent) Handle(row *goCodeRow) error {
	prefix := inst.prefix
	text := row.text
	ts := row.reader.currentTypeStruct
	errBadCom := fmt.Errorf("bad starter component: %s", text)
	if ts == nil {
		return errBadCom
	}
	if !strings.HasPrefix(text, prefix) {
		return errBadCom
	}
	part2 := strings.TrimSpace(text[len(prefix):])
	if part2 == "" {
		// ok
	} else if strings.HasPrefix(part2, "(") && strings.HasSuffix(part2, ")") {
		atts, err := gocode.ParseConfigenParams(part2)
		if err != nil {
			return err
		}
		ts.ComAtts = atts
	} else {
		return errBadCom
	}
	ts.IsComponent = true
	inst.loadComAtts(ts)
	return nil
}

func (inst *rowHandlerForStarterComponent) loadComAtts(ts *gocode.TypeStruct) {

	if ts == nil {
		return
	}
	atts := ts.ComAtts
	if atts == nil {
		return
	}

	all := atts.GetItems()
	table := make(map[string]*gocode.ConfigenParam)
	for _, item := range all {
		table[item.Name] = item
	}

	id := inst.getAttr(table, "id")
	if id == "" {
		id = inst.getAttr(table, "")
	}

	ts.ComID = id
	ts.ComClass = inst.getAttr(table, "class")
	ts.ComAlias = inst.getAttr(table, "alias")
	ts.ComScope = inst.getAttr(table, "scope")
}

func (inst *rowHandlerForStarterComponent) getAttr(table map[string]*gocode.ConfigenParam, name string) string {
	item := table[name]
	if item == nil {
		return ""
	}
	return item.Value
}

////////////////////////////////////////////////////////////////////////////////

type rowHandlerForStarterInject struct{}

func (inst *rowHandlerForStarterInject) Init() {
}

func (inst *rowHandlerForStarterInject) Accept(row *goCodeRow) bool {
	const keyword = "//starter:inject"
	return strings.Contains(row.text, keyword)
}

func (inst *rowHandlerForStarterInject) Handle(row *goCodeRow) error {
	// const keyword = "//starter:inject("xxxxx")"

	p1, _, p3, err := inst.getPart123(row)
	if err != nil {
		return err
	}

	f1, err := inst.parsePart1(row, p1)
	if err != nil {
		return err
	}

	f3, err := inst.parsePart3(row, p3)
	if err != nil {
		return err
	}

	field := f1
	field.Name = row.wordAt(0, "")
	field.Injection = f3.Injection

	ts := row.reader.currentTypeStruct
	ts.Fields.Add(field)

	return nil // todo ...
}

//  func ( inst *rowHandlerForStarterInject ) loadTypeInfo (elements *gocode.Words ,  * ImportSet) ComplexType {

//  }

func (inst *rowHandlerForStarterInject) parsePart1(row *goCodeRow, part1 []string) (*gocode.Field, error) {

	size := len(part1)
	if size < 2 {
		return nil, fmt.Errorf("bad starter:inject row: %s", row.text)
	}

	elements := gocode.NewWords(part1[1:])
	typeInfo, err := gocode.CreateComplexType(elements, &row.source.ImportSet)
	if err != nil {
		return nil, fmt.Errorf("bad starter:inject row: %s", row.text)
	}

	f := &gocode.Field{}
	f.Name = part1[0]
	f.Type = *typeInfo
	return f, nil
}

func (inst *rowHandlerForStarterInject) parsePart3(row *goCodeRow, part3 []string) (*gocode.Field, error) {
	elements := gocode.NewWords(part3)
	if elements.HasPattern("(", "*", ")") {
		f := &gocode.Field{}
		f.Injection = elements.WordAt(1, "#")
		return f, nil
	}
	return nil, fmt.Errorf("bad starter:inject row")
}

func (inst *rowHandlerForStarterInject) getPart123(row *goCodeRow) ([]string, string, []string, error) {

	const token = "//starter:inject"
	words := row.words.List()
	list1 := make([]string, 0)  // part-1
	list3 := make([]string, 0)  // part-3
	part2 := ""                 // part-2
	part2b := strings.Builder{} // part-2-builder
	part := 1                   // ={1,2,3}

	for _, w := range words {
		switch part {
		case 1:
			if w == "/" {
				part2b.WriteString(w)
				part++
			} else {
				list1 = append(list1, w)
			}
			break
		case 2:
			part2b.WriteString(w)
			part2 = part2b.String()
			if part2 == token {
				part++
			}
			break
		case 3:
			list3 = append(list3, w)
			break
		}
	}

	if part2 != token {
		const format = "bad starter:inject row:%s"
		return nil, "", nil, fmt.Errorf(format, row.text)
	}

	return list1, part2, list3, nil
}

////////////////////////////////////////////////////////////////////////////////

type rowHandlerForStarterAs struct {
	keyword string
}

func (inst *rowHandlerForStarterAs) Init() {
	inst.keyword = "//starter:as"
}

func (inst *rowHandlerForStarterAs) Accept(row *goCodeRow) bool {
	return strings.Contains(row.text, inst.keyword)
}

func (inst *rowHandlerForStarterAs) Handle(row *goCodeRow) error {
	// name func(t1,t2,t3) //starter:as("#",".","#.")
	kw := inst.keyword
	elements := row.words.List()
	builder := strings.Builder{}
	part1 := make([]string, 0)
	part2 := make([]string, 0)
	part3 := make([]string, 0)
	ipart := 1

	for _, el := range elements {
		switch ipart {
		case 1:
			if el == "/" {
				ipart++
				part2 = append(part2, el)
				builder.WriteString(el)
			} else {
				part1 = append(part1, el)
			}
			break
		case 2:
			part2 = append(part2, el)
			builder.WriteString(el)
			if builder.String() == kw {
				ipart++
			}
			break
		case 3:
			part3 = append(part3, el)
			break
		}
	}

	part1words := gocode.NewWords(part1)
	params, err := gocode.CreateConfigenParams(gocode.NewWords(part3))
	if err != nil {
		return err
	}
	return inst.parseStarterAsList(row, part1words, params)
}

func (inst *rowHandlerForStarterAs) parseStarterAsList(row *goCodeRow, fn *gocode.Words, params *gocode.ConfigenParams) error {

	// check
	size := fn.Len()
	iend := size - 1
	ibegin := 2
	name := fn.WordAt(0, "")
	tFunc := fn.WordAt(1, "")       // 'func'
	tBegin := fn.WordAt(ibegin, "") // '('
	tEnd := fn.WordAt(iend, "")     // ')'
	if tFunc != "func" || tBegin != "(" || tEnd != ")" {
		return fmt.Errorf("bad starter:as func: %s", name)
	}

	fragment := make([]string, 0)
	iParam := 0

	for i := ibegin + 1; i <= iend; i++ {
		word := fn.WordAt(i, "")
		if word == "," || word == ")" {
			param := params.GetItemAt(iParam)
			err := inst.parseStarterAsItem(row, fragment, param)
			if err != nil {
				return err
			}
			fragment = make([]string, 0)
			iParam++
		} else {
			fragment = append(fragment, word)
		}
	}

	return nil
}

func (inst *rowHandlerForStarterAs) parseStarterAsItem(row *goCodeRow, fragment []string, param *gocode.ConfigenParam) error {

	fragmentWords := gocode.NewWords(fragment)
	imports := &row.source.ImportSet
	ct, err := gocode.CreateComplexType(fragmentWords, imports)
	if err != nil {
		return err
	}

	impl := &gocode.Implementation{}
	impl.Type = *ct
	if param != nil {
		impl.Injection = param.Value
	}

	ts := row.reader.currentTypeStruct
	ts.As.Add(impl)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
