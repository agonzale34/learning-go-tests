package iterations

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleRepeat() {
	fmt.Println(Repeat("1", 5))
	// Output: 11111
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	//goland:noinspection SpellCheckingInspection
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkRepeatStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 5)
	}
}
