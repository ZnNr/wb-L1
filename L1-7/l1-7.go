package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать конкурентную запись данных в map.

*/

/*
Конкурентная запись в map в языке Go может вызвать проблемы с гонками данных, поскольку стандартные map не являются потокобезопасными. Чтобы реализовать конкурентную запись в map, можно использовать sync.Mutex для синхронизации доступа к map. Также можно использовать sync.RWMutex, если вы хотите позволить многим читателям одновременно получать доступ к map, но иметь эксклюзивный доступ для записей.

Вот пример реализации с использованием sync.Mutex:

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

/*
Объяснение кода:
SafeMap: Мы создаем структуру, которая содержит map и sync.Mutex. Методы Set и Get синхронизируют доступ к map.
Set: Этот метод блокирует мьютекс перед записью в map, гарантируя, что одновременно только одна горутина может изменять его.
Get: Этот метод также блокирует мьютекс для чтения из map, чтобы обеспечить консистентность данных.
Горутины: Мы создаем несколько горутин для записи и чтения данных из SafeMap. Каждая запись и чтение выполняются в отдельной горутине.
sync.WaitGroup: Используется для ожидания завершения всех горутин.
Этот подход гарантирует, что доступ к map будет выполнен безопасно в конкурентной среде.

*/
