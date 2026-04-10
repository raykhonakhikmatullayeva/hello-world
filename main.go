package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func normalize(input string) (string, error) {
	result := strings.ReplaceAll(input, " ", "")
	result = strings.ReplaceAll(result, "-", "")

	for _, ch := range result {
		if ch < '0' || ch > '9' {
			return "", fmt.Errorf("номер содержит недопустимые символы")
		}
	}

	if len(result) != 16 {
		return "", fmt.Errorf("неверный формат карты")
	}

	return result, nil
}

func mask(number string) string {
	return number[0:4] + " **** **** " + number[12:16]
}

func main() {
	cardFlag := flag.String("card", "", "номер карты")
	flag.Parse()

	var input string

	if *cardFlag != "" {
		input = *cardFlag
	} else {
		fmt.Print("Введите номер карты: ")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n') // reads until newline
		input = strings.TrimSpace(line)    // strip the trailing \n (and \r on Windows)
	}

	number, err := normalize(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Нормализованный номер:", number)
	fmt.Println("Длина:", len(number), "✓")

	fmt.Println("Карта:", mask(number))
}
