package algs

import "testing"

func TestLRUCache_Put(t *testing.T) {
	testCases := []struct {
		key  string
		val  interface{}
		size int32
	}{
		{
			key:  "key1",
			val:  1,
			size: 1,
		},
		{
			key:  "key2",
			val:  2,
			size: 2,
		},
		{
			key:  "key1",
			val:  100,
			size: 2,
		},
	}

	cache := NewLRUCache()
	for _, cc := range testCases {
		cache.Put(cc.key, cc.val)
		sz := cache.Size()
		if sz != cc.size {
			t.Errorf("expect size: %d, get: %d", cc.size, sz)
		}

		v := cache.Get(cc.key)
		if v != cc.val {
			t.Errorf("expect val: %v, get: %v", cc.val, v)
		}
	}
}

func TestLRUCache_Remove(t *testing.T) {
	keys := []string{"abc", "123", "def", "hjk"}
	testCases := []struct {
		key  string
		size int32
	}{
		{
			key:  "abc",
			size: 3,
		},
		{
			key:  "456",
			size: 3,
		},
		{
			key:  "def",
			size: 2,
		},
	}

	cache := NewLRUCache()
	for _, k := range keys {
		cache.Put(k, true)
	}

	for _, cc := range testCases {
		cache.Remove(cc.key)
		sz := cache.Size()

		if sz != cc.size {
			t.Errorf("expect cache size: %d, get: %d\n", cc.size, sz)
		}

		v := cache.Get(cc.key)
		if v != nil {
			t.Errorf("cache should be nil, get: %v\n", v)
		}
	}
}
