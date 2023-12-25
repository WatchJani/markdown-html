package main

import (
	"fmt"
)

func main() {
	text := []byte("##red##  ~~more je fino~~ [Chat GPT](https://chat.openai.com/c/10793460-973e-49ac-99f1-42a4af15a9da)")

	hashMap := NewHash(128)

	hashMap.Append('*', 10794)
	hashMap.Append('`', 0)
	hashMap.Append('[', 0)
	hashMap.Append('_', 95)
	hashMap.Append('#', 8995)
	hashMap.Append('>', 15904)
	hashMap.Append('~', 32382)

	Read(text, NewBuf(100), hashMap)

	fmt.Println(string(text))
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

type Hash struct {
	hash []int
}

func NewHash(capacity int) *Hash {
	return &Hash{
		hash: make([]int, capacity),
	}
}

func Read(data []byte, buffer *Buffer, hash *Hash) {
	for i := 0; i < len(data); i++ {
		if code, ok := hash.GetValue(data[i]); ok {

			//img & url
			if data[i] == '[' {
				var c uint8 //img => 1 or url => 0
				// var start int = i

				if i != 0 && data[i]-1 == '!' {
					c++
				}

				for data[i] != ']' && i < len(data) {
					buffer.Append(data[i])
					i++
				}

				buffer.setHalf(buffer.pointer)

				if data[i+1] != '(' {
					buffer.Reset()
					continue
				}

				i += 2
				for data[i] != ')' && i < len(data) {
					buffer.Append(data[i])
					i++
				}

				buffer.Reset()
				continue
			}

			//bold & heading & quote & strikeTrough
			if i <= len(data)-2 && load16(data[i:i+2]) == uint16(code) {
				continue
			}
		}
	}
}

func load16(b []byte) uint16 {
	return uint16(b[1])<<8 | uint16(b[0])
}

type Buffer struct {
	buf      []byte
	half     int
	pointer  int
	capacity int
}

func NewBuf(capacity int) *Buffer {
	return &Buffer{
		capacity: capacity,
		buf:      make([]byte, capacity),
	}
}

func (b *Buffer) Append(value byte) {
	if b.pointer == b.capacity {
		panic("to much character") //fix this error
	}

	b.buf[b.pointer] = value
	b.pointer++
}

func (b Buffer) Read() {
	fmt.Println(string(b.buf[1:b.pointer]))
}

func (b *Buffer) Reset() {
	b.pointer = 0
}

func (b *Buffer) setHalf(half int) {
	b.half = half
}

func (b Buffer) Img() {
	fmt.Println(string(b.buf[1:b.half]), "|", string(b.buf[b.half:]))
}
