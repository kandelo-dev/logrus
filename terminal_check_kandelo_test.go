package logrus

import (
	"os"
	"testing"
)

func TestKandeloTerminalCheck(t *testing.T) {
	terminal, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer terminal.Close()
	if !isTerminal(int(terminal.Fd())) {
		t.Fatal("PTY master is not a terminal")
	}
	regular, err := os.CreateTemp("", "kandelo-logrus-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(regular.Name())
	defer regular.Close()
	if isTerminal(int(regular.Fd())) {
		t.Fatal("regular file reported as terminal")
	}
}
