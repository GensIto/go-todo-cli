package main

import (
	"fmt"
	"go-todo-cli/crud"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "HI😎 I'm todo app!!",
		Items: []string{"create todo", "get todo", "update todo", "delete todo"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "create todo":
		crud.Create()
	case "get todo":
		crud.Read()
	case "update todo":
		crud.Update()
	case "delete todo":
		crud.Delete()
	}
}
