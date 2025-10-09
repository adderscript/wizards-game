package entities

import (
	"log"
	"math"

	"wizards/config"
	"wizards/libs/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	Tag string

	X, Y          float64
	Width, Height float64
	Rotation      float64

	speed      float64
	velX, velY float64

	damage int

	Collider    *engine.CircleCollider
	isColliding bool

	sprite *ebiten.Image

	sceneManager *engine.SceneManager
}

func NewBullet(sceneManager *engine.SceneManager, x, y, rotation float64) *Bullet {

	// load image
	img, _, err := ebitenutil.NewImageFromFile("assets/bullet.png")
	if err != nil {
		log.Fatal(err)
	}

	bullet := &Bullet{
		Tag: "bullet",

		X: x, Y: y,
		Width: 1.0, Height: 1.0,
		Rotation: rotation,

		speed: 2.5,

		damage: 1,

		Collider: engine.NewCircleCollider(x, y, 1.5),

		sprite: img,

		sceneManager: sceneManager,
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
		b.sceneManager.DeleteEntity(b)
	}

	// check collision
	b.Collider.Move(b.X, b.Y)
	enemies := b.sceneManager.GetEntities("enemy")
	for _, entity := range enemies {
		enemy, ok := entity.(*Enemy)
		if !ok {
			continue
		}
		if engine.CheckCollisionCircles(*enemy.Collider, *b.Collider) {
			enemy.TakeDamage(b.damage)
			b.sceneManager.DeleteEntity(b)
			break
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

func (b *Bullet) GetTag() string {
	return b.Tag
}
