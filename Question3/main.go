package main

import (
	"bufio"
	"fmt"

	// "go/scanner"
	"os"
	// "strings"
)

func main() {
	bufio.NewScanner(os.Stdin)

	fmt.Println("What is your Name")
}
