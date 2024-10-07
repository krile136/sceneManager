package effectType

type SceneEffectType int

const (
	Immediately SceneEffectType = iota
	FadeIn
	FadeOut
	CircularClosing
	CircularOpening
	CircularFocusClosing
	CircularFocusOpening
)
