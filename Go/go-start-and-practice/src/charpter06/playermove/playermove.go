package main

import (
	"fmt"
	"math"
)

type Vec2 struct {
	X, Y float32
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}

	return Vec2{0, 0}
}

type Player struct {
	currPos   Vec2    // 当前位置
	targetPos Vec2    // 目标位置
	speed     float32 // 移动速度
}

// 设置玩家移动的位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// 设置当前位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// 判断是否到达目的地
func (p *Player) IsArried() bool {
	// 通过计算当前玩家位置与目标位置的距离不超过移动的步长，判断已经达到目标点
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

// 更新玩家的位置
func (p *Player) Update() {
	if !p.IsArried() {
		// 计算出当前位置指向目标位置的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()
		newPos := p.currPos.Add(dir.Scale(p.speed))
		p.currPos = newPos
	}
}

func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}

func main() {
	p := NewPlayer(0.5)
	p.MoveTo(Vec2{3, 1})

	for !p.IsArried() {
		p.Update()
		fmt.Println(p.Pos())
	}
}
