package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	players := 4

	for i := range 4 {
		go waitPlayer(cond, &players, i)
		time.Sleep(1 * time.Second)
	}

}

func waitPlayer(cond *sync.Cond, player *int, id int) {
	cond.L.Lock()

	*player -= 1

	if *player == 0 {
		cond.Broadcast()
	}

	if *player > 0 {
		cond.Wait()
		// fmt.Println("unlock0", id)
	}
	fmt.Println("unlock", id)
	cond.L.Unlock()
	fmt.Println("unlock1", id)
}
