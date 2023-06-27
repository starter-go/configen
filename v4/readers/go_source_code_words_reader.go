package readers

import "strings"

// GoCodeWordsReader go 代码分词读取
type GoCodeWordsReader struct {
}

func (inst *GoCodeWordsReader) Read(row string) []string {
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
	return words
}

func (inst *GoCodeWordsReader) appendWord(words []string, sb *strings.Builder) []string {
	if sb.Len() > 0 {
		word := sb.String()
		words = append(words, word)
		sb.Reset()
	}
	return words
}

func (inst *GoCodeWordsReader) isStringCloserChar(r rune) bool {
	return (r == '"') || (r == '\'') || (r == '`')
}

func (inst *GoCodeWordsReader) isSpaceChar(r rune) bool {
	return (r == ' ') || (r == '\t')
}

func (inst *GoCodeWordsReader) isWordChar(r rune) bool {
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
