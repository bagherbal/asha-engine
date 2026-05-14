package boundaryselector

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func BoundaryScaleOperatorAbsoluteCouplingSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-BOUNDARY-SCALE-OPERATOR-ABSOLUTE-COUPLING-SEARCH"
	const name = "boundary-scale operator and absolute coupling unit search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build boundary selector audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 103 finite boundary seed inherited", Passed: a.FiniteBoundarySeedSelected && a.RelativeKineticNormalizationComplete, Detail: fmt.Sprintf("k_Y=%.10f; sin²_*=%.10f; K_* trace=%.10f", a.BoundaryKY, a.BoundarySin2, traceEmbeddedBoundary(a.RG))},
			{Name: "candidate operator inventory", Passed: a.CandidateCount >= 8, Detail: FormatCandidateOperators(a.CandidateOperators)},
			{Name: "topological action seal audited", Passed: a.TopologicalSeal > 0 && a.InstantonWeight > 0 && a.TopologicalSealAsScaleRejected, Detail: fmt.Sprintf("S_top=%.10f; exp(-S_top)=%.3e; dimensionful scale derived=%t", a.TopologicalSeal, a.InstantonWeight, a.Coupling.DimensionfulScaleDerived)},
			{Name: "unit-trace coupling rejected as convention", Passed: a.UnitTraceConventionRejected && a.UnitTraceCouplingSq > 0, Detail: fmt.Sprintf("g_unit²=%.10f; alpha_unit^{-1}=%.10f; trace normalization derived=%t", a.UnitTraceCouplingSq, a.UnitTraceInverseAlpha, a.Coupling.TraceNormalizationDerived)},
			{Name: "all candidates are dimensionless", Passed: a.AllCandidateOperatorsDimensionless && !a.DimensionfulOperatorFound, Detail: "no candidate carries physical mass/length units"},
			{Name: "absolute coupling selector absent", Passed: !a.AbsoluteCouplingOperatorFound && !a.BoundaryCouplingDerived, Detail: "no finite-to-continuum gauge-action prefactor or trace normalization has been derived"},
			{Name: "boundary-scale selector absent", Passed: !a.BoundaryScaleOperatorFound && !a.BoundaryScaleDerived, Detail: "no M* or physical scale unit is selected by the finite dimensionless invariants"},
			{Name: "threshold selector absent", Passed: !a.ThresholdSelectorFound && !a.ThresholdRuleDerived, Detail: "threshold inventory exists, but no activation/decoupling operator is selected"},
			{Name: "residual equation nullity exposed", Passed: a.EquationAudit.Nullity >= 3 && a.EquationAudit.IndependentEquationsForPhysicalFlow == 0, Detail: FormatEquationAudit(a.EquationAudit)},
			{Name: "residual symmetries remain", Passed: len(a.ResidualSymmetries) == 3, Detail: FormatResidualSymmetries(a.ResidualSymmetries)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, thetaW, W/Z masses, Higgs vev, Higgs mass, fermion masses, or observed matching inputs are used"},
		}, Notes: []string{
			a.TruthStatement,
			"rejected claims: " + Join(a.RejectedClaims),
			"remaining unknowns: " + Join(a.RemainingUnknowns),
			"Next: " + a.RecommendedNextGate,
		}}
	}}
}
