package hw04lrucache

import "sync"

// Key ...
type Key string

// Cache ...
type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

// NewCache ...
func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	if key == "" || value == nil {
		return false
	}
	lc.mu.Lock()
	defer lc.mu.Unlock()
	if v, ok := lc.items[key]; ok {
		v.Value = value
		lc.queue.MoveToFront(v)
		return true
	}
	lc.items[key] = lc.queue.PushFront(value)
	if lc.queue.Len() > lc.capacity {
		pntr := lc.queue.Back()
		for k, v := range lc.items {
			if v == pntr {
				delete(lc.items, k)
				break
			}
		}
		lc.queue.Remove(pntr)
	}

	return false
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	if key == "" {
		return nil, false
	}
	lc.mu.Lock()
	defer lc.mu.Unlock()
	if v, ok := lc.items[key]; ok {
		lc.queue.MoveToFront(v)
		return v.Value, ok
	}
	return nil, false
}

func (lc *lruCache) Clear() {
	lc.mu.Lock()
	lc.queue = NewList()
	lc.items = make(map[Key]*ListItem, lc.capacity)
	lc.mu.Unlock()
}
