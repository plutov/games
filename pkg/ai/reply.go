package ai

import "math/rand"

var cantRecognizeReply = []string{"Sorry, what do you mean?", "What did you say?", "What do you mean?", "I don't understand.", "Excuse me, I didn't get it.", "Excuse me, can you please repeat it?", "Sorry, I did not catch that.", "I missed that.", "Can you please speak slowly?", "I don't get it."}

var byeReply = []string{"Got it! See you later, alligator!", "Got it! Bye!", "See you again!", "See you later", "Okay, see you later!"}

func chooseReply(replies []string) string {
	n := rand.Int() % len(replies)
	return replies[n]
}
