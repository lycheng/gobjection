package doublylinkedlist

import "github.com/lycheng/gobjection/pkg/ds/container"

// List holds the elements, where each element points to the next and previous element
type List struct {
	head *item
	tail *item
	size int
}

type item struct {
	value interface{}
	prev  *item
	next  *item
}

// Iterator use to traverse List
type Iterator struct {
	curr *item
}

// ReverseIterator use for reverse loop
type ReverseIterator struct {
	curr *item
}

// Next returns next value of the list or nil for the end
func (i *Iterator) Next() interface{} {
	if i.curr == nil {
		return nil
	}
	rv := i.curr.value
	i.curr = i.curr.next
	return rv
}

// Next returns prev value of the list or nil for the head
func (i *ReverseIterator) Next() interface{} {
	if i.curr == nil {
		return nil
	}
	rv := i.curr.value
	i.curr = i.curr.prev
	return rv
}

// New instantiates a new list and adds the passed values, if any, to the list
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Append(values...)
	}
	return list
}

// Append adds values to the end of the list
func (list *List) Append(values ...interface{}) {
	if len(values) == 0 {
		return
	}
	for _, val := range values {
		ni := &item{value: val, prev: list.tail}
		if list.head == nil {
			list.head = ni
			list.tail = list.head
		} else {
			ni.prev = list.tail
			list.tail.next = ni
			list.tail = ni
		}
	}
	list.size += len(values)
}

// Pop returns the item at index and delete it
func (list *List) Pop(i int) (interface{}, bool) {
	if list.size == 0 || i < 0 || list.size <= i {
		return nil, false
	}

	var prev *item
	var curr *item = list.head
	for j := 0; j != i; j++ {
		prev = curr
		curr = curr.next
	}

	if curr == list.head {
		list.head = curr.next
		curr.prev = nil
	}

	if curr == list.tail {
		list.tail = curr
		curr.next = nil
	}

	if prev != nil {
		prev.next = curr.next
		curr.next.prev = prev
	}
	list.size--
	return curr.value, true
}

// Get returns the item at index
func (list *List) Get(i int) (interface{}, bool) {
	if list.size == 0 || i < 0 || list.size <= i {
		return nil, false
	}

	curr := list.head
	for j := 0; j < i; j++ {
		curr = curr.next
	}
	return curr.value, true
}

// IsEmpty checks the list empty or not
func (list *List) IsEmpty() bool {
	return list.size == 0
}

// Size returns number of items
func (list *List) Size() int {
	return list.size
}

// Clear removes all items
func (list *List) Clear() {
	list.size = 0
	list.head = nil
	list.tail = nil
}

// Iterator returns list iterator
func (list *List) Iterator() container.Iterator {
	return &Iterator{curr: list.head}
}

// ReverseIterator returns list iterator
func (list *List) ReverseIterator() container.Iterator {
	return &ReverseIterator{curr: list.tail}
}

// Values returns list iterator
func (list *List) Values() []interface{} {
	rv := make([]interface{}, list.size)
	for i, j := 0, list.head; j != nil; i, j = i+1, j.next {
		rv[i] = j.value
	}
	return rv
}
