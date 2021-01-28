# ansicsi

[![PkgGoDev](https://pkg.go.dev/badge/github.com/pgavlin/ansicsi)](https://pkg.go.dev/github.com/pgavlin/ansicsi)
[![codecov](https://codecov.io/gh/pgavlin/ansicsi/branch/master/graph/badge.svg)](https://codecov.io/gh/pgavlin/ansicsi)
[![Go Report Card](https://goreportcard.com/badge/github.com/pgavlin/ansicsi)](https://goreportcard.com/report/github.com/pgavlin/ansicsi)
[![Test](https://github.com/pgavlin/ansicsi/workflows/Test/badge.svg)](https://github.com/pgavlin/ansicsi/actions?query=workflow%3ATest)

ansicsi provides a Go package that decodes and encodes ANSI control sequences as defined in ECMA-48/ANSI X3.64.

The high-level decoder currently only supports the Set Graphics Rendition control function. All other control
functions are returned as a tuple of (parameter bytes, intermediate bytes, final byte).

The decoder can be called in a loop in order to separate control sequences from normal text:

```go
for len(bytes) > 0 {
	if cmd, size := Decode(bytes); size > 0 {
		switch cmd := cmd.(type) {
			// Handle control functions here
		}
		bytes = bytes[size:]
		continue
	}

	// Handle plain text here

	bytes = bytes[1:]
}
```

A command can be encoded using its Encode method:

```go
resetCommand := SetGraphicsRendition{Command: SGRReset}
sz, err := resetCommand.Encode(w)
```
