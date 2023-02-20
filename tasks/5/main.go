package main

import (
	"fmt"
	"time"
)

// producer
func sendData(c chan int) { // записываем данные в канал
	for i := 0; ; i++ { 
		c <- i
	}
}

// consumer
func readData(c chan int) { // считываем данные
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	N := 1                   // количество секунд после которых программа должна завершиться
	c := make(chan int)      // создаем канал для отправки записи
	go readData(c)           // можем читать данные из канала (горутина блокируется при попытке прочитать данные из пустого канала и возобновляет действие при записи в канал)
	go sendData(c) // отпрвляем
	time.Sleep(time.Duration(N) * time.Second)
}
