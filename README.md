# DoTCP

[![API documentation](https://godoc.org/github.com/claygod/dotcp?status.svg)](https://godoc.org/github.com/claygod/dotcp)
[![Go Report Card](https://goreportcard.com/badge/github.com/claygod/dotcp)](https://goreportcard.com/report/github.com/claygod/dotcp)

TCP server with RPC mode for JSON. Coverage 80.0%

TCH server for the exchange of JSON dates. Before starting the server, the procedures are registered, which will then be called in a RPC style. The server validates the received data.

The library has two main entities: a server and an client.

### Server

- works in RPC mode
- valid incoming JSON

### Client

- sends requests to the server

## Usage

Blah-blah Blah-blah

## Example

```go
import (
	"fmt"
	"net"

	"github.com/claygod/dotcp"
)

func main() {
	fmt.Println("Begin app")

	s := dotcp.Server(dotcp.NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)
	s.Register("handleArticle", handleArticle, newArticle, schemeArticle)
	s.Start()

	msg := []byte(`{"method": "handleArticle", "query": {"id": 777, "text": "Hello world!"}}`)
	c := dotcp.Client(dotcp.NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)

	res, err := c.Send(msg)

	if err != nil {
		// error handling
		fmt.Println(err)
	} else {
		fmt.Println("Server response:", string(res))
	}

	fmt.Println("End app")
}
```	

### Struct (for example)

Blah-blah

```go
type Article struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}
```	

### New struct (for example)

Blah-blah

```go
func newArticle() interface{} {
	return &Article{}
}
```

### Scheme (for example)

Blah-blah

```go
var schemeArticle string = `{
	"additionalProperties":false,
	"properties": {
		"id": {"type": "integer", "minimum": 0, "maximum": 4294967295},
		"text": {"type": "string", "minLength":2, "maxLength":10}
	},
	"required": ["id", "text"]
}`
```

### Handle (for example)

Blah-blah

```go
func handleArticle(d interface{}) []byte {
	return []byte{77, 77, 77}
}
```



## API

- Server()
- Server.IP()
- Server.Port()
- Server.Zone()
- Server.Register()
- Client()
- Client.IP()
- Client.Port()
- Client.Zone()
- Client.Send()

## F.A.Q.

Blah-blah
- Blah-blah

Blah-blah?
- Blah-blah

## ToDo

- [x] API
- [ ] More tests

## Bench

i3-5005U:

- BenchmarkSequence-2   	   50000	     32839 ns/op
- BenchmarkParallel-2   	  100000	     15691 ns/op
