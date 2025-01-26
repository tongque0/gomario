package item

import (
	"gomario/assets"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ItemGraphics struct {
	Image  *ebiten.Image
	Width  int
	Height int
}

// 从精灵图中切割出指定位置和尺寸的子图像
func loadSubImage(spriteSheet *ebiten.Image, x, y, width, height int) (*ebiten.Image, int, int) {
	subImage := spriteSheet.SubImage(image.Rect(x, y, x+width, y+height)).(*ebiten.Image)
	return subImage, width, height
}

func NewItemGraphics(x, y, width, height int) *ItemGraphics {
	ig := &ItemGraphics{}
	spriteSheet, _, err := ebitenutil.NewImageFromFileSystem(assets.Assets, "graphics/item_objects.png")
	if err != nil {
		log.Fatal(err)
	}
	ig.Image, ig.Width, ig.Height = loadSubImage(spriteSheet, x, y, width, height)
	return ig
}

// 创建蘑菇的图像资源
func NewMushroom() *ItemGraphics {
	return NewItemGraphics(0, 0, 16, 16)
}

// 创建花的图像资源
func NewFlower() *ItemGraphics {
	return NewItemGraphics(0, 32, 16, 16)
}

// 创建星星的图像资源
func NewStar() *ItemGraphics {
	return NewItemGraphics(0, 64, 16, 16)
}

// 创建金币的图像资源
func NewCoin() *ItemGraphics {
	return NewItemGraphics(0, 96, 16, 16)
}
