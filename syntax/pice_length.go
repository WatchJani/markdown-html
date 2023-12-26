package syntax

type PiceLength struct {
	Start int
	End   int
}

func NewPiceLength(start, end int) PiceLength {
	return PiceLength{
		Start: start,
		End:   end,
	}
}
