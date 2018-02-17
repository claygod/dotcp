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
	//"time"
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
newTcpServer - create new tcpServer.

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
*/

/*
Server - create a new tcp server.
*/
func Server(network networkType) *tcpServer { // TCPAddr
	ts := &tcpServer{
		network:    string(network),
		addr:       new(net.TCPAddr),
		procedures: make(map[string]*procedure),
	}
	return ts
}

/*
IP - set IP.
*/
func (t *tcpServer) IP(ip net.IP) *tcpServer {
	t.addr.IP = ip
	return t
}

/*
Port - set Port.
*/
func (t *tcpServer) Port(port int) *tcpServer {
	t.addr.Port = port
	return t
}

/*
Zone - set Zone.
*/
func (t *tcpServer) Zone(zone string) *tcpServer {
	t.addr.Zone = zone
	return t
}

/*
Register - register handler.
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

	//addr := &net.TCPAddr{
	//	IP:   net.IPv4(127, 0, 0, 1), // use net.ParseCIDR for string
	//	Port: 9999,
	//	Zone: "", // IPv6 scoped addressing zone
	//}
	//s := newTcpServer(NetworkTsp, addr)
	//if s == nil {
	//	log.Fatalf("Server not started")
	//	return
	//}
	s := Server(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)

	s.Register("RGB", dummy1, newRGB)
	s.Register("YCbCr", dummy2, newYCbCr)

	s.Start()

	fmt.Println(s.addr.String())

	//fmt.Println("Start")

	//client7()
	//fmt.Println("------------------------------------------------------------")
	//fmt.Println("------------------------------------------------------------")

	var msg = []byte(`[
		{"Method": "YCbCr", "Query": {"Y": 255, "Cb": 0, "Cr": -10}},
		{"Method": "RGB",   "Query": {"R": 98, "G": 218, "B": 255, "X":0}}
	]`)

	c := Client(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)
	c.Send(msg)

	//time.Sleep(1 * time.Second)
}

/*
procedure - stores the called function and the structure for the marshaling.
*/
type procedure struct {
	name      string
	method    func(interface{}) []byte
	getStruct func() interface{}
}
