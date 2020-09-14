package main

import (
	"fmt"

	"../chapter_2/tempconv"
)

func main() {
	fmt.Printf("Брррр! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}

func delta(old, new int) int {
	return new - old
}
