package entities

import (
	"log"
	"math"

	"wizards/config"
	"wizards/libs/collision"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	X, Y          float64
	Width, Height float64
	Rotation      float64

	speed      float64
	velX, velY float64

	Collider    *collision.CircleCollider
	isColliding bool

	sprite *ebiten.Image

	enemyManager *EnemyManager
	ShouldRemove bool
}

func NewBullet(x, y, rotation float64, enemyManager *EnemyManager) *Bullet {
	// load image
	img, _, err := ebitenutil.NewImageFromFile("assets/bullet.png")
	if err != nil {
		log.Fatal(err)
	}

	bullet := &Bullet{
		X: x, Y: y,
		Width: 1.0, Height: 1.0,
		Rotation: rotation,

		speed: 3.5,

		Collider: collision.NewCircleCollider(x, y, 1.5),

		sprite: img,

		enemyManager: enemyManager,
	}

	return bullet
}

func (b *Bullet) Update() {
	// move in facing directory
	b.X += math.Cos(b.Rotation) * b.speed
	b.Y += math.Sin(b.Rotation) * b.speed

	// delete if off screen
	if (b.X <= -config.ScreenOffset || b.X >= config.ScreenWidth+config.ScreenOffset) ||
		(b.Y <= -config.ScreenOffset || b.Y >= config.ScreenHeight+config.ScreenOffset) {
		b.ShouldRemove = true
	}

	// check collision
	b.Collider.Move(b.X, b.Y)
	enemies := b.enemyManager.Enemies
	for i := len(enemies) - 1; i >= 0; i-- {
		enemy := enemies[i]
		if collision.CheckCollisionCircles(*enemy.Collider, *b.Collider) {
			enemy.ShouldRemove = true
			b.ShouldRemove = true
		}
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	spriteSize := b.sprite.Bounds().Size()
	op.GeoM.Translate(-float64(spriteSize.X)/2, -float64(spriteSize.Y)/2)
	op.GeoM.Rotate(b.Rotation)
	op.GeoM.Scale(b.Width, b.Height)
	op.GeoM.Translate(b.X, b.Y)

	screen.DrawImage(b.sprite, op)
}
