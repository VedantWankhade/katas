package sequence

import (
	"bytes"
	"fmt"
	"iter"
)

type node[T any] struct {
	data T
	next *node[T]
	prev *node[T]
}

type linkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func NewLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (ll *linkedList[T]) Len() int {
	return ll.length
}

func (ll *linkedList[T]) Add(t T) {
	newNode := &node[T]{
		data: t,
		next: nil,
	}
	if ll.head == nil {
		newNode.prev = nil
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.prev = ll.tail
		ll.tail.next = newNode
		ll.tail = newNode
	}
	ll.length++
}

func (ll *linkedList[T]) AddAt(index int, t T) bool {
	if index < 0 || index > ll.Len() {
		return false
	}
	if index == ll.Len() {
		ll.Add(t)
		return true
	}
	if index == 0 {
		newNode := &node[T]{
			data: t,
			next: ll.head,
			prev: nil,
		}
		ll.head = newNode
		ll.length++
		return true
	}
	var n *node[T] = ll.head
	for i := 0; i < index; i, n = i+1, n.next {
	}
	newNode := &node[T]{
		data: t,
		next: n,
		prev: n.prev,
	}
	n.prev.next = newNode
	n.prev = newNode
	ll.length++
	return true
}

func (ll *linkedList[T]) Get(index int) (T, bool) {
	if index >= ll.Len() {
		return *new(T), false
	}

	n := ll.head

	for i := 0; i < index; i++ {
		n = n.next
	}

	return n.data, true
}

func (ll *linkedList[T]) String() string {
	var out bytes.Buffer
	out.WriteString("[")
	for n := ll.head; n != nil; n = n.next {
		out.WriteString(fmt.Sprintf("%v, ", n.data))
	}
	out.WriteString("]")
	return out.String()
}

func (ll *linkedList[T]) Set(index int, t T) bool {
	if index < 0 || index >= ll.Len() {
		return false
	}
	n := ll.head
	for i := 0; i < index; i, n = i+1, n.next {

	}
	n.data = t
	return true
}

func (ll *linkedList[T]) Iter() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, n := 0, ll.head; n != nil; i, n = i+1, n.next {
			if !yield(i, n.data) {
				return
			}
		}
	}
}
