package configs

import (
	"GifToTelegramSticker/consoleIO"
)

type Config struct {
	GifPath string
	SpeedUp bool
	Cut     bool
}

func GetConfig(IO *consoleIO.ManagerIO) (*Config, error) {
	path, err := IO.Request("Enter path to gif file:")
	if err != nil {
		return nil, err
	}

	var speedUp bool
	var cut bool
	for {
		mode, err := IO.Request("Speed up gif or cut to 3s? [s/c]:")
		if err != nil {
			return nil, err
		}
		if len(mode) == 1 {
			if mode == "s" {
				speedUp = true
				break
			} else if mode == "c" {
				cut = true
				break
			}
		}
	}

	return &Config{
		GifPath: path,
		SpeedUp: speedUp,
		Cut:     cut,
	}, nil
}
