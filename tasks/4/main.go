package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	dataChan := make(chan string) // делаем канал для загрузки данных
	numWorkers := 5 // количество воркеров
	var wg sync.WaitGroup
	wg.Add(numWorkers) // на каждый воркер по горутине будет

	for i := 0; i < numWorkers; i++ {
		go func() {
			for data := range dataChan {
				fmt.Println(data)
			}
			wg.Done()
		}()
	}

	sigChan := make(chan os.Signal, 1) // слушаем сигнал SIGINT для завершения программы
	signal.Notify(sigChan, syscall.SIGINT) 

	// бесконечный цикл для прослушки сигнала в бесконечном цикле
	for {
		select {
		case <-sigChan:
			close(dataChan) // закрываем канал

			
			wg.Wait() // ждем пока доработают все воркеры
			fmt.Println("Program terminated")
			return
		default:
			
			var input string // читаем данные из stdin
			fmt.Scanln(&input)
			dataChan <- input // пишем данные в канал
		}
	}
}
