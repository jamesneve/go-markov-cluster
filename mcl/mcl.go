package mcl

import (
	"github.com/gonum/matrix/mat64"
	"errors"
	"fmt"
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

func (m *MCL) GenerateClusters(input *mat64.Dense) error {
	r, c := input.Dims()
	if r != c {
		return errors.New("Input matrix must be an adjacency matrix with an equal number of rows and columns")
	}

	m.addSelfLoops(input, r)
	fmt.Println("Initial adjacency matrix:", input)
	m.normalize(input, r)

	for i := 0; i < m.loops; i++ {
		m.takePower(input)
		m.takeInflation(input, r)
	}

	fmt.Println("Cluster discovery:", input)

	return nil
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