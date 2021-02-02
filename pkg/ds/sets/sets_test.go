package sets

import (
	"testing"

	"github.com/lycheng/gobjection/pkg/ds/sets/hashset"
)

func TestSetsImplementation(t *testing.T) {
	var _ Set = (*hashset.Set)(nil)
}
