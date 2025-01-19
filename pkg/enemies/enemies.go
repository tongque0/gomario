package enemies

type Enemies struct {
	X, Y                 float64          // 位置
	VelocityX, VelocityY float64          // 速度
	Width, Height        float64          // 尺寸
	Kind                 int              // 种类
	
	IsDead               bool             // 是否死亡
	Graphics             *EnemiesGraphics // 敌人的图像资源
}
