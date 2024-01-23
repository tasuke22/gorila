package validation

import (
	"github.com/stretchr/testify/assert"
	"github.com/tasuke/go-mux-task/models"
	"testing"
)

func TestTodoValidation_MutationTodoValidate(t *testing.T) {
	tests := []struct {
		name    string
		request models.MutationTodoRequest
		wantErr bool
	}{
		{
			name: "正常系 - タイトルとコメントが適切",
			request: models.MutationTodoRequest{
				Title:   "適切なタイトル",
				Comment: "適切なコメント",
			},
			wantErr: false,
		},
		{
			name: "異常系 - タイトルが空",
			request: models.MutationTodoRequest{
				Title:   "",
				Comment: "適切なコメント",
			},
			wantErr: true,
		},
		{
			name: "異常系 - タイトルが長すぎる",
			request: models.MutationTodoRequest{
				Title:   "これは非常に長いタイトルです",
				Comment: "適切なコメント",
			},
			wantErr: true,
		},
		{
			name: "異常系 - コメントが空",
			request: models.MutationTodoRequest{
				Title:   "適切なタイトル",
				Comment: "",
			},
			wantErr: true,
		},
	}

	tv := NewTodoValidation()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tv.MutationTodoValidate(tt.request)
			assert.Equal(t, tt.wantErr, err != nil, "MutationTodoValidate() の戻り値が期待と異なります")
		})
	}
}
