package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// Горутина с использованием канала
func workerWithChannel(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Горутина с каналом остановлена")
			return
		default:
			fmt.Println("Работаю с каналом...")
			time.Sleep(1 * time.Second)
		}
	}
}

// Горутина с использованием контекста
func workerWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина с контекстом остановлена")
			return
		default:
			fmt.Println("Работаю с контекстом...")
			time.Sleep(1 * time.Second)
		}
	}
}

// Горутина с использованием sync.WaitGroup
func workerWithWaitGroup(wg *sync.WaitGroup, stop chan bool) {
	defer wg.Done() // Уменьшаем счетчик при завершении
	for {
		select {
		case <-stop:
			fmt.Println("Горутина с WaitGroup остановлена")
			return
		default:
			fmt.Println("Работаю с WaitGroup...")
			time.Sleep(1 * time.Second)
		}
	}
}

// Горутина с использованием флага
func workerWithFlag(running *bool) {
	for *running {
		fmt.Println("Работаю с флагом...")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Горутина с флагом остановлена")
}

func main() {
	// Реализация с использованием канала
	fmt.Println("Запуск горутины с каналом")
	stopChan := make(chan bool)
	go workerWithChannel(stopChan)

	time.Sleep(5 * time.Second) // Даем время на выполнение
	stopChan <- true            // Отправляем сигнал остановки
	time.Sleep(1 * time.Second) // Ждем немного перед завершением

	// Реализация с использованием контекста
	fmt.Println("Запуск горутины с контекстом")
	ctx, cancel := context.WithCancel(context.Background())
	go workerWithContext(ctx)

	time.Sleep(5 * time.Second) // Даем время на выполнение
	cancel()                    // Отмена контекста
	time.Sleep(1 * time.Second) // Ждем немного перед завершением

	// Реализация с использованием sync.WaitGroup
	fmt.Println("Запуск горутины с WaitGroup")
	var wg sync.WaitGroup
	stopWG := make(chan bool)
	wg.Add(1)
	go workerWithWaitGroup(&wg, stopWG)

	time.Sleep(5 * time.Second) // Даем время на выполнение
	stopWG <- true              // Отправляем сигнал остановки
	wg.Wait()                   // Ожидаем завершения горутины
	time.Sleep(1 * time.Second) // Ждем немного перед завершением

	// Реализация с использованием флага
	fmt.Println("Запуск горутины с флагом")
	running := true
	go workerWithFlag(&running)

	time.Sleep(5 * time.Second) // Даем время на выполнение
	running = false             // Установка флага в false
	time.Sleep(1 * time.Second) // Ждем немного перед завершением
}
