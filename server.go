package main

// Do TCP
// Server
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>
import (
	//"encoding/json"
	"fmt"
	"log"
	"net"
	//"strconv"
	//"bufio"
	"errors"
	//"os"
	"time"
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
total - current account balance.
*/
func (t *tcpServer) Register(name string, method func(interface{}) []byte, getStruct func() interface{}) error {
	if _, ok := t.procedures[name]; ok {
		return errors.New("This procedure has already been registered")
	}

	t.procedures[name] = &procedure{
		name:      name,
		method:    method,
		getStruct: getStruct,
	}
	return nil
}

/*
procedure - stores the called function and the structure for the marshaling.
*/
type procedure struct {
	name      string
	method    func(interface{}) []byte
	getStruct func() interface{}
}

/*
newTcpServer - create new tcpServer.
*/
func newTcpServer(network networkType, addr *net.TCPAddr) *tcpServer { // TCPAddr
	//if _, err := strconv.Atoi() {
	//	return nil
	//}
	ts := &tcpServer{
		network:    string(network),
		addr:       addr,
		procedures: make(map[string]*procedure),
	}
	return ts
}

/*
Start - start the server.
*/
func (t *tcpServer) Start() {
	go func() {
		lstnr, err := net.ListenTCP(t.network, t.addr)
		if err != nil {
			log.Println(err)
			return
		}
		defer lstnr.Close()

		for {
			con, err := lstnr.AcceptTCP()
			if err != nil {
				log.Println(err)
				continue
			}
			t.handle(con)
		}

	}()
}

func main() {
	addr := &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1), // use net.ParseCIDR for string
		Port: 9999,
		Zone: "", // IPv6 scoped addressing zone
	}
	s := newTcpServer(NetworkTsp, addr)
	if s == nil {
		log.Fatalf("Server not started")
		return
	}
	s.Register("RGB", dummy, newRGB)
	s.Register("YCbCr", dummy, newYCbCr)
	s.Start()

	fmt.Println("Start")
	time.Sleep(10 * time.Millisecond)
	client7()

	//time.Sleep(1 * time.Second)
}
