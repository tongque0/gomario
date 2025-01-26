package item

import (
	"gomario/assets"
	"gomario/internal/camera.go"
	"gomario/pkg/sound"

	"github.com/hajimehoshi/ebiten/v2"
)

var _ camera.Drawable = (*Item)(nil)

func (i *Item) GetPosition() (x, y float64) {
	return i.X, i.Y
}

func (i *Item) GetSize() (width, height float64) {
	return i.Width, i.Height
}

type Item struct {
	X, Y           float64
	ScaleX, ScaleY float64
	Width, Height  float64
	Kind           string
	IsDestroy      bool
	Graphics       *ItemGraphics
}

func NewItem(gridX, gridY int, kind string) *Item {
	graphics := &ItemGraphics{}
	switch kind {
	case "mushroom":
		graphics = NewMushroom()
	case "flower":
		graphics = NewFlower()
	case "star":
		graphics = NewStar()
	case "coin":
		graphics = NewCoin()
	default:
		graphics = NewCoin()
	}
	// 计算缩放比例
	scaleX := assets.CellSize / float64(graphics.Width)
	scaleY := assets.CellSize / float64(graphics.Height)

	return &Item{
		X:         float64(gridX) * assets.CellSize,  // 使用网格坐标计算实际坐标
		Y:         float64(gridY) * assets.CellSize,  // 使用网格坐标计算实际坐标
		Width:     float64(graphics.Width) * scaleX,  // 根据图片宽度和缩放比例设置宽度
		Height:    float64(graphics.Height) * scaleY, // 根据图片高度和缩放比例设置高度
		Kind:      kind,
		Graphics:  graphics,
		IsDestroy: false,
		ScaleX:    scaleX,
		ScaleY:    scaleY,
	}
}

func (i *Item) Update() {
	if i.IsDestroy {
		return
	}
}
func (i *Item) Draw(screen *ebiten.Image, camera *camera.Camera) {
	// 计算屏幕坐标
	screenX := i.X - camera.X
	screenY := i.Y - camera.Y

	// 如果地形被销毁，则不绘制
	if i.IsDestroy {
		return
	}

	// 创建绘制选项
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(i.ScaleX, i.ScaleY)
	op.GeoM.Translate(screenX, screenY)

	// 绘制物品图形
	screen.DrawImage(i.Graphics.Image, op)

}

func (i *Item) OnMarioCollision(direction int, isBig bool) {
	switch i.Kind {
	case "mushroom":
		if !isBig {
			// 马里奥变大
			// mario.IsBig = true
			// mario.Graphics = marioGraphicsBig
			// mario.Height = float64(marioGraphicsBig.BigHeight) * mario.ScaleY
			// mario.Y -= mario.Height - float64(marioGraphics.SmallHeight)*mario.ScaleY
		}
	case "flower":
		// 马里奥变为火焰马里奥
		// mario.IsFire = true
		// mario.Graphics = marioGraphicsFire
	case "star":
		// 马里奥变为无敌状态
		// mario.IsInvincible = true
	case "coin":
		// 马里奥获得金币
		// mario.Coin++
		sound.NewSfxPlayer("coin")
	}
	if i.IsDestroy {
		return
	}
	i.IsDestroy = true
}
