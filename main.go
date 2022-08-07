package main

import (
	"GifToTelegramSticker/configs"
	"GifToTelegramSticker/consoleIO"
	"GifToTelegramSticker/ffmpeg"
	"GifToTelegramSticker/files"
	"fmt"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	IO := consoleIO.InitIO()

	userConfig, err := configs.GetConfig(IO)
	if err != nil {
		fmt.Printf("Get config error: %s", err.Error())
		return
	}

	gifConfig, err := files.GetImageConfig(userConfig.GifPath)
	if err != nil {
		fmt.Printf("Get image error: %s", err.Error())
	}

	cfg := ffmpeg.GetConfigForGif(userConfig, gifConfig)

	cfg.Run(IO)

}
