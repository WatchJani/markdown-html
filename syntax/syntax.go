package syntax

import (
	"errors"
	"fmt"
	"root/sbuff"
)

const (
	URL_LENGTH         = 100 //max url length for image and url HTML tag
	DESCRIPTION_LENGTH = 50  //max number of character for description url or img HTML tag
	MAX_TAG_LENGTH     = 16  //max length of HTML tag
)

var (
	MaxLengthURL         = errors.New(fmt.Sprintf("Max url length %d", URL_LENGTH))
	MaxLengthDescription = errors.New(fmt.Sprintf("Max description length %d", DESCRIPTION_LENGTH))
)

// Generate complex HTML tag (img, ulr)
// help us for parsing markdown to HTML
type Syntax struct {
	option [2][3]PiceLength
	code   []byte
	*sbuff.SBuff
	// sync.RWMutex
}

func NewSyntax() *Syntax {
	var option = [2][3]PiceLength{
		{NewPiceLength(0, 9), NewPiceLength(9, 11), NewPiceLength(11, 15)},    //url
		{NewPiceLength(15, 25), NewPiceLength(25, 32), NewPiceLength(32, 34)}, //img
	}

	return &Syntax{
		option: option,
		code:   []byte("<a href=\"\"><\\a><img src=\"\" alt=\"\">"),
		SBuff:  sbuff.NewSBuff(URL_LENGTH + DESCRIPTION_LENGTH + MAX_TAG_LENGTH), // not good for concurrency, data race problem
	}
}

func (s *Syntax) Parser(syntax int, url, description []byte) []byte {
	s.Append(s.code[s.option[syntax][0].Start:s.option[syntax][0].End])
	s.Append(url)
	s.Append(s.code[s.option[syntax][1].Start:s.option[syntax][1].End])
	s.Append(description)
	s.Append(s.code[s.option[syntax][2].Start:s.option[syntax][2].End])

	return s.SBuff.Read(0, 120)
}
