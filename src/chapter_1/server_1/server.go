// Минимальный echo - сервер
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler) // Каждый запрос вызывает обработчик
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissa", lissa)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработчик возвращает компонент пути из URL запроса
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// Счетчик, возвращающий количество сделанных запросов
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// Построение функции Лисажу по данным пользователя из строки запроса
func lissa(w http.ResponseWriter, r *http.Request) {
	vals := make(map[string]int)
	cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
	if err != nil {
		cycles = 5
	}
	vals["cycles"] = cycles
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		size = 100
	}
	vals["size"] = size
	nframes, err := strconv.Atoi(r.URL.Query().Get("nframes"))
	if err != nil {
		nframes = 64
	}
	vals["nframes"] = nframes
	delay, err := strconv.Atoi(r.URL.Query().Get("delay"))
	if err != nil {
		delay = 8
	}
	vals["delay"] = delay
	lissajous(w, vals)
}
