package main

import (
	"container/list"
	"fmt"
)

type Entry struct {
	key   interface{}
	value interface{}
}

type Cache struct {
	maxEntries int
	ll         *list.List
	cache      map[interface{}]*list.Element
}

func New(maxEntries int) *Cache {
	return &Cache{
		maxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

func (c *Cache) Add(key interface{}, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}

	if _, ok := c.cache[key]; ok {
		c.Remove(key)
	}

	ele := c.ll.PushFront(&Entry{key: key, value: value})
	c.cache[key] = ele

	if c.maxEntries != 0 && c.ll.Len() > c.maxEntries {
		c.RemoveOldest()
	}

}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	if c.cache == nil {
		return nil, false
	}
	if ele, ok := c.cache[key]; ok {
		c.ll.PushFront(ele)
		return ele.Value.(*Entry).value, true
	}
	return nil, false
}

func (c *Cache) Remove(key interface{}) bool {
	if c.cache == nil {
		return false
	}

	if ele, ok := c.cache[key]; ok {
		entry := ele.Value.(*Entry)
		delete(c.cache, entry.key)
		c.ll.Remove(ele)
		return true
	}
	return false
}

func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}

	ele := c.ll.Back()
	c.ll.Remove(ele)
	entry := ele.Value.(*Entry)
	delete(c.cache, entry.key)
}

func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}

	return c.ll.Len()
}

func main() {
	var maxEntries int
	fmt.Scanf("%d", &maxEntries)
	cache := New(maxEntries)
	cache.Add(1, 2)
	fmt.Println(cache.cache)
	cache.Add(6, 10)
	fmt.Println(cache.cache)
	cache.Add(1, 15)
	fmt.Println(cache.cache)
	cache.Add(2, 4)
	fmt.Println(cache.cache)
	cache.Add(3, 5)
	fmt.Println(cache.cache)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}
