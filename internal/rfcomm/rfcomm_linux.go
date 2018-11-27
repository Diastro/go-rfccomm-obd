package rfcomm

import (
	"golang.org/x/sys/unix"
	"syscall"
)

type Socket struct {
	fd int
}

func NewSocket(addr [6]uint8, port uint8) Socket {
	fd, err := unix.Socket(syscall.AF_BLUETOOTH, syscall.SOCK_STREAM, unix.BTPROTO_RFCOMM)
	if err != nil {
		panic(err.Error())
	}

	//https://github.com/golang/sys/blob/master/unix/syscall_linux.go
	socket := &unix.SockaddrRFCOMM{
		//Addr:  [6]uint8{0x72,0xA2,0x9F,0x3E,0x04,0x00}
		//Addr:  [6]uint8{0xF1,0xBE,0x6D,0x70,0xF3,0x5C}
		Addr:    addr,
		Channel: uint8(port),
	}
	err = unix.Connect(fd, socket)
	if err != nil {
		panic(err.Error())
	}

	return Socket{fd: fd}
}

func (s *Socket) Write(buffer []byte) (int, error) {
	return unix.Write(s.fd, buffer)
}

func (s *Socket) Read() ([]byte, error) {
	buf := make([]byte, 1024)
	l, err := unix.Read(s.fd, buf)
	if err != nil {
		return nil, err
	}
	return buf[0:l], nil
}
