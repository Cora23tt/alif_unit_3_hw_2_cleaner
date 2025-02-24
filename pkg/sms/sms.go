package sms

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"unicode/utf8"
)

func Run() {
	name := flag.String("n", "", "Имя пользователя")
	code := flag.String("c", "", "Код доступа (5 цифр)")
	lang := flag.String("l", "ru", "Язык вывода (ru/en)")
	flag.Parse()

	message, err := GenerateMessage(*name, *code, *lang)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println(message)
}

func GenerateMessage(name, code, lang string) (string, error) {
	if name == "" || code == "" {
		return "", errors.New("недостаточно аргументов. Укажите имя и код доступа")
	}

	if utf8.RuneCountInString(name) < 3 {
		return "", errors.New("имя должно содержать не менее 3 символов")
	}

	if len(code) != 5 || !isNumeric(code) {
		return "", errors.New("код доступа должен состоять из 5 цифр")
	}

	var message string
	switch strings.ToLower(lang) {
	case "ru":
		message = fmt.Sprintf("Добро пожаловать, %s! Ваш код доступа: %s.", name, code)
	case "en":
		message = fmt.Sprintf("Welcome, %s! Your access code is: %s.", name, code)
	default:
		return "", errors.New("неподдерживаемый язык. Используйте 'ru' или 'en'")
	}

	return message, nil
}

func isNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
