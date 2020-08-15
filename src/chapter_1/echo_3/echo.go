// Упражнение 1.1. Выводим имя выполняемой команды

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Command: ", os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
