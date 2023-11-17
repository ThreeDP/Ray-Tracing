package tuples

import (
	"testing"
	"math"
)

const EPSILON = 0.00001

func isEqual(x, y, epsilon float64) bool {
	if math.Abs(x - y) < epsilon {
		return true
	}
	return false
}

func compareTuples(t *testing.T, exp, res Tuple) {
	t.Helper()
	e := exp.GetAxes();
	r := res.GetAxes();
	axes := axes{math.Abs(e.x - r.x), math.Abs(e.y - r.y), math.Abs(e.z - r.z), math.Abs(e.w - r.w)}
	if ( axes.x >= EPSILON || axes.y >= EPSILON || axes.z >= EPSILON || axes.w >= EPSILON) {
		t.Errorf("\nexpected Tuple:\t[x: %.2f, y: %.2f, z: %.2f, w: %.2f],\nresult Tuple:\t[x: %.2f, y: %.2f, z: %.2f, w: %.2f]", e.x, e.y, e.z, e.w, r.x, r.y, r.z, r.w)
	}
}
		
func TestCreate(t *testing.T) {
	t.Run("Create a Point", func(t *testing.T) {
		var p Point
		p.Create(4.0, -4.0, 3.0)
		compareTuples(t, &p, &Point{axes{4.0, -4.0, 3.0, 1.0}})
	})

	t.Run("Create a Vector", func(t *testing.T) {
		var v vector
		v.Create(4.0, -4.0, 3.0)
		compareTuples(t, &v, &vector{axes{4.0, -4.0, 3.0, 0.0}})
	})
}

func TestSumTwoTuples(t *testing.T) {
	t.Run("Adding two tuples", func(t *testing.T) {
		var a Point
		var b vector
		a.Create(3.0, -2.0, 5.0)
		b.Create(-2.0, 3.0, 1.0)
		res := a.Add(&b)
		compareTuples(t, res, &Point{axes{1.0, 1.0, 6.0, 1.0}})
		if _, ok := res.(*Point); !ok {
			t.Errorf("The type expected is Point, but has a vector\n")
		}
	})
}

func TestSubtracting(t *testing.T) {
	t.Run("Subtracting two points", func(t *testing.T) {
		var p1, p2 Point
		p1.Create(3, 2, 1)
		p2.Create(5, 6, 7)
		res := p1.Subtracting(&p2)
		compareTuples(t, res, &vector{axes{-2.0, -4.0, -6.0, 0.0}})
		if _, ok := res.(*vector); !ok {
			t.Errorf("The type expected is vector, but has a Point\n")
		}
	})

	t.Run("Subtracting a vector from a Point", func(t *testing.T) {
		var p Point
		var v vector
		p.Create(3, 2, 1)
		v.Create(5, 6, 7)
		res := p.Subtracting(&v)
		compareTuples(t, res, &Point{axes{-2.0, -4.0, -6.0, 1.0}})
		if _, ok := res.(*Point); !ok {
			t.Errorf("The type expected is Point, but has a vector\n")
		}
	})

	t.Run("Subtracting two vectors", func(t *testing.T) {
		var v1, v2 vector
		v1.Create(3, 2, 1)
		v2.Create(5, 6, 7)
		res := v1.Subtracting(&v2)
		compareTuples(t, res, &vector{axes{-2.0, -4.0, -6.0, 0.0}})
		if _, ok := res.(*vector); !ok {
			t.Errorf("The type expected is Point, but has a vector\n")
		}
	})

	t.Run("Subtracting a vector from the zero vector(0, 0, 0, 0)", func(t *testing.T) {
		var zero, v vector
		zero.Create(0, 0, 0)
		v.Create(1, -2, 3)
		res := zero.Subtracting(&v)
		compareTuples(t, res, &vector{axes{-1, 2, -3, 0}})
		if _, ok := res.(*vector); !ok {
			t.Errorf("The type expected is Point, but has a vector\n")
		}
	})
}

func TestNegate(t *testing.T) {
	t.Run("Negating a Tuple", func(t *testing.T) {
		a := Point{axes{1, -2, 3, -4}}
		a.Negate()
		compareTuples(t, &a, &Point{axes{-1, 2, -3, 4}})
	})
}

func TestMultiply(t *testing.T) {
	t.Run("Multiplying a Tuple by a scalar", func(t *testing.T) {
		a := Point{axes{1, -2, 3, -4}}
		a.Multiply(3.5)
		compareTuples(t, &a, &Point{axes{3.5, -7, 10.5, -14}})
	})
	t.Run("Multiplying a Tuple by a faction", func(t *testing.T) {
		a := Point{axes{1, -2, 3, -4}}
		a.Multiply(0.5)
		compareTuples(t, &a, &Point{axes{0.5, -1, 1.5, -2}})
	})
}

func TestDivide(t *testing.T) {
	t.Run("Dividing a Tuple by a scalar", func(t *testing.T) {
		a := Point{axes{1, -2, 3, -4}}
		a.Divide(2)
		compareTuples(t, &a, &Point{axes{0.5, -1, 1.5, -2}})
	})
}

func TestMagnitude(t *testing.T) {
	t.Run("Computing the Magnitude of a vector(0, 1, 0)", func(t *testing.T) {
		var v vector
		v.Create(0, 1, 0)
		res := v.Magnitude()
		if (!isEqual(res, 1.0, EPSILON)) {
			t.Errorf("Expected Magnitude %f, but has %f", 1.0, res)
		}
	})

	t.Run("Computing the Magnitude of vector(1, 0, 0)", func(t *testing.T) {
		var v vector
		v.Create(1,0,0)
		res := v.Magnitude()
		if (!isEqual(res, 1.0, EPSILON)) {
			t.Errorf("Expected Magnitude %f, but has %f", 1.0, res)
		}
	})

	t.Run("Computing the Magnitude of vector(0, 0, 1)", func(t *testing.T) {
		var v vector
		v.Create(0,0,1)
		res := v.Magnitude()
		if (!isEqual(res, 1.0, EPSILON)) {
			t.Errorf("Expected Magnitude %f, but has %f", 1.0, res)
		}
	})

	t.Run("Computing the Magnitude of vector(0.7071, 0.7071, 0)", func(t *testing.T) {
		var v vector
		v.Create(0.7071, 0.7071, 0)
		res := v.Magnitude()
		if (!isEqual(res, 1.0, EPSILON)) {
			t.Errorf("Expected Magnitude %f, but has %f", 1.0, res)
		}
	})

	t.Run("Computing the Magnitude of vector(1, 2, 3)", func(t *testing.T) {
		var v vector
		v.Create(1, 2, 3)
		res := v.Magnitude()
		if (!isEqual(res, math.Sqrt(14), EPSILON)) {
			t.Errorf("Expected Magnitude %f, but has %f", math.Sqrt(14), res)
		}
	})

	t.Run("Computing the Magnitude of vector(-1, -2, -3)", func(t *testing.T) {
		var v vector
		v.Create(1, 2, 3)
		res := v.Magnitude()
		if (!isEqual(res, math.Sqrt(14), EPSILON)) {
			t.Errorf("Expected Magnitude %f, but has %f", math.Sqrt(14), res)
		}
	})
}

func TestNormalize(t *testing.T) {
	t.Run("Normalizing vector(4, 0, 0) gives vector(1, 0, 0)", func(t *testing.T) {
		var v vector
		v.Create(4, 0, 0)
		res := v.Normalize()
		compareTuples(t, &res, &vector{axes{1, 0, 0, 0}})
	})

	t.Run("Normalizing vector (1, 2, 3) gives vector(0.26726, 0.53452, 0.80178)", func(t *testing.T) {
		var v vector
		v.Create(1, 2, 3)
		res := v.Normalize()
		compareTuples(t, &res, &vector{axes{0.26726, 0.53452, 0.80178, 0}})
	})

	t.Run("The Magnitude of a normalized vector", func(t *testing.T) {
		var v vector
		v.Create(1, 2, 3)
		norm := v.Normalize()
		res := norm.Magnitude()
		if (!isEqual(res, 1.0, EPSILON)) {
			t.Errorf("Expected Magnitude %f, but has %f\n", math.Sqrt(14), res)
		}
	})
}

func TestDotProduct(t *testing.T) {
	t.Run("The dot product of two tuples", func(t *testing.T) {
		var a, b vector
		a.Create(1, 2, 3)
		b.Create(2, 3, 4)
		res := a.DotProduct(&b)
		if !isEqual(res, 20.0, EPSILON) {
			t.Errorf("Expected dot product %f, but has %f\n", 20.0, res)
		}
	})
}

func TestCrossProduct(t *testing.T) {
	var a, b vector
	a.Create(1, 2, 3)
	b.Create(2, 3, 4)
	t.Run("The cross product of vector(1, 2, 3) and vector(2, 3, 4)", func(t *testing.T) {
		res := a.CrossProduct(&b)
		compareTuples(t, &res, &vector{axes{-1, 2, -1, 0}})
	})

	t.Run("The cross product of vector(2, 3, 4) and vector(1, 2, 3)", func(t *testing.T) {
		res := b.CrossProduct(&a)
		compareTuples(t, &res, &vector{axes{1, -2, 1, 0}})
	})
}