package ai

// Intent strcuct
type Intent struct {
	Name string
	// Samples can contain placeholders. Example: "log in with {id} and {region}"
	Samples []string
}

var yesIntent = Intent{
	Name:    "Yes",
	Samples: []string{"yes", "sure", "certainly", "no doubt", "for sure", "right", "all right", "okay", "true", "yea", "yeah", "fine", "of course", "yep", "aye", "good", "yes i want"},
}

var noIntent = Intent{
	Name:    "No",
	Samples: []string{"no", "not", "nope", "not at all", "no way", "no thanks", "next time", "nix", "i don't want", "no i don't", "no i'm not", "not today", "not really", "not this time"},
}
