# Ebitengine シーンマネージャー

このパッケージは、Ebitengineを使用したゲーム開発において、シーンの管理と切り替えを簡単に行うためのツールです。シーンの遷移をスムーズにし、コードの可読性と保守性を向上させることができます。


## 特徴

- シンプルなAPIでシーンの切り替えを実現
- シーンごとの初期化とクリーンアップを自動化
- 柔軟なシーン遷移の制御が可能

## インストール

以下のコマンドでインストールできます:

```sh
go get github.com/krile136/sceneManager/game
```

## 使い方

基本的な使い方は以下の通りです:

### シーンの登録

1. Init, Update, Drawメソッドを持つシーンを作成します。ここでは、TopとNextの２つのシーンを作ります。

```go
import (
  "github.com/hajimehoshi/ebiten/v2"
)

// Top Scence
type Top struct {}

func (t *Top) Init() error {
  return nil
}

func (t *Top) Update() error {
  return nil
}

func (t *Top) Draw(screen *ebiten.Image) {
  ebitenutil.DebugPrint(screen, "top!")
}

// Next Scene
type Next struct {}

func (n *Next) Init() error {
  return nil
}

func (n *Next) Update() error {
  return nil
}

func (n *Next) Draw(screen *ebiten.Image) {
  ebitenutil.DebugPrint(screen, "Next!")
}

```

2. main.goで1で作ったシーンを登録します。

```go
package main

import (
  "log"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/krile136/sceneManager/game"
)


func main() {
  game := game.MakeGame(320, 240)
  
  game.SetInitScene(&Top{})     // 最初に表示したいシーンはSetInitSCeneを使用して登録します
  game.AddScene(&Next{})        // 最初以外のシーンはAddSceneを使用して登録します

  ebiten.SetWindowSize(640, 480)
  ebiten.SetWindowTitle("Hello, SceneManager!")
  if err := ebiten.RunGame(game); err != nil {
	log.Fatal(err)
  }
}
```

3. sceneの切り替えを行えるようにします

```go
// Top Scence Update method
func (t *Top) Update() error {
  if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
	game.Change("Next", nil)
  }

  return nil
}



// Next Scene Update method
func (n *Next) Update() error {
  if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
    outSceneEffect := &effects.SceneEffect{
	  Type:  effectType.FadeOut,
	  Clr:   color.RGBA{R: 0, G: 0, B: 0, A: 255},
	  Frame: 30,
	}
	inSceneEffect := &effects.SceneEffect{
	  Type:  effectType.FadeIn,
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
```

4. ゲームを実行します

Top画面でEnterを押すとNext画面に即時遷移し、Next画面でEnterを押すとTop画面へフェードアウト/フェードインをしながら遷移します。
