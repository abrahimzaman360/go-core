package pointers

import (
	"fmt"
	"runtime"
)

func Pointers() {
	var m runtime.MemStats


	// Pointer Example
	counter := 0
	ptr_increment(&counter)
	ptr_increment(&counter)
	ptr_increment(&counter)
	fmt.Printf("%d\n", counter)




	// Before
    runtime.ReadMemStats(&m)
    fmt.Printf("Before: Alloc = %v KB\n", m.Alloc/1024)

	// Non Ptr example:
	counter_ := 0
	counter_ = increment(counter_)
	counter_ = increment(counter_)
	counter_ = increment(counter_)
	fmt.Printf("%d\n", counter_)


    // After
    runtime.ReadMemStats(&m)
    fmt.Printf("After: Alloc = %v KB\n", m.Alloc/1024)
}

func ptr_increment(counter *int) {
	*counter += 1
	fmt.Printf("Pointer Address: %v\n", counter);
}

func increment(counter int) (int) {
	counter += 1
	fmt.Printf("Non-Pointer Address: %v\n", &counter);
	return counter
}