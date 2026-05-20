package generation2boundaryrawmomentresponsepolynomialclosureaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryRawMomentResponsePolynomialClosureAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 733 — Boundary Raw-Moment Response Polynomial Closure Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate733 polynomial closure audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate732 raw moment coordinate audit", Passed: a.Gate732.Inherited && a.Gate732.RawMomentCoordinateActive && a.Gate732.RawM3BestCompression && a.Gate732.NoNativeRawMomentTheorem && strings.Contains(a.Gate732.Verdict, StatusGate732RawMomentCoordinateInherited), Detail: FormatGate732(a.Gate732)},
			{Name: "define cubic raw moment response polynomial", Passed: near(a.Polynomial.LeadingTerm, a.Gate732.M1Wall, 1e-18) && near(a.Polynomial.QuadraticTerm, a.Gate732.KappaE*a.Gate732.M2Wall, 1e-18) && near(a.Polynomial.CubicTerm, -2*a.Gate732.P_K7*a.Gate732.M3Wall, 1e-18) && strings.Contains(a.Polynomial.FactoredFormula, "p_K7 S") && strings.Contains(a.Polynomial.Verdict, StatusCubicRawMomentResponsePolynomialDefined), Detail: FormatPolynomial(a.Polynomial)},
			{Name: "compute cubic polynomial closure residual", Passed: near(a.Closure.Residual, a.Gate732.DBase-a.Polynomial.Value, 1e-18) && near(a.Closure.Residual, a.Gate732.RawCubicResidual, 1e-18) && a.Closure.StrongCompression && a.Closure.CompressionFactor > 1000 && strings.Contains(a.Closure.Verdict, StatusCubicPolynomialStronglyCompressesResidual), Detail: FormatClosure(a.Closure)},
			{Name: "compute fourth order required coefficient without typed source", Passed: near(a.FourthOrder.M4Wall, a.Gate732.P_K7*a.Gate732.SSplit*a.Gate732.SSplit*a.Gate732.SSplit*a.Gate732.SSplit, 1e-18) && near(a.FourthOrder.RequiredCoeff, a.Closure.Residual/a.FourthOrder.M4Wall, 1e-18) && !a.FourthOrder.TypedSourceFound && !a.FourthOrder.PromoteFourthOrder && strings.Contains(a.FourthOrder.Verdict, StatusNoTypedFourthOrderCoefficientSource), Detail: FormatFourthOrder(a.FourthOrder)},
			{Name: "audit stop condition", Passed: !a.Stop.ProjectorPowersSupplyNewDirections && a.Stop.HigherMomentsOnlyScalarPowers && !a.Stop.TypedFourthOrderSourceFound && a.Stop.StoppingAtCubicMoreLawful && strings.Contains(a.Stop.Verdict, StatusStoppingAtCubicMoreLawfulThanUntypedM4Fit), Detail: FormatStop(a.Stop)},
			{Name: "record polynomial source type", Passed: strings.Contains(a.SourceType.Compact, "pS") && strings.Contains(a.SourceType.Cubic, "double-K7") && strings.Contains(a.SourceType.Verdict, StatusPolynomialSourceTypeRecorded), Detail: FormatSourceType(a.SourceType)},
			{Name: "audit generating function candidate", Passed: a.Generating.CandidateTruncationSupported && !a.Generating.NativeGeneratingFunction && strings.Contains(a.Generating.GWallTruncation, "...") && strings.Contains(a.Generating.Verdict, StatusNoNativeBoundaryResponseGeneratingFunction), Detail: FormatGenerating(a.Generating)},
			{Name: "propagate cubic polynomial residual to scalar runtime", Passed: near(a.Runtime.RuntimeResidual, a.Runtime.LambdaProxy*a.Runtime.L*a.Closure.Residual, 1e-18) && a.Runtime.NearEliminated && strings.Contains(a.Runtime.Verdict, StatusScalarRuntimeResidualPropagatedCubicPolynomial), Detail: FormatRuntime(a.Runtime)},
			{Name: "preserve noncircularity firewall", Passed: a.NonCircular.KappaEPartiallyDependent && !a.NonCircular.DoubleK7CoefficientNative && !a.NonCircular.BoundaryMomentExpansionNative && strings.Contains(a.NonCircular.Verdict, StatusKappaECoefficientPartiallyDependent) && strings.Contains(a.NonCircular.Verdict, StatusNoNativeReasonForDoubleK7CubicCoefficient), Detail: FormatNoncircularity(a.NonCircular)},
			{Name: "preserve physical firewalls", Passed: !a.Firewall.ScalarRuntimeTheoremNative && !a.Firewall.HiggsMassTheoremNative && !a.Firewall.YukawaTheoremNative && strings.Contains(a.Firewall.Verdict, StatusGate733Boundary), Detail: FormatFirewall(a.Firewall)},
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
