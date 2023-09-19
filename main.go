package main

import (
	"day11/monkeyparty"
	"fmt"
)

func main() {
	monkeys := []monkeyparty.Monkey{monkeyparty.InitMonkey([]int{79, 98}, monkeyparty.Multiply{K: 19}, 23),
		monkeyparty.InitMonkey([]int{54, 65, 75, 74}, monkeyparty.Add{K: 6}, 19),
		monkeyparty.InitMonkey([]int{79, 60, 97}, monkeyparty.Square{}, 13),
		monkeyparty.InitMonkey([]int{74}, monkeyparty.Add{K: 3}, 17)}

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

	party(&monkeys)

	for i := range monkeys {
		fmt.Println(monkeys[i])
	}
}

func party(party *[]monkeyparty.Monkey) {
	for index := range *party {
		fmt.Printf("index : %d \n", index)
		//m := monkeys[index]
		(*party)[index].Monkeybusiness()
	}
}
