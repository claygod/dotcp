package main

// Do TCP
// Server
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>
import (
	//"encoding/json"
	"fmt"
	//"log"
	"net"
	//"strconv"
	//"bufio"
	//"errors"
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
		fmt.Println(err, " !!")
		return []byte{}, err
	}

	// послать сообщение

	fmt.Println("Имеется массив - ")
	fmt.Println(msg)

	fmt.Println(" Клиент пытается отправить серверу:")
	fmt.Println(c.Write(msg))

	//time.Sleep(1000 * time.Millisecond)

	// получение сообщения
	buf := make([]byte, 0)
	fmt.Println(" Клиент  пытается получить:")

	for {
		xx := make([]byte, bufSize)
		res, err := c.Read(xx)
		fmt.Println("H3", res, err)
		buf = append(buf, xx[:res]...)

		if err != nil || res < bufSize { // res == 0 ||
			fmt.Println(" Клиент  получил:")
			fmt.Println(res, err)
			break
		}

	}
	fmt.Println("===----------------", buf)

	c.Close()
	return buf, nil
}
