package markdown

import (
	"root/hash"
	"root/sbuff"
)

type Markdown struct {
	*hash.Hash
	*Syntax
}

func NewMarkdown() *Markdown {
	hashMap := hash.NewHash(128)

	hashMap.Append('*', 10794)
	hashMap.Append('`', 0) //specific case (dynamic size)
	hashMap.Append('[', 0) //specific case (dynamic size)
	hashMap.Append('_', 95)
	hashMap.Append('#', 8995)
	hashMap.Append('>', 15904)
	hashMap.Append('~', 32382)

	return &Markdown{
		Hash:   hashMap,
		Syntax: NewSyntax(),
	}
}

func (m *Markdown) ToHTML(data []byte) []byte {
	buffer := sbuff.NewSBuff(len(data) + len(data)/2)

	for i := 0; i < len(data); i++ {
		if code, ok := m.GetValue(data[i]); ok {
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

				buffer.Append(data[start:i])
				start = i
				// buffer.setHalf(buffer.pointer)

				if data[i+1] != '(' {
					buffer.Reset()
					continue
				}

				i += 2
				for data[i] != ')' && i < len(data) {
					buffer.Append(data[start:i])
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

	return buffer.GetBuff()
}

func load16(b []byte) uint16 {
	return uint16(b[1])<<8 | uint16(b[0])
}
