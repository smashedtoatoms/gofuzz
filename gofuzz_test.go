package gofuzz

import (
	"fmt"
)

// The actual tests are located in the phon and sim subdirectories.  These are
// only here to provide examples for the documentation.

func ExampleMetaphone() {
	result, err := Metaphone("Colorado")
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(result)
	// Output: klrt
}

func ExampleMetaphoneMetric() {
	result, err := MetaphoneMetric("Colorado", "Kolorado")
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(result)
	// Output: true
}

func ExampleJaro() {
	result, err := Jaro("bob", "bobby")
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(result)
	// Output: 0.75555557
}

func ExampleJaroWinkler() {
	result, err := JaroWinkler("bob", "bobby")
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(result)
	// Output: 0.8288889
}

func ExampleOverlap() {
	result, err := Overlap("Chipotle", "Chipotle Mexican Grill", 2)
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(result)
	// Output: 1
}
