package heap

type Item interface {
	ID() uint64
	Less(Item) bool
}

type Heap struct {
	items     []Item
	positions map[uint64]int
}

func NewHeap() *Heap {
	return &Heap{
		items:     []Item{},
		positions: map[uint64]int{},
	}
}

func (h *Heap) Push(item Item) {
	if item == nil {
		return
	}
	if _, ok := h.positions[item.ID()]; ok {
		h.Remove(item.ID())
	}
	h.items = append(h.items, item)
	h.positions[item.ID()] = h.Len() - 1
	h.siftUp(h.Len() - 1)
}

func (h *Heap) siftUp(i int) {
	if i >= h.Len() {
		return
	}
	for i > 0 {
		j := (i - 1) / 2
		if h.items[i].Less(h.items[j]) {
			h.swap(i, j)
			i = j
		} else {
			return
		}
	}
}

func (h *Heap) Len() int {
	return len(h.items)
}

func (h *Heap) swap(i, j int) {
	h.positions[h.items[i].ID()] = j
	h.positions[h.items[j].ID()] = i
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap) Remove(id uint64) {
	p, ok := h.positions[id]
	if !ok {
		return
	}
	h.swap(p, h.Len()-1)
	h.items = h.items[:h.Len()-1]
	delete(h.positions, id)
	h.siftDown(p)
	h.siftUp(p)
}

func (h *Heap) siftDown(i int) {
	j := (i+1)*2 - 1
	for j < h.Len() {
		j1 := (i + 1) * 2
		if j1 < h.Len() && h.items[j1].Less(h.items[j]) {
			j = j1
		}
		if h.items[j].Less(h.items[i]) {
			h.swap(i, j)
			i = j
			j = (i+1)*2 - 1
		} else {
			return
		}
	}
}

func (h *Heap) Pop() Item {
	if h.Len() == 0 {
		return nil
	}
	result := h.items[0] // nil is not inserted
	h.Remove(result.ID())
	return result
}

func (h *Heap) Peek() Item {
	if h.Len() == 0 {
		return nil
	}
	return h.items[0]
}
