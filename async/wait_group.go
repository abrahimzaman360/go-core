package async

import (
	"fmt"
	"sync"
)

func WaitGroups() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
	}()
}
