package moonraker

import (
	"context"
	"log"
	"net"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	wsrpc "github.com/sourcegraph/jsonrpc2/websocket"
)

// Empty struct for rpc session
type Session struct {
	Context       *context.Context
	RPCConnection *jsonrpc2.Conn
	ContextCancel context.CancelFunc
}

func New(connectionString string, handler jsonrpc2.Handler) (*Session, error) {
	moonrakerContext, cancel := context.WithCancel(context.Background())

	var stream jsonrpc2.ObjectStream
	connectionUrl, err := url.Parse(connectionString)

	if err != nil {
		log.Panicf("URL format not valid %v", err)
	}

	switch scheme := connectionUrl.Scheme; scheme {

	case "http":
		fallthrough
	case "ws":
		fallthrough
	case "https":
		fallthrough
	case "wss":
		var u url.URL
		if scheme == "http" || scheme == "ws" {
			u = url.URL{Scheme: "ws", Host: connectionUrl.Host, Path: "/websocket"}
		} else {
			u = url.URL{Scheme: "wss", Host: connectionUrl.Host, Path: "/websocket"}
		}
		log.Printf("Connecting with URL: %s", u.String())

		dialer := websocket.DefaultDialer
		wsConn, _, err := dialer.Dial(u.String(), nil)
		if err != nil {
			log.Fatalf("WebSocket connection failed: %v", err)
		}

		stream = wsrpc.NewObjectStream(wsConn)
	case "unix":
		conn, err := net.Dial("unix", connectionUrl.Path)
		if err != nil {
			log.Fatalf("Failed to dial Unix socket: %v", err)
		}
		stream = jsonrpc2.NewPlainObjectStream(conn)
	default:
		log.Fatalf("%s scheme not supported", scheme)
	}

	rpcConn := jsonrpc2.NewConn(moonrakerContext, stream, handler)

	s := &Session{
		Context:       &moonrakerContext,
		RPCConnection: rpcConn,
		ContextCancel: cancel,
	}

	return s, nil
}
