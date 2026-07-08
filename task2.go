package main

import (
	"fmt"
)

type CacheEntry struct {
	Key   string
	Value any
	prev  *CacheEntry
	next  *CacheEntry
}

type CacheStats struct {
	Hits      int
	Misses    int
	Evictions int
}

type LRUCache struct {
	capacity int
	items    map[string]*CacheEntry
	head     *CacheEntry
	tail     *CacheEntry
	stats    *CacheStats
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*CacheEntry),
		stats:    &CacheStats{},
	}
}

func (c *LRUCache) moveToFront(node *CacheEntry) {

	if node == c.head {
		return
	}

	// вырезаем ноду и переназначаем текущих сосодей
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	// переназначаем хвост
	if node == c.tail {
		c.tail = node.prev
	}

	// Вставляем узел в начало списка.
	node.prev = nil
	node.next = c.head

	// переназначаем голову
	if c.head != nil {
		c.head.prev = node
	}

	c.head = node

}

func (c *LRUCache) Get(key string) (any, bool) {
	node, ok := c.items[key]

	if !ok {
		c.stats.Misses++
		return nil, false
	}
	c.moveToFront(node)
	c.stats.Hits++
	return node.Value, true

}

func (c *LRUCache) Put(key string, value any) {
	// обновление элемента
	node, ok := c.items[key]

	if ok {
		node.Value = value
		c.moveToFront(node)
		return
	}

	// удалить tail, если переполнен capacity
	if len(c.items) == c.capacity {
		c.stats.Evictions++
		delete(c.items, c.tail.Key)
		if c.tail.prev != nil {
			c.tail = c.tail.prev
			c.tail.next = nil
		} else {
			c.head = nil
			c.tail = nil
		}
	}

	newEntry := &CacheEntry{Key: key, Value: value}
	newEntry.next = c.head
	c.items[newEntry.Key] = newEntry

	if c.head != nil {
		c.head.prev = newEntry
	}

	c.head = newEntry

	if c.tail == nil {
		c.tail = newEntry
	}
}

func (c *LRUCache) Len() int {
	return len(c.items)
}

func (c *LRUCache) Stats() CacheStats {
	return *c.stats
}

func (c *LRUCache) HitRate() float64 {
	total := c.stats.Hits + c.stats.Misses
	if total == 0 {
		return 0
	}

	return float64(c.stats.Hits) / float64(total) * 100

}

func (c *LRUCache) Keys() []string {
	keys := make([]string, 0, len(c.items))
	for key := range c.items {
		keys = append(keys, key)
	}
	return keys
}

func (c *LRUCache) Clear() {
	c.items = make(map[string]*CacheEntry)
	c.head = nil
	c.tail = nil
	c.stats = &CacheStats{}
}

func (c *LRUCache) Delete(key string) bool {
	deleteNode, ok := c.items[key]
	if ok {
		if deleteNode == c.head {
			c.head = c.head.next
		}

		if deleteNode == c.tail {
			c.tail = c.tail.prev
		}

		if deleteNode.prev != nil {
			deleteNode.prev.next = deleteNode.next
		}

		if deleteNode.next != nil {
			deleteNode.next.prev = deleteNode.prev
		}

		delete(c.items, key)
		return true
	}
	return false
}

func (c *LRUCache) Contains(key string) bool {
	_, ok := c.items[key]
	return ok
}

func (c *LRUCache) Clone() *LRUCache {
	cloneCache := NewLRUCache(c.capacity)

	cloneCache.stats = &CacheStats{
		Hits:      c.stats.Hits,
		Misses:    c.stats.Misses,
		Evictions: c.stats.Evictions,
	}

	var prevEntry *CacheEntry

	for node := c.head; node != nil; node = node.next {
		cloneEntry := &CacheEntry{
			Key:   node.Key,
			Value: node.Value,
		}

		cloneCache.items[cloneEntry.Key] = cloneEntry
		cloneEntry.prev = prevEntry

		if prevEntry == nil {
			cloneCache.head = cloneEntry
		} else {
			prevEntry.next = cloneEntry
		}

		if node.next == nil {
			cloneCache.tail = cloneEntry
		}

		prevEntry = cloneEntry
	}

	return cloneCache
}

func main_task2() {
	cache := NewLRUCache(3)

	cache.Put("a", "one")
	cache.Put("b", "two")
	cache.Put("c", "three")

	fmt.Println(cache.Get("a")) // one true

	cache.Put("d", "four") // вытесняет "b"

	fmt.Println(cache.Get("b")) // <nil> false
	fmt.Println(cache.Get("a")) // one true
	fmt.Println(cache.Len())    // 3
	fmt.Println(cache.Get("c"))

	fmt.Printf("%+v\n", cache.Stats())
	fmt.Println(cache.HitRate())

	fmt.Println(cache.Contains("a"))

	clone := cache.Clone()

	fmt.Println(clone.Contains("a"))

}
