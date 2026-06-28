package moonraker

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sebasptsch/discraker/moonraker/comms"
	"github.com/sourcegraph/jsonrpc2"
	wsrpc "github.com/sourcegraph/jsonrpc2/websocket"
)

// Empty struct for rpc session
type Session struct {
	// The RPC Background Context
	Context *context.Context
	// The RPC Connection itself
	RPCConnection *jsonrpc2.Conn
	// The cancel function to abort any in-progress RPC requests
	ContextCancel context.CancelFunc
}

// Create a new Moonraker client session
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
		stream = comms.NewETXObjectStream(conn)
		// stream = NewETXObjectStream(conn)
	default:
		return s, errors.ErrUnsupported
	}

	s.RPCConnection = jsonrpc2.NewConn(moonrakerContext, stream, handler)

	return s, nil
}

// The close function that cancells running requests and frees up resources
func (s *Session) Close() {
	if s.RPCConnection != nil {
		s.RPCConnection.Close()
	}

	s.ContextCancel()
}

// A generic function to abstract away moonraker rpc calling functionality
func rpc[T any](s *Session, method string, params any) (T, error) {
	var reply T

	// 5 second timeout
	ctx, cancel := context.WithTimeout(*s.Context, 5*time.Second)
	defer cancel()

	if err := s.RPCConnection.Call(ctx, method, params, &reply); err != nil {
		var zero T
		return zero, err
	}

	return reply, nil
}
