package main

// Do TCP
// Server
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>
import (
	//"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net"
	//"strconv"
	"bufio"
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
			t.handleServerConnection(con)
		}

	}()
}

/*
func server() {
	// слушать порт
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// принятие соединения
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// обработка соединения
		c.LocalAddr()
		go handleServerConnection(c)
	}
}
*/

func (t *tcpServer) handleServerConnection(c net.Conn) {
	fmt.Println("server")
	fmt.Println(c.LocalAddr())
	fmt.Println(c.RemoteAddr())
	//
	/*
		r := bufio.NewReader(c)
		w := bufio.NewWriter(c)
		scaner := bufio.NewScanner(r)
		scanned := scaner.Scan()
		if !scanned {
			log.Println("ERRROORRR!", c.RemoteAddr())
			return
		}
		fmt.Println(scaner.Bytes())
		fmt.Println(r)
		return
	*/
	//r := bufio.NewReader(c)
	//scaner := bufio.NewScanner(r)
	//scanned := scaner.Scan()
	//if !scanned {
	//	log.Println("ERRROORRR!", c.RemoteAddr())
	//	return
	//}
	//fmt.Println(r.ReadLine())
	//fmt.Println(scaner)
	// получение сообщения
	buf := make([]byte, 0)
	xx := make([]byte, 2)
	for {
		res, err := c.Read(xx)

		// fmt.Println(res, "  ++")
		if err != nil {
			break
		}
		buf = append(buf, xx[:res]...)
	}
	fmt.Println("===----------------", buf)

	//w.WriteString("PESEC")
	//w.Flush()

	//c.Write([]byte(`Message received.---`))
	//c.Close()
	//return

	//c.Close()
	//fmt.Println("===----------------", buf)
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
	//connRE, _ := net.Dial("tcp", c.RemoteAddr())
	// fmt.Println(c.Write([]byte{5, 5, 5}))
	//w := bufio.NewWriter(c)
	//w.WriteString("Golos")
	//w.Flush()
	c.Write([]byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2})

	fmt.Println(" ++++++++++++")
	//c.Write([]byte{5, 5, 5})
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

func client() {
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
	fmt.Println("client")
	fmt.Println(c.LocalAddr())
	fmt.Println(c.RemoteAddr())
	/*
		r := bufio.NewReader(c)
		w := bufio.NewWriter(c)

		w.Write(j)
		w.Flush()
		return
		scaner := bufio.NewScanner(r)
		scanned := scaner.Scan()

		if !scanned {
			log.Println("ERRROORRR! client", c.RemoteAddr())

		}
		fmt.Println(scaner.Text())
	*/
	c.Write(j)
	//buf := make([]byte, 2)
	fmt.Println(" &&&")

	//fmt.Println(c.Read(buf))
	c.Close()
	//fmt.Println("===----------------", buf)
}

func client3() {
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
	fmt.Println("client")
	fmt.Println(c.LocalAddr())
	fmt.Println(c.RemoteAddr())

	r := bufio.NewReader(c)
	w := bufio.NewWriter(c)
	//w.WriteString("qqqqqqqqq")
	//c.Write([]byte{7, 7, 7, 7})
	w.Write(j)
	w.Flush()
	/*
		buf := make([]byte, 0)
		xx := make([]byte, 2)
		for {
			res, err := c.Read(xx)

			// fmt.Println(res, "  ++")
			if err != nil {
				break
			}
			buf = append(buf, xx[:res]...)
		}
		fmt.Println("===----------------", buf)
	*/
	fmt.Println(" &&&")
	scaner := bufio.NewScanner(r)
	fmt.Println(scaner)
	//fmt.Println(r)
	//if scaner.Scan(r)

	//buf := make([]byte, 2)
	fmt.Println(" &&&")
	c.Close()

}
func client2() {
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:9999")
	for {
		// read in input from stdin
		//reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		//text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, "ttttttttt"+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}

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

	//go server()
	time.Sleep(10 * time.Millisecond)
	go client3()
	//go client()
	//go client2()

	time.Sleep(10 * time.Millisecond)
	// fmt.Scanln()
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
