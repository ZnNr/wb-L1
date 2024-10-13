package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно
отправлять значения в канал,
а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func producer(ch chan<- int, duration time.Duration) {
	num := 0
	startTime := time.Now()

	for time.Since(startTime) < duration {
		ch <- num // Отправляем значение в канал
		fmt.Printf("Производитель отправил: %d\n", num)
		num++
		time.Sleep(1 * time.Second) // Ждем 1 секунду перед отправкой следующего значения
	}
	close(ch) // Закрываем канал, когда закончено
}

func consumer(ch <-chan int, duration time.Duration) {
	startTime := time.Now()

	for time.Since(startTime) < duration {
		select {
		case item, ok := <-ch:
			if ok {
				fmt.Printf("Потребитель получил: %d\n", item)
			} else {
				return // Если канал закрыт, выходим из функции
			}
		case <-time.After(1 * time.Second):
			// Ждем 1 секунду, если ничего нет в канале
			continue
		}
	}
}

func main() {
	duration := 10 * time.Second // Установка времени работы программы
	ch := make(chan int)

	go producer(ch, duration) // Запуск горутины производителя
	go consumer(ch, duration) // Запуск горутины потребителя

	time.Sleep(duration) // Ждем завершения программы
	fmt.Println("Программа завершена.")
}
