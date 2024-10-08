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

3. sceneの切り替えを行えるよう、それぞれのシーンのUpdateメソッドでシーン切り替えAPIを呼びます

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

### シーンの初期化

各シーンは切り替え時にInit関数が一度だけ呼び出されます。
シーン内で使用する変数などの初期化にご利用ください。

### シーン切り替えエフェクトのみの使用

シーンの切り替えエフェクトのみを切り出して使用することもできます。

```go
  op := &effects.SceneEffect{
    Type:  effectType.CircularFocusClosing,
    Focus: effects.Focus{X: 120, Y: 120},
    Clr:   color.RGBA{R: 255, G: 252, B: 219, A: 255},
    Tick:  30 + 10*math.Sin(rad),
    Frame: 60,
  }
  game.ExecuteEffect(op)
```
このようにすると(120,120)を中心に円が表示されます。
radをフレームごとに変化させると半径が経過時間に応じて変化する様子を再現することもできます。
エフェクトの進み具合は`Tick / Frame`で表されます。
例えばTickが30、Frameが60の時は、円で画面を閉じるエフェクトの50%の進捗状況を表します。
Tickを59 にすると画面のほとんどがエフェクトで覆い尽くされており、(120,120)の点に小さな円が残っているのが見えるでしょう。


### シーンの切り替えのオプション値について

```go
  op := &effects.SceneEffect{
    Type:  effectType.CircularFocusClosing,             // シーンの切り替えタイプを選びます
    Focus: effects.Focus{X: 120, Y: 120},               // 特定のエフェクトではフォーカスする点を指定できます
    Clr:   color.RGBA{R: 255, G: 252, B: 219, A: 255},  // 画面を覆っていくエフェクトの色を指定します
    Tick:  30 + 10*math.Sin(rad),                       // Frameに対して現在の進捗の値を指定します（ExecuteEffectの際にのみ使用されます）
    Frame: 60,                                          // エフェクトがスタート→完了するまでのFrame数を指定します
  }
```

### 使用できるシーンの切り替えタイプ

- effectType.Immediately
- effectType.FadeIn
- effectType.FadeOut
- effectType.CircularClosing
- effectType.CircularOpening
- effectType.CircularFocusClosing
- effectType.CircularFocusOpening:
