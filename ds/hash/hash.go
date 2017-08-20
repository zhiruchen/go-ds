package hash

import (
	"container/list"
	"hash/fnv"
)

const bucketSize uint64 = 1 << 4
const loadFactor float64 = .75

type kv struct {
	Key string
	Val interface{}
}

// KVPairs hold the slice of kv pairs
type KVPairs struct {
	KVs []*kv
}

// Map map keys to values
type Map struct {
	Buckets    []*list.List
	Count      uint64  // 实际有值的桶个数
	LoadFactor float64 // 装载因子
	Size       uint64  // 所有桶的个数
}

// NewMap create a new Map
func NewMap() *Map {
	buckets := []*list.List{}
	var i uint64
	for i < bucketSize {
		buckets = append(buckets, list.New())
		i++
	}
	return &Map{Buckets: buckets, Count: 0, Size: bucketSize}
}

// calculate index of key
func hash(key string) uint64 {
	h := fnv.New64()
	h.Write([]byte(key))
	return h.Sum64()
}

// Put insert a kv pair to buckets
func (m *Map) Put(key string, val interface{}) {
	i := hash(key) % m.Size
	l := m.Buckets[i]
	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(*kv)
		if p.Key == key {
			if p.Val == val {
				return
			}
			e.Value = &kv{Key: p.Key, Val: val}
			return

		}
	}
	l.PushBack(&kv{Key: key, Val: val})
	m.Buckets[i] = l
	m.Count++
	m.LoadFactor = float64(m.Count) / float64(m.Size)
	if m.LoadFactor > loadFactor {
		m.reHash()
	}
}

// 当装载因子>=.75时, 扩展桶个数为原来的2倍
func (m *Map) reHash() {
	m.Size *= 2
	buckets := []*list.List{}
	var i uint64
	for i < m.Size {
		buckets = append(buckets, list.New())
		i++
	}

	for _, l := range m.Buckets {
		for e := l.Front(); e != nil; e = e.Next() {
			p := e.Value.(*kv)
			i := hash(p.Key) % m.Size
			buckets[i].PushBack(&kv{Key: p.Key, Val: p.Val})
		}
	}

	m.Buckets = buckets
}

// Get get a val from map by key
func (m *Map) Get(key string) interface{} {
	i := hash(key) % m.Size
	l := m.Buckets[i]
	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(*kv)
		if p.Key == key {
			return p.Val
		}
	}
	return nil
}

// Delete delete a key from map
func (m *Map) Delete(key string) interface{} {
	i := hash(key) % m.Size
	l := m.Buckets[i]
	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(*kv)
		if p.Key == key {
			l.Remove(e)
			m.Buckets[i] = l
			m.Count--
			m.LoadFactor = float64(m.Count) / float64(m.Size)
			return p.Val
		}
	}
	return nil
}

// GetKVPairs get all kv pairs of map
func (m *Map) GetKVPairs() *KVPairs {
	kvp := &KVPairs{}
	kvp.KVs = make([]*kv, m.Count)
	for _, l := range m.Buckets {
		for e := l.Front(); e != nil; e = e.Next() {
			kvp.KVs = append(kvp.KVs, e.Value.(*kv))
		}
	}
	return kvp
}
