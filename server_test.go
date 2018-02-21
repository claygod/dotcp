package dotcp

// Do TCP
// Server test
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net"
	"testing"
)

var ip net.IP = net.IPv4(127, 0, 0, 1)
var port int = 49001

func TestCreateServer(t *testing.T) {
	s := Server(NetworkTsp)
	if s == nil {
		t.Error("Error create server")
	}
}

func TestMethodTrue1(t *testing.T) {

	s := Server(NetworkTsp).IP(ip).Port(port)
	s.Register("handleArticle", handleArticle, newArticle, schemeArticle)
	s.Start()

	msg := []byte(`{"method": "handleArticle", "query": {"id": 777, "text": "Blah-blah"}}`)
	c := Client(NetworkTsp).IP(ip).Port(port)
	res, err := c.Send(msg)
	if err != nil {
		t.Error(err)
	}
	if len(res) != 3 || res[0] != 77 {
		t.Error(res)
		t.Error("The handler returned incorrect data")
	}
}

func TestMethodTrue2(t *testing.T) {

	c := initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 777, "text": "Blah-blah"}}`)
	res, err := c.Send(msg)
	if err != nil {
		t.Error(err)
	}
	if len(res) != 3 || res[0] != 77 {
		t.Error(res)
		t.Error("The handler returned incorrect data")
	}
}

func TestFalseNameField(t *testing.T) {

	c := initServerClient(ip, port)
	msg := []byte(`{"methodFALSE": "handleArticle", "query": {"id": 777, "text": "Blah-blah"}}`)
	_, err := c.Send(msg)
	if err == nil {
		t.Error("An error is expected due to an invalid one field name")
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

	c := Client(NetworkTsp).IP(ip).Port(port + 1)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestPortMin(t *testing.T) {
	initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 7, "text": "Blah-blah"}}`)

	c := Client(NetworkTsp).IP(ip).Port(portsLimitMin - 1)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestPortMax(t *testing.T) {
	initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 7, "text": "Blah-blah"}}`)

	c := Client(NetworkTsp).IP(ip).Port(portsLimitMax + 1)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestIp(t *testing.T) {
	initServerClient(ip, port)
	msg := []byte(`{"method": "handleArticle", "query": {"id": 7, "text": "Blah-blah"}}`)

	c := Client(NetworkTsp).IP(net.IPv4(127, 0, 0, 2)).Port(port)

	if _, err := c.Send(msg); err == nil {
		t.Error(err)
	}
}

func TestRegister(t *testing.T) {
	s := Server(NetworkTsp).IP(ip).Port(port)

	if s.Register("handleArticle", handleArticle, newArticle, schemeArticle) != nil {
		t.Error("Registration error")
	}
}

func TestReRegister(t *testing.T) {
	s := Server(NetworkTsp).IP(ip).Port(port)
	s.Register("handleArticle", handleArticle, newArticle, schemeArticle)

	if s.Register("handleArticle", handleArticle, newArticle, schemeArticle) == nil {
		t.Error("Re-registration is not possible")
	}
}

func TestServerZone(t *testing.T) { //2001:DB0:0:123A::30
	z := "2001:DB0:0:123A::30"
	s := Server(NetworkTsp).Zone(z)

	if s.addr.Zone != z {
		t.Error("Server: address (Zone) error")
	}
}

func TestServerAddress(t *testing.T) {
	Server(NetworkTsp).IP(ip).Port(port).Start()

	if Server(NetworkTsp).IP(ip).Port(port).Start() == nil {
		t.Error("Server: start - must return an error")
	}
}
