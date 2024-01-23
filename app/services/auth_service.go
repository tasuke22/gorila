//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package services

import (
	"encoding/json"
	"errors"
	"github.com/tasuke/go-mux-task/models"
	"github.com/tasuke/go-mux-task/repositories"
	"github.com/tasuke/go-mux-task/utils/logic"
	"github.com/tasuke/go-mux-task/utils/validation"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
)

type IAuthService interface {
	GetUserIdFromToken(w http.ResponseWriter, r *http.Request) (int, error)
	SignIn(w http.ResponseWriter, r *http.Request) (models.User, error)
	SignUp(w http.ResponseWriter, r *http.Request) (models.User, error)
	SendAuthResponse(w http.ResponseWriter, user *models.User, code int)
}

type authService struct {
	ur repositories.IUserRepository
	al logic.IAuthLogic
	rl logic.IResponseLogic
	jl logic.IJWTLogic
	av validation.IAuthValidation
}

func NewAuthService(
	ur repositories.IUserRepository,
	al logic.IAuthLogic,
	rl logic.IResponseLogic,
	jl logic.IJWTLogic,
	av validation.IAuthValidation,
) IAuthService {
	return &authService{ur, al, rl, jl, av}
}

func (as authService) GetUserIdFromToken(w http.ResponseWriter, r *http.Request) (int, error) {
	// トークンからユーザーIDを取得
	userId, err := as.al.GetUserIdFromContext(r)
	if err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("トークンからユーザーIDを抽出できません"), http.StatusUnauthorized)
		return 0, err
	}
	return userId, nil
}

func (as authService) SignIn(w http.ResponseWriter, r *http.Request) (models.User, error) {
	// リクエストボディを読み込み
	reqBody, _ := io.ReadAll(r.Body)
	var signInRequestParam models.SignInRequest

	// リクエストパラメータのJSONを構造体にマッピング
	if err := json.Unmarshal(reqBody, &signInRequestParam); err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("リクエストパラメータを構造体へ変換処理でエラー発生"), http.StatusInternalServerError)
		return models.User{}, err
	}

	// リクエストパラメータのバリデーション
	if err := as.av.SignInValidate(signInRequestParam); err != nil {
		// バリデーションエラーの場合はエラーレスポンスを返却
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse(err.Error()), http.StatusBadRequest)
		return models.User{}, err
	}

	// ユーザーの存在確認
	var user models.User
	if err := as.ur.GetUserByEmail(&user, signInRequestParam.Email); err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("ユーザーが存在しません"), http.StatusUnauthorized)
		return models.User{}, err
	}

	// パスワードの一致確認
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInRequestParam.Password)); err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("パスワードが間違っています。"), http.StatusUnauthorized)
		return models.User{}, err
	}

	return user, nil
}

func (as authService) SignUp(w http.ResponseWriter, r *http.Request) (models.User, error) {
	// リクエストボディを読み込み
	reqBody, _ := io.ReadAll(r.Body)
	log.Printf("reqBody: %v", string(reqBody))
	var signUpRequestParam models.SignUpRequest
	if err := json.Unmarshal(reqBody, &signUpRequestParam); err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("リクエストパラメータを構造体へ変換処理でエラー発生"), http.StatusInternalServerError)
		return models.User{}, err
	}

	// リクエストパラメータのバリデーション
	if err := as.av.SignUpValidate(signUpRequestParam); err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorResponse(err), http.StatusBadRequest)
		return models.User{}, err
	}

	// emailに紐づくユーザーをチェック, ユーザーの重複チェック
	var users []models.User
	if err := as.ur.GetAllUserByEmail(&users, signUpRequestParam.Email); err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("DBエラー"), http.StatusInternalServerError)
		return models.User{}, err
	}

	if len(users) != 0 {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("入力されたメールアドレスは既に登録されています。"), http.StatusUnauthorized)
		return models.User{}, errors.New("ユーザーは既に登録されています。")
	}

	// パスワードのハッシュ化して登録データを作成
	var newUser models.User
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(signUpRequestParam.Password), bcrypt.DefaultCost)
	newUser.Name = signUpRequestParam.Name
	newUser.Email = signUpRequestParam.Email
	newUser.Password = string(hashPassword)

	// ユーザー登録処理
	if err := as.ur.CreateUser(&newUser); err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("ユーザー登録処理に失敗"), http.StatusInternalServerError)
		return models.User{}, err
	}

	return newUser, nil
}

func (as authService) SendAuthResponse(w http.ResponseWriter, user *models.User, code int) {
	// JWTトークンの作成
	token, err := as.jl.CreateJwtToken(user)
	if err != nil {
		as.rl.SendResponse(w, as.rl.CreateErrorStringResponse("トークン作成処理に失敗"), http.StatusInternalServerError)
		return
	}

	// レスポンスボディの作成
	var response models.AuthResponse
	response.Token = token
	response.User.ID = user.ID
	response.User.Name = user.Name
	response.User.Email = user.Email
	response.User.CreatedAt = user.CreatedAt
	response.User.UpdatedAt = user.UpdatedAt
	response.User.DeletedAt = user.DeletedAt

	// レスポンスの送信
	responseBody, _ := json.Marshal(response)
	as.rl.SendResponse(w, responseBody, code)
}
