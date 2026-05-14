package familyboundarysourceaxiom

import (
	"strings"
	"testing"
)

func TestGate415CandidateLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Candidates) < 5 {
		t.Fatalf("missing candidates")
	}
	var sectorOK, observedRejected, universalBlind bool
	for _, c := range a.Candidates {
		if c.Name == "charge-sector source boundary" && c.CKMCapacity && c.PMNSCapacity && c.EmpiricalIndependent && !c.NativeToCurrentAsha && !c.FixesCoefficientValues {
			sectorOK = true
		}
		if c.Name == "observed Yukawa matrix source" && c.ImportsObservedYukawa && c.FixesCoefficientValues && !c.EmpiricalIndependent && !c.PromotedToTheorem {
			observedRejected = true
		}
		if c.Name == "universal family source" && c.DiagonalOnly && !c.CKMCapacity {
			universalBlind = true
		}
	}
	if !sectorOK || !observedRejected || !universalBlind {
		t.Fatalf("bad candidates:\n%s", RenderCandidateSummary(a.Candidates))
	}
}

func TestGate415RankingCapacityFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Ranking.LeastCostName != "charge-sector source boundary" || a.Ranking.LeastCost != 2 || !a.Ranking.LeastCostStillAxiom || a.Ranking.LeastCostFixesAngles || !a.Ranking.NoCandidateNative {
		t.Fatalf("bad ranking: %s", FormatRanking(a.Ranking))
	}
	if !a.Capacity.ConditionalCKMAvailable || !a.Capacity.ConditionalPMNSAvailable || a.Capacity.AnyCandidateFixesAngles || a.Capacity.AnyCandidateNative || !a.Capacity.AnyCandidateCurveFitting {
		t.Fatalf("bad capacity: %s", FormatCapacity(a.Capacity))
	}
	if !a.Firewall.AxiomsQuarantined || !a.Firewall.NoNativeDerivationClaimed || !a.Firewall.NoYukawaMatricesInserted {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate415ModuliStatusesAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Moduli.NativeReductionBelow13 || !a.Moduli.ConditionalMixingCapacity || !a.Moduli.CoefficientsRemainFree || !a.Moduli.FirewallPreserved {
		t.Fatalf("bad moduli: %s", FormatModuli(a.Moduli))
	}
	joined := strings.Join(Statuses(a), "\n")
	for _, needle := range []string{StatusSourceAxiomLedgerCompiled, StatusMinimalityRankingAudited, StatusCKMPMNSCapacityAudited, StatusFailedNoNativeBoundary, StatusFailedSelectorRequiresAxiom, StatusFailedDiscreteUnderdetermines, StatusFailedObservedYukawaRejected, StatusFirewallPreserved13Moduli} {
		if !strings.Contains(joined, needle) {
			t.Fatalf("missing %q in\n%s", needle, joined)
		}
	}
	res := FamilyBoundaryConditionSectorSourceAxiomMinimalitySieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate415Markdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 415 Registry Audit", "Candidate axiom ledger", "charge-sector source boundary", StatusFailedNoNativeBoundary, StatusFailedObservedYukawaRejected, StatusFirewallPreserved13Moduli, "gate=416"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
