package state_test

import (
	. "github.com/konapun/qwirkle/internal/state"
)

type TestModel struct {
	key    string
	String string
	Int    int
}

func NewTestModel(key string) *TestModel {
	return &TestModel{
		key: key,
	}
}

func (t *TestModel) Key() string {
	return t.key
}

func (t *TestModel) Diff(other *TestModel) Diff {
	// TODO: Implement the actual diff logic here
	return Diff{}
}

func (t *TestModel) Clone() *TestModel {
	return &TestModel{
		key:    t.key,
		String: t.String,
		Int:    t.Int,
	}
}
