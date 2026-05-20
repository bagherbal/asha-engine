package generation2higgsradialselectorsourcecandidateandvacuumdirectionfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate737CandidateSourcesAndSymmetry(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate736.Inherited || !a.Gate736.RhoPlusMaxEntropy || !a.Gate736.RadialWeightCertified || !a.Gate736.RhoPlusDoesNotSelectPRad || !a.Gate736.RhoPlusDoesNotSelectN {
		t.Fatalf("bad Gate736 inheritance: %+v", a.Gate736)
	}
	if a.Problem.Rank != 1 || !a.Problem.NeedsLineInsideK7Plus || a.Problem.CurrentlyNative {
		t.Fatalf("bad radial selector problem: %+v", a.Problem)
	}
	if len(a.Candidates.Candidates) != 8 || a.Candidates.AnyNativeSelectorFound || a.Candidates.BoundaryScalarsContainVector {
		t.Fatalf("bad candidate source audit: %+v", a.Candidates)
	}
	for _, c := range a.Candidates.Candidates {
		if c.SuppliesPRad {
			t.Fatalf("candidate unexpectedly supplies P_rad: %+v", c)
		}
	}
	if !a.Symmetry.RequiresVacuumSelector || a.Symmetry.CurrentDataSelectsLine {
		t.Fatalf("bad symmetry obstruction: %+v", a.Symmetry)
	}
}

func TestGate737SealHistoryLoopAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Seal.SealNames) != 3 || !a.Seal.DistinctFromN || !a.Seal.DistinctFromQ || !a.Seal.DistinctFromRhoPlus {
		t.Fatalf("bad seal classification: %+v", a.Seal)
	}
	if !a.HistoryLoop.RhoPlusSuppliesWeight || !a.HistoryLoop.PRadSuppliesEvent || !a.HistoryLoop.NSuppliesPhaseLoop || a.HistoryLoop.QSuppliesPRad || !a.HistoryLoop.ConditionalWithoutPRad {
		t.Fatalf("bad HistoryLoop dependence: %+v", a.HistoryLoop)
	}
	if a.Firewall.PRadIsHiggsVacuumTheorem || a.Firewall.PRadIsElectroweakBreakingTheorem || a.Firewall.PhaseTransverseGoldstoneTheorem || a.Firewall.PRadIsHiggsMassTheorem || a.Firewall.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("physical firewall failed: %+v", a.Firewall)
	}
	res := Generation2HiggsRadialSelectorSourceCandidateAndVacuumDirectionFirewallAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
