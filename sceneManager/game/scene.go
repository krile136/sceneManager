package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	current            string
	next               string
	isTransitionFinish bool
	tick               int
	transitionFrame    int
	outTransitionType  TransitionType
	inTransitionType   TransitionType
	clr                color.RGBA
)

type sceneInterface interface {
	Init() error
	Update() error
	Draw(screen *ebiten.Image)
}

// 現在のシーンを最初からやり直します
func Reload() {
	current = ""
}

func Change(scenaName string, op *TransitionOptions) {
	if op == nil {
		op = &TransitionOptions{
			OutTransitionType: Immediately,
			InTransitionType:  Immediately,
			TransitionFrame:   0,
			Clr:               color.RGBA{R: 0, G: 0, B: 0, A: 255},
		}
	}

	next = scenaName
	isTransitionFinish = false
	tick = 0
	transitionFrame = op.TransitionFrame
	outTransitionType = op.OutTransitionType
	inTransitionType = op.InTransitionType
	clr = op.Clr
}
