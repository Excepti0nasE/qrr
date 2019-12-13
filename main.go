package main

import (
	"fmt"
	"github.com/tuotoo/qrcode"
	"os"
)

const appname = "qrr"
const version = "0.0.1"
const platform = "darwin/amd64"

func main() {
	s := os.Args[1]
	switch s {
	case "--version":
		fmt.Printf("%v verion %v %v\n", appname, version, platform)
	case "--help":
		fmt.Println("Just show some help infos")
	default:
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
}
