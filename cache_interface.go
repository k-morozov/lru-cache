package lru_cache

import "container/list"

type LruCache interface {
	Exists(key interface{}) bool

	Add(key, value interface{}) (ok bool)

	Get(key interface{}) (value interface{}, ok bool)

	removeOldest()

	remove(element *list.Element)

	Clear()
}
