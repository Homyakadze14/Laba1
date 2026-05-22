package main

import (
	"fmt"
	"laba4/internal/entity"
	"laba4/internal/infra/file"
	"laba4/internal/infra/mapbase"
	"laba4/internal/usecase"
)

func main() {
	repo := mapbase.NewDB()
	fileStorage := file.NewFileStore()
	uc := usecase.NewJournalUsecase(repo, fileStorage)

	fmt.Println("=== 1. Добавление студентов ===")
	_ = uc.AddStudent("Alice")
	_ = uc.AddStudent("Bob")
	_ = uc.AddStudent("Charlie")
	_ = uc.AddStudent("Diana")
	_ = uc.AddStudent("Eve")

	fmt.Println("Студенты после добавления:")
	printStudents(uc.GetAll())

	fmt.Println("\n=== 2. Проверка существования ===")
	fmt.Printf("Alice exists: %v\n", uc.FoundStudent("Alice"))
	fmt.Printf("Unknown exists: %v\n", uc.FoundStudent("Unknown"))

	fmt.Println("\n=== 3. Выставление оценок ===")
	_ = uc.UpdateMark("Alice", 85)
	_ = uc.UpdateMark("Bob", 72)
	_ = uc.UpdateMark("Charlie", 95)
	_ = uc.UpdateMark("Diana", 60)
	_ = uc.UpdateMark("Eve", 40)

	fmt.Println("Оценки:")
	printMarks(uc, []string{"Alice", "Bob", "Charlie", "Diana", "Eve"})

	fmt.Println("\n=== 4. Обновление оценки ===")
	_ = uc.UpdateMark("Alice", 90)
	mark, _ := uc.GetMark("Alice")
	fmt.Printf("Новая оценка Alice: %d\n", mark)

	fmt.Println("\n=== 5. Получение оценки несуществующего студента ===")
	_, err := uc.GetMark("Nobody")
	fmt.Printf("Ошибка: %v\n", err)

	fmt.Println("\n=== 6. Вывод всех студентов ===")
	printStudents(uc.GetAll())

	fmt.Println("\n=== 7. Средний балл ===")
	fmt.Printf("Средний балл: %.2f\n", uc.GetAvg())

	fmt.Println("\n=== 8. Максимальный балл ===")
	maxStud := uc.GetMax()
	if maxStud != nil {
		fmt.Printf("Максимум: %s (%d)\n", maxStud.Name, maxStud.Mark)
	}

	fmt.Println("\n=== 9. Минимальный балл ===")
	minStud := uc.GetMin()
	if minStud != nil {
		fmt.Printf("Минимум: %s (%d)\n", minStud.Name, minStud.Mark)
	}

	fmt.Println("\n=== 10. Студенты с баллом >= 75 ===")
	high := uc.GetAllMoreThen(75)
	printStudents(high)

	fmt.Println("\n=== 11. Студенты с баллом < 60 (не прошли порог) ===")
	low := uc.CountBelow(60)
	fmt.Println("Количество студентов с баллом < 60:", low)

	fmt.Println("\n=== 12. Сортировка по имени ===")
	sorted := uc.GetAllSortedByName()
	printStudents(sorted)

	fmt.Println("\n=== 13. Удаление студента ===")
	_ = uc.DeleteStudent("Eve")
	fmt.Println("После удаления Eve:")
	printStudents(uc.GetAll())

	fmt.Println("\n=== 14. Сохранение в файл ===")
	errSave := uc.SaveToFile("test_journal.json")
	if errSave != nil {
		fmt.Printf("Ошибка сохранения: %v\n", errSave)
	} else {
		fmt.Println("Журнал сохранён в test_journal.json")
	}

	fmt.Println("\n=== 15. Загрузка из файла ===")
	errLoad := uc.LoadFromFile("test_journal.json")
	if errLoad != nil {
		fmt.Printf("Ошибка загрузки: %v\n", errLoad)
	} else {
		fmt.Println("Журнал загружен из test_journal.json")
		printStudents(uc.GetAll())
	}
}

func printStudents(students []entity.Stud) {
	if len(students) == 0 {
		fmt.Println("Нет студентов.")
		return
	}
	for _, s := range students {
		fmt.Printf("  %s: %d\n", s.Name, s.Mark)
	}
}

func printMarks(uc *usecase.JournalUsecase, names []string) {
	for _, name := range names {
		mark, err := uc.GetMark(name)
		if err != nil {
			fmt.Printf("  %s: ошибка (%v)\n", name, err)
		} else {
			fmt.Printf("  %s: %d\n", name, mark)
		}
	}
}
