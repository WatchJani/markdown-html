package hash

import (
	"root/syntax"
)

type Hash struct {
	hash []syntax.Code
}

func NewHash(capacity int) *Hash {
	return &Hash{
		hash: make([]syntax.Code, capacity),
	}
}

func (h *Hash) Append(key byte, value syntax.Code) {
	h.hash[hashFn(key)] = value
}

func (h Hash) GetValue(value byte) (syntax.Code, bool) {
	val := h.hash[hashFn(value)]
	return val, val.Code != 0
}

func hashFn(key byte) int8 {
	return int8(key - 1&0b11111)
}
