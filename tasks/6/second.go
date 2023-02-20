package main

import (
	"context"
	"fmt"
	"time"
)

func myWorker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return // останавливаем горутину
		default:
			fmt.Println("Working...")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // можем остановить гоуртину с помощью контекста
	go myWorker(ctx)
	time.Sleep(5 * time.Second)
	cancel() // подаем сигнал для остановки
	time.Sleep(1 * time.Second) // ждем чтобы горутина успела завершиться не раньше main
	fmt.Println("Program terminated")
}
