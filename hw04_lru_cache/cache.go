package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Len() int
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (ci *cacheItem) GetKey() Key {
	return ci.key
}

func (ci *cacheItem) GetValue() interface{} {
	return ci.value
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, ok := l.items[key]
	if ok {
		l.items[key].Value = &cacheItem{key, value}
		l.queue.MoveToFront(l.items[key])

		return true
	}

	item := l.queue.PushFront(&cacheItem{key, value})
	l.items[key] = item
	if l.queue.Len() <= l.capacity {
		return false
	}

	back := l.queue.Back()
	ci, ok := back.Value.(*cacheItem)
	if !ok {
		panic("invalidate cache value")
	}
	delete(l.items, ci.key)
	l.queue.Remove(back)

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	value, ok := l.items[key]
	if !ok {
		return nil, false
	}

	l.queue.MoveToFront(l.items[key])
	ci, ok := value.Value.(*cacheItem)
	if !ok {
		panic("invalidate cache value")
	}

	return ci.GetValue(), true
}

func (l *lruCache) Len() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.queue.Len()
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.items = map[Key]*ListItem{}
	l.queue.Clear()
}
