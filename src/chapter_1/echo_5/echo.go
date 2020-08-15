package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Version 1:")
	start := time.Now()
	s := ""
	for i := 0; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}
	fmt.Println(s)
	sec := time.Since(start).Seconds()
	fmt.Println("Time elapsed: ", sec)
	fmt.Println("Version 2:")
	start = time.Now()
	s = ""
	for _, arg := range os.Args[0:] {
		s += arg + " "
	}
	fmt.Println(s)
	sec = time.Since(start).Seconds()
	fmt.Println("Time elapsed: ", sec)
	fmt.Println("Version 3:")
	start = time.Now()
	s = ""
	fmt.Println(strings.Join(os.Args[0:], " "))
	sec = time.Since(start).Seconds()
	fmt.Println("Time Elapsed: ", sec)
}
