package markdown

import (
	"fmt"
	"root/hash"
	"root/sbuff"
	"root/syntax"
)

type Markdown struct {
	*hash.Hash
	*syntax.Syntax
	HTML []byte
	find int
}

func NewMarkdown() *Markdown {
	hashMap := hash.NewHash(128)

	hashMap.Append('*', syntax.NewCode(10794, 43, 46)) //done
	hashMap.Append('`', syntax.NewCode(0, 7, 12))      //specific case (dynamic size)
	hashMap.Append('[', syntax.NewCode(0, 0, 0))       //specific case (dynamic size)
	hashMap.Append('_', syntax.NewCode(95, 16, 19))    //done
	hashMap.Append('#', syntax.NewCode(8995, 7, 11))   //done
	hashMap.Append('>', syntax.NewCode(15904, 26, 39)) //bug
	hashMap.Append('~', syntax.NewCode(32382, 50, 53)) //done

	return &Markdown{
		Hash:   hashMap,
		Syntax: syntax.NewSyntax(),
		HTML:   []byte("<p></p><h2></h2><i></i><code></code><q></q><b></b><s></s>"),
	}
}

func (m *Markdown) ToHTML(data []byte) []byte {
	buffer := sbuff.NewSBuff(len(data) + len(data)/2)
	m.find = 0 //for testing
	literal := 0

	for i := 0; i < len(data); i++ {

		if code, ok := m.GetValue(data[i]); ok {
			buffer.Append(data[literal:i])
			//img & url
			if data[i] == '[' {
				var c uint8 //img => 1 or url => 0
				var start int = i

				if i != 0 && data[i]-1 == '!' {
					c++
				}

				for data[i] != ']' && i < len(data) {
					i++
				}

				m.SBuff.Append(data[start:i])
				start = i
				// buffer.setHalf(buffer.pointer)
				if data[i+1] != '(' {
					buffer.Reset()
					continue
				}

				i += 2
				for data[i] != ')' && i < len(data) {
					m.SBuff.Append(data[start:i])
					i++
				}

				fmt.Println(string(m.SBuff.Read(0, 150)))

				buffer.Reset()
			} else if i <= len(data)-2 && load16(data[i:i+2]) == uint16(code.Code) {
				buffer.Append(m.HTML[code.Start+code.IsOpen : code.End+code.IsOpen])
				code.Update()
				i += 2
			}

			literal = i
		}
	}

	buffer.Append(data[literal:])

	return buffer.GetBuff()
}

func load16(b []byte) uint16 {
	return uint16(b[1])<<8 | uint16(b[0])
}
