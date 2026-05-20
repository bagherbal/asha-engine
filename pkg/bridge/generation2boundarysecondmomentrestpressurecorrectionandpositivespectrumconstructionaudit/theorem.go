package generation2boundarysecondmomentrestpressurecorrectionandpositivespectrumconstructionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-813-BOUNDARY-SECOND-MOMENT-RESTPRESSURE-POSITIVE-SPECTRUM-CONSTRUCTION"
	theoremName = "Gate 813 — Boundary Second-Moment RestPressure Correction and Positive Spectrum Construction Audit"
)

func Generation2BoundarySecondMomentRestPressureCorrectionAndPositiveSpectrumConstructionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 811 and Gate 812 rest-pressure target", Passed: a.Inheritance.Gate811Inherited && a.Inheritance.Gate812Inherited && a.Inheritance.BoundarySecondMomentSelected && math.Abs(a.Inheritance.M2-1.624013231638281e-7) < 1e-20 && math.Abs(a.Inheritance.C2Obs-5.8299915725) < 1e-9 && containsAll(a.Inheritance.Verdicts, []string{StatusGate811Inherited, StatusGate812Inherited, StatusBoundaryM2Selected}), Detail: FormatInheritance(a.Inheritance)},
			{Name: "define exact top/rest positivity framework", Passed: a.Framework.Defined && strings.Contains(a.Framework.Formula, "N_eff") && containsAll(a.Framework.Supports, []string{StatusPositiveBetaBounds}) && containsAll(a.Framework.Failures, []string{StatusFirstOrderInsufficient}), Detail: a.Framework.Inequality},
			{Name: "reject naive alpha=(3/10)s as exact positive-rest law", Passed: a.NaiveAlpha.Audited && a.NaiveAlpha.Alpha > 0 && a.NaiveAlpha.BetaRequired < 0 && !a.NaiveAlpha.Positive && containsAll(a.NaiveAlpha.Failures, []string{StatusAlphaThreeTenNegativeBeta, StatusNaiveAlphaNotPositiveExact}), Detail: FormatAlphaCandidate(AlphaCandidate{Name: "alpha_B1", CAlpha: 0, Alpha: a.NaiveAlpha.Alpha, Beta: a.NaiveAlpha.BetaRequired, QRest: math.NaN(), ValidBeta: false, ValidQRest: false})},
			{Name: "compute positive lower-bound alpha and M2-sized correction", Passed: a.LowerBound.Computed && math.Abs(a.LowerBound.AlphaMin-0.0003878160447268) < 1e-15 && a.LowerBound.Correction > 0 && math.Abs(a.LowerBound.CorrectionOverM2-0.5086108926) < 1e-9 && containsAll(a.LowerBound.Supports, []string{StatusMinimalCorrectionOrderM2, StatusMinimalCorrectionHalfM2}) && containsAll(a.LowerBound.Failures, []string{StatusHalfM2NotExact, StatusAlphaMinFromAggregate}), Detail: "correction/M2 near half: " + FormatAlphaCandidate(a.AlphaFamily.ObservedMin)},
			{Name: "audit corrected alpha family against beta and q_rest positivity", Passed: a.AlphaFamily.Defined && a.AlphaFamily.HalfM2.Beta < 0 && !a.AlphaFamily.HalfM2.ValidBeta && a.AlphaFamily.ThreeFifths.Beta > 0 && a.AlphaFamily.ThreeFifths.ValidQRest && a.AlphaFamily.SixElevenths.Beta > 0 && a.AlphaFamily.SixElevenths.ValidQRest && containsAll(a.AlphaFamily.Supports, []string{StatusAlphaAboveHalfM2, StatusThreeFifthsPositiveCandidate, StatusSixEleventhsPositiveCandidate}) && containsAll(a.AlphaFamily.Failures, []string{StatusHalfM2StillNegative, StatusNoNativeCAlpha}), Detail: FormatAlphaFamily(a.AlphaFamily)},
			{Name: "audit direct Delta_N second-moment closure", Passed: a.DirectDelta.Audited && math.Abs(a.DirectDelta.Residual+2.76095936e-8) < 1e-15 && a.DirectDelta.ResidualImprovement > 30 && a.DirectDelta.PositiveBandExists && containsAll(a.DirectDelta.Supports, []string{StatusC2SixSharpClosure, StatusC2SixTypedCandidate, StatusDirectDeltaNeedsSpectrum, StatusDirectB2HasAbstractPositiveBand}) && containsAll(a.DirectDelta.Failures, []string{StatusDirectDeltaNotEnough, StatusC2SixNotNative}), Detail: FormatDirectDelta(a.DirectDelta)},
			{Name: "define positive rest spectrum construction and existence condition", Passed: a.RestConstruction.Defined && strings.Contains(a.RestConstruction.Condition, "q_rest") && len(a.RestConstruction.Examples) >= 3 && containsAll(a.RestConstruction.Supports, []string{StatusSpectrumExistsByQ}) && containsAll(a.RestConstruction.Failures, []string{StatusPositiveNoSectors, StatusPositiveNotNativeYukawa, StatusNoRestAtomCount}), Detail: a.RestConstruction.Condition},
			{Name: "reaudit boundary coefficient source typing", Passed: a.Coefficient.Audited && strings.Contains(a.Coefficient.NineOverFive, "9/5") && strings.Contains(a.Coefficient.Six, "6") && containsAll(a.Coefficient.Supports, []string{StatusCoefficientsTypedCandidates}) && containsAll(a.Coefficient.Failures, []string{StatusCoeffTypingNotTraceTheorem, StatusNoColorHyperchargeTraceMap}), Detail: a.Coefficient.NineOverFive + "; " + a.Coefficient.Six},
			{Name: "define exact missing BoundaryToTraceMagnitudeRestMap", Passed: a.BoundaryMap.Defined && len(a.BoundaryMap.Objects) >= 10 && strings.Contains(a.BoundaryMap.Target, "C_Yukawa") && containsAll(a.BoundaryMap.Supports, []string{StatusExactMissingMap}) && containsAll(a.BoundaryMap.Failures, []string{StatusNoBoundaryTraceRestMap, StatusNoTopSelectorFromBoundary, StatusNoRestConcentrationLaw}), Detail: strings.Join(a.BoundaryMap.Objects, "; ")},
			{Name: "compute second-moment corrected FN spurion", Passed: a.Spurion.Defined && math.Abs(a.Spurion.EpsilonN-0.21964195823344188) < 2e-15 && math.Abs(a.Spurion.EpsilonB2-0.21964260964006385) < 2e-15 && math.Abs(a.Spurion.B2Diff) < math.Abs(a.Spurion.B1Diff) && containsAll(a.Spurion.Supports, []string{StatusB2SpurionSharper}) && containsAll(a.Spurion.Failures, []string{StatusSpurionNotNative, StatusSpurionNoSectors}), Detail: FormatSpurion(a.Spurion)},
			{Name: "audit candidate C_Yukawa and C_Higgs impact without updating ledger", Passed: a.Impact.Audited && math.Abs(a.Impact.CYukawaBoundaryB2-0.9992248096849) < 1e-10 && math.Abs(a.Impact.CHiggsBoundaryB2-1.037220510859) < 1e-9 && containsAll(a.Impact.Supports, []string{StatusCertifiedMapWouldReduceSeal}) && containsAll(a.Impact.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "record partial-success outcome and next branch", Passed: a.Outcome.Recorded && strings.Contains(a.Outcome.Selected, "Outcome 2") && containsAll(a.Outcome.Supports, []string{StatusExpectedPartialSuccess}) && a.Branch.Recorded && strings.Contains(a.Branch.Next, "Gate 814") && strings.Contains(a.Branch.Next, "BoundaryToTraceMagnitudeRestMap"), Detail: a.Outcome.Selected + " -> " + a.Branch.Next},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoApproxAsTheorem && a.Firewalls.NoCoefficientFit && a.Firewalls.NoSectorAssignment && a.Firewalls.NoNativeYukawa && a.Firewalls.NoLedgerUpdate && a.Firewalls.NoPoleMass && a.Firewalls.Verdict == StatusFirewallGate813, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatInheritance(a.Inheritance), a.Framework.Inequality, FormatAlphaFamily(a.AlphaFamily), FormatDirectDelta(a.DirectDelta), a.RestConstruction.Condition, a.Coefficient.NineOverFive, a.Coefficient.Six, FormatSpurion(a.Spurion), FormatImpact(a.Impact), a.Outcome.Selected, a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
