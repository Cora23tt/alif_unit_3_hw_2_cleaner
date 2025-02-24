package cleaner // Aziz Rustamov

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"unicode"
)

func Run() error {
	phones := flag.String("p", "", "Phone numbers separated by comma")
	flag.Parse()

	if *phones == "" {
		return errors.New("номер телефона не указан")
	}

	numbers := strings.Split(*phones, ",")
	for _, number := range numbers {
		normalized, err := normalizePhoneNumber(strings.TrimSpace(number))
		if err != nil {
			fmt.Printf("Ошибка для номера '%s': %v\n", number, err)
			continue
		}
		fmt.Println("Нормализованный номер:", normalized)
	}

	return nil
}

func normalizePhoneNumber(phone string) (string, error) {
	var cleaned strings.Builder
	hasPlus := false

	// Цикл по каждому символу строки ввода
	for _, r := range phone {
		if unicode.IsDigit(r) {
			cleaned.WriteRune(r)
		} else if r == '+' && !hasPlus {
			cleaned.WriteRune(r)
			hasPlus = true
		}
	}

	result := cleaned.String()

	// проверить нормализованный номер
	if !strings.HasPrefix(result, "+998") {
		return "", errors.New("номер должен начинаться с +998")
	}
	if len(result) != 13 { // +998 and 12 digits
		return "", errors.New("номер должен содержать 12 цифр")
	}

	return result, nil
}
