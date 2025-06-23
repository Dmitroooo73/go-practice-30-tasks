package main

import "fmt"

// removeElement удаляет элемент по заданному индексу из среза строк.
func removeElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice // Если индекс неверный, возвращаем исходный срез
	}
	// Удаляем элемент, объединяя часть до и после индекса
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	words := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("Исходный срез:", words)

	// Удаляем элемент с индексом 2 ("cherry")
	words = removeElement(words, 2)
	fmt.Println("После удаления:", words)
}
