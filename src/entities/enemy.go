package entities

import (
	"log"

	"wizards/libs"
	"wizards/libs/engine"
	"wizards/libs/engine/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	Tag string

	*components.Transform

	velX, velY float64
	speed      float64

	*components.HealthManager

	sprite     *ebiten.Image
	facingLeft bool

	Collider *engine.CircleCollider

	sceneManager *engine.SceneManager
	player       *Player
}

func NewEnemy(sceneManager *engine.SceneManager, x, y float64) *Enemy {
	player, ok := sceneManager.GetEntity("player").(*Player)
	if !ok {
		log.Fatal("entity 'player' is not of type Player")
	}

	// load image
	img, _, err := ebitenutil.NewImageFromFile("assets/enemy.png")
	if err != nil {
		log.Fatal(err)
	}

	enemy := &Enemy{
		Tag: "enemy",

		Transform: components.NewTransform(
			x, y,
			1.0, 1.0,
			0.0,
		),

		speed: 0.75,

		HealthManager: components.NewHealthManager(2.0),

		sprite: img,

		Collider: engine.NewCircleCollider(x, y, 5.0),

		sceneManager: sceneManager,
		player:       player,
	}

	return enemy
}

func (e *Enemy) Update() {
	// move towards player
	dirX, dirY := libs.Normalize(e.player.X-e.X, e.player.Y-e.Y)
	e.X += dirX * e.speed
	e.Y += dirY * e.speed

	// check collision
	e.Collider.Move(e.X, e.Y)
	if engine.CheckCollisionCircles(*e.Collider, *e.player.Collider) {
		e.player.TakeDamage(1)
	}

	// check if health is depleted
	if e.Health <= 0 {
		e.sceneManager.DeleteEntity(e)
	}
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

func (e *Enemy) TakeDamage(amount float64) {
	e.Health -= amount
}

func (e *Enemy) GetTag() string {
	return e.Tag
}
