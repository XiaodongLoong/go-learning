package cgo

/*
#include <stdio.h>
void printHello() {
     printf("hello world!");
}
*/
import "C"

func MyPrint() {
	C.printHello()
}
