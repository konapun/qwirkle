package state_test

import (
	. "github.com/konapun/quirkle/internal/state"
)

type TestModel struct {
	key string
}

func (t *TestModel) Key() string {
	return t.key
}

// FIXME: other should be a *TestModel type, not Model interface
func (t *TestModel) Diff(other Model) Diff {
	return Diff{}
}
