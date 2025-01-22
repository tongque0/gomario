package physics

import (
	"gomario/pkg/enemies"
	"gomario/pkg/item"
	"gomario/pkg/mario"
	"gomario/pkg/terrain"
)

type Rectangle struct {
	X, Y, Width, Height float64
}

type CollisionDirection int

const (
	None CollisionDirection = iota
	Top
	Bottom
	Left
	Right
)

// CheckCollision 检查两个矩形是否相交，并处理碰撞
func CheckCollision(rect1, rect2 Rectangle) CollisionDirection {
	// 判断是否发生碰撞
	if rect1.X < rect2.X+rect2.Width && rect1.X+rect1.Width > rect2.X &&
		rect1.Y < rect2.Y+rect2.Height && rect1.Y+rect1.Height > rect2.Y {

		// 确定碰撞方向
		overlapTop := rect2.Y + rect2.Height - rect1.Y
		overlapBottom := rect1.Y + rect1.Height - rect2.Y
		overlapLeft := rect2.X + rect2.Width - rect1.X
		overlapRight := rect1.X + rect1.Width - rect2.X

		// 底部碰撞
		if overlapBottom < overlapTop && overlapBottom < overlapLeft && overlapBottom < overlapRight {
			return Bottom
		}
		// 顶部碰撞
		if overlapTop < overlapBottom && overlapTop < overlapLeft && overlapTop < overlapRight {
			return Top
		}
		// 右侧碰撞
		if overlapRight < overlapLeft && overlapRight < overlapTop && overlapRight < overlapBottom {
			return Right
		}
		// 左侧碰撞
		if overlapLeft < overlapRight && overlapLeft < overlapTop && overlapLeft < overlapBottom {
			return Left
		}
	}
	return None
}

// CheckPlayerTerrainCollision 检查马里奥与地形的碰撞，并处理碰撞
func CheckPlayerTerrainCollision(player *mario.Mario, terrain []*terrain.Terrain) {
	playerRect := Rectangle{player.X, player.Y, player.Width, player.Height}
	for _, t := range terrain {
		terrainRect := Rectangle{t.X, t.Y, t.Width, t.Height}
		if direction := CheckCollision(playerRect, terrainRect); direction != None {
			t.OnMarioCollision(int(direction), player.IsBig, player.IsJumping)
			player.OnTerrainCollision(int(direction), t.X, t.Y, t.Width, t.Height)
		}
	}
}

// CheckPlayerEnemyCollision 检查马里奥与敌人的碰撞，并处理碰撞
func CheckPlayerEnemyCollision(player *mario.Mario, enemies []*enemies.Enemies) {
	playerRect := Rectangle{player.X, player.Y, player.Width, player.Height}
	for _, enemy := range enemies {
		enemyRect := Rectangle{enemy.X, enemy.Y, enemy.Width, enemy.Height}
		if direction := CheckCollision(playerRect, enemyRect); direction != None {
			// player.OnEnemyCollision(enemy, direction)
		}
	}
}

// CheckPlayerItemCollision 检查马里奥与物品的碰撞，并处理碰撞
func CheckPlayerItemCollision(player *mario.Mario, items []*item.Item) {
	playerRect := Rectangle{player.X, player.Y, player.Width, player.Height}
	for _, item := range items {
		itemRect := Rectangle{item.X, item.Y, item.Width, item.Height}
		if direction := CheckCollision(playerRect, itemRect); direction != None {
			// player.OnItemCollision(item)
		}
	}
}
