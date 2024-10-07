package effects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Immediately struct{}

func (ie *Immediately) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	img := ebiten.NewImage(w, h)
	clr := sceneEffect.Clr
	img.Fill(clr)
	screen.DrawImage(img, nil)
}
