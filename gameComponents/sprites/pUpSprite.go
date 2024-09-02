package sprites

import "github.com/hajimehoshi/ebiten/v2"

type PUpSprite struct {
	*Sprite
	currentSpriteName string
	messageText       string
	nextSpriteName    string
	HasNextSprite     bool
	spawnType         string
}

func (pu *PUpSprite) SetPosition(x, y float64) {
	pu.X = x
	pu.Y = y
}

func (pu *PUpSprite) GetMessageText() string {
	return pu.messageText
}

func (pu *PUpSprite) SetMessageText(messageText string) {
	pu.messageText = messageText
}

func (pu *PUpSprite) SetNextSpriteName(spriteName string) {
	pu.nextSpriteName = spriteName
}

func (pu *PUpSprite) GoToNextSprite() {
	if pu.HasNextSprite {
		pu.currentSpriteName = pu.nextSpriteName
		//ks.Frame.EffectImageToDraw = SpriteMaker.Get(ks.nextSpriteName)
	}
}

func (pu *PUpSprite) GetSpawnType() string {
	return pu.spawnType
}

func (pu *PUpSprite) SetSpawnType(typeName string) {
	pu.spawnType = typeName
}

func (pu *PUpSprite) GetCurrentSpriteName() string {
	return pu.currentSpriteName
}

func (pu *PUpSprite) AddToPlayerStatus(player *Player) {
	player.Status.AddKey("LevelKey")
}

func (pu *PUpSprite) SetSpriteFrameImage(image *ebiten.Image) {
	pu.Sprite.Frame.ImageToDraw = image
}
