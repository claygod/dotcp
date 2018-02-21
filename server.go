/*
TCP server with RPC mode for JSON.
TCH server for the exchange of JSON dates. Before starting the server,
the procedures are registered, which will then be called in a RPC style.
The server validates the received data.
*/

package dotcp

// Do TCP
// Server
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/xeipuuv/gojsonschema"
)

/*
tcpServer - TCP server.
*/
type tcpServer struct {
	network    string
	addr       *net.TCPAddr
	procedures map[string]*procedure
}

/*
Server - create a new tcp server.
*/
func Server(network networkType) *tcpServer {
	ts := &tcpServer{
		network:    string(network),
		addr:       new(net.TCPAddr),
		procedures: make(map[string]*procedure),
	}
	return ts
}

/*
IP - set IP.
An IP is a single IP address, a slice of bytes.
Functions in this package accept either 4-byte
(IPv4) or 16-byte (IPv6) slices as input.
*/
func (t *tcpServer) IP(ip net.IP) *tcpServer {
	t.addr.IP = ip
	return t
}

/*
Port - set Port.
0 - 65535
*/
func (t *tcpServer) Port(port int) *tcpServer {
	t.addr.Port = port
	return t
}

/*
Zone - set Zone.
IPv6 scoped addressing zone
*/
func (t *tcpServer) Zone(zone string) *tcpServer {
	t.addr.Zone = zone
	return t
}

/*
Register - register handler.
*/
func (t *tcpServer) Register(name string, method func(interface{}) []byte, getStruct func() interface{}, scheme string) error {
	if _, ok := t.procedures[name]; ok {
		return errors.New("This procedure has already been registered")
	}
	t.procedures[name] = &procedure{
		name:      name,
		method:    method,
		getStruct: getStruct,
		schema:    gojsonschema.NewStringLoader(scheme),
	}
	return nil
}

/*
Start - start the server.
*/
func (t *tcpServer) Start() error {
	if t.addr.Port > portsLimitMax || t.addr.Port < portsLimitMin {
		return errors.New("Uncorrect port")
	}
	lstnr, err := net.ListenTCP(t.network, t.addr)
	if err != nil {
		return err
	}

	go func(lstnr *net.TCPListener) {

		defer lstnr.Close()

		for {
			con, err := lstnr.AcceptTCP()
			if err != nil {
				log.Println(err)
				continue
			}
			t.handle(con)
		}

	}(lstnr)

	//time.Sleep(startPause * time.Millisecond)
	return nil
}

/*
handle - processing request.
*/
func (t *tcpServer) handle(c net.Conn) {
	reply := make([]byte, 1)
	buf := make([]byte, 0)

	// load msg
	for {
		xx := make([]byte, bufSize)
		res, err := c.Read(xx)
		buf = append(buf, xx[:res]...)
		if err != nil || res < bufSize {
			break
		}
	}

	var b Box
	err := json.Unmarshal(buf, &b)
	if err != nil {
		reply[0] = reError
		reply = append(reply, []byte(err.Error())...)
		goto sendReply
	}
	if p, ok := t.procedures[b.Method]; ok {
		doc := gojsonschema.NewBytesLoader(b.Query)
		result, err := gojsonschema.Validate(p.schema, doc)

		if err != nil {
			reply[0] = reError
			reply = append(reply, []byte(err.Error())...)
			goto sendReply
		}
		if !result.Valid() {
			var errStr string
			for _, desc := range result.Errors() {
				errStr = fmt.Sprintf("%s %v; ", errStr, desc)
			}
			reply[0] = reError
			reply = append(reply, []byte(errStr)...)
			goto sendReply
		}

		if err != nil {
			reply[0] = reError
			reply = append(reply, []byte(err.Error())...)
			goto sendReply
		}

		dst := p.getStruct()

		if err := json.Unmarshal(b.Query, dst); err != nil {
			reply[0] = reError
			reply = append(reply, []byte(err.Error())...)
		} else {
			reply[0] = reOk
			reply = append(reply, p.method(dst)...)
		}
		goto sendReply
	} else {
		reply[0] = reError
		reply = append(reply, []byte("No procedure")...)
		goto sendReply
	}
sendReply:
	c.Write(reply)
}
