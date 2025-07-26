package main

import (
	"fmt"
)

type SparseArray struct {
	sz int
	m  map[int]interface{}
}

func NewSparseArray(arr []interface{}, sz int) *SparseArray {
	a := &SparseArray{
		sz: sz,
		m:  make(map[int]interface{}),
	}
	for i := range arr {
		a.set(i, arr[i])
	}
	return a
}

func (a *SparseArray) set(i int, val interface{}) {
	if i < 0 || i >= a.sz {
		str := fmt.Sprintf("Invalid index: %v", i)
		panic(str)
	}
	if val != nil {
		a.m[i] = val
	} else {
		delete(a.m, i)
	}
}

func (a *SparseArray) get(i int) interface{} {
	if i < 0 || i >= a.sz {
		str := fmt.Sprintf("Invalid index: %v", i)
		panic(str)
	}
	return a.m[i]
}

func (a *SparseArray) print() {
	fmt.Printf("SparseArray: [")
	for i := 0; i < a.sz; i++ {
		fmt.Printf("%v ", a.get(i))
	}
	fmt.Println("]")
}

func convertIntSliceToInterfaceSlice(intSlice []int) []interface{} {
	interfaceSlice := make([]interface{}, len(intSlice))
	for i, val := range intSlice {
		interfaceSlice[i] = val
	}
	return interfaceSlice
}

func main() {
	a := NewSparseArray(convertIntSliceToInterfaceSlice([]int{10, 3, 4, 6, 99}), 30)
	a.set(25, 1000)
	a.print()
}
