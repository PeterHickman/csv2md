package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func header(text string) string {
	parts := strings.Split(text, ",")
	return "|" + strings.Repeat("---|", len(parts))
}

func convert(text string) string {
	parts := strings.Split(text, ",")
	return "|" + strings.Join(parts, "|") + "|"
}

func main() {
	firstLine := true

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		fmt.Println(convert(text))
		if firstLine {
			firstLine = false
			fmt.Println(header(text))
		}
	}
}
