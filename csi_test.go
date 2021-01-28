package ansicsi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSI(t *testing.T) {
	bytes := []byte("\x1b[38;5;128m\x1b[1mHello, world!\x1b[0m")

	expected := []struct {
		size int
		cmd  Command
	}{
		{size: 11, cmd: &SetGraphicsRendition{Command: SGRForegroundColor, Parameters: []int{5, 128}}},
		{size: 4, cmd: &SetGraphicsRendition{Command: SGRBold, Parameters: []int{}}},
		{size: 0},
		{size: 0},
		{size: 0},
		{size: 0},
		{size: 0},
		{size: 0},
		{size: 0},
		{size: 0},
		{size: 0},
		{size: 0},
		{size: 0},
		{size: 0},
		{size: 0},
		{size: 4, cmd: &SetGraphicsRendition{Command: SGRReset, Parameters: []int{}}},
	}
	for ; len(bytes) > 0; expected = expected[1:] {
		cmd, size := Decode(bytes)
		if !assert.Equal(t, expected[0].size, size) {
			return
		}
		if size == 0 {
			bytes = bytes[1:]
			continue
		}
		bytes = bytes[size:]

		if !assert.Equal(t, expected[0].cmd, cmd) {
			return
		}
	}
}
