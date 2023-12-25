package main

import (
	"testing"
)

func Benchmark(b *testing.B) {
	text := []byte("##red##  ~~more je fino~~ [super novi text, hvala](url) Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")

	hashMap := NewHash(128)
	hashMap.Append('*', 1)
	hashMap.Append('`', 2)
	hashMap.Append('[', 3)
	hashMap.Append('_', 4)
	hashMap.Append('#', 5)
	hashMap.Append('>', 6)
	hashMap.Append('!', 7)

	buf := NewBuf(100)

	for i := 0; i < b.N; i++ {
		Read(text, buf, hashMap)
	}
}
