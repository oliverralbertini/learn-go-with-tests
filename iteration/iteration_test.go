package iteration

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	got := Reverse("abcdefg")
	expected := "gfedcba"
	if got != expected {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}
}

func TestRepeat(t *testing.T) {
	got := Repeat(4, "a")
	expected := "aaaa"
	if got != expected {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(7, "a")
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("abcdefghijklmnopqrstuvwxyz")
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat(2, "abc"))
	// Output: abcabc
}
