package ai

import (
	"errors"
	"time"
)

// Game struct
type Game struct {
	GreetingDone    bool      `json:"greeting_done"`
	UserWantsToPlay bool      `json:"user_wants_to_play"`
	Started         bool      `json:"started"`
	Completed       bool      `json:"completed"`
	Canceled        bool      `json:"canceled"`
	CreatedAt       time.Time `json:"created_at"`
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

// GetGameBySessionID func
func (gm *GameManager) GetGameBySessionID(sessionID string) *Game {
	game, gameExists := gm.Games[sessionID]
	if !gameExists {
		return nil
	}

	return game
}

// Reply func
func (gm *GameManager) Reply(sessionID string, userInput string) (string, error) {
	game := gm.GetGameBySessionID(sessionID)
	if nil == game || game.Canceled {
		return "", errors.New("Sorry, game not found, please try to refresh a page.")
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
		if !game.UserWantsToPlay || !game.Started {
			game.Canceled = true
			return chooseReply(byeReply), nil
		}
	}

	return chooseReply(cantRecognizeReply), nil
}
