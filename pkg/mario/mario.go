package mario

import (
	"gomario/assets"
	"gomario/internal/camera.go"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mario struct {
	X, Y                 float64        // 位置
	VelocityX, VelocityY float64        // 速度
	ScaleX, ScaleY       float64        // 缩放比例
	Width, Height        float64        // 尺寸
	State                int            // 当前状态（例如：站立、奔跑、跳跃等）
	IsBig                bool           // 是否是大马里奥
	IsDead               bool           // 是否死亡
	IsJumping            bool           // 是否在跳跃
	IsFalling            bool           // 是否在下落
	IsDucking            bool           // 是否在下蹲
	IsInvincible         bool           // 是否无敌
	Direction            int            // 方向（例如：左、右）
	Score                int            // 分数
	Lives                int            // 生命数
	Coins                int            // 收集的金币数
	PowerUp              string         // 当前的能力（例如：火焰、星星等）
	AnimationFrame       int            // 当前动画帧
	AnimationTimer       time.Time      // 动画计时器
	Graphics             *MarioGraphics // 马里奥的图像资源
}

const (
	StateIdle = iota
	StateRunning
	StateJumping
	StateFalling
	StateDucking
	StateDead
)

const (
	DirectionLeft = iota
	DirectionRight
)

var _ camera.Drawable = (*Mario)(nil)

func NewMario(gridX, gridY int) *Mario {
	marioGraphics := NewMarioGraphics()

	// 计算缩放比例
	scaleX := assets.CellSize / float64(marioGraphics.SmallWidth)
	scaleY := assets.CellSize / float64(marioGraphics.SmallHeight)

	mario := Mario{
		X:         float64(gridX) * assets.CellSize, // 使用网格坐标计算实际坐标
		Y:         float64(gridY) * assets.CellSize, // 使用网格坐标计算实际坐标
		State:     StateFalling,
		Width:     float64(marioGraphics.BigWidth) * scaleX,
		Height:    float64(marioGraphics.BigHeight) * scaleY,
		Direction: DirectionLeft,
		IsBig:     true,
		Graphics:  marioGraphics,
		ScaleX:    scaleX,
		ScaleY:    scaleY,
	}
	return &mario
}

// GetPosition 返回马里奥的世界坐标
func (m *Mario) GetPosition() (float64, float64) {
	return m.X, m.Y
}

// GetSize 返回马里奥的尺寸
func (m *Mario) GetSize() (float64, float64) {
	return m.Width, m.Height
}
func (m *Mario) Update() {
	// 更新马里奥的位置
	m.handleInput()
	m.X += m.VelocityX
	m.Y += m.VelocityY
	// 应用重力
	m.VelocityY += 0.5 // 重力加速度

	// 更新马里奥的状态
	switch m.State {
	case StateRunning:
		if time.Since(m.AnimationTimer) > 100*time.Millisecond { // 每100毫秒更新一次动画帧
			m.AnimationFrame++
			if m.AnimationFrame >= len(m.Graphics.SmallWalkingRightImages) {
				m.AnimationFrame = 0
			}
			m.AnimationTimer = time.Now()
		}
	case StateJumping:
		if m.VelocityY > 0 {
			m.State = StateFalling
		}
	case StateFalling:
		//落地后状态变为Idle
		if m.VelocityY == 0 {
			m.State = StateIdle
		}
	case StateDucking:

	case StateDead:
		// 处理死亡逻辑
	}
}
func (m *Mario) Draw(screen *ebiten.Image, camera *camera.Camera) {
	op := &ebiten.DrawImageOptions{}

	// 计算屏幕坐标
	screenX := m.X - camera.X
	screenY := m.Y - camera.Y

	// 绘制偏移量，处理一些特殊状态下的绘制位置偏移，例如下蹲
	drawOffsetY := float64(0)
	if m.State == StateDucking && m.IsBig {
		// 下蹲时，将绘制位置向上偏移半个身位
		drawOffsetY = -(float64(m.Graphics.BigHeight) - float64(m.Graphics.SmallHeight)) * m.ScaleY * 0.75
	}

	// 根据方向设置缩放和位置
	if m.Direction == DirectionLeft {
		op.GeoM.Scale(-m.ScaleX, m.ScaleY)
		op.GeoM.Translate(screenX+m.Width, screenY+drawOffsetY)
	} else {
		op.GeoM.Scale(m.ScaleX, m.ScaleY)
		op.GeoM.Translate(screenX, screenY+drawOffsetY)
	}

	// 选择当前帧
	var img *ebiten.Image
	if m.IsBig {
		switch m.State {
		case StateIdle:
			img = m.Graphics.BigIdleRightImages
		case StateRunning:
			if m.IsJumping || m.IsFalling {
				img = m.Graphics.BigJumpingRightImages
			} else {
				img = m.Graphics.BigWalkingRightImages[m.AnimationFrame]
			}
		case StateJumping:
			img = m.Graphics.BigJumpingRightImages
		case StateFalling:
			img = m.Graphics.BigJumpingRightImages
		case StateDucking:
			img = m.Graphics.BigDuckingRightImages
		}
	} else {
		switch m.State {
		case StateIdle:
			img = m.Graphics.SmallIdleRightImages
		case StateRunning:
			if m.IsJumping || m.IsFalling {
				img = m.Graphics.SmallJumpingRightImages
			} else {
				img = m.Graphics.SmallWalkingRightImages[m.AnimationFrame]
			}
		case StateJumping:
			img = m.Graphics.SmallJumpingRightImages
		case StateFalling:
			img = m.Graphics.SmallJumpingRightImages
		case StateDucking:
			img = m.Graphics.SmallSkiddingRightImages
		case StateDead:
			img = m.Graphics.SmallDeathRightImages
		}
	}

	// 绘制图像
	if img != nil {
		screen.DrawImage(img, op)
	}

	// 调试模式下绘制碰撞框
	if assets.IsDebug {
		blue := color.RGBA{0, 0, 255, 255}
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				screen.Set(int(screenX)+dx, int(screenY)+dy, blue)                  // Top-left
				screen.Set(int(screenX+m.Width)+dx, int(screenY)+dy, blue)          // Top-right
				screen.Set(int(screenX)+dx, int(screenY+m.Height)+dy, blue)         // Bottom-left
				screen.Set(int(screenX+m.Width)+dx, int(screenY+m.Height)+dy, blue) // Bottom-right
			}
		}
	}
}

// dirction: 1-top, 2-bottom, 3-left, 4-right
// 待重构：将碰撞检测逻辑移动到碰撞检测模块
func (m *Mario) OnTerrainCollision(direction int, terrainX, terrainY, terrainWidth, terrainHeight float64) {
	// 处理马里奥与地形的碰撞
	//上下碰撞,Y轴速度清零;左右碰撞,X轴速度清零;
	switch direction {
	case 1: // Top
		// fmt.Println("top")
		m.VelocityY = 0
		m.Y = terrainY + terrainHeight
		m.State = StateFalling
		m.IsJumping = false
		m.IsFalling = true
	case 2: // Bottom
		// fmt.Println("bottom")
		m.VelocityY = 0
		m.Y = terrainY - m.Height
		m.IsJumping = false
		m.IsFalling = false
		m.State = StateIdle
	case 3: // Left
		// fmt.Println("left")
		m.VelocityX = 0
		m.X = terrainX + terrainWidth
	case 4: // Right
		// fmt.Println("right")
		m.VelocityX = 0
		m.X = terrainX - m.Width
	default:
		return
	}
}
