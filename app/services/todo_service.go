//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package services

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/tasuke/go-mux-task/models"
	"github.com/tasuke/go-mux-task/repositories"
	"github.com/tasuke/go-mux-task/utils/logic"
	"github.com/tasuke/go-mux-task/utils/validation"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type TodoService interface {
	GetAllTodos(w http.ResponseWriter, userId int) ([]models.BaseTodoResponse, error)
	GetTodo(w http.ResponseWriter, r *http.Request, userId int) (models.BaseTodoResponse, error)
	CreateTodo(w http.ResponseWriter, r *http.Request, userId int) (models.BaseTodoResponse, error)
	DeleteTodo(w http.ResponseWriter, r *http.Request, userId int) error
	UpdateTodo(w http.ResponseWriter, r *http.Request, userId int) (models.BaseTodoResponse, error)
	SendAllTodoResponse(w http.ResponseWriter, todos *[]models.BaseTodoResponse)
	SendTodoResponse(w http.ResponseWriter, todo *models.BaseTodoResponse)
	SendCreateTodoResponse(w http.ResponseWriter, todo *models.BaseTodoResponse)
	SendDeleteTodoResponse(w http.ResponseWriter)
}

type todoService struct {
	tr repositories.ITodoRepository
	tl logic.ITodoLogic
	rl logic.IResponseLogic
	tv validation.ITodoValidation
}

func NewTodoService(
	tr repositories.ITodoRepository,
	tl logic.ITodoLogic,
	rl logic.IResponseLogic,
	tv validation.ITodoValidation,
) TodoService {
	return &todoService{tr, tl, rl, tv}
}

func (ts *todoService) GetAllTodos(w http.ResponseWriter, userId int) ([]models.BaseTodoResponse, error) {
	var todos []models.Todo
	if err := ts.tr.GetAllTodos(&todos, userId); err != nil {
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("データ取得に失敗"), http.StatusInternalServerError)
	}
	responseTodos := ts.tl.CreateAllTodoResponse(&todos)
	return responseTodos, nil
}

func (ts *todoService) GetTodo(w http.ResponseWriter, r *http.Request, userId int) (models.BaseTodoResponse, error) {
	// リクエストURLからパラメータを抽出する
	vars := mux.Vars(r)
	id := vars["id"]
	var todo models.Todo

	if err := ts.tr.GetTodo(&todo, id, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("該当データは存在しません"), http.StatusBadRequest)
		} else {
			ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("データ取得に失敗しました"), http.StatusInternalServerError)
		}
		return models.BaseTodoResponse{}, err
	}
	responseTodos := ts.tl.CreateTodoResponse(&todo)
	return responseTodos, nil
}

func (ts *todoService) CreateTodo(w http.ResponseWriter, r *http.Request, userId int) (models.BaseTodoResponse, error) {
	// リクエストボディを構造体へ変換
	reqBody, _ := io.ReadAll(r.Body)
	var mutationTodoRequest models.MutationTodoRequest
	if err := json.Unmarshal(reqBody, &mutationTodoRequest); err != nil {
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("リクエストパラメータを構造体へ変換処理できません"), http.StatusBadRequest)
		return models.BaseTodoResponse{}, err
	}
	// バリデーションチェック
	if err := ts.tv.MutationTodoValidate(mutationTodoRequest); err != nil {
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse(err.Error()), http.StatusBadRequest)
		return models.BaseTodoResponse{}, err
	}

	// リクエストボディを構造体へ変換
	var todo models.Todo
	todo.Title = mutationTodoRequest.Title
	todo.Comment = mutationTodoRequest.Comment
	todo.UserId = userId

	// todo新規登録
	if err := ts.tr.CreateTodo(&todo); err != nil {
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("データ作成に失敗しました"), http.StatusInternalServerError)
		return models.BaseTodoResponse{}, err
	}

	// todo登録後のデータを取得
	if err := ts.tr.GetTodoLast(&todo, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("該当データは存在しません"), http.StatusBadRequest)
		} else {
			ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("データ取得に失敗しました"), http.StatusInternalServerError)
		}
		return models.BaseTodoResponse{}, err
	}

	// レスポンスデータを作成
	responseTodos := ts.tl.CreateTodoResponse(&todo)
	return responseTodos, nil
}

func (ts *todoService) DeleteTodo(w http.ResponseWriter, r *http.Request, userId int) error {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := ts.tr.DeleteTodo(id, userId); err != nil {
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("データ削除に失敗"), http.StatusInternalServerError)
		return err
	}
	return nil
}

func (ts *todoService) UpdateTodo(w http.ResponseWriter, r *http.Request, userId int) (models.BaseTodoResponse, error) {
	// GetパラメータからIDを取得
	vars := mux.Vars(r)
	id := vars["id"]
	// request bodyから値を取得
	reqBody, _ := io.ReadAll(r.Body)

	var mutationTodoRequest models.MutationTodoRequest
	if err := json.Unmarshal(reqBody, &mutationTodoRequest); err != nil {
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("リクエストパラメータを構造体へ変換処理でエラー発生"), http.StatusInternalServerError)
		return models.BaseTodoResponse{}, err
	}
	// バリデーション
	if err := ts.tv.MutationTodoValidate(mutationTodoRequest); err != nil {
		// バリデーションエラーのレスポンスを送信
		ts.rl.SendResponse(w, ts.rl.CreateErrorResponse(err), http.StatusBadRequest)
		return models.BaseTodoResponse{}, err
	}

	// 更新用データ用意
	var updateTodo models.Todo
	updateTodo.Title = mutationTodoRequest.Title
	updateTodo.Comment = mutationTodoRequest.Comment

	// todoデータ新規登録処理
	if err := ts.tr.UpdateTodo(&updateTodo, id, userId); err != nil {
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("データ更新に失敗しました。"), http.StatusInternalServerError)
		return models.BaseTodoResponse{}, err
	}

	// 更新データを取得
	var todo models.Todo
	if err := ts.tr.GetTodo(&todo, id, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("該当データは存在しません"), http.StatusBadRequest)
		} else {
			ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("データ取得に失敗しました"), http.StatusInternalServerError)
		}
		return models.BaseTodoResponse{}, err
	}

	// レスポンスデータを作成
	responseTodos := ts.tl.CreateTodoResponse(&todo)

	return responseTodos, nil
}

func (ts *todoService) SendAllTodoResponse(w http.ResponseWriter, todos *[]models.BaseTodoResponse) {
	var response models.AllTodoResponse
	// todos スライスのポインタをデリファレンス => アドレスが指す実際のデータにアクセスする操作
	response.Todos = *todos
	// response オブジェクトをJSON形式のバイトスライス[]byteに変換 => オブジェクトをJSONテキスト形式にシリアライズ（変換）し、その結果をバイトの配列（スライス）として取得するプロセス
	responseBody, _ := json.Marshal(response)
	ts.rl.SendResponse(w, responseBody, http.StatusOK)
}

func (ts *todoService) SendTodoResponse(w http.ResponseWriter, todo *models.BaseTodoResponse) {
	var response models.TodoResponse
	response.Todo = *todo
	responseBody, _ := json.Marshal(response)
	ts.rl.SendResponse(w, responseBody, http.StatusOK)
}

func (ts *todoService) SendCreateTodoResponse(w http.ResponseWriter, todo *models.BaseTodoResponse) {
	var response models.TodoResponse
	response.Todo = *todo
	responseBody, _ := json.Marshal(response)
	ts.rl.SendResponse(w, responseBody, http.StatusCreated)
}

func (ts *todoService) SendDeleteTodoResponse(w http.ResponseWriter) {
	ts.rl.SendNotBodyResponse(w)
}
