package raycast

import "github.com/faiface/pixel/pixelgl"

// KeyPressed is for checking if a update should be done
func KeyPressed(win *pixelgl.Window) bool {
	expr1 := win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyS)
	expr2 := win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyD)
	return expr1 || expr2
}

// CastAllRays draws a collection of rays in the screen
func CastAllRays(p *Player, w *pixelgl.Window, g *Map) {
	column := 0.0
	rayAngle := p.rotationAngle - (fov / 2)

	for i := 0; i < int(numRays); i += 2 {
		ray := NewRay(rayAngle, w, p, g)
		ray.Cast(column)
		rayAngle += fov / numRays
		column++
	}
}
