package lru_cache

import (
	"container/list"
	"errors"
)

type Cache struct {
	size  int
	elems *list.List
	items map[interface{}]*list.Element
}

func NewLruCache(size int) (*Cache, error) {
	if size < 0 {
		return nil, errors.New("must provides a positive size")
	}

	cache := &Cache{
		size:  size,
		elems: list.New(),
		items: make(map[interface{}]*list.Element),
	}

	return cache, nil
}

func (c *Cache) Exists(key interface{}) bool {
	_, b := c.items[key]
	return b
}

func (c *Cache) Add(key, value interface{}) (ok bool) {
	if !c.Exists(key) {
		temp := c.elems.PushFront(value)
		c.items[key] = temp

		if c.elems.Len() > c.size {
			c.elems.Remove(c.elems.Back())
		}
	} else {
		if oldValue, ok := c.items[key]; ok {
			oldValue.Value = value
			c.elems.MoveToFront(oldValue)
		}
	}

	return true
}

func (c *Cache) Get(key interface{}) (value interface{}, ok bool) {
	value, ok = c.items[key]
	return
}

func (c *Cache) Remove(key interface{}) (ok bool) {
	return
}

func (c *Cache) Clear() {

}
