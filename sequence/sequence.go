/*
Sequence is a collection of items that has some order.
*/
package sequence

import "iter"

type Sequence[T any] interface {
	Add(t T)
	AddAt(index int, t T) bool
	Len() int
	Get(index int) (T, bool)
	Set(index int, t T) bool
	Iter() iter.Seq[T]
	String() string
}
