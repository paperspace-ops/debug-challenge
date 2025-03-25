package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <n>")
		return
	}
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Println("fib:", fib(n))
}