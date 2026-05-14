package familybundleaxiomledger

import (
	"strings"
	"testing"
)

func TestGate411InheritanceAndLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate410NoNativeFamilyBundle || !a.Inheritance.Gate410RequiresNewAxiom || !a.Inheritance.Gate409TrivialU3Multiplicity || !a.Inheritance.Gate408ScalarFlavorBlind || a.Inheritance.Gate372ChargedModuliDim != Gate372ChargedFlavorModuliDim {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Ledger.Executed || a.Ledger.CandidatesAudited < 5 || a.Ledger.PromotedAxioms != 0 || a.Ledger.LowestCost != 2 || len(a.Ledger.LeastCostNames) == 0 {
		t.Fatalf("bad ledger: %s", FormatLedger(a.Ledger))
	}
}

func TestGate411CapacityAndEmpiricalFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Capacity.NativeNoncommutingPairs != 0 || a.Capacity.CKMNative || a.Capacity.PMNSNative || a.Capacity.ConditionalNoncommutingPairs == 0 || !a.Capacity.CKMConditional || !a.Capacity.PMNSConditional {
		t.Fatalf("bad capacity: %s", FormatCapacity(a.Capacity))
	}
	if !a.EmpiricalIndependence.NoObservedMassesImported || !a.EmpiricalIndependence.NoCKMImported || !a.EmpiricalIndependence.NoPMNSImported || !a.EmpiricalIndependence.NoYukawaMatricesInserted || a.EmpiricalIndependence.CandidatesCanBePureRules < 4 || a.EmpiricalIndependence.CandidatesCollapseToFitting < 1 {
		t.Fatalf("bad empirical audit: %s", FormatEmpiricalIndependence(a.EmpiricalIndependence))
	}
}

func TestGate411RankingBoundaryModuli(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ranking.Executed || len(a.Ranking.Rows) < 5 || a.Ranking.Rows[0].Name != "minimal modular family Hamiltonian axiom" || a.Ranking.Rows[0].Cost != 2 {
		t.Fatalf("bad ranking: %s", FormatRanking(a.Ranking))
	}
	if !a.Boundary.LawSpaceNative || a.Boundary.FamilyBundleNative || !a.Boundary.NewAxiomRequiredForFamilies || a.Boundary.CurrentASHAFlavorComplete {
		t.Fatalf("bad boundary: %s", FormatBoundary(a.Boundary))
	}
	if a.Moduli.StartDim != Gate372ChargedFlavorModuliDim || a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Moduli.NativeReductionBelow13 || !a.Moduli.FirewallPreserved {
		t.Fatalf("bad moduli: %s", FormatModuli(a.Moduli))
	}
}

func TestGate411StatusesAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(Statuses(a), "\n")
	for _, needle := range []string{StatusAxiomLedgerCompiled, StatusEpistemologicalBoundaryDocumented, StatusAxiomCostRankingAudited, StatusFailedNoAxiomPromoted, StatusFailedModularKMSNeedsHamiltonianAxiom, StatusFailedFamilyConnectionNeedsAxiom, StatusFailedPrimitiveIdealNeedsAlgebraAxiom, StatusFailedTrialityLocalSystemNeedsFunctor, StatusFirewallPreserved13Moduli} {
		if !strings.Contains(joined, needle) {
			t.Fatalf("missing %q in\n%s", needle, joined)
		}
	}
	res := AxiomCandidateLedgerNontrivialFamilyBundleExtensionsTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate411Markdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 411 Registry Audit", "Axiom candidate ledger", "Cost ranking", StatusAxiomLedgerCompiled, StatusFirewallPreserved13Moduli, "gate=412"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
