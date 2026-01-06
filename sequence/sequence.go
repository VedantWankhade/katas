/*
Sequence is a collection of items that has some order.
*/
package sequence

import "iter"

type Sequence[T any] interface {
	Add(t T)
	Len() int
	Get(index int) T
	Set(index int, t T)
	Iter() iter.Seq[T]
}
