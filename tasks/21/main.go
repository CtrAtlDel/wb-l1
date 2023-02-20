package main

import "fmt"

type Target interface { // Интерфейс, которому будет соответствовать адаптер.
	Request() string
}

type Adaptee struct { // Адаптер, который мы хотим применить к  интерфейсу.
	AdapteeRequest string
}

func (a *Adaptee) SpecificRequest() string {
	return a.AdapteeRequest
}

type Adapter struct { // Адаптер, который оборачивает адаптер и реализует интерфейс.
	Adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{AdapteeRequest: "Adaptee request"}
	adapter := &Adapter{Adaptee: adaptee}

	fmt.Println(adapter.Request()) // "Adaptee request"
}
