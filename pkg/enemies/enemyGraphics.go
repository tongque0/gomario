package enemies

import "github.com/hajimehoshi/ebiten/v2"

type EnemiesGraphics struct {
	AliveImages   []*ebiten.Image // 存活状态的图像
	DeathImages   []*ebiten.Image // 死亡状态的图像
	SpecialImages []*ebiten.Image // 特殊状态的图像
	AttackImages  []*ebiten.Image // 攻击状态的图像
}
