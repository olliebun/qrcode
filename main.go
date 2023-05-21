// qrcode is a utility to generate a QR code image from a URL.
// Run qrcode -h to see usage.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	opts := parseFlags()
	png, err := qrcode.Encode(opts.url, qrcode.Low, int(opts.size))
	exitOnErr(err)
	_, err = os.Stdout.Write(png)
	exitOnErr(err)
}

type Opts struct {
	url  string
	size uint
}

func parseFlags() Opts {
	var o Opts
	flag.StringVar(&o.url, "url", "", "the url to encode")
	flag.UintVar(&o.size, "size", 1024, "width/height in pixels")
	flag.Parse()

	if o.url == "" {
		flag.Usage()
		os.Exit(1)
	}
	return o
}

func exitOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
