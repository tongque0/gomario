package terrain

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Terrain struct {
	X, Y          float64
	Width, Height float64
	Kind          int
	Graphics      *TerrainGraphics
}

func NewTerrain(x, y float64, Kind int) *Terrain {
	terrain := Terrain{
		X:    x,
		Y:    y,
		Kind: 0,
	}
	switch Kind {
	case 0:
		terrain.Graphics = NewTerrainGraphics()
	default:
		terrain.Graphics = NewTerrainGraphics()
	}
	return &terrain
}

func (t *Terrain) Update() {
}
func (t *Terrain) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	fmt.Println(t.X, t.Y)
	op.GeoM.Translate(t.X, t.Y)
	screen.DrawImage(t.Graphics.Images[0], op)
}
