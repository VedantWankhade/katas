package dsa

type Queue[T any] interface {
	PushBack(t T)
	PopFront() (T, bool)
	GetFront() (T, bool)
	Len() int
}

type Deque[T any] interface {
	Queue[T]
	PushFront(t T)
	PopBack() (T, bool)
	GetBack() (T, bool)
}
