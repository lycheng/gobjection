package singlylinkedlist

// List holds the items, where each item has a pointer point to the next item
type List struct {
	head *item
	tail *item
	size int
}

type item struct {
	value interface{}
	next  *item
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
		ni := &item{value: val}
		if list.head == nil {
			list.head = ni
			list.tail = list.head
		} else {
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
	}

	if curr == list.tail {
		list.tail = curr
	}

	if prev != nil {
		prev.next = curr.next
	}
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
