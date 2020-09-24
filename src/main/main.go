package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

type Game struct{}

const (
	screenWidth, screenHeight = 1020, 700
)

var (
	err error
	background *ebiten.Image
	spaceShip *ebiten.Image
	image     *ebiten.Image
)

type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	spaceShip, _, err = ebitenutil.NewImageFromFile("assets/rocket.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

var playerOne = player{
	image: spaceShip,
	xPos:  screenHeight / 2,
	yPos:  screenWidth / 2,
	speed: 4,
}

func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		playerOne.yPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		playerOne.yPos += playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		playerOne.xPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		playerOne.xPos += playerOne.speed
	}
}
func (g *Game) Update(screen *ebiten.Image) error {
	movePlayer()
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	_ = screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	_ = screen.DrawImage(spaceShip, playerOp)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	_ = screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Scale(0.3, 0.3)
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	_ = screen.DrawImage(spaceShip, playerOp)

}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simple go game")

	game := &Game{}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
