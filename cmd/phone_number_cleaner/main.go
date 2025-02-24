// Aziz Rustamov 
// Урок 3: Компиляция и пакеты
// ДЗ №2: Программа "Cleaner"
package main

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
