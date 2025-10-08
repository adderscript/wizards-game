package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneManager struct {
	entities []Entity
}

func NewSceneManager() *SceneManager {
	sceneManager := &SceneManager{}

	return sceneManager
}

func (sc *SceneManager) Update() {
	for i := len(sc.entities) - 1; i >= 0; i-- {
		if i >= len(sc.entities) {
			continue
		}
		entity := sc.entities[i]
		entity.Update()
	}
}

func (sc *SceneManager) Draw(screen *ebiten.Image) {
	for i := len(sc.entities) - 1; i >= 0; i-- {
		entity := sc.entities[i]
		entity.Draw(screen)
	}
}

func (sc *SceneManager) AddEntity(entity Entity) {
	sc.entities = append(sc.entities, entity)
}

func (sc *SceneManager) DeleteEntity(entity Entity) {
	for i, e := range sc.entities {
		if e == entity {
			sc.entities = append(sc.entities[:i], sc.entities[i+1:]...)
			break
		}
	}
}

func (sc *SceneManager) GetEntity(tag string) Entity {
	for _, entity := range sc.entities {
		if entity.GetTag() == tag {
			return entity
		}
	}
	return nil
}

func (sc *SceneManager) GetEntities(tag string) []Entity {
	var entities []Entity
	for _, entity := range sc.entities {
		if entity.GetTag() == tag {
			entities = append(entities, entity)
		}
	}
	return entities
}
