package mario

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Mario struct {
	X, Y  float64
	State int
}

func NewMario() *Mario {
	return &Mario{
		X:     0,
		Y:     0,
		State: 0,
	}
}
func (m *Mario) Update() {
	m.X++
	m.Y++
}

func (m *Mario) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Hello, World!", int(m.X), int(m.Y))
}
