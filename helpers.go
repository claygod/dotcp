package main

// Do TCP
// Server add
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>
import (
	//"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func (t *tcpServer) handle(c net.Conn) {
	fmt.Println("H1")
	// получение сообщения
	buf := make([]byte, 0, 100)

	for {
		xx := make([]byte, 100)
		res, err := c.Read(xx)
		fmt.Println("H3", res, err)
		buf = append(buf, xx[:res]...)
		if err != nil || res == 0 || res < 100 {
			break
		}

	}

	fmt.Println("H4")
	fmt.Println("===----------------", buf)
	fmt.Println(" Сервер  получил:")
	fmt.Println(len(buf))

	fmt.Println("К маршаллингу:", buf)
	var colors []Color
	err := json.Unmarshal(buf, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}

	for _, item := range colors {
		var dst interface{}

		if p, ok := t.procedures[item.Method]; ok {
			dst = p.getStruct()
		} else {
			break // To Do
		}
		/*
			switch c.Method {
			case "RGB":
				dst = new(RGB)
			case "YCbCr":
				dst = new(YCbCr)
			}
		*/
		err := json.Unmarshal(item.Query, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}
		fmt.Println(item.Method, dst)
	}
	fmt.Println(" ++++++++++++")
	fmt.Println(" Сервер  пытается отправить клиенту:")
	fmt.Println(c.Write([]byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}))
	fmt.Println(" ++++++++++++")
	//c.Close()
}

func client7() {
	// соединиться с сервером
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err, " !!")
		return
	}

	// послать сообщение
	var j = []byte(`[
		{"Method": "YCbCr", "Query": {"Y": 255, "Cb": 0, "Cr": -10}},
		{"Method": "RGB",   "Query": {"R": 98, "G": 218, "B": 255, "X":0}}
	]`)

	fmt.Println("Отправляется массив - ")
	fmt.Println(j)

	fmt.Println(" Клиент пытается отправить серверу:")
	fmt.Println(c.Write(j))

	//time.Sleep(1000 * time.Millisecond)

	// получение сообщения
	buf := make([]byte, 0)
	fmt.Println(" Клиент  пытается получить:")

	for {
		xx := make([]byte, 100)
		res, err := c.Read(xx)
		fmt.Println("H3", res, err)
		buf = append(buf, xx[:res]...)

		if err != nil || res == 0 || res < 100 {
			fmt.Println(" Клиент  получил:")
			fmt.Println(res, err)
			break
		}

	}
	fmt.Println("===----------------", buf)

	c.Close()
	return
}

func handle1() []byte {
	return []byte{1, 1, 1}
}

func handle2() []byte {
	return []byte{2, 2, 2}
}

func dummy(interface{}) []byte {
	return []byte{7, 7, 7}
}

type Color struct {
	Method string
	Query  json.RawMessage // delay parsing until we know the color space
}
type RGB struct {
	R uint8
	G uint8
	B uint8
	X uint8
}

func newRGB() interface{} {
	//item := &RGB{}
	//i := item.(interface{})
	return &RGB{}
}

type YCbCr struct {
	Y  uint8
	Cb int8
	Cr int8
}

func newYCbCr() interface{} {
	//item := &RGB{}
	//i := item.(interface{})
	return &YCbCr{}
}
