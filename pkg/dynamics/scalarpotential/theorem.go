package scalarpotential

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func EffectivePotentialTheorem() theorem.Theorem {
	const id = "DYN-SCALAR-EFFECTIVE-POTENTIAL"
	const name = "finite scalar-sector effective potential normal form"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerDynamics, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerDynamics, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct scalar effective potential audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerDynamics, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "active scalar doublet dimension", Passed: a.ActiveRealDimension == 4 && a.ComplexDoubletDimension == 2, Detail: fmt.Sprintf("%d real active directions = %d complex scalar doublet components", a.ActiveRealDimension, a.ComplexDoubletDimension)},
			{Name: "protected direction resonance", Passed: a.ProtectedDirectionCount == 3, Detail: fmt.Sprintf("protected contact directions=%d; resonance with three would-be broken gauge directions, not yet a gauge-eating theorem", a.ProtectedDirectionCount)},
			{Name: "pair-degenerate scalar spectrum", Passed: a.PairDegenerate, Detail: fmt.Sprintf("active spectrum=%s; pair residual=%.3e", formatSlice(a.ActiveSpectrum), a.PairDegeneracyResidual)},
			{Name: "finite vacuum radius", Passed: a.VacuumRadiusSquared > 0, Detail: fmt.Sprintf("r0²=Tr(M_K)=%.10f, r0=%.10f", a.VacuumRadiusSquared, a.VacuumRadius)},
			{Name: "quartic shape invariant", Passed: a.LambdaShape > 0, Detail: fmt.Sprintf("λ_shape=Tr(M_K²)/Tr(M_K)²=%.10f, Tr(M_K²)=%.10f", a.LambdaShape, a.QuarticTrace)},
			{Name: "shifted normal form", Passed: a.ShiftedNormalFormAvailable, Detail: fmt.Sprintf("%s; expanded quadratic coefficient=%.10f, constant=%.10f", a.PotentialForm, a.NormalFormQuadraticCoeff, a.NormalFormConstant)},
			{Name: "dimensionless radial curvature", Passed: a.DimensionlessRadialMassSq > 0, Detail: fmt.Sprintf("m_radial²_hat=8λ_shape r0²=%.10f; dimensionless only", a.DimensionlessRadialMassSq)},
			{Name: "finite tachyonic mass sign", Passed: a.FiniteTachyonicMassDerived, Detail: "not derived as an exact finite action sign; negative coefficient appears only after choosing the shifted normal form"},
			{Name: "electroweak scale bridge", Passed: a.ElectroweakScaleDerived, Detail: "not derived; no comparison to v=246 GeV is allowed in this gate"},
			{Name: "Higgs mass bridge", Passed: a.HiggsMassDerived, Detail: "not derived; no comparison to 125 GeV is allowed in this gate"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("pair split Δλ=%.10f between high pair %.10f and low pair %.10f", a.PairSplitting, a.HighPairEigenvalue, a.LowPairEigenvalue),
			fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns),
		}}
	}}
}

func formatSlice(values []float64) string {
	out := "["
	for i, v := range values {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%.10g", v)
	}
	out += "]"
	return out
}
