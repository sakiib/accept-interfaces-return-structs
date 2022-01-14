package user_test

import (
	"testing"

	"github.com/sakiib/go-pattern/db"
	"github.com/sakiib/go-pattern/user"
)

func TestCreateUser(t *testing.T) {
	mp := make(map[string]string)
	store, _ := db.NewStore(mp)

	service := user.NewUserService(store)

	testCases := []struct {
		name          string
		key           string
		value         string
		expectedError bool
	}{
		{
			name:          "user should be created successfully, expectedError nil",
			key:           "x",
			value:         "y",
			expectedError: false,
		},
		{
			name:          "user shouldn't be created successfully, expectedError not nil",
			key:           "x",
			value:         "",
			expectedError: true,
		},
		{
			name:          "user shouldn't be created successfully, expectedError not nil",
			key:           "",
			value:         "y",
			expectedError: true,
		},
		{
			name:          "user shouldn't be created successfully, expectedError not nil",
			key:           "",
			value:         "",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := service.CreateUser(tc.key, tc.value); err != nil && tc.expectedError == false {
				t.Errorf("error: %v, expected: %v", err, tc.expectedError)
			}
		})
	}
}

func TestRetrieveUser(t *testing.T) {
	mp := make(map[string]string)
	store, _ := db.NewStore(mp)

	service := user.NewUserService(store)

	service.CreateUser("x", "y")

	testCases := []struct {
		name  string
		key   string
		value string
	}{
		{
			name:  "user should be retrieved successfully, expectedError nil",
			key:   "x",
			value: "y",
		},
		{
			name:  "user shouldn't be retrieved successfully, expectedError not nil",
			key:   "z",
			value: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := service.RetrieveUser(tc.key); tc.value != got {
				t.Errorf("error: %v, got: %v, expected: %v", err, got, tc.value)
			}
		})
	}
}
