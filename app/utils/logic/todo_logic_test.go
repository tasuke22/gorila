package logic

import (
	"github.com/stretchr/testify/assert"
	"github.com/tasuke/go-mux-task/models"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestCreateAllTodoResponse(t *testing.T) {
	todoLogic := NewTodoLogic()

	todos := []models.Todo{
		{
			Model:   gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Title:   "Test Todo 1",
			Comment: "Test Comment 1",
		},
		{
			Model:   gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Title:   "Test Todo 2",
			Comment: "Test Comment 2",
		},
	}

	responses := todoLogic.CreateAllTodoResponse(&todos)

	assert.Len(t, responses, 2)
	assert.Equal(t, uint(1), responses[0].ID)
	assert.Equal(t, "Test Todo 1", responses[0].Title)
	assert.Equal(t, "Test Todo 2", responses[1].Title)
}

func TestCreateTodoResponse(t *testing.T) {
	todoLogic := NewTodoLogic()

	// テストデータの準備
	todo := models.Todo{
		Model:   gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Title:   "Test Todo",
		Comment: "Test Comment",
	}

	// メソッドの実行
	response := todoLogic.CreateTodoResponse(&todo)

	// 結果の検証
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, "Test Todo", response.Title)
	assert.Equal(t, "Test Comment", response.Comment)
}
