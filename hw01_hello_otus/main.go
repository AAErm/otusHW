package main

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

func main() {
	str := "Hello, OTUS!"
	fmt.Print(stringutil.Reverse(str))
}
