package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/tasuke/go-mux-task/models"
)

type IAuthValidation interface {
	SignInValidate(signInRequest models.SignInRequest) error
	SignUpValidate(signUpRequest models.SignUpRequest) error
}

type authValidation struct{}

func NewAuthValidation() IAuthValidation {
	return &authValidation{}
}

var emailRules = []validation.Rule{
	validation.Required.Error("メールアドレスは必須入力です。"),
	validation.RuneLength(5, 40).Error("メールアドレスは 5～40 文字です"),
	is.Email.Error("メールアドレスの形式が間違っています。"),
}

var passwordRules = []validation.Rule{
	validation.Required.Error("パスワードは必須入力です。"),
	validation.RuneLength(6, 20).Error("パスワードは 6~20 文字です。"),
	is.Alphanumeric.Error("パスワードは英数字で入力してください。"),
}

// SignInValidate ログインパラメータのバリデーション
func (av *authValidation) SignInValidate(signInRequest models.SignInRequest) error {
	return validation.ValidateStruct(&signInRequest,
		validation.Field(
			&signInRequest.Email,
			emailRules...,
		),
		validation.Field(
			&signInRequest.Password,
			passwordRules...,
		),
	)
}

// SignUpValidate 会員登録パラメータのバリデーション
func (av *authValidation) SignUpValidate(signUpRequest models.SignUpRequest) error {
	return validation.ValidateStruct(&signUpRequest,
		validation.Field(
			&signUpRequest.Name,
			validation.Required.Error("お名前は必須入力です。"),
			validation.RuneLength(5, 10).Error("お名前は 5～10 文字です"),
		),
		validation.Field(
			&signUpRequest.Email,
			emailRules...,
		),
		validation.Field(
			&signUpRequest.Password,
			passwordRules...,
		),
	)
}
