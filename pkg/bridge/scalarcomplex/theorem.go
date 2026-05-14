package scalarcomplex

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ScalarComplexQuaternionicStructureSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-SCALAR-COMPLEX-QUATERNIONIC-STRUCTURE-SEARCH"
	const name = "scalar-contact complex/quaternionic structure search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct scalar complex/quaternionic audit", Passed: false, Detail: err.Error()}}}
		}
		eps := 1e-9
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "active scalar/contact pair frame", Passed: a.ActiveRealDimension == 4 && a.PairDegenerate, Detail: fmt.Sprintf("active=%d real directions, pair planes=%d, pair split=%.10f", a.ActiveRealDimension, a.PairPlaneCount, a.PairSplit)},
			{Name: "pair-compatible complex candidate", Passed: a.PairCompatibleComplexAvailable, Detail: fmt.Sprintf("J²=-I residual=%.3e, skew residual=%.3e, ||[S_Φ,J]||=%.3e", a.ComplexSquareResidual, a.ComplexSkewResidual, a.ComplexCommutesWithScalar)},
			{Name: "canonical complex orientation", Passed: a.CanonicalComplexDerived, Detail: fmt.Sprintf("not derived; scalar pair data leaves %d sign/orientation choices across two active planes", a.PairOrientationChoices)},
			{Name: "abstract quaternionic triple", Passed: a.QuaternionicTripleAvailable, Detail: fmt.Sprintf("square residual=%.3e, closure residual=%.3e", a.QuaternionicSquareResidual, a.QuaternionicClosureResidual)},
			{Name: "quaternionic triple selected by scalar response", Passed: a.QuaternionicTripleSelected, Detail: fmt.Sprintf("commutators with S_Φ: I=%.6e, J=%.6e, K=%.6e", a.QuaternionicCommIWithScalar, a.QuaternionicCommJWithScalar, a.QuaternionicCommKWithScalar)},
			{Name: "full scalar SU(2)_L recovered", Passed: a.FullScalarSU2Recovered, Detail: fmt.Sprintf("requires max quaternionic scalar commutator < %.1e; max=%.6e", eps, a.MaxQuaternionicScalarComm)},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
