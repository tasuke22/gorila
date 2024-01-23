package logic

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/tasuke/go-mux-task/models"
	"os"
	"testing"
)

func TestCreateJwtToken(t *testing.T) {
	jl := NewJWTLogic()
	user := &models.User{
		Name: "Test User",
	}

	// テスト用の固定キー
	testKey := "test_jwt_key"

	// 環境変数の設定（テスト用）
	os.Setenv("JWT_KEY", testKey)

	tokenString, err := jl.CreateJwtToken(user)
	assert.NoError(t, err, "CreateJwtToken should not return an error")
	assert.NotEmpty(t, tokenString, "Token string should not be empty")

	// トークンの解析と検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(testKey), nil
	})
	assert.NoError(t, err, "Failed to parse the token")
	assert.True(t, token.Valid, "The token should be valid")

	// クレームの検証
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		assert.Equal(t, user.Name, claims["name"], "Name should match")
	}
}
