package main

import (
	"GifToTelegramSticker/configs"
	"GifToTelegramSticker/consoleIO"
	"fmt"
)

func main() {

	IO := consoleIO.InitIO()

	config, err := configs.GetConfig(IO)
	if err != nil {
		fmt.Printf("Get config error: %s", err.Error())
		return
	}

	fmt.Println(config.GifPath)
	fmt.Println(config.SpeedUp)
	fmt.Println(config.Cut)
}
