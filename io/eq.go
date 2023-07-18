package io

import (
	EQ "github.com/IBM/fp-go/eq"
	G "github.com/IBM/fp-go/io/generic"
)

// Eq implements the equals predicate for values contained in the IO monad
func Eq[A any](e EQ.Eq[A]) EQ.Eq[IO[A]] {
	return G.Eq[IO[A]](e)
}