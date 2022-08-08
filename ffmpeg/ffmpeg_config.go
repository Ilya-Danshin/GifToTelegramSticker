package ffmpeg

import (
	"fmt"
	"image"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"GifToTelegramSticker/configs"
	"GifToTelegramSticker/consoleIO"
)

type Config struct {
	file     string
	speed    int
	fps      int
	width    int
	height   int
	quality  int
	codec    string
	duration time.Duration
	outFile  string
}

var stickerWidthOrHeight = 512
var stickerDuration = 3 * time.Second
var stickerExtension = ".webm"
var stickerQuality = 1024
var stickerSize = int64(256) * 1024

var defaultCfg = Config{
	file:     "",
	speed:    1,
	fps:      30,
	width:    stickerWidthOrHeight,
	height:   stickerWidthOrHeight,
	quality:  stickerQuality,
	codec:    "libvpx-vp9",
	duration: stickerDuration,
	outFile:  "",
}

func GetConfigForGif(config *configs.Config, gifCfg *image.Config) *Config {
	cfg := defaultCfg
	cfg.file = config.GifPath
	cfg.width, cfg.height = calculateStickerSize(gifCfg)
	cfg.outFile = changeFileExtension(config.GifPath)

	if config.Duration != time.Duration(0) {
		cfg.duration = config.Duration
	}

	return &cfg
}

func changeFileExtension(file string) string {
	name := strings.TrimSuffix(file, filepath.Ext(file))
	return name + stickerExtension
}

func calculateStickerSize(gifCfg *image.Config) (int, int) {
	var width, height int

	if gifCfg.Width > gifCfg.Height {
		scale := float64(stickerWidthOrHeight) / float64(gifCfg.Width)
		width = stickerWidthOrHeight
		height = int(math.Round(scale * float64(gifCfg.Height)))
	} else {
		scale := float64(stickerWidthOrHeight) / float64(gifCfg.Height)
		height = stickerWidthOrHeight
		width = int(math.Round(scale * float64(gifCfg.Width)))
	}

	return width, height
}

func (cfg *Config) Run(io *consoleIO.ManagerIO) error {

	var end bool
	for !end {
		err := cfg.TryToConvert(io)
		if err != nil {
			return err
		}
		end, err = cfg.CorrectQualityBySize()
		if err != nil {
			return err
		}
	}

	return nil
}

func (cfg *Config) CorrectQualityBySize() (bool, error) {
	fileInfo, err := os.Stat(cfg.outFile)
	if err != nil {
		return false, err
	}

	if stickerSize-fileInfo.Size() < 0 { // Need to less file
		cfg.quality /= 2
	} else {
		if stickerSize-fileInfo.Size() < 16*1024 {
			return true, nil
		} else {
			cfg.quality += cfg.quality / 2
		}
	}

	return false, nil
}

func (cfg *Config) TryToConvert(io *consoleIO.ManagerIO) error {
	args := cfg.toArgs()

	out, err := runCMD("ffmpeg/ffmpeg.exe", args)
	if err != nil {
		return err
	}

	err = io.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) toArgs() []string {
	speed := strconv.Itoa(cfg.speed)
	pts := speed + "*PTS"
	fps := strconv.Itoa(cfg.fps)
	scale := strconv.Itoa(cfg.width) + "x" + strconv.Itoa(cfg.height)
	quality := strconv.Itoa(cfg.quality)
	duration := getFormatDuration(cfg.duration)

	args := []string{"-y", "-i", cfg.file, "-vf", "setpts=" + pts, "-r", fps, "-vf", "scale=" + scale, "-b:v", quality + "k",
		"-c:v", cfg.codec, "-c:a", "libopus", "-an", "-metadata:s:v:0", "alpha_mode=\"1\"", "-ss", "00:00:00.000",
		"-t", duration, cfg.outFile}

	return args
}

func getFormatDuration(d time.Duration) string {
	d = d.Round(time.Millisecond)
	h := d / time.Hour
	d -= h * time.Hour
	min := d / time.Minute
	d -= min * time.Minute
	s := d / time.Second
	d -= s * time.Second
	mil := d / time.Millisecond
	return fmt.Sprintf("%02d:%02d:%02d.%03d", h, min, s, mil)
}

func runCMD(exe string, args []string) (out string, err error) {
	cmd := exec.Command(exe, args...)

	var b []byte
	b, err = cmd.CombinedOutput()
	out = string(b)

	return
}
