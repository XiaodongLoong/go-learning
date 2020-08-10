package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type player struct {
	name    string
	count   int
	round   chan int
	wg	*sync.WaitGroup
}

func play(p *player) {
	defer p.wg.Done()
	for {	
		ball,ok :=  <- p.round
		if ok { // come in
			p.count=ball+1
		}else {
			fmt.Printf("player %s won\n",p.name)
			return
		}
		// miss
		n := rand.Intn(100)
		if n%13 ==0 {
			fmt.Printf("player %s miss\n",p.name)
			close(p.round)
			return
		}
		fmt.Printf("%s < %d \n",p.name,p.count)
		// hit back
		p.round <- p.count
	}
}

func main() {
	round := make(chan int)
	var wait sync.WaitGroup
	wait.Add(2)

	a :=&player {
		name: "A",
		count: 0,
		round: round,
		wg: &wait,
	}
	// player A ready
	go play(a)

	b :=&player {
                name: "B",
                count: 0,
		round: round,
		wg: &wait,
        }
        // player B ready
        go play(b)

	round <- 1

	wait.Wait() 
}
