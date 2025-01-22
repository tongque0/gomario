package mario

import (
	"gomario/assets"

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
	IsInvincible         bool           // 是否无敌
	Direction            int            // 方向（例如：左、右）
	Score                int            // 分数
	Lives                int            // 生命数
	Coins                int            // 收集的金币数
	PowerUp              string         // 当前的能力（例如：火焰、星星等）
	AnimationFrame       int            // 当前动画帧
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

func NewMario(gridX, gridY int) *Mario {
	marioGraphics := NewMarioGraphics()

	// 计算缩放比例
	scaleX := assets.CellSize / float64(marioGraphics.SmallWidth)
	scaleY := assets.CellSize / float64(marioGraphics.SmallHeight)

	mario := Mario{
		X:         float64(gridX) * assets.CellSize, // 使用网格坐标计算实际坐标
		Y:         float64(gridY) * assets.CellSize, // 使用网格坐标计算实际坐标
		State:     StateFalling,
		Width:     float64(marioGraphics.SmallWidth) * scaleX,
		Height:    float64(marioGraphics.SmallHeight) * scaleY,
		Direction: DirectionLeft,
		IsBig:     false,
		Graphics:  marioGraphics,
		ScaleX:    scaleX,
		ScaleY:    scaleY,
	}
	return &mario
}
func (m *Mario) Update() {
	// 更新马里奥的位置
	m.X += m.VelocityX
	m.Y += m.VelocityY
	// fmt.Println(m.State)
	// 更新马里奥的状态
	switch m.State {
	case StateRunning:
		m.AnimationFrame++
		if m.AnimationFrame >= len(m.Graphics.SmallWalkingRightImages) {
			m.AnimationFrame = 0
		}
	case StateJumping:
		// 处理跳跃逻辑
	case StateFalling:
		m.VelocityY = 10
	case StateDead:
		// 处理死亡逻辑
	}
}
func (m *Mario) Draw(screen *ebiten.Image) {
	// 根据马里奥的状态和方向绘制马里奥

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(m.ScaleX, m.ScaleY)
	op.GeoM.Translate(m.X, m.Y)
	var img *ebiten.Image
	if m.IsBig {
		switch m.State {
		case StateIdle:
			img = m.Graphics.BigIdleRightImages
		case StateRunning:
			img = m.Graphics.BigWalkingRightImages[m.AnimationFrame]
		case StateJumping:
			img = m.Graphics.BigJumpingRightImages
		case StateFalling:
			img = m.Graphics.BigJumpingRightImages
		case StateDucking:
			img = m.Graphics.BigSkiddingRightImages
		}
	} else {
		switch m.State {
		case StateIdle:
			img = m.Graphics.SmallIdleRightImages
		case StateRunning:
			img = m.Graphics.SmallWalkingRightImages[m.AnimationFrame]
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

	if m.Direction == DirectionLeft {
		op.GeoM.Scale(-1, 1)          // 镜像水平翻转
		op.GeoM.Translate(m.Width, 0) // 调整位置
	}

	if img != nil {
		screen.DrawImage(img, op)
	}

}

// dirction: 1-top, 2-bottom, 3-left, 4-right
func (m *Mario) OnTerrainCollision(direction int) {
	// 处理马里奥与地形的碰撞
	//上下碰撞,Y轴速度清零;左右碰撞,X轴速度清零;
	switch direction {
	case 1:
		m.VelocityY = 0
		m.State = StateFalling
		m.IsJumping = false
		m.IsFalling = true
	case 2:
		m.VelocityY = 0
		m.IsJumping = false
		m.IsFalling = false
		m.State = StateIdle
	case 3:
		m.VelocityX = 0
		m.X -= m.Width
	case 4:
		m.VelocityX = 0
		m.X += m.Width
	default:
		return
	}
}
