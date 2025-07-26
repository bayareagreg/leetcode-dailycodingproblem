package main

import (
	"fmt"
	"math/rand"
)

// I am assuming since its a set, duplicates are not allowed

type RandomizedSet struct {
	m  map[int]int
	ar []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:  make(map[int]int),
		ar: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	ix, ok := this.m[val]
	if !ok {
		if len(this.m) < len(this.ar) {
			this.ar[len(this.m)] = val
			this.m[val] = len(this.m)
		} else {
			this.ar = append(this.ar, val)
			this.m[val] = len(this.ar) - 1
		}
	} else {
		this.ar[ix] = val
	}
	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	ix, ok := this.m[val]
	if ok {
		delete(this.m, val)
		// move just the end element to where the gap is
		if ix != len(this.m) {
			this.ar[ix] = this.ar[len(this.m)]
			this.m[this.ar[ix]] = ix
		}
	}
	return ok
}

func (this *RandomizedSet) GetRandom() int {
	ix := rand.Intn(len(this.m))
	return this.ar[ix]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	obj := Constructor()
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Insert(0))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Insert(0))

	/*
		fmt.Println(obj.m)
		fmt.Println(obj.ar)

		fmt.Println(obj.Remove(0))
		fmt.Println(obj.m)
		fmt.Println(obj.ar)
	*/

	//fmt.Println(obj.Insert(2))
	//fmt.Println(obj.Remove(1))

}
