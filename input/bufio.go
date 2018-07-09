package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the input: ")
	// input, err := inputReader.ReadSlice('a')
	input, err := inputReader.ReadString('a')
	if err == nil {
		fmt.Printf("The input is %v", input)
	}
}
