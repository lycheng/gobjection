package container

// Container is base interface that all data structures implement.
type Container interface {
	IsEmpty() bool
	Size() int
	Clear()
	GetIterator() Iterator
}

// Iterator interface that use to traverse container
type Iterator interface {
	Next() interface{}
}
