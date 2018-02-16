package main

// Do TCP
// Config
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

type networkType string

// The network must be a TCP network name
const (
	NetworkTsp  networkType = "tcp"
	NetworkTsp4 networkType = "tcp4"
	NetworkTsp6 networkType = "tcp6"
)
