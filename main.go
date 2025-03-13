package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var summCompleted int = 0

type Task struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

var tasks []Task
var dataFile = "tasks.json"

func loadTasks() error { // Загружаем все задачи из файла tasks.json
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &tasks)
}

func saveTasks() error { //Сохраняем все задачи в tasks.json файл
	file, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, file, 0644)
}

func addTasks(text string) { // Добавляем новую задачу и сохраняем её в tasks.json
	task := Task{
		ID:        len(tasks) + 1,
		Text:      text,
		Completed: false,
	}
	tasks = append(tasks, task)
	fmt.Printf("Задача успешно добавлена: %s\n", text)
}

func listTasks(c int) { // Выводим список всех задач или выполненых задач

	switch c {
	case 1:
		if len(tasks) == 0 { // Здесь выводиться список всех задач
			fmt.Println("Задач нет")
			return
		} else {
			for _, task := range tasks {
				status := " "
				if task.Completed {
					status = "X"
				}
				fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Text)
			}
		}
	case 2:
		if summCompleted == 0 {
			fmt.Println("Выполненых задач нет")
		} else {
			for _, task := range tasks { // А здесь выводиться список всех выполненых задач
				status := " "
				if task.Completed {
					status = "X"
					fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Text)
				}
			}
		}
	}
}

func compliteTasks(id int) { // Помечаем выбранную по ID задачу как выполненную
	for i, task := range tasks {
		if task.ID == id {
			if tasks[i].Completed {
				fmt.Printf("Задача номер %v уже выполнена '%s'\n", task.ID, task.Text)
				return
			} else {
				tasks[i].Completed = true
				fmt.Printf("Задача номер %v выполнена '%s'\n", task.ID, task.Text)
				summCompleted++
				return
			}

		}
	}
	fmt.Println("Задача не найдена.")
}

func deliteTasks(id int) { // Удаляем выбранную по ID задачу
	for i, task := range tasks {
		if task.ID == id {
			if task.Completed {
				summCompleted--
			}
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Задача номер %v удалена '%s'\n", task.ID, task.Text)
			return

		}
	}
	fmt.Println("Задача не найдена.")
}

func main() {
	loadTasks() // Загружаем все задачи из файла tasks.json
	for _, task := range tasks {
		if task.Completed {
			summCompleted = 1
		}
	}
	reader := bufio.NewReader(os.Stdin)
	for { // Даем пользователю выбор действия
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать все задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Показать только выполненые задачи")
		fmt.Println("5. Удалить задачу")
		fmt.Println("6. Выйти")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("\nЗапишите задачу")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			addTasks(text)
		case "2":
			fmt.Println("\nЗадачи:")
			listTasks(1)
		case "3":
			fmt.Println("\nВведите ID задачи")
			var id int
			fmt.Scanln(&id)
			compliteTasks(id)
		case "4":
			fmt.Println("\nВыполненые задачи")
			listTasks(2)
		case "5":
			fmt.Println("\nВведите ID задачи")
			var id int
			fmt.Scanln(&id)
			deliteTasks(id)
		case "6":
			fmt.Println("\nПока!)")
			if err := saveTasks(); err != nil {
				fmt.Println("Ошибка сохранения задач:", err)
			}
			return
		default:
			fmt.Printf("Вы ввели неккоретктное значение")
		}
	}
}
