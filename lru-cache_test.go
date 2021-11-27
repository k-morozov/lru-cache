package lru_cache

import (
	"strconv"
	"testing"
)

func TestCache_Add(t *testing.T) {
	cache, err := NewLruCache(10)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	type entry struct {
		key   int
		value string
	}

	const Count = 10
	testValues := make([]entry, Count)

	for i := 0; i < Count; i++ {
		testValues[i].key = i
		testValues[i].value = strconv.Itoa(i)
	}

	for _, temp := range testValues {
		if cache.Exists(temp.key) {
			t.Fatalf("%v exists in empty cache", temp.key)
		}
	}

	for _, temp := range testValues {
		cache.Add(temp.key, temp.value)
		if !cache.Exists(temp.key) {
			t.Fatalf("%v not exists in full cache", temp.key)
		}
	}

	for _, temp := range testValues {
		if !cache.Exists(temp.key) {
			t.Fatalf("%v not exists in full cache", temp.key)
		}
	}

	for i := 0; i < Count; i++ {
		testValues[i].key += Count
		testValues[i].value = strconv.Itoa(i + Count)
	}

	for _, temp := range testValues {
		if cache.Exists(temp.key) {
			t.Fatalf("%v exists in empty cache", temp.key)
		}
	}

	for _, temp := range testValues {
		cache.Add(temp.key, temp.value)
		if !cache.Exists(temp.key) {
			t.Fatalf("%v not exists in full cache", temp.key)
		}
	}

	if cache.elems.Len() != Count {
		t.Fatalf("size=%v, expected=%v", cache.elems.Len(), Count)
	}

	for i := 0; i < Count; i++ {
		if cache.Exists(i) {
			t.Fatalf("%v not exists in full cache", i)
		}
	}

	for _, temp := range testValues {
		if !cache.Exists(temp.key) {
			t.Fatalf("%v not exists in full cache", temp.key)
		}
	}
}
