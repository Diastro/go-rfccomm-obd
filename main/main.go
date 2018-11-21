package main

import (
	"fmt"
	"github.com/Diastro/go-rfccomm-obd/internal/rfcomm"
	"github.com/Diastro/go-rfccomm-obd/pkg/elm327"
)

func main() {
	//https://github.com/brendan-w/python-OBD/tree/master/obd
	sock := rfcomm.NewScoket([6]uint8{0xF1, 0xBE, 0x6D, 0x70, 0xF3, 0x5C}, 0)
	elm := elm327.NewElm327(&sock)
	resp, err := elm.Send([]byte("010C"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Print(resp)
}
