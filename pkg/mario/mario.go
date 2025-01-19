package mario

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Mario struct {
	X, Y                 float64        // 位置
	VelocityX, VelocityY float64        // 速度
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

func NewMario() *Mario {
	return &Mario{
		X:        0,
		Y:        0,
		State:    0,
		Graphics: NewMarioGraphics(),
	}
}
func (m *Mario) Update() {
	// 更新马里奥的位置
	m.X += m.VelocityX
	m.Y += m.VelocityY

	// 更新马里奥的状态
	switch m.State {
	case StateRunning:
		// 处理奔跑逻辑
	case StateJumping:
		// 处理跳跃逻辑
	case StateFalling:
		// 处理下落逻辑
	case StateDead:
		// 处理死亡逻辑
	}
}
func (m *Mario) Draw(screen *ebiten.Image) {
	// 根据马里奥的状态和方向绘制马里奥

	// switch m.State {
	// case StateIdle:
	// 	img = idleImage
	// case StateRunning:
	// 	img = runningImages[m.AnimationFrame]
	// case StateJumping:
	// 	img = jumpingImage
	// case StateFalling:
	// 	img = fallingImage
	// case StateDead:
	// 	img = deadImage
	// }

}

func (m *Mario) OnTerrainCollision(direction int) {
	// 处理马里奥与地形的碰撞
	switch direction {
	case 1:
		m.VelocityY = 0
		m.Y -= m.Height
	case 2:
		m.VelocityY = 0
		m.Y += m.Height
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
