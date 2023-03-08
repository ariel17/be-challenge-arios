package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariel17/be-challenge-arios/pkg/repositories"
)

func TestRealStatusService_GetStatus(t *testing.T) {

	testCases := []struct {
		name      string
		isSuccess bool
	}{
		{"ok", true},
		{"failed", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := repositories.MockRepository{}

			service := realStatusService{
				Repository: &mockRepository,
			}

			on := mockRepository.On("GetStatus")
			if tc.isSuccess {
				on.Return(nil)
			} else {
				on.Return(errors.New("mocked error"))
			}

			err := service.GetStatus()
			assert.Equal(t, err == nil, tc.isSuccess)
		})
	}
}