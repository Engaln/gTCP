package model

type Serialize interface {
	Encode()
	Decode()
}

type Body struct {
	len  uint16
	sign uint16
	data string
}

func (b *Body) Encode() {}

func (b *Body) Decode() {}
