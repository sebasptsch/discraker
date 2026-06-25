package moonraker

import (
	"context"
	"fmt"
	"log/slog"
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
	s := &Session{}

	moonrakerContext, cancel := context.WithCancel(context.Background())

	s.ContextCancel = cancel
	s.Context = &moonrakerContext

	var stream jsonrpc2.ObjectStream
	connectionUrl, err := url.Parse(connectionString)

	if err != nil {
		return s, err
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
		slog.Info(fmt.Sprintf("Connecting with URL: %s", u.String()))

		dialer := websocket.DefaultDialer
		wsConn, _, err := dialer.Dial(u.String(), nil)
		if err != nil {
			return s, err
		}

		stream = wsrpc.NewObjectStream(wsConn)
	case "unix":
		conn, err := net.Dial("unix", connectionUrl.Path)
		if err != nil {
			return s, err
		}
		stream = NewETXObjectStream(conn)
		// stream = NewETXObjectStream(conn)
	default:
		return s, err
	}

	rpcConn := jsonrpc2.NewConn(moonrakerContext, stream, handler)

	s.RPCConnection = rpcConn

	return s, nil
}
