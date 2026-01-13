package dsa

type deque[T any] struct {
	ll Sequence[T]
}

func NewDeque[T any]() Deque[T] {
	return &deque[T]{
		ll: NewLinkedList[T](),
	}
}

func NewQueue[T any]() Queue[T] {
	return &deque[T]{
		ll: NewLinkedList[T](),
	}
}

func (d *deque[T]) Len() int {
	return d.ll.Len()
}

func (d *deque[T]) PushBack(t T) {
	d.ll.Add(t)
}

func (d *deque[T]) PushFront(t T) {
	d.ll.AddAt(0, t)
}

func (d *deque[T]) GetBack() (T, bool) {
	return d.ll.Get(d.ll.Len() - 1)
}

func (d *deque[T]) GetFront() (T, bool) {
	return d.ll.Get(0)
}

func (d *deque[T]) PopBack() (T, bool) {
	return d.ll.Remove(d.ll.Len() - 1)
}

func (d *deque[T]) PopFront() (T, bool) {
	return d.ll.Remove(0)
}
