package main

import (
	"fmt"
	"os"
	"github.com/tuotoo/qrcode"
)

func main() {
	s := os.Args[1]
	fi, err := os.Open(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(qrmatrix.Content)
}
