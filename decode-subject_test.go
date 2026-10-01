package main

import (
	"strings"
	"testing"
)

func TestDecodeSubject(t *testing.T) {
	tests := []struct {
		name    string
		subject string
		want    string
	}{
		{"plain ascii", "Hello world", "Hello world"},
		{"empty", "", ""},
		{"utf-8 base64", "=?UTF-8?B?SGFsbG8gd8OpcmVsZA==?=", "Hallo wéreld"},
		{"utf-8 quoted-printable", "=?UTF-8?Q?Hallo_w=C3=A9reld?=", "Hallo wéreld"},
		{"iso-8859-1", "=?ISO-8859-1?Q?Caf=E9?=", "Café"},
		{"windows-1252", "=?windows-1252?Q?Caf=E9_=80_price?=", "Café € price"},
		{"iso-8859-15", "=?ISO-8859-15?Q?=A4uro?=", "€uro"},
		{"shift_jis", "=?Shift_JIS?B?k/qWe4zq?=", "日本語"},
		{"koi8-r", "=?KOI8-R?B?8NLJ18XU?=", "Привет"},
		{"lowercase charset label", "=?utf-8?q?caf=C3=A9?=", "café"},
		{"mixed encoded and plain", "Re: =?UTF-8?Q?caf=C3=A9?= order", "Re: café order"},
		{"adjacent encoded-words", "=?UTF-8?Q?Hallo_?= =?UTF-8?Q?w=C3=A9reld?=", "Hallo wéreld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := "Subject: " + tt.subject + "\r\n\r\nbody\r\n"
			got, err := decodeSubject(strings.NewReader(msg))
			if err != nil {
				t.Fatalf("decodeSubject() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("decodeSubject() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDecodeSubjectNoSubjectHeader(t *testing.T) {
	got, err := decodeSubject(strings.NewReader("From: a@example.com\r\n\r\nbody\r\n"))
	if err != nil {
		t.Fatalf("decodeSubject() error = %v", err)
	}
	if got != "" {
		t.Errorf("decodeSubject() = %q, want empty string", got)
	}
}

func TestDecodeSubjectErrors(t *testing.T) {
	tests := []struct {
		name string
		msg  string
	}{
		{"unsupported charset", "Subject: =?x-bogus?Q?hi?=\r\n\r\nbody\r\n"},
		{"malformed header", "not a header\r\n\r\nbody\r\n"},
		{"empty input", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := decodeSubject(strings.NewReader(tt.msg)); err == nil {
				t.Error("decodeSubject() error = nil, want error")
			}
		})
	}
}
