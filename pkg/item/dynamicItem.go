// description: 动态物品
// 作用：这些元素是动态生成的，不在预定义地图之中，如碰撞出现蘑菇，所以采用了类似销毁动画的实现方式。

package item

import (
	"gomario/internal/camera.go"

	"github.com/hajimehoshi/ebiten/v2"
)

var DynamicItems []*Item

func AddDynamicItem(item *Item) {
	DynamicItems = append(DynamicItems, item)
}

type DynamicItem struct {
}

func NewDynamicItem() *DynamicItem {
	return &DynamicItem{}
}

func (d *DynamicItem) Update() {
	for i := 0; i < len(DynamicItems); i++ {
		if DynamicItems[i].IsDestroy {
			DynamicItems = append(DynamicItems[:i], DynamicItems[i+1:]...)
			i--
		} else {
			DynamicItems[i].Update()
		}
	}
}
func (d *DynamicItem) Draw(screen *ebiten.Image, camera *camera.Camera) {
	for i := 0; i < len(DynamicItems); i++ {
		DynamicItems[i].Draw(screen, camera)
	}
}
