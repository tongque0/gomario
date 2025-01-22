package mario

import (
	"gomario/pkg/sound"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (m *Mario) handleInput() {
	// 跳跃
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyK) {
		if !m.IsJumping && !m.IsFalling {
			m.VelocityY = -10 // 设置向上的速度
			m.IsJumping = true
			m.State = StateJumping
		}
		if m.IsBig {
			sound.NewSfxPlayer("big_jump")
		}else {
			sound.NewSfxPlayer("small_jump")
		}
	}
	// 左移
	if inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
		m.State = StateRunning
		m.Direction = DirectionLeft
		m.VelocityX = -2
	}
	// 右移
	if inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
		m.State = StateRunning
		m.Direction = DirectionRight
		m.VelocityX = 2
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		m.VelocityX = 0
	}
	// 下蹲
	if inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
		if m.IsBig {
			m.IsDucking = true
			m.State = StateDucking
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		bottomY := m.Y + m.Height
		// 更新碰撞箱为小马里奥高度
		m.Height = float64(m.Graphics.SmallHeight) * m.ScaleY
		// 调整位置，保持底部Y不变
		m.Y = bottomY - m.Height
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyS) {
		m.IsDucking = false
		bottomY := m.Y + m.Height
		// 恢复为大马里奥高度
		m.Height = float64(m.Graphics.BigHeight) * m.ScaleY
		// 调整位置，保持底部Y不变
		m.Y = bottomY - m.Height
	}
	//停止移动
	if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) {
		m.VelocityX = 0
		if m.State == StateRunning {
			m.State = StateIdle
		}
	}
}
