package channel

import (
	"fmt"
	"time"
)

func TimeoutTest() {
	var c <-chan int
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}

}
