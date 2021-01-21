package main

import (
	"testing"
)

func BenchmarkNoGoroutines(b *testing.B) {
	inputs, err := getInputs()
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		noGoroutines(inputs)
	}
}

func BenchmarkWithGoroutines(b *testing.B) {
	inputs, err := getInputs()
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		withGoroutines(inputs)
	}
}
