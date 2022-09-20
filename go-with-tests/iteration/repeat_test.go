package iteration

import (
	"fmt"
	"testing"
)

// repeat character in count fnd return new string
func TestRepeat(t *testing.T) {
	repeated := Repeat("b", 5)
	expected := "bbbbb"

	if repeated != expected {
		t.Errorf("expected %q but got %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("п", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("g", 6)
	fmt.Println(repeated)
	// Output: gggggg
}
