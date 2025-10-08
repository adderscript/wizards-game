package entities

import (
	"wizards/config"
	"wizards/libs"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyManager struct {
	Enemies           []*Enemy
	spawnDelaySeconds float64
	spawnTimer        float64
	spawnOffset       float64

	player *Player
}

func NewEnemyManager(player *Player) *EnemyManager {
	enemyMamager := &EnemyManager{
		spawnDelaySeconds: 2.0,
		spawnOffset:       10.0,

		player: player,
	}

	return enemyMamager
}

func (em *EnemyManager) Update() {
	// update all enemies
	for i := len(em.Enemies) - 1; i >= 0; i-- {
		enemy := em.Enemies[i]

		enemy.Update()
		if enemy.ShouldRemove {
			em.Enemies = append(em.Enemies[:i], em.Enemies[i+1:]...)
		}
	}

	// spawn enemy every spawnDelaySeconds
	if em.spawnTimer >= em.spawnDelaySeconds {
		em.spawnEnemy()
		em.spawnTimer = 0.0
	} else {
		em.spawnTimer += 1.0 / 60.0
	}
}

func (em *EnemyManager) Draw(screen *ebiten.Image) {
	// draw all enemies
	for _, enemy := range em.Enemies {
		enemy.Draw(screen)
	}
}

func (em *EnemyManager) spawnEnemy() {
	// set initial position
	side := libs.RandRangeI(0, 3)
	var enemyX, enemyY float64
	switch side {
	case 0: // left
		enemyX = -config.ScreenOffset
		enemyY = float64(libs.RandRangeI(0.0, config.ScreenHeight))
	case 1: // right
		enemyX = config.ScreenWidth + config.ScreenOffset
		enemyY = float64(libs.RandRangeI(0.0, config.ScreenHeight))
	case 2:
		enemyX = float64(libs.RandRangeI(0.0, config.ScreenWidth))
		enemyY = -config.ScreenOffset
	case 3:
		enemyX = float64(libs.RandRangeI(0.0, config.ScreenWidth))
		enemyY = config.ScreenHeight + config.ScreenOffset
	}

	enemy := NewEnemy(enemyX, enemyY, em.player)
	em.Enemies = append(em.Enemies, enemy)
}
