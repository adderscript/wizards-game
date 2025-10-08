package collision

import "math"

func CheckCollisionRects(a RectangleCollider, b RectangleCollider) bool {
	return a.X < b.X+b.Width &&
		a.X+a.Width > b.X &&
		a.Y < b.Y+b.Height &&
		a.Y+a.Height > b.Y
}

func CheckCollisionCircles(a CircleCollider, b CircleCollider) bool {
	distance := math.Sqrt((b.X-a.X)*(b.X-a.X) + (b.Y-a.Y)*(b.Y-a.Y))
	return distance <= (a.Radius + b.Radius)
}
