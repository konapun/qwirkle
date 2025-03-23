package tui

import (
	"bufio"
	"os"
)

// Input is a text-based user interface input device.
type Input struct {
	reader *bufio.Reader
}

func NewInput() *Input {
	reader := bufio.NewReader(os.Stdin)
	return &Input{reader}
}

// Read reads input from the user.
func (i *Input) Read() (string, error) {
	return i.reader.ReadString('\n') // Read until newline
}
