package katas

import (
	"fmt"
	"iter"
)

type companies []company

type company struct {
	name string
	code string
}

func (c companies) iter() iter.Seq[string] {
	return func(yeild func(string) bool) {
		for _, com := range c {
			fmt.Println("iter", com)
			if !yeild(com.name) {
				return
			}
		}
	}
}

func iterators() {
	coms := companies{
		company{"hello", "h"},
		company{"bye", "b"},
	}

	for x := range coms.iter() {
		fmt.Println(x)
	}

	for f := range fibonacci() {
		if f > 16 {
			break
		}
		fmt.Print(f, " ")
	}
	fmt.Println()
}

func fibonacci() iter.Seq[int] {
	return func(yeild func(int) bool) {
		i := 0
		j := 1
		for {
			if !yeild(i + j) {
				return
			}
			temp := j
			j = i + j
			i = temp
		}
	}
}
