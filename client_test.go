package dotcp

// Do TCP
// Client test
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net"
	"testing"
)

func TestClientZone(t *testing.T) { //2001:DB0:0:123A::30
	z := "2001:DB0:0:123A::30"
	c := Client(NetworkTsp).Zone(z)

	if c.addr.Zone != z {
		t.Error("Server: address (Zone) error")
	}
}

func TestClientDial(t *testing.T) { //2001:DB0:0:123A::30
	c := Client(NetworkTsp).IP(ip).Port(65535)
	_, err := c.Send([]byte{})

	if err == nil {
		t.Error("Client dial - must return an error")
	}
}

func BenchmarkSequence(b *testing.B) {
	b.StopTimer()
	s := Server(NetworkTsp).IP(ip).Port(port + 1)
	s.Register("handleArticle", handleArticle, newArticle, schemeArticle)
	c := Client(NetworkTsp).IP(ip).Port(port + 1)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 7, "text": "Blah-blah"}}`)
	//u := 0
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Send(msg)
	}
}

func BenchmarkParallel(b *testing.B) {
	b.StopTimer()

	s := Server(NetworkTsp).IP(ip).Port(port + 2)
	s.Register("handleArticle", handleArticle, newArticle, schemeArticle)
	var clients [256]*tspClient
	for i := 0; i < 256; i++ {
		clients[uint8(i)] = Client(NetworkTsp).IP(ip).Port(port + 2)
	}
	msg := []byte(`{"method": "handleArticle", "query": {"id": 7, "text": "Blah-blah"}}`)

	var u uint8
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			clients[u].Send(msg)
			u++
		}
	})
}

func initServerClient(ip net.IP, port int) *tspClient { // (*tcpServer,
	s := Server(NetworkTsp).IP(ip).Port(port)
	s.Register("handleArticle", handleArticle, newArticle, schemeArticle)
	s.Start()
	c := Client(NetworkTsp).IP(ip).Port(port)
	return c
}

type Article struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

func newArticle() interface{} {
	return &Article{}
}

var schemeArticle string = `{
	"additionalProperties":false,
	"properties": {
		"id": {"type": "integer", "minimum": 0, "maximum": 4294967295},
		"text": {"type": "string", "minLength":2, "maxLength":10}
	},
	"required": ["id", "text"]
}`

func handleArticle(d interface{}) []byte {
	return []byte{77, 77, 77}
}
