package lru_cache

type LruCache interface {
	Exists(key interface{}) bool

	Add(key, value interface{}) (ok bool)

	Get(key interface{}) (value interface{}, ok bool)

	Remove(key interface{}) (ok bool)

	Clear()
}
