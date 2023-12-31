package main

import (
	"day11/monkeyparty"
	"day11/operations"
	"day11/parse"
	"fmt"
)

func main() {

	//monkeyParty := createFromExample()
	monkeyParty := createFromFile("input.txt")

	fmt.Println(monkeyParty)
	monkeyParty.Party2()
	print("after one round\n")
	fmt.Println(monkeyParty.GetInspecteds())

	for i := 1; i < 20; i++ {
		monkeyParty.Party2()
	}

	print("after 20 round\n")
	print("inspected numbers\n")
	fmt.Println(monkeyParty.GetInspecteds())

	for i := 20; i < 10000; i++ {
		monkeyParty.Party2()
	}

	print("after 10000 round\n")
	fmt.Println(monkeyParty)
	print("inspected numbers\n")
	fmt.Println(monkeyParty.GetInspecteds())
	print("Monkeybuissines\n")
	fmt.Println(monkeyParty.MonkeyBusiness())

}

func createFromExample() monkeyparty.MonkeyParty {
	infos := []parse.InputInfo{
		{Items: []int{79, 98}, Operation: operations.Multiply{K: 19}, Devider: 23, PositiveReciverIndex: 2, NegativeReciverindex: 3},
		{Items: []int{54, 65, 75, 74}, Operation: operations.Add{K: 6}, Devider: 19, PositiveReciverIndex: 2, NegativeReciverindex: 0},
		{Items: []int{79, 60, 97}, Operation: operations.Square{}, Devider: 13, PositiveReciverIndex: 1, NegativeReciverindex: 3},
		{Items: []int{74}, Operation: operations.Add{K: 3}, Devider: 17, PositiveReciverIndex: 0, NegativeReciverindex: 1}}

	return monkeyparty.Init(infos)
}

func createFromFile(filepath string) monkeyparty.MonkeyParty {
	infos, err := parse.ParseInput(filepath)
	if err != nil {
		panic("Fileparse failed")
	}
	return monkeyparty.Init(infos)
}
