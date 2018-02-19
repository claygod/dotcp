package main

// Do TCP
// Config
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"time"
)

type networkType string
type errCode byte

const deadline time.Duration = 100 * time.Millisecond // millisecond
const bufSize int = 1024

// The network must be a TCP network name
const (
	NetworkTsp  networkType = "tcp"
	NetworkTsp4 networkType = "tcp4"
	NetworkTsp6 networkType = "tcp6"
)

// Reply codes
const (
	Ok byte = iota
	Err
	ErrUnmarshal
	ErrUnregistered
	ErrDial
	ErrSendingMsg
	ErrGettingMsg
	ErrOther
)
