package main

import (
	"log"
	"os"
	"time"

	"github.com/zwj186/qrcode"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	startAt := time.Now()
	fi, err := os.Open("qrcode11.png")
	if !check(err) {
		return
	}
	defer fi.Close()
	qrMatrix, err := qrcode.Decode(fi)
	if !check(err) {
		return
	}
	logger.Println(qrMatrix.Content)
	logger.Println(time.Now().Sub(startAt))
}

func check(err error) bool {
	if err != nil {
		logger.Println(err)
	}
	return err == nil
}
