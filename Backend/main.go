package main

import (
	"fmt"
)

func main() {
	addTodo("test")
	test, err := findTodoByID(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test)
	setPriority(1, PriorityUrgent)
	setStatus(1, ProgressionInProgress)
	fmt.Println(test)
}
