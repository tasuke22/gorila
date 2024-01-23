package logic

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/tasuke/go-mux-task/models"
	"os"
	"time"
)

type IJWTLogic interface {
	CreateJwtToken(user *models.User) (string, error)
}

type jwtLogic struct{}

func NewJWTLogic() IJWTLogic {
	return &jwtLogic{}
}

// CreateJwtToken jwtトークンの新規作成
func (jl *jwtLogic) CreateJwtToken(user *models.User) (string, error) {
	// トークンの生成
	token := jwt.New(jwt.SigningMethodHS256)

	// 環境変数を取得
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	// Claimsのセットとトークンの生成
	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // 1時間後に期限切れ
	}
	// トークンにClaimsをセット
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンの署名
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", fmt.Errorf("failed to sign the token: %w", err)
	}

	return tokenString, err
}
