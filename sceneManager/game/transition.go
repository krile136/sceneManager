package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/sceneManager/sceneManager/game/effectType"
	"github.com/krile136/sceneManager/sceneManager/game/effects"
)

type TransitionOptions struct {
	OutSceneEffect effects.SceneEffect
	InSceneEffect  effects.SceneEffect
}

var transitionHandlerMap = map[effectType.SceneEffectType]transitionInterface{
	effectType.Immediately:          &effects.Immediately{},
	effectType.FadeIn:               &effects.FadeIn{},
	effectType.FadeOut:              &effects.FadeOut{},
	effectType.CircularClosing:      &effects.CircularClosing{},
	effectType.CircularOpening:      &effects.CircularOpening{},
	effectType.CircularFocusClosing: &effects.CircularFocusClosing{},
	effectType.CircularFocusOpening: &effects.CircularFocusOpening{},
}

func getOutTransitionHandler() transitionInterface {
	return transitionHandlerMap[outSceneEffect.Type]
}
func getInTransitionHandler() transitionInterface {
	return transitionHandlerMap[inSceneEffect.Type]
}

type transitionInterface interface {
	Draw(sceneEffect effects.SceneEffect, w, h int, screen *ebiten.Image)
}
