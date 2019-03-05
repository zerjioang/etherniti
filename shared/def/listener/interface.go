// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package listener

type ListenerInterface interface {
	Listen() error
	RunMode(string, bool)
}

type ServiceType uint8

const (
	HttpMode ServiceType = iota
	UnixMode
	UnknownMode
)
