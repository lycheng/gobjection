package hashset

import "testing"

func TestSet(t *testing.T) {
	set := New()
	set.Add(1)

	if set.Size() != 1 {
		t.Error("set size not equals to 1")
	}

	if !set.Has(1) {
		t.Error("set not contains 1")
	}

	set.Add(2)
	if set.Size() != 2 {
		t.Error("set size not equals to 1")
	}

	set.Add(2)
	if set.Size() != 2 {
		t.Error("duplicate item")
	}

	set.Remove(2)
	if set.Size() != 1 || set.Has(2) {
		t.Error("remove item failed")
	}

	set.Remove(1)
	if set.Size() != 0 || !set.IsEmpty() {
		t.Error("clear item failed")
	}
}

func TestSetIterator(t *testing.T) {
	set := New()
	set.Add(1, 2, 3)

	rv := make(map[int]int)
	sum := 0

	iter := set.Iterator()
	for v := iter.Next(); v != nil; v = iter.Next() {
		rv[v.(int)] = 1
		sum += v.(int)
	}

	if len(rv) != 3 || sum != 6 {
		t.Error("invalid iterator values")
	}
}
