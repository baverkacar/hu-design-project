package util

import (
	"testing"
)

func Test_createVerificationCode(t *testing.T) {
	code, err := createVerificationCode()
	if err != nil {
		t.Fatalf("createVerificationCode() failed: %v", err)
	}
	if len(code) != 6 {
		t.Errorf("Expected code to have 6 digits, got %d", len(code))
	}
	for _, r := range code {
		if r < '0' || r > '9' {
			t.Errorf("Expected code to contain only digits, got %s", code)
			break
		}
	}
}
