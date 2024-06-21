package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	current            string
	next               string
	isTransitionFinish bool
	tick               float64
	outSceneEffect     SceneEffect
	inSceneEffect      SceneEffect
  customEffect       *SceneEffect
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
		outSceneEffect := &SceneEffect{
			Type:  Immediately,
			Clr:   color.RGBA{R: 0, G: 0, B: 0, A: 255},
			Tick:  0,
			Frame: 30,
		}
		inSceneEffect := outSceneEffect

		op = &TransitionOptions{
			OutSceneEffect: *outSceneEffect,
			InSceneEffect:  *inSceneEffect,
		}
	}

	next = scenaName
	isTransitionFinish = false
	tick = 0
	outSceneEffect = op.OutSceneEffect
	inSceneEffect = op.InSceneEffect
}

func ExecuteEffect(effect *SceneEffect){
  customEffect = effect
}
