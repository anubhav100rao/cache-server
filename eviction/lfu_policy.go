package eviction

import (
	"container/heap"
)

// LFUCacheItem represents an item in the LFU cache with its frequency count.
type LFUCacheItem struct {
	key       string
	frequency int
	index     int
}

// LFUCachePolicy implements the EvictionPolicy interface for LFU.
type LFUCachePolicy struct {
	items     map[string]*LFUCacheItem
	frequency *LFUHeap
}

// NewLFUCachePolicy creates a new LFU eviction policy.
func NewLFUCachePolicy() *LFUCachePolicy {
	return &LFUCachePolicy{
		items:     make(map[string]*LFUCacheItem),
		frequency: &LFUHeap{},
	}
}

func (p *LFUCachePolicy) RecordAccess(key string) {
	if item, exists := p.items[key]; exists {
		p.frequency.update(item, item.frequency+1)
	}
}

func (p *LFUCachePolicy) Add(key string) {
	item := &LFUCacheItem{key: key, frequency: 1}
	heap.Push(p.frequency, item)
	p.items[key] = item
}

func (p *LFUCachePolicy) Evict() string {
	if p.frequency.Len() == 0 {
		return ""
	}
	item := heap.Pop(p.frequency).(*LFUCacheItem)
	delete(p.items, item.key)
	return item.key
}

func (p *LFUCachePolicy) Remove(key string) {
	if item, exists := p.items[key]; exists {
		heap.Remove(p.frequency, item.index)
		delete(p.items, key)
	}
}

func (p *LFUCachePolicy) Clear() {
	p.items = make(map[string]*LFUCacheItem)
	p.frequency = &LFUHeap{}
}

// LFUHeap is a min-heap based on the frequency of cache items.
type LFUHeap []*LFUCacheItem

func (h LFUHeap) Len() int           { return len(h) }
func (h LFUHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h LFUHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *LFUHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*LFUCacheItem)
	item.index = n
	*h = append(*h, item)
}

func (h *LFUHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

func (h *LFUHeap) update(item *LFUCacheItem, frequency int) {
	item.frequency = frequency
	heap.Fix(h, item.index)
}
