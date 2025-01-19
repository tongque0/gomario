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
		Mario: mario.NewMario(),
		// Enemies: []*enemies.Enemies{
		// 	enemies.NewEnemies(),
		// },
		// Item: []*item.Item{
		// 	item.NewItem(),
		// },
		Terrain: []*terrain.Terrain{
			terrain.NewTerrain(0, 0, 0),
			// terrain.NewTerrain(16, 16, 0),
			// terrain.NewTerrain(32, 32, 0),
			// terrain.NewTerrain(48, 48, 0),
			// terrain.NewTerrain(64, 64, 0),
		},
	}
}
func (l *Level) Update() {
	physics.CheckPlayerEnemyCollision(l.Mario, l.Enemies)
	// physics.CheckPlayerTerrainCollision(l.Mario, l.Item)

}
func (l *Level) Draw(screen *ebiten.Image) {
	l.Mario.Draw(screen)
	// for _, enemy := range l.Enemies {
	// 	enemy.Draw(screen)
	// }
	// for _, item := range l.Item {
	// 	item.Draw(screen)
	// }
	for _, terrain := range l.Terrain {
		terrain.Draw(screen)
	}
}
