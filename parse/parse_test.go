package parse

import (
	"testing"
)

func TestParseItem(t *testing.T) {
	input := "Starting items: 76, 66"
	out := parseItems(input)
	expected := []int{76, 66}
	for i := range expected {
		if out[i] != expected[i] {
			t.Errorf("wrong parsed items, was: %d , should be %d", out[i], expected[i])
		}
	}
}

func TestParseTest(t *testing.T) {
	input := "Test: divisible by 13"
	out := parseTest(input)
	expected := 13
	if out != expected {
		t.Errorf("wrong parsed testLins,was %d, should be %d", out, expected)
	}
}

func TestParseIndex(t *testing.T) {
	input := "If true: throw to monkey 7"
	out := parseIndex(input)
	expected := 7
	if out != expected {
		t.Errorf("wrong")
	}
}

func TestParseOperAdd(t *testing.T) {
	input := "Operation: new = old + 5"
	out := parseOperation(input)
	if out.Opr(2) != 7 {
		t.Errorf("Wrong operation returned")
	}
}

func TestParseOperAddOldOld(t *testing.T) {
	input := "Operation: new = old + old"
	out := parseOperation(input)
	if out.Opr(3) != 6 {
		t.Errorf("Wrong operation returned")
	}
}

func TestParseOperMult(t *testing.T) {
	input := "Operation: new = old * 11"
	out := parseOperation(input)
	if out.Opr(5) != 55 {
		t.Errorf("error")
	}
}

func TestParseOperMultOldOld(t *testing.T) {
	input := "Operation: new = old * old"
	out := parseOperation(input)
	if out.Opr(5) != 25 {
		t.Errorf("error")
	}
}

func TestParseEverything(t *testing.T) {
	file := "../input.txt"
	_, err := ParseInput(file)
	if err != nil {
		t.Errorf("error")
	}
}
