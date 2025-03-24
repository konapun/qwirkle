package tui

import (
	"bufio"
	"fmt"
	"os"
)

// Input translates text-based user input into actions.
type Input struct {
	reader *bufio.Reader
}

func NewInput() *Input {
	reader := bufio.NewReader(os.Stdin)
	return &Input{reader}
}

func (i *Input) Read() (string, error) {
	fmt.Print("> ")
	return i.reader.ReadString('\n')
}
