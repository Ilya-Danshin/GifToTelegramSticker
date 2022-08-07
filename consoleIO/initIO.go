package consoleIO

import (
	"bufio"
	"os"
)

type ManagerIO struct {
	in  *bufio.Reader
	out *bufio.Writer
}

func InitIO() *ManagerIO {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	return &ManagerIO{
		in:  in,
		out: out,
	}
}
