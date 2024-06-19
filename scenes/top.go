package scenes

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/sceneManager/sceneManager/game"
)

type Top struct {
	image *ebiten.Image
	hole  *ebiten.Image
	tick  float64
}

func (t *Top) Init() error {
	t.image = ebiten.NewImage(20, 20)
	t.tick = 0
	clr := color.RGBA{R: 255, G: 0, B: 0, A: uint8(t.tick)}
	t.image.Fill(clr)

	sourceImage := ebiten.NewImage(200, 200)
	sourceImage.Fill(color.RGBA{255, 0, 0, 255}) // 赤色で塗りつぶします

	resultImage := ebiten.NewImageFromImage(sourceImage)

	// 丸い穴を作成します
	holeDiameter := 100
	holeX, holeY := 50, 50
	holeImage := createCircleImage(holeDiameter)

	// 描画オプションを設定します
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(holeX), float64(holeY))
	op.Blend = ebiten.BlendSourceIn
	resultImage.DrawImage(holeImage, op)

	t.hole = resultImage

	return nil
}

func createCircleImage(diameter int) *ebiten.Image {
	circleImage := ebiten.NewImage(diameter, diameter)
	circleImage.Fill(color.RGBA{255, 0, 0, 255}) // 透明で塗りつぶします

	radius := float64(diameter) / 2
	for y := 0; y < diameter; y++ {
		for x := 0; x < diameter; x++ {
			dx := float64(x) - radius
			dy := float64(y) - radius
			if dx*dx+dy*dy <= radius*radius {
				circleImage.Set(x, y, color.RGBA{0, 0, 0, 0}) // 不透明にする
			}
		}
	}

	return circleImage
}

func (t *Top) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		op := &game.TransitionOptions{
			OutTransitionType: game.CircularClosing,
			InTransitionType:  game.CircularOpening,
			TransitionFrame:   60,
      Clr: color.RGBA{R: 24, G: 235, B: 249, A: 255},
		}
		game.Change("Next", op)
		//scene.Reload()
	}

	t.tick = math.Min(255, t.tick+3)
	clr := color.RGBA{R: 0, G: 0, B: 0, A: uint8(t.tick)}
	t.image.Fill(clr)

	return nil
}

func (t *Top) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "top!")

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(20, 20)

	img := ebiten.NewImage(50, 50)
	clr := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	img.Fill(clr)
	screen.DrawImage(img, op)

	screen.DrawImage(t.image, op)
	screen.DrawImage(t.hole, op)
}
