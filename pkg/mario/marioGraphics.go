package mario

import (
	"gomario/assets"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MarioGraphics struct {
	// 小马里奥的图像资源(右)
	SmallIdleRightImages     *ebiten.Image   // 站立
	SmallWalkingRightImages  []*ebiten.Image // 奔跑
	SmallJumpingRightImages  *ebiten.Image   // 跳跃
	SmallSkiddingRightImages *ebiten.Image   // 滑行
	SmallDeathRightImages    *ebiten.Image   // 死亡
	// 大马里奥的图像资源(右)
	BigIdleRightImages     *ebiten.Image   // 站立
	BigWalkingRightImages  []*ebiten.Image // 奔跑
	BigJumpingRightImages  *ebiten.Image   // 跳跃
	BigSkiddingRightImages *ebiten.Image   // 滑行
	BigDuckingRightImages  *ebiten.Image   // 蹲下
	//图形资源的宽高
	SmallWidth, SmallHeight int
	BigWidth, BigHeight     int
}

func NewMarioGraphics() *MarioGraphics {
	mg := &MarioGraphics{}
	mg.loadResources()
	return mg
}

// 加载mario精灵图
func (mg *MarioGraphics) loadResources() {

	spriteSheet, _, err := ebitenutil.NewImageFromFileSystem(assets.Assets, "graphics/mario_bros.png")
	if err != nil {
		log.Fatal(err)
	}
	//加载小马里奥的图像资源(右)
	mg.SmallWidth, mg.SmallHeight = 16, 16
	mg.SmallIdleRightImages = loadSubImage(spriteSheet, 178, 32, 12, 16)
	mg.SmallWalkingRightImages = append(mg.SmallWalkingRightImages, loadSubImage(spriteSheet, 80, 32, 15, 16))
	mg.SmallWalkingRightImages = append(mg.SmallWalkingRightImages, loadSubImage(spriteSheet, 96, 32, 16, 16))
	mg.SmallWalkingRightImages = append(mg.SmallWalkingRightImages, loadSubImage(spriteSheet, 112, 32, 16, 16))
	mg.SmallJumpingRightImages = loadSubImage(spriteSheet, 144, 32, 16, 16)
	mg.SmallSkiddingRightImages = loadSubImage(spriteSheet, 130, 32, 14, 16)
	mg.SmallDeathRightImages = loadSubImage(spriteSheet, 160, 32, 15, 16)

	//加载大马里奥的图像资源(右)
	mg.BigWidth, mg.BigHeight = 16, 32
	mg.BigIdleRightImages = loadSubImage(spriteSheet, 176, 0, 16, 32)
	mg.BigWalkingRightImages = append(mg.BigWalkingRightImages, loadSubImage(spriteSheet, 81, 0, 16, 32))
	mg.BigWalkingRightImages = append(mg.BigWalkingRightImages, loadSubImage(spriteSheet, 97, 0, 15, 32))
	mg.BigWalkingRightImages = append(mg.BigWalkingRightImages, loadSubImage(spriteSheet, 113, 0, 15, 32))
	mg.BigJumpingRightImages = loadSubImage(spriteSheet, 144, 0, 16, 32)
	mg.BigSkiddingRightImages = loadSubImage(spriteSheet, 128, 0, 16, 32)
	mg.BigDuckingRightImages = loadSubImage(spriteSheet, 160, 0, 16, 32)

}

// 从精灵图中切割出指定位置和尺寸的子图像
func loadSubImage(spriteSheet *ebiten.Image, x, y, width, height int) *ebiten.Image {
	subImage := spriteSheet.SubImage(image.Rect(x, y, x+width, y+height)).(*ebiten.Image)
	return subImage
}
