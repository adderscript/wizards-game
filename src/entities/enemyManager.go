package entities

import (
	"wizards/config"
	"wizards/libs"
	"wizards/libs/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyManager struct {
	Tag string

	Enemies           []*Enemy
	spawnDelaySeconds float64
	spawnTimer        float64
	spawnOffset       float64

	sceneManager *engine.SceneManager
	player       *Player
}

func NewEnemyManager(sceneManager *engine.SceneManager) *EnemyManager {
	enemyManager := &EnemyManager{
		Tag: "enemyManager",

		spawnDelaySeconds: 2.0,
		spawnOffset:       10.0,

		sceneManager: sceneManager,
	}

	return enemyManager
}

func (em *EnemyManager) Update() {
	// spawn enemy every spawnDelaySeconds
	if em.spawnTimer >= em.spawnDelaySeconds {
		em.spawnEnemy()
		em.spawnTimer = 0.0
	} else {
		em.spawnTimer += 1.0 / 60.0
	}
}

func (em *EnemyManager) Draw(screen *ebiten.Image) {}

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

	enemy := NewEnemy(em.sceneManager, enemyX, enemyY)
	em.sceneManager.AddEntity(enemy)
}

func (em *EnemyManager) GetTag() string {
	return em.Tag
}
