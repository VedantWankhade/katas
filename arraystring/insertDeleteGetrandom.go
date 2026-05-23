package arraystring

import (
	"math/rand"
)

// Problem: https://leetcode.com/problems/insert-delete-getrandom-o1/description/

type RandomizedSet struct {
	hash map[int]struct {
		index int
	}
	list []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hash: make(map[int]struct{ index int }),
		list: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hash[val]; ok {
		return false
	}
	this.list = append(this.list, val)
	this.hash[val] = struct {
		index int
	}{
		index: len(this.list) - 1,
	}
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.hash[val]; ok {
		lastElement := this.list[len(this.list)-1]
		indexOfElementToRemove := this.hash[val].index
		this.list[indexOfElementToRemove] = lastElement
		this.hash[lastElement] = struct{ index int }{
			index: indexOfElementToRemove,
		}
		this.list = this.list[:len(this.list)-1]
		delete(this.hash, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	randIndex := rand.Intn(len(this.hash))
	return this.list[randIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
