package tuples

import (
	"math"
	"rt/defines"
)

type Axes struct {
	x, y, z, w float64
}

type Point struct {
	Axes
}

type Vector struct {
	Axes
}

type Tuple interface {
	Create(x, y, z float64)
	GetAxes() Axes
	Add(b Tuple) Tuple
	Subtracting(b Tuple) Tuple
	Negate()
	Multiply(m float64)
	Divide(m float64)
	DotProduct(t Tuple) float64
	Comp(t Tuple) bool
}

func (p *Point) Create(x, y, z float64) {
	p.x = x
	p.y = y
	p.z = z
	p.w = 1.0
}

func (p *Point) GetAxes() Axes {
	return Axes{p.x, p.y, p.z, p.w};
}

func (a *Point) Add(b Tuple) Tuple {
	b1 := b.GetAxes()
	w := a.w + b1.w 
	if w > 0 {
		return &Point{Axes{a.x + b1.x, a.y + b1.y, a.z + b1.z, w}}
	}
	return &Vector{Axes{a.x + b1.x, a.y + b1.y, a.z + b1.z, w}}
}

func (a *Point) Comp(b Tuple) bool {
	e := a.GetAxes();
	r := b.GetAxes();
	axes := Axes{math.Abs(e.x - r.x), math.Abs(e.y - r.y), math.Abs(e.z - r.z), math.Abs(e.w - r.w)}
	if ( axes.x >= defines.EPSILON || axes.y >= defines.EPSILON || axes.z >= defines.EPSILON || axes.w >= defines.EPSILON) {
		return false
	}
	return true
}

func (a *Point) Subtracting(b Tuple) Tuple {
	b1 := b.GetAxes()
	w := math.Abs(a.w - b1.w)
	if w != 0 {
		return &Point{Axes{a.x - b1.x, a.y - b1.y, a.z - b1.z, w}}
	}
	return &Vector{Axes{a.x - b1.x, a.y - b1.y, a.z - b1.z, w}}
}

func (p *Point) Negate() {
	p.x = -p.x
	p.y = -p.y
	p.z = -p.z
	p.w = -p.w
}

func (p *Point) Multiply(m float64) {
	p.x = p.x * m
	p.y = p.y * m
	p.z = p.z * m
	p.w = p.w * m
}

func (p *Point) Divide(m float64) {
	p.x = p.x / m
	p.y = p.y / m
	p.z = p.z / m
	p.w = p.w / m
}

func (a *Point) DotProduct(t Tuple) float64 {
	b := t.GetAxes()
	return (a.x * b.x + a.y * b.y + a.z * b.z + a.w * a.w) 
}

func (v *Vector) Create(x, y, z float64) {
	v.x = x
	v.y = y
	v.z = z
	v.w = 0.0
}

func (v *Vector) GetAxes() Axes {
	return Axes{v.x, v.y, v.z, v.w};
}

func (a *Vector) Comp(b Tuple) bool {
	e := a.GetAxes();
	r := b.GetAxes();
	axes := Axes{math.Abs(e.x - r.x), math.Abs(e.y - r.y), math.Abs(e.z - r.z), math.Abs(e.w - r.w)}
	if ( axes.x >= defines.EPSILON || axes.y >= defines.EPSILON || axes.z >= defines.EPSILON || axes.w >= defines.EPSILON) {
		return false
	}
	return true
}

func (a *Vector) Add(b Tuple) Tuple {
	b1 := b.GetAxes()
	w := a.w + b1.w 
	if ( w > 0) {
		return &Point{Axes{a.x + b1.x, a.y + b1.y, a.z + b1.z, w}}
	}
	return &Vector{Axes{a.x + b1.x, a.y + b1.y, a.z + b1.z, w}}
}

func (a *Vector) Subtracting(b Tuple) Tuple {
	b1 := b.GetAxes()
	w := math.Abs(a.w - b1.w)
	if w > 0 {
		return &Point{Axes{a.x - b1.x, a.y - b1.y, a.z - b1.z, w}}
	}
	return &Vector{Axes{a.x - b1.x, a.y - b1.y, a.z - b1.z, w}}
}

func (v *Vector) Negate() {
	v.x = -v.x
	v.y = -v.y
	v.z = -v.z
	v.w = -v.w
}

func (v *Vector) Multiply(m float64) {
	v.x = v.x * m
	v.y = v.y * m
	v.z = v.z * m
	v.w = v.w * m
}

func (v *Vector) Divide(m float64) {
	v.x = v.x / m
	v.y = v.y / m
	v.z = v.z / m
	v.w = v.w / m
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2) + math.Pow(v.w, 2))
}

func (v *Vector) Normalize() Vector {
	return (Vector{Axes{v.x / v.Magnitude(), v.y / v.Magnitude(), v.z / v.Magnitude(), v.w / v.Magnitude()}})
}

func (a *Vector) DotProduct(t Tuple) float64 {
	b := t.GetAxes()
	return (a.x * b.x + a.y * b.y + a.z * b.z + a.w * a.w) 
}

func (a *Vector) CrossProduct(t Tuple) Vector {
	b := t.GetAxes()
	return Vector{Axes{a.y * b.z - a.z * b.y, a.z * b.x - a.x * b.z, a.x * b.y - a.y * b.x, 0.0}}
}