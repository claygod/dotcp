# DoTCP


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

### Example

```go
	s := Server(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)
	s.Register("handleArticle", handleArticle, newArticle, schemeArticle)
	s.Start()

	msg := []byte(`{"method": "handleArticle", "query": {"id": 777, "text": "Blah-blah"}}`)
	c := Client(NetworkTsp).IP(net.IPv4(127, 0, 0, 1)).Port(9999)
	res, err := c.Send(msg)
	if err != nil {
		// error handling
	}
	if len(res) != 3 || res[0] != 77 {
		t.Error(res)
		t.Error("The handler returned incorrect data")
	}
```	

### Struct (Example)

Blah-blah

```go
type Article struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}
```	

### New struct (Example)

Blah-blah

```go
func newArticle() interface{} {
	return &Article{}
}
```

### Scheme (Example)

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

### Handle (Example)

Blah-blah

```go
func handleArticle(d interface{}) []byte {
	return []byte{77, 77, 77}
}
```



## API

- New
- Load ("path")
- Start ()
- AddUnit(ID)
- Begin().Debit(ID, key, amount).End()
- Begin().Credit(ID, key, amount).End()
- TotalUnit(ID)
- TotalAccount(ID, key)
- DelUnit(ID)
- Stop ()
- Save ("path")

## F.A.Q.

Blah-blah
- Blah-blah

Blah-blah?
- Blah-blah

## ToDo

- [x] Blah-blah
- [ ] More tests

## Bench

i7-6700T:

- BenchmarkTotalUnitSequence-8        	 3000000	       419 ns/op
- BenchmarkTotalUnitParallel-8        	10000000	       185 ns/op
- BenchmarkCreditSequence-8           	 5000000	       311 ns/op
- BenchmarkCreditParallel-8           	10000000	       175 ns/op
- BenchmarkDebitSequence-8            	 5000000	       314 ns/op
- BenchmarkDebitParallel-8            	10000000	       178 ns/op
- BenchmarkTransferSequence-8         	 3000000	       417 ns/op
- BenchmarkTransferParallel-8         	 5000000	       277 ns/op
- BenchmarkBuySequence-8              	 2000000	       644 ns/op
- BenchmarkBuyParallel-8              	 5000000	       354 ns/op
