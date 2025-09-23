package main

import (
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

const (
	targetURL  = "https://forms.gle/XMGnVaB3mr7vxFjy6"
	inputFile  = "gophers.jpeg"
	outputFile = "gophers_with_qr.png"
)

func main() {

	qrc, err := qrcode.NewWith(targetURL,
		qrcode.WithEncodingMode(qrcode.EncModeByte),
		qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionHighest),
	)
	if err != nil {
		panic(err)
	}

	w, err := standard.New(outputFile,
		standard.WithHalftone(inputFile),
		standard.WithQRWidth(30),
	)
	if err != nil {
		panic(err)
	}

	if err = qrc.Save(w); err != nil {
		panic(err)
	}
}
