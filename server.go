package main

// Do TCP
// Server
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>
import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func server() {
	// слушать порт
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	//for {
	// принятие соединения
	c, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		//continue
	}
	// обработка соединения
	c.LocalAddr()
	go handleServerConnection(c)
	//}
}

func handleServerConnection(c net.Conn) {
	// получение сообщения
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

func client() {
	// соединиться с сервером
	c, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
		return
	}

	// послать сообщение
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}
func main() {
	fmt.Printf("Start")
	go server()
	client()
	time.Sleep(10 * time.Millisecond)
	// fmt.Scanln()
}
