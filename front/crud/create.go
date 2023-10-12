package crud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-todo-cli/model"
	"log"
	"net/http"

	"github.com/manifoldco/promptui"
)

func Create() {
	validateTitle := func(input string) error {
		if len(input) == 0 {
			return errors.New("Input should not be empty")
		}
		if len(input) > 20 {
			return errors.New("Input should be no more than 20 characters long")
		}
		return nil
	}
	promptTitle := promptui.Prompt{
		Label:    "Title",
		Validate: validateTitle,
	}
	titleResult, err := promptTitle.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	validateDescription := func(input string) error {
		if len(input) == 0 {
			return errors.New("Input should not be empty")
		}
		if len(input) > 50 {
			return errors.New("Input should be no more than 50 characters long")
		}
		return nil
	}
	promptDescription := promptui.Prompt{
		Label:    "Description",
		Validate: validateDescription,
	}
	descriptionResult, err := promptDescription.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	todo := model.Todo{
		Title:       titleResult,
		Description: descriptionResult,
	}

	json, err := json.Marshal(todo)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/todos", "application/json", bytes.NewBuffer(json))
	if err != nil {
		log.Fatalf("Error making the request: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println("Todo created successfully")
}
