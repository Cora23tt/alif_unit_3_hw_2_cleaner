package cleaner

import (
	"testing"
)

func TestNormalizePhoneNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"+998901234567", "+998901234567", false},
		{"+998 (90) 123-45-67", "+998901234567", false},
		{"998901234567", "", true},                     // Нет +
		{"901234567", "", true},                        // Нет кода страны
		{"+997901234567", "", true},                    // Неверный код страны
		{"+99890123456", "", true},                     // Недостаточно цифр
		{"+9989012345678", "", true},                   // Лишняя цифра
		{"", "", true},                                 // Пустая строка
		{" , +998901234567, ", "+998901234567", false}, // Лишние пробелы и запятые
	}

	for _, test := range tests {
		result, err := normalizePhoneNumber(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Ошибка на входе %q: ожидалась ошибка %v, но получили %v", test.input, test.hasError, err)
		}
		if result != test.expected {
			t.Errorf("Ошибка на входе %q: ожидали %q, но получили %q", test.input, test.expected, result)
		}
	}
}
