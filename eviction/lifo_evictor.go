package eviction

// LIFOQueuePolicy implements LIFO eviction policy
type LIFOQueuePolicy struct {
	stack []string
}

// NewLIFOQueuePolicy creates a new LIFO cache policy
func NewLIFOQueuePolicy() *LIFOQueuePolicy {
	return &LIFOQueuePolicy{
		stack: []string{},
	}
}

func (p *LIFOQueuePolicy) RecordAccess(key string) {
	// No need to record access for LIFO
}

func (p *LIFOQueuePolicy) Add(key string) {
	p.stack = append(p.stack, key)
}

func (p *LIFOQueuePolicy) Evict() string {
	if len(p.stack) > 0 {
		lastIndex := len(p.stack) - 1
		key := p.stack[lastIndex]
		p.stack = p.stack[:lastIndex]
		return key
	}
	return ""
}

func (p *LIFOQueuePolicy) Remove(key string) {
	for i, k := range p.stack {
		if k == key {
			p.stack = append(p.stack[:i], p.stack[i+1:]...)
			break
		}
	}
}

func (p *LIFOQueuePolicy) Clear() {
	p.stack = []string{}
}
