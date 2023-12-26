package main

import (
	"fmt"
)

func main() {

}

func url(description, url string) []byte {
	return []byte(fmt.Sprintf("<a href=\"%s\">%s</a>", url, description))
}
