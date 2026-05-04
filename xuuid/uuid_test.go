package xuuid

import (
	"fmt"
	"testing"
)

func TestNewUUID(t *testing.T) {
	id := NewUUID()
	fmt.Println(id)
	if id == "" {
		t.Fatal("CreateUUID fail")
	}
}

func BenchmarkNewUUID(b *testing.B) {
	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewUUID()
	}
}
