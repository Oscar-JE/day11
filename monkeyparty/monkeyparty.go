package monkeyparty

import (
	"day11/operations"
	"day11/parse"
	"fmt"
)

type MonkeyParty struct {
	monkeys    []Monkey
	normalizer int
}

func Init(infos []parse.InputInfo) MonkeyParty {
	monkeys := []Monkey{}
	n := 1
	for _, info := range infos {
		m := InitMonkey(info.Items, info.Operation, info.Devider)
		n = n * info.Devider
		monkeys = append(monkeys, m)
	}
	for index, info := range infos {
		monkeys[index].SetPositiveReceiver(&monkeys[info.PositiveReciverIndex])
		monkeys[index].SetNegativeReceiver(&monkeys[info.NegativeReciverindex])
	}
	return MonkeyParty{monkeys: monkeys, normalizer: n}
}

func (m *MonkeyParty) Party() {
	for index := range m.monkeys {
		m.monkeys[index].Shenanigans()
	}
}

func (m *MonkeyParty) Party2() {
	for index := range m.monkeys {
		m.monkeys[index].Shenanigans2(m.normalizer)
	}
}

func (m MonkeyParty) String() string {
	return fmt.Sprint(m.monkeys)
}

func (m MonkeyParty) GetInspecteds() []int {
	inspecteds := []int{}
	for _, mon := range m.monkeys {
		inspecteds = append(inspecteds, mon.nrOfInspects)
	}
	return inspecteds
}

func (m MonkeyParty) MonkeyBusiness() int {
	inpected := m.GetInspecteds()
	return operations.MultiplyLargestAndSecondLargest(inpected)
}
