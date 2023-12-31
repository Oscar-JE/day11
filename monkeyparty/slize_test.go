package monkeyparty

import (
	"testing"
)

func setToTen(s *[]int) {
	(*s)[0] = 10
}

func crop(s *[]int) {
	(*s) = (*s)[1:]
}

func TestModifySlizeInFunction(t *testing.T) {
	mySlize := []int{1, 2, 3}
	setToTen(&mySlize)
	if mySlize[0] != 10 {
		t.Errorf("wtf1 ")
	}
	crop(&mySlize)
	if len(mySlize) != 2 {
		t.Errorf("wtf2")
	}
}

type Wraper struct {
	S []int
}

func (w *Wraper) setToEleven() {
	(w.S)[0] = 11 // förstår inte hur detta fungerar
}

func (w *Wraper) cropInStruct() {
	(w.S) = (w.S)[1:]
}

func TestModifySlizeInStructPassedToFunction(t *testing.T) {
	w := Wraper{[]int{1, 2, 3}}
	w.setToEleven()
	if (w.S)[0] != 11 {
		t.Errorf("first element not changed")
	}
	w.cropInStruct()
	if len(w.S) != 2 {
		t.Errorf("internal slize not cropped")
	}
}

type RefWraper struct {
	S *[]int
}

func (w *RefWraper) setToEleven() {
	(*w.S)[0] = 11 // förstår inte hur detta fungerar
}

func (w *RefWraper) cropInStruct() {
	(*w.S) = (*w.S)[1:]
}

func TestModifyWithSlizeRefInStructPassedToFunction(t *testing.T) {
	w := RefWraper{&[]int{1, 2, 3}}
	w.setToEleven()
	if (*w.S)[0] != 11 {
		t.Errorf("first element not changed")
	}
	w.cropInStruct()
	if len(*w.S) != 2 {
		t.Errorf("internal slize not cropped")
	}
}

func TestModifyInLoop(t *testing.T) {
	list := []Wraper{Wraper{[]int{1, 2, 3}}, Wraper{[]int{4, 5, 6}}}
	for index, _ := range list {
		list[index].cropInStruct()
	}
	for _, element := range list {
		if len(element.S) != 2 {
			t.Errorf("not all lengts in structs are 2")
		}
	}
}
