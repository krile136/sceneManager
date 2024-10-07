package effects

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type CircularClosing struct{}

func (cce *CircularClosing) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	sceneEffect.Focus = Focus{X: w / 2, Y: h / 2}
	r := math.Sqrt(math.Pow(float64(w)/2, 2) + math.Pow(float64(h)/2, 2))
	d := (r - r*(float64(sceneEffect.Tick)/float64(sceneEffect.Frame))) * 2

	circleImage := generateCircleImage(sceneEffect, w, h, d)
	screen.DrawImage(circleImage, nil)
}

type CircularOpening struct{}

func (coe *CircularOpening) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	sceneEffect.Focus = Focus{X: w / 2, Y: h / 2}
	r := math.Sqrt(math.Pow(float64(w)/2, 2) + math.Pow(float64(h)/2, 2))
	d := r * (float64(sceneEffect.Tick) / float64(sceneEffect.Frame)) * 2

	circleImage := generateCircleImage(sceneEffect, w, h, d)
	screen.DrawImage(circleImage, nil)
}

type CircularFocusClosing struct{}

func (cfce *CircularFocusClosing) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	x := math.Max(float64(sceneEffect.Focus.X), float64(w-sceneEffect.Focus.X))
	y := math.Max(float64(sceneEffect.Focus.Y), float64(h-sceneEffect.Focus.Y))
	r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	d := (r - r*(sceneEffect.Tick/sceneEffect.Frame)) * 2

	circleImage := generateCircleImage(sceneEffect, w, h, d)
	screen.DrawImage(circleImage, nil)
}

type CircularFocusOpening struct{}

func (cfoe *CircularFocusOpening) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	x := math.Max(float64(sceneEffect.Focus.X), float64(w-sceneEffect.Focus.X))
	y := math.Max(float64(sceneEffect.Focus.Y), float64(h-sceneEffect.Focus.Y))
	r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	d := r * (sceneEffect.Tick / sceneEffect.Frame) * 2

	circleImage := generateCircleImage(sceneEffect, w, h, d)
	screen.DrawImage(circleImage, nil)
}

func generateCircleImage(se SceneEffect, w, h int, d float64) (img *ebiten.Image) {
	img = ebiten.NewImage(w, h)
	clr := se.Clr
	clr.A = uint8(255)
	img.Fill(clr)

	if d != 0 {
		holeImg := createHoleImage(int(d), clr)

		holeW := float64(holeImg.Bounds().Dx())
		holeH := float64(holeImg.Bounds().Dy())
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-holeW/2+float64(se.Focus.X), -holeH/2+float64(se.Focus.Y))
		op.Blend = ebiten.BlendSourceIn
		img.DrawImage(holeImg, op)
	}
	return
}

// createHoleImage creates an image with the specified diameter and color,
// with a transparent circular hole in the center.
//
// Parameters:
//   - d: The diameter of the image. The image will be a square with this value as both the width and height.
//   - clr: The color to fill the entire image with, before creating the transparent circular hole.
//
// Returns:
//   - *ebiten.Image: The created image.
func createHoleImage(d int, clr color.RGBA) (circleImg *ebiten.Image) {
	circleImg = ebiten.NewImage(d, d)
	circleImg.Fill(clr)

	r := float64(d) / 2
	for y := 0; y < d; y++ {
		for x := 0; x < d; x++ {
			dx := float64(x) - r
			dy := float64(y) - r
			if dx*dx+dy*dy <= r*r {
				circleImg.Set(x, y, color.RGBA{0, 0, 0, 0})
			}
		}
	}

	return
}
