package obd

import (
	"fmt"
	"strconv"
	"strings"
)

type Elm327 interface {
	Send(buffer []byte) ([]byte, error)
}

type Provider struct{
	elm327 Elm327
}

func NewObdProvider(elm327 Elm327) Provider {
	return Provider{elm327: elm327}
}

func (p *Provider) GetRpm() (int, error){
	resp, err := p.elm327.Send([]byte("010C\r"))
	if err != nil{
		return 0, err
	}

	if !strings.HasPrefix(string(resp), "410C") {
		return 0, fmt.Errorf("Invalid response : %v", string(resp))
	}

	a, err := strconv.ParseUint(strings.TrimSuffix(string(resp[4:9]),"\r"), 16, 32)
	if err != nil{
		panic(err.Error())
	}

	return int(a) / 4, nil
}
