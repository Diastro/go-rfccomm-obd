package main

import (
	"fmt"
	"github.com/Diastro/go-rfccomm-obd/internal/rfcomm"
	"github.com/Diastro/go-rfccomm-obd/pkg/elm327"
	"github.com/Diastro/go-rfccomm-obd/pkg/obd"
)

func main() {
	sock := rfcomm.NewSocket([6]uint8{0xF1, 0xBE, 0x6D, 0x70, 0xF3, 0x5C}, 2)
	//sock := rfcomm.NewSocket([6]uint8{0x72,0xA2,0x9F,0x3E,0x04,0x00}, 1)
	elm := elm327.NewElm327(&sock)


	odb := obd.NewObdProvider(elm)
	resp, err := odb.GetRpm()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("RPM : %v\n", resp)
}
