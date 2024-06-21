package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type sceneEffectType int

const (
	Immediately sceneEffectType = iota
	FadeIn
	FadeOut
	CircularClosing
	CircularOpening
	CircularFocusClosing
  CircularFocusOpening
)

type TransitionOptions struct {
	OutSceneEffect SceneEffect
	InSceneEffect  SceneEffect
}

type SceneEffect struct {
	Type  sceneEffectType
	Focus Focus
	Clr   color.RGBA
	Tick  float64
	Frame float64
}

type Focus struct {
	X int
	Y int
}

var transitionHandlerMap = map[sceneEffectType]transitionInterface{
	Immediately:          &immediatelyEffect{},
	FadeIn:               &fadeInEffect{},
	FadeOut:              &fadeOutEffect{},
	CircularClosing:      &circularClosingEffect{},
	CircularOpening:      &circularOpeningEffect{},
	CircularFocusClosing: &circularFocusClosingEffect{},
  CircularFocusOpening: &circularFocusOpeningEffect{},
}

func getOutTransitionHandler() transitionInterface {
	return transitionHandlerMap[outSceneEffect.Type]
}
func getInTransitionHandler() transitionInterface {
	return transitionHandlerMap[inSceneEffect.Type]
}

type transitionInterface interface {
	Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image)
}

type immediatelyEffect struct{}

func (ie *immediatelyEffect) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	img := ebiten.NewImage(w, h)
	clr := sceneEffect.Clr
	img.Fill(clr)
	screen.DrawImage(img, nil)
}


