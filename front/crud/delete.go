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

func Delete() {
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
		Label: "Which todo do you want to delete?",
		Items: todos,
	}

	selectedIndex, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	selectedTodo := todos[selectedIndex]
	url := fmt.Sprintf("http://localhost:8080/todos/%d", selectedTodo.ID)

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(nil))
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Error making the DELETE request: %v", err)
	}

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("Error sending the PUT request: %v", err)
	}

	fmt.Println("Todo deleted successfully")
}
