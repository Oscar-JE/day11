package monkeyparty

import (
	"day11/operations"
	"fmt"
	"testing"
)

func TestShortExample(t *testing.T) {
	monkeys := []Monkey{InitMonkey([]int{79, 98}, operations.Multiply{19}, 23),
		InitMonkey([]int{54, 65, 75, 74}, operations.Add{6}, 19),
		InitMonkey([]int{79, 60, 97}, operations.Square{}, 13),
		InitMonkey([]int{74}, operations.Add{3}, 17)}

	// settig upp throws
	monkeys[0].SetPositiveReceiver(&monkeys[2])
	monkeys[0].SetNegativeReceiver(&monkeys[3])
	monkeys[1].SetPositiveReceiver(&monkeys[2])
	monkeys[1].SetNegativeReceiver(&monkeys[0])
	monkeys[2].SetPositiveReceiver(&monkeys[1])
	monkeys[2].SetNegativeReceiver(&monkeys[3])
	monkeys[3].SetPositiveReceiver(&monkeys[0])
	monkeys[3].SetNegativeReceiver(&monkeys[1])

	for i := range monkeys {
		fmt.Println(monkeys[i])
	}
	for index := range monkeys {
		fmt.Printf("index : %d \n", index)
		//m := monkeys[index]
		monkeys[index].Shenanigans()
	}

	for i := range monkeys {
		fmt.Println(monkeys[i])
	}
}
