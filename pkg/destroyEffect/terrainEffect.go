// Description: 此文件定义了销毁地形的动画效果
package destroyeffect

import (
	"time"

	"gomario/internal/camera.go"

	"github.com/hajimehoshi/ebiten/v2"
)

type TerrainEffect struct {
	Debris         []*TerrainDebris
	Image          ebiten.Image
	ScaleX, ScaleY float64
	StartTime      time.Time
	IsFinished     bool
}
type TerrainDebris struct {
	X, Y                 float64
	VelocityX, VelocityY float64
	Width, Height        float64
	Rotation             float64
	RotationSpeed        float64
}

func NewTerrainEffect(x, y, width, height float64, image ebiten.Image, scaleX, scaleY float64) {
	debrisWidth := width / 4
	debrisHeight := height / 4
	centerX := x + width/2
	centerY := y + height/2

	effect := &TerrainEffect{
		Image:     image,
		ScaleX:    scaleX,
		ScaleY:    scaleY,
		StartTime: time.Now(),
		Debris: []*TerrainDebris{
			{X: centerX - debrisWidth, Y: centerY - debrisHeight,
				Width: debrisWidth, Height: debrisHeight,
				VelocityX: -3, VelocityY: -10, RotationSpeed: 0.15},
			{X: centerX, Y: centerY - debrisHeight,
				Width: debrisWidth, Height: debrisHeight,
				VelocityX: 3, VelocityY: -10, RotationSpeed: -0.15},
			{X: centerX - debrisWidth, Y: centerY,
				Width: debrisWidth, Height: debrisHeight,
				VelocityX: -3, VelocityY: -8, RotationSpeed: -0.15},
			{X: centerX, Y: centerY,
				Width: debrisWidth, Height: debrisHeight,
				VelocityX: 3, VelocityY: -8, RotationSpeed: 0.15},
		},
	}
	// 将销毁地形的动画效果添加到销毁动画列表中
	terrainEffect = append(terrainEffect, effect)
}

func (e *TerrainEffect) Update() {
	for _, debris := range e.Debris {
		debris.VelocityY += 0.5
		debris.X += debris.VelocityX
		debris.Y += debris.VelocityY
		debris.Rotation += debris.RotationSpeed
	}

	if time.Since(e.StartTime) > time.Duration(0.66*float64(time.Second)) {
		e.IsFinished = true
	}
}

func (e *TerrainEffect) Draw(screen *ebiten.Image, camera *camera.Camera) {
	for _, debris := range e.Debris {
		// 计算屏幕坐标
		screenX := debris.X - camera.X
		screenY := debris.Y - camera.Y

		// 创建绘制选项
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-debris.Width/2, -debris.Height/2) // 旋转中心点
		op.GeoM.Rotate(debris.Rotation)                      // 应用旋转
		op.GeoM.Translate(debris.Width/2, debris.Height/2)   // 恢复位置
		op.GeoM.Scale(e.ScaleX/2, e.ScaleY/2)                // 应用缩放
		op.GeoM.Translate(screenX, screenY)                  // 应用屏幕坐标

		// 绘制碎片
		screen.DrawImage(&e.Image, op)
	}
}
