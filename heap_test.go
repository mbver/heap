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

func TestHeap_PushPop(t *testing.T) {
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

func TestHeap_EmptyHeap(t *testing.T) {
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

func TestHeap_Remove(t *testing.T) {
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
	h.Remove(&testItem{8, 19})
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
