package generation2cubicstresspullcoefficientsourcetypeanddoubleeventweightaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2CubicStressPullCoefficientSourceTypeAndDoubleEventWeightAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 731 — Cubic Stress-Pull Coefficient Source-Type and Double-Event Weight Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate731 cubic coefficient source-type audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate730 cubic stress-pull correction", Passed: a.Gate730.Inherited && a.Gate730.SevenOver36CompressedResidual && a.Gate730.KappaEPartiallyDependent && a.Gate730.NoNativeBoundaryMomentExpansion && strings.Contains(a.Gate730.Verdict, StatusGate730CubicStressPullInherited), Detail: FormatGate730(a.Gate730)},
			{Name: "rewrite 7/36 as two times K7 event weight", Passed: near(a.DoubleEvent.CubicCoefficient, 7.0/36.0, 1e-18) && near(a.DoubleEvent.K7EventProbability, 7.0/72.0, 1e-18) && near(a.DoubleEvent.DoubleK7Weight, a.DoubleEvent.CubicCoefficient, 1e-18) && a.DoubleEvent.IdentityExact && strings.Contains(a.DoubleEvent.Verdict, StatusCubicCoeffRewrittenAsTwoTimesK7Weight), Detail: FormatDoubleEvent(a.DoubleEvent)},
			{Name: "audit boundary-pair source candidate", Passed: near(a.BoundaryPair.BoundaryPairDimension, 2, 1e-18) && near(a.BoundaryPair.BoundaryPairTimesK7Weight, a.DoubleEvent.CubicCoefficient, 1e-18) && a.BoundaryPair.EqualsCubicCoefficient && strings.Contains(a.BoundaryPair.Verdict, StatusBoundaryPairSourceCandidateAudited), Detail: FormatBoundaryPair(a.BoundaryPair)},
			{Name: "audit two-wall stress-pull source candidate", Passed: a.StressPull.TwoSidedBoundaryLegs && a.StressPull.ArbitraryFitRejected && strings.Contains(a.StressPull.Verdict, StatusTwoWallStressPullSourceCandidateAudited), Detail: FormatStressPull(a.StressPull)},
			{Name: "record kinetic-to-amplitude warning", Passed: a.KineticFactor.FactorTwoResonance && !a.KineticFactor.DerivesCubicCoeff && strings.Contains(a.KineticFactor.Verdict, StatusKineticToAmplitudeDoesNotDeriveCubicCoeff), Detail: FormatKinetic(a.KineticFactor)},
			{Name: "audit typed alternatives without arbitrary rational search", Passed: a.Alternatives.NoArbitrarySearch && a.Alternatives.ClosestAccepted && a.Alternatives.ClosestName == "2p_K7=7/36" && strings.Contains(a.Alternatives.Verdict, StatusTwoPK7BestTypedSourceForCubicCoeff), Detail: FormatAlternatives(a.Alternatives)},
			{Name: "rewrite moment polynomial with event-weight source", Passed: a.Polynomial.UsesDoubleEventForm && near(a.Polynomial.CubicTerm, a.DoubleEvent.DoubleK7Weight*a.Gate730.M3Wall, 1e-18) && strings.Contains(a.Polynomial.MomentForm, "2p_K7 M3_wall") && strings.Contains(a.Polynomial.Verdict, StatusMomentPolynomialRewrittenWithEventWeightSource), Detail: FormatPolynomial(a.Polynomial)},
			{Name: "preserve noncircularity firewall", Passed: a.NonCircular.KappaEPartiallyDependent && a.NonCircular.TwoPK7TypedButUnexplained && !a.NonCircular.BoundaryPairStressPullNative && !a.NonCircular.MomentExpansionTheoremNative && strings.Contains(a.NonCircular.Verdict, StatusNoNativeReasonCubicCoeffEqualsTwoPK7), Detail: FormatNoncircularity(a.NonCircular)},
			{Name: "preserve physical firewalls", Passed: !a.Firewall.ClaimsNativeScalarRuntime && !a.Firewall.ClaimsHiggsMassTheorem && !a.Firewall.ClaimsYukawaTheorem && !a.Firewall.ClaimsCKMPMNSTheorem && !a.Firewall.ClaimsHistoryLoopTheorem && strings.Contains(a.Firewall.Verdict, StatusGate731Boundary), Detail: FormatFirewall(a.Firewall)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
