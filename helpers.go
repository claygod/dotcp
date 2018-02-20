package dotcp

// Do TCP
// Helpers
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"encoding/json"

	"github.com/xeipuuv/gojsonschema"
)

/*
Box - container for arguments of the called method.
*/
type Box struct {
	Method string          `json:"method"`
	Query  json.RawMessage `json:"query"`
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
