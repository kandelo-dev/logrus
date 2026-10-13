package logrus

import "syscall"

func isTerminal(fd int) bool {
	_, err := syscall.GetTermios(fd)
	return err == nil
}
