package services

import (
	"github.com/tasuke/go-mux-task/repositories"
	"github.com/tasuke/go-mux-task/utils/logic"
	"github.com/tasuke/go-mux-task/utils/validation"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tasuke/go-mux-task/models"
)

func TestTodoService_GetAllTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 依存関係のモックを作成
	mockRepo := repositories.NewMockITodoRepository(ctrl)
	mockTodoLogic := logic.NewMockITodoLogic(ctrl)
	mockRespLogic := logic.NewMockIResponseLogic(ctrl)
	mockValidation := validation.NewMockITodoValidation(ctrl)

	// TodoServiceのインスタンスを作成
	todoService := NewTodoService(mockRepo, mockTodoLogic, mockRespLogic, mockValidation)

	// 正常系のテストケース
	t.Run("Successfully retrieved todos", func(t *testing.T) {
		w := httptest.NewRecorder()
		userId := 1
		dummyBaseTodos := []models.BaseTodoResponse{{
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

		mockRepo.EXPECT().GetAllTodos(gomock.Any(), userId).Return(nil)
		mockTodoLogic.EXPECT().CreateAllTodoResponse(gomock.Any()).Return(dummyBaseTodos)

		responseTodos, err := todoService.GetAllTodos(w, userId)

		assert.NoError(t, err)
		assert.Len(t, responseTodos, 2)
	})
}
