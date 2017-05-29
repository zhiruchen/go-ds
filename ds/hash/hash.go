package hash

import (
	"container/list"
)

type kv struct {
	Key string
	Val interface{}
}

// KVPair map keys to values
type KVPair struct {
	TableSize uint64
	Buckets   []*list.List
	Count     uint64
}
