package entities

import (
	"log"

	"wizards/libs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	X, Y          float64
	Width, Height float64

	velX, velY float64
	speed      float64

	sprite     *ebiten.Image
	facingLeft bool

	player       *Player
	ShouldRemove bool
}

func NewEnemy(x, y float64, player *Player) *Enemy {
	// load image
	img, _, err := ebitenutil.NewImageFromFile("assets/enemy.png")
	if err != nil {
		log.Fatal(err)
	}

	enemy := &Enemy{
		X: x, Y: y,
		Width: 1.0, Height: 1.0,

		speed: 0.70,

		sprite: img,

		player: player,
	}

	return enemy
}

func (e *Enemy) Update() {
	// move towards player
	dirX, dirY := libs.Normalize(e.player.X-e.X, e.player.Y-e.Y)
	e.X += dirX * e.speed
	e.Y += dirY * e.speed
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	spriteSize := e.sprite.Bounds().Size()
	op.GeoM.Translate(-float64(spriteSize.X)/2, -float64(spriteSize.Y)/2)
	op.GeoM.Scale(e.Width, e.Height)

	// flip x if facing left
	if e.facingLeft {
		op.GeoM.Scale(-1, 1)
	} else {
		op.GeoM.Scale(1, 1)
	}

	op.GeoM.Translate(e.X, e.Y)

	screen.DrawImage(e.sprite, op)
}
