package files

import (
	"image"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func GetImageConfig(path string) (*image.Config, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	gifImage, _, err := image.DecodeConfig(reader)
	if err != nil {
		return nil, err
	}

	return &gifImage, nil
}
