package elm327

import "fmt"

type Elm327 struct {
	socket Socket
}

type Socket interface {
	Write([]byte) (int, error)
	Read() ([]byte, error)
}

func NewElm327(socket Socket) *Elm327 {
	elm := Elm327{socket: socket}
	elm.initialize()
	return &elm
}

func (e *Elm327) Send(buffer []byte) ([]byte, error) {
	_, err := e.socket.Write(buffer)
	if err != nil {
		return nil, err
	}

	return e.read()
}

func (e *Elm327) initialize() error {
	_, err := e.send([]byte("ATZ"))
	if err != nil {
		return err
	}
	_, err = e.read()
	if err != nil {
		return err
	}

	_, err = e.send([]byte("ATE0"))
	if err != nil {
		return err
	}
	_, err = e.read() // should be ok
	if err != nil {
		return err
	}

	_, err = e.send([]byte("ATH1"))
	if err != nil {
		return err
	}
	_, err = e.read() // should be ok
	if err != nil {
		return err
	}

	_, err = e.send([]byte("ATL0"))
	if err != nil {
		return err
	}
	_, err = e.read() // should be ok
	if err != nil {
		return err
	}

	return nil
}

func (e *Elm327) send(buffer []byte) (int, error) {
	return e.socket.Write(buffer)
}

func (e *Elm327) read() ([]byte, error) {
	buf := make([]byte, 1024)

	for {
		r, err := e.socket.Read()
		if err != nil {
			return nil, err
		}

		// stop on terminal >
		for i, n := range r {
			if 0x3C == n {
				buf = append(buf, r[:i]...)
				break
			}

			buf = append(buf, r...)
		}
	}

	fmt.Println(string(buf))
	return buf, nil
}