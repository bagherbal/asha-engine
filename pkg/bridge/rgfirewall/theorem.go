package rgfirewall

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteRGFlowBoundaryScaleFirewallTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-RG-FLOW-BOUNDARY-SCALE-FIREWALL"
	const name = "finite RG flow and boundary-scale selection firewall"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build finite RG firewall", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 102 embedded boundary inherited", Passed: a.BoundaryDataDerived, Detail: fmt.Sprintf("K_*=%s; k_Y=%.10f; sin²_*=%.10f", FormatMatrix(a.EmbeddedBoundaryHessian), a.BoundaryKY, a.BoundarySin2)},
			{Name: "finite-spectrum beta diagnostic attached", Passed: a.BetaDiagnosticAvailable && !a.Beta.ImportedSMBetaTable && !a.Beta.HiddenObservedCouplingsUsed, Detail: fmt.Sprintf("b1=%.10f, b2=%.10f, b3=%.10f; slopes=(%.10f, %.10f, %.10f)", a.B1, a.B2, a.B3, a.B1Slope, a.B2Slope, a.B3Slope)},
			{Name: "formal one-loop flow family constructed", Passed: a.FormalRGFamilyConstructed, Detail: a.Flow.Convention + "; " + a.Flow.U1InverseExpression + "; " + a.Flow.SU2InverseExpression},
			{Name: "electromagnetic observables remain symbolic", Passed: a.Flow.EMInverseExpression != "" && a.Flow.Sin2Expression != "" && a.Flow.AlphaExpression != "", Detail: a.Flow.EMInverseExpression + "; " + a.Flow.Sin2Expression},
			{Name: "free variables exposed", Passed: a.FreeVariableCount >= 5 && !a.BoundaryCouplingDerived && !a.BoundaryScaleDerived, Detail: FormatMissing(a.MissingVariables)},
			{Name: "non-uniqueness witnessed", Passed: a.NonUniquenessWitnessed && !a.SampleA.Physical && !a.SampleB.Physical, Detail: FormatSample(a.SampleA) + " | " + FormatSample(a.SampleB)},
			{Name: "boundary-scale dimension firewall", Passed: a.BoundaryScaleDimensionNoGo && !a.BoundaryScaleDerived, Detail: "dimensionless finite Hessian data cannot by itself select an energy/length scale M*"},
			{Name: "threshold firewall", Passed: a.ThresholdFirewallClosed && !a.ThresholdRuleDerived && !a.BetaThresholdCorrected, Detail: fmt.Sprintf("activation candidates=%d, continuum candidates=%d, corrected beta derived=%t", a.Thresholds.CandidateCount, a.Thresholds.ContinuumFieldCandidateCount, a.Thresholds.ThresholdCorrectedBetaDerived)},
			{Name: "native finite RG theorem still sealed", Passed: !a.FiniteRGTheoremDerived && a.TwoParameterUnderdetermined, Detail: "current flow uses a stated continuum one-loop assumption and remains a two-parameter family in u and L"},
			{Name: "physical constants remain unclaimed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived, Detail: "no physical thetaW, alpha, g2, gY, W/Z masses, Higgs vev, or fermion masses are computed"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed electroweak coupling, weak angle, mass, GUT scale, or threshold scale was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			"rejected claims: " + Join(a.RejectedClaims),
			"remaining unknowns: " + Join(a.RemainingUnknowns),
			"Next: " + a.RecommendedNextGate,
		}}
	}}
}
