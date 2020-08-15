package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("index: ", i, " arg: ", os.Args[i])
	}
	fmt.Println("Alternative version:")
	for key, arg := range os.Args[0:] {
		fmt.Println("index: ", key, " arg: ", arg)
	}
}
