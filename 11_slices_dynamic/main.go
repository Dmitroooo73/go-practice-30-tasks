package main

import "fmt"

func main() {
	// Создаем пустой срез строк с начальной емкостью 10
	slice := make([]string, 0, 10)

	// Динамически добавляем элементы с помощью append
	slice = append(slice, "Первый")
	slice = append(slice, "Второй")
	slice = append(slice, "Третий")
	slice = append(slice, "Четвертый")

	fmt.Println("Элементы среза:")
	for _, elem := range slice {
		fmt.Println(elem)
	}
}
