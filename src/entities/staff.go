package entities

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Staff struct {
	player  *Player
	bullets []*Bullet

	X, Y          float64
	Width, Height float64
	Rotation      float64

	distance         float64
	offsetX, offsetY float64
	rotationOffset   float64

	shootDelay float64
	shootTimer float64

	sprite         *ebiten.Image
	spriteRotation float64

	enemyManager *EnemyManager
}

func NewStaff(player *Player, enemyManager *EnemyManager) *Staff {
	// load image
	img, _, err := ebitenutil.NewImageFromFile("assets/staff.png")
	if err != nil {
		log.Fatal(err)
	}

	staff := &Staff{
		player: player,

		X: player.X, Y: player.Y,
		Width: 1.0, Height: 1.0,

		distance: 10.0,
		offsetX:  0.0, offsetY: 1.0,
		rotationOffset: math.Pi * 0.25,

		shootDelay: 0.25,

		sprite: img,

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

	// update all bullets
	for i := len(s.bullets) - 1; i >= 0; i-- {
		bullet := s.bullets[i]

		bullet.Update()
		if bullet.ShouldRemove {
			s.bullets = append(s.bullets[:i], s.bullets[i+1:]...)
		}
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

	// draw all bullets
	for _, bullet := range s.bullets {
		bullet.Draw(screen)
	}

	println(len(s.bullets))
}

func (s *Staff) shoot() {
	bullet := NewBullet(s.X, s.Y, s.Rotation, s.enemyManager)
	s.bullets = append(s.bullets, bullet)
}
