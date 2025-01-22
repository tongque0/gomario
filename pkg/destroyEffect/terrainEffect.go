// Description: 此文件定义了销毁地形的动画效果
package destroyeffect

import (
	"time"

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

func (e *TerrainEffect) Draw(screen *ebiten.Image) {
	for _, debris := range e.Debris {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-debris.Width/2, -debris.Height/2)
		op.GeoM.Rotate(debris.Rotation)
		op.GeoM.Translate(debris.Width/2, debris.Height/2)
		op.GeoM.Scale(e.ScaleX/2, e.ScaleY/2)
		op.GeoM.Translate(debris.X, debris.Y)
		screen.DrawImage(&e.Image, op)
	}
}
