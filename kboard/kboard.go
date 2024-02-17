package kboard

import (
	"github.com/go-vgo/robotgo"
)

func New() KBoard { return KBoard{} }

type KBoard struct{}

func (m KBoard) Type(text string) {
	robotgo.TypeStr(text)
}

func (m KBoard) KeyPress(key string) {
	robotgo.KeyPress(key)
}

func (m KBoard) KeyTap(key string, args ...interface{}) {
	robotgo.KeyTap(key, args...)
}

func (m KBoard) CtrlA() {
	robotgo.KeyTap("a", "ctrl")
}

func (m KBoard) PressEnter() {
	m.KeyPress("enter")
}
