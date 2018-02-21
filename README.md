# DoTCP

[![API documentation](https://godoc.org/github.com/claygod/dotcp?status.svg)](https://godoc.org/github.com/claygod/dotcp)
[![Go Report Card](https://goreportcard.com/badge/github.com/claygod/dotcp)](https://goreportcard.com/report/github.com/claygod/dotcp)

TCP server with RPC mode for JSON. Coverage 83.5%

TCH server for the exchange of JSON dates. Before starting the server, the procedures are registered, which will then be called in a RPC style. The server validates the received data.

The library has two main entities: a server and an client.

### Server

- works in RPC mode
- valid incoming JSON

### Client

- sends requests to the server

## Usage

When using the server, the correct port number is important.

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


```go
type Article struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}
```	

### New struct (for example)


```go
func newArticle() interface{} {
	return &Article{}
}
```

### Scheme (for example)


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

Why is your server slower than the standard TCP server?
- For the RPC-realization a JSON is used. The operation of unmarshalling is slow. In addition, the JSON should be validated. This is necessary for safety.

Which ports are best to use?
- This should be asked from the administrator)) Ports that are used dynamically: 49152 - 65535. Read more on [WiKi](https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers)

## ToDo

- [x] API
- [x] add example
- [x] F.A.Q.
- [ ] increase coverage
- [ ] more tests

## Bench

i3-5005U (2.0 GHz):

- BenchmarkSequence-2   	   50000	     26403 ns/op
- BenchmarkParallel-2   	  100000	     15819 ns/op
