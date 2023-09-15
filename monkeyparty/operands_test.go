package monkeyparty

import (
	"testing"
)

func TestOperations(t *testing.T) {
	addition := Add{k: 2}
	multi := Multiply{k: 10}
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

func TestDivisibility(t *testing.T) {
	denominators := []int{0, 1, 2, 3, 4, 5, 6, 7}
	enumerators := []int{2, 2, 1, 6, 7, 8, 12, 50}
	expected := []bool{false, true, false, true, false, false, true, false}
	for index, _ := range expected {
		res := fistDevideSecond(denominators[index], enumerators[index])
		exp := expected[index]
		assertion(exp, res, index, t)
	}
}

func assertion(expect bool, res bool, index int, t *testing.T) {
	if expect != res {
		t.Errorf("error at index %d", index)
	}
}
