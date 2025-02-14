package main // Aziz Rustamov

import (
	"fmt"
	"unit_3_hw_2_cleaner/pkg/cleaner"
)

func main() {
	err := cleaner.Run()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
