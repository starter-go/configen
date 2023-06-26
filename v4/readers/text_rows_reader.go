package readers

import (
	"strings"

	"github.com/starter-go/afs"
)

// ReadRows 读取文本文件中的每一行
func ReadRows(file afs.Path) ([]string, error) {

	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}

	const ch1 = "\r"
	const ch2 = "\n"
	text = strings.ReplaceAll(text, ch1, ch2)
	rows := strings.Split(text, ch2)

	for i := len(rows) - 1; i >= 0; i-- {
		row := rows[i]
		rows[i] = strings.TrimSpace(row)
	}

	return rows, nil
}
