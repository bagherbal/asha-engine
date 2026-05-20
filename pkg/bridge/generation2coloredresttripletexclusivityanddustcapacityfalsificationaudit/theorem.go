package generation2coloredresttripletexclusivityanddustcapacityfalsificationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-821-COLORED-REST-TRIPLET-EXCLUSIVITY-DUST-CAPACITY-FALSIFICATION"
	theoremName = "Gate 821 — Colored RestTriplet Exclusivity and Dust-Capacity Falsification Audit"
)

func Generation2ColoredRestTripletExclusivityAndDustCapacityFalsificationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 820 triplet ledger", Passed: math.Abs(a.Ledger.AlphaB-0.0003878958469680527) < 1e-18 && math.Abs(a.Ledger.BOverT-0.0003877453837799576) < 1e-18 && math.Abs(a.Ledger.SqrtBOverT-0.019691251452864992) < 1e-16, Detail: FormatLedger(a.Ledger)},
			{Name: "derive total rest, triplet use, and dust capacity", Passed: math.Abs(a.Ledger.TotalRestOverT-0.001163687540904158) < 1e-18 && math.Abs(a.Ledger.TripletTraceOverT-0.001163236151339873) < 1e-18 && math.Abs(a.Ledger.DustOverT-4.513895642851889e-7) < 1e-19, Detail: FormatLedger(a.Ledger)},
			{Name: "compute second colored triplet and uncolored dust bounds", Passed: math.Abs(a.Ledger.SecondColoredPerColorBound-1.5046318809506294e-7) < 1e-20 && math.Abs(a.Ledger.SecondColoredSqrtBound-a.Ledger.AlphaB) < 1e-18 && math.Abs(a.Ledger.UncoloredSqrtBound-0.0006718553149936293) < 1e-16, Detail: FormatLedger(a.Ledger)},
			{Name: "record dust-capacity consequence", Passed: containsAll(a.Capacity.Supports, []string{SupportOneLargeTripletOnly, SupportOtherColoredDustBound, SupportDustStronger}) && containsAll(a.Capacity.Failures, []string{FailureMatchAloneInsufficient, FailureSecondColoredAboveDust, FailureUncoloredAboveDust}), Detail: FormatCapacity(a.Capacity)},
			{Name: "define bottom, charm, abstract, and failure branches", Passed: len(a.Branches) == 4 && strings.Contains(a.Branches[0].Name, "bottom") && strings.Contains(a.Branches[1].Name, "charm") && strings.Contains(a.Branches[3].Name, "failure"), Detail: FormatBranches(a.Branches)},
			{Name: "define external ledger falsification protocol", Passed: a.Protocol.CanFalsify && len(a.Protocol.Tests) == 6 && strings.Contains(strings.Join(a.Protocol.Tests, " "), "R_k <= alpha_B^2") && strings.Contains(strings.Join(a.Protocol.Tests, " "), "L_i/T <= 3 alpha_B^2"), Detail: FormatProtocol(a.Protocol)},
			{Name: "audit native sources without sector promotion", Passed: len(a.Native) == 7 && containsAll(a.Native[1].Failures, []string{FailureProjectiveNotTheorem}) && containsAll(a.Native[3].Failures, []string{FailureAlphaNotSector}) && containsAll(a.Native[4].Failures, []string{FailureBFNNotOperator}), Detail: FormatNative(a.Native)},
			{Name: "classify strengthened partial R2 status", Passed: strings.Contains(a.Status.Level, "strengthened partial R2") && !a.Status.NativeSourceFound && !a.Status.ExternalLedgerSupplied && !a.Status.CanUpdateCYukawa, Detail: a.Status.Outcome + " — " + a.Status.Level},
			{Name: "preserve C_Yukawa and C_Higgs firewall", Passed: math.Abs(a.Impact.CandidateCYukawa-0.9992248096922658) < 1e-15 && math.Abs(a.Impact.CandidateCHiggs-1.0372205108665146) < 2e-15 && containsAll(a.Impact.Failures, []string{FailureNoUpdateCYukawa, FailureCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.MatchAloneInsufficient && a.Firewalls.BottomNeedsDust && a.Firewalls.CharmNeedsDust && a.Firewalls.SecondColoredBound && a.Firewalls.UncoloredBound && a.Firewalls.ProjectiveNotTheorem && a.Firewalls.AlphaNotSector && a.Firewalls.BFNNotOperator && a.Firewalls.ExternalNotNative && a.Firewalls.NoCYukawaUpdate && a.Firewalls.CHiggsLevelB && a.Firewalls.TreeProxyNotPole && a.Firewalls.Verdict == StatusFirewallGate821, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatCapacity(a.Capacity), FormatBranches(a.Branches), FormatProtocol(a.Protocol), FormatNative(a.Native), a.Status.Outcome, a.Status.Level, FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
