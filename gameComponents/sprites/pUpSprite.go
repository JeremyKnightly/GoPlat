package sprites

import "github.com/hajimehoshi/ebiten/v2"

type PUpSprite struct {
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

func (pu *PUpSprite) SetPosition(x, y float64) {
	pu.X = x
	pu.Y = y
}

func (pu *PUpSprite) GetPosition() (float64, float64) {
	return pu.X, pu.Y
}

func (pu *PUpSprite) GetMessageText() string {
	return pu.messageText
}

func (pu *PUpSprite) SetMessageText(messageText string) {
	pu.messageText = messageText
}

func (pu *PUpSprite) SetNextSpriteName(spriteName string) {
	pu.NextSpriteName = spriteName
}

func (pu *PUpSprite) GoToNextSprite() {
	if pu.CurrentSpriteName == pu.NextSpriteName {
		return
	}

	if pu.HasNextSprite {
		pu.CurrentSpriteName = pu.NextSpriteName
	} else {
		pu.Exists = false
	}
}
func (pu *PUpSprite) GoToPrevSprite() {
	if pu.CurrentSpriteName == pu.FirstSpriteName {
		return
	}

	pu.CurrentSpriteName = pu.FirstSpriteName
	pu.Exists = true
}

func (pu *PUpSprite) GetSpawnType() string {
	return pu.SpawnType
}

func (pu *PUpSprite) SetSpawnType(typeName string) {
	pu.SpawnType = typeName
}

func (pu *PUpSprite) GetCurrentSpriteName() string {
	if len(pu.CurrentSpriteName) > 0 {
		return pu.CurrentSpriteName
	} /*else if pu.Exists {
		println("loading prev pup")
		pu.GoToPrevSprite()
		return pu.CurrentSpriteName
	}

	println("InvalidCurrentSpriteName")*/
	return ""
}

func (pu *PUpSprite) AddToPlayerStatus(player *Player) {
	player.Status.AddKey("LevelKey")
}

func (pu *PUpSprite) SetSpriteFrameImage(image *ebiten.Image) {
	pu.Sprite.Frame.ImageToDraw = image
}

func (pu *PUpSprite) GetSpriteFrameImage() *ebiten.Image {
	return pu.Sprite.Frame.ImageToDraw
}

func (pu *PUpSprite) DoesExist() bool {
	return pu.Exists
}
