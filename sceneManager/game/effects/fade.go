package effects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type FadeIn struct{}

func (fie *FadeIn) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	alpha := 255 - 255*(sceneEffect.Tick/sceneEffect.Frame)
	img := generateFadeImage(sceneEffect, w, h, alpha)
	screen.DrawImage(img, nil)
}

type FadeOut struct{}

func (foe *FadeOut) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	alpha := 255 * (sceneEffect.Tick / sceneEffect.Frame)
	img := generateFadeImage(sceneEffect, w, h, alpha)
	screen.DrawImage(img, nil)
}

func generateFadeImage(se SceneEffect, w, h int, alpha float64) (img *ebiten.Image) {
	img = ebiten.NewImage(w, h)
	clr := se.Clr
	clr.A = uint8(alpha)
	img.Fill(clr)
	return
}
