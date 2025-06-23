package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. Go is expressive, concise, clean, and efficient."

	// Приводим текст к нижнему регистру и разбиваем на слова
	words := strings.Fields(strings.ToLower(text))

	frequency := make(map[string]int)
	for _, word := range words {
		// Убираем знаки препинания (если они есть)
		word = strings.Trim(word, ".,")
		frequency[word]++
	}

	fmt.Println("Частота слов:")
	for word, count := range frequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}
