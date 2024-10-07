package heap

import (
	"testing"
)

type testItem struct {
	id    int
	value int
}

func (i *testItem) ID() int {
	return i.id
}

func (i *testItem) Less(item Item) bool {
	i1 := item.(*testItem)
	return i.value < i1.value
}

func Test_PushPop(t *testing.T) {
	items := []*testItem{
		{1, 10},
		{2, 30},
		{3, 2},
		{4, 5},
		{6, 1},
		{7, 21},
		{8, 19},
		{9, 15},
		{10, 27},
	}
	expected := []*testItem{
		{6, 1},
		{3, 2},
		{4, 5},
		{1, 10},
		{9, 15},
		{8, 19},
		{7, 21},
		{10, 27},
		{2, 30},
	}

	h := NewHeap()
	for _, item := range items {
		h.Push(item)
	}
	for _, item := range expected {
		item1 := h.Pop().(*testItem)
		if item1.id != item.id || item1.value != item.value {
			t.Errorf("pop item: expected: %v, got: %v", item, item1)
		}
	}
}

func Test_EmptyHeap(t *testing.T) {
	h := NewHeap()
	item := h.Pop()
	if item != nil {
		t.Errorf("empty heap expect nil, got: %v", item)
	}
	item = h.Peek()
	if item != nil {
		t.Errorf("empty heap expect nil, got: %v", item)
	}
}

func Test_Remove(t *testing.T) {
	items := []*testItem{
		{1, 10},
		{2, 30},
		{3, 2},
		{4, 5},
		{6, 1},
		{7, 21},
		{8, 19},
		{9, 15},
		{10, 27},
	}

	expected := []*testItem{
		{6, 1},
		{3, 2},
		{4, 5},
		{1, 10},
		{9, 15},
		{7, 21},
		{10, 27},
		{2, 30},
	}
	h := NewHeap()
	for _, item := range items {
		h.Push(item)
	}
	h.Remove(8)
	if h.Len() != len(items)-1 || len(h.positions) != len(items)-1 {
		t.Errorf("expect heap size: %d, got: %d, %d", len(items)-1, h.Len(), len(h.positions))
	}
	for _, item := range expected {
		item1 := h.Pop().(*testItem)
		if item1.id != item.id || item1.value != item.value {
			t.Errorf("pop item: expected: %v, got: %v", item, item1)
		}
	}
}

func Test_NilPushRemove(t *testing.T) {
	h := NewHeap()
	h.Push(nil)
	if h.Len() != 0 {
		t.Errorf("expect nil is not inserted")
	}
	h.Remove(0)
}

func Test_SameIDs(t *testing.T) {
	h := NewHeap()
	for i := 0; i < 3; i++ {
		h.Push(&testItem{id: 0, value: 10 + i})
	}
	if h.Len() != 1 {
		t.Errorf("item of same id should replace old item, expect: %d, got: %d", 1, h.Len())
	}
}
