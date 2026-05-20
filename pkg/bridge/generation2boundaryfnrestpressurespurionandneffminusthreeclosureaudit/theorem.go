package generation2boundaryfnrestpressurespurionandneffminusthreeclosureaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-810-BOUNDARY-FN-RESTPRESSURE-SPURION-NEFF-MINUS-THREE-CLOSURE"
	theoremName = "Gate 810 — Boundary-FN RestPressure Spurion and N_eff-Minus-Three Closure Audit"
)

func Generation2BoundaryFNRestPressureSpurionAndNEffMinusThreeClosureAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 809 Boundary-FN candidate", Passed: a.Inheritance.Gate809Inherited && a.Inheritance.CandidateSelected && math.Abs(a.Inheritance.DeltaN-DeltaN) < 1e-15 && containsAll(a.Inheritance.Verdicts, []string{StatusGate809Inherited, StatusBoundaryFNSelected}), Detail: "Boundary-FN rest-pressure candidate selected"},
			{Name: "compute direct boundary closure residual", Passed: a.Direct.Computed && math.Abs(a.Direct.CObs-1.8007325638446063) < 5e-14 && math.Abs(a.Direct.Residual-9.4679834536684e-7) < 2e-18 && math.Abs(a.Direct.RelativeResidual-0.0004068143483817) < 5e-16 && containsAll(a.Direct.Supports, []string{StatusDeltaNApproxNineFifths, StatusBoundaryScaleCapable}) && containsAll(a.Direct.Failures, []string{StatusNineFifthsNotExact, StatusNumericalNotTheorem}), Detail: FormatDirect(a.Direct)},
			{Name: "define boundary-FN spurion and compute epsilons", Passed: a.Spurion.Defined && math.Abs(a.Spurion.EpsilonN-0.21964195823344188) < 3e-15 && math.Abs(a.Spurion.EpsilonB-0.21961961644976352) < 1e-15 && math.Abs(a.Spurion.RelativeDifference-0.0001017191062078) < 5e-16 && containsAll(a.Spurion.Supports, []string{StatusBoundarySpurionScale, StatusEpsilonClose}) && containsAll(a.Spurion.Failures, []string{StatusEpsilonBNotNative, StatusNoBoundarySpurionMap, StatusCabibboNotTheorem}), Detail: FormatSpurion(a.Spurion)},
			{Name: "audit 9/5 coefficient source", Passed: a.Coefficient.Audited && strings.Contains(a.Coefficient.Coefficient, "9/5") && containsAll(a.Coefficient.Supports, []string{StatusNineFifthsTypedCandidate, StatusFiveThirdsNonarbitrary}) && containsAll(a.Coefficient.Failures, []string{StatusNoColorHyperchargeTheorem, StatusInverseHyperchargeNotAuto, StatusNoRationalFit}), Detail: FormatCoeff(a.Coefficient)},
			{Name: "compute top/rest alpha boundary closure", Passed: a.Alpha.Computed && math.Abs(a.Alpha.AlphaApprox-0.0003878912453691245) < 1e-15 && math.Abs(a.Alpha.AlphaBoundary-0.00038773344564488885) < 1e-15 && math.Abs(a.Alpha.Residual-1.5779972422780667e-7) < 1e-18 && containsAll(a.Alpha.Supports, []string{StatusAlphaApproxThreeTenths, StatusThreeTenthsTypedCandidate}) && containsAll(a.Alpha.Failures, []string{StatusThreeTenthsNotNative, StatusNoAlphaReadoutMap, StatusAlphaApproxNotExact}), Detail: FormatAlpha(a.Alpha)},
			{Name: "perform exact top/rest positivity audit", Passed: a.Positivity.Completed && math.Abs(a.Positivity.BetaRequired-(-1.651341154285823e-7)) < 3e-16 && math.Abs(a.Positivity.AlphaMinOverS-0.3000639091748843) < 5e-14 && a.Positivity.BetaRequired < 0 && containsAll(a.Positivity.Supports, []string{StatusThreeTenthsCloseNotExact, StatusCorrectionAboveThreeTenths}) && containsAll(a.Positivity.Failures, []string{StatusAlphaThreeTenthsNotExact, StatusFirstOrderNotTheorem, StatusPositiveBlocksExact}), Detail: FormatPositivity(a.Positivity)},
			{Name: "audit rest concentration regimes", Passed: a.Concentration.Audited && len(a.Concentration.Regimes) == 3 && math.Abs(a.Concentration.Regimes[1].QRest-0.33307493962706697) < 5e-10 && math.Abs(a.Concentration.Regimes[2].AlphaOverS-0.3002387347866694) < 5e-13 && containsAll(a.Concentration.Supports, []string{StatusNarrowCorridor, StatusConcentrationControls}) && containsAll(a.Concentration.Failures, []string{StatusAggregateNoQRest, StatusNoRestAtoms, StatusConcentrationNotNative}), Detail: FormatConcentration(a.Concentration)},
			{Name: "define BoundaryFNRestPressureMap requirement", Passed: a.Map.Defined && strings.Contains(a.Map.Name, "BoundaryFNRestPressureMap") && containsAll(a.Map.Chain, []string{"s", "Delta_N=N_eff-3", "C_Yukawa=3/N_eff"}) && containsAll(a.Map.Supports, []string{StatusMapWouldReduceNEff, StatusMapPreciseMissing}) && containsAll(a.Map.Failures, []string{StatusNoBoundaryFNMap, StatusNoSectorTraceRule, StatusNoPositiveConcentration, StatusNoScaleStability}), Detail: FormatMap(a.Map)},
			{Name: "define alternative coefficient controls", Passed: a.Controls.Defined && len(a.Controls.Controls) == 4 && strings.Contains(a.Controls.BestTyped, "c=9/5") && math.Abs(a.Controls.Controls[3].AbsResidual-9.4679834536684e-7) < 2e-18 && containsAll(a.Controls.Supports, []string{StatusBestTypedCandidate}) && containsAll(a.Controls.Failures, []string{StatusLowDenominatorNotTheorem, StatusBestNumericalNeedsType}), Detail: FormatControls(a.Controls)},
			{Name: "audit C_Yukawa/C_Higgs boundary-FN rewrite candidate", Passed: a.CHiggs.Defined && math.Abs(a.CHiggs.CYukawaCandidate-0.9992251339916449) < 1e-15 && math.Abs(a.CHiggs.CYukawaResidual-3.151104440712871e-7) < 2e-18 && math.Abs(a.CHiggs.CHiggsCandidate-1.0372208474974351) < 1e-15 && containsAll(a.CHiggs.Supports, []string{StatusCertifiedMapReducesSeal}) && containsAll(a.CHiggs.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB, StatusCandidateNotLevelC}), Detail: FormatCHiggs(a.CHiggs)},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.Next, "Gate 811") && strings.Contains(a.Branch.Next, "Hypercharge-Color") && containsAll(a.Branch.Supports, []string{StatusNextHyperchargeColor}), Detail: a.Branch.Next},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoNativeNEff && a.Firewalls.NoNativeFN && a.Firewalls.NoRationalFit && a.Firewalls.NoAlphaExact && a.Firewalls.NoYukawaSpectrum && a.Firewalls.NoCertifiedRewrite && a.Firewalls.NoLevelC && a.Firewalls.NoPoleMass && a.Firewalls.NoD4KoideGJ && a.Firewalls.Verdict == StatusFirewallGate810, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatDirect(a.Direct), FormatSpurion(a.Spurion), FormatCoeff(a.Coefficient), FormatAlpha(a.Alpha), FormatPositivity(a.Positivity), FormatConcentration(a.Concentration), FormatMap(a.Map), FormatControls(a.Controls), FormatCHiggs(a.CHiggs), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
