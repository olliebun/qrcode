// qrcode is a utility to generate a QR code image from a URL.
// Run qrcode -h to see usage.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	f := parseFlags()
	png, err := qrcode.Encode(f.url, qrcode.RecoveryLevel(f.recoveryLevel), int(f.size))
	exitOnErr(err)
	_, err = os.Stdout.Write(png)
	exitOnErr(err)
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

func exitOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
