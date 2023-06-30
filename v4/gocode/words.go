package gocode

import "strings"

// 定义几个特殊的单词符号
const (
	WordAnyString = "*"     // 表示任意字符串
	WordAnyRune   = "?"     // 表示任意字符
	WordHex       = "[*]"   // 表示 '*' 符号
	WordMore      = "[...]" // 表示后缀为任意长度的任意内容
)

// ParseWords 把字符串解析为单词序列
func ParseWords(text string) *Words {
	r := &wordsReader{}
	return r.Read(text)
}

// NewWords 把字符串解析为单词序列
func NewWords(items []string) *Words {
	if items == nil {
		items = make([]string, 0)
	}
	return &Words{list: items}
}

////////////////////////////////////////////////////////////////////////////////

// Words 是一组由单词（或者符号）构成的序列
type Words struct {
	list []string
}

func (inst *Words) String() string {
	b := strings.Builder{}
	for _, item := range inst.list {
		b.WriteString(item)
	}
	return b.String()
}

// WordAt 通过索引取词
func (inst *Words) WordAt(index int, def string) string {
	list := inst.list
	size := len(list)
	if 0 <= index && index < size {
		return list[index]
	}
	return def
}

// HasPattern 判断序列是否符合指定的模式
func (inst *Words) HasPattern(pattern ...string) bool {

	items := inst.list
	sizeH := len(items)   // have:items
	sizeW := len(pattern) // want:pattern
	sizeMax := sizeH
	if sizeMax < sizeW {
		sizeMax = sizeW
	}

	for i := 0; i < sizeMax; i++ {
		have := ""
		want := ""

		if i < sizeH {
			have = items[i]
		} else {
			return false
		}

		if i < sizeW {
			want = pattern[i]
		} else {
			return false
		}

		if want == WordAnyString {
			// continue
		} else if want == WordHex && have == "*" {
			// continue
		} else if want == have {
			// continue
		} else if want == WordMore {
			return true
		} else {
			return false
		}
	}

	return true
}

// List 返回列表
func (inst *Words) List() []string {
	return inst.list
}

// Len 返回列表长度
func (inst *Words) Len() int {
	return len(inst.list)
}

////////////////////////////////////////////////////////////////////////////////

// wordsReader go 代码分词读取
type wordsReader struct {
}

func (inst *wordsReader) Read(row string) *Words {
	words := make([]string, 0)
	sb := &strings.Builder{}
	chs := []rune(row)
	inAString := false
	inAStringCloser := '`'
	for _, ch := range chs {
		// col := i + 1
		if inAString {
			if ch == inAStringCloser {
				// the string end
				inAString = false
				words = inst.appendWord(words, sb)
			} else {
				sb.WriteRune(ch)
			}
		} else {
			if inst.isWordChar(ch) {
				sb.WriteRune(ch)
			} else if inst.isSpaceChar(ch) {
				words = inst.appendWord(words, sb)
			} else if inst.isStringCloserChar(ch) {
				// the string begin
				inAString = true
				inAStringCloser = ch
				words = inst.appendWord(words, sb)
			} else /* other marks */ {
				words = inst.appendWord(words, sb)
				words = append(words, string(ch))
			}
		}
	}
	words = inst.appendWord(words, sb)
	return &Words{list: words}
}

func (inst *wordsReader) appendWord(words []string, sb *strings.Builder) []string {
	if sb.Len() > 0 {
		word := sb.String()
		words = append(words, word)
		sb.Reset()
	}
	return words
}

func (inst *wordsReader) isStringCloserChar(r rune) bool {
	return (r == '"') || (r == '\'') || (r == '`')
}

func (inst *wordsReader) isSpaceChar(r rune) bool {
	return (r == ' ') || (r == '\t')
}

func (inst *wordsReader) isWordChar(r rune) bool {
	if ('0' <= r) && (r <= '9') {
		return true
	} else if ('a' <= r) && (r <= 'z') {
		return true
	} else if ('A' <= r) && (r <= 'Z') {
		return true
	} else if r == '_' {
		return true
	}
	return false
}
