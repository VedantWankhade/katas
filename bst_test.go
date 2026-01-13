package dsa_test

import (
	"fmt"
	"testing"

	"github.com/vedantwankhade/katas/dsa"
)

func TestBST(t *testing.T) {
	tr := dsa.NewBST[string]()
	tr.Add("A")
	tr.Add("B")
	tr.Add("X")
	tr.Add("D")
	tr.Add("C")
	tr.Add("Y")

	fmt.Println(tr.Size())

	tr.Dfs()
	fmt.Println("-------")

	for t := range tr.DfsIter() {
		fmt.Println(t)
	}
	fmt.Println("-------")
	tr.DfsIterative()
	fmt.Println("-------")
	tr.BfsIterative()
}
