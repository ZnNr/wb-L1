package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают
произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.


Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров..
*/

// Функция receive отправляет целые числа в канал ingest с задержкой 100 мс между отправками.
// Цель - Вечный цикл будет генерировать последовательные целые числа.
func receive(ingest chan<- int) {
	defer close(ingest) // Закрываем канал после завершения работы функции
	i := 0
	for {
		ingest <- i                        // Отправляем текущее число в канал
		i++                                // Увеличиваем число на 1
		time.Sleep(100 * time.Millisecond) // Пауза перед следующей отправкой
	}
}

// Функция starting принимает данные из канала ingest и пересылает их в канал jobs.
// Если контекст завершен, то закрываем канал jobs и прекращаем выполнение.
// цель функции - Пересылка данных из канала ingest в канал jobs с учетом контекста завершения.
func starting(ctx context.Context, jobs chan int, ingest chan int) {
	for {
		select {
		case job := <-ingest:
			jobs <- job // Читаем данные из канала ingest
		case <-ctx.Done(): // Передаем данные из канала ingest в канал jobs.
			fmt.Println("closing jobs")
			close(jobs) // Закрываем канал jobs, если контекст завершен.
			fmt.Println("closed jobs")
			return
		}
	}
}

// Функция worker обрабатывает задания из канала jobs.
// После завершения работы (пустой канал) выводит сообщение о прерывании.
func worker(id int, wg *sync.WaitGroup, jobs chan int) {
	defer wg.Done() // Уменьшаем счетчик горутины по завершении.
	for j := range jobs {
		fmt.Printf("Worker %d processed job %d\n", id, j)
		time.Sleep(time.Second) // Имитируем обработку задания.
	}
	fmt.Printf("Worker %d interrupted\n", id)
}
func main() {
	var numWorkers int
	fmt.Print("Please provide the number of workers: ")

	// Пользователь задает количество воркеров.
	if _, err := fmt.Scan(&numWorkers); err != nil || numWorkers < 1 {
		fmt.Println("Invalid number of workers")
		return
	}

	var wg sync.WaitGroup
	sigs := make(chan os.Signal, 1)
	ingest := make(chan int, 1)
	jobs := make(chan int, 100)
	ctx, cancel := context.WithCancel(context.Background())

	// Функция signal.Notify регистрирует получение сигналов SIGINT и SIGTERM.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Запуск горутин для приема и распределения заданий.
	//Функции receive и starting запускаются как отдельные горутины.
	go receive(ingest)
	go starting(ctx, jobs, ingest)

	// Запуск указанного пользователем количества воркеров.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs)
	}

	// Блокируем основной поток до получения сигнала завершения.
	<-sigs    //Программа ждет получения сигнала SIGINT или SIGTERM для завершения всех процессов.
	cancel()  // Отмена контекста, что завершает работу функции starting.
	wg.Wait() // Ожидание завершения всех воркеров.
	fmt.Println("All workers have completed their tasks.")
}
