package dsa

import "fmt"

type entry[K comparable, V any] struct {
	key   K
	value V
}

type HashMap[K comparable, V any] struct {
	bucket dynamicArray[*linkedList[entry[K, V]]]
	size   int
}

func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		bucket: *NewDynamicArray[*linkedList[entry[K, V]]](8),
	}
}

func (hm *HashMap[K, V]) Size() int {
	return hm.size
}

func (hm *HashMap[K, V]) Get(key K) (V, bool) {
	index := hm.index(key)
	ll, _ := hm.bucket.Get(index)
	if ll == nil {
		return *new(V), false
	}
	e, ok := ll.Search(
		entry[K, V]{
			key: key,
		},
		func(t1, t2 entry[K, V]) int {
			if t1.key == t2.key {
				return 0
			}
			return -1
		},
	)
	if !ok {
		return *new(V), false
	}
	return e.value, true
}

func (hm *HashMap[K, V]) Delete(key K) {
	index := hm.index(key)
	if ll, _ := hm.bucket.Get(index); ll != nil {
		ok := ll.RemoveItem(entry[K, V]{
			key: key,
		}, func(e1, e2 entry[K, V]) int {
			if e1.key == e2.key {
				return 0
			}
			return -1
		})
		if ok {
			hm.size--
		}
	}
}

func (hm *HashMap[K, V]) Put(key K, value V) {
	index := hm.index(key)
	if ll, _ := hm.bucket.Get(index); ll == nil {
		fmt.Println("YOOO")
		fmt.Println(hm.bucket.Set(index, NewLinkedList[entry[K, V]]()))
	}
	fmt.Println(hm.bucket)
	ll, _ := hm.bucket.Get(index)
	fmt.Println(ll)
	ll.Add(entry[K, V]{
		key:   key,
		value: value,
	})
	hm.size++
}

func (hm *HashMap[K, V]) hash(key K) uint64 {
	switch t := any(key).(type) {
	case int:
		return uint64(t)
	case string:
		var h uint64
		for i := 0; i < len(t); i++ {
			h = h*31 + uint64(t[i])
		}
		return h
	default:
		panic("unsupported key type")
	}
}

func (hm *HashMap[K, V]) index(key K) int {
	return int(hm.hash(key)) % hm.bucket.Size()
}
