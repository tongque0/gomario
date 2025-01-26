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
	DynamicItem   *item.DynamicItem // 动态物品
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
		DynamicItem:   item.NewDynamicItem(),
	}
	// 设置地图边界
	lvl.Camera.SetBounds(config.MapConfig.Width, config.MapConfig.Height)

	// 初始化地形
	for typeStr, positions := range config.Terrain {
		terrainType, _ := strconv.Atoi(typeStr) // 转换类型为int
		for _, pos := range positions {
			lvl.Terrain = append(lvl.Terrain, terrain.NewTerrain(pos.X, pos.Y, terrainType))
		}
	}

	// 初始化敌人
	for enemyType, positions := range config.Enemies {
		for _, pos := range positions {
			lvl.Enemies = append(lvl.Enemies, enemies.NewEnemies(pos.X, pos.Y, enemyType))
		}
	}

	// 初始化物品
	for itemType, positions := range config.Items {
		for _, pos := range positions {
			lvl.Item = append(lvl.Item, item.NewItem(pos.X, pos.Y, itemType))
		}
	}
	return lvl
}
func (l *Level) Update() {
	//销毁特效
	l.destroyeffect.Update()

	//碰撞系统
	physics.CheckPlayerTerrainCollision(l.Mario, l.Terrain)
	physics.CheckPlayerEnemyCollision(l.Mario, l.Enemies)
	physics.CheckPlayerItemCollision(l.Mario, l.Item)

	// 更新所有对象
	l.DynamicItem.Update()
	for i := 0; i < len(l.Terrain); i++ {
		if l.Terrain[i].Destroyed {
			l.Terrain = append(l.Terrain[:i], l.Terrain[i+1:]...)
			i--
		} else {
			l.Terrain[i].Update()
		}
	}
	for _, enemy := range l.Enemies {
		enemy.Update()
	}
	for i := 0; i < len(l.Item); i++ {
		if l.Item[i].IsDestroy {
			l.Item = append(l.Item[:i], l.Item[i+1:]...)
			i--
		} else {
			l.Item[i].Update()
		}
	}
	l.Mario.Update()
	l.Camera.Update(l.Mario.X, l.Mario.Y)

}
func (l *Level) Draw(screen *ebiten.Image) {
	l.destroyeffect.Draw(screen, l.Camera)
	l.DynamicItem.Draw(screen, l.Camera)
	drawables := make([]camera.Drawable, 0)

	for _, enemy := range l.Enemies {
		drawables = append(drawables, enemy)
	}

	for _, item := range l.Item {
		drawables = append(drawables, item)
	}

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

// isVisible 检查对象是否在摄像头视野内
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
