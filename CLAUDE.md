# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

Requires Go 1.26 or later.

```sh
go build -o decode-subject .                # build (binary is gitignored)
go test ./...                               # run all tests
go test -run '^TestDecodeSubject$/koi8-r' . # run a single table-driven subtest
go vet ./...
```

## Architecture

A single-file CLI (`package main`) that reads an email message from stdin and prints its `Subject` header with RFC 2047 encoded-words decoded to UTF-8.

- `main` only wires stdin/stdout/stderr and the exit code; all logic lives in `decodeSubject(io.Reader)`, which is what the tests exercise.
- Parsing uses the standard library (`net/mail.ReadMessage`, `mime.WordDecoder`). The only external dependency is `golang.org/x/net/html/charset`, plugged in as `WordDecoder.CharsetReader` so that any WHATWG Encoding Standard label (ISO-8859-*, windows-125x, Shift_JIS, KOI8-R, aliases like `latin1`/`cp1252`, case-insensitive) is supported beyond the stdlib's built-in UTF-8/ISO-8859-1/US-ASCII.

## Behavior contract

The README documents these and tests cover them, so keep them stable:

- Missing `Subject` header → prints an empty line, exit 0.
- Unparseable message or unsupported charset → error on stderr, exit 1.

Tests in `decode-subject_test.go` are table-driven: add new charset/encoding cases as rows in `TestDecodeSubject` and failure cases in `TestDecodeSubjectErrors`.
