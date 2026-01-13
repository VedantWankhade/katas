package dsa

import (
	"bytes"
	"fmt"
	"iter"
)

const load_factor = 2

type dynamicArray[T any] struct {
	arr  []T
	len  int
	size int
}

// Returns a dynamic array with initial size of 8
func NewDynamicArray[T any](initialSize int) *dynamicArray[T] {
	return &dynamicArray[T]{
		arr:  make([]T, initialSize),
		len:  0,
		size: initialSize,
	}
}

func (da *dynamicArray[T]) Add(t T) {
	if da.len == da.size {
		da.resize()
	}
	da.arr[da.len] = t
	da.len++
}

func (da *dynamicArray[T]) Remove(index int) (T, bool) {
	ret, ok := da.Get(index)
	if !ok {
		return *new(T), false
	}
	if index == da.Len()-1 {
		da.arr[index] = *new(T)
		da.len--
		return ret, true
	}
	for i := index + 1; i < da.Len(); i++ {
		da.arr[i-1] = da.arr[i]
	}
	da.arr[da.Len()-1] = *new(T)
	da.len--
	return ret, true
}

func (da *dynamicArray[T]) AddAt(index int, t T) bool {
	if index < 0 || index > da.Len() {
		return false
	}
	if da.len == da.size {
		da.resize()
	}
	for i := da.len - 1; i >= index; i-- {
		da.arr[i+1] = da.arr[i]
	}
	da.arr[index] = t
	da.len++
	return true
}

func (da *dynamicArray[T]) resize() {
	newSize := load_factor * da.size
	newArr := make([]T, newSize)
	copy(newArr, da.arr)
	da.size = newSize
	da.arr = newArr
}

func (da *dynamicArray[T]) Len() int {
	return da.len
}

func (da *dynamicArray[T]) Get(i int) (T, bool) {
	if i >= da.Len() || i < 0 {
		return *new(T), false
	}
	return da.arr[i], true
}

func (da *dynamicArray[T]) Size() int {
	return da.size
}

func (da *dynamicArray[T]) Set(i int, t T) bool {
	if i < da.Len() && i >= 0 {
		da.arr[i] = t
		return true
	}
	return false
}

func (da *dynamicArray[T]) Iter() iter.Seq2[int, T] {
	return func(yeild func(int, T) bool) {
		for i, t := range da.arr[0:da.Len()] {
			if !yeild(i, t) {
				return
			}
		}
	}
}

func (da *dynamicArray[T]) String() string {
	var out bytes.Buffer
	out.WriteString("[")
	for i, t := range da.arr[0:da.len] {
		out.WriteString(fmt.Sprintf("%v", t))
		if i < da.Len() {
			out.WriteString(", ")
		}
	}
	out.WriteString("]")

	return out.String()
}
