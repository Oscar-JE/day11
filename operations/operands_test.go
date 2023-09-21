package operations

import (
	"testing"
)

func TestOperations(t *testing.T) {
	addition := Add{K: 2}
	multi := Multiply{K: 10}
	var opr1 Operation = &addition
	var opr2 Operation = &multi
	res1 := opr1.Opr(1)
	exp1 := 3
	res2 := opr2.Opr(2)
	exp2 := 20
	if res1 != exp1 {
		t.Errorf("Expexted : %d , was : %d ", exp1, res1)
	}
	if res2 != exp2 {
		t.Errorf("Expexted : %d , was : %d ", exp1, res1)
	}
}

func TestDivisibilityMany(t *testing.T) {
	denominators := []int{0, 1, 2, 3, 4, 5, 6, 7}
	enumerators := []int{2, 2, 1, 6, 7, 8, 12, 50}
	expected := []bool{false, true, false, true, false, false, true, false}
	for index, _ := range expected {
		res := FirstDevideSecond(denominators[index], enumerators[index])
		exp := expected[index]
		assertion(exp, res, index, t)
	}
}

func TestOneDivisibility(t *testing.T) {
	denominator := 0
	enumerator := 1
	res := FirstDevideSecond(denominator, enumerator)
	assertion(false, res, 51, t)
}

func TestDivisibilityCase2(t *testing.T) {
	denominator := 1
	enumerator := 2
	res := FirstDevideSecond(denominator, enumerator)
	assertion(true, res, 51, t)
}

func assertion(expect bool, res bool, index int, t *testing.T) {
	if expect != res {
		t.Errorf("error at index %d", index)
	}
}

func TestMultiplyLargestAndSecondLargest(t *testing.T) {
	values := []int{101, 95, 7, 105}
	res := MultiplyLargestAndSecondLargest(values)
	if res != 10605 {
		t.Errorf("fail in Multiply largest and second largest")
	}
}
