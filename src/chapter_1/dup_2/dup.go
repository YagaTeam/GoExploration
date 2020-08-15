// выводит текст каждой строки, которая появляется во входных данных
// более одного раза. Программа читает стандартный ввод или список
// именованных файлов.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fileList := make(map[string]int)
	if len(files) == 0 {
		countLines(os.Stdin, counts, nil)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts, fileList)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileList map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "Q" || input.Text() == "q" {
			counts[input.Text()]++
		}
	}
	if fileList != nil {
		if len(counts) > 1 {
			fileList[f.Name()]++
		}
	}
	// Ignore all potential errors from input.Err()
}
