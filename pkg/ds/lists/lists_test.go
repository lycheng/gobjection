package lists

import (
	"testing"

	"github.com/lycheng/gobjection/pkg/ds/lists/singlylinkedlist"
)

func TestListsImplementation(t *testing.T) {
	var _ List = (*singlylinkedlist.List)(nil)
}
