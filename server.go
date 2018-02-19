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
		schema:    gojsonschema.NewStringLoader(schemeRGB),
	}
	return nil
}

/*
Start - start the server.
*/
func (t *tcpServer) Start() error {
	lstnr, err := net.ListenTCP(t.network, t.addr)
	if err != nil {
		//log.Fatalln(err)
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

	time.Sleep(10 * time.Millisecond)
	return nil
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

	//var msg = []byte(`[
	//	{"Method": "YCbCr", "Query": {"Y": 255, "Cb": 0, "Cr": -10}},
	//	{"Method": "RGB",   "Query": {"R": 98, "G": 218, "B": 255, "X":0}}
	//]`)

	//inRGB := `{"R": 98, "G": 218, "B": 255, "X":0}`

	//msg := []byte(`{"Method": "YCbCr", "Query": {"Y": 255, "Cb": 0, "Cr": -10}}`)
	msg2 := []byte(`{"Method": "RGB",   "Query": {"R": 98, "G": 218, "B": 255, "k": 6}}`)
	//msg3 := []byte(`{"Method": "RGB",   "Query":` + inRGB + `}`)
	//time.Sleep(10 * time.Millisecond)

	c := Client(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)

	/*
		reply, err := c.Send(msg)
		fmt.Println("_APP_get1: ", reply, err)
		reply2, err2 := c.Send(msg2)
		fmt.Println("_APP_get2: ", reply2, err2)
	*/
	reply3, err3 := c.Send(msg2)
	fmt.Println("_APP_get3: ", reply3, err3)

	//
}

/*
procedure - stores the called function and the structure for the marshaling.
*/
type procedure struct {
	name      string
	method    func(interface{}) []byte
	getStruct func() interface{}
	schema    gojsonschema.JSONLoader
}

var schemeRGB string = `{
	"additionalProperties":false,
	"properties": {
		"R": {"type": "integer"},
		"G": {"type": "integer"},
		"B": {"type": "integer"},
		"X": {"type": "integer"}
	},
	"required": ["R", "G", "B", "X"]
}`
