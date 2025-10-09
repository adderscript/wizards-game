package components

type HealthManager struct {
	MaxHealth float64
	Health    float64
}

func NewHealthManager(maxHealth float64) HealthManager {
	health := HealthManager{
		MaxHealth: maxHealth,
		Health:    maxHealth,
	}

	return health
}
