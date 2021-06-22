package main

import (
	"fmt"

	"github.com/asasmoyo/demo-golang/mymath"
)

func main() {
	var a, b = 10, 20
	c := mymath.Add(a, b)
	fmt.Println("hasil:", c)
}
