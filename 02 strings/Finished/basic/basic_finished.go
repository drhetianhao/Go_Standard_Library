package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"

	// Basic string operations

	// Length of string
	fmt.Println(len(s))

	// iterate over each character
	for _, ch := range s {
		fmt.Print(string(ch), ",")
	}
	fmt.Println()

	// Using operators < > == !=
	fmt.Println("dog" < "cat")
	fmt.Println("dog" < "horse")
	fmt.Println("dog" == "Dog")

	// Comparing two strings
	result := strings.Compare("dog", "cat")
	fmt.Println(result)
	result = strings.Compare("dog", "dog")
	fmt.Println(result)

	// EqualFold tests using Unicode case-folding
	fmt.Println(strings.EqualFold("Hello", "hello")) // ignore the case

	// ToUpper, ToLower, Title
	s1 := strings.ToUpper(s)
	s2 := strings.ToLower(s)
	s3 := strings.Title(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
/*
43
T,h,e, ,q,u,i,c,k, ,b,r,o,w,n, ,f,o,x, ,j,u,m,p,s, ,o,v,e,r, ,t,h,e, ,l,a,z,y, ,d,o,g,
false
true
false
1
0
true
THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG
the quick brown fox jumps over the lazy dog
The Quick Brown Fox Jumps Over The Lazy Dog
*/