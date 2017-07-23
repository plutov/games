package env

import "github.com/plutov/games/pkg/config"

// Context struct
type Context struct {
	Config *config.Structure
}

// New func
func New() (Context, error) {
	c, configErr := config.Get()
	if configErr != nil {
		return Context{}, configErr
	}

	return Context{
		Config: c,
	}, nil
}
