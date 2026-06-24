package moonraker

import (
	"bufio"
	"encoding/json"
	"io"
	"sync"
)

const etx byte = 0x03

type ETXObjectStream struct {
	conn io.ReadWriteCloser
	r    *bufio.Reader
	w    *bufio.Writer
	mu   sync.Mutex
}

func NewETXObjectStream(conn io.ReadWriteCloser) *ETXObjectStream {
	return &ETXObjectStream{
		conn: conn,
		r:    bufio.NewReader(conn),
		w:    bufio.NewWriter(conn),
	}
}

func (s *ETXObjectStream) WriteObject(obj interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	data = append(data, etx)

	if _, err := s.w.Write(data); err != nil {
		return err
	}

	return s.w.Flush()
}

func (s *ETXObjectStream) ReadObject(v interface{}) error {
	data, err := s.r.ReadBytes(etx)
	if err != nil {
		return err
	}

	data = data[:len(data)-1] // strip ETX

	return json.Unmarshal(data, v)
}

func (s *ETXObjectStream) Close() error {
	return s.conn.Close()
}
