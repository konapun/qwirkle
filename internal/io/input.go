package io

type Input[T any] interface {
	Read() T
}
