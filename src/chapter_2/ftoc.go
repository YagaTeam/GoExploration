// Выводит результаты преобразования двух температур
// по Фаренгейту в температуру по Цельсию

package main

import "fmt"

func ftocPrint() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func incr(p *int) int {
	*p++ //Увеличивает значение, на которое указывает p;
	//Не изменяет значение p
	return *p
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
