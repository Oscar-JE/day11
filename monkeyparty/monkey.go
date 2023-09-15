package monkeyparty

type Monkey struct {
	items           []int
	oper            Operation
	divider         int
	positivResiver  *Monkey
	negativeResiver *Monkey
}
