package async

import (
	"fmt"
	"time"
)

func Termeric(id int) {
	fmt.Println("termeric", id)
}

func Goroutines() {
	go Termeric(1)
	go Termeric(2)
	go Termeric(3)
	go Termeric(4)

	time.Sleep(1 * time.Second) // Wait for goroutines to finish
	fmt.Println("Goroutines finished")
}
