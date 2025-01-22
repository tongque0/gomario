package level

import (
	"gomario/pkg/enemies"
	"gomario/pkg/item"
	"gomario/pkg/mario"
	"gomario/pkg/physics"
	"gomario/pkg/terrain"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	Mario   *mario.Mario
	Enemies []*enemies.Enemies
	Item    []*item.Item
	Terrain []*terrain.Terrain
}

func NewLevel() *Level {
	return &Level{
		Mario: mario.NewMario(0, 0),
		// Enemies: []*enemies.Enemies{
		// 	enemies.NewEnemies(),
		// },
		// Item: []*item.Item{
		// 	item.NewItem(),
		// },
		Terrain: []*terrain.Terrain{
			terrain.NewTerrain(0, 14, 2),
			terrain.NewTerrain(1, 14, 2),
			terrain.NewTerrain(2, 14, 2),
			terrain.NewTerrain(3, 14, 2),
			terrain.NewTerrain(4, 14, 2),
			terrain.NewTerrain(5, 14, 2),
			terrain.NewTerrain(6, 14, 2),
			terrain.NewTerrain(7, 14, 2),
			terrain.NewTerrain(8, 14, 2),
			terrain.NewTerrain(9, 14, 2),
			terrain.NewTerrain(10, 14, 2),
			terrain.NewTerrain(11, 14, 2),
			terrain.NewTerrain(12, 14, 2),
			terrain.NewTerrain(13, 14, 2),
			terrain.NewTerrain(14, 14, 2),
			terrain.NewTerrain(15, 14, 2),
			terrain.NewTerrain(3, 11, 1),
			terrain.NewTerrain(4, 11, 1),
			terrain.NewTerrain(5, 11, 1),
			terrain.NewTerrain(6, 10, 1),
			terrain.NewTerrain(7, 10, 1),
		},
	}
}
func (l *Level) Update() {
	//遍历物品，销毁已经被销毁的物品
	for i := 0; i < len(l.Terrain); i++ {
		if l.Terrain[i].Destroyed {
			l.Terrain = append(l.Terrain[:i], l.Terrain[i+1:]...)
			i--
		}
	}
	physics.CheckPlayerEnemyCollision(l.Mario, l.Enemies)
	// physics.CheckPlayerTerrainCollision(l.Mario, l.Item)
	physics.CheckPlayerTerrainCollision(l.Mario, l.Terrain)
	for _, terrain := range l.Terrain {
		terrain.Update()
	}
	l.Mario.Update()

}
func (l *Level) Draw(screen *ebiten.Image) {
	// for _, enemy := range l.Enemies {
	// 	enemy.Draw(screen)
	// }
	// for _, item := range l.Item {
	// 	item.Draw(screen)
	// }
	for _, terrain := range l.Terrain {
		terrain.Draw(screen)
	}
	l.Mario.Draw(screen)
}
