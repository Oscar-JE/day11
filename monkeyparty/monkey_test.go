package monkeyparty

import (
	"fmt"
	"testing"
)

func TestShortExample(t *testing.T) {
	monkeys := []Monkey{InitMonkey([]int{79, 98}, Multiply{19}, 23),
		InitMonkey([]int{54, 65, 75, 74}, Add{6}, 19),
		InitMonkey([]int{79, 60, 97}, Square{}, 13),
		InitMonkey([]int{74}, Add{3}, 17)}

	// settig upp throws
	monkeys[0].SetPositiveReceiver(&monkeys[2])
	monkeys[0].SetNegativeReceiver(&monkeys[3])
	monkeys[1].SetPositiveReceiver(&monkeys[2])
	monkeys[1].SetNegativeReceiver(&monkeys[0])
	monkeys[2].SetPositiveReceiver(&monkeys[1])
	monkeys[2].SetNegativeReceiver(&monkeys[3])
	monkeys[3].SetPositiveReceiver(&monkeys[0])
	monkeys[3].SetNegativeReceiver(&monkeys[1])

	for _, m := range monkeys {
		fmt.Println(m)
	}
	for _, m := range monkeys {
		m.Monkeybusiness()
		m.items = []int{}
	}

	for _, m := range monkeys {
		fmt.Println(m)
	}
}
