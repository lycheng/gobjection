package sets

import "github.com/lycheng/gobjection/pkg/ds/container"

// Set interface that all sets implement
type Set interface {
	Add(items ...interface{})
	Remove(items ...interface{})
	Has(item interface{}) bool

	container.Container
	// IsEmpty() bool
	// Size() int
	// Clear()
	// Iterator() Iterator
	// Values() []interface{}
}
