package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func generateFadeImage(se SceneEffect, w, h int, alpha float64) (img *ebiten.Image) {
	img = ebiten.NewImage(w, h)
	clr := se.Clr
	clr.A = uint8(alpha)
	img.Fill(clr)
	return
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
