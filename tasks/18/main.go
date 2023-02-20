package main

import (
	"fmt"
	"sync"
)

type counter struct{
	mu sync.Mutex
	c uint
}

func (c *counter) inc(wg *sync.WaitGroup) {
	c.mu.Lock() // блокируем доступ, что инкрементировать могла только одна горутина
	defer c.mu.Unlock() // разблокируем доступ 
	defer wg.Done() //уменьшим счетчик горутин
	c.c+=1
}

func main() {
	var c = counter{c: 0}
	var size = 100
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		wg.Add(1) // инкрементируем счетчик горутин
		go c.inc(&wg)
	}
	wg.Wait() // ожидаем выполнения горутин, чтобы завершились не раньше потока main
	fmt.Println(c.c)
}
