package ai

import (
	"strings"
)

func sanitizeInput(userInput string) string {
	return strings.TrimSpace(strings.ToLower(userInput))
}

// Search in Intent's samples to determine what userInput means
func inputBelongsToIndent(userInput string, intent Intent) bool {
	for _, s := range intent.Samples {
		if s == userInput {
			return true
		}
	}

	return false
}
