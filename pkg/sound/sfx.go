package sound

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var sfxPlayer *audio.Player

func NewSfxPlayer(name string) {
	switch name {
	case "big_jump":
		sfxPlayer = loadMusic("sound/big_jump.ogg")
	case "small_jump":
		sfxPlayer = loadMusic("sound/small_jump.ogg")
	case "brick_smash":
		sfxPlayer = loadMusic("sound/brick_smash.ogg")
	}
	sfxPlayer.Play()
}
