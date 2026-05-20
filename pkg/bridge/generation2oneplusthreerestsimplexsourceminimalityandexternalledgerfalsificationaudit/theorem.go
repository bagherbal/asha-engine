package generation2oneplusthreerestsimplexsourceminimalityandexternalledgerfalsificationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-819-ONE-PLUS-THREE-REST-SIMPLEX-SOURCE-MINIMALITY-EXTERNAL-LEDGER-FALSIFICATION"
	theoremName = "Gate 819 — OnePlusThree RestSimplex Source Minimality and External Ledger Falsification Audit"
)

func Generation2OnePlusThreeRestSimplexSourceMinimalityAndExternalLedgerFalsificationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 818 simplex ledger", Passed: math.Abs(a.Ledger.AlphaB-0.0003878958469680527) < 1e-18 && math.Abs(a.Ledger.QSimplex-0.3330749367196054) < 1e-15 && containsAll(a.Ledger.Supports, []string{SupportSharpestSimplex, SupportBoundaryAlphaNatural}), Detail: FormatLedger(a.Ledger)},
			{Name: "define minimal one-plus-three source seal", Passed: !a.SourceSeal.CurrentSupplied && len(a.SourceSeal.Components) == 10 && containsAll(a.SourceSeal.Failures, []string{FailureNoNativeSimplex}), Detail: FormatSourceSeal(a.SourceSeal)},
			{Name: "audit projective, K7, boundary and external lanes", Passed: len(a.Candidates) == 4 && !a.Candidates[0].SuppliesTraceReadout && !a.Candidates[1].SuppliesTraceReadout && a.Candidates[2].SuppliesAlpha && a.Candidates[3].SuppliesSectorAtoms, Detail: FormatCandidates(a.Candidates)},
			{Name: "define external ledger falsification protocol", Passed: a.Protocol.CanUpgradeExternalR3 && len(a.Protocol.PrimaryTests) == 5 && strings.Contains(strings.Join(a.Protocol.PrimaryTests, " "), "c2_ext") && containsAll(a.Protocol.Supports, []string{SupportExternalFalsifies, SupportLedgerUpgradesR3}), Detail: FormatProtocol(a.Protocol)},
			{Name: "enforce noncircularity for simplex test", Passed: a.NonCircular.Enforced && len(a.NonCircular.FrozenCoefficients) == 3 && containsAll(a.NonCircular.Failures, []string{FailureShapeMustNotBeForced, FailureCoefficientsFrozen}), Detail: strings.Join(a.NonCircular.Forbidden, " | ")},
			{Name: "classify R-status without native promotion", Passed: strings.Contains(a.Status.Level, "strengthened partial R2") && a.Status.ExternalR3Ready && !a.Status.NativeSourceFound && !a.Status.CanUpdateCYukawa, Detail: a.Status.Outcome + " — " + a.Status.Level},
			{Name: "record C_Yukawa/C_Higgs impact without ledger update", Passed: math.Abs(a.Impact.CYukawaCandidate-0.9992248096922658) < 1e-15 && math.Abs(a.Impact.CHiggsCandidate-1.0372205108665146) < 2e-15 && containsAll(a.Impact.Failures, []string{FailureNoUpdateCYukawa, FailureCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "record next branch", Passed: strings.Contains(a.Branch.NextGate, "Gate 820") && strings.Contains(a.Branch.Reason, "trace-ledger"), Detail: a.Branch.NextGate},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.SimplexNotNative && a.Firewalls.ProjectiveNotYukawa && a.Firewalls.K7NotYukawa && a.Firewalls.BoundaryAlphaNotYukawa && a.Firewalls.AbstractNotSector && a.Firewalls.ExternalNotNative && a.Firewalls.NoCYukawaUpdate && a.Firewalls.CHiggsLevelB && a.Firewalls.TreeProxyNotPole && a.Firewalls.Verdict == StatusFirewallGate819, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatSourceSeal(a.SourceSeal), FormatCandidates(a.Candidates), FormatProtocol(a.Protocol), a.Status.Outcome, a.Status.Level, FormatImpact(a.Impact), a.Branch.NextGate, a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
