package ai

// Reply func
func Reply(userInput string) (string, error) {
	userInput = sanitizeInput(userInput)

	// Note: all these conditions are just for testing purpose, not a final implementation
	if len(userInput) == 0 {
		return "Hi!", nil
		//return "Hi, my name is Alex. Do you want to play a game?", nil
	}
	if inputBelongsToIndent(userInput, noIntent) {
		return "Negative answer!", nil
		//return "Got it. See you later, alligator.", nil
	}
	if inputBelongsToIndent(userInput, yesIntent) {
		return "Positive answer!", nil
		//return "Today we have a game called '20 questions'. Let's play?", nil
	}

	return "", nil
}
