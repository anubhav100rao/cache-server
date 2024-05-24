package eviction

import "container/list"

// LRUCachePolicy implements LRU eviction policy
type LRUCachePolicy struct {
	cache map[string]*list.Element
	list  *list.List
}

// NewLRUCachePolicy creates a new LRU cache policy
func NewLRUCachePolicy() *LRUCachePolicy {
	return &LRUCachePolicy{
		cache: make(map[string]*list.Element),
		list:  list.New(),
	}
}

func (p *LRUCachePolicy) RecordAccess(key string) {
	if elem, exists := p.cache[key]; exists {
		p.list.MoveToFront(elem)
	}
}

func (p *LRUCachePolicy) Add(key string) {
	elem := p.list.PushFront(key)
	p.cache[key] = elem
}

func (p *LRUCachePolicy) Evict() string {
	if p.list.Len() > 0 {
		back := p.list.Back()
		p.list.Remove(back)
		key := back.Value.(string)
		delete(p.cache, key)
		return key
	}
	return ""
}

func (p *LRUCachePolicy) Remove(key string) {
	if elem, exists := p.cache[key]; exists {
		p.list.Remove(elem)
		delete(p.cache, key)
	}
}

func (p *LRUCachePolicy) Clear() {
	p.list.Init()
	p.cache = make(map[string]*list.Element)
}
