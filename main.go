package main

import (
	"fmt"
	"root/markdown"
)

func main() {
	md := []byte("##red## ~~more je fino~~  [super novi text, hvala](url)")
	markdown := markdown.NewMarkdown()

	fmt.Println(string(markdown.ToHTML(md)))

}

func url(description, url string) []byte {
	return []byte(fmt.Sprintf("<a href=\"%s\">%s</a>", url, description))
}
