package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать конкурентную запись данных в map.
*/

type SafeMap struct {
	m  map[int]string
	mu sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[int]string),
	}
}

func (sm *SafeMap) Set(key int, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key int) string {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.m[key]
}

func main() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	// Горутины для записи в map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Set(i, fmt.Sprintf("value%d", i))
			time.Sleep(time.Millisecond * 50) // Симуляция работы
		}(i)
	}

	// Горутины для чтения из map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value := safeMap.Get(i)
			fmt.Printf("Key: %d, Value: %s\n", i, value)
			time.Sleep(time.Millisecond * 50) // Симуляция работы
		}(i)
	}

	wg.Wait()
}
