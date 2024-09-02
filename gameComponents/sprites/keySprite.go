package sprites

import "github.com/hajimehoshi/ebiten/v2"

type KeySprite struct {
	*Sprite
	firstSpriteName   string
	currentSpriteName string
	messageText       string
	nextSpriteName    string
	HasNextSprite     bool
	spawnType         string
	Exists            bool
}

func (ks *KeySprite) SetPosition(x, y float64) {
	ks.X = x
	ks.Y = y
}

func (ks *KeySprite) GetMessageText() string {
	return ks.messageText
}

func (ks *KeySprite) SetMessageText(messageText string) {
	ks.messageText = messageText
}

func (ks *KeySprite) SetNextSpriteName(spriteName string) {
	ks.nextSpriteName = spriteName
}

func (ks *KeySprite) GoToNextSprite() {
	if ks.currentSpriteName == ks.nextSpriteName {
		return
	}

	if ks.HasNextSprite {
		ks.currentSpriteName = ks.nextSpriteName
	} else {
		ks.Exists = false
	}
}

func (ks *KeySprite) GoToPrevSprite() {
	if ks.currentSpriteName == ks.firstSpriteName {
		return
	}

	ks.currentSpriteName = ks.firstSpriteName
	ks.Exists = true
}

func (ks *KeySprite) GetSpawnType() string {
	return ks.spawnType
}

func (ks *KeySprite) SetSpawnType(typeName string) {
	ks.spawnType = typeName
}

func (ks *KeySprite) GetCurrentSpriteName() string {
	return ks.currentSpriteName
}

func (ks *KeySprite) AddToPlayerStatus(player *Player) {
	player.Status.AddKey("LevelKey")
}

func (ks *KeySprite) SetSpriteFrameImage(image *ebiten.Image) {
	ks.Sprite.Frame.ImageToDraw = image
}

func (ks *KeySprite) DoesExist() bool {
	return ks.Exists
}
