package game

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/krile136/sceneManager/game/effectType"
)

type Game struct {
	outsideWidth  int
	outsideHeight int
	route         map[string]sceneInterface
	scene         sceneInterface
}

// Game本体を作成します
func MakeGame(outsideWidth, outsideHeight int) *Game {
	game := &Game{
		outsideWidth:  outsideWidth,
		outsideHeight: outsideHeight,
		route:         map[string]sceneInterface{},
	}
	return game
}

// 初期シーンを設定します
func (g *Game) SetInitScene(s sceneInterface) {
	sceneName := g.getSceneName(s)
	g.AddScene(s)
	current = ""
	next = sceneName
}

// 初期シーンを呼び出す際の名前を指定して設定します
func (g *Game) SetInitSceneAsNamed(s sceneInterface, sceneName string) {
	g.AddSceneAsNamed(s, sceneName)
	current = ""
	next = sceneName
}

// シーンを追加します
func (g *Game) AddScene(s sceneInterface) {
	sceneName := g.getSceneName(s)
	g.route[sceneName] = s
}

// シーンを呼び出す際の名前を指定して追加します
func (g *Game) AddSceneAsNamed(s sceneInterface, sceneName string) {
	g.route[sceneName] = s
}

// シーンの名称を取得します
func (g *Game) getSceneName(s sceneInterface) string {
	tv := reflect.TypeOf(s)
	if tv.Kind() == reflect.Ptr {
		// ポインタ型の場合、要素型を取得し直す
		tv = tv.Elem()
	}
	return tv.Name()
}

// シーンのレイアウト決定
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.outsideWidth, g.outsideHeight
}

// 通常のDrawと同様に描画を行います
func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
	if customEffect != nil {
		transitionHandlerMap[customEffect.Type].Draw(*customEffect, g.outsideWidth, g.outsideHeight, screen)
		customEffect = nil
	}
	if !isTransitionFinish {
		if tick <= outSceneEffect.Frame {
			outSceneEffect.Tick = tick
			getOutTransitionHandler().Draw(outSceneEffect, g.outsideWidth, g.outsideHeight, screen)
		} else {
			inSceneEffect.Tick = tick - outSceneEffect.Tick
			getInTransitionHandler().Draw(inSceneEffect, g.outsideWidth, g.outsideHeight, screen)
		}
	}
}

// 通常のUpdateと同様に更新を行います
func (g *Game) Update() error {
	if isTransitionFinish {
		g.scene.Update()
	}

	// g.scene.Update()でisTransitionFinishが書き換わるので
	// 単純にelseで繋げてはいけない
	finishFrame := outSceneEffect.Frame + inSceneEffect.Frame
	if !isTransitionFinish {
		tick += 1

		// シーン切り替え前のトランジションがImmediatelyの場合はtickを待たず即時切り替え
		if tick < outSceneEffect.Frame && outSceneEffect.Type == effectType.Immediately {
			tick = outSceneEffect.Frame
		}

		if tick >= outSceneEffect.Frame {
			if current != next {
				g.scene = g.route[next]
				current = next
				g.scene.Init()
			}
			// シーン切り替え後のトランジションがImmediatelyの場合はtickを待たず終了
			if inSceneEffect.Type == effectType.Immediately {
				tick = finishFrame
			}
		}
		if tick >= finishFrame {
			isTransitionFinish = true
		}
	}

	return nil
}
