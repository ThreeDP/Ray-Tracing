package tuples

import (
	"math"
)

type axes struct {
	x, y, z, w float64
}

type point struct {
	axes
}

type vector struct {
	axes
}

type tuple interface {
	Create(x, y, z float64)
	GetAxes() axes
	Add(b tuple) tuple
	Subtracting(b tuple) tuple
	Negate()
	Multiply(m float64)
	Divide(m float64)
	DotProduct(t tuple) float64
}

func (p *point) Create(x, y, z float64) {
	p.x = x
	p.y = y
	p.z = z
	p.w = 1.0
}

func (p *point) GetAxes() axes {
	return axes{p.x, p.y, p.z, p.w};
}

func (a *point) Add(b tuple) tuple {
	b1 := b.GetAxes()
	w := a.w + b1.w 
	if w > 0 {
		return &point{axes{a.x + b1.x, a.y + b1.y, a.z + b1.z, w}}
	}
	return &vector{axes{a.x + b1.x, a.y + b1.y, a.z + b1.z, w}}
}

func (a *point) Subtracting(b tuple) tuple {
	b1 := b.GetAxes()
	w := math.Abs(a.w - b1.w)
	if w != 0 {
		return &point{axes{a.x - b1.x, a.y - b1.y, a.z - b1.z, w}}
	}
	return &vector{axes{a.x - b1.x, a.y - b1.y, a.z - b1.z, w}}
}

func (p *point) Negate() {
	p.x = -p.x
	p.y = -p.y
	p.z = -p.z
	p.w = -p.w
}

func (p *point) Multiply(m float64) {
	p.x = p.x * m
	p.y = p.y * m
	p.z = p.z * m
	p.w = p.w * m
}

func (p *point) Divide(m float64) {
	p.x = p.x / m
	p.y = p.y / m
	p.z = p.z / m
	p.w = p.w / m
}

func (a *point) DotProduct(t tuple) float64 {
	b := t.GetAxes()
	return (a.x * b.x + a.y * b.y + a.z * b.z + a.w * a.w) 
}

func (v *vector) Create(x, y, z float64) {
	v.x = x
	v.y = y
	v.z = z
	v.w = 0.0
}

func (v *vector) GetAxes() axes {
	return axes{v.x, v.y, v.z, v.w};
}

func (a *vector) Add(b tuple) tuple {
	b1 := b.GetAxes()
	w := a.w + b1.w 
	if ( w > 0) {
		return &point{axes{a.x + b1.x, a.y + b1.y, a.z + b1.z, w}}
	}
	return &vector{axes{a.x + b1.x, a.y + b1.y, a.z + b1.z, w}}
}

func (a *vector) Subtracting(b tuple) tuple {
	b1 := b.GetAxes()
	w := math.Abs(a.w - b1.w)
	if w > 0 {
		return &point{axes{a.x - b1.x, a.y - b1.y, a.z - b1.z, w}}
	}
	return &vector{axes{a.x - b1.x, a.y - b1.y, a.z - b1.z, w}}
}

func (v *vector) Negate() {
	v.x = -v.x
	v.y = -v.y
	v.z = -v.z
	v.w = -v.w
}

func (v *vector) Multiply(m float64) {
	v.x = v.x * m
	v.y = v.y * m
	v.z = v.z * m
	v.w = v.w * m
}

func (v *vector) Divide(m float64) {
	v.x = v.x / m
	v.y = v.y / m
	v.z = v.z / m
	v.w = v.w / m
}

func (v *vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2) + math.Pow(v.w, 2))
}

func (v *vector) Normalize() vector {
	return (vector{axes{v.x / v.Magnitude(), v.y / v.Magnitude(), v.z / v.Magnitude(), v.w / v.Magnitude()}})
}

func (a *vector) DotProduct(t tuple) float64 {
	b := t.GetAxes()
	return (a.x * b.x + a.y * b.y + a.z * b.z + a.w * a.w) 
}

func (a *vector) CrossProduct(t tuple) vector {
	b := t.GetAxes()
	return vector{axes{a.y * b.z - a.z * b.y, a.z * b.x - a.x * b.z, a.x * b.y - a.y * b.x, 0.0}}
}