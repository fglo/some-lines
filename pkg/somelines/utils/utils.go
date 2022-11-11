package utils

import "math"

func ColorPixel(x, y, screenWidth, screenHeight int, pixels []byte) {
	if x >= 0 && x < screenWidth && y > 0 && y < screenHeight {
		i := getPixelsIndex(x, y, screenWidth, screenHeight)
		pixels[4*i] = 0xf0
		pixels[4*i+1] = 0xf0
		pixels[4*i+2] = 0xf0
		pixels[4*i+3] = 0xff
	}
}

func Color3DPixel(x, y int, z float64, screenWidth, screenHeight int, pixels []byte) {
	if x >= 0 && x < screenWidth && y > 0 && y < screenHeight {
		colorAt0 := byte(0xe0)
		color := colorAt0

		modifier := 0.3 * z
		switch {
		case modifier > 200:
			modifier = 200
		case modifier < -30:
			modifier = -30
		}

		color = colorAt0 - byte(modifier)

		i := getPixelsIndex(x, y, screenWidth, screenHeight)
		if pixels[4*i] < color {
			pixels[4*i] = color
			pixels[4*i+1] = color
			pixels[4*i+2] = color
			pixels[4*i+3] = 0xff
		}
	}
}

func getPixelsIndex(x, y, screenWidth, screenHeight int) int {
	x = keepCoordBetweenMinAndMax(x, 0, screenWidth)
	y = keepCoordBetweenMinAndMax(y, 0, screenHeight)

	max := screenWidth * screenHeight
	i := keepCoordBetweenMinAndMax(y*screenWidth+x, 0, max)
	return i
}

func keepCoordBetweenMinAndMax(coord, min, max int) int {
	for coord < min || coord >= max {
		switch {
		case coord < min:
			coord += max
		case coord >= max:
			coord -= max
		}
	}
	return coord
}

func Deg2Rad(deg float64) float64 {
	deg = math.Mod(deg, 360)
	return deg * math.Pi / 180.0
}
