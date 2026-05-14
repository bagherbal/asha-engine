package exterior

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/combinatorics"
)

type GradeSpace struct {
	Grade     int
	Dimension int
}

type Algebra struct {
	VectorDimension int
	Grades          []GradeSpace
}

// NewAlgebra constructs the exterior algebra Λ(R^vectorDimension).
func NewAlgebra(vectorDimension int) (Algebra, error) {
	dims, err := combinatorics.GradeDimensions(vectorDimension)
	if err != nil {
		return Algebra{}, err
	}

	grades := make([]GradeSpace, len(dims))
	for k, d := range dims {
		grades[k] = GradeSpace{Grade: k, Dimension: d}
	}

	return Algebra{VectorDimension: vectorDimension, Grades: grades}, nil
}

func (a Algebra) TotalDimension() int {
	total := 0
	for _, g := range a.Grades {
		total += g.Dimension
	}
	return total
}

func (a Algebra) Grade(k int) (GradeSpace, error) {
	if k < 0 || k >= len(a.Grades) {
		return GradeSpace{}, fmt.Errorf("grade %d outside exterior algebra of dimension %d", k, a.VectorDimension)
	}
	return a.Grades[k], nil
}

func (a Algebra) MiddleGrade() (GradeSpace, error) {
	if a.VectorDimension%2 != 0 {
		return GradeSpace{}, fmt.Errorf("middle grade is unique only for even vector dimension: %d", a.VectorDimension)
	}
	return a.Grade(a.VectorDimension / 2)
}
