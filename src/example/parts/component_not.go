package parts

// ComNot 这不是一个组件
type ComNot struct {
	A int //starter:inject("666")
	B bool
	C rune
	D []byte
	E string
	F float32
}
