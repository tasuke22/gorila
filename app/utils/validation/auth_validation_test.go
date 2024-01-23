package validation

import (
	"github.com/tasuke/go-mux-task/models"
	"testing"
)

func Test_authValidation_SignUpValidate(t *testing.T) {
	tests := []struct {
		name    string
		request models.SignUpRequest
		wantErr bool
	}{
		{
			name: "正常系 - 全てのフィールドが適切",
			request: models.SignUpRequest{
				Name:     "ValidName",
				Email:    "valid@example.com",
				Password: "ValidPass123",
			},
			wantErr: false,
		},
		{
			name: "異常系 - Nameが短すぎる",
			request: models.SignUpRequest{
				Name:     "Abc",
				Email:    "valid@example.com",
				Password: "ValidPass123",
			},
			wantErr: true,
		},
		{
			name: "異常系 - Nameが長すぎる",
			request: models.SignUpRequest{
				Name:     "ThisNameIsWayTooLongForTheField",
				Email:    "valid@example.com",
				Password: "ValidPass123",
			},
			wantErr: true,
		},
		{
			name: "異常系 - Emailが長すぎる",
			request: models.SignUpRequest{
				Name:     "ValidName",
				Email:    "verylongemailaddressssssssssssssssssssss@example.com",
				Password: "ValidPass123",
			},
			wantErr: true,
		},
		{
			name: "異常系 - Passwordが長すぎる",
			request: models.SignUpRequest{
				Name:     "ValidName",
				Email:    "valid@example.com",
				Password: "ThisPasswordIsWayTooLong",
			},
			wantErr: true,
		},
		{
			name: "異常系 - Emailにドメインが含まれていない",
			request: models.SignUpRequest{
				Name:     "ValidName",
				Email:    "invalidemail@",
				Password: "ValidPass123",
			},
			wantErr: true,
		},
		{
			name: "異常系 - Passwordが英数字以外を含む",
			request: models.SignUpRequest{
				Name:     "ValidName",
				Email:    "valid@example.com",
				Password: "Password#123",
			},
			wantErr: true,
		},
		{
			name: "異常系 - Emailが無効",
			request: models.SignUpRequest{
				Name:     "ValidName",
				Email:    "invalid-email",
				Password: "ValidPass123",
			},
			wantErr: true,
		},
		{
			name: "異常系 - Passwordが短すぎる",
			request: models.SignUpRequest{
				Name:     "ValidName",
				Email:    "valid@example.com",
				Password: "123",
			},
			wantErr: true,
		},
	}

	av := NewAuthValidation()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := av.SignUpValidate(tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUpValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
