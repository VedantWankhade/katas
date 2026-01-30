package dsa

import (
	"bytes"
	"fmt"
)

type heapElementType interface {
	~int | ~string | ~float64
}

type heap[T heapElementType] struct {
	list []*T
	comp func(T, T) int
	size int
}

func NewHeap[T heapElementType](comp func(T, T) int) *heap[T] {
	return &heap[T]{
		list: make([]*T, 0),
		comp: comp,
		size: 0,
	}
}

func (h *heap[T]) Size() int {
	return h.size
}

func (h *heap[T]) add(t T) {
	h.list = append(h.list, &t)
	currIndex := h.Size() - 1
	parentIndex := (currIndex - 1) / 2
	for currIndex > 0 && h.comp(*h.list[currIndex], *h.list[parentIndex]) > 0 {
		h.list[currIndex], h.list[parentIndex] = h.list[parentIndex], h.list[currIndex]
		currIndex = parentIndex
		parentIndex = (currIndex - 1) / 2
	}
	h.size++
}

func (h *heap[T]) peek() T {
	return *h.list[0]
}

func (h *heap[T]) get() *T {
	if h.Size() <= 0 {
		return nil
	}
	ret := h.list[0]
	h.list[0] = h.list[h.Size()-1]
	h.list[h.Size()-1] = nil
	h.size--

	currIndex := 0

	for {
		leftIndex := currIndex*2 + 1
		rightIndex := currIndex*2 + 2
		maxIndex := currIndex

		if leftIndex < h.Size() && h.comp(*h.list[maxIndex], *h.list[leftIndex]) < 0 {
			maxIndex = leftIndex
		}
		if rightIndex < h.Size() && h.comp(*h.list[maxIndex], *h.list[rightIndex]) < 0 {
			maxIndex = rightIndex
		}
		if maxIndex == currIndex {
			break
		}
		h.list[currIndex], h.list[maxIndex] = h.list[maxIndex], h.list[currIndex]
		currIndex = maxIndex
	}

	return ret
}

func (h *heap[T]) String() string {
	var out bytes.Buffer
	for _, t := range h.list {
		out.WriteString(fmt.Sprintf("%v", *t))
	}
	return out.String()
}
