package main

import "fmt"

type circle struct {
	radius int
	border int
}

func main() {
	x := 20
	f := 123.45

	// Basic formatting
	fmt.Printf("%d\n", x) // decimal
	fmt.Printf("%x\n", x) // hex

	// Booleans can be printed as "true" or "false"
	fmt.Printf("%t\n", x > 10) // boolean

	// floating point numbers
	fmt.Printf("%f\n", f) // float
	fmt.Printf("%e\n", f) // exp

	// Using explicit argument indexes
	fmt.Printf("%[2]d %[1]d\n", 52, 40) // => 40, 52
	// Argument indexes can be used to print values repeatedly
	fmt.Printf("%d %#[1]o %#[1]x\n", 52) // decimal, oct, hex

	// Print a value in default format
	c := circle{
		radius: 20,
		border: 5,
	}
	fmt.Printf("%v\n", c) // default value
	fmt.Printf("%+v\n", c) // with the key and value
	fmt.Printf("%T\n", c) // type

	// Sprintf is the same as Printf, but returns a string
	s := fmt.Sprintf("%[2]d %[1]d\n", 52, 40)
	fmt.Println(s)
}

/*
20
14
true
123.450000
1.234500e+02
40 52
52 064 0x34
{20 5}
{radius:20 border:5}
main.circle
40 52
*/