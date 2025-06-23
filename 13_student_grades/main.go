package main

import "fmt"

// addGrade добавляет или обновляет оценку студента.
func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
}

// getGrade ищет оценку по имени студента.
func getGrade(grades map[string]int, name string) (int, bool) {
	grade, exists := grades[name]
	return grade, exists
}

// deleteGrade удаляет оценку студента по имени.
func deleteGrade(grades map[string]int, name string) {
	delete(grades, name)
}

func main() {
	grades := make(map[string]int)

	addGrade(grades, "Алексей", 5)
	addGrade(grades, "Мария", 4)
	addGrade(grades, "Иван", 3)

	fmt.Println("Карта оценок:", grades)

	if grade, found := getGrade(grades, "Мария"); found {
		fmt.Printf("Оценка Марии: %d\n", grade)
	} else {
		fmt.Println("Оценка не найдена.")
	}

	deleteGrade(grades, "Иван")
	fmt.Println("После удаления Ивана:", grades)
}
