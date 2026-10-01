package main

import (
	"fmt"
	"io"
	"mime"
	"net/mail"
	"os"

	"golang.org/x/net/html/charset"
)

func main() {
	subject, err := decodeSubject(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(subject)
}

// decodeSubject reads a mail message from r and returns its Subject header
// with any RFC 2047 encoded-words decoded.
func decodeSubject(r io.Reader) (string, error) {
	msg, err := mail.ReadMessage(r)
	if err != nil {
		return "", err
	}

	subject := msg.Header.Get("Subject")

	decoder := &mime.WordDecoder{CharsetReader: charset.NewReaderLabel}
	return decoder.DecodeHeader(subject)
}
