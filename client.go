package main

// Do TCP
// Server
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>
import (
	//"encoding/json"
	//"fmt"
	//"log"
	"net"
	//"strconv"
	//"bufio"
	"errors"
	//"os"
	//"time"
)

/*
tspClient - TCP server.
*/
type tspClient struct {
	network string
	addr    *net.TCPAddr
}

/*
Client - create a new tcp client.
*/
func Client(network networkType) *tspClient { // TCPAddr
	ts := &tspClient{
		network: string(network),
		addr:    new(net.TCPAddr),
	}
	return ts
}

/*
IP - set IP.
*/
func (t *tspClient) IP(ip net.IP) *tspClient {
	t.addr.IP = ip
	return t
}

/*
Port - set Port.
*/
func (t *tspClient) Port(port int) *tspClient {
	t.addr.Port = port
	return t
}

/*
Zone - set Zone.
*/
func (t *tspClient) Zone(zone string) *tspClient {
	t.addr.Zone = zone
	return t
}

/*
Send - send a message to the server.
*/
func (t *tspClient) Send(msg []byte) ([]byte, error) {

	c, err := net.Dial(t.network, t.addr.String()) // "127.0.0.1:9999"
	if err != nil {
		return []byte{}, err
	}
	defer c.Close()
	// послать сообщение

	n, err := c.Write(msg)
	if err != nil {
		return []byte{}, err
	}
	if n != len(msg) {
		return []byte{}, errors.New("The message has not been sent in full")
	}

	// получение сообщения
	buf := make([]byte, 0)

	for {
		xx := make([]byte, bufSize)
		res, err := c.Read(xx)
		buf = append(buf, xx[:res]...)

		if err != nil || res < bufSize { // res == 0 ||
			break
		}
	}
	// Анализ полученноего сообщения
	if buf[0] == Ok {
		return buf[1:], nil
	}
	str := string(buf[1:])
	return []byte{}, errors.New(str)
}
