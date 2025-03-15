package scene

import "errors"

var (
	ErrInvalidTransition = errors.New("invalid transition")
	ErrInvalidAction     = errors.New("invalid action")
	ErrNoPlayers         = errors.New("no players")
)
