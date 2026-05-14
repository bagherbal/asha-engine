package sourceaction

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func SourceTensorActionTheorem() theorem.Theorem {
	const id = "MATTER-SOURCE-TENSOR-VARIATIONAL-SELECTION"
	const name = "source tensor action and variational selection"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.OpenTest, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct source tensor action", Passed: false, Detail: err.Error()}}}
		}
		status := theorem.OpenTest
		if a.NonzeroStationaryFound {
			status = theorem.BridgeRequired
		}
		actionDetails := make([]string, 0, len(a.Actions))
		for _, act := range a.Actions {
			actionDetails = append(actionDetails, FormatAction(act))
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: status, Checks: []theorem.Check{
			{Name: "source tensor variational domain", Passed: a.GenerationDimension == 3 && a.ActiveDimension == 4 && a.TensorDimension == 12, Detail: fmt.Sprintf("M:H_active→H_generation is a %dx%d tensor; dim Hom=%d", a.GenerationDimension, a.ActiveDimension, a.TensorDimension)},
			{Name: "minimal source action", Passed: a.NaturalHessianPositive && a.NaturalUniqueMinimum, Detail: "S[M]=1/2||M||²−<J,M> has positive Hessian and stationary equation M=J"},
			{Name: "geometric source term", Passed: a.NaturalSourceNorm > 1e-8, Detail: fmt.Sprintf("canonical source norm ||J||=%.6e from connection/BF/source candidates", a.NaturalSourceNorm)},
			{Name: "stationary tensor selected", Passed: a.NonzeroStationaryFound, Detail: fmt.Sprintf("||M_*||=%.6e, minimum action=%.6e", a.NaturalStationaryNorm, a.NaturalMinimumAction)},
			{Name: "zero-map theorem", Passed: a.NaturalSelectsZero, Detail: "with the current finite data, the unique stable stationary tensor is M=0"},
			{Name: "arbitrary nonzero tensor rejected", Passed: a.ArbitraryMapRejected, Detail: fmt.Sprintf("Hom space has %d entries, but arbitrary J would be fitting", a.TensorDimension)},
			{Name: "symmetry-breaking source not derived", Passed: !a.TachyonicTermDerived, Detail: "no finite theorem derives a negative quadratic term, fixed radius, or orientation for nonzero M"},
		}, Notes: []string{
			a.TruthStatement,
			"This is a variational no-go: positive source-tensor dynamics does not rescue generation mixing unless a real finite source J appears.",
			fmt.Sprintf("action audit: %s", strings.Join(actionDetails, " | ")),
			fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns),
		}}
	}}
}
