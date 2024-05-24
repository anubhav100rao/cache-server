package eviction

import (
	"math/rand"
	"time"
)

type RandomEvictionPolicy struct {
	keys []string
	rand *rand.Rand
}

func NewRandomEvictionPolicy() *RandomEvictionPolicy {
	return &RandomEvictionPolicy{
		keys: make([]string, 0),
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (p *RandomEvictionPolicy) RecordAccess(key string) {
	// No need to record access for Random policy
}

func (p *RandomEvictionPolicy) Add(key string) {
	p.keys = append(p.keys, key)
}

func (p *RandomEvictionPolicy) Evict() string {
	if len(p.keys) == 0 {
		return ""
	}
	index := p.rand.Intn(len(p.keys))
	evictKey := p.keys[index]
	p.keys = append(p.keys[:index], p.keys[index+1:]...)
	return evictKey
}

func (p *RandomEvictionPolicy) Remove(key string) {
	for i, k := range p.keys {
		if k == key {
			p.keys = append(p.keys[:i], p.keys[i+1:]...)
			break
		}
	}
}

func (p *RandomEvictionPolicy) Clear() {
	p.keys = make([]string, 0)
}
