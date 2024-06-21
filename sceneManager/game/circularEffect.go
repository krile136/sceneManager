package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type circularClosingEffect struct{}

func (cce *circularClosingEffect) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	sceneEffect.Focus = Focus{X: w / 2, Y: h / 2}
	r := math.Sqrt(math.Pow(float64(w)/2, 2) + math.Pow(float64(h)/2, 2))
	d := (r - r*(float64(sceneEffect.Tick)/float64(sceneEffect.Frame))) * 2

	circleImage := generateCircleImage(sceneEffect, w, h, d)
	screen.DrawImage(circleImage, nil)
}

type circularOpeningEffect struct{}

func (coe *circularOpeningEffect) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	sceneEffect.Focus = Focus{X: w / 2, Y: h / 2}
	r := math.Sqrt(math.Pow(float64(w)/2, 2) + math.Pow(float64(h)/2, 2))
	d := r * (float64(sceneEffect.Tick) / float64(sceneEffect.Frame)) * 2

	circleImage := generateCircleImage(sceneEffect, w, h, d)
	screen.DrawImage(circleImage, nil)
}

type circularFocusClosingEffect struct{}

func (cfce *circularFocusClosingEffect) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	x := math.Max(float64(sceneEffect.Focus.X), float64(w-sceneEffect.Focus.X))
	y := math.Max(float64(sceneEffect.Focus.Y), float64(h-sceneEffect.Focus.Y))
	r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	d := (r - r*(sceneEffect.Tick/sceneEffect.Frame)) * 2

	circleImage := generateCircleImage(sceneEffect, w, h, d)
	screen.DrawImage(circleImage, nil)
}

type circularFocusOpeningEffect struct{}

func (cfoe *circularFocusOpeningEffect) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	x := math.Max(float64(sceneEffect.Focus.X), float64(w-sceneEffect.Focus.X))
	y := math.Max(float64(sceneEffect.Focus.Y), float64(h-sceneEffect.Focus.Y))
	r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	d := r * (sceneEffect.Tick / sceneEffect.Frame) * 2

	circleImage := generateCircleImage(sceneEffect, w, h, d)
	screen.DrawImage(circleImage, nil)
}
