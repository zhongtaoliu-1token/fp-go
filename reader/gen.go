// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2023-09-12 10:51:22.8683246 +0200 CEST m=+0.010727801

package reader

import (
	G "github.com/IBM/fp-go/reader/generic"
)

// From0 converts a function with 1 parameters returning a [R] into a function with 0 parameters returning a [Reader[C, R]]
// The first parameter is considered to be the context [C] of the reader
func From0[F ~func(C) R, C, R any](f F) func() Reader[C, R] {
	return G.From0[Reader[C, R]](f)
}

// From1 converts a function with 2 parameters returning a [R] into a function with 1 parameters returning a [Reader[C, R]]
// The first parameter is considered to be the context [C] of the reader
func From1[F ~func(C, T0) R, T0, C, R any](f F) func(T0) Reader[C, R] {
	return G.From1[Reader[C, R]](f)
}

// From2 converts a function with 3 parameters returning a [R] into a function with 2 parameters returning a [Reader[C, R]]
// The first parameter is considered to be the context [C] of the reader
func From2[F ~func(C, T0, T1) R, T0, T1, C, R any](f F) func(T0, T1) Reader[C, R] {
	return G.From2[Reader[C, R]](f)
}

// From3 converts a function with 4 parameters returning a [R] into a function with 3 parameters returning a [Reader[C, R]]
// The first parameter is considered to be the context [C] of the reader
func From3[F ~func(C, T0, T1, T2) R, T0, T1, T2, C, R any](f F) func(T0, T1, T2) Reader[C, R] {
	return G.From3[Reader[C, R]](f)
}

// From4 converts a function with 5 parameters returning a [R] into a function with 4 parameters returning a [Reader[C, R]]
// The first parameter is considered to be the context [C] of the reader
func From4[F ~func(C, T0, T1, T2, T3) R, T0, T1, T2, T3, C, R any](f F) func(T0, T1, T2, T3) Reader[C, R] {
	return G.From4[Reader[C, R]](f)
}

// From5 converts a function with 6 parameters returning a [R] into a function with 5 parameters returning a [Reader[C, R]]
// The first parameter is considered to be the context [C] of the reader
func From5[F ~func(C, T0, T1, T2, T3, T4) R, T0, T1, T2, T3, T4, C, R any](f F) func(T0, T1, T2, T3, T4) Reader[C, R] {
	return G.From5[Reader[C, R]](f)
}

// From6 converts a function with 7 parameters returning a [R] into a function with 6 parameters returning a [Reader[C, R]]
// The first parameter is considered to be the context [C] of the reader
func From6[F ~func(C, T0, T1, T2, T3, T4, T5) R, T0, T1, T2, T3, T4, T5, C, R any](f F) func(T0, T1, T2, T3, T4, T5) Reader[C, R] {
	return G.From6[Reader[C, R]](f)
}

// From7 converts a function with 8 parameters returning a [R] into a function with 7 parameters returning a [Reader[C, R]]
// The first parameter is considered to be the context [C] of the reader
func From7[F ~func(C, T0, T1, T2, T3, T4, T5, T6) R, T0, T1, T2, T3, T4, T5, T6, C, R any](f F) func(T0, T1, T2, T3, T4, T5, T6) Reader[C, R] {
	return G.From7[Reader[C, R]](f)
}

// From8 converts a function with 9 parameters returning a [R] into a function with 8 parameters returning a [Reader[C, R]]
// The first parameter is considered to be the context [C] of the reader
func From8[F ~func(C, T0, T1, T2, T3, T4, T5, T6, T7) R, T0, T1, T2, T3, T4, T5, T6, T7, C, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7) Reader[C, R] {
	return G.From8[Reader[C, R]](f)
}

// From9 converts a function with 10 parameters returning a [R] into a function with 9 parameters returning a [Reader[C, R]]
// The first parameter is considered to be the context [C] of the reader
func From9[F ~func(C, T0, T1, T2, T3, T4, T5, T6, T7, T8) R, T0, T1, T2, T3, T4, T5, T6, T7, T8, C, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8) Reader[C, R] {
	return G.From9[Reader[C, R]](f)
}

// From10 converts a function with 11 parameters returning a [R] into a function with 10 parameters returning a [Reader[C, R]]
// The first parameter is considered to be the context [C] of the reader
func From10[F ~func(C, T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) R, T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, C, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) Reader[C, R] {
	return G.From10[Reader[C, R]](f)
}
