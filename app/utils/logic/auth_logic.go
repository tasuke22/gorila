package logic

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strings"
)

type IAuthLogic interface {
	GetUserIdFromContext(r *http.Request) (int, error)
}

type authLogic struct{}

func NewAuthLogic() IAuthLogic {
	return &authLogic{}
}

func (al *authLogic) GetUserIdFromContext(r *http.Request) (int, error) {
	// リクエストヘッダーからトークンを取得 例：Authorization： Bearer eyJhbGciO...
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return 0, errors.New("認証トークンが提供されていません")
	}

	// Bearer トークンのフォーマットをチェック 例：Bearer eyJhbGciO...
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, errors.New("無効なトークンのフォーマットです")
	}

	// トークンの解析 ヘッダー、ペイロード、シグネチャーを検証している。シグネチャの検証には、シークレットキーが必要。
	token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(err)
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// クレームからユーザーIDを取得
		if userId, ok := claims["id"].(float64); ok {
			return int(userId), nil
		}
	}

	return 0, errors.New("トークンからユーザーIDを抽出できません")
}
