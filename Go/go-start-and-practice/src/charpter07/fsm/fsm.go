package main

import (
	"errors"
	"fmt"
	"reflect"
)

// fsm，Finite-State Machine, 有限状态机，表示有限个状态以及
// 在这些状态之间转移和动作等行为的数字模型

type State interface {
	Name() string // 同类型状态转换
	EnableSameTransit() bool
	OnBegin() // 响应状态开始
	OnEnd()   // 响应状态结束

	CanTransitTo(name string) bool
}

func StateName(s State) string {
	if s == nil {
		return "none"
	}
	// 使用反射获取状态名称
	return reflect.TypeOf(s).Elem().Name()
}

type StateInfo struct {
	name string
}

// 通过 StateInfo 实现 State 接口
func (s *StateInfo) Name() string {
	return s.name
}

func (s *StateInfo) setName(name string) {
	s.name = name
}

func (s *StateInfo) EnableSameTransit() bool {
	return false
}

func (s *StateInfo) OnBegin() {

}

func (s *StateInfo) OnEnd() {

}

func (s *StateInfo) CanTransitTo(name string) bool {
	return true
}

// 状态管理器
type StateManager struct {
	stateByName map[string]State
	OnChange    func(from State, to State)
	curr        State
}

func (sm *StateManager) Add(s State) {
	name := StateName(s)
	// 将 s 转换为能设置名字的接口，然后调用该接口
	// State 有很多方法，包括 SetName, 此处好比转换成了 State 的父类
	// 此父类包含 SetName 方法
	s.(interface {
		setName(name string)
	}).setName(name)

	if sm.Get(name) != nil {
		panic("duplicate state:" + name)
	}
	sm.stateByName[name] = s
}

func (sm *StateManager) Get(name string) State {
	if v, ok := sm.stateByName[name]; ok {
		return v
	}
	return nil
}

func NewStateManager() *StateManager {
	return &StateManager{
		stateByName: make(map[string]State),
	}
}

var ErrStateNotFound = errors.New("state not found")
var ErrForbidSameStateTransit = errors.New("forbid same transit")
var ErrCannotTransitToState = errors.New("cannot transit to state")

func (sm *StateManager) CurrState() State {
	return sm.curr
}

func (sm *StateManager) CanCurrTransitTo(name string) bool {
	if sm.curr == nil {
		return true
	}

	// 状态相同
	if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
		return false
	}

	return sm.curr.CanTransitTo(name)
}

func (sm *StateManager) Transit(name string) error {
	next := sm.Get(name)
	if next == nil {
		return ErrStateNotFound
	}

	pre := sm.curr
	if sm.curr != nil {
		if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
			return ErrForbidSameStateTransit
		}

		if !sm.curr.CanTransitTo(name) {
			return ErrCannotTransitToState
		}

		sm.curr.OnEnd()
	}

	sm.curr = next
	sm.curr.OnBegin()

	if sm.OnChange != nil {
		sm.OnChange(pre, sm.curr)
	}

	return nil
}

type IdleState struct {
	StateInfo
}

func (i *IdleState) OnBegin() {
	fmt.Println("idle State begin")
}

func (i *IdleState) OnEnd() {
	fmt.Println("idle State end")
}

type MoveState struct {
	StateInfo
}

func (m *MoveState) OnBegin() {
	fmt.Println("Move State begin")
}

func (m *MoveState) EnableSameTransit() bool {
	return true
}

type JumpState struct {
	StateInfo
}

func (j *JumpState) OnBegin() {
	fmt.Println("Jump State begin")
}

func (j *JumpState) CanTransitTo(name string) bool {
	return name != "MoveState"
}

func transitAndReport(sm *StateManager, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("FAILED! %s --> %s, %s\n\n", sm.CurrState().Name(), target, err.Error())
	}
}

func main() {
	sm := NewStateManager()
	sm.OnChange = func(from State, to State) {
		fmt.Printf("%s-----> %s\n\n", StateName(from), StateName(to))
	}

	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))

	transitAndReport(sm, "IdleState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "IdleState")
}
