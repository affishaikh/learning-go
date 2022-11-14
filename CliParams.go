package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Split(os.Args[1], ","))
}
