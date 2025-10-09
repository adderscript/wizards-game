package entities

import (
	"log"

	"wizards/libs/engine"
	"wizards/libs/engine/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Tag string

	components.Transform

	velX, velY float64
	speed      float64

	components.HealthManager

	Collider *engine.CircleCollider

	sprite     *ebiten.Image
	facingLeft bool

	sceneManager *engine.SceneManager
}

func NewPlayer(sceneManager *engine.SceneManager, x, y float64) *Player {
	// load image
	img, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	player := &Player{
		Tag: "player",

		Transform: components.NewTransform(
			x, y,
			1.0, 1.0,
			0.0,
		),

		speed: 1.0,

		HealthManager: components.NewHealthManager(3.0),

		Collider: engine.NewCircleCollider(x, y, 3.0),

		sprite:     img,
		facingLeft: false,

		sceneManager: sceneManager,
	}

	return player
}

func (p *Player) Update() {
	// take input
	p.velX, p.velY = 0.0, 0.0
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.velX = 1.0
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.velX = -1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.velY = 1.0
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.velY = -1.0
	}

	// apply velocity
	p.X += p.velX * p.speed
	p.Y += p.velY * p.speed

	// check if facing left
	mouseX, _ := ebiten.CursorPosition()
	dirX := p.X - float64(mouseX)
	if dirX > 0.0 {
		p.facingLeft = true
	} else if dirX < 0.0 {
		p.facingLeft = false
	}

	p.Collider.Move(p.X, p.Y)
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	spriteSize := p.sprite.Bounds().Size()
	op.GeoM.Translate(-float64(spriteSize.X)/2, -float64(spriteSize.Y)/2)
	op.GeoM.Scale(p.Width, p.Height)

	// flip x if facing left
	if p.facingLeft {
		op.GeoM.Scale(-1, 1)
	} else {
		op.GeoM.Scale(1, 1)
	}

	op.GeoM.Translate(p.X, p.Y)

	screen.DrawImage(p.sprite, op)
}

func (p *Player) GetTag() string {
	return p.Tag
}

func (p *Player) TakeDamage(amount float64) {
	p.Health -= amount
	println(p.Health)
}
