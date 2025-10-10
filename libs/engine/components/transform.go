package components

type Transform struct {
	X, Y          float64
	Width, Height float64
	Rotation      float64
}

func NewTransform(x, y, width, height, rotation float64) *Transform {
	Transform := &Transform{
		X: x, Y: y,
		Width: width, Height: height,
		Rotation: rotation,
	}

	return Transform
}
