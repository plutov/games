package game

import (
	"strings"
)

func sanitizeInput(userInput string) string {
	return strings.ToLower(userInput)
}
