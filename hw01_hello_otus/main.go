package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	reverseHelloOTUS := stringutil.Reverse("Hello, OTUS!")

	fmt.Println(reverseHelloOTUS)
}
