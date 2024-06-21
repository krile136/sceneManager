package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type fadeInEffect struct{}

func (fie *fadeInEffect) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	alpha := 255 - 255*(sceneEffect.Tick/sceneEffect.Frame)
	img := generateFadeImage(sceneEffect, w, h, alpha)
	screen.DrawImage(img, nil)
}

type fadeOutEffect struct{}

func (foe *fadeOutEffect) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	alpha := 255 * (sceneEffect.Tick / sceneEffect.Frame)
	img := generateFadeImage(sceneEffect, w, h, alpha)
	screen.DrawImage(img, nil)
}
