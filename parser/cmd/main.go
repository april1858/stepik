package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Создаем сканер для стандартного ввода

	comands := make(map[string]bool)
	var line string
	for scanner.Scan() {
		line = scanner.Text() // Получаем текущую строку

		if !strings.Contains(line, ";") { // Проверяем, содержит ли строка ';'
			break
		}

		if unicode.IsUpper(rune(line[0])) {
			fmt.Println("error")
		}

		line = line[:len(line)-1]
		parts := strings.Split(line, "=")

		if parts[1] == "" {
			parts[1] = "False"
		}

		v, err := strconv.ParseBool(parts[1])
		if err != nil {
			fmt.Println("parse bool ", err)
		}

		comands[parts[0]] = v

	}

	// Проверка на ошибки при сканировании
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении ввода:", err)
	}

	fmt.Println(comands)
	fmt.Println(line)
	fmt.Println(false == true)
}
