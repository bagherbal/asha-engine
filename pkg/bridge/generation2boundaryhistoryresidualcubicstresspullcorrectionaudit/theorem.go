package generation2boundaryhistoryresidualcubicstresspullcorrectionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryHistoryResidualCubicStressPullCorrectionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 730 — Boundary-History Residual Cubic Stress-Pull Correction Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate730 cubic stress-pull correction audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate729 second-moment residual", Passed: a.Gate729.Inherited && near(a.Gate729.P_K7, 7.0/72.0, 1e-18) && a.Gate729.KappaEPartiallyDependent && a.Gate729.NoNativeSecondOrderBoundaryTheorem && strings.Contains(a.Gate729.Verdict, StatusGate729SecondMomentResidualInherited), Detail: FormatGate729(a.Gate729)},
			{Name: "compute cubic wall moment and coefficient ratio", Passed: near(a.CubicMoment.M3Wall, a.Gate729.P_K7*math.Pow(a.Gate729.SSplit, 3), 1e-18) && near(a.CubicMoment.NegativeE2OverM3, -a.Gate729.E2Residual/a.CubicMoment.M3Wall, 1e-18) && a.CubicMoment.SecondResidualCubicScale && strings.Contains(a.CubicMoment.Verdict, StatusCubicWallMomentComputed), Detail: FormatCubicMoment(a.CubicMoment)},
			{Name: "audit typed cubic coefficient candidates", Passed: a.Coefficients.SevenOver36Closest && a.Coefficients.NoArbitrarySearch && a.Coefficients.ClosestName == "7/36" && strings.Contains(a.Coefficients.Verdict, StatusTypedCubicCoefficientCandidatesAudited), Detail: FormatCoefficients(a.Coefficients)},
			{Name: "test 7/36 cubic stress-pull correction", Passed: near(a.CubicCorr.QuadraticTerm, a.Gate729.KappaE*a.Gate729.M2Wall, 1e-18) && near(a.CubicCorr.CubicStressPullTerm, (7.0/36.0)*a.CubicMoment.M3Wall, 1e-18) && near(a.CubicCorr.ResidualAfterCubicCorrection, a.Gate729.EWall-a.CubicCorr.CombinedCorrection, 1e-18) && a.CubicCorr.ImprovesSecondOrderResidual && a.CubicCorr.NotExact && strings.Contains(a.CubicCorr.Verdict, StatusSevenOverThirtySixCompressesResidual), Detail: FormatCubicCorrection(a.CubicCorr)},
			{Name: "propagate cubic correction to scalar runtime residual", Passed: near(a.Runtime.CubicCorrectedRuntimeResidual, a.Gate729.LambdaProxy*a.Gate729.L*a.CubicCorr.ResidualAfterCubicCorrection, 1e-18) && a.Runtime.CompressedToNearFloatScale && a.Runtime.ImprovesSecondOrderRuntime && strings.Contains(a.Runtime.Verdict, StatusRuntimeCompressedByTypedCubicCorrection), Detail: FormatRuntime(a.Runtime)},
			{Name: "record source-type interpretation without native moment theorem", Passed: !a.SourceType.MomentExpansionTheoremNative && strings.Contains(a.SourceType.Verdict, StatusSourceTypeInterpretationRecorded) && strings.Contains(a.SourceType.Verdict, StatusNoNativeBoundaryMomentExpansionTheorem), Detail: FormatSourceType(a.SourceType)},
			{Name: "preserve noncircularity firewall", Passed: a.NonCircular.DBaseContainsKappaE && a.NonCircular.KappaEUsedAsQuadraticCoeff && a.NonCircular.CubicCoeffTypedButUnexplained && !a.NonCircular.NativeExpansionTheorem && strings.Contains(a.NonCircular.Verdict, StatusKappaEQuadraticCoefficientDependent), Detail: FormatNoncircularity(a.NonCircular)},
			{Name: "preserve physical theorem firewalls", Passed: !a.Firewall.ClaimsNativeBoundaryHistory && !a.Firewall.ClaimsNativeMomentExpansion && !a.Firewall.ClaimsNativeScalarRuntime && !a.Firewall.ClaimsHiggsMassTheorem && !a.Firewall.ClaimsYukawaTheorem && !a.Firewall.ClaimsCKMPMNSTheorem && strings.Contains(a.Firewall.Verdict, StatusGate730Boundary), Detail: FormatFirewall(a.Firewall)},
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
