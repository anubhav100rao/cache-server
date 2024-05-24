package eviction

import "container/list"

// FIFOCachePolicy implements FIFO eviction policy
type FIFOCachePolicy struct {
	queue *list.List
}

// NewFIFOCachePolicy creates a new FIFO cache policy
func NewFIFOCachePolicy() *FIFOCachePolicy {
	return &FIFOCachePolicy{
		queue: list.New(),
	}
}

func (p *FIFOCachePolicy) RecordAccess(key string) {
	// No need to record access for FIFO
}

func (p *FIFOCachePolicy) Add(key string) {
	p.queue.PushBack(key)
}

func (p *FIFOCachePolicy) Evict() string {
	if p.queue.Len() > 0 {
		front := p.queue.Front()
		p.queue.Remove(front)
		return front.Value.(string)
	}
	return ""
}

func (p *FIFOCachePolicy) Remove(key string) {
	for e := p.queue.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == key {
			p.queue.Remove(e)
			break
		}
	}
}

func (p *FIFOCachePolicy) Clear() {
	p.queue.Init()
}
