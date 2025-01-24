// Description: 这个包定义了销毁实体的动画效果
// Created by tongque on 25-1-22
// 为什么定义这个包？
// 1. 我发现我的结构设计，将销毁动画放在实体，不是很优雅(其实就是有bug)
// 2. 我将销毁动画的逻辑单独抽出来，这样可以更好的控制销毁动画的逻辑
// 3. 写到这里的时候，这个游戏的逻辑已经有些复杂了，但是首次编写，我希望以现在的思路将其写完，如果有机会再进行重构

// 在编写这个包的时候，我突然意识到，让这个包独立运行更加好一点
package destroyeffect

import (
	"gomario/internal/camera.go"

	"github.com/hajimehoshi/ebiten/v2"
)

var terrainEffect = []*TerrainEffect{} //待销毁地形的动画效果

type DestroyEffect struct {

}

func NewDestroyEffect() *DestroyEffect {
	return &DestroyEffect{}
}

func (e *DestroyEffect) Update() {
	// 遍历销毁动画列表，更新销毁动画
	for i := 0; i < len(terrainEffect); i++ {
		if terrainEffect[i].IsFinished {
			terrainEffect = append(terrainEffect[:i], terrainEffect[i+1:]...)
			i--
		} else {
			terrainEffect[i].Update()
		}
	}
}

func (e *DestroyEffect) Draw(screen *ebiten.Image,camera *camera.Camera) {
	for i := 0; i < len(terrainEffect); i++ {
		terrainEffect[i].Draw(screen,camera)
	}
}
