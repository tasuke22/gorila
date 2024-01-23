package logic

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const wantUserID = 123

func Test_authLogic_GetUserIdFromContext(t *testing.T) {
	al := NewAuthLogic()

	godotenv.Load()

	secretKey := os.Getenv("JWT_KEY")
	testToken, _ := GenerateTestToken(secretKey)

	tests := []struct {
		name       string
		token      string
		wantUserID int
		wantErr    bool
	}{
		{
			name:       "valid token",
			token:      testToken,
			wantUserID: wantUserID,
			wantErr:    false,
		},
		{
			name:       "invalid token",
			token:      "invalid token",
			wantUserID: 0,
			wantErr:    true,
		},
		{
			name:       "empty token",
			token:      "",
			wantUserID: 0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("Authorization", "Bearer "+tt.token)

			gotUserID, err := al.GetUserIdFromContext(req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantUserID, gotUserID)
			}
		})
	}
}

func GenerateTestToken(secretKey string) (string, error) {
	// テストユーザーのクレームをセットアップ
	claims := jwt.MapClaims{
		"id":  wantUserID,
		"exp": time.Now().Add(time.Hour * 1).Unix(), // 1時間後に期限切れ
	}

	// JWTトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
