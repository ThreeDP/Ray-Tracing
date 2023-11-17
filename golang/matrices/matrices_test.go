package rtmatrices

import (
	"testing"
	"math"
	"rt/tuples"
)

// func TestCreatematrix(t *testing.T) {
// 	t.Run("Constructing and inspecting a 4x4 matrix", func(t *testing.T) {

// 	})
// }

const EPSILON = 0.00001

func isEqual(x, y, epsilon float64) bool {
	if math.Abs(x - y) < epsilon {
		return true
	}
	return false
}

func comapreMatrix(t *testing.T, res, exp *Matrices) {
	t.Helper()
	if (res.x != exp.x && res.y != exp.y) {
		t.Error("\nmatrix of different size\n")
	}
	for i := uint8(0); i < exp.y; i++ {
		for j := uint8(0); j < exp.x; j++ {
			if (!isEqual(res.node[i][j], exp.node[i][j], EPSILON)) {
				t.Errorf("\nexpected:\n" +
					"\t[%f, %f, %f, %f]\n" +
					"\t[%f, %f, %f, %f]\n" + 
					"\t[%f, %f, %f, %f]\n" +
					"\t[%f, %f, %f, %f]\n" +
					"\nbut has:\n" +
					"\t[%f, %f, %f, %f]\n" +
					"\t[%f, %f, %f, %f]\n" +
					"\t[%f, %f, %f, %f]\n" +
					"\t[%f, %f, %f, %f]\n\n",
					exp.node[0][0], exp.node[0][1], exp.node[0][2], exp.node[0][3],
					exp.node[1][0], exp.node[1][1], exp.node[1][2], exp.node[1][3],
					exp.node[2][0], exp.node[2][1], exp.node[2][2], exp.node[2][3],
					exp.node[3][0], exp.node[3][1], exp.node[3][2], exp.node[3][3],
					res.node[0][0], res.node[0][1], res.node[0][2], res.node[0][3],
					res.node[1][0], res.node[1][1], res.node[1][2], res.node[1][3],
					res.node[2][0], res.node[2][1], res.node[2][2], res.node[2][3],
					res.node[3][0], res.node[3][1], res.node[3][2], res.node[3][3])
					
			}
		}
	}
}

func TestMatrixOperations(t *testing.T) {
	t.Run("Multiplying two matrices", func(t *testing.T) {
		a := Matrices{
			x: 4,
			y: 4,
			node: [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2}}}
		b := Matrices{
			x: 4,
			y: 4,
			node: [][]float64{
				{-2, 1, 2, 3},
				{3, 2, 1, -1},
				{4, 3, 6, 5},
				{1, 2, 7, 8}}}
		exp := Matrices{
			x: 4,
			y: 4,
			node: [][]float64{
				{20, 22, 50, 48},
				{44, 54, 114, 108},
				{40, 58, 110, 102},
				{16, 26, 46, 42}}}
		res := a.Mul(&b)
		comapreMatrix(t, &res, &exp)
	})

	t.Run("A matrix multiplied by a tuple", func(t *testing.T) {
		a := Matrices{
			x: 4,
			y: 4,
			node: [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2}}}
		var b tuples.Point
		b.Create(1, 2, 3)
		if (a.MulByObj(b).Comp(&Point{Axes{18, 24, 33, 1}}) == 0) {
			
		}
	})
}