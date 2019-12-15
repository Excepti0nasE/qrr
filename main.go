package main

import (
	"flag"
	"fmt"
	"github.com/tuotoo/qrcode"
	"os"
)

var (
	h bool
	v bool
	d string
)

const (
	appname  = "qrr"
	version  = "0.1.0"
	platform = "darwin/amd64"
)

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}

	if v {
		fmt.Printf("%v %v %v\n", appname, version, platform)
	}

	if d != "" {
		fi, err := os.Open(d)
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

func usage() {
	fmt.Fprintf(os.Stderr, `qrr version: qrr 0.1.0
Usage: qrr [-d decode] [-h help] [-v version]
	
Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.StringVar(&d, "d", "", "`decode` following qrcode")
	flag.BoolVar(&h, "h", false, "show qrr help")
	flag.BoolVar(&v, "v", false, "show version")
	flag.Usage = usage
}
