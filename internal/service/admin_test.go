package service

import (
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"testing"
)

func TestAdminService_CheckCredentials(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	adminService := NewAdminService()

	tests := []struct {
		name         string
		email        string
		password     string
		expectResult bool
	}{
		{
			name:         "should success account admin",
			email:        "admin@kopoksu.com",
			password:     "adminkopoksu",
			expectResult: true,
		},
		{
			name:         "should success account master",
			email:        "master@kopoksu.com",
			password:     "masterkopoksu",
			expectResult: true,
		},
		{
			name:         "should failed wrong account",
			email:        "salah@kopoksu.com",
			password:     "salahkopoksu",
			expectResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := adminService.CheckCredentials(tt.email, tt.password)

			assert.Equal(t, result, tt.expectResult)
		})
	}

}
