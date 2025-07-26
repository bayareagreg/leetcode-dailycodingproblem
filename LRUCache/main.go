package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	prev  *Node
	key   interface{}
	value interface{}
}

func NewNode(key interface{}, value interface{}, prev *Node, next *Node) *Node {
	return &Node{
		key:   key,
		value: value,
		next:  next,
		prev:  prev,
	}
}

type LRUCache struct {
	sz    int
	m     map[interface{}]*Node
	first *Node
	last  *Node
}

func NewLRUCache(n int) *LRUCache {
	return &LRUCache{
		sz:    n,
		m:     make(map[interface{}]*Node, n),
		first: nil,
		last:  nil,
	}
}

func (c *LRUCache) set(key interface{}, value interface{}) {
	node, ok := c.m[key]
	if !ok {
		node := NewNode(key, value, nil, nil)
		c.m[key] = node
		if c.first != nil {
			c.first.prev = node
			node.next = c.first
		}
		c.first = node
		if c.last == nil {
			c.last = node
		}
	} else {
		node.value = value
		c.moveNodeToFront(node)
	}

	// remove eldest if new size > c.sz
	if len(c.m) > c.sz {
		delete(c.m, c.last.key)
		c.last = c.last.prev
		c.last.next = nil
	}
}

func (c *LRUCache) get(key interface{}) interface{} {
	node, ok := c.m[key]
	if !ok {
		return nil
	} else {
		// cache hit - move to the front
		c.moveNodeToFront(node)
		return node.value
	}
}

func (c *LRUCache) moveNodeToFront(node *Node) {
	// cache hit - move to the front
	if c.first != node {
		// first point last
		if c.last == node {
			c.last = c.last.prev
		}
		// now, remove node from middle or end of the list
		node.prev.next = node.next
		if node.next != nil {
			node.next.prev = node.prev
		}
		// then point first
		node.next = c.first
		node.next.prev = node
		node.prev = nil
		c.first = node
	}
}

func (c *LRUCache) Print() {
	fmt.Printf("LRUCache:\n")
	fmt.Printf("  max size: %v\n", c.sz)
	fmt.Printf("  map: %v\n", c.m)
	fmt.Printf("  first: ")
	for node := c.first; node != nil; node = node.next {
		fmt.Printf("%v,", node.key)
	}
	fmt.Println()
	fmt.Printf("  last: ")
	for node := c.last; node != nil; node = node.next {
		fmt.Printf("%v,", node.key)
	}
	fmt.Println()
}

func main() {
	/*
		c := NewLRUCache(3)
		c.set("foo", "bar")
		c.Print()

		c.set("zot", "xyz")
		c.Print()

		c.set("baz", "foo")
		c.Print()

		c.set("coo", "koo")
		c.Print()

		c.get("zot")
		c.Print()
		c.get("zot")
		c.Print()
		c.get("zot")
		c.Print()
	*/
	c := NewLRUCache(3)
	c.set(1, 1)
	c.set(2, 2)
	c.set(3, 3)
	c.set(4, 4)
	v := c.get(4)
	fmt.Println(v)
	c.Print()
	v = c.get(3)
	fmt.Println(v)
	c.Print()
	v = c.get(2)
	fmt.Println(v)
	v = c.get(1)
	fmt.Println(v)
	c.Print()
	c.set(5, 5)
	c.Print()

	/*
		c.Print()

		v := c.get(2)
		fmt.Println(v)
	*/

}
