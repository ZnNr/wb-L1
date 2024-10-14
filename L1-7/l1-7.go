package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать конкурентную запись данных в map.
*/

type SafeMap struct { //структура мапы с мьютексом, который позволяет провести конкурентную запись в мап, так как мап для записи не потокобезопасен
	m  map[int]string //m — сама мапа, где данные будут храниться.
	mu sync.Mutex
}

// Cоздаёт и возвращает указатель на новый экземпляр SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[int]string),
	}
}

// Записывает пару "ключ-значение" в  мапу
func (sm *SafeMap) Set(key int, value string) { //запись в мапу
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

// возвращает значение по ключу
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
	for i := 0; i < 10; i++ { ///Создаются 10 горутин (по одной на каждое значение i).
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value := safeMap.Get(i)
			fmt.Printf("Key: %d, Value: %s\n", i, value)
			time.Sleep(time.Millisecond * 50) // Симуляция работы
		}(i)
	}

	wg.Wait() // Блокирует выполнение основной функции до завершения всех горутин, что гарантирует, что все записи и чтения завершены до выхода из программы.
}
