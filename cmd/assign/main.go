package main

import "fmt"

func main() {
	a := []int{1,2,3}
	for x := range a {
		fmt.Println(x)
		fmt.Println(&x)
		x := x
		fmt.Println(&x)
	}
}
