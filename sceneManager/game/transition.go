package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type sceneEffectType int

const (
	Immediately sceneEffectType = iota
	FadeIn
	FadeOut
	CircularClosing
	CircularOpening
)

type TransitionOptions struct {
	OutSceneEffect SceneEffect
	InSceneEffect  SceneEffect
}

type SceneEffect struct {
	Type  sceneEffectType
	Clr   color.RGBA
	Tick  int
	Frame int
}

var transitionHandlerMap = map[sceneEffectType]transitionInterface{
	Immediately:     &immediatelyEffect{},
	FadeIn:          &fadeInEffect{},
	FadeOut:         &fadeOutEffect{},
	CircularClosing: &circularClosingEffect{},
	CircularOpening: &circularOpeningEffect{},
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

type fadeInEffect struct{}

func (fie *fadeInEffect) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	alpha := 255 - 255*(float64(sceneEffect.Tick)/float64(sceneEffect.Frame))
	img := ebiten.NewImage(w, h)
	clr := sceneEffect.Clr
	clr.A = uint8(alpha)
	img.Fill(clr)
	screen.DrawImage(img, nil)
}

type fadeOutEffect struct{}

func (foe *fadeOutEffect) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	alpha := 255 * (float64(sceneEffect.Tick) / float64(sceneEffect.Frame))
	img := ebiten.NewImage(w, h)
	clr := sceneEffect.Clr
	clr.A = uint8(alpha)
	img.Fill(clr)
	screen.DrawImage(img, nil)
}

type circularClosingEffect struct{}

func (cce *circularClosingEffect) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	img := ebiten.NewImage(w, h)
	clr := sceneEffect.Clr
	clr.A = uint8(255)
	img.Fill(clr)

	radius := math.Sqrt(math.Pow(float64(w)/2, 2) + math.Pow(float64(h)/2, 2))
	diameter := (radius - radius*(float64(sceneEffect.Tick)/float64(sceneEffect.Frame))) * 2
	if diameter != 0 {
		holeImage := createCircleImage(int(diameter), clr)

		holeW := float64(holeImage.Bounds().Dx())
		holeY := float64(holeImage.Bounds().Dy())
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-holeW/2+float64(w)/2, -holeY/2+float64(h)/2)
		op.Blend = ebiten.BlendSourceIn
		img.DrawImage(holeImage, op)
	}

	screen.DrawImage(img, nil)
}

type circularOpeningEffect struct{}

func (coe *circularOpeningEffect) Draw(sceneEffect SceneEffect, w, h int, screen *ebiten.Image) {
	img := ebiten.NewImage(w, h)
	clr := sceneEffect.Clr
	clr.A = uint8(255)
	img.Fill(clr)

	radius := math.Sqrt(math.Pow(float64(w)/2, 2) + math.Pow(float64(h)/2, 2))
	diameter := radius * (float64(sceneEffect.Tick) / float64(sceneEffect.Frame)) * 2
	holeImage := createCircleImage(int(diameter), clr)

	holeW := float64(holeImage.Bounds().Dx())
	holeY := float64(holeImage.Bounds().Dy())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-holeW/2+float64(w)/2, -holeY/2+float64(h)/2)
	op.Blend = ebiten.BlendSourceIn
	img.DrawImage(holeImage, op)

	screen.DrawImage(img, nil)
}

func createCircleImage(diameter int, clr color.RGBA) *ebiten.Image {
	circleImage := ebiten.NewImage(diameter, diameter)
	circleImage.Fill(clr)

	radius := float64(diameter) / 2
	for y := 0; y < diameter; y++ {
		for x := 0; x < diameter; x++ {
			dx := float64(x) - radius
			dy := float64(y) - radius
			if dx*dx+dy*dy <= radius*radius {
				circleImage.Set(x, y, color.RGBA{0, 0, 0, 0})
			}
		}
	}

	return circleImage
}
