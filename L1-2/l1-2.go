package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Написать программу, которая конкурентно рассчитает
значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout.
*/

// Функция для расчета квадратов чисел
func calculateSquare(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	square := num * num
	fmt.Printf("Время: %s | Квадрат числа %d равен %d\n", time.Now().Format("15:04:05.999999999"), num, square)
}

func main() {
	// массив чисел
	numbers := []int{2, 4, 6, 8, 10}
	start := time.Now()

	// Инициализируем WaitGroup
	var wg sync.WaitGroup

	// Запускаем горутину для каждого числа в массиве
	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(&wg, num)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// замеряем общее время после ожидапния горутин
	elapsedTime := time.Since(start)
	fmt.Printf("Общее время выполнения: %d ms или %d ns\n", elapsedTime.Milliseconds(), elapsedTime.Nanoseconds())
}
