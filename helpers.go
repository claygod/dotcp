package main

// Do TCP
// Server add
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>
import (
	//"encoding/gob"
	"encoding/json"
	"fmt"
	// "log"
	"net"
	// "time"
	"github.com/xeipuuv/gojsonschema"
)

func (t *tcpServer) handle(c net.Conn) {
	reply := make([]byte, 1)
	buf := make([]byte, 0)

	// load msg
	for {
		xx := make([]byte, bufSize)
		res, err := c.Read(xx)
		buf = append(buf, xx[:res]...)
		if err != nil || res < bufSize { // res == 0 ||
			break
		}
	}

	var b Box
	err := json.Unmarshal(buf, &b)
	if err != nil {
		reply[0] = ErrUnmarshal
		reply = append(reply, []byte(err.Error())...)
		goto sendReply
	}
	//fmt.Println("======")
	if p, ok := t.procedures[b.Method]; ok {
		doc := gojsonschema.NewBytesLoader(b.Query)
		//fmt.Println(b.Query)
		//fmt.Println(doc)
		result, err := gojsonschema.Validate(p.schema, doc)

		if err != nil {
			reply[0] = Err
			reply = append(reply, []byte(err.Error())...)
			goto sendReply
		}

		if !result.Valid() {
			//fmt.Printf("The document is not valid. see errors :\n")
			for _, desc := range result.Errors() {
				fmt.Printf("- %s\n", desc)
			}
		}

		//fmt.Println(resul.Valid())
		//fmt.Println(result)

		if err != nil {
			//for _, desc := range resul.Errors() {
			//	fmt.Printf("- %s\n", desc)
			//}
			reply[0] = Err
			reply = append(reply, []byte(err.Error())...)
			goto sendReply
		}

		dst := p.getStruct()

		if err := json.Unmarshal(b.Query, dst); err != nil {
			reply[0] = ErrUnmarshal
			reply = append(reply, []byte(err.Error())...)
		} else {
			reply[0] = Ok
			reply = append(reply, p.method(dst)...)
		}
		goto sendReply
	} else {
		reply[0] = ErrUnregistered
		reply = append(reply, []byte("No procedure")...)
		goto sendReply
	}
sendReply:
	fmt.Println(c.Write(reply))
}

func dummy1(d interface{}) []byte {
	fmt.Println(" ++++ сработал dummy 1 ++++", d)
	return []byte{71, 71, 71}
}

func dummy2(d interface{}) []byte {
	fmt.Println(" ++++ сработал dummy 2 ++++", d)
	return []byte{72, 72, 72}
}

type Box struct {
	Method string
	Query  json.RawMessage
}
type Color2 struct {
	Method string
	Query  []byte
}

type RGB struct {
	R uint8 `jsonschema:"required"`
	G uint8 `json:"G" jsonschema:"required"`
	B uint8 `jsonschema:"required"`
	X uint8 `jsonschema:"required"`
}

func newRGB() interface{} {
	return &RGB{}
}

type YCbCr struct {
	Y  uint8 `jsonschema:"required"`
	Cb int8  `jsonschema:"required"`
	Cr int8  `jsonschema:"required"`
}

func newYCbCr() interface{} {
	return &YCbCr{}
}
