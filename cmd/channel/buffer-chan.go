package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(tasks chan string) {
	defer wg.Done()
	
	for {
		t,ok := <-tasks
		if !ok {
			fmt.Printf("task is null \n")
			return
		}
		fmt.Printf("work start : %s \n",t)
		
		time.Sleep(2*time.Second)

		fmt.Printf("work complete : %s \n",t)
	}
}

func main() {
	tasks := make(chan string, 10)
	wg.Add(1)
	
	go worker(tasks)
	
	for post := 1; post <= 10; post++ {
		tasks <- fmt.Sprintf("task id:%d", post)
	}
	fmt.Printf("task post complete \n")

	close(tasks)
	wg.Wait()
}
