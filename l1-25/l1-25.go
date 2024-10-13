package main

import "time"

/*
Реализовать собственную функцию sleep.
*/

// Sleep приостанавливает выполнение на заданное количество миллисекунд
func Sleep(milliseconds int) {
	// Создаем отдельный канал для указания завершения
	done := make(chan struct{})

	// Запускаем горутину, которая завершится через заданное время
	go func() {
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)
		close(done) // Закрываем канал по истечении времени
	}()

	// Ожидаем завершения горутины
	<-done
}

func main() {
	// Пример использования нашей функции Sleep
	println("Начинаю спать...")
	Sleep(2000) // спим 2000 миллисекунд (2 секунды)
	println("Пробуждение!")
}
