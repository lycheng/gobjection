package singlylinkedlist

import "testing"

func TestNewList(t *testing.T) {
	a := New(1, 2, 3)
	if a.Size() != 3 {
		t.Error("list size not equals to 3")
	}

	b := New()
	if b.Size() != 0 {
		t.Error("list size not equals to 0")
	}
}
