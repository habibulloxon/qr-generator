package main

import (
	"fmt"
	"github.com/rs/xid"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
)

func qrGenerator() {
	uniqueId := xid.New().String()

	qrCode, err := qrcode.Encode(uniqueId, qrcode.Medium, 256)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	fmt.Println(err)

	err = ioutil.WriteFile(uniqueId+".png", qrCode, 0644)
	if err != nil {
		fmt.Println("Error saving QR code to file:", err)
		return
	}

	fmt.Println("QR code generated successfully and saved")
}

func main() {
	var countQr int
	fmt.Println("Enter the number of QR:")
	fmt.Scanln(&countQr)

	for i := 0; i < countQr; i++ {
		qrGenerator()
	}
}
