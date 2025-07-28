package katas

import "fmt"

type exam struct {
	name  string
	marks int
}

func (e exam) changeMarks() {
	e.marks = 99
}

func (e *exam) changeMarksPtr() {
	e.marks = 99
}

func methods() {
	e := exam{
		"Math",
		20,
	}
	fmt.Println(e.marks)
	e.changeMarks()
	fmt.Println(e.marks)
	e.changeMarksPtr()
	fmt.Println(e.marks)
}
