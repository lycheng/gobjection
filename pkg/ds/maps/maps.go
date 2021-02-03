package maps

import "github.com/lycheng/gobjection/pkg/ds/container"

// Map interface that all maps implement
type Map interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) (value interface{}, found bool)
	Has(key interface{}) bool
	Remove(key interface{})
	Keys() []interface{}

	container.Container
	// IsEmpty() bool
	// Size() int
	// Clear()
	// Iterator() Iterator
	// Values() []interface{}
}

// KV stores kev, value pair
type KV struct {
	key interface{}
	val interface{}
}

// KVPair creates KV instance
func KVPair(key interface{}, val interface{}) *KV {
	return &KV{key: key, val: val}
}

// Key returns the key
func (kv *KV) Key() interface{} {
	return kv.key
}

// Value returns the key
func (kv *KV) Value() interface{} {
	return kv.val
}
