package mcl

import (
	"github.com/gonum/matrix/mat64"
	"errors"
	"math"
)

type MCL struct {
	power int
	inflation int
	loops int
}

func NewMCL(p, i, l int) MCL {
	return MCL{
		power: p,
		inflation: i,
		loops: l,
	}
}

func (m *MCL) GenerateClusters(input *mat64.Dense) (*[][]int, error) {
	r, c := input.Dims()
	if r != c {
		return nil, errors.New("Input matrix must be an adjacency matrix with an equal number of rows and columns")
	}

	m.addSelfLoops(input, r)
	m.normalize(input, r)

	for i := 0; i < m.loops; i++ {
		m.takePower(input)
		m.takeInflation(input, r)
	}

	clusters := m.getClusters(input, r)

	return &clusters, nil
}

func (m *MCL) addSelfLoops(input *mat64.Dense, size int) {
	for i := 0; i < size; i++ {
		input.Set(i, i, 1)
	}
}

func (m *MCL) normalize(input *mat64.Dense, size int) {
	for i := 0; i < size; i++ {
		var a float64
		for j := 0; j < size; j++ {
			a += input.At(i, j)
		}

		for j := 0; j < size; j++ {
			ijx := input.At(i, j) / a
			input.Set(i, j, ijx)
		}
	}
}

func (m *MCL) takePower(input *mat64.Dense) {
	if m.power <= 1 {
		return
	}

	input.Pow(input, m.power)
}

func (m *MCL) takeInflation(input *mat64.Dense, size int) {
	if m.inflation <= 1 {
		return
	}

	m.powValues(input, size)
	for i := 0; i < size; i++ {
		var a float64
		for j := 0; j < size; j++ {
			a += input.At(i, j)
		}

		for j := 0; j < size; j++ {
			ijx := input.At(i,j) / a
			input.Set(i, j, ijx)
		}
	}
}

func (m *MCL) powValues(input *mat64.Dense, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			inf := math.Pow(input.At(i, j), float64(m.inflation))
			input.Set(i, j, inf)
		}
	}
}

func (m *MCL) getClusters(input *mat64.Dense, size int) [][]int {
	clusters := make([][]int, 0, size)
	for j := 0; j < size; j++ {
		cluster := make([]int, 0, size)
		for i := 0; i < size; i++ {
			if input.At(i, j) != float64(0) {
				cluster = append(cluster, i)
			}
		}

		if len(cluster) > 1 && !m.clustersContains(clusters, cluster) {
			clusters = append(clusters, cluster)
		}
	}

	return clusters
}

func (m *MCL) clustersContains(clusters [][]int, newCluster []int) bool {
	for _, c := range clusters {
		if m.arraysAreEqual(c, newCluster) {
			return true
		}
	}
	return false
}

func (m *MCL) arraysAreEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}