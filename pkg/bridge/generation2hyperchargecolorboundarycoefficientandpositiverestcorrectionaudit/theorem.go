package generation2hyperchargecolorboundarycoefficientandpositiverestcorrectionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-811-HYPERCHARGE-COLOR-BOUNDARY-COEFFICIENT-POSITIVE-REST-CORRECTION"
	theoremName = "Gate 811 — Hypercharge-Color Boundary Coefficient and Positive-Rest Correction Audit"
)

func Generation2HyperchargeColorBoundaryCoefficientAndPositiveRestCorrectionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 810 hypercharge-color coefficient target", Passed: a.Inheritance.Gate810Inherited && a.Inheritance.CoeffTargetSelected && a.Inheritance.CorrectionTargetSelected && math.Abs(a.Inheritance.DeltaN-DeltaN) < 1e-15 && containsAll(a.Inheritance.Verdicts, []string{StatusGate810Inherited, StatusCoeffTargetSelected, StatusRestCorrectionTarget}), Detail: "Gate 810 closure and positivity target inherited"},
			{Name: "audit 9/5 coefficient factorization", Passed: a.NineFifths.Audited && strings.Contains(a.NineFifths.Expression, "9/5") && strings.Contains(a.NineFifths.Expression, "3/5") && containsAll(a.NineFifths.Verdicts, []string{StatusNineFiveFactor, StatusLedgerExistence}) && containsAll(a.NineFifths.Supports, []string{StatusNineFiveTypedCandidate, StatusRestCoeffColorHypercharge}) && containsAll(a.NineFifths.Failures, []string{StatusExistenceNotTheorem, StatusNoColorHyperchargeMap, StatusInverseHyperchargeNotYukawa}), Detail: FormatFactor(a.NineFifths)},
			{Name: "audit 3/10 coefficient factorization", Passed: a.ThreeTenths.Audited && strings.Contains(a.ThreeTenths.Expression, "3/10") && strings.Contains(a.ThreeTenths.Expression, "1/2") && containsAll(a.ThreeTenths.Supports, []string{StatusThreeTenthsTypedCandidate, StatusHalfBoundaryPairCandidate}) && containsAll(a.ThreeTenths.Failures, []string{StatusNoBoundaryAverageTheorem, StatusThreeTenthsNotNative, StatusBoundaryPairNotReadout}), Detail: FormatFactor(a.ThreeTenths)},
			{Name: "compute exact positive-rest correction scale", Passed: a.Correction.Defined && math.Abs(a.Correction.DeltaAlpha-8.2599081954e-8) < 5e-17 && math.Abs(a.Correction.DeltaAlphaOverS-0.0000639091748843) < 5e-14 && math.Abs(a.Correction.DeltaAlphaOverAlphaB-0.0002130) < 5e-7 && math.Abs(a.Correction.M2-1.624013231638281e-7) < 1e-20 && math.Abs(a.Correction.HalfM2-8.120066158191404e-8) < 1e-20 && containsAll(a.Correction.Supports, []string{StatusCorrectionOrderM2, StatusHalfM2Approx, StatusCorrectionMayBeM2}) && containsAll(a.Correction.Failures, []string{StatusCorrectionNotExactHalfM2, StatusNoPositiveCorrectionTheorem, StatusM2NotConcentrationLaw}), Detail: FormatCorrection(a.Correction)},
			{Name: "test corrected alpha candidate with half-M2", Passed: a.AlphaCorr.Defined && a.AlphaCorr.Tested && !a.AlphaCorr.BetaNonnegative && !a.AlphaCorr.QRestValid && math.Abs(a.AlphaCorr.AlphaCorr-0.00038781464630647074) < 2e-18 && math.Abs(a.AlphaCorr.AlphaCorrMinusAlphaMin-(-1.398420347890287e-9)) < 5e-18 && math.Abs(a.AlphaCorr.BetaCorr-(-2.795756404161409e-9)) < 5e-17 && containsAll(a.AlphaCorr.Supports, []string{StatusAlphaHalfM2LawfulToTest}) && containsAll(a.AlphaCorr.Failures, []string{StatusAlphaCorrNeedsBeta, StatusAlphaCorrNotNative, StatusQRestNotSector, StatusQRestNotNativeSpectrum, StatusNoAtomLedgerFromQRest}), Detail: FormatAlphaCorr(a.AlphaCorr)},
			{Name: "compute direct Delta_N second-moment correction", Passed: a.DeltaCorr.Defined && a.DeltaCorr.Computed && math.Abs(a.DeltaCorr.C2Obs-5.8299915722461693) < 5e-13 && math.Abs(a.DeltaCorr.CandidateResidual-(-2.7609593616136768e-8)) < 5e-18 && math.Abs(a.DeltaCorr.CandidateRelErr-(-1.1863116249617649e-5)) < 5e-16 && containsAll(a.DeltaCorr.Supports, []string{StatusDeltaResidualOrderM2, StatusC2ObsCloseToSix}) && containsAll(a.DeltaCorr.Failures, []string{StatusC2ObsNotNative, StatusSecondMomentNeedsMap, StatusResidualFittingForbidden}), Detail: FormatDeltaCorr(a.DeltaCorr)},
			{Name: "define hypercharge-color boundary rest-pressure package", Passed: a.Package.Defined && strings.Contains(a.Package.Name, "HyperchargeColor") && len(a.Package.Components) >= 6 && containsAll(a.Package.Supports, []string{StatusPackageSharpestCandidate}) && containsAll(a.Package.Failures, []string{StatusPackageNeedsTraceMap, StatusPackageNeedsSpectrum, StatusPackageNeedsScale}), Detail: FormatPackage(a.Package)},
			{Name: "preserve alternative coefficient controls", Passed: a.Controls.Defined && len(a.Controls.Controls) == 5 && strings.Contains(a.Controls.Controls[4].Name, "6ps") && math.Abs(a.Controls.Controls[4].AbsResidual-2.7609593616136768e-8) < 5e-18 && containsAll(a.Controls.Supports, []string{StatusCorrectedModelControl}) && containsAll(a.Controls.Failures, []string{StatusRationalClosenessForbidden, StatusModelNeedsSourcePositivity}), Detail: FormatControls(a.Controls)},
			{Name: "audit C_Yukawa and C_Higgs candidate impact", Passed: a.Impact.Defined && math.Abs(a.Impact.CYukawaBoundary-0.9992248096922658) < 1e-15 && math.Abs(a.Impact.CYukawaResidual-(-9.188935057302672e-9)) < 5e-18 && math.Abs(a.Impact.CHiggsBoundary-1.0372205108665145) < 1e-15 && containsAll(a.Impact.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoCoeffShortcut && a.Firewalls.NoAlphaExact && a.Firewalls.NoHalfM2Native && a.Firewalls.NoSixFit && a.Firewalls.NoPackageNative && a.Firewalls.NoBoundarySpectrum && a.Firewalls.NoQRestSector && a.Firewalls.NoCertifiedRewrite && a.Firewalls.NoPoleMass && a.Firewalls.Verdict == StatusFirewallGate811, Detail: a.Firewalls.Verdict},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.Next, "Gate 812") && containsAll(a.Branch.Supports, []string{StatusNextSecondMomentCorrection}), Detail: a.Branch.Next},
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
		notes := []string{a.Truth, FormatFactor(a.NineFifths), FormatFactor(a.ThreeTenths), FormatCorrection(a.Correction), FormatAlphaCorr(a.AlphaCorr), FormatDeltaCorr(a.DeltaCorr), FormatPackage(a.Package), FormatControls(a.Controls), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
