//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package repositories

import (
	"github.com/tasuke/go-mux-task/models"
	"golang.org/x/crypto/openpgp/errors"
	"gorm.io/gorm"
)

type ITodoRepository interface {
	GetAllTodos(todos *[]models.Todo, userId int) error
	GetTodo(todo *models.Todo, id string, userId int) error
	GetTodoLast(todo *models.Todo, userId int) error
	CreateTodo(todo *models.Todo) error
	DeleteTodo(id string, userId int) error
	UpdateTodo(todo *models.Todo, id string, userId int) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &todoRepository{db}
}

func (tr *todoRepository) GetAllTodos(todos *[]models.Todo, userId int) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) GetTodo(todo *models.Todo, id string, userId int) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(&todo, id).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) GetTodoLast(todo *models.Todo, userId int) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Last(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) CreateTodo(todo *models.Todo) error {
	if err := tr.db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) DeleteTodo(id string, userId int) error {
	if err := tr.db.Where("id=? AND user_id=?", id, userId).Delete(&models.Todo{}).Error; err != nil {
		return err
	}
	if tr.db.RowsAffected < 1 {
		return errors.InvalidArgumentError("No record found")
	}
	return nil
}

func (tr *todoRepository) UpdateTodo(todo *models.Todo, id string, userId int) error {
	if err := tr.db.Model(&todo).Where("id=? AND user_id=?", id, userId).Updates(
		map[string]interface{}{
			"title":   todo.Title,
			"comment": todo.Comment,
		}).Error; err != nil {
		return err
	}
	return nil
}
