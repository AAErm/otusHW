package hw04lrucache

import "sync"

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый элемент списка
	Back() *ListItem                   // последний элемент списка
	PushFront(v interface{}) *ListItem // добавить значение в начало
	PushBack(v interface{}) *ListItem  // добавить значение в конец
	Remove(i *ListItem)                // удалить элемент
	MoveToFront(i *ListItem)           // переместить элемент в начало
	Clear()
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first  *ListItem
	end    *ListItem
	length int
	mu     sync.RWMutex
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.length
}

func (l *list) Front() *ListItem {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.first
}

func (l *list) Back() *ListItem {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.end
}

// func (l *list) PushFront(v interface{}) *ListItem {
// 	l.mu.Lock()
// 	defer l.mu.Unlock()

// 	item := &ListItem{Value: v}
// 	l.length++
// 	if l.length == 1 {
// 		l.first = item
// 		l.end = item

// 		return item
// 	}

// 	item.Next = l.first
// 	l.first.Prev = item
// 	l.first = item

// 	return item§
// }

func (l *list) PushFront(v interface{}) *ListItem {
	l.mu.Lock()
	defer l.mu.Unlock()

	item := &ListItem{Value: v, Next: l.first}
	if l.first != nil {
		l.first.Prev = item
	}
	l.first = item
	if l.end == nil {
		l.end = item
	}
	l.length++
	return item
}

// func (l *list) PushBack(v interface{}) *ListItem {
// 	l.mu.Lock()
// 	defer l.mu.Unlock()

// 	item := &ListItem{Value: v}
// 	l.length++
// 	if l.length == 1 {
// 		l.first = item
// 		l.end = item

// 		return item
// 	}

// 	item.Prev = l.end
// 	l.end.Next = item
// 	l.end = item

// 	return item
// }

func (l *list) PushBack(v interface{}) *ListItem {
	l.mu.Lock()
	defer l.mu.Unlock()

	item := &ListItem{Value: v, Prev: l.end}
	if l.end != nil {
		l.end.Next = item
	}
	l.end = item
	if l.first == nil {
		l.first = item
	}
	l.length++
	return item
}

func (l *list) Remove(i *ListItem) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.first = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.end = i.Prev
	}
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if i == l.first {
		return
	}

	if l.length == 1 {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.first = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.end = i.Prev
	}

	i.Prev = nil
	i.Next = l.first
	l.first = i

	if l.first.Next != nil {
		l.first.Next.Prev = l.first
	}
}

func (l *list) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.first = nil
	l.end = nil
	l.length = 0
}
