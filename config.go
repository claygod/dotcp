package dotcp

// Do TCP
// Config
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

type networkType string
type errCode byte

// const deadline time.Duration = 100 * time.Millisecond // millisecond
// const startPause int = 100
const bufSize int = 1024
const portsLimitMax int = 65535
const portsLimitMin int = 0

// The network must be a TCP network name
const (
	NetworkTsp  networkType = "tcp"
	NetworkTsp4 networkType = "tcp4"
	NetworkTsp6 networkType = "tcp6"
)

// Reply codes
const (
	reOk byte = iota
	reError
)
