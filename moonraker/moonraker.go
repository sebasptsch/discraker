package moonraker

import (
	"context"
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
	context *context.Context
	// The RPC Connection itself
	rpc_connection *jsonrpc2.Conn
	// The cancel function to abort any in-progress RPC requests
	context_cancel context.CancelFunc
	// Connection URL
	connectionParameters *ConnectionParameters
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

type ConnectionParameters struct {
	// Used for HTTP-only operations such as downloading and uploading files
	HttpURL *string
	// Connection URL is used for all RPC operations
	SocketURL *string
	// API Key is used for both HTTP and RPC operations
	APIKey *string
	// Token
	AccessToken *string
}

// Create a new Moonraker client session
func New(params *ConnectionParameters, handler jsonrpc2.Handler) (*Session, error) {
	s := &Session{
		connectionParameters: params,
	}

	moonrakerContext, cancel := context.WithCancel(context.Background())

	s.context_cancel = cancel
	s.context = &moonrakerContext

	var stream jsonrpc2.ObjectStream

	u, err := url.Parse(*s.connectionParameters.SocketURL)

	if err != nil {
		return nil, err
	}

	switch scheme := u.Scheme; scheme {
	case "ws":
		fallthrough
	case "wss":
		slog.Info(fmt.Sprintf("Connecting with URL: %s", u.String()))

		connectionHeaders := &http.Header{}

		if s.connectionParameters.APIKey != nil {
			connectionHeaders.Add("X-API-Key", *params.APIKey)
		}

		if s.connectionParameters.AccessToken != nil {
			connectionHeaders.Add("Authorization", fmt.Sprintf("Bearer %s", *s.connectionParameters.AccessToken))
		}

		dialer := websocket.DefaultDialer
		wsConn, _, err := dialer.Dial(u.String(), *connectionHeaders)
		if err != nil {
			return s, err
		}

		stream = wsrpc.NewObjectStream(wsConn)
	case "unix":
		conn, err := net.Dial("unix", expandPath(u.Path))
		if err != nil {
			return s, err
		}
		stream = comms.NewETXObjectStream(conn)
		// stream = NewETXObjectStream(conn)
	default:
		return s, fmt.Errorf("unsupported scheme used %s", scheme)
	}

	s.rpc_connection = jsonrpc2.NewConn(moonrakerContext, stream, handler)

	return s, nil
}

// The close function that cancells running requests and frees up resources
func (s *Session) Close() {
	if s.rpc_connection != nil {
		s.rpc_connection.Close()
	}

	s.context_cancel()
}

// A generic function to abstract away moonraker rpc calling functionality
func rpc[T any](s *Session, method string, params any) (T, error) {
	var reply T

	// 5 second timeout
	ctx, cancel := context.WithTimeout(*s.context, 5*time.Second)
	defer cancel()

	if err := s.rpc_connection.Call(ctx, method, params, &reply); err != nil {
		var zero T
		return zero, err
	}

	return reply, nil
}

func newRequest(s *Session, method string, path string, body io.Reader) (*http.Request, error) {
	fullUrl, err := url.JoinPath(*s.connectionParameters.HttpURL, path)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(*s.context, method, fullUrl, body)

	if err != nil {
		return nil, err
	}

	if s.connectionParameters.APIKey != nil {
		request.Header.Add("X-API-Key", *s.connectionParameters.APIKey)
	}

	if s.connectionParameters.AccessToken != nil {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *s.connectionParameters.AccessToken))
	}

	return request, nil
}
