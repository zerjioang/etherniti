// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package trycatch

type Error struct {
	err string
}

var (
	nilErr = Error{}
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
	return Error{err: msg}
}

func (stack Error) Error() string {
	return stack.err
}

func (stack Error) Occur() bool {
	return stack.err != ""
}
func (stack Error) None() bool {
	return stack.err == ""
}
