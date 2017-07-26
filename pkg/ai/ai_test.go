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

func TestInitSession(t *testing.T) {
	gm := NewGameManager()
	if nil == gm {
		t.Errorf("NewGameManager() expected non-nil return value")
	}
	gm.InitSession("test")
	if len(gm.Games) != 1 {
		t.Errorf("Games expected to have 1 item")
	}
	if nil == gm.GetGameBySessionID("test") {
		t.Errorf("GetGameBySessionID(test) expected non-nil return value")
	}
}
