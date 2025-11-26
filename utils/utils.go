package utils

import "fmt"

func PrintError(msg string) {
	fmt.Println(msg)
}

func IsValidStatus(s string) bool {
	switch s {
	case "new", "progress", "completed":
		return true
	default:
		return false
	}
}
