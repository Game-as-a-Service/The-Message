// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
)

// var db *sql.DB

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		getTodos(w, r)
// 	case http.MethodPost:
// 		createTodo(w, r)
// 	}
// }

// func getTodos(w http.ResponseWriter, r *http.Request) {
// 	// 執行 SQL 查詢來檢索待辦事項列表
// 	rows, err := db.Query("SELECT id, task FROM todos")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var todos []map[string]interface{}
// 	for rows.Next() {
// 		var id int
// 		var task string
// 		if err := rows.Scan(&id, &task); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		todo := map[string]interface{}{
// 			"id":   id,
// 			"task": task,
// 		}
// 		todos = append(todos, todo)
// 	}

// 	// 回傳待辦事項列表
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(todos)
// }

// func createTodo(w http.ResponseWriter, r *http.Request) {
// 	// 解析請求 JSON
// 	var input map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// 執行 SQL 插入語句以新增待辦事項
// 	result, err := db.Exec("INSERT INTO todos (task) VALUES (?)", input["task"])
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// 回傳新增待辦事項的 ID
// 	id, _ := result.LastInsertId()
// 	response := map[string]interface{}{
// 		"id": id,
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name") // get URL query string
		content := fmt.Sprintf("hello, %s", name)
		fmt.Fprint(w, content) // write out content
	})
	http.ListenAndServe(":8080", nil)
}
