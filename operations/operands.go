package operations

type Operation interface {
	Opr(int) int
}

type Add struct {
	K int
}

func (add Add) Opr(worry int) int {
	return add.K + worry
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

func FirstDevideSecond(denominator int, numerator int) bool {
	if denominator == 0 {
		return false
	}
	modulus := numerator % denominator
	return modulus == 0
}

func MultiplyLargestAndSecondLargest(values []int) int {
	largest := 0
	nextLargest := 0
	for _, nr := range values {
		if nr > largest && nr > nextLargest {
			nextLargest = largest
			largest = nr
		} else if nr > nextLargest {
			nextLargest = nr
		}
	}
	return largest * nextLargest
}
