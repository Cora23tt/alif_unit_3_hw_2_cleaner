package sms

import (
	"testing"
)

func TestGenerateMessage(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		lang     string
		expected string
		err      bool
	}{
		{"Alice", "12345", "ru", "Добро пожаловать, Alice! Ваш код доступа: 12345.", false},
		{"Bob", "54321", "en", "Welcome, Bob! Your access code is: 54321.", false},
		{"", "12345", "ru", "", true},      // Missing name
		{"Alice", "", "ru", "", true},      // Missing code
		{"Al", "12345", "ru", "", true},    // Name too short
		{"Alice", "1234a", "ru", "", true}, // Code not numeric
		{"Alice", "12345", "fr", "", true}, // Unsupported language
	}

	for _, test := range tests {
		message, err := GenerateMessage(test.name, test.code, test.lang)

		if (err != nil) != test.err {
			t.Errorf("GenerateMessage(%q, %q, %q) error = %v, wantErr %v", test.name, test.code, test.lang, err, test.err)
			continue
		}

		if message != test.expected {
			t.Errorf("GenerateMessage(%q, %q, %q) = %q, want %q", test.name, test.code, test.lang, message, test.expected)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"12345", true},
		{"54321a", false},
		{"12.345", false},
		{"1234567890", true},
	}

	for _, test := range tests {
		result := isNumeric(test.input)

		if result != test.expected {
			t.Errorf("isNumeric(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
