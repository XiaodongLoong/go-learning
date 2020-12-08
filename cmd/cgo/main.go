package main

/*
#include <stdio.h>
void printHello() {
     printf("hello world!");
}
*/
import "C"

func main() {
	C.printHello()
}
