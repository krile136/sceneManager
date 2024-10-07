package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/sceneManager/game"
	"github.com/krile136/sceneManager/example/scenes"
)


func main() {
  game := game.MakeGame(320, 240)
  
  game.SetInitScene(&scenes.Top{})
  game.AddScene(&scenes.Next{})

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
