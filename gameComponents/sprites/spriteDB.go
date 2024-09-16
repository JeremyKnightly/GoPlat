package sprites

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func CreateNamedSprite(name string, filepath string, x1, y1, x2, y2 int) (NamedSprite, error) {
	fullPng, _, err := ebitenutil.NewImageFromFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	subImage := fullPng.SubImage(
		image.Rect(x1, y1, x2, y2),
	).(*ebiten.Image)

	return NamedSprite{
		Sprite: subImage,
		Name:   name,
	}, nil
}

func (npsm *NPSpriteManager) CreateSpriteDB() {
	tilesetsFilepath := "Assets/Maps/Tilesets/"
	itemsFilepath := "Assets/Images/Items/"

	blank, err := CreateNamedSprite("Blank",
		tilesetsFilepath+"Dungeon Tile Set.png", 0, 0, 1, 1)
	if err != nil {
		log.Fatal(err)
	}

	keyVertical, err := CreateNamedSprite("KeyVertical",
		tilesetsFilepath+"Dungeon Tile Set.png", 0, 256, 16, 272)
	if err != nil {
		log.Fatal(err)
	}

	keyHorizontal, err := CreateNamedSprite("KeyHorizontal",
		tilesetsFilepath+"Dungeon Tile Set.png", 176, 144, 192, 160)
	if err != nil {
		log.Fatal(err)
	}

	chestClosed, err := CreateNamedSprite("ChestClosed",
		tilesetsFilepath+"Dungeon Tile Set.png", 192, 144, 208, 160)
	if err != nil {
		log.Fatal(err)
	}

	chestOpen, err := CreateNamedSprite("ChestOpen",
		tilesetsFilepath+"Dungeon Tile Set.png", 208, 144, 224, 160)
	if err != nil {
		log.Fatal(err)
	}

	scroll, err := CreateNamedSprite("Scroll",
		itemsFilepath+"Scroll.png", 0, 0, 16, 16)
	if err != nil {
		log.Fatal(err)
	}

	keySilver, err := CreateNamedSprite("KeySilver",
		itemsFilepath+"Item Sheet.png", 0, 0, 16, 16)
	if err != nil {
		log.Fatal(err)
	}

	keyBronze, err := CreateNamedSprite("KeyBronze",
		itemsFilepath+"Item Sheet.png", 16, 0, 32, 16)
	if err != nil {
		log.Fatal(err)
	}

	keyGold, err := CreateNamedSprite("KeyGold",
		itemsFilepath+"Item Sheet.png", 32, 0, 48, 16)
	if err != nil {
		log.Fatal(err)
	}

	lockSilver, err := CreateNamedSprite("LockSilver",
		itemsFilepath+"Item Sheet.png", 48, 0, 64, 16)
	if err != nil {
		log.Fatal(err)
	}

	lockBronze, err := CreateNamedSprite("LockBronze",
		itemsFilepath+"Item Sheet.png", 64, 0, 80, 16)
	if err != nil {
		log.Fatal(err)
	}

	lockGold, err := CreateNamedSprite("LockGold",
		itemsFilepath+"Item Sheet.png", 80, 0, 96, 16)
	if err != nil {
		log.Fatal(err)
	}

	letter1, err := CreateNamedSprite("Letter1",
		itemsFilepath+"Item Sheet.png", 0, 16, 16, 32)
	if err != nil {
		log.Fatal(err)
	}

	letter2, err := CreateNamedSprite("Letter2",
		itemsFilepath+"Item Sheet.png", 16, 16, 32, 32)
	if err != nil {
		log.Fatal(err)
	}

	letter3, err := CreateNamedSprite("Letter3",
		itemsFilepath+"Item Sheet.png", 32, 16, 48, 32)
	if err != nil {
		log.Fatal(err)
	}

	potionCyan1, err := CreateNamedSprite("PotionCyan1",
		itemsFilepath+"Item Sheet.png", 64, 16, 80, 32)
	if err != nil {
		log.Fatal(err)
	}

	potionRed1, err := CreateNamedSprite("PotionRed1",
		itemsFilepath+"Item Sheet.png", 80, 16, 96, 32)
	if err != nil {
		log.Fatal(err)
	}

	potionAzure1, err := CreateNamedSprite("PotionAzure1",
		itemsFilepath+"Item Sheet.png", 0, 32, 16, 48)
	if err != nil {
		log.Fatal(err)
	}

	potionGreen1, err := CreateNamedSprite("PotionGreen1",
		itemsFilepath+"Item Sheet.png", 16, 32, 32, 48)
	if err != nil {
		log.Fatal(err)
	}

	potionCyan2, err := CreateNamedSprite("PotionCyan2",
		itemsFilepath+"Item Sheet.png", 32, 32, 48, 48)
	if err != nil {
		log.Fatal(err)
	}

	potionRed2, err := CreateNamedSprite("PotionRed2",
		itemsFilepath+"Item Sheet.png", 48, 32, 64, 48)
	if err != nil {
		log.Fatal(err)
	}

	potionAzure2, err := CreateNamedSprite("PotionAzure2",
		itemsFilepath+"Item Sheet.png", 64, 32, 80, 48)
	if err != nil {
		log.Fatal(err)
	}

	potionGreen2, err := CreateNamedSprite("PotionGreen2",
		itemsFilepath+"Item Sheet.png", 80, 32, 96, 48)
	if err != nil {
		log.Fatal(err)
	}

	moneyBag, err := CreateNamedSprite("MoneyBag",
		itemsFilepath+"Item Sheet.png", 0, 48, 16, 64)
	if err != nil {
		log.Fatal(err)
	}

	cog, err := CreateNamedSprite("Cog",
		itemsFilepath+"Item Sheet.png", 16, 48, 32, 64)
	if err != nil {
		log.Fatal(err)
	}

	lanternOff, err := CreateNamedSprite("LanternOff",
		itemsFilepath+"Item Sheet.png", 32, 48, 48, 64)
	if err != nil {
		log.Fatal(err)
	}

	lanternOn, err := CreateNamedSprite("LanternOn",
		itemsFilepath+"Item Sheet.png", 48, 48, 64, 64)
	if err != nil {
		log.Fatal(err)
	}

	npsm.SpriteDB = append(npsm.SpriteDB, blank, keyVertical, keyHorizontal,
		chestClosed, chestOpen, scroll, keySilver, keyBronze, keyGold,
		lockSilver, lockBronze, lockGold, letter1, letter2, letter3, potionCyan1,
		potionRed1, potionAzure1, potionGreen1, potionCyan2, potionRed2, potionAzure2,
		potionGreen2, moneyBag, cog, lanternOff, lanternOn)
}
