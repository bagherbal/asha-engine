package generation2lowscaleyukawadustcapstresstestandsectorreadingdowngradeaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-822-LOW-SCALE-YUKAWA-DUST-CAP-STRESS-TEST-SECTOR-READING-DOWNGRADE"
	theoremName = "Gate 822 — Low-Scale Yukawa Dust-Cap Stress Test and Sector-Reading Downgrade Audit"
)

func Generation2LowScaleYukawaDustCapStressTestAndSectorReadingDowngradeAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 821 dust-capacity ledger", Passed: math.Abs(a.Ledger.AlphaB-0.0003878958469680527) < 1e-18 && math.Abs(a.Ledger.BOverT-0.0003877453837799576) < 1e-18 && math.Abs(a.Ledger.SqrtBOverT-0.019691251452864992) < 1e-16, Detail: FormatLedger(a.Ledger)},
			{Name: "compute colored and uncolored dust caps", Passed: math.Abs(a.Ledger.ExtraColoredTraceCap-1.5046318809506294e-7) < 1e-20 && math.Abs(a.Ledger.ExtraColoredYukawaCap-a.Ledger.AlphaB) < 1e-18 && math.Abs(a.Ledger.UncoloredTraceCap-4.513895642851889e-7) < 1e-19 && math.Abs(a.Ledger.UncoloredYukawaCap-0.0006718553149936293) < 1e-16, Detail: FormatLedger(a.Ledger)},
			{Name: "define bottom, charm, and abstract stress tests", Passed: len(a.StressRules) == 3 && strings.Contains(a.StressRules[0].Name, "bottom") && strings.Contains(a.StressRules[1].Name, "charm") && strings.Contains(a.StressRules[2].Name, "abstract"), Detail: FormatStressRules(a.StressRules)},
			{Name: "define external low-scale ratio ledger requirements", Passed: containsAll(a.Requirement.Verdicts, []string{StatusExternalLedgerReq}) && strings.Contains(FormatRequirement(a.Requirement), "y_b/y_t") && strings.Contains(FormatRequirement(a.Requirement), "neutrino convention"), Detail: FormatRequirement(a.Requirement)},
			{Name: "define kill-switch protocol", Passed: a.Protocol.CanFalsify && containsAll(a.Protocol.Verdicts, []string{StatusTestProtocol, StatusKillSwitch}) && strings.Contains(strings.Join(a.Protocol.KillSwitch, " "), "literal low-scale sector simplex is falsified") && strings.Contains(strings.Join(a.Protocol.Tests, " "), "q_rest"), Detail: FormatProtocol(a.Protocol)},
			{Name: "preserve native source firewalls", Passed: len(a.Native) == 7 && containsAll(a.Native[1].Failures, []string{FailureProjectiveNotTheorem}) && containsAll(a.Native[2].Failures, []string{FailureK7NotTraceAtom}) && containsAll(a.Native[3].Failures, []string{FailureAlphaNotSector}) && containsAll(a.Native[4].Failures, []string{FailureBFNNotOperator}), Detail: FormatNative(a.Native)},
			{Name: "select no-ledger outcome C", Passed: strings.Contains(a.Status.Outcome, "Outcome C") && strings.Contains(a.Status.Level, "strengthened partial R2") && !a.Status.ExternalLedgerSupplied && !a.Status.CanUpdateCYukawa, Detail: a.Status.Outcome + " — " + a.Status.Level},
			{Name: "preserve C_Yukawa and C_Higgs firewall", Passed: math.Abs(a.Impact.CandidateCYukawa-0.9992248096922658) < 1e-15 && math.Abs(a.Impact.CandidateCHiggs-1.0372205108665146) < 2e-15 && containsAll(a.Impact.Failures, []string{FailureNoUpdateCYukawa, FailureCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.BottomMatchAlone && a.Firewalls.CharmMatchAlone && a.Firewalls.ExtraColoredCap && a.Firewalls.UncoloredCap && a.Firewalls.ExternalNotNative && a.Firewalls.LowScaleNotAggregate && a.Firewalls.HighScaleNeedsRG && a.Firewalls.NoCYukawaUpdate && a.Firewalls.CHiggsLevelB && a.Firewalls.TreeProxyNotPole && a.Firewalls.Verdict == StatusFirewallGate822, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatStressRules(a.StressRules), FormatRequirement(a.Requirement), FormatProtocol(a.Protocol), FormatNative(a.Native), a.Status.Outcome, a.Status.Level, FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
