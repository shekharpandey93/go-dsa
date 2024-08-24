package main

import (
	"container/list"
	"fmt"
)

// LRUCache represents the LRU cache
type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*list.Element
	queue    *list.List
}

// entry represents a key-value pair in the cache
type entry[K comparable, V any] struct {
	key   K
	value V
}

// NewLRUCache creates a new LRU cache with the given capacity
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*list.Element),
		queue:    list.New(),
	}
}

// Get retrieves a value from the cache
func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	if elem, ok := c.cache[key]; ok {
		c.queue.MoveToFront(elem)
		return elem.Value.(*entry[K, V]).value, true
	}
	var zero V
	return zero, false
}

// Put adds a value to the cache
func (c *LRUCache[K, V]) Put(key K, value V) {
	if elem, ok := c.cache[key]; ok {
		c.queue.MoveToFront(elem)
		elem.Value.(*entry[K, V]).value = value
		return
	}

	if c.queue.Len() >= c.capacity {
		oldest := c.queue.Back()
		if oldest != nil {
			c.queue.Remove(oldest)
			delete(c.cache, oldest.Value.(*entry[K, V]).key)
		}
	}

	elem := c.queue.PushFront(&entry[K, V]{key, value})
	c.cache[key] = elem
}

func main() {
	cache := NewLRUCache[string, int](2)

	cache.Put("a", 1)
	cache.Put("b", 2)
	fmt.Println(cache.Get("a")) // Output: 1 true

	cache.Put("c", 3)           // This will evict "b"
	fmt.Println(cache.Get("b")) // Output: 0 false
	fmt.Println(cache.Get("c")) // Output: 3 true

	cache.Put("d", 4)           // This will evict "a"
	fmt.Println(cache.Get("a")) // Output: 0 false
	fmt.Println(cache.Get("c")) // Output: 3 true
	fmt.Println(cache.Get("d")) // Output: 4 true
}
