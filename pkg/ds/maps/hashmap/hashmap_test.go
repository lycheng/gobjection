package hashmap

import (
	"reflect"
	"sort"
	"testing"

	"github.com/lycheng/gobjection/pkg/ds/maps"
)

func assertSortedIntListEquals(values []interface{}, expected []int) bool {
	rv := make([]int, 0)
	for _, i := range values {
		rv = append(rv, i.(int))
	}
	sort.Ints(rv)
	return reflect.DeepEqual(rv, expected)
}

func TestMapsImplementation(t *testing.T) {
	var _ maps.Map = (*Map)(nil)
}

func TestHashMap(t *testing.T) {
	m := New()
	m.Put(1, 1)

	if !m.Has(1) || m.Size() != 1 {
		t.Error("hashmap put error")
	}

	v, f := m.Get(1)
	if v.(int) != 1 || !f {
		t.Error("hashmap get error")
	}

	m.Remove(1)
	v, f = m.Get(1)
	if v != nil || f {
		t.Error("hashmap remove error")
	}

	if m.Size() != 0 || !m.IsEmpty() {
		t.Error("hashmap size and IsEmpty checks error")
	}
}

func TestHashMapKeysAndValues(t *testing.T) {
	m := New()
	m.Put(1, 10)
	m.Put(2, 9)
	m.Put(3, 8)

	if !assertSortedIntListEquals(m.Keys(), []int{1, 2, 3}) {
		t.Error("hashmap keys error")
	}

	if !assertSortedIntListEquals(m.Values(), []int{8, 9, 10}) {
		t.Error("hashmap values error")
	}
}

func TestHashMapIterator(t *testing.T) {
	m := New()
	m.Put(1, 2)

	iter := m.Iterator()
	v := iter.Next()
	if v == nil {
		t.Error("hashmap iterator next is nil")
	}

	kv := v.(*maps.KV)
	if kv.Key().(int) != 1 || kv.Value().(int) != 2 {
		t.Error("hashmap kv value error")
	}

	v = iter.Next()
	if v != nil {
		t.Error("hashmap iterator next is not nil")
	}
}
