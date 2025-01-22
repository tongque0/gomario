package terrain

import (
	"gomario/assets"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Terrain struct {
	X, Y           float64
	Width, Height  float64
	Kind           int
	Graphics       *TerrainGraphics
	ScaleX, ScaleY float64
	AnimationTimer time.Time
	IsAnimating    bool
	Destroyed      bool
}

// NewTerrain 创建地形
func NewTerrain(gridX, gridY int, kind int) *Terrain {
	var graphics *TerrainGraphics
	switch kind {
	case 0:
		graphics = NewUnbreakableWall()
	case 1:
		graphics = NewBreakableWall()
	case 2:
		graphics = NewBreakableWall2()
	default:
		graphics = NewTerrainGraphics(32, 0, 16, 16)
	}

	// 计算缩放比例
	scaleX := assets.CellSize / float64(graphics.Width)
	scaleY := assets.CellSize / float64(graphics.Height)

	terrain := Terrain{
		X:         float64(gridX) * assets.CellSize, // 使用网格坐标计算实际坐标
		Y:         float64(gridY) * assets.CellSize, // 使用网格坐标计算实际坐标
		Kind:      kind,
		Graphics:  graphics,
		Width:     float64(graphics.Width) * scaleX,  // 根据图片宽度和缩放比例设置宽度
		Height:    float64(graphics.Height) * scaleY, // 根据图片高度和缩放比例设置高度
		ScaleX:    scaleX,
		ScaleY:    scaleY,
		Destroyed: false,
	}

	return &terrain
}

func (t *Terrain) Update() {
	if t.IsAnimating {
		if time.Since(t.AnimationTimer) > 500*time.Millisecond {
			t.IsAnimating = false
		}
	}
}
func (t *Terrain) Draw(screen *ebiten.Image) {
	if t.Destroyed {
		return // 如果地形已被销毁，则不绘制
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(t.ScaleX, t.ScaleY)
	op.GeoM.Translate(t.X, t.Y)

	// // 如果正在动画中，则向上移动 4 像素
	// if t.IsAnimating {
	// 	op.GeoM.Translate(0, -4) // 向上移动 4 像素
	// }

	screen.DrawImage(t.Graphics.Image, op)

	// Debug mode: 绘制地形边框
	if assets.IsDebug {
		red := color.RGBA{255, 0, 0, 255}
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				screen.Set(int(t.X)+dx, int(t.Y)+dy, red)                  // Top-left
				screen.Set(int(t.X+t.Width)+dx, int(t.Y)+dy, red)          // Top-right
				screen.Set(int(t.X)+dx, int(t.Y+t.Height)+dy, red)         // Bottom-left
				screen.Set(int(t.X+t.Width)+dx, int(t.Y+t.Height)+dy, red) // Bottom-right
			}
		}
	}
}

// StartAnimation 开始动画
func (t *Terrain) StartAnimation() {
	t.IsAnimating = true
	t.AnimationTimer = time.Now()
}

// Destroy 销毁地形
func (t *Terrain) Destroy() {
	t.Destroyed = true
}

// OnMarioCollision 处理马里奥与地形的碰撞
// 默认情况下,地形不会有变化
func (t *Terrain) OnMarioCollision(direction int) {
	switch t.Kind {
	case 0:
		return
	default:
		return
	}
}
