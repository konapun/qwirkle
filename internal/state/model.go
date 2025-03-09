package state

import (
	"github.com/konapun/statekit/state"
)

type Change struct {
	old any
	new any
}

type Diff struct {
	// Define the methods for the Diff type
	changed map[string]Change
}

func (d *Diff) SetChanged(field string, from any, to any) {
	d.changed[field] = Change{old: from, new: to}
}

type Model[T any] interface {
	state.Model[T]
	Diff(other T) Diff
}
