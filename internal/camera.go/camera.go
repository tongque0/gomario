package camera

import (
	"gomario/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	X, Y                          float64
	ViewportWidth, ViewportHeight int
	MaxX, MaxY                    float64 // 地图边界
}

type Drawable interface {
	Draw(screen *ebiten.Image, camera *Camera)
	GetPosition() (x, y float64)
	GetSize() (width, height float64)
}

func NewCamera(viewportW, viewportH int) *Camera {
	return &Camera{
		ViewportWidth:  viewportW,
		ViewportHeight: viewportH,
	}
}

// 更新摄像头位置，跟随目标并限制边界
func (c *Camera) Update(targetX, targetY float64) {
	// 将目标置于视口中心
	newX := targetX - float64(c.ViewportWidth)/2
	newY := targetY - float64(c.ViewportHeight)/2

	// 限制边界
	if newX < 0 {
		newX = 0
	} else if newX > c.MaxX {
		newX = c.MaxX
	}

	if newY < 0 {
		newY = 0
	} else if newY > c.MaxY {
		newY = c.MaxY
	}

	c.X = newX
	c.Y = newY
}

// 设置地图边界（在关卡加载时调用）
func (c *Camera) SetBounds(MaxX, MaxY int) {
	c.MaxX = float64(MaxX*assets.CellSize) - float64(c.ViewportWidth)
	if c.MaxX < 0 {
		c.MaxX = 0
	}
	c.MaxY = float64(MaxY*assets.CellSize) - float64(c.ViewportHeight)
	if c.MaxY < 0 {
		c.MaxY = 0
	}
}
