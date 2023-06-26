package vo

import "github.com/starter-go/configen/v4/dto"

// Configen 用于解析 "configen.json" 文件
type Configen struct {
	Configen dto.Configen `json:"configen"`
}
