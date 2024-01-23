package repositories

import (
	"github.com/stretchr/testify/assert"
	"github.com/tasuke/go-mux-task/models"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_todoRepository_GetAllTodos(t *testing.T) {
	// テスト対象のメソッドに渡す引数を設定
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// インターフェースのモックインスタンスを作成
	mockRepo := NewMockITodoRepository(ctrl)

	// モックの期待値を設定
	userId := uint(1)
	// 引数として何らかのスライスと userId を受け取り、nil を返すことを期待 gomock.Any() はtodos の具体的な値に関心がないために使用
	mockRepo.EXPECT().GetAllTodos(gomock.Any(), userId).Return(nil)

	// テスト対象メソッドの呼び出し
	var todos []models.Todo
	err := mockRepo.GetAllTodos(&todos, userId)

	assert.NoError(t, err)
}

func TestCreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockITodoRepository(ctrl)
	todo := models.Todo{Title: "New Todo", Comment: "New Comment"}

	mockRepo.EXPECT().CreateTodo(&todo).Return(nil)

	err := mockRepo.CreateTodo(&todo)
	assert.NoError(t, err)
}

func TestGetTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockITodoRepository(ctrl)
	var todo models.Todo
	todoID, userID := uint(1), uint(1)

	mockRepo.EXPECT().GetTodo(&todo, todoID, userID).Return(nil)

	err := mockRepo.GetTodo(&todo, todoID, userID)
	assert.NoError(t, err)
}

func TestDeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockITodoRepository(ctrl)
	todoID, userID := uint(1), uint(1)

	mockRepo.EXPECT().DeleteTodo(todoID, userID).Return(nil)

	err := mockRepo.DeleteTodo(todoID, userID)
	assert.NoError(t, err)
}

func TestUpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockITodoRepository(ctrl)
	todo := models.Todo{Title: "Updated Todo", Comment: "Updated Comment"}
	todoID, userID := "1", 1

	mockRepo.EXPECT().UpdateTodo(&todo, todoID, userID).Return(nil)

	err := mockRepo.UpdateTodo(&todo, todoID, userID)
	assert.NoError(t, err)
}
