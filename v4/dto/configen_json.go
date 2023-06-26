package dto

// SourceID ...
type SourceID string

// DestinationID ...
type DestinationID string

// Source 表示源文件夹
type Source struct {
	ID        SourceID `json:"id"`
	Path      string   `json:"path"`
	Recursive bool     `json:"r"` // 是否递归
}

// Destination 表示目标文件夹
type Destination struct {
	ID      DestinationID `json:"id"`
	Path    string        `json:"path"`
	Sources []SourceID    `json:"sources"`
}

// Configen  表示 ConfigenJSON 内部的 root 对象
type Configen struct {
	Version      string         `json:"version"`
	Destinations []*Destination `json:"destinations"`
	Sources      []*Source      `json:"sources"`
}
