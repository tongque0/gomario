package level

import (
	"encoding/json"
	"gomario/assets"
	"gomario/internal/camera.go"
	destroyeffect "gomario/pkg/destroyEffect"
	"gomario/pkg/enemies"
	"gomario/pkg/item"
	"gomario/pkg/mario"
	"gomario/pkg/physics"
	"gomario/pkg/terrain"
	"io"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type LevelConfig struct {
	MapConfig MapConfig             `json:"mapconfig"`
	Mario     MarioConfig           `json:"mario"`
	Terrain   map[string][]Position `json:"terrain"`
	Enemies   map[string][]Position `json:"enemies"`
	Items     map[string][]Position `json:"items"`
}
type MapConfig struct {
	Width  int `json:"width"`
	Height int `json:"height"`
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
	Camera        *camera.Camera
	Mario         *mario.Mario
	Enemies       []*enemies.Enemies
	Item          []*item.Item
	Terrain       []*terrain.Terrain
	destroyeffect *destroyeffect.DestroyEffect
}

// 根据配置文件，创建关卡
// level: 关卡编号
func NewLevel(level int) *Level {
	file, err := assets.Assets.Open("level/level" + strconv.Itoa(level) + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var config LevelConfig
	if err := json.Unmarshal(byteValue, &config); err != nil {
		log.Fatal("Error parsing level config:", err)
	}

	lvl := &Level{
		destroyeffect: destroyeffect.NewDestroyEffect(),
		Camera:        camera.NewCamera(assets.ScreenWidth, assets.ScreenHeight),
		Mario:         mario.NewMario(config.Mario.StartX, config.Mario.StartY),
	}
	lvl.Camera.SetBounds(config.MapConfig.Width, config.MapConfig.Height)
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
	l.destroyeffect.Update()
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
	l.Camera.Update(l.Mario.X, l.Mario.Y)

}
func (l *Level) Draw(screen *ebiten.Image) {
	l.destroyeffect.Draw(screen, l.Camera)
	drawables := make([]camera.Drawable, 0)
	// for _, enemy := range l.Enemies {
	// 	drawables = append(drawables, enemy)
	// }
	// for _, item := range l.Item {
	// 	drawables = append(drawables, item)
	// }
	for _, terrain := range l.Terrain {
		drawables = append(drawables, terrain)
	}

	drawables = append(drawables, l.Mario)

	// 绘制所有对象
	for _, d := range drawables {
		if isVisible(d, l.Camera) {
			d.Draw(screen, l.Camera)
		}
	}

}
func isVisible(d camera.Drawable, c *camera.Camera) bool {
	x, y := d.GetPosition()
	w, h := d.GetSize()

	left := x - c.X
	right := left + w
	top := y - c.Y
	bottom := top + h

	return right > 0 &&
		left < float64(c.ViewportWidth) &&
		bottom > 0 &&
		top < float64(c.ViewportHeight)
}
