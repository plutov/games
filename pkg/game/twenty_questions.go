package game

// Reply func
func Reply(userInput string) (string, error) {
	userInput = sanitizeInput(userInput)

	// Note: all these conditions are just for testing purpose, not a final implementation
	if len(userInput) == 0 {
		return "Hi, my name is Alex. Do you want to play a game?", nil
	}
	if userInput == "no" {
		return "Got it. See you later, alligator.", nil
	}
	if userInput == "yes" || userInput == "what games do you have" {
		return "Today we have a game called '20 questions'. Let's play?", nil
	}
	if userInput == "can you explain me the rules of this game" {
		return "It's a simple and fun game where you should guess a noun by asking a maximum of 20 questions. Question can be answered only with 'Yes', 'No' and Don't know'.", nil
	}
	if userInput == "is it an animal" {
		return "Yes", nil
	}
	if userInput == "does it have two legs" {
		return "No", nil
	}
	if userInput == "is it a crocodile" {
		return "Yes, you are right! Congratulations! Want to play again?", nil
	}

	return "", nil
}
