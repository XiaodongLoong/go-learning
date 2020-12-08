package main

import (
	"fmt"
	"os"
)

func foo(a ,b int) int {
	return a+b
}

func max(a, b  int) int {
	if a>=b {
		return a
	}else  {
		return b
	}
}

func main() {
	fmt.Println(os.Getenv("LOONG"))
}
