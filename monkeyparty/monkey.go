package monkeyparty

import "fmt"

type Monkey struct {
	items            []int
	oper             Operation
	divider          int
	positivReceiver  *Monkey
	negativeReceiver *Monkey
	nrOfInspects     int
}

func InitMonkey(items []int, opr Operation, devider int) Monkey {
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
	return firstDevideSecond(m.divider, item)
}

func sendTo(other *Monkey, item int) {
	other.items = append(other.items, item)

}

func (m *Monkey) Monkeybusiness() { // frågan är om man ska stadga någonting för att inte appor ska kasta flera saker per lop ?
	nrItem := len(m.items)
	i := 0
	for i < nrItem {
		activeItem := m.items[0]
		m.items = m.items[1:] // varför fungerar inte denna ??
		activeItem = m.oper.Opr(activeItem) / 3
		m.nrOfInspects++
		if firstDevideSecond(m.divider, activeItem) {
			sendTo(m.positivReceiver, activeItem)
		} else {
			sendTo(m.negativeReceiver, activeItem)
		}
		i++
	}
}

func (m Monkey) String() string {
	return fmt.Sprintf("%d", m.items)
}
