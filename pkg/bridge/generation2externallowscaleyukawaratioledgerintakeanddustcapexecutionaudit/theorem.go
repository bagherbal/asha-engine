package generation2externallowscaleyukawaratioledgerintakeanddustcapexecutionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-823-EXTERNAL-LOW-SCALE-YUKAWA-RATIO-LEDGER-INTAKE-DUST-CAP-EXECUTION"
	theoremName = "Gate 823 — External Low-Scale Yukawa Ratio Ledger Intake and Dust-Cap Execution Audit"
)

func Generation2ExternalLowScaleYukawaRatioLedgerIntakeAndDustCapExecutionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 822 dust caps", Passed: math.Abs(a.Ledger.AlphaB-0.0003878958469680527) < 1e-18 && math.Abs(a.Ledger.SqrtBOverT-0.019691251452864992) < 1e-16 && math.Abs(a.Ledger.ExtraColoredCap-a.Ledger.AlphaB) < 1e-18 && math.Abs(a.Ledger.UncoloredCap-0.0006718553149936293) < 1e-16, Detail: FormatLedger(a.Ledger)},
			{Name: "search active ledger and return DATA_REQUIRED when absent", Passed: !a.Search.Found && !a.Search.ConventionLocked && containsAll(a.Search.Verdicts, []string{StatusDataRequiredLedger}) && containsAll(a.Search.Failures, []string{FailureNoLedgerFound, FailureNoLedger}), Detail: FormatSearch(a.Search)},
			{Name: "define convention requirements", Passed: strings.Contains(FormatSearch(a.Search), "scale_mu") && strings.Contains(FormatSearch(a.Search), "r_b") && strings.Contains(FormatSearch(a.Search), "neutrino_convention"), Detail: FormatSearch(a.Search)},
			{Name: "define branch execution tests", Passed: len(a.Branches) == 3 && strings.Contains(a.Branches[0].Name, "bottom") && strings.Contains(a.Branches[1].Name, "charm") && strings.Contains(a.Branches[2].Name, "abstract") && containsAll(a.Branches[0].Failures, []string{FailureBottomMatchAlone}), Detail: FormatBranches(a.Branches)},
			{Name: "define violation margin diagnostics and downgrade protocol", Passed: containsAll(a.Protocol.Verdicts, []string{StatusMarginsDefined, StatusDowngradeRule, StatusHighScaleFirewall}) && strings.Contains(FormatProtocol(a.Protocol), "c2") == false && strings.Contains(FormatProtocol(a.Protocol), "y_f/y_t <= alpha_B") && strings.Contains(FormatProtocol(a.Protocol), "high-scale"), Detail: FormatProtocol(a.Protocol)},
			{Name: "freeze literal sector assignment without ledger", Passed: a.Status.DataRequired && a.Status.LiteralSectorFrozen && !a.Status.ExternalR3 && !a.Status.CanUpdateCYukawa && containsAll(a.Status.Verdicts, []string{StatusDataRequiredLedger}), Detail: a.Status.Outcome + " — " + a.Status.Level},
			{Name: "preserve native source audit", Passed: len(a.Native) == 7 && strings.Contains(FormatNative(a.Native), "finite spectral triple") && strings.Contains(FormatNative(a.Native), "D4/triality"), Detail: FormatNative(a.Native)},
			{Name: "preserve C_Yukawa and C_Higgs firewall", Passed: math.Abs(a.Impact.CandidateCYukawa-0.9992248096922658) < 1e-15 && math.Abs(a.Impact.CandidateCHiggs-1.0372205108665146) < 2e-15 && math.Abs(a.Impact.OfficialCYukawa-CYukawa) < 1e-18 && containsAll(a.Impact.Failures, []string{FailureNoUpdateCYukawa, FailureCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoInferenceWithoutLedger && a.Firewalls.NoIncompleteLedger && a.Firewalls.NoMatchAlone && a.Firewalls.DustOverflow && a.Firewalls.NoHighScaleEscape && a.Firewalls.NoCYukawaUpdate && a.Firewalls.CHiggsLevelB && a.Firewalls.ExternalNotNative && a.Firewalls.Verdict == StatusFirewallGate823, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatSearch(a.Search), FormatBranches(a.Branches), FormatProtocol(a.Protocol), FormatNative(a.Native), a.Status.Outcome, a.Status.Level, FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
