package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type TransitionOptions struct {
	OutTransitionType TransitionType
	InTransitionType  TransitionType
	TransitionFrame   int
	Clr               color.RGBA
}

type TransitionType int

const (
	Immediately TransitionType = iota
	FadeIn
	FadeOut
	CircularClosing
	CircularOpening
)

var transitionHandlerMap = map[TransitionType]TransitionInterface{
	Immediately:     &TransitionImmediately{},
	FadeIn:          &TransitionFadeIn{},
	FadeOut:         &TransitionFadeOut{},
	CircularClosing: &TransitionCircularClosing{},
	CircularOpening: &TransitionCircularOpening{},
}

func getOutTransitionHandler() TransitionInterface {
	return transitionHandlerMap[outTransitionType]
}
func getInTransitionHandler() TransitionInterface {
	return transitionHandlerMap[inTransitionType]
}

type TransitionInterface interface {
	Draw(tick int, transitionFrame int, w int, h int, clr color.RGBA, screen *ebiten.Image)
}

type TransitionImmediately struct{}

func (ti *TransitionImmediately) Draw(tick, transitionFrame, w, h int, clr color.RGBA, screen *ebiten.Image) {
	img := ebiten.NewImage(w, h)
	clr = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	img.Fill(clr)
	screen.DrawImage(img, nil)
}

type TransitionFadeIn struct{}

func (tfi *TransitionFadeIn) Draw(tick, transitionFrame, w, h int, clr color.RGBA, screen *ebiten.Image) {
	alpha := 255 - 255*(float64(tick)/float64(transitionFrame))
	img := ebiten.NewImage(w, h)
  clr.A = uint8(alpha)
	img.Fill(clr)
	screen.DrawImage(img, nil)
}

type TransitionFadeOut struct{}

func (tfo *TransitionFadeOut) Draw(tick, transitionFrame, w, h int, clr color.RGBA, screen *ebiten.Image) {
	alpha := 255 * (float64(tick) / float64(transitionFrame))
	img := ebiten.NewImage(w, h)
  clr.A = uint8(alpha)
	img.Fill(clr)
	screen.DrawImage(img, nil)
}

type TransitionCircularClosing struct{}

func (tcc *TransitionCircularClosing) Draw(tick, transitionFrame, w, h int, clr color.RGBA, screen *ebiten.Image) {
	img := ebiten.NewImage(w, h)
  clr.A = uint8(255)
	img.Fill(clr)

	radius := math.Sqrt(math.Pow(float64(w)/2, 2) + math.Pow(float64(h)/2, 2))
	diameter := (radius - radius*(float64(tick)/float64(transitionFrame))) * 2
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

type TransitionCircularOpening struct{}

func (tco *TransitionCircularOpening) Draw(tick, transitionFrame, w, h int, clr color.RGBA, screen *ebiten.Image) {
	img := ebiten.NewImage(w, h)
  clr.A = uint8(255)
	img.Fill(clr)

	radius := math.Sqrt(math.Pow(float64(w)/2, 2) + math.Pow(float64(h)/2, 2))
	diameter := radius * (float64(tick) / float64(transitionFrame)) * 2
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
