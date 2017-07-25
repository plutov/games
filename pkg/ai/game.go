package ai

import (
	"errors"
	"time"
)

// Game struct
type Game struct {
	GreetingDone    bool
	UserWantsToPlay bool
	Started         bool
	Completed       bool
	Canceled        bool
	CreatedAt       time.Time
}

// GameManager struct
type GameManager struct {
	Games map[string]*Game
}

// NewGameManager func is initialized in env.go
func NewGameManager() *GameManager {
	gm := new(GameManager)
	gm.Games = make(map[string]*Game)

	return gm
}

// InitSession func
func (gm *GameManager) InitSession(sessionID string) {
	gm.Games[sessionID] = new(Game)
	gm.Games[sessionID].CreatedAt = time.Now()
}

// Reply func
func (gm *GameManager) Reply(sessionID string, userInput string) (string, error) {
	game, gameExists := gm.Games[sessionID]
	if !gameExists {
		return "", errors.New("Sorry, game not found, please try to refresh a page.")
	}

	if game.Canceled {
		return "", nil
	}

	userInput = sanitizeInput(userInput)

	if !game.GreetingDone {
		game.GreetingDone = true
		return "Hi, my name is Alex. Do you want to play a game?", nil
	}

	if inputBelongsToIndent(userInput, yesIntent) {
		if !game.UserWantsToPlay {
			game.UserWantsToPlay = true
			return "I have a game called '20 questions'. Let's play?", nil
		}
		if game.UserWantsToPlay && !game.Started {
			game.Started = true
			return "It's a simple fun game where you should guess a noun by asking a maximum of 20 questions. Question can be answered only with Yes, No and Don't know. Try to ask something!", nil
		}
	}

	if inputBelongsToIndent(userInput, noIntent) {
		if !game.UserWantsToPlay {
			game.Canceled = true
			return "Got it! See you later, alligator!", nil
		}
	}

	return "Sorry, what do you mean?", nil
}
