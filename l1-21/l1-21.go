package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// AudioPlayer - целевой интерфейс
type AudioPlayer interface {
	Play(audioType string, fileName string)
}

// Mp3Player - реализация для MP3 формата
type Mp3Player struct{}

// Play - реализация метода воспроизведения для MP3
func (mp *Mp3Player) Play(audioType string, fileName string) {
	fmt.Printf("Воспроизведение MP3 файла: %s\n", fileName)
}

// WavPlayer - класс, который мы должны адаптировать
type WavPlayer struct{}

// PlayWav - метод воспроизведения для WAV формата
func (wp *WavPlayer) PlayWav(fileName string) {
	fmt.Printf("Воспроизведение WAV файла: %s\n", fileName)
}

// WavAdapter - адаптер для WavPlayer
type WavAdapter struct {
	wavPlayer *WavPlayer
}

// Play - метод, который адаптирует вызов для WAV формата
func (wa *WavAdapter) Play(audioType string, fileName string) {
	if audioType == "wav" {
		wa.wavPlayer.PlayWav(fileName)
	} else {
		fmt.Printf("Невозможно воспроизвести файл формата: %s\n", audioType)
	}
}

// AudioAdapter - основной аудиоплеер, комбинирующий разные игроки
type AudioAdapter struct {
	mp3Player  *Mp3Player
	wavAdapter *WavAdapter
}

// Play - метод для воспроизведения аудиофайлов
func (aa *AudioAdapter) Play(audioType string, fileName string) {
	if audioType == "mp3" {
		aa.mp3Player.Play(audioType, fileName)
	} else if audioType == "wav" {
		aa.wavAdapter.Play(audioType, fileName)
	} else {
		fmt.Println("Неподдерживаемый формат аудиа.")
	}
}

func main() {
	mp3Player := &Mp3Player{}
	wavPlayer := &WavPlayer{}
	wavAdapter := &WavAdapter{wavPlayer: wavPlayer}
	audioAdapter := &AudioAdapter{mp3Player: mp3Player, wavAdapter: wavAdapter}

	audioAdapter.Play("mp3", "song.mp3")
	audioAdapter.Play("wav", "track.wav")
	audioAdapter.Play("ogg", "audio.ogg")
}
