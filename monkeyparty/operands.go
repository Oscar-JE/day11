package monkeyparty

type Operation interface {
	Opr(int) int
}

type Add struct {
	K int
}

func (a Add) Opr(worry int) int {
	return a.K + worry
}

type Multiply struct {
	K int
}

func (m Multiply) Opr(worry int) int {
	return m.K * worry
}

type Square struct {
}

func (s Square) Opr(worry int) int {
	return worry * worry
}

func firstDevideSecond(denominator int, numerator int) bool {
	if denominator == 0 {
		return false
	}
	modulus := numerator % denominator
	return modulus == 0
}
