package scenes

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/sceneManager/sceneManager/game"
)

type Next struct {
	showEffect bool
  rad float64 
}

func (n *Next) Init() error {
	return nil
}

func (n *Next) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
    n.showEffect = !n.showEffect
	}
  if(n.showEffect){
    n.rad += 0.2

    cf := &game.SceneEffect{
      Type: game.CircularFocusClosing,
      Focus : game.Focus{X: 120, Y:120},
			Clr:   color.RGBA{R: 255, G: 252, B: 219, A: 255},
      Tick: 30 + 10 * math.Sin(n.rad),
      Frame: 60,
    }
    game.ExecuteEffect(cf)
  }

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {

		outSceneEffect := &game.SceneEffect{
			Type:  game.FadeOut,
			Clr:   color.RGBA{R: 0, G: 0, B: 0, A: 255},
			Frame: 30,
		}
		inSceneEffect := &game.SceneEffect{
			Type:  game.FadeIn,
			Clr:   color.RGBA{R: 0, G: 0, B: 0, A: 255},
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
