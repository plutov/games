package ai

import (
	"testing"
)

var intentCases = []struct {
	userInput string
	intent    Intent
	expected  bool
}{
	{"yes", yesIntent, true},
	{"no", noIntent, true},
	{"fuck yeah", yesIntent, false},
}

func TestSanitizeInput(t *testing.T) {
	if "i don't know" != sanitizeInput("I don't know ") {
		t.Errorf("sanitizeInput(%s) expected %s", "I don't know ", "i don't know")
	}
}

func TestInputBelongsToIndent(t *testing.T) {
	for _, i := range intentCases {
		actual := inputBelongsToIndent(i.userInput, i.intent)
		if actual != i.expected {
			t.Errorf("inputBelongsToIndent(%s, %v) expected %t, got %t", i.userInput, i.intent, i.expected, actual)
		}
	}
}
