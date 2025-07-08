package pointers

import (
	"fmt"
	"runtime"
	"time"
)

type Big struct {
	data [1024]int64 // ~8KB
}

// Value receiver (copies entire struct)
func updateByValue(b Big) {
	b.data[0]++
	b.data[512] += 255
	b.data[1023] += 1024
}

// Pointer receiver (modifies in-place)
func (b *Big) updateByPointer() {
	b.data[0]++
	b.data[512] += 255
	b.data[1023] += 1024
}

func Benchmark() {
	var m runtime.MemStats

	b1 := Big{}
	b2 := Big{}

	runtime.ReadMemStats(&m)
	fmt.Printf("Before: Alloc = %v KB\n", m.Alloc/1024)

	// Value benchmark
	start := time.Now()
	for i := 0; i < 100_000; i++ {
		updateByValue(b1)
		updateByValue(b1)
		updateByValue(b1)
	}
	fmt.Println("Value time:", time.Since(start))

	runtime.ReadMemStats(&m)
	fmt.Printf("Mid: Alloc = %v KB\n", m.Alloc/1024)

	// Pointer benchmark
	start = time.Now()
	for i := 0; i < 100_000; i++ {
		b2.updateByPointer()
		b2.updateByPointer()
		b2.updateByPointer()
	}
	fmt.Println("Pointer time:", time.Since(start))

	runtime.ReadMemStats(&m)
	fmt.Printf("After: Alloc = %v KB\n", m.Alloc/1024)
}
