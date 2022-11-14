package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("length -----> ", len(os.Args))
	fmt.Println("args -------->", os.Args)
	fmt.Println("1 to length-1", os.Args[1:len(os.Args)])
	fmt.Println("1 to length-1", os.Args[1:])
	fmt.Println("0 to 2", os.Args[:3])
}
