package terrain

import (
	"fmt"
	"gomario/assets"
	destroyeffect "gomario/pkg/destroyEffect"
	"gomario/pkg/sound"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Terrain struct {
	X, Y           float64
	Width, Height  float64
	Kind           int
	Graphics       *TerrainGraphics
	ScaleX, ScaleY float64
	Destroyed      bool
}
type TerrainDebris struct {
	X, Y                 float64
	VelocityX, VelocityY float64
	Width, Height        float64
	Rotation             float64
	RotationSpeed        float64
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

}
func (t *Terrain) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(t.ScaleX, t.ScaleY)

	op.GeoM.Translate(t.X, t.Y)
	if t.Destroyed {
		return
	}

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

// Destroy 销毁地形
func (t *Terrain) Destroy() {
	t.Destroyed = true
	destroyeffect.NewTerrainEffect(t.X, t.Y, t.Width, t.Height, *t.Graphics.Image, t.ScaleX, t.ScaleY)
}

// OnMarioCollision 处理马里奥与地形的碰撞
// 默认情况下,地形不会有变化
func (t *Terrain) OnMarioCollision(direction int, isbig, isjumping bool) {
	switch t.Kind {
	case 0: // 不可破坏的墙壁
		fmt.Println("Unbreakable wall")
	case 1: // 可破坏的墙壁
		if direction == 1 && isbig && isjumping {
			sound.NewSfxPlayer("brick_smash")
			t.Destroy()
		}
	default:
		return
	}
}
