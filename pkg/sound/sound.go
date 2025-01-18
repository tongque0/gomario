package sound

import (
	"gomario/assets"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

func NewBgSoundPlayer() *audio.Player {

	// 创建音频播放器
	audioContext := audio.NewContext(44100)

	// 打开背景音乐文件（OGG格式）
	file, err := assets.Assets.Open("music/main_theme.ogg")
	if err != nil {
		log.Fatal(err)
	}

	// 解码OGG音频文件
	audioData, err := vorbis.DecodeWithSampleRate(44100, file)
	if err != nil {
		log.Fatal(err)
	}

	// 创建音频播放器
	bgSound, err := audioContext.NewPlayer(audioData)
	if err != nil {
		log.Fatal(err)
	}

	return bgSound
}
