package main

import "fmt"

func main() {
	text := "Hello, 世界!"
	for i, c := range text {
		fmt.Println(i, c, string(c+3), string(100000+i))
	}
}
