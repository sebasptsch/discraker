package moonraker

import (
	"context"
	"log"
	"time"

	"github.com/sebasptsch/discraker/moonraker/structs"
)

func (s *Session) PrinterInfo() (structs.Status, error) {
	var reply structs.Status

	// log.Println("Sending 'Arith.Multiply' request...")
	ctx, cancel := context.WithTimeout(*s.Context, 5*time.Second)
	defer cancel()

	err := s.RPCConnection.Call(ctx, "printer.info", nil, &reply)
	if err != nil {
		log.Fatalf("RPC Call error: %v", err)
	}

	return reply, err
}

func (s *Session) WebcamsList() (structs.WebcamList, error) {
	var reply structs.WebcamList

	// log.Println("Sending 'Arith.Multiply' request...")
	ctx, cancel := context.WithTimeout(*s.Context, 5*time.Second)
	defer cancel()

	err := s.RPCConnection.Call(ctx, "server.webcams.list", nil, &reply)
	if err != nil {
		log.Fatalf("RPC Call error: %v", err)
	}

	return reply, err
}

func (s *Session) Close() {
	if s.RPCConnection != nil {
		s.RPCConnection.Close()
	}

	s.ContextCancel()
}
