# decode-subject

Reads an email message from stdin and prints its `Subject` header with any
[RFC 2047](https://datatracker.ietf.org/doc/html/rfc2047) encoded-words
decoded to UTF-8.

```console
$ cat message.eml
Subject: =?UTF-8?B?SGFsbG8gd8OpcmVsZA==?=

body
$ decode-subject < message.eml
Hallo wéreld
```

## Supported charsets

Any charset label defined by the
[WHATWG Encoding Standard](https://encoding.spec.whatwg.org/#names-and-labels)
is supported, including UTF-8, ISO-8859-*, windows-125x, KOI8-R, Shift_JIS,
EUC-JP, GBK and Big5. Labels are matched case-insensitively and common aliases
such as `latin1` and `cp1252` work too.

## Install

### Download a binary

Prebuilt binaries for macOS, Linux and Windows (x86-64 and arm64) are attached
to each [release](https://github.com/breedeweg/decode-subject/releases/latest).
They need no Go installation. Download the archive for your platform, extract
it and put `decode-subject` somewhere on your `PATH`.

Verify a download against the release's `checksums.txt`:

```sh
shasum -a 256 -c checksums.txt --ignore-missing
```

The macOS binaries are not signed. If you download one with a browser and
macOS refuses to run it, remove the quarantine flag:

```sh
xattr -d com.apple.quarantine decode-subject
```

### Install with Go

Requires Go 1.26 or later.

```sh
go install github.com/breedeweg/decode-subject@latest
```

Or build from a checkout:

```sh
go build -o decode-subject .
```

## Test

```sh
go test ./...
```

## Exit status

| Code | Meaning |
|------|---------|
| 0 | Subject printed (an empty line if the message has no `Subject` header) |
| 1 | The message could not be parsed or the subject uses an unsupported charset; the error is written to stderr |
