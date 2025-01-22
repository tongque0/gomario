// Description: 用于加载地形图像的包
// Created by tongque on 25-1-20
// 创建 TerrainGraphics 实例时，传入对应种类序号，即可加载对应的地形图像
// 以下为地形种类序号：
// 0: 不可破坏的墙壁 1: 可破坏的墙壁 2:
package terrain

import (
	"gomario/assets"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TerrainGraphics struct {
	Image  *ebiten.Image
	Width  int
	Height int
}

// 从精灵图中切割出指定位置和尺寸的子图像
func loadSubImage(spriteSheet *ebiten.Image, x, y, width, height int) (*ebiten.Image, int, int) {
	subImage := spriteSheet.SubImage(image.Rect(x, y, x+width, y+height)).(*ebiten.Image)
	return subImage, width, height
}

// 创建 TerrainGraphics 实例
func NewTerrainGraphics(x, y, width, height int) *TerrainGraphics {
	tg := &TerrainGraphics{}
	spriteSheet, _, err := ebitenutil.NewImageFromFileSystem(assets.Assets, "graphics/tile_set.png")
	if err != nil {
		log.Fatal(err)
	}
	tg.Image, tg.Width, tg.Height = loadSubImage(spriteSheet, x, y, width, height)
	return tg
}

// 创建不可破坏的墙壁
func NewUnbreakableWall() *TerrainGraphics {
	return NewTerrainGraphics(0, 16, 16, 16)
}

// 创建可破坏的墙壁
func NewBreakableWall() *TerrainGraphics {
	return NewTerrainGraphics(16, 0, 16, 16)
}

func NewBreakableWall2() *TerrainGraphics {
	return NewTerrainGraphics(0, 0, 16, 16)
}
