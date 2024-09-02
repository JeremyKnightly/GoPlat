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

	lockSilver, err := CreateNamedSprite("lockSilver",
		itemsFilepath+"Item Sheet.png", 48, 0, 64, 16)
	if err != nil {
		log.Fatal(err)
	}

	lockBronze, err := CreateNamedSprite("lockBronze",
		itemsFilepath+"Item Sheet.png", 64, 0, 80, 16)
	if err != nil {
		log.Fatal(err)
	}

	lockGold, err := CreateNamedSprite("lockGold",
		itemsFilepath+"Item Sheet.png", 80, 0, 96, 16)
	if err != nil {
		log.Fatal(err)
	}

	letter1, err := CreateNamedSprite("letter1",
		itemsFilepath+"Item Sheet.png", 0, 16, 16, 32)
	if err != nil {
		log.Fatal(err)
	}

	letter2, err := CreateNamedSprite("letter2",
		itemsFilepath+"Item Sheet.png", 16, 16, 32, 32)
	if err != nil {
		log.Fatal(err)
	}

	letter3, err := CreateNamedSprite("letter3",
		itemsFilepath+"Item Sheet.png", 32, 16, 48, 32)
	if err != nil {
		log.Fatal(err)
	}

	potionCyan1, err := CreateNamedSprite("potionCyan1",
		itemsFilepath+"Item Sheet.png", 64, 16, 80, 32)
	if err != nil {
		log.Fatal(err)
	}

	potionRed1, err := CreateNamedSprite("potionRed1",
		itemsFilepath+"Item Sheet.png", 80, 16, 96, 32)
	if err != nil {
		log.Fatal(err)
	}

	potionAzure1, err := CreateNamedSprite("potionAzure1",
		itemsFilepath+"Item Sheet.png", 0, 32, 16, 48)
	if err != nil {
		log.Fatal(err)
	}

	potionGreen1, err := CreateNamedSprite("potionGreen1",
		itemsFilepath+"Item Sheet.png", 16, 32, 32, 48)
	if err != nil {
		log.Fatal(err)
	}

	potionCyan2, err := CreateNamedSprite("potionCyan2",
		itemsFilepath+"Item Sheet.png", 32, 32, 48, 48)
	if err != nil {
		log.Fatal(err)
	}

	potionRed2, err := CreateNamedSprite("potionRed2",
		itemsFilepath+"Item Sheet.png", 48, 32, 64, 48)
	if err != nil {
		log.Fatal(err)
	}

	potionAzure2, err := CreateNamedSprite("potionAzure2",
		itemsFilepath+"Item Sheet.png", 64, 32, 80, 48)
	if err != nil {
		log.Fatal(err)
	}

	potionGreen2, err := CreateNamedSprite("potionGreen2",
		itemsFilepath+"Item Sheet.png", 80, 32, 96, 48)
	if err != nil {
		log.Fatal(err)
	}

	moneyBag, err := CreateNamedSprite("moneyBag",
		itemsFilepath+"Item Sheet.png", 0, 48, 16, 64)
	if err != nil {
		log.Fatal(err)
	}

	cog, err := CreateNamedSprite("cog",
		itemsFilepath+"Item Sheet.png", 16, 48, 32, 64)
	if err != nil {
		log.Fatal(err)
	}

	lanternOff, err := CreateNamedSprite("lanternOff",
		itemsFilepath+"Item Sheet.png", 32, 48, 48, 64)
	if err != nil {
		log.Fatal(err)
	}

	lanternOn, err := CreateNamedSprite("lanternOn",
		itemsFilepath+"Item Sheet.png", 48, 48, 64, 64)
	if err != nil {
		log.Fatal(err)
	}

	npsm.SpriteDB = append(npsm.SpriteDB, keyVertical, keyHorizontal,
		chestClosed, chestOpen, scroll, keySilver, keyBronze, keyGold,
		lockSilver, lockBronze, lockGold, letter1, letter2, letter3, potionCyan1,
		potionRed1, potionAzure1, potionGreen1, potionCyan2, potionRed2, potionAzure2,
		potionGreen2, moneyBag, cog, lanternOff, lanternOn)
}
