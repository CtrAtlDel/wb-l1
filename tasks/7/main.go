package main

import "sync"

type concurrentMap struct {
	mu sync.RWMutex
	cm map[int]int
}

// Будем использовать mutex для корректной записи данных в map, чтобы горутины не перезаписывали значение
func (c *concurrentMap) setValue(key, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cm[key] = value
}

func main() {}
