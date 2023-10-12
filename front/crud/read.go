package crud

import (
	"encoding/json"
	"fmt"
	"go-todo-cli/model"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
)

func Read() {
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

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.FilterHTML)
	fmt.Fprintln(w, "ID\tTitle\tDescription\tStatus")
	for _, todo := range todos {
		fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", todo.ID, todo.Title, todo.Description, todo.Status)
	}
	w.Flush()
}
