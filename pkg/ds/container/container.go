package container

// Container is base interface that all data structures implement.
type Container interface {
	IsEmpty() bool
	Size() int
	Clear()
}
