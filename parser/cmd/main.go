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
		line = scanner.Text()

		if !strings.Contains(line, ";") {
			break
		}

		if unicode.IsUpper(rune(line[0])) {
			fmt.Println("[error]")
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

	line = strings.ReplaceAll(line, " ", "")

	fmt.Println(groups(line))

	answer := one(comands, line)
	fmt.Println(answer)
}

func groups(line string) (int, []int) {
	sumL := 0
	sumR := 0
	group := 0
	ind := 0
	indexes := make([]int, 0)
	for i, r := range line {
		if r == '(' {
			sumL++
		}
		if r == ')' {
			sumR++
		}
		if sumL == sumR && sumL != 0 {
			indexes = append(indexes, i)
			ind = i
			sumL = 0
			sumR = 0
			group++

		}
		if sumL == 1 && sumR == 0 && i > ind {
			indexes = append(indexes, i)
		}
	}
	return group, indexes
}

func one(c map[string]bool, l string) bool {
	answer := false

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
