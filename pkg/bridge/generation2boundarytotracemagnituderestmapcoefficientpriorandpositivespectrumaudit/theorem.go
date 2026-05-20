package generation2boundarytotracemagnituderestmapcoefficientpriorandpositivespectrumaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-816-BOUNDARY-TO-TRACE-MAGNITUDE-REST-MAP-COEFFICIENT-PRIOR-POSITIVE-SPECTRUM"
	theoremName = "Gate 816 — BoundaryToTraceMagnitudeRestMap Coefficient-Prior and Positive-Spectrum Construction Audit"
)

func Generation2BoundaryToTraceMagnitudeRestMapCoefficientPriorAndPositiveSpectrumAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "recompute inherited Boundary-FN numerical ledger", Passed: math.Abs(a.Ledger.DeltaNBFN-0.002327375081808316) < 1e-18 && math.Abs(a.Ledger.ResidualBFN+2.7609593616e-8) < 1e-16 && math.Abs(a.Ledger.C2Obs-5.8299915725) < 1e-9 && math.Abs(a.Ledger.EpsilonBFN-0.2196426096400638) < 1e-15 && containsAll(a.Ledger.Verdicts, []string{StatusGate815Inherited, StatusNumericalLedger}), Detail: FormatLedger(a.Ledger)},
			{Name: "audit 9/5 and 6 coefficient priors before residual use", Passed: a.Coeff95.Classification != "" && a.Coeff6.Classification != "" && containsAll(a.Coeff95.Supports, []string{StatusCoeff95Prior, StatusCoeffPriorBridge}) && containsAll(a.Coeff6.Supports, []string{StatusCoeff6Prior, StatusCoeffPriorBridge}) && containsAll(a.Coeff95.Failures, []string{StatusCoeff95NotTheorem}) && containsAll(a.Coeff6.Failures, []string{StatusCoeff6NotTheorem}), Detail: FormatCoefficient(a.Coeff95) + " :: " + FormatCoefficient(a.Coeff6)},
			{Name: "compute alpha candidate positivity table", Passed: len(a.AlphaRows) == 4 && containsAlpha(a.AlphaRows, "1/2", false, false) && containsAlpha(a.AlphaRows, "3/5", true, true) && containsAlpha(a.AlphaRows, "6/11", true, true) && containsAlpha(a.AlphaRows, "1", true, true), Detail: FormatAlphaRows(a.AlphaRows)},
			{Name: "verify half-M2 correction remains slightly beta-negative", Passed: alphaByName(a.AlphaRows, "1/2").AgainstInherited.BetaRequired < 0 && alphaByName(a.AlphaRows, "1/2").AgainstBFN.BetaRequired < 0 && math.Abs(alphaByName(a.AlphaRows, "1/2").AgainstInherited.QRestRequired+0.00619626039) < 1e-9, Detail: FormatAlphaRows(a.AlphaRows)},
			{Name: "compute beta/q_rest and Delta_N reconstruction table", Passed: len(a.QRestCandidates) == 6 && len(a.DeltaRows) == 24 && math.Abs(a.DeltaRows[0].ResidualToBFN) < 1e-8 && strings.Contains(FormatDeltaRows(a.DeltaRows, 4), "alpha=1"), Detail: FormatDeltaRows(a.DeltaRows, 6)},
			{Name: "classify positive-spectrum realizability level", Passed: a.Positive.AbstractExistenceAny && strings.Contains(a.Positive.Level, "partial R2") && containsAll(a.Positive.Supports, []string{StatusAbstractPositive, StatusNoMapConstructed, StatusLevelPartialR2}) && containsAll(a.Positive.Failures, []string{StatusPositiveNoSectors, StatusNoTraceAtoms, StatusNoYukawaOperator}), Detail: strings.Join(a.Positive.BestRows, " | ") + " :: " + a.Positive.Level},
			{Name: "prove coefficient-prior package is not a BoundaryToTraceMagnitudeRestMap", Passed: a.NoGo.PackageDefined && a.NoGo.ProducesScalarClosure && !a.NoGo.ConstructsAlphaBetaQ && containsAll(a.NoGo.Failures, []string{StatusCoeffPackageNoMap, StatusScalarNotTraceMap}), Detail: strings.Join(a.NoGo.Ingredients, " | ")},
			{Name: "enforce noncircularity requirements", Passed: a.Protocol.Enforced && len(a.Protocol.Forbidden) == 7 && len(a.Protocol.Allowed) == 6 && containsAll(a.Protocol.Failures, []string{StatusNoRetune, StatusNoTopSolve, StatusNoHiggsData, StatusQRestChosenNotTheorem, StatusNoFNCharges}), Detail: strings.Join(a.Protocol.Forbidden, " | ")},
			{Name: "record C_Yukawa and C_Higgs candidate impact without update", Passed: a.Impact.Recorded && math.Abs(a.Impact.CYukawaBFN-0.9992248096922658) < 1e-15 && math.Abs(a.Impact.CHiggsBFN-1.0372205108665146) < 1e-15 && containsAll(a.Impact.Failures, []string{StatusGate816NoCYukawaUpdate, StatusCHiggsLevelB, StatusTreeProxyNotPole}), Detail: FormatImpact(a.Impact)},
			{Name: "classify outcome branch", Passed: a.Branch.Recorded && strings.Contains(a.Branch.Outcome, "partial success") && strings.Contains(a.Branch.NextGate, "Gate 817") && containsAll(a.Branch.Supports, []string{StatusNextNoGo}), Detail: a.Branch.Outcome + " -> " + a.Branch.NextGate},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.CoeffPriorNotTheorem && a.Firewalls.ScalarClosureNotMap && a.Firewalls.PositiveExistenceNotLedger && a.Firewalls.SectorLedgerNotNative && a.Firewalls.BoundaryNotYukawa && a.Firewalls.FNLikeNotChargeOperator && a.Firewalls.HyperchargeNotRestLaw && a.Firewalls.ColorNotGeneration && a.Firewalls.CHiggsLevelB && a.Firewalls.TreeProxyNotPole && a.Firewalls.Verdict == StatusFirewallGate816, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatCoefficient(a.Coeff95), FormatCoefficient(a.Coeff6), FormatAlphaRows(a.AlphaRows), FormatDeltaRows(a.DeltaRows, 8), strings.Join(a.Positive.BestRows, " | "), a.Positive.Level, strings.Join(a.NoGo.Ingredients, " | "), FormatImpact(a.Impact), a.Branch.Outcome, a.Branch.NextGate, a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func containsAlpha(rows []AlphaCandidate, name string, inhPos, bfnPos bool) bool {
	for _, r := range rows {
		if r.Name == name {
			return r.AgainstInherited.PositiveCompatible == inhPos && r.AgainstBFN.PositiveCompatible == bfnPos
		}
	}
	return false
}

func alphaByName(rows []AlphaCandidate, name string) AlphaCandidate {
	for _, r := range rows {
		if r.Name == name {
			return r
		}
	}
	return AlphaCandidate{}
}
