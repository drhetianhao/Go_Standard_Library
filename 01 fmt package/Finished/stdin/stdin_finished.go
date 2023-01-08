package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // read from key input

	s, _ := reader.ReadString('\n')
	fmt.Println(s)
}
