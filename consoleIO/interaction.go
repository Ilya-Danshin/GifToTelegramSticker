package consoleIO

import (
	"fmt"
	"strings"
)

func (io *ManagerIO) Write(str string) error {
	n, err := io.out.WriteString(str)
	if err != nil {
		return err
	}

	if n != len(str) {
		return fmt.Errorf("in string %s was writed %d symbols", str, n)
	}

	err = io.out.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (io *ManagerIO) Read() (string, error) {
	in, err := io.in.ReadString('\n')
	if err != nil {
		return "", err
	}
	in = strings.Trim(in, "\r\n")

	if len(in) == 0 {
		return "", fmt.Errorf("empty input")
	}

	return in, nil
}

func (io *ManagerIO) Request(question string) (string, error) {
	err := io.Write(question)
	if err != nil {
		return "", err
	}

	answer, err := io.Read()
	if err != nil {
		return "", err
	}

	return answer, nil
}
