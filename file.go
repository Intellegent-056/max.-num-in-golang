package main

import "fmt"

func main() {
	var s float64 = 99999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999 // It's max.
	fmt.Printf("%v %T", s, s)
}

// run: go run file.go
// build: go build file.go
