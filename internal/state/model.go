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

func NewDiff() Diff {
	return Diff{
		changed: make(map[string]Change),
	}
}

func (d *Diff) SetChanged(field string, from any, to any) {
	d.changed[field] = Change{old: from, new: to}
}

func (d *Diff) HasChanged(field string) bool {
	_, ok := d.changed[field]
	return ok
}

func (d *Diff) GetChangeSet() map[string]Change {
	return d.changed
}

type Model[T any] interface {
	state.Model[T]
	Diff(other T) Diff
}
