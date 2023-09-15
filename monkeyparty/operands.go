package monkeyparty

type Operation interface {
	Opr(int) int
}

type Add struct {
	k int
}

func (a Add) Opr(worry int) int {
	return a.k + worry
}

type Multiply struct {
	k int
}

func (m Multiply) Opr(worry int) int {
	return m.k * worry
}

func fistDevideSecond(denominator int, numerator int) bool {
	if denominator == 0 || denominator > numerator {
		return false
	}
	if denominator == 1 {
		return true
	}
	var fraction int = numerator / denominator
	if fraction == 1 {
		return true
	} else {
		return fistDevideSecond(denominator, fraction)
	}
}
