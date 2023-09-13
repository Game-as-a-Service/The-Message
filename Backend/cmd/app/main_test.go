package main

import (
	// "bytes"
	// "encoding/json"
	// "net/http"
	// "net/http/httptest"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/DATA-DOG/go-sqlmock"
	httpHandler "github.com/Game-as-a-Service/The-Message/service/delivery/http/v1"
	mysqlRepo "github.com/Game-as-a-Service/The-Message/service/repository/mysql"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetGameByIdE2E(t *testing.T) {
	// 創建一個虛擬的MySQL數據庫連接
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// 替換GORM的數據庫連接為虛擬的連接
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM database: %v", err)
	}

	// 設置預期的模擬數據庫查詢和操作
	mock.ExpectQuery("SELECT `id`,`name` FROM `games` WHERE id = ? AND `games`.`deleted_at` IS NULL ORDER BY `games`.`id` LIMIT 1").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Test Game"))

	// 創建Gin路由和HTTP伺服器
	router := gin.Default()
	gameHandler := &httpHandler.Game{
		GameRepo: mysqlRepo.NewGameRepositoryRepository(gdb),
	}
	router.GET("/api/v1/game/:gameId", gameHandler.GetGameById)

	// 準備HTTP GET請求
	req, _ := http.NewRequest("GET", "/api/v1/game/1", nil)
	recorder := httptest.NewRecorder()

	// 執行HTTP請求
	router.ServeHTTP(recorder, req)

	// 檢查HTTP響應
	assert.Equal(t, http.StatusOK, recorder.Code)

	// 解析HTTP響應
	var response mysqlRepo.Game
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	// 檢查模擬的數據庫調用是否符合預期
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

	// 檢查數據庫操作是否正確，這可以通過模擬數據庫操作來實現
}

// func TestCreateGameE2E(t *testing.T) {
// 	// 創建一個虛擬的MySQL數據庫連接
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock database: %v", err)
// 	}
// 	defer db.Close()

// 	// 設置預期的模擬數據庫查詢和操作
// 	mock.ExpectExec("INSERT INTO games").WillReturnResult(sqlmock.NewResult(1, 1))

// 	// 初始化GORM數據庫連接
// 	gdb, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      db,
// 		DriverName:                "mysql",
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})

// 	if err != nil {
// 		t.Fatalf("Error opening GORM database: %v", err)
// 	}

// 	// 創建Gin路由和HTTP伺服器
// 	router := gin.Default()
// 	gameHandler := &httpHandler.Game{
// 		GameRepo: mysqlRepo.NewGameRepositoryRepository(gdb),
// 	}
// 	router.POST("/api/v1/game", gameHandler.CreateGame)

// 	// 準備HTTP POST請求
// 	reqBody := []byte(`{ "name": "Test Game" }`)
// 	req, _ := http.NewRequest("POST", "/api/v1/game", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Content-Type", "application/json")
// 	recorder := httptest.NewRecorder()

// 	// 執行HTTP請求
// 	router.ServeHTTP(recorder, req)

// 	// 檢查HTTP響應
// 	assert.Equal(t, http.StatusOK, recorder.Code)

// 	// 解析HTTP響應
// 	var response mysqlRepo.Game
// 	err = json.Unmarshal(recorder.Body.Bytes(), &response)
// 	assert.NoError(t, err)

// 	// 檢查模擬的數據庫調用是否符合預期
// 	// if err := mock.ExpectationsWereMet(); err != nil {
// 	// 	t.Errorf("Unfulfilled expectations: %s", err)
// 	// }
// }
