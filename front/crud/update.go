package crud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-todo-cli/model"
	"log"
	"net/http"

	"github.com/manifoldco/promptui"
)

func Update() {
	resp, err := http.Get("http://localhost:8080/todos")
	if err != nil {
		log.Fatalf("Error making the request: %v", err)
	}
	defer resp.Body.Close()

	todos := []model.Todo{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&todos); err != nil {
		log.Fatalf("Error decoding the response: %v", err)
	}

	prompt := promptui.Select{
		Label: "Which todo do you want to update status?",
		Items: todos,
	}

	selectedIndex, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	selectedTodo := todos[selectedIndex]
	selectedTodo.Status = !selectedTodo.Status
	url := fmt.Sprintf("http://localhost:8080/todos/%d", selectedTodo.ID)

	jsonData, err := json.Marshal(selectedTodo)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	defer resp.Body.Close()
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalf("Error making the PUT request: %v", err)
	}

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("Error sending the PUT request: %v", err)
	}

	fmt.Println("Todo updated successfully")
}
