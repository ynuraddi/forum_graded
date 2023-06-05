package sqlite

import (
	"context"
	"testing"

	"graded/model"
	"graded/pkg/erring"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	ctx := context.Background()
	userRepository := InitUserRepo(testDB)

	testUser := model.User{
		ID:       1,
		Login:    "aboba",
		Email:    "aboba@mail.ru",
		Password: "12345",
		IsActive: true,
		Version:  1,
	}

	// Create
	casesCreate := []struct {
		name    string
		input   model.User
		wantInt int64
		wantErr error
	}{
		// {
		// 	name:    "Create nil",
		// 	input:   testUser,
		// 	wantInt: 1,
		// 	wantErr: nil,
		// },
		{
			name:    "Create errDuplicate",
			input:   testUser,
			wantInt: -1,
			wantErr: erring.ErrDuplicateValue,
		},
	}

	for _, test := range casesCreate {
		t.Run(test.name, func(t *testing.T) {
			getInt, getErr := userRepository.Create(ctx, test.input)
			assert.ErrorIs(t, getErr, test.wantErr)
			assert.Equal(t, getInt, test.wantInt)
		})
	}

	// GetByID
	casesGetByID := []struct {
		name     string
		input    int64
		wantUser model.User
		wantErr  error
	}{
		{
			name:     "GetByID nil",
			input:    1,
			wantUser: testUser,
			wantErr:  nil,
		},
		{
			name:     "GetById noValue",
			input:    2,
			wantUser: model.User{},
			wantErr:  erring.ErrNoData,
		},
		{
			name:     "Negative ID",
			input:    -1,
			wantUser: model.User{},
			wantErr:  erring.ErrNoData,
		},
	}

	for _, test := range casesGetByID {
		t.Run(test.name, func(t *testing.T) {
			getUser, getErr := userRepository.GetByID(ctx, test.input)
			assert.ErrorIs(t, getErr, test.wantErr)
			assert.Equal(t, getUser, test.wantUser)
		})
	}
}
