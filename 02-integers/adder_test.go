package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	fmt.Println(Add(1, 5))
	// Output: 6
}

func TestAdder(t *testing.T) {
	got := Add(10, 16)
	want := 26

	if got != want {
		t.Errorf("expected '%d' but got '%d'", got, want)
	}
}
