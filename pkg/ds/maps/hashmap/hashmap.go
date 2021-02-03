package hashmap

import (
	"github.com/lycheng/gobjection/pkg/ds/container"
	"github.com/lycheng/gobjection/pkg/ds/maps"
)

// Map holds the elements in go's native map
type Map struct {
	m map[interface{}]interface{}
}

// Iterator use to traverse map
type Iterator struct {
	kvs   []*maps.KV
	index int
}

// Next returns next value of the map or nil for the end
func (i *Iterator) Next() interface{} {
	if i.index >= len(i.kvs) {
		return nil
	}
	kv := i.kvs[i.index]
	i.index++
	return kv
}

// New instantiates a hash map.
func New() *Map {
	return &Map{m: make(map[interface{}]interface{})}
}

// Put inserts element into the map.
func (m *Map) Put(key interface{}, val interface{}) {
	m.m[key] = val
}

// Get returns the value in the map by key or nil if key is not found
func (m *Map) Get(key interface{}) (val interface{}, found bool) {
	val, found = m.m[key]
	return
}

// Has checks if key exist
func (m *Map) Has(key interface{}) (found bool) {
	_, found = m.m[key]
	return
}

// Remove removes the element from the map by key.
func (m *Map) Remove(key interface{}) {
	delete(m.m, key)
}

// Keys returns all keys
func (m *Map) Keys() []interface{} {
	keys := make([]interface{}, 0, m.Size())
	for key := range m.m {
		keys = append(keys, key)
	}
	return keys
}

// Values returns all values
func (m *Map) Values() []interface{} {
	vals := make([]interface{}, 0, m.Size())
	for _, val := range m.m {
		vals = append(vals, val)
	}
	return vals
}

// IsEmpty returns true if map does not contain any elements
func (m *Map) IsEmpty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return len(m.m)
}

// Clear will erasure all the items
func (m *Map) Clear() {
	m.m = make(map[interface{}]interface{})
}

// Iterator returns map iterator
func (m *Map) Iterator() container.Iterator {
	kvs := make([]*maps.KV, 0, m.Size())
	for key, val := range m.m {
		kv := maps.KVPair(key, val)
		kvs = append(kvs, kv)
	}
	return &Iterator{kvs: kvs, index: 0}
}
