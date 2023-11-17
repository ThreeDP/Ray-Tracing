package rtcolors

import (
	"testing"
	"math"
	"image/color"
	"rt/defines"
)


func isEqual(x, y float32, epsilon float64) bool {
	if math.Abs(float64(x - y)) < epsilon {
		return true
	}
	return false
}

func compareColors(t *testing.T, a, b RtColorRGB) {
	t.Helper()

    if !isEqual(a.R, b.R, defines.EPSILON) || !isEqual(a.G, b.G, defines.EPSILON) || !isEqual(a.B, b.B, defines.EPSILON){
        t.Errorf("Expected color(R: %f, G: %f, B: %f), but has color(R: %f, G: %f, B: %f)", a.R, a.G, a.B, b.R, b.G, b.B)
    }
}

func TestColor(t *testing.T) {
	t.Run("Create a Color", func(t *testing.T) {
		var rgb RtColorRGB
		rgb.RtColorRGB(-0.5, 0.4, 1.7)
		compareColors(t, rgb, RtColorRGB{-0.5, 0.4, 1.7})
	})

	t.Run("Adding colors", func(t *testing.T) {
		var c1, c2 RtColorRGB
		c1.RtColorRGB(0.9, 0.6, 0.75)
		c2.RtColorRGB(0.7, 0.1, 0.25)
		cRes := c1.Add(c2)
		compareColors(t, cRes, RtColorRGB{1.6, 0.7, 1.0})
	})

	t.Run("Subtracting colors", func(t *testing.T) {
		var c1, c2 RtColorRGB
		c1.RtColorRGB(0.9, 0.6, 0.75)
		c2.RtColorRGB(0.7, 0.1, 0.25)
		cRes := c1.Sub(c2)
		compareColors(t, cRes, RtColorRGB{0.2, 0.5, 0.5})
	})

	t.Run("Multiplying a color by scalar", func(t *testing.T) {
		var c1 RtColorRGB
		c1.RtColorRGB(0.2, 0.3, 0.4)
		c1.Scalar(2)
		compareColors(t, c1, RtColorRGB{0.4, 0.6, 0.8})
	})

	t.Run("Multiplying colors", func(T *testing.T) {
		var c1, c2 RtColorRGB
		c1.RtColorRGB(1, 0.2, 0.4)
		c2.RtColorRGB(0.9, 1, 0.1)
		cRes := c1.Mul(c2)
		compareColors(t, cRes, RtColorRGB{0.9, 0.2, 0.04})
	})

	t.Run("Get pixel", func(T *testing.T) {
		var c1 RtColorRGB
		c1.RtColorRGB(1, 0.4, 0.4)
		cRes := c1.GetPixel()
		cExp := color.RGBA{255, 102, 102, 255}
		if cRes != cExp {
			t.Errorf("Expected color(R: %d, G: %d, B: %d), but has color(R: %d, G: %d, B: %d)", cRes.R, cRes.G, cRes.B, cExp.R, cExp.G, cExp.B)
		}
	})
}