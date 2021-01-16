package arraylist

import (
	"github.com/lycheng/gobjection/pkg/ds/container"
)

const (
	defaultCap = 10
)

// List holds the elements in a slice
type List struct {
	items []interface{}
}

// Iterator use to traverse List
type Iterator struct {
	list  *List
	index int
}

// Next returns next value of the list or nil for the end
func (i *Iterator) Next() interface{} {
	if i.list.IsEmpty() || i.index >= i.list.Size() {
		return nil
	}
	rv := i.list.items[i.index]
	i.index++
	return rv
}

// New instantiates a new list and adds the passed values, if any, to the list
func New(values ...interface{}) *List {
	c := defaultCap
	if len(values) > c {
		c = len(values) * 2
	}
	list := &List{}
	list.items = make([]interface{}, 0, c)
	list.Append(values...)
	return list
}

func (list *List) withinRange(i int) bool {
	if len(list.items) == 0 || i < 0 || len(list.items) <= i {
		return false
	}
	return true
}

// Append values at the end of the list
func (list *List) Append(values ...interface{}) {
	if len(values) == 0 {
		return
	}
	list.items = append(list.items, values...)
}

// Pop returns the item at index and delete it
func (list *List) Pop(i int) (interface{}, bool) {
	if !list.withinRange(i) {
		return nil, false
	}

	val := list.items[i]
	copy(list.items[i:], list.items[i+1:])
	list.items = list.items[:len(list.items)-1]
	return val, true
}

// Get returns the item at index
func (list *List) Get(i int) (interface{}, bool) {
	if !list.withinRange(i) {
		return nil, false
	}
	return list.items[i], true
}

// IsEmpty checks the list empty or not
func (list *List) IsEmpty() bool {
	return len(list.items) == 0
}

// Size returns number of items
func (list *List) Size() int {
	return len(list.items)
}

// Capacity returns the capacity of the list
func (list *List) Capacity() int {
	return cap(list.items)
}

// Clear removes all items
func (list *List) Clear() {
	list.items = make([]interface{}, 0, defaultCap)
}

// Iterator returns list iterator
func (list *List) Iterator() container.Iterator {
	return &Iterator{list: list, index: 0}
}

// Values returns list iterator
func (list *List) Values() []interface{} {
	rv := make([]interface{}, len(list.items))
	copy(rv, list.items)
	return rv
}
