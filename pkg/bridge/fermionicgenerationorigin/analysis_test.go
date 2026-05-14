package fermionicgenerationorigin

import (
	"strings"
	"testing"
)

func TestGate409Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate408ScalarFlavorBlind || !a.Inheritance.Gate407FullHphiCapacityButNoSelector || !a.Inheritance.Gate395SpinorSplitTwoSector || !a.Inheritance.Gate393TrialityDomainNotAdmitted || !a.Inheritance.Gate394StaticGenerationAddressCentral || a.Inheritance.Gate372ChargedModuliDim != Gate372ChargedFlavorModuliDim || !a.Inheritance.NoEmpiricalInputsImported {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
}

func TestGate409CarrierInventory(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inventory.Executed || a.Inventory.NativeCarrierCount < 7 || a.Inventory.NativeGenerationCarrierCount != 0 || a.Inventory.NativeNoncentralGenerationActions != 0 || a.Inventory.NativeNoncommutingOperatorPairs != 0 || a.Inventory.ColorThreefoldCount < 2 {
		t.Fatalf("bad inventory: %s", FormatInventory(a.Inventory))
	}
}

func TestGate409PrimitiveIdempotents(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Idempotents.Executed || a.Idempotents.NativeThreeBlockCandidates < 1 || a.Idempotents.NativeGenerationLabelCandidates != 0 || a.Idempotents.ColorOrSpeciesRejected < 4 || a.Idempotents.NoncommutingNativePairs != 0 {
		t.Fatalf("bad idempotents: %s", FormatIdempotentAudit(a.Idempotents))
	}
}

func TestGate409CommutantTriality(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Commutant.ContainsU3GenerationFreedom || a.Commutant.GenerationFreedomCanonicalSelector || a.Commutant.NativeDiagonalGenerationOperator || a.Commutant.NativeNoncommutingGenerationPair || !a.Commutant.CentralBroadcastOverGeneration {
		t.Fatalf("bad commutant: %s", FormatCommutant(a.Commutant))
	}
	if a.Triality.TrialityDomainAdmitted || !a.Triality.FermionTo8SRepresentativeFound || !a.Triality.FermionTo8CRepresentativeFound || a.Triality.FermionTo8VRepresentativeFound || a.Triality.GenerationLabelsDerivedNotInserted || !a.Triality.OnePlusTwoDegeneracy {
		t.Fatalf("bad triality: %s", FormatTriality(a.Triality))
	}
}

func TestGate409BilinearsSourcesCKM(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Bilinears.NativeGenerationSensitiveFamilies != 0 || a.Bilinears.NativeNoncommutingFamilies != 0 || a.Bilinears.NativeModuliReducingFamilies != 0 {
		t.Fatalf("bad bilinears: %s", FormatBilinearAudit(a.Bilinears))
	}
	if a.Sources.NativeGenerationHamiltonians != 0 || a.Sources.SealedOrCircularSources < 2 || a.Sources.WrongDomainSources < 1 {
		t.Fatalf("bad sources: %s", FormatSourceAudit(a.Sources))
	}
	if a.CKM.NativeNoncommutingPairs != 0 || a.CKM.CKMCapacityNative || a.CKM.PMNSCapacityNative || !a.CKM.CKMCapacityConditional || a.CKM.SealedNoncommutingPairs == 0 {
		t.Fatalf("bad ckm: %s", FormatCKM(a.CKM))
	}
}

func TestGate409ModuliFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Moduli.StartDim != Gate372ChargedFlavorModuliDim || a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Moduli.NativeReductionBelow13 || !a.Moduli.FirewallPreserved {
		t.Fatalf("bad moduli: %s", FormatModuli(a.Moduli))
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaAmplitudesInserted || !a.Firewall.NoScalarSelectorPromoted || !a.Firewall.NoTauEtaInserted || !a.Firewall.NoNDiagInserted || !a.Firewall.NoManualGenerationLabels || !a.Firewall.NoColorAsGenerationPromoted || !a.Firewall.NoSpeciesAsGenerationPromoted || !a.Firewall.NoModuliReductionClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate409StatusesAndNext(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(Statuses(a), "\n")
	for _, needle := range []string{StatusFermionicCarrierInventoryAudited, StatusFailedTrivialGenerationCopy, StatusFailedSpinorSplitChirality, StatusFailedTrialityExactDegeneracy, StatusFailedColorAsGeneration, StatusFailedSpeciesAsGeneration, StatusFailedNoNativeCKMCapacity, StatusFirewallPreserved13Moduli} {
		if !strings.Contains(joined, needle) {
			t.Fatalf("missing %q in\n%s", needle, joined)
		}
	}
	if a.Next.Gate != 410 || !strings.Contains(a.Next.Title, "Family Bundle") {
		t.Fatalf("bad next: %s", FormatNext(a.Next))
	}
}

func TestGate409Theorem(t *testing.T) {
	res := FermionicMatterCarrierOriginNontrivialGenerationRepresentationSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate409RenderMarkdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 409 Registry Audit", "not an empirical Yukawa seal", StatusFailedTrivialGenerationCopy, StatusFailedNoNativeCKMCapacity, StatusFirewallPreserved13Moduli, "gate=410"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
