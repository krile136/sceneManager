package scenes

import (
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
		op := &game.TransitionOptions{
			OutTransitionType: game.FadeOut,
			InTransitionType:  game.FadeIn,
			TransitionFrame:   30,
		}
    game.Change("Top", op)
	}
	return nil
}

func (n *Next) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Next Scene")
}
