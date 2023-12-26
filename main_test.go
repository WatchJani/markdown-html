package main

import (
	"root/markdown"
	"root/syntax"
	"testing"
)

// kill
func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url("ChatGPT", "https://github.com/WatchJani/bug-free-journey")
	}
}

func BenchmarkByte(b *testing.B) {
	syn := syntax.NewSyntax()

	for i := 0; i < b.N; i++ {
		syn.Parser(1, []byte("https://chat.openai.com/c/ce412eb4-51f4-41bb-a0a2-17e051db363b"), []byte("ChatGPT"))
	}
}

func BenchmarkParser(b *testing.B) {
	md := []byte("##red##  ~~more je fino~~ [super novi text, hvala](url) Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")

	markdown := markdown.NewMarkdown()

	for i := 0; i < b.N; i++ {
		markdown.ToHTML(md)
	}
}
