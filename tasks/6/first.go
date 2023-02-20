package main

import (
	"fmt"
	"time"
)

func checkChan(c chan bool) {
	for { // Выполняется бесконечный цикл, из которого можно выйти с помощью return, если пришел нужный флаг
		select {
		case <-c:
			fmt.Println("Stop gorutione...")
			return // останавливаем горутину
		default:
			fmt.Println("Its working...")
			time.Sleep(1 * time.Second)
		}
	}

}

func main() {
	c := make(chan bool) // Остановка горутины с помощью флага из канала
	go checkChan(c)
	time.Sleep(5 * time.Second) // через 5 секунд в горутину отправялется флаг
	c <- true
	time.Sleep(2 * time.Second) // добавим еще время чтобы main поток не завершился раньше выполнения горутины
}
