package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	var str = "Hello, OTUS!"
	fmt.Println(reverse.String(str))
}
