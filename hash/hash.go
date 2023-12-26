package hash

type Hash struct {
	hash []int
}

func NewHash(capacity int) *Hash {
	return &Hash{
		hash: make([]int, capacity),
	}
}

func (h *Hash) Append(key byte, value int) {
	h.hash[hashFn(key)] = value
}

func (h Hash) GetValue(value byte) (int, bool) {
	val := h.hash[hashFn(value)]
	return val, val != 0
}

func hashFn(key byte) int8 {
	return int8(key - 1&0b11111)
}
