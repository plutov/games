package ai

import "errors"

// Game struct
type Game struct {
	GreetingDone bool
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
}

// Reply func
func (gm *GameManager) Reply(sessionID string, userInput string) (string, error) {
	game, gameExists := gm.Games[sessionID]
	if !gameExists {
		return "", errors.New("Sorry, game not found, please try to refresh a page.")
	}

	userInput = sanitizeInput(userInput)

	if !game.GreetingDone {
		game.GreetingDone = true
		return "Hi, my name is Alex. Do you want to play a game?", nil
	}
	if inputBelongsToIndent(userInput, noIntent) {
		return "Negative answer!", nil
		//return "Got it. See you later, alligator.", nil
	}
	if inputBelongsToIndent(userInput, yesIntent) {
		return "Positive answer!", nil
		//return "Today we have a game called '20 questions'. Let's play?", nil
	}

	return "Sorry, what do you mean?", nil
}
