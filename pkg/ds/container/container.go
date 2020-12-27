package container

// Container is base interface that all data structures implement.
type Container interface {
	IsEmpty() bool
	Size() int
	Clear()
	Iterator() Iterator
	Values() []interface{}
}

// Iterator interface that use to traverse container
type Iterator interface {
	Next() interface{}
}
