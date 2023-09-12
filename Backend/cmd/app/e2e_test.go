// e2e_test.go

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var serverURL string

func TestE2EWithMySQL(t *testing.T) {
	// 初始化資料庫連接
	db, err := InitDB()
	if err != nil {
		t.Fatalf("無法初始化資料庫: %v", err)
	}
	defer db.Close()

	// 清空測試資料庫
	db.Exec("DELETE FROM todos")

	// 啟動測試用的 HTTP 伺服器，模擬後端應用程式

	log.Println("Starting server...")
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8081", nil)

	// 等待伺服器啟動（這部分可以根據實際情況調整）
	time.Sleep(15 * time.Second)

	// 設定全域變數以供測試使用
	serverURL = "http://localhost:8081" // 注意這裡的 URL

	// 模擬 POST 請求，新增待辦事項
	payload := []byte(`{"task": "完成 E2E 測試 with MySQL"}`)
	resp, err := http.Post(serverURL+"/todos", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("HTTP POST 請求失敗: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, resp.StatusCode)
	}

	// 模擬 GET 請求，檢索所有待辦事項
	resp, err = http.Get(serverURL + "/todos")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	// 解析回應 JSON
	var todos []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		t.Fatal(err)
	}

	// 驗證回應中是否包含我們新增的待辦事項
	found := false
	for _, todo := range todos {
		if task, ok := todo["task"].(string); ok && task == "完成 E2E 測試 with MySQL" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected task '完成 E2E 測試 with MySQL' not found in response")
	}
}
