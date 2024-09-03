package sprites

import "github.com/hajimehoshi/ebiten/v2"

type ItemSprite struct {
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

func (is *ItemSprite) SetPosition(x, y float64) {
	is.X = x
	is.Y = y
}

func (is *ItemSprite) GetPosition() (float64, float64) {
	return is.X, is.Y
}

func (is *ItemSprite) GetMessageText() string {
	return is.messageText
}

func (is *ItemSprite) SetMessageText(messageText string) {
	is.messageText = messageText
}

func (is *ItemSprite) SetNextSpriteName(spriteName string) {
	is.NextSpriteName = spriteName
}

func (is *ItemSprite) GoToNextSprite() {
	if is.CurrentSpriteName == is.NextSpriteName {
		return
	}

	if is.HasNextSprite {
		is.CurrentSpriteName = is.NextSpriteName
	} else {
		is.Exists = false
	}
}

func (is *ItemSprite) GoToPrevSprite() {
	if is.CurrentSpriteName == is.FirstSpriteName {
		return
	}

	is.CurrentSpriteName = is.FirstSpriteName
	is.Exists = true
}

func (is *ItemSprite) GetSpawnType() string {
	return is.SpawnType
}

func (is *ItemSprite) SetSpawnType(typeName string) {
	is.SpawnType = typeName
}

func (is *ItemSprite) GetCurrentSpriteName() string {
	if len(is.CurrentSpriteName) > 0 {
		return is.CurrentSpriteName
	} /*else if is.Exists {
		is.GoToPrevSprite()
		return is.CurrentSpriteName
	}

	println("InvalidCurrentSpriteName")*/
	return ""
}

func (is *ItemSprite) AddToPlayerStatus(player *Player) {
	player.Status.AddKey("LevelKey")
}

func (is *ItemSprite) SetSpriteFrameImage(image *ebiten.Image) {
	is.Sprite.Frame.ImageToDraw = image
}

func (is *ItemSprite) GetSpriteFrameImage() *ebiten.Image {
	return is.Sprite.Frame.ImageToDraw
}

func (is *ItemSprite) DoesExist() bool {
	return is.Exists
}
