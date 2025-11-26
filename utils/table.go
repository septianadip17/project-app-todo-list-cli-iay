package utils

import (
	"fmt"
	"strings"
)

func PrintTable(headers []string, rows [][]string) {
	widths := make([]int, len(headers))

	for i, h := range headers {
		widths[i] = len(h)
	}

	for _, row := range rows {
		for i, col := range row {
			if len(col) > widths[i] {
				widths[i] = len(col)
			}
		}
	}

	line := "+"
	for _, w := range widths {
		line += strings.Repeat("-", w+2) + "+"
	}

	fmt.Println(line)

	fmt.Print("|")
	for i, h := range headers {
		fmt.Printf(" %-*s |", widths[i], h)
	}
	fmt.Println()
	fmt.Println(line)

	for _, row := range rows {
		fmt.Print("|")
		for i, col := range row {
			fmt.Printf(" %-*s |", widths[i], col)
		}
		fmt.Println()
		fmt.Println(line)
	}
}
