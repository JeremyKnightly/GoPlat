package sprites

import "github.com/hajimehoshi/ebiten/v2"

type KeySprite struct {
	*Sprite
	FirstSpriteName   string
	CurrentSpriteName string
	messageText       string
	NextSpriteName    string
	HasNextSprite     bool
	SpawnType         string
	Exists            bool
	UniqueName        string
}

func (ks *KeySprite) SetPosition(x, y float64) {
	ks.Sprite.SetPosition(x, y)
}

func (ks *KeySprite) GetPosition() (float64, float64) {
	return ks.Sprite.X, ks.Sprite.Y
}

func (ks *KeySprite) GetMessageText() string {
	return ks.messageText
}

func (ks *KeySprite) SetMessageText(messageText string) {
	ks.messageText = messageText
}

func (ks *KeySprite) SetNextSpriteName(spriteName string) {
	ks.NextSpriteName = spriteName
}

func (ks *KeySprite) GoToNextSprite() {
	if ks.CurrentSpriteName == ks.NextSpriteName {
		return
	}

	if ks.HasNextSprite {
		ks.CurrentSpriteName = ks.NextSpriteName
	} else {
		ks.Exists = false
	}
}

func (ks *KeySprite) GoToPrevSprite() {
	if ks.CurrentSpriteName == ks.FirstSpriteName {
		return
	}

	ks.CurrentSpriteName = ks.FirstSpriteName
	ks.Exists = true
}

func (ks *KeySprite) GetSpawnType() string {
	return ks.SpawnType
}

func (ks *KeySprite) SetSpawnType(typeName string) {
	ks.SpawnType = typeName
}

func (ks *KeySprite) GetCurrentSpriteName() string {
	if len(ks.CurrentSpriteName) > 0 {
		return ks.CurrentSpriteName
	}

	return ""
}

func (ks *KeySprite) AddToPlayerStatus(player *Player) {
	player.Status.AddKey("LevelKey")
}

func (ks *KeySprite) SetSpriteFrameImage(image *ebiten.Image) {
	ks.Sprite.Frame.ImageToDraw = image
}

func (ks *KeySprite) GetSpriteFrameImage() *ebiten.Image {
	return ks.Sprite.Frame.ImageToDraw
}

func (ks *KeySprite) DoesExist() bool {
	return ks.Exists
}

func (ks *KeySprite) GetSpawnName() string {
	return ks.UniqueName
}

func (ks *KeySprite) SetSpawnName(name string) {
	ks.UniqueName = name
}
