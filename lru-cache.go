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

type Node struct {
	key   interface{}
	value interface{}
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
	if oldValue, ok := c.items[key]; ok {
		oldValue.Value.(*Node).value = value
		c.elems.MoveToFront(oldValue)
	} else {
		node := &Node{key: key, value: value}
		temp := c.elems.PushFront(node)
		if nil == temp {
			return false
		}
		c.items[key] = temp

		if c.elems.Len() > c.size {
			c.removeOldest()
		}
	}
	return true
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	node, ok := c.items[key]
	if !ok {
		return nil, ok
	}
	return node.Value.(*Node).value, ok
}

func (c *Cache) removeOldest() {
	oldest := c.elems.Back()
	if nil == oldest {
		return
	}

	c.remove(oldest)
}

func (c *Cache) remove(elem *list.Element) {
	node := c.elems.Remove(elem)
	if nil == node {
		return
	}
	delete(c.items, node.(*Node).key)
}

func (c *Cache) Clear() {

}
