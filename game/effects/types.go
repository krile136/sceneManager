package effects

import (
	"image/color"

  "github.com/krile136/sceneManager/game/effectType"
)

type SceneEffect struct {
	Type  effectType.SceneEffectType
	Focus Focus
	Clr   color.RGBA
	Tick  float64
	Frame float64
}

type Focus struct {
	X int
	Y int
}
