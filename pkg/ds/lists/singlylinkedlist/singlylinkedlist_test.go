package singlylinkedlist

import (
	"testing"

	"github.com/lycheng/gobjection/pkg/ds/lists"
)

func assertListEquals(values []interface{}, expected []int) bool {
	for idx, i := range values {
		if expected[idx] != i.(int) {
			return false
		}
	}
	return true
}

func TestListsImplementation(t *testing.T) {
	var _ lists.List = (*List)(nil)
}

func TestNewList(t *testing.T) {
	a := New(1, 2, 3)

	if a.Size() != 3 {
		t.Error("list size not equals to 3")
	}

	if a.IsEmpty() {
		t.Error("list should not empty")
	}

	if !assertListEquals(a.Values(), []int{1, 2, 3}) {
		t.Error("list values error")
	}

	b := New()
	if b.Size() != 0 {
		t.Error("list size not equals to 0")
	}

	a.Clear()
	if a.Size() != 0 {
		t.Error("list size must equals to 0")
	}

	if !a.IsEmpty() {
		t.Error("list has to be empty")
	}
}

func TestListIterator(t *testing.T) {
	a := New(1, 2, 3)
	iter := a.Iterator()
	if iter.Next() != 1 {
		t.Error("list iterator return not equals to 1")
	}
	if iter.Next() != 2 {
		t.Error("list iterator return not equals to 2")
	}
	if iter.Next() != 3 {
		t.Error("list iterator return not equals to 3")
	}
	if iter.Next() != nil {
		t.Error("list iterator return not equals to nil")
	}
}

func TestListGet(t *testing.T) {
	a := New()

	_, flag := a.Get(1)
	if flag {
		t.Error("list has no value at index 1")
	}

	a = New(1, 2, 3)
	v, flag := a.Get(1)
	if !flag || v != 2 {
		t.Error("list get item at index 1 error")
	}

	v, flag = a.Get(-1)
	if flag || v != nil {
		t.Error("list get item at index -1 must return nil")
	}

	v, flag = a.Get(3)
	if flag || v != nil {
		t.Error("list get item at index 3 must return nil")
	}
}

func TestListPopAndAppend(t *testing.T) {
	a := New()
	a.Append(1)

	v, flag := a.Get(0)
	if v != 1 || !flag || a.Size() != 1 {
		t.Error("list append 1 error")
	}

	a.Append(2, 3)
	v, flag = a.Get(2)
	if v != 3 || !flag || a.Size() != 3 {
		t.Error("list append 2, 3 error")
	}

	v, flag = a.Pop(3)
	if v != nil || flag || a.Size() != 3 {
		t.Error("list pop index 3 error")
	}

	v, flag = a.Pop(1)
	if v != 2 || !flag || a.Size() != 2 {
		t.Error("list pop index 2 error")
	}

	if !assertListEquals(a.Values(), []int{1, 3}) {
		t.Error("list values error")
	}

	v, flag = a.Get(1)
	if v != 3 || !flag {
		t.Error("list get tail item error")
	}
}
