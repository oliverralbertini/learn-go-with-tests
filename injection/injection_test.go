package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Oliver")

	have := buffer.String()
	want := "Hello, Oliver!"

	if have != want {
		t.Errorf("have '%s' want '%s'", have, want)
	}
}
