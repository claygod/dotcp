package dotcp

// Do TCP
// Server test
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net"
	"testing"
)

var ip net.IP = net.IPv4(127, 0, 0, 1)
var port int = 9999

func TestCreateServer(t *testing.T) {
	s := Server(NetworkTsp)
	if s == nil {
		t.Error("Error create server")
	}
}

func TestMethodTrue(t *testing.T) {

	s := Server(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)
	s.Register("handleArticle", handleArticle, newArticle, schemeArticle)
	s.Start()

	msg := []byte(`{"method": "handleArticle", "query": {"id": 777, "text": "Blah-blah"}}`)
	c := Client(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)
	res, err := c.Send(msg)
	if err != nil {
		t.Error(err)
	}
	if len(res) != 3 || res[0] != 77 {
		t.Error(res)
		t.Error("The handler returned incorrect data")
	}
}

func TestMethodFalse(t *testing.T) {
	c := initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticleFALSE", "query": {"id": 777, "text": "Blah-blah"}}`)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestArgLong(t *testing.T) {
	c := initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 777, "text": "Blah-blah-blah-blah"}}`)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestArgShort(t *testing.T) {
	c := initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 777, "text": "!"}}`)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestArgMin(t *testing.T) {
	c := initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": -5, "text": "Blah-blah"}}`)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestArgMax(t *testing.T) {
	c := initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 4294967296, "text": "Blah-blah"}}`)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestArgMore(t *testing.T) {
	c := initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 7, "text": "Blah-blah", "tag": "news"}}`)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestArgRequired(t *testing.T) {
	c := initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 7}}`)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestPortDiscrepancy(t *testing.T) {
	initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 7, "text": "Blah-blah"}}`)

	c := Client(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9998)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestPortMin(t *testing.T) {
	initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 7, "text": "Blah-blah"}}`)

	c := Client(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(portsLimitMin - 1)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestPortMax(t *testing.T) {
	initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 7, "text": "Blah-blah"}}`)

	c := Client(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(portsLimitMax + 1)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestIp(t *testing.T) {
	initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 7, "text": "Blah-blah"}}`)

	c := Client(NetworkTsp).IP(net.IPv4(127, 0, 0, 2)).Port(9999)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestRegister(t *testing.T) {
	s := Server(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)

	if s.Register("handleArticle", handleArticle, newArticle, schemeArticle) != nil {
		t.Error("Registration error")
	}
}

func TestReRegister(t *testing.T) {
	s := Server(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)
	s.Register("handleArticle", handleArticle, newArticle, schemeArticle)

	if s.Register("handleArticle", handleArticle, newArticle, schemeArticle) == nil {
		t.Error("Re-registration is not possible")
	}
}

func initServerClient(ip net.IP, port int) *tspClient { // (*tcpServer,
	s := Server(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)
	s.Register("handleArticle", handleArticle, newArticle, schemeArticle)
	s.Start()
	c := Client(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)
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
