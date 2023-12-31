package monkeyparty

import (
	"day11/operations"
	"fmt"
)

type Monkey struct {
	items            []int
	oper             operations.Operation
	divider          int
	positivReceiver  *Monkey
	negativeReceiver *Monkey
	nrOfInspects     int
}

func InitMonkey(items []int, opr operations.Operation, devider int) Monkey {
	return Monkey{items: items, oper: opr,
		divider: devider, positivReceiver: nil, negativeReceiver: nil, nrOfInspects: 0}
}

func (m *Monkey) SetPositiveReceiver(other *Monkey) {
	m.positivReceiver = other
}

func (m *Monkey) SetNegativeReceiver(other *Monkey) {
	m.negativeReceiver = other
}

func (m *Monkey) inspectItem() {
	aktiveItem := m.items[0]
	res := m.oper.Opr(aktiveItem)
	m.items[0] = res
}

func (m Monkey) preformeTest(item int) bool {
	return operations.FirstDevideSecond(m.divider, item)
}

func sendTo(other *Monkey, item int) {
	other.items = append(other.items, item)

}

func (m *Monkey) Shenanigans() { // frågan är om man ska stadga någonting för att inte appor ska kasta flera saker per lop ?
	nrItem := len(m.items)
	i := 0
	for i < nrItem {
		activeItem := m.items[i]
		activeItem = m.oper.Opr(activeItem) / 3
		m.nrOfInspects++
		if operations.FirstDevideSecond(m.divider, activeItem) {
			sendTo(m.positivReceiver, activeItem)
		} else {
			sendTo(m.negativeReceiver, activeItem)
		}
		i++
	}
	m.items = []int{}
}

func (m *Monkey) Shenanigans2(normalizer int) { // frågan är om man ska stadga någonting för att inte appor ska kasta flera saker per lop ?
	nrItem := len(m.items)
	i := 0
	for i < nrItem {
		activeItem := m.items[i]
		activeItem = m.oper.Opr(activeItem) % normalizer
		m.nrOfInspects++
		if operations.FirstDevideSecond(m.divider, activeItem) {
			sendTo(m.positivReceiver, activeItem)
		} else {
			sendTo(m.negativeReceiver, activeItem)
		}
		i++
	}
	m.items = []int{}
}

func (m Monkey) String() string {
	return fmt.Sprintf("%d", m.items)
}
