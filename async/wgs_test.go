package async

import "testing"

func BenchmarkHeavyAlloc(b *testing.B) {
	for b.Loop() {
		WaitGroups()
	}
}
