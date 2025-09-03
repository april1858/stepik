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
			return
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

	answer := parse(comands, line)
	fmt.Println(answer)
}

func parse(c map[string]bool, l string) bool {
	answer := false
	l = strings.ReplaceAll(l, " ", "")
	switch {
	case strings.Contains(l, "and"):
		index := strings.Index(l, "and")
		answer = c[l[:index]] && c[l[index+3:]]
	case strings.Contains(l, "or"):
		index := strings.Index(l, "or")
		answer = c[l[:index]] || c[l[index+2:]]
	case strings.Contains(l, "xor"):
		index := strings.Index(l, "xor")
		answer = c[l[:index]] != c[l[index+3:]]
	case strings.Contains(l, "not"):
		answer = !c[l[3:]]
	}
	return answer
}
