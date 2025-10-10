package entities

import (
	"log"
	"math"
	"wizards/libs/engine"
	"wizards/libs/engine/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Staff struct {
	Tag string

	*components.Transform

	distance         float64
	offsetX, offsetY float64
	rotationOffset   float64

	bullets    []*Bullet
	shootDelay float64
	shootTimer float64

	sprite         *ebiten.Image
	spriteRotation float64

	sceneManager *engine.SceneManager
	player       *Player
	enemyManager *EnemyManager
}

func NewStaff(sceneManager *engine.SceneManager) *Staff {
	player, ok := sceneManager.GetEntity("player").(*Player)
	if !ok {
		log.Fatal("entity 'player' is not of type Player")
	}
	enemyManager, ok := sceneManager.GetEntity("enemyManager").(*EnemyManager)
	if !ok {
		log.Fatal("entity 'enemyManager' is not of type EnemyManager")
	}

	// load image
	img, _, err := ebitenutil.NewImageFromFile("assets/staff.png")
	if err != nil {
		log.Fatal(err)
	}

	staff := &Staff{
		Tag: "staff",

		Transform: components.NewTransform(
			player.X, player.X,
			1.0, 1.0,
			0.0,
		),

		distance: 10.0,
		offsetX:  0.0, offsetY: 1.0,
		rotationOffset: math.Pi * 0.25,

		shootDelay: 0.25,

		sprite: img,

		sceneManager: sceneManager,
		player:       player,
		enemyManager: enemyManager,
	}

	return staff
}

func (s *Staff) Update() {
	// rotate towards mouse
	mouseX, mouseY := ebiten.CursorPosition()
	dirX, dirY := float64(mouseX)-s.player.X, float64(mouseY)-s.player.Y
	s.Rotation = math.Atan2(dirY, dirX)

	s.spriteRotation = s.Rotation + s.rotationOffset

	// set position
	s.X = s.player.X + s.offsetX + math.Cos(s.Rotation)*s.distance
	s.Y = s.player.Y + s.offsetY + math.Sin(s.Rotation)*s.distance

	// shoot with lmb
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && s.shootTimer >= s.shootDelay {
		s.shoot()
		s.shootTimer = 0.0
	} else {
		s.shootTimer += 1.0 / 60.0
	}
}

func (s *Staff) Draw(screen *ebiten.Image) {
	// draw staff
	op := &ebiten.DrawImageOptions{}
	spriteSize := s.sprite.Bounds().Size()
	op.GeoM.Translate(-float64(spriteSize.X)/2, -float64(spriteSize.Y)/2)
	op.GeoM.Rotate(s.spriteRotation)
	op.GeoM.Scale(s.Width, s.Height)
	op.GeoM.Translate(s.X, s.Y)

	screen.DrawImage(s.sprite, op)
}

func (s *Staff) shoot() {
	bullet := NewBullet(s.sceneManager, s.X, s.Y, s.Rotation)
	s.sceneManager.AddEntity(bullet)
}

func (s *Staff) GetTag() string {
	return s.Tag
}
