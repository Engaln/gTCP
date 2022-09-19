package model

type Serialize interface {
	Encode()
	Decode()
}

type Body struct {
	len uint16
}
