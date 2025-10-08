package src

import (
	"wizards/config"
	"wizards/libs/engine"
	"wizards/src/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneManager *engine.SceneManager
}

func NewGame() *Game {
	sceneManager := engine.NewSceneManager()

	player := entities.NewPlayer(sceneManager, config.ScreenWidth/2, config.ScreenHeight/2)
	sceneManager.AddEntity(player)

	enemyManager := entities.NewEnemyManager(sceneManager)
	sceneManager.AddEntity(enemyManager)

	staff := entities.NewStaff(sceneManager)
	sceneManager.AddEntity(staff)

	game := &Game{
		sceneManager: sceneManager,
	}

	return game
}

func (g *Game) Update() error {
	g.sceneManager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
