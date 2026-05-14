package exterior

import (
	"fmt"
	"reflect"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func GradeStructureTheorem(vectorDimension int) theorem.Theorem {
	id := fmt.Sprintf("ALG-EXT-GRADES-R%d", vectorDimension)
	return theorem.Theorem{
		ID:     id,
		Name:   fmt.Sprintf("Exterior grade structure of R^%d", vectorDimension),
		Layer:  theorem.LayerAlgebra,
		Status: theorem.ExactFinite,
		Verify: func() theorem.Result {
			alg, err := NewAlgebra(vectorDimension)
			if err != nil {
				return theorem.Result{
					ID: id, Name: "Exterior grade structure", Layer: theorem.LayerAlgebra, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct exterior algebra", Passed: false, Detail: err.Error()}},
				}
			}

			expectedTotal := 1 << vectorDimension
			checks := []theorem.Check{
				{
					Name:   "total exterior dimension",
					Passed: alg.TotalDimension() == expectedTotal,
					Detail: fmt.Sprintf("sum_k dim Λ^k R^%d = %d; expected 2^%d = %d", vectorDimension, alg.TotalDimension(), vectorDimension, expectedTotal),
				},
			}

			if vectorDimension == 8 {
				got := make([]int, len(alg.Grades))
				for i, g := range alg.Grades {
					got[i] = g.Dimension
				}
				want := []int{1, 8, 28, 56, 70, 56, 28, 8, 1}
				checks = append(checks, theorem.Check{
					Name:   "canonical R8 grade sequence",
					Passed: reflect.DeepEqual(got, want),
					Detail: fmt.Sprintf("grades=%v", got),
				})
				mid, _ := alg.MiddleGrade()
				checks = append(checks, theorem.Check{
					Name:   "middle chamber",
					Passed: mid.Grade == 4 && mid.Dimension == 70,
					Detail: fmt.Sprintf("Λ^%d R^8 has dimension %d", mid.Grade, mid.Dimension),
				})
			}

			return theorem.Result{
				ID:     id,
				Name:   fmt.Sprintf("Exterior grade structure of R^%d", vectorDimension),
				Layer:  theorem.LayerAlgebra,
				Status: theorem.ExactFinite,
				Checks: checks,
				Notes:  []string{"Mathematical dimensions are derived from binomial coefficients, not hard-coded constants."},
			}
		},
	}
}
