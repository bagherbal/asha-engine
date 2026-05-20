package generation2selfconsistentrestconcentrationlawandboundaryalphamapaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-817-SELF-CONSISTENT-REST-CONCENTRATION-LAW-BOUNDARY-ALPHA-MAP"
	theoremName = "Gate 817 — Self-Consistent Rest Concentration Law and Boundary Alpha Map Audit"
)

func Generation2SelfConsistentRestConcentrationLawAndBoundaryAlphaMapAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "define boundary alpha map and exact 6 alpha compression", Passed: math.Abs(a.Ledger.AlphaB-0.0003878958469680527) < 1e-18 && math.Abs(a.Ledger.DeltaNBFN-a.Ledger.SixAlphaB) < 1e-18 && a.Alpha.DeltaEqualsSixAlphaExactly && containsAll(a.Alpha.Supports, []string{StatusAlphaSourceShape, StatusDeltaEqualsSixAlpha}) && containsAll(a.Alpha.Failures, []string{StatusAlphaNotTheorem}), Detail: FormatLedger(a.Ledger) + " :: " + FormatAlpha(a.Alpha)},
			{Name: "prove exact beta identity and self-consistent q law", Passed: math.Abs(a.Closure.BetaBySimplifiedIdentity-1.5034655049e-7) < 1e-16 && math.Abs(a.Closure.BetaIdentityResidual) < 1e-15 && math.Abs(a.Closure.QRest-(1.0/a.Ledger.NEffBFN)) < 1e-18 && math.Abs(a.Closure.InverseQRest-a.Ledger.NEffBFN) < 1e-15 && containsAll(a.Closure.Supports, []string{StatusBetaExactIdentity, StatusQInverseNEff, StatusPositiveQ}) && containsAll(a.Closure.Failures, []string{StatusQNoIndependentSource}), Detail: FormatClosure(a.Closure)},
			{Name: "audit abstract positive rest spectrum constructions", Passed: len(a.Spectra) == 4 && spectrumByName(a.Spectra, "diffuse three-rest construction").Exact == false && spectrumByName(a.Spectra, "concentrated one-rest construction").Exact == false && spectrumByName(a.Spectra, "mixed four-rest construction small-support branch").Exact && spectrumByName(a.Spectra, "mixed four-rest construction heavy-support branch").Exact && math.Abs(spectrumByName(a.Spectra, "mixed four-rest construction small-support branch").Q-a.Ledger.QRestB) < 1e-15, Detail: FormatSpectra(a.Spectra)},
			{Name: "classify q_rest as self-consistency rather than independent theorem", Passed: !a.SelfAudit.IndependentSourceCertified && strings.Contains(a.SelfAudit.Classification, "self-consistent") && containsAll(a.SelfAudit.Failures, []string{StatusQNoIndependentSource, StatusQMayBeAlgebraic}), Detail: a.SelfAudit.Classification + " :: " + strings.Join(a.SelfAudit.CandidateSourceLanes, " | ")},
			{Name: "update BoundaryToTraceMagnitudeRestMap status", Passed: a.Map.ConstructsAlphaBetaQ && a.Map.ConstructsPositiveSpectrum && !a.Map.ConstructsSectorLedger && !a.Map.NativeYukawaTheorem && strings.Contains(a.Map.Level, "partial R2") && containsAll(a.Map.Failures, []string{StatusNoTraceAtoms, StatusNoR3, StatusNoR4}), Detail: a.Map.CandidateExpression + " :: " + a.Map.Level},
			{Name: "record C_Yukawa and C_Higgs candidate impact without official update", Passed: a.Impact.Recorded && math.Abs(a.Impact.CYukawaBFN-0.9992248096922658) < 1e-15 && math.Abs(a.Impact.CHiggsBFN-1.0372205108665146) < 1e-15 && containsAll(a.Impact.Failures, []string{StatusNoUpdate, StatusCHiggsLevelB, StatusTreeProxyNotPole}), Detail: FormatImpact(a.Impact)},
			{Name: "record outcome and next branch", Passed: a.Branch.Recorded && strings.Contains(a.Branch.Outcome, "partial R2") && strings.Contains(a.Branch.NextGate, "Gate 818"), Detail: a.Branch.Outcome + " -> " + a.Branch.NextGate},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.DeltaSixAlphaNotTheorem && a.Firewalls.QRestNotTheorem && a.Firewalls.PositiveNotLedger && a.Firewalls.SectorLedgerNotNative && a.Firewalls.BoundaryNotYukawa && a.Firewalls.FNLikeNotChargeOperator && a.Firewalls.HyperchargeNotRestLaw && a.Firewalls.ColorNotGeneration && a.Firewalls.CHiggsLevelB && a.Firewalls.TreeProxyNotPole && a.Firewalls.Verdict == StatusFirewallGate817, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatAlpha(a.Alpha), FormatClosure(a.Closure), FormatSpectra(a.Spectra), a.SelfAudit.Classification, a.Map.Level, FormatImpact(a.Impact), a.Branch.Outcome, a.Branch.NextGate, a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func spectrumByName(rows []SpectrumConstruction, name string) SpectrumConstruction {
	for _, r := range rows {
		if r.Name == name {
			return r
		}
	}
	return SpectrumConstruction{}
}
