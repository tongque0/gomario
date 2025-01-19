package game

import (
	"gomario/internal/level"
	"gomario/pkg/sound"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Game struct {
	Level   *level.Level
	BgSound *audio.Player
}

func (g *Game) Update() error {
	// 如果背景音频还没有开始播放，则开始播放
	if g.BgSound != nil && !g.BgSound.IsPlaying() {
		g.BgSound.Play()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Level.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// 保持 16:9 的宽高比
	const aspectRatio = 16.0 / 9.0
	width := outsideWidth
	height := int(float64(outsideWidth) / aspectRatio)
	if height > outsideHeight {
		height = outsideHeight
		width = int(float64(outsideHeight) * aspectRatio)
	}
	return width, height
}

func Run() {
	// 初始化音频上下文
	bgsound := sound.NewBgSoundPlayer()
	// 初始化游戏并设置背景音频
	game := &Game{
		Level:   level.NewLevel(),
		BgSound: bgsound,
	}

	// 设置窗口尺寸
	ebiten.SetWindowSize(1280, 720) // 初始窗口大小
	ebiten.SetWindowTitle("gomario")

	// 运行游戏
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
