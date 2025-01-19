package terrain

import (
	"gomario/assets"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TerrainGraphics struct {
	Images []*ebiten.Image
}

// 从精灵图中切割出指定位置和尺寸的子图像
func loadSubImage(spriteSheet *ebiten.Image, x, y, width, height int) *ebiten.Image {
	subImage := spriteSheet.SubImage(image.Rect(x, y, x+width, y+height)).(*ebiten.Image)
	return subImage
}

func NewTerrainGraphics() *TerrainGraphics {
	tg := &TerrainGraphics{}
	spriteSheet, _, err := ebitenutil.NewImageFromFileSystem(assets.Assets, "graphics/tile_set.png")
	if err != nil {
		log.Fatal(err)
	}
	tg.Images = append(tg.Images, loadSubImage(spriteSheet, 0, 0, 16, 16))
	return tg
}
