package dsa

import "github.com/vedantwankhade/katas/dsa/sequence"

type linkedStack[T any] struct {
	ll sequence.Sequence[T]
}

func NewStack[T any]() Stack[T] {
	return &linkedStack[T]{
		ll: sequence.NewLinkedList[T](),
	}
}

func (ls *linkedStack[T]) Push(t T) {
	ls.ll.Add(t)
}

func (ls *linkedStack[T]) Pop() (T, bool) {
	return ls.ll.Remove(ls.ll.Len() - 1)
}

func (ls *linkedStack[T]) Peek() (T, bool) {
	return ls.ll.Get(ls.ll.Len() - 1)
}

func (ls *linkedStack[T]) Len() int {
	return ls.ll.Len()
}
