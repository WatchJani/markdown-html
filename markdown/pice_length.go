package markdown

type PiceLength struct {
	start int
	end   int
}

func NewPiceLength(start, end int) PiceLength {
	return PiceLength{
		start: start,
		end:   end,
	}
}
