package scalarsu2

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ScalarContactSU2ActionSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-SCALAR-CONTACT-SU2-ACTION-SEARCH"
	const name = "scalar-contact SU(2)_L action search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct scalar-contact SU(2) audit", Passed: false, Detail: err.Error()}}}
		}
		eps := 1e-9
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "four-real-dimensional scalar frame", Passed: a.ActiveRealDimension == 4 && a.ComplexDoubletDimension == 2, Detail: fmt.Sprintf("active scalar/contact directions=%d = %d complex doublet components", a.ActiveRealDimension, a.ComplexDoubletDimension)},
			{Name: "active scalar response spectrum", Passed: a.PairDegenerate, Detail: fmt.Sprintf("S_Φ spectrum=%s with pair split %.10f", FormatSpectrum(a.ActiveSpectrum), a.PairSplit)},
			{Name: "abstract SU(2) doublet representation", Passed: a.AbstractDoubletRepresentation, Detail: fmt.Sprintf("skew residual=%.3e, su(2) closure residual=%.3e", a.SkewResidual, a.SU2ClosureResidual)},
			{Name: "scalar response commutator with T1", Passed: a.ScalarResponseCommT1Norm < eps, Detail: fmt.Sprintf("||[S_Φ,T1]||_F = %.6e", a.ScalarResponseCommT1Norm)},
			{Name: "scalar response commutator with T2", Passed: a.ScalarResponseCommT2Norm < eps, Detail: fmt.Sprintf("||[S_Φ,T2]||_F = %.6e", a.ScalarResponseCommT2Norm)},
			{Name: "scalar response commutator with T3", Passed: a.ScalarResponseCommT3Norm < eps, Detail: fmt.Sprintf("||[S_Φ,T3]||_F = %.6e", a.ScalarResponseCommT3Norm)},
			{Name: "full SU(2) selected by scalar data", Passed: a.FullSU2SelectedByScalarData, Detail: fmt.Sprintf("requires all ||[S_Φ,T_a]|| < %.1e; max=%.6e", eps, a.MaxFullSU2CommNorm)},
			{Name: "commuting pair-rotation subgroup selected", Passed: a.U1PairRotationSelected, Detail: "the anisotropic pair spectrum preserves a T3-like U(1) rotation but not the full nonabelian SU(2)"},
			{Name: "canonical complex/quaternionic structure", Passed: a.CanonicalComplexStructure, Detail: "not derived; the abstract doublet action still needs a finite contact-geometric complex structure"},
			{Name: "finite covariant derivative", Passed: a.CovariantDerivativeDerived, Detail: "not derived; DΦ and kinetic normalization are still required"},
			{Name: "gauge-eating theorem", Passed: a.GaugeEatingTheoremDerived, Detail: "not derived; scalar SU(2) representation alone does not produce W/Z masses"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
