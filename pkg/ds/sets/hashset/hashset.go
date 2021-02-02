package hashset

import "github.com/lycheng/gobjection/pkg/ds/container"

// Set holds elements in go's native map
type Set struct {
	items map[interface{}]struct{}
}

var placeHolder = struct{}{}

// Iterator use to traverse Set
type Iterator struct {
	keys  []interface{}
	index int
}

// Next returns next value of the set or nil for the end
func (i *Iterator) Next() interface{} {
	if i.index >= len(i.keys) {
		return nil
	}
	rv := i.keys[i.index]
	i.index++
	return rv
}

// New returns an empty set
func New() *Set {
	return &Set{items: make(map[interface{}]struct{})}
}

// Iterator returns set iterator
func (set *Set) Iterator() container.Iterator {
	keys := set.Values()
	return &Iterator{keys: keys, index: 0}
}

// Add items (one or more) to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = placeHolder
	}
}

// Remove items in the set.
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Has check if item is present in the set.
func (set *Set) Has(item interface{}) bool {
	_, contains := set.items[item]
	return contains
}

// IsEmpty returns true if set does not contain any elements.
func (set *Set) IsEmpty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.items = make(map[interface{}]struct{})
}

// Values returns all items in the set.
func (set *Set) Values() []interface{} {
	values := make([]interface{}, set.Size())
	idx := 0
	for item := range set.items {
		values[idx] = item
		idx++
	}
	return values
}
