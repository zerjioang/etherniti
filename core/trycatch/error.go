// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package trycatch

type Error string

var (
	nilErr = Error("")
)

func Nil() Error {
	return nilErr
}

func Ret(e error) Error {
	if e == nil {
		return nilErr
	}
	return New(e.Error())
}

func New(msg string) Error {
	return Error(msg)
}

func (stack Error) Error() string {
	return string(stack)
}

func (stack Error) Occur() bool {
	return stack != ""
}
func (stack Error) None() bool {
	return stack == ""
}
