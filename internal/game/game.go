package game

import (
	"gomario/assets"
	"gomario/internal/level"
	destroyeffect "gomario/pkg/destroyEffect"
	"gomario/pkg/sound"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Game struct {
	Level         *level.Level
	BgSound       *audio.Player
	destroyeffect *destroyeffect.DestroyEffect
}

func (g *Game) Update() error {
	// 如果背景音频还没有开始播放，则开始播放
	if g.BgSound != nil && !g.BgSound.IsPlaying() {
		g.BgSound.Play()
	}
	g.Level.Update()
	g.destroyeffect.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Level.Draw(screen)
	g.destroyeffect.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func Run() {
	// 初始化音频上下文
	bgsound := sound.NewBgSoundPlayer()
	// 初始化游戏并设置背景音频
	game := &Game{
		Level:         level.NewLevel(),
		BgSound:       bgsound,
		destroyeffect: destroyeffect.NewDestroyEffect(),
	}

	// 设置窗口尺寸
	ebiten.SetWindowSize(assets.ScreenWidth, assets.ScreenHeight) // 初始窗口大小
	ebiten.SetWindowTitle("gomario")

	// 运行游戏
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
