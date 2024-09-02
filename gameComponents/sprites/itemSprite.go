package sprites

import "github.com/hajimehoshi/ebiten/v2"

type ItemSprite struct {
	*Sprite
	firstSpriteName   string
	currentSpriteName string
	messageText       string
	nextSpriteName    string
	HasNextSprite     bool
	spawnType         string
	Exists            bool
}

func (is *ItemSprite) SetPosition(x, y float64) {
	is.X = x
	is.Y = y
}

func (is *ItemSprite) GetMessageText() string {
	return is.messageText
}

func (is *ItemSprite) SetMessageText(messageText string) {
	is.messageText = messageText
}

func (is *ItemSprite) SetNextSpriteName(spriteName string) {
	is.nextSpriteName = spriteName
}

func (is *ItemSprite) GoToNextSprite() {
	if is.currentSpriteName == is.nextSpriteName {
		return
	}

	if is.HasNextSprite {
		is.currentSpriteName = is.nextSpriteName
	} else {
		is.Exists = false
	}
}

func (is *ItemSprite) GoToPrevSprite() {
	if is.currentSpriteName == is.firstSpriteName {
		return
	}

	is.currentSpriteName = is.firstSpriteName
	is.Exists = true
}

func (is *ItemSprite) GetSpawnType() string {
	return is.spawnType
}

func (is *ItemSprite) SetSpawnType(typeName string) {
	is.spawnType = typeName
}

func (is *ItemSprite) GetCurrentSpriteName() string {
	return is.currentSpriteName
}

func (is *ItemSprite) AddToPlayerStatus(player *Player) {
	player.Status.AddKey("LevelKey")
}

func (is *ItemSprite) SetSpriteFrameImage(image *ebiten.Image) {
	is.Sprite.Frame.ImageToDraw = image
}

func (is *ItemSprite) DoesExist() bool {
	return is.Exists
}
