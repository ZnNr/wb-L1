package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают
//произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//
//
//Программа должна завершаться по нажатию Ctrl+C.
//Выбрать и обосновать способ завершения работы всех воркеров..

// Функция receive генерирует последовательные числа и отправляет их в канал ingest.
func receive(ingest chan<- int) {
	defer close(ingest) // Закрываем канал после завершения работы функции
	i := 0
	for {
		ingest <- i                        // Отправляем текущее число в канал
		i++                                // Увеличиваем число на 1
		time.Sleep(100 * time.Millisecond) // Пауза перед следующей отправкой
	}
}

// Функция starting читает числа из канала ingest и передает их в канал jobs для обработки воркерами.
func starting(ctx context.Context, ingest <-chan int, jobs chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении функции
	for {
		select {
		case num, ok := <-ingest: // Читаем данные из канала ingest
			if !ok { // Если канал закрыт
				return // Выход из функции
			}
			jobs <- num // Отправляем число в канал jobs
		case <-ctx.Done(): // Проверяем, был ли вызван контекст cancel
			return // Выход из функции при завершении контекста
		}
	}
}

// Функция worker обрабатывает числа из канала jobs.
func worker(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()         // Уменьшаем счетчик WaitGroup при завершении функции
	for num := range jobs { // Перебираем все числа в канале jobs
		fmt.Println("Processed:", num) // Выводим обработанное число
		time.Sleep(1 * time.Second)    // Эмуляция обработки (пауза в 1 секунду)
	}
}

// Основная функция, запускающая всех воркеров и управляемая контекстом.
func main() {
	if len(os.Args) < 2 { // Проверка наличия аргументов командной строки
		fmt.Println("Please provide the number of workers") // Сообщение об ошибке при отсутствии аргументов
		return
	}
	workerCount, err := strconv.Atoi(os.Args[1]) // Преобразование аргумента в целое число
	if err != nil || workerCount <= 0 {          // Проверка на корректность числа воркеров
		fmt.Println("Invalid number of workers") // Сообщение об ошибке
		return
	}

	ingest := make(chan int) // Канал для отправки данных
	jobs := make(chan int)   // Канал для обработки данных
	var wg sync.WaitGroup    // WaitGroup для ожидания завершения всех горутин

	ctx, cancel := context.WithCancel(context.Background()) // Создание контекста для управления задержкой завершения

	go receive(ingest) // Запуск функции receive как горутины

	wg.Add(1)                           // Увеличиваем счетчик WaitGroup для функции starting
	go starting(ctx, ingest, jobs, &wg) // Запуск функции starting как горутины

	for i := 0; i < workerCount; i++ { // Запуск указанного количества воркеров
		wg.Add(1)            // Увеличиваем счетчик WaitGroup для каждого воркера
		go worker(jobs, &wg) // Запуск функции worker как горутины
	}

	// Настройка канала для обнаружения сигналов завершения
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan // Ожидание сигнала завершения
	cancel()     // Отменяем контекст для уведомления воркеров о завершении

	close(jobs) // Закрываем канал jobs, чтобы оповестить воркеров о завершении

	wg.Wait()                                              // Ждем завершения всех воркеров
	fmt.Println("All workers have completed their tasks.") // Финальное сообщение
}
