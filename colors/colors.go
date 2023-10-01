package rtcolors

import (
	"image/color"
)

type RtColorRGB struct {
	R, G, B float32
}

func (c1 *RtColorRGB) Add(c2 RtColorRGB) RtColorRGB {
	return RtColorRGB{c1.R + c2.R, c1.G + c2.G, c1.B + c2.B}
}

func (c1 *RtColorRGB) Sub(c2 RtColorRGB) RtColorRGB {
	return RtColorRGB{c1.R - c2.R, c1.G - c2.G, c1.B - c2.B}
}

func (c1 *RtColorRGB) Mul(c2 RtColorRGB) RtColorRGB {
	return RtColorRGB{c1.R * c2.R, c1.G * c2.G, c1.B * c2.B}
}

func (c *RtColorRGB)GetPixel() color.RGBA {
	var base float32
	base = 255.0
	return color.RGBA{uint8(c.R * base), uint8(c.G * base), uint8(c.B * base), 255}
}

func (c1 *RtColorRGB) Scalar(s float32) {
	c1.R = c1.R * s
	c1.G = c1.G * s
	c1.B = c1.B * s
}

func (c *RtColorRGB) RtColorRGB(r, g, b float32) {
	c.R = r
	c.G = g
	c.B = b
}
