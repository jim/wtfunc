package main

import "fmt"

func main() {
	fmt.Println("hi")
}

type Contraption struct{}

func (c Contraption) valueReceiver() {
	fmt.Println("hi")
}

func (c *Contraption) pointerReceiver() {
	fmt.Println("hi")
}
