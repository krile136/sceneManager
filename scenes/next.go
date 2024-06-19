package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/sceneManager/sceneManager/game"
)

type Next struct{}

func (n *Next) Init() error {
	return nil
}

func (n *Next) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {

		outSceneEffect := &game.SceneEffect{
			Type:  game.FadeOut,
			Clr:   color.RGBA{R: 0, G: 0, B: 0, A: 255},
			Tick:  0,
			Frame: 30,
		}
		inSceneEffect := &game.SceneEffect{
			Type:  game.FadeIn,
			Clr:   color.RGBA{R: 0, G: 0, B: 0, A: 255},
			Tick:  0,
			Frame: 30,
		}

		op := &game.TransitionOptions{
			OutSceneEffect: *outSceneEffect,
			InSceneEffect:  *inSceneEffect,
		}
		game.Change("Top", op)
	}
	return nil
}

func (n *Next) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Next Scene")
}
