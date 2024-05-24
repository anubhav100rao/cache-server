package eviction

type EvictionPolicy interface {
	RecordAccess(key string)
	Add(key string)
	Evict() string
	Remove(key string)
	Clear()
}
