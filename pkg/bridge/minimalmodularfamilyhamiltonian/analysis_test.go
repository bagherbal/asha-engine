package minimalmodularfamilyhamiltonian

import (
	"strings"
	"testing"
)

func TestGate412InheritanceAndHamiltonian(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate411LeastCostKGen || !a.Inheritance.Gate411NoAxiomPromoted || !a.Inheritance.Gate410NoNativeFamilyBundle || !a.Inheritance.Gate409TrivialU3Multiplicity || !a.Inheritance.Gate408ScalarFlavorBlind {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if a.Hamiltonian.NativeInCurrentAsha || !a.Hamiltonian.ExplicitAxiom || !a.Hamiltonian.Hermitian || !a.Hamiltonian.Traceless || a.Hamiltonian.DistinctEigenvalues != 3 || a.Hamiltonian.MinimalPolynomialDegree != 3 || !a.Hamiltonian.DiagonalOnly {
		t.Fatalf("bad Hamiltonian: %s", FormatHamiltonian(a.Hamiltonian))
	}
}

func TestGate412KMSAndCompatibility(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.KMS.Positive || !a.KMS.Normalized || a.KMS.Tracial || !a.KMS.ModularFlowActive || a.KMS.MaxWeightRatio <= 1 {
		t.Fatalf("bad KMS: %s", FormatKMS(a.KMS))
	}
	if !a.Compatibility.CommutesWithAF || !a.Compatibility.CommutesWithGaugeCharges || !a.Compatibility.CommutesWithHypercharge || !a.Compatibility.CommutesWithSU2L || !a.Compatibility.CommutesWithBL || !a.Compatibility.RequiresFamilyFiberAxiom {
		t.Fatalf("bad compatibility: %s", FormatCompatibility(a.Compatibility))
	}
}

func TestGate412MixingSectorMapAndFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Mixing.NativeNoncommutingPairs != 0 || a.Mixing.ConditionalNoncommutingPairs != 0 || a.Mixing.CKMNative || a.Mixing.PMNSNative || a.Mixing.CKMConditional || a.Mixing.PMNSConditional || !a.Mixing.DiagonalOnly {
		t.Fatalf("bad mixing: %s", FormatMixing(a.Mixing))
	}
	if !a.SectorMap.MassHierarchyCapacity || !a.SectorMap.SectorSpecificMapsNeeded || a.SectorMap.UpSectorMapNative || a.SectorMap.DownSectorMapNative || a.SectorMap.LeptonSectorMapNative || a.SectorMap.ObservedYukawasInserted {
		t.Fatalf("bad sector map: %s", FormatSectorMap(a.SectorMap))
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaMatricesInserted || !a.Firewall.KGenPromotedAsAxiomOnly || !a.Firewall.NoNativeDerivationClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate412ModuliStatusesAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Moduli.NativeReductionBelow13 || !a.Moduli.ConditionalHierarchy || a.Moduli.ConditionalCKMPMNS || !a.Moduli.FirewallPreserved {
		t.Fatalf("bad moduli: %s", FormatModuli(a.Moduli))
	}
	joined := strings.Join(Statuses(a), "\n")
	for _, needle := range []string{StatusMinimalHamiltonianAxiomFormalized, StatusNontracialKMSFamilyStateActivated, StatusHierarchyCapacityActivated, StatusFailedNotNativeDerivation, StatusFailedSingleHamiltonianDiagonalOnly, StatusFailedNoNativeCKMCapacity, StatusFailedNoNativePMNSCapacity, StatusFirewallPreserved13Moduli} {
		if !strings.Contains(joined, needle) {
			t.Fatalf("missing %q in\n%s", needle, joined)
		}
	}
	res := MinimalModularFamilyHamiltonianAxiomConsistencySieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate412Markdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 412 Registry Audit", "Minimal Hamiltonian axiom", "Nontracial KMS state", StatusFailedSingleHamiltonianDiagonalOnly, StatusFirewallPreserved13Moduli, "gate=413"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
