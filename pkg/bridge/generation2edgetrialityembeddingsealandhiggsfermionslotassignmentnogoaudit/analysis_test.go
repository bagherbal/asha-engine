package generation2edgetrialityembeddingsealandhiggsfermionslotassignmentnogoaudit

import (
	"strings"
	"testing"
)

func TestGate805InheritanceSealAndCanonicalCandidate(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.ArityOnly || a.Inheritance.CarrierEmbeddingFound || !containsAll(a.Inheritance.Verdicts, []string{StatusGate804Inherited, StatusAritySelected}) {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Seal.Defined || a.Seal.Name != "EdgeTrialityEmbeddingSeal" || !containsAll(a.Seal.Components, []string{"Higgs-slot embedding", "left-fermion slot embedding", "right-fermion slot embedding", "real-form descent", "gauge-label preservation", "chirality compatibility"}) {
		t.Fatalf("bad seal: %s", FormatSeal(a.Seal))
	}
	if a.Canonical.Assignment["Higgs"] != "V_C" || a.Canonical.Assignment["psi_L"] != "S_plus_C" || a.Canonical.Assignment["psi_R"] != "S_minus_C" {
		t.Fatalf("bad canonical assignment: %s", FormatSlotCandidate(a.Canonical))
	}
	if !containsAll(a.Canonical.Failures, []string{StatusNoHiggsToVector, StatusNoFermionsToSpinors}) {
		t.Fatalf("missing canonical failures: %s", FormatSlotCandidate(a.Canonical))
	}
}

func TestGate805SlotPermutationHiggsAndFermionFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Permutations.Audited || !containsAll(a.Permutations.Failures, []string{StatusPermutationRoleFail, StatusParityFail}) {
		t.Fatalf("bad permutation audit: %s", FormatPermutation(a.Permutations))
	}
	if !containsAll(a.Higgs.Failures, []string{StatusK7PlusC2NotD4C8, StatusNoHiggsC2ToD4C8, StatusHiggsOneFormNotD4}) {
		t.Fatalf("bad Higgs audit: %s", FormatEmbedding(a.Higgs))
	}
	if !containsAll(a.Fermion.Failures, []string{StatusNoSMFermionToD4Spinor, StatusChiralityNotD4, StatusNoSectorEmbeddings}) {
		t.Fatalf("bad fermion audit: %s", FormatEmbedding(a.Fermion))
	}
	if !containsAll(a.Chirality.Failures, []string{StatusSMChiralityNotD4, StatusNoChiralitySeal}) {
		t.Fatalf("bad chirality firewall: %s", FormatFirewall(a.Chirality))
	}
}

func TestGate805GaugeHermitianRealFormAndCandidateTable(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.Gauge.Failures, []string{StatusD4NoGaugeLabels, StatusNoGaugePreservingMap, StatusNoHyperchargeFromSlot}) {
		t.Fatalf("bad gauge firewall: %s", FormatFirewall(a.Gauge))
	}
	if !containsAll(a.Hermitian.Failures, []string{StatusEmbeddingNoGenOperator, StatusKernelNoYF, StatusNoYdaggerY}) {
		t.Fatalf("bad Hermitian obstruction: %s", FormatFirewall(a.Hermitian))
	}
	if !containsAll(a.RealForm.Failures, []string{StatusComplexNotNative, StatusNoRealDescent}) {
		t.Fatalf("bad real-form obstruction: %s", FormatFirewall(a.RealForm))
	}
	if !a.Table.Recorded || len(a.Table.Rows) != 4 || !containsAll([]string{FormatTable(a.Table)}, []string{"Candidate A", "Candidate D", StatusNoCandidateCertified}) {
		t.Fatalf("bad candidate table: %s", FormatTable(a.Table))
	}
}

func TestGate805PackageCHiggsBranchAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Package.Updated || !containsAll(a.Package.NotSupplied, []string{"EdgeTrialityEmbeddingSeal", "GenerationCarrierSeal", "TraceAtomExtractionSeal"}) {
		t.Fatalf("bad package status: %s", FormatPackage(a.Package))
	}
	if !a.CHiggs.Preserved || !containsAll(a.CHiggs.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}) {
		t.Fatalf("bad C_Higgs firewall: %s", FormatCHiggs(a.CHiggs))
	}
	if !strings.Contains(a.Branch.Next, "Gate 806") || a.Branch.Seal != "GenerationOperatorSeal" {
		t.Fatalf("bad branch: %+v", a.Branch)
	}
	res := Generation2EdgeTrialityEmbeddingSealAndHiggsFermionSlotAssignmentNoGoAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
