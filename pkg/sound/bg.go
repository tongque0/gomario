package sound

import (
	"gomario/assets"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

var audioContext *audio.Context

func init() {
	audioContext = audio.NewContext(44100)
}
func NewBgSoundPlayer() *audio.Player {
	return loadMusic("music/main_theme.ogg")
}

func loadMusic(filename string) *audio.Player {
	// 打开背景音乐文件（OGG格式）
	file, err := assets.Assets.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// 解码OGG音频文件
	audioData, err := vorbis.DecodeWithSampleRate(44100, file)
	if err != nil {
		log.Fatal(err)
	}

	// 创建音频播放器
	player, err := audioContext.NewPlayer(audioData)
	if err != nil {
		log.Fatal(err)
	}
	return player
}

