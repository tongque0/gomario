package level

import (
	"encoding/json"
	"gomario/assets"
	"gomario/pkg/enemies"
	"gomario/pkg/item"
	"gomario/pkg/mario"
	"gomario/pkg/physics"
	"gomario/pkg/terrain"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type LevelConfig struct {
	Mario   MarioConfig           `json:"mario"`
	Terrain map[string][]Position `json:"terrain"`
	Enemies map[string][]Position `json:"enemies"`
	Items   map[string][]Position `json:"items"`
}
type MarioConfig struct {
	StartX int `json:"startX"`
	StartY int `json:"startY"`
}
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Level struct {
	Mario   *mario.Mario
	Enemies []*enemies.Enemies
	Item    []*item.Item
	Terrain []*terrain.Terrain
}

// 根据配置文件，创建关卡
// level: 关卡编号
func NewLevel(level int) *Level {
	file, err := assets.Assets.Open("level/level" + strconv.Itoa(level) + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var config LevelConfig
	if err := json.Unmarshal(byteValue, &config); err != nil {
		log.Fatal("Error parsing level config:", err)
	}

	lvl := &Level{
		Mario: mario.NewMario(config.Mario.StartX, config.Mario.StartY),
	}

	for typeStr, positions := range config.Terrain {
		terrainType, _ := strconv.Atoi(typeStr) // 转换类型为int
		for _, pos := range positions {
			lvl.Terrain = append(lvl.Terrain, terrain.NewTerrain(pos.X, pos.Y, terrainType))
		}
	}

	// // 初始化敌人
	// for enemyType, positions := range config.Enemies {
	// 	for _, pos := range positions {
	// 		enemy := enemies.NewEnemy(enemyType)
	// 		enemy.SetPosition(pos.X, pos.Y)
	// 		lvl.Enemies = append(lvl.Enemies, enemy)
	// 	}
	// }

	// // 初始化物品
	// for itemType, positions := range config.Items {
	// 	for _, pos := range positions {
	// 		item := item.NewItem(itemType)
	// 		item.SetPosition(pos.X, pos.Y)
	// 		lvl.Items = append(lvl.Items, item)
	// 	}
	// }
	return lvl
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
