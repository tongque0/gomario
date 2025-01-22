// Description: 用于嵌入资源文件，如图片、音频等, 以及定义一些常量
// Created by tongque on 25-1-21
package assets

import (
	"embed"
)

//go:embed **/*
var Assets embed.FS

const (
	ScreenWidth, ScreenHeight = 512, 480                             // 屏幕尺寸,比例最好为4:3
	CellSize                  = min(ScreenWidth/16, ScreenHeight/15) // 图像单元的尺寸，根据屏幕尺寸计算
	IsDebug                   = true                                 // 是否开启调试模式
)

//游戏单元格说明
//将屏幕划分为单元格，方便进行坐标的定位，从左上角开始，横向为x轴，纵向为y轴
//摄像头包含的单元格数为16*15，即屏幕尺寸的大小
//如果用坐标数组表示，即
//[
// [0,0],[1,0],[2,0]...[15,0]
// [0,1],[1,1],[2,1]...[15,1]
// ...
// [0,14],[1,14],[2,14]...[15,14]
// ]
//其中[0,0]为左上角，[15,14]为右下角
