package rfcomm

type RfcommSocket struct {
	fd int
}

func NewSocket(addr [6]uint8, port uint8) RfcommSocket {

	return RfcommSocket{fd: 0}
}

func (s *RfcommSocket) Write(buffer []byte) (int, error) {
	return 0, nil
}

func (s *RfcommSocket) Read() ([]byte, error) {
	return nil, nil
}
