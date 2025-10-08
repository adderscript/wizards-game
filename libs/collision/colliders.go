package collision

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// rectangle collider
type RectangleCollider struct {
	X, Y          float64
	Width, Height float64
}

func NewRectangleCollider(x, y, width, height float64) *RectangleCollider {
	rectangleCollider := &RectangleCollider{
		X: x, Y: y,
		Width: width, Height: height,
	}

	return rectangleCollider
}

func (rc *RectangleCollider) Move(x, y float64) {
	rc.X = x
	rc.Y = y
}

func (rc *RectangleCollider) Draw(screen *ebiten.Image, color color.Color) {
	vector.DrawFilledRect(screen, float32(rc.X), float32(rc.Y), float32(rc.Width), float32(rc.Height), color, true)
}

// circle collider
type CircleCollider struct {
	X, Y   float64
	Radius float64
}

func NewCircleCollider(x, y, radius float64) *CircleCollider {
	circleCollider := &CircleCollider{
		X: x, Y: y,
		Radius: radius,
	}

	return circleCollider
}

func (cc *CircleCollider) Move(x, y float64) {
	cc.X = x
	cc.Y = y
}

func (rc *CircleCollider) Draw(screen *ebiten.Image, color color.Color) {
	vector.DrawFilledCircle(screen, float32(rc.X), float32(rc.Y), float32(rc.Radius), color, true)
}
