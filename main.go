// qrcode is a utility to generate a QR code image from a URL.
// Run qrcode -h to see usage.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	log.SetFlags(0)
	f := parseFlags()
	checkURL(f.url)
	png, err := qrcode.Encode(f.url, qrcode.RecoveryLevel(f.recoveryLevel), int(f.size))
	exitOnErr(err, "encode URL")
	_, err = os.Stdout.Write(png)
	exitOnErr(err, "write image to output")
}

type Flags struct {
	url           string
	size          uint
	recoveryLevel recoveryLevel
}

func parseFlags() Flags {
	var f Flags
	flag.StringVar(&f.url, "url", "", "the url to encode")
	flag.UintVar(&f.size, "size", 1024, "width/height in pixels")
	flag.Var(&f.recoveryLevel, "level", "recovery level: low|medium|high|highest")
	flag.Parse()

	if f.url == "" {
		flag.Usage()
		os.Exit(1)
	}
	return f
}

type recoveryLevel qrcode.RecoveryLevel

const (
	recoveryLevelLow     recoveryLevel = recoveryLevel(qrcode.Low)
	recoveryLevelMedium  recoveryLevel = recoveryLevel(qrcode.Medium)
	recoveryLevelHigh    recoveryLevel = recoveryLevel(qrcode.High)
	recoveryLevelHighest recoveryLevel = recoveryLevel(qrcode.Highest)
)

var recoveryLevelStrings = map[recoveryLevel]string{
	recoveryLevelLow:     "low",
	recoveryLevelMedium:  "medium",
	recoveryLevelHigh:    "high",
	recoveryLevelHighest: "highest",
}

func (r *recoveryLevel) Set(val string) error {
	for level, levelString := range recoveryLevelStrings {
		if levelString == val {
			*r = level
			return nil
		}
	}
	return fmt.Errorf("invalid recovery level: %q", val)
}

func (r recoveryLevel) String() string {
	s, ok := recoveryLevelStrings[r]
	if !ok {
		return fmt.Sprintf("recoveryLevel(%d)", r)
	}
	return s
}

// checkURL logs a warning if val does not look like a value HTTP request URI.
func checkURL(val string) {
	_, err := url.ParseRequestURI(val)
	if err != nil {
		log.Printf("warning: %q is not a valid URL: %s", val, err)
	}
}

func exitOnErr(err error, message string) {
	if err != nil {
		log.Fatal(fmt.Errorf("%s: %w", message, err))
	}
}
