package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("four times", func(t *testing.T) {
		repeated := Repeat("a", 4)
		expected := strings.Repeat("a", 4)
		assertRepetition(repeated, expected, t)
	})

	t.Run("two times", func(t *testing.T) {
		repeated := Repeat("x", 2)
		expected := "xx"
		assertRepetition(repeated, expected, t)
	})
}

func assertRepetition(repeated string, expected string, t *testing.T) {
	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("!", 3)
	fmt.Println(repeated)
	// Output: !!!
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
