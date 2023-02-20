package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	<-time.After(t) // функция блокируется до тех пор, пока будет не получено текущее значение, а получено оно будет через t секунд
}

func getSleep() {
	fmt.Println("I am sleeping...")
	sleep(5 * time.Second) // для примера взято 5 секунд
	fmt.Println("I am wakeup")
}

func main() {
	getSleep()
}
