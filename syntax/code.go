package syntax

type Code struct {
	Code int
	PiceLength
	IsOpen int //is tag open, if open then close them
}

func NewCode(code, min, max int) Code {
	return Code{
		Code:       code,
		PiceLength: NewPiceLength(min, max),
	}
}

func (c *Code) Update() {
	if c.IsOpen > 0 {
		c.IsOpen = 0
	}
	
	c.IsOpen = c.PiceLength.End
}
