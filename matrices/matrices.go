package rtmatrices

type Matrices struct {
	x, y uint8
	node [][]float64
}

func (a *Matrices) MulAux(y, x int, b *Matrices) float64 {
	return a.node[y][0] * b.node[0][x] +
		a.node[y][1] * b.node[1][x] +
		a.node[y][2] * b.node[2][x] +
		a.node[y][3] * b.node[3][x]
}

func (m1 *Matrices) Mul(m2 *Matrices) Matrices {
	var res Matrices
	for i := uint8(0); i < m1.y; i++ {
		res.node = append(res.node,
			[]float64{m1.MulAux(int(i), 0, m2),
			m1.MulAux(int(i), 1, m2),
			m1.MulAux(int(i), 2, m2),
			m1.MulAux(int(i), 3, m2)})
	}
	res.x = m1.x
	res.y = m1.y 
	return res
}