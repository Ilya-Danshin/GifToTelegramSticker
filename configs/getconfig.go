package configs

import (
	"GifToTelegramSticker/consoleIO"
	"strconv"
	"time"
)

type Config struct {
	GifPath  string
	Duration time.Duration
}

func GetConfig(IO *consoleIO.ManagerIO) (*Config, error) {
	path, err := IO.Request("Enter path to gif file: ")
	if err != nil {
		return nil, err
	}

	cut, err := IO.Request("Duration, seconds: ")
	if err != nil {
		return nil, err
	}

	cutInt, err := strconv.Atoi(cut)
	if err != nil {
		return nil, err
	}
	duration := time.Duration(cutInt) * time.Second

	return &Config{
		GifPath:  path,
		Duration: duration,
	}, nil
}
