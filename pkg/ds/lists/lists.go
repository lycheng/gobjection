package lists

import "github.com/lycheng/gobjection/pkg/ds/container"

// List interface that all lists implement
type List interface {
	Append(values ...interface{})
	Pop(i int) (interface{}, bool)
	Get(i int) (interface{}, bool)

	container.Container
	// IsEmpty() bool
	// Size() int
	// Clear()
	// GetIterator() Iterator
}
