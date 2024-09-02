package sprites

import "github.com/hajimehoshi/ebiten/v2"

type KeySprite struct {
	*Sprite
	currentSpriteName string
	messageText       string
	nextSpriteName    string
	HasNextSprite     bool
	spawnType         string
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
	if ks.HasNextSprite {
		ks.currentSpriteName = ks.nextSpriteName
		//ks.Frame.EffectImageToDraw = SpriteMaker.Get(ks.nextSpriteName)
	}
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
