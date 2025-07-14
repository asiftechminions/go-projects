package main

import "fmt"

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func InitiateCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		cache:    make(map[int]*Node),
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, exists := lru.cache[key]; exists {
		lru.MoveToFront(node)
		return node.value
	}
	return -1

}

func (lru *LRUCache) Put(key, value int) {
	if node, exist := lru.cache[key]; exist {
		node.value = value
		lru.MoveToFront(node)
	} else if lru.capacity == len(lru.cache) {
		lastCache := lru.RemoveLast()
		delete(lru.cache, lastCache.key)
	}
	node := &Node{
		key:   key,
		value: value,
	}
	lru.cache[key] = node
	lru.AddToFront(node)
}

func (lru *LRUCache) RemoveNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (lru *LRUCache) RemoveLast() *Node {
	node := lru.tail.prev
	lru.RemoveNode(node)
	return node
}

func (lru *LRUCache) MoveToFront(node *Node) {
	lru.RemoveNode(node)
	lru.AddToFront(node)
}

func (lru *LRUCache) AddToFront(node *Node) {
	currFront := lru.head.next

	node.prev = lru.head
	node.next = currFront
	currFront.prev = node
	lru.head.next = node

}

func main() {
	lru := InitiateCache(3)

	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(3))
	lru.Put(4, 4)
	fmt.Println(lru.Get(2))
	lru.Put(5, 5)
	fmt.Println(lru.Get(1))
}
