package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

// Функция для расчета квадрата числа и добавления к общей сумме
func calculateSquareAndAdd(num int, sum *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	// время начала вычисления
	start := time.Now()

	square := num * num

	mu.Lock()
	*sum += square
	mu.Unlock()

	elapsed := time.Since(start)
	fmt.Printf("Время выполнения для числа %d: %v наносекунд\n", num, elapsed.Nanoseconds())
}

func main() {
	// массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	var sum int
	var wg sync.WaitGroup
	var mu sync.Mutex

	//общее время
	totalStart := time.Now()

	for _, num := range numbers {
		wg.Add(1)
		go calculateSquareAndAdd(num, &sum, &mu, &wg)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	fmt.Printf("Сумма квадратов чисел: %d\n", sum)

	// Измеряем и выводим общее время выполнения программы
	totalElapsed := time.Since(totalStart)
	fmt.Printf("Общее время выполнения программы: %v наносекунд\n", totalElapsed.Nanoseconds())
}
