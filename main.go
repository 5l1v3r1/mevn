package main

import (
	"fmt"
	"os"
)

func main() {
	cs := os.Getenv("CIRCLE_SHA1")
	fmt.Print(cs)
}
