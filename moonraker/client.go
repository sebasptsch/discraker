package moonraker

import (
	"context"
	"time"

	"github.com/sebasptsch/discraker/moonraker/structs"
)

func (s *Session) PrinterInfo() (structs.Status, error) {
	var reply structs.Status

	// slog.Println("Sending 'Arith.Multiply' request...")
	ctx, cancel := context.WithTimeout(*s.Context, 5*time.Second)
	defer cancel()

	err := s.RPCConnection.Call(ctx, "printer.info", nil, &reply)
	if err != nil {
		return structs.Status{}, err
	}

	return reply, err
}

func (s *Session) WebcamsList() (structs.WebcamList, error) {
	var reply structs.WebcamList

	// slog.Println("Sending 'Arith.Multiply' request...")
	ctx, cancel := context.WithTimeout(*s.Context, 5*time.Second)
	defer cancel()

	err := s.RPCConnection.Call(ctx, "server.webcams.list", nil, &reply)
	if err != nil {
		return structs.WebcamList{}, err
	}

	return reply, err
}

func (s *Session) Close() {
	if s.RPCConnection != nil {
		s.RPCConnection.Close()
	}

	s.ContextCancel()
}
