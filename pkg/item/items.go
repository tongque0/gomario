package item

import (
	"gomario/internal/camera.go"

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
	X, Y          float64
	Width, Height float64
	Kind          string
	IsDestroy     bool
}

func NewItem(x, y int, kind string) *Item {
	return &Item{
		X:    float64(x),
		Y:    float64(y),
		Kind: kind,
	}
}

func (i *Item) Update() {

}
func (i *Item) Draw(screen *ebiten.Image, camera *camera.Camera) {

}
