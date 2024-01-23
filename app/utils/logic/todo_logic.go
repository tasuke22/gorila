package logic

import "github.com/tasuke/go-mux-task/models"

type ITodoLogic interface {
	CreateAllTodoResponse(todos *[]models.Todo) []models.BaseTodoResponse
	CreateTodoResponse(todo *models.Todo) models.BaseTodoResponse
}

type todoLogic struct{}

func NewTodoLogic() ITodoLogic {
	return &todoLogic{}
}

func (tl *todoLogic) CreateAllTodoResponse(todos *[]models.Todo) []models.BaseTodoResponse {
	var responseTodos []models.BaseTodoResponse
	for _, todo := range *todos {
		var newTodo models.BaseTodoResponse
		newTodo.ID = todo.ID
		newTodo.CreatedAt = todo.CreatedAt
		newTodo.UpdatedAt = todo.UpdatedAt
		newTodo.DeletedAt = todo.DeletedAt
		newTodo.Title = todo.Title
		newTodo.Comment = todo.Comment
		responseTodos = append(responseTodos, newTodo)
	}

	return responseTodos
}

func (tl *todoLogic) CreateTodoResponse(todo *models.Todo) models.BaseTodoResponse {
	var responseTodo models.BaseTodoResponse
	responseTodo.ID = todo.ID
	responseTodo.CreatedAt = todo.CreatedAt
	responseTodo.UpdatedAt = todo.UpdatedAt
	responseTodo.DeletedAt = todo.DeletedAt
	responseTodo.Title = todo.Title
	responseTodo.Comment = todo.Comment

	return responseTodo
}
