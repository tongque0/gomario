package enemies

import (
	"gomario/internal/camera.go"

	"github.com/hajimehoshi/ebiten/v2"
)

var _ camera.Drawable = (*Enemies)(nil)

type Enemies struct {
	X, Y                 float64 // 位置
	VelocityX, VelocityY float64 // 速度
	Width, Height        float64 // 尺寸
	Kind                 string  // 种类

	IsDead   bool             // 是否死亡
	Graphics *EnemiesGraphics // 敌人的图像资源
}

func NewEnemies(x, y int, kind string) *Enemies {
	return &Enemies{
		X:    float64(x),
		Y:    float64(y),
		Kind: kind,
	}
}

func (e *Enemies) Update() {

}

func (e *Enemies) Draw(screen *ebiten.Image, camera *camera.Camera) {
	// // 计算相对于摄像机的位置
	// x, y := e.GetPosition()
	// x -= camera.X
	// y -= camera.Y

	// // 绘制敌人
	// opts := &ebiten.DrawImageOptions{}
	// opts.GeoM.Translate(x, y)
	// opts.GeoM.Scale(e.ScaleX, e.ScaleY)
	// screen.DrawImage(e.Graphics.Image, opts)
}

func (e *Enemies) GetPosition() (x, y float64) {
	return e.X, e.Y
}

func (e *Enemies) GetSize() (width, height float64) {
	return e.Width, e.Height
}
