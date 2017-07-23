package main

import (
	"os"

	"github.com/plutov/games/pkg/env"
	"github.com/plutov/games/pkg/shared"
	"github.com/plutov/games/pkg/web"
)

func main() {
	e, err := env.New()
	shared.LogErr(err)
	if err != nil {
		os.Exit(1)
	}

	web.ListenAndServe(e)
}
