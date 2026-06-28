package moonraker

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
	// HTTP Client for downloading and uploading
	HTTPClient *http.Client
	// Connection URL
	ConnectionURL *url.URL
	apiKey        *string
}

func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			return filepath.Join(home, path[2:])
		}
	}
	return path
}

// Create a new Moonraker client session
func New(connectionString string, apiKey *string, handler jsonrpc2.Handler) (*Session, error) {
	s := &Session{
		HTTPClient: &http.Client{},
		apiKey:     apiKey,
	}

	moonrakerContext, cancel := context.WithCancel(context.Background())

	s.ContextCancel = cancel
	s.Context = &moonrakerContext

	var stream jsonrpc2.ObjectStream
	connectionUrl, err := url.Parse(connectionString)

	if err != nil {
		return s, err
	}

	s.ConnectionURL = connectionUrl

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

		connectionHeaders := &http.Header{}

		if apiKey != nil {
			connectionHeaders.Add("X-API-Key", *apiKey)
		}

		dialer := websocket.DefaultDialer
		wsConn, _, err := dialer.Dial(u.String(), *connectionHeaders)
		if err != nil {
			return s, err
		}

		stream = wsrpc.NewObjectStream(wsConn)
	case "unix":
		conn, err := net.Dial("unix", expandPath(connectionUrl.Path))
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

func newRequest(s *Session, method string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequestWithContext(*s.Context, method, url, body)

	if err != nil {
		return nil, err
	}

	request.Header.Add("X-API-Key", *s.apiKey)

	return request, nil
}
