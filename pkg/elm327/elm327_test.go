package elm327

import (
	"fmt"
	"testing"
)

type socketMock struct {
	write_f func([]byte) (int, error)
	read_r  func() ([]byte, error)
}

func (s *socketMock) Write(b []byte) (int, error) {
	return s.write_f(b)
}

func (s *socketMock) Read() ([]byte, error) {
	return s.read_r()
}

func TestElm327_Send(t *testing.T) {
	socket := socketMock{}

	socket.read_r = func() ([]byte, error) {
		return []byte("OK\r\n>\r\n"), nil
	}

	socket.write_f = func(bytes []byte) (int, error) {
		return 0, nil
	}

	elm := NewElm327(&socket)

	r, err := elm.Send([]byte("0104"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(r))
}
