package generation2bottomcolorresttripletcandidateandalphabyukawaratiofalsificationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-820-BOTTOM-COLOR-REST-TRIPLET-CANDIDATE-ALPHAB-YUKAWA-RATIO-FALSIFICATION"
	theoremName = "Gate 820 — BottomColor RestTriplet Candidate and AlphaB Yukawa-Ratio Falsification Audit"
)

func Generation2BottomColorRestTripletCandidateAndAlphaBYukawaRatioFalsificationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 819 simplex ledger", Passed: math.Abs(a.Ledger.AlphaB-0.0003878958469680527) < 1e-18 && math.Abs(a.Ledger.QSimplex-0.3330749367196054) < 1e-15, Detail: FormatLedger(a.Ledger)},
			{Name: "derive triplet and dust ratios", Passed: math.Abs(a.Ledger.BOverT-0.0003877453837799576) < 1e-18 && math.Abs(a.Ledger.SqrtBOverT-0.019691251452864992) < 1e-16 && math.Abs(a.Ledger.DustOverT-4.513895642851888e-7) < 1e-19, Detail: FormatLedger(a.Ledger)},
			{Name: "define interpretation as diagnostic only", Passed: a.Interpretation.DiagnosticOnly && containsAll(a.Interpretation.Supports, []string{SupportSimplexAsTriplet, SupportAlphaPredictsBOverT}) && containsAll(a.Interpretation.Failures, []string{FailureTripletNotNative, FailureAbstractNoSector}), Detail: FormatInterpretation(a.Interpretation)},
			{Name: "audit bottom, charm, abstract, and failure candidates", Passed: len(a.Candidates) == 4 && !a.Candidates[0].AllowedToIdentifyNow && !a.Candidates[1].AllowedToIdentifyNow && !a.Candidates[2].AllowedToIdentifyNow && a.Candidates[3].AllowedToIdentifyNow, Detail: FormatCandidates(a.Candidates)},
			{Name: "define external ledger falsification tests", Passed: a.Protocol.CanFalsify && a.Protocol.CanUpgradeExternalR3 && len(a.Protocol.Tests) == 5 && strings.Contains(strings.Join(a.Protocol.Tests, " "), "B_f/T"), Detail: FormatProtocol(a.Protocol)},
			{Name: "audit native source lanes without promotion", Passed: len(a.NativeSources) == 6 && containsAll(a.NativeSources[1].Failures, []string{FailureProjectiveNotTheorem}) && containsAll(a.NativeSources[2].Failures, []string{FailureK7NotTriplet}) && containsAll(a.NativeSources[3].Failures, []string{FailureAlphaNotRatioTheorem}), Detail: FormatNativeSources(a.NativeSources)},
			{Name: "classify strengthened partial R2 status", Passed: strings.Contains(a.Status.Level, "strengthened partial R2") && !a.Status.NativeSourceFound && !a.Status.ExternalLedgerSupplied && !a.Status.CanUpdateCYukawa, Detail: a.Status.Outcome + " — " + a.Status.Level},
			{Name: "preserve C_Yukawa and C_Higgs firewall", Passed: math.Abs(a.Impact.CandidateCYukawa-0.9992248096922658) < 1e-15 && math.Abs(a.Impact.CandidateCHiggs-1.0372205108665146) < 2e-15 && containsAll(a.Impact.Failures, []string{FailureNoUpdateCYukawa, FailureCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.TripletNotNative && a.Firewalls.BottomNotAllowed && a.Firewalls.CharmNotAllowed && a.Firewalls.ProjectiveNotTheorem && a.Firewalls.K7NotTriplet && a.Firewalls.AlphaNotRatioTheorem && a.Firewalls.AbstractNotSector && a.Firewalls.ExternalNotNative && a.Firewalls.NoCYukawaUpdate && a.Firewalls.CHiggsLevelB && a.Firewalls.TreeProxyNotPole && a.Firewalls.Verdict == StatusFirewallGate820, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatInterpretation(a.Interpretation), FormatCandidates(a.Candidates), FormatProtocol(a.Protocol), FormatNativeSources(a.NativeSources), a.Status.Outcome, a.Status.Level, FormatImpact(a.Impact), a.Branch.NextGate, a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
