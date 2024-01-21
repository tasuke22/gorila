package repositories

import (
	"github.com/stretchr/testify/assert"
	"github.com/tasuke/go-mux-task/models"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetUserByEmail(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIUserRepository(ctrl)

	// テスト対象のメソッドに渡す引数を設定
	user := &models.User{Email: "johndoe@example.com"}
	email := "johndoe@example.com"

	// メソッド呼び出しの期待値を設定
	mockRepo.EXPECT().GetUserByEmail(gomock.Any(), email).Return(nil).Times(1)

	// テスト対象のメソッドを実行
	err := mockRepo.GetUserByEmail(user, email)

	// アサーション
	assert.NoError(t, err)
}

func TestUserRepository_GetAllUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIUserRepository(ctrl)
	var users []models.User
	email := "johndoe@example.com"

	// 期待される動作の設定
	// 2人のユーザー（異なるEメールアドレス）を返すようにモックを設定
	mockRepo.EXPECT().GetAllUserByEmail(&users, email).Do(func(u *[]models.User, e string) {
		*u = append(*u, models.User{Email: "johndoe@example.com"}, models.User{Email: "johndoe2@example.com"})
	}).Return(nil).Times(1)

	// メソッドの実行
	err := mockRepo.GetAllUserByEmail(&users, email)

	// アサーション
	assert.NoError(t, err)
	assert.Len(t, users, 2) // 2人のユーザーが返されることを確認
	assert.Equal(t, "johndoe@example.com", users[0].Email)
	assert.Equal(t, "johndoe2@example.com", users[1].Email)
}

func TestUserRepository_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIUserRepository(ctrl)
	user := &models.User{Email: "johndoe@example.com"}

	// 期待される動作の設定
	mockRepo.EXPECT().CreateUser(user).Return(nil).Times(1)

	// メソッドの実行
	err := mockRepo.CreateUser(user)

	// アサーション
	assert.NoError(t, err)
}
