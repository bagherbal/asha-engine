package generation2boundaryalphaoneplusthreerestsimplexandconcentrationsourceaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-818-BOUNDARY-ALPHA-ONE-PLUS-THREE-REST-SIMPLEX-CONCENTRATION-SOURCE"
	theoremName = "Gate 818 — Boundary-Alpha 1+3 Rest Simplex and Concentration Source Audit"
)

func Generation2BoundaryAlphaOnePlusThreeRestSimplexAndConcentrationSourceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "define boundary alpha 1+3 rest simplex", Passed: a.Formula.Formula != "" && containsAll(a.Formula.Supports, []string{StatusProjectiveRelevant}) && containsAll(a.Formula.Failures, []string{StatusSimplexNotYukawa}), Detail: a.Formula.Formula},
			{Name: "audit t equals alpha_B prior simplex and fifth-order residual", Passed: math.Abs(a.Prior.T-a.Ledger.AlphaB) < 1e-18 && math.Abs(a.Prior.NEffResidual) < 3e-16 && math.Abs(a.Prior.SymbolicResidual+2.107593378826735e-16) < 1e-27 && a.Prior.QSimplex > a.Prior.QRestB && containsAll(a.Prior.Supports, []string{StatusSimplexSourcesQ, StatusFifthOrderClosure}), Detail: FormatPrior(a.Prior)},
			{Name: "audit exact t_star branch and target-solved firewall", Passed: math.Abs(a.TStar.TStar-a.Ledger.AlphaB-2.336e-10) < 5e-14 && math.Abs(a.TStar.QResidual) < 1e-15 && containsAll(a.TStar.Supports, []string{StatusTStarExact}) && containsAll(a.TStar.Failures, []string{StatusTStarMayBeTargetSolved, StatusSquareRootNotNative}), Detail: FormatTStar(a.TStar)},
			{Name: "reaudit three-rest and one-rest controls", Passed: len(a.Controls) == 2 && !controlByName(a.Controls, "three equal rest atoms").Exact && controlByName(a.Controls, "three equal rest atoms").Q > a.Ledger.QRestB && !controlByName(a.Controls, "one concentrated rest atom").Exact && controlByName(a.Controls, "one concentrated rest atom").Q == 1, Detail: FormatControls(a.Controls)},
			{Name: "construct abstract positive rest atoms from simplex", Passed: a.Spectrum.Realizable && math.Abs(a.Spectrum.Sum-1) < 1e-15 && math.Abs(a.Spectrum.Q-a.Prior.QSimplex) < 1e-18 && containsAll(a.Spectrum.Failures, []string{StatusNoSectorAssignment, StatusNotNativeYukawa}), Detail: FormatSpectrum(a.Spectrum)},
			{Name: "audit structural source lanes without promotion", Passed: !a.Structural.IndependentNativeSource && containsAll(a.Structural.Supports, []string{StatusProjectiveRelevant, StatusK7Resonance}) && containsAll(a.Structural.Failures, []string{StatusProjectiveNotTheorem, StatusK7NotTheorem}), Detail: strings.Join(a.Structural.CandidateLanes, " | ")},
			{Name: "classify strengthened partial R2 status", Passed: strings.Contains(a.Status.Level, "strengthened partial R2") && a.Status.ConstructsConcentration && a.Status.ConstructsPositiveSpectrum && !a.Status.ConstructsSectorLedger && !a.Status.NativeYukawaTheorem, Detail: a.Status.Level},
			{Name: "record C_Yukawa and C_Higgs candidate impact without official update", Passed: math.Abs(a.Impact.CYukawaCandidate-0.999224809692266) < 1e-15 && math.Abs(a.Impact.CHiggsCandidate-1.0372205108665148) < 1e-15 && containsAll(a.Impact.Failures, []string{StatusNoCertifiedMap, StatusCHiggsLevelB, StatusTreeProxyNotPole}), Detail: FormatImpact(a.Impact)},
			{Name: "record next branch", Passed: strings.Contains(a.Branch.NextGate, "Gate 819") && strings.Contains(a.Branch.Outcome, "partial R2"), Detail: a.Branch.Outcome + " -> " + a.Branch.NextGate},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.SimplexNotTheorem && a.Firewalls.TStarNotSource && a.Firewalls.AbstractNotSector && a.Firewalls.SectorLedgerNotNative && a.Firewalls.BoundaryNotYukawa && a.Firewalls.ProjectiveNotYukawa && a.Firewalls.K7NotYukawa && a.Firewalls.CHiggsLevelB && a.Firewalls.TreeProxyNotPole && a.Firewalls.Verdict == StatusFirewallGate818, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), a.Formula.Formula, FormatPrior(a.Prior), FormatTStar(a.TStar), FormatControls(a.Controls), FormatSpectrum(a.Spectrum), strings.Join(a.Structural.CandidateLanes, " | "), a.Status.Level, FormatImpact(a.Impact), a.Branch.Outcome, a.Branch.NextGate, a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func controlByName(rows []Control, name string) Control {
	for _, r := range rows {
		if r.Name == name {
			return r
		}
	}
	return Control{}
}
