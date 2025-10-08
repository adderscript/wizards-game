package src

import (
	"wizards/config"
	"wizards/src/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player       *entities.Player
	staff        *entities.Staff
	enemyManager *entities.EnemyManager
}

func NewGame() *Game {
	game := &Game{}

	game.player = entities.NewPlayer(config.ScreenWidth/2, config.ScreenHeight/2)
	game.enemyManager = entities.NewEnemyManager(game.player)
	game.staff = entities.NewStaff(game.player, game.enemyManager)

	return game
}

func (g *Game) Update() error {
	g.player.Update()
	g.staff.Update()
	g.enemyManager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	g.staff.Draw(screen)
	g.enemyManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
