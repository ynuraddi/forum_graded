package sqlite

import (
	"context"
	"testing"

	"graded/model"
	"graded/pkg/errors"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	ctx := context.Background()
	userRepository := InitUserRepo(testDB)

	cases := []struct {
		name    string
		input   model.User
		wantInt int64
		wantErr error
	}{
		{
			name: "Create User",
			input: model.User{
				Login:    "aboba",
				Email:    "aboba@mail.ru",
				Password: "1234567890",
			},
			wantInt: 1,
			wantErr: nil,
		},
		{
			name: "Error Duplicate User",
			input: model.User{
				Login:    "aboba",
				Email:    "aboba@mail.ru",
				Password: "1234567890",
			},
			wantInt: -1,
			wantErr: errors.ErrDuplicateValue,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			getInt, getErr := userRepository.Create(ctx, test.input)
			assert.ErrorIs(t, getErr, test.wantErr)
			assert.Equal(t, getInt, test.wantInt)
		})
	}
}
