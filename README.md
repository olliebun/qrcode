# qrcode

![Test](https://github.com/github/docs/actions/workflows/test.yml/badge.svg)

Simple command line utility to generate a QR code PNG image from a URL.

Thin wrapper around [`skip2/go-qrcode`](https://github.com/skip2/go-qrcode).

## Installation

Install from the latest [release](https://github.com/olliebun/qrcode/releases). Please open an [issue](https://github.com/olliebun/qrcode/issues) if your platform is not supported.

Or install with Go:

```sh
go install github.com/olliebun/qrcode
```

## Usage

The `-url` flag specifies the URL to encode. The image is printed to stdout.

```sh
qrcode -url "https://example.com" > output.png
```

The default image is `1024x1024` pixels. Use `-size` to change it.

Use `-level` to set the error detection / recovery level. Valid values are `low`, `medium`, `high` and `highest`. The default value is `low`.

```
qrcode -h
Usage of ./qrcode:
  -level value
        recovery level: low|medium|high|highest
  -size uint
        width/height in pixels (default 1024)
  -url string
        the url to encode
```

## Errata

- `-url` can really be any string
