package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
)

func main() {
	var link string

	fmt.Println("Enter link to generate QR: ")
	fmt.Scanln(&link)

	qrCode, err := qrcode.Encode(link, qrcode.Medium, 256)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	// Save the QR code to a file
	err = ioutil.WriteFile("qrcode.png", qrCode, 0644)
	if err != nil {
		fmt.Println("Error saving QR code to file:", err)
		return
	}

	fmt.Println("QR code generated successfully and saved to 'qrcode.png'!")
}
