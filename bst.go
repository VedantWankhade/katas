package dsa

import (
	"fmt"
	"iter"
)

type treenode[T ~int | ~string | ~float64] struct {
	data  T
	left  *treenode[T]
	right *treenode[T]
}

type BST[T ~int | ~string | ~float64] struct {
	root *treenode[T]
	size int
}

func NewBST[T ~int | ~string | ~float64]() *BST[T] {
	return &BST[T]{
		root: nil,
		size: 0,
	}
}

func (t *BST[T]) Size() int {
	return t.size
}

func (t *BST[T]) Add(item T) {
	newNode := &treenode[T]{
		data:  item,
		left:  nil,
		right: nil,
	}
	if t.size == 0 {
		t.root = newNode
		t.size++
		return
	}

	for n := t.root; n != nil; {
		if item > n.data {
			if n.right == nil {
				n.right = newNode
				break
			} else {
				n = n.right
			}
		} else {
			if n.left == nil {
				n.left = newNode
				break
			} else {
				n = n.left
			}
		}
	}
	t.size++
}

func (t *BST[T]) dfs(n *treenode[T]) {
	if n == nil {
		return
	}
	t.dfs(n.left)
	fmt.Println(n.data)
	t.dfs(n.right)
}

func (t *BST[T]) dfsIter(n *treenode[T], yeild func(T) bool) {
	if n == nil {
		return
	}
	t.dfs(n.left)
	if !yeild(n.data) {
		return
	}
	t.dfs(n.right)
}

func (t *BST[T]) Dfs() {
	t.dfs(t.root)
}

func (t *BST[T]) DfsIter() iter.Seq[T] {
	return func(yield func(T) bool) {
		t.dfsIter(t.root, yield)
	}
}

func (t *BST[T]) DfsIterative() {
	pushLeftSubtree := func(n *treenode[T], s Stack[*treenode[T]]) {
		for n != nil {
			s.Push(n)
			n = n.left
		}
	}

	s := NewStack[*treenode[T]]()
	pushLeftSubtree(t.root, s)
	n, ok := s.Pop()
	for ok {
		fmt.Println(n.data)
		pushLeftSubtree(n.right, s)
		n, ok = s.Pop()
	}
}

func (t *BST[T]) BfsIterative() {
	q := NewQueue[*treenode[T]]()
	q.PushBack(t.root)
	n, ok := q.PopFront()
	for ok || n != nil {
		fmt.Println(n.data)
		if n.left != nil {
			q.PushBack(n.left)
		}
		if n.right != nil {
			q.PushBack(n.right)
		}
		n, ok = q.PopFront()
	}
}
