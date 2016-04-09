package natsproxy

import (
	"log"

	"encoding/json"

	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/nats"
)

const (
	// GET method constant
	GET = "GET"
	// POST method constant
	POST = "POST"
	// PUT method constant
	PUT = "PUT"
	// DELETE method constant
	DELETE = "DELETE"
)

// NatsHandler handles the
// tranforrmed HTTP request from
// NatsProxy. The context c wraps the
// request and response.
type NatsHandler func(c *Context)

// NatsHandlers is an
// array of NatsHandler functions.
// This type primary function is to group
// filters in NatsClient.
type NatsHandlers []NatsHandler

// Connector is the interface for
// generic pub/sub client.
type Connector interface {
	Subscribe(url string, handler NatsHandler)
	UnSubscribe(url string, handler NatsHandler)
}

// NatsClient serves as Connector
// to NATS messaging. Allows to subscribe
// for an specific url or url pattern.
type NatsClient struct {
	conn    *nats.Conn
	filters NatsHandlers
}

// NewNatsClient creates new NATS client
// from given connection. The connection must be
// connected or the function will
// return error ErrNatsClientNotConnected.
func NewNatsClient(conn *nats.Conn) (*NatsClient, error) {
	if err := testConnection(conn); err != nil {
		return nil, err
	}
	return &NatsClient{
		conn,
		make([]NatsHandler, 0),
	}, nil
}

// Use will add the middleware NatsHandler
// for a client.
func (nc *NatsClient) Use(middleware NatsHandler) {
	nc.filters = append(nc.filters, middleware)
}

// GET subscribes the client
// for an url with GET method.
func (nc *NatsClient) GET(url string, handler NatsHandler) {
	nc.Subscribe(GET, url, handler)
}

// POST subscribes the client
// for an url with POST method.
func (nc *NatsClient) POST(url string, handler NatsHandler) {
	nc.Subscribe(POST, url, handler)
}

// PUT subscribes the client
// for an url with PUT method.
func (nc *NatsClient) PUT(url string, handler NatsHandler) {
	nc.Subscribe(PUT, url, handler)
}

// DELETE subscribes the client
// for an url with DELETE method.
func (nc *NatsClient) DELETE(url string, handler NatsHandler) {
	nc.Subscribe(DELETE, url, handler)
}

// Subscribe is a generic subscribe function
// for any http method. It also
// wraps the processing of the context.
func (nc *NatsClient) Subscribe(method, url string, handler NatsHandler) {
	subscribeURL := SubscribeURLToNats(method, url)
	nc.conn.Subscribe(subscribeURL, func(m *nats.Msg) {
		request := &Request{}
		if err := request.UnmarshallFrom(m.Data); err != nil {
			log.Println(err)
			return
		}
		response := NewResponse()
		c := newContext(url, response, request)

		// Iterate through filters
		for _, filter := range nc.filters {
			filter(c)
			c.index++
		}

		// If request is aborted do
		// not proceed to handler.
		if !c.IsAborted() {
			handler(c)
		}
		bytes, err := proto.Marshal(c.Response)
		if err != nil {
			log.Println(err)
			return
		}
		nc.conn.Publish(m.Reply, bytes)
	})
}

func (nc *NatsClient) HandleWebsocket(webSocketID string, handler nats.MsgHandler) {
	nc.conn.Subscribe(ws_IN+webSocketID, handler)
}

// WriteWebsocketJSON writes struct
// serialized to JSON to registered
// websocketID NATS subject.
func (nc *NatsClient) WriteWebsocketJSON(websocketID string, msg interface{}) error {
	if data, err := json.Marshal(msg); err == nil {
		return nc.WriteWebsocket(websocketID, data)
	} else {
		return err
	}
}

func (nc *NatsClient) WriteWebsocket(websocketID string, data []byte) error {
	return nc.conn.Publish(ws_OUT+websocketID, data)
}
