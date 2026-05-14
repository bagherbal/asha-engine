package fermionicfamilybundleextension

import (
	"strings"
	"testing"
)

func TestGate410Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate409FermionicCarrierTrivial || !a.Inheritance.Gate409U3GenerationFreedomUnselected || !a.Inheritance.Gate409NoNativeCKMCapacity || !a.Inheritance.Gate408ScalarFlavorBlind || a.Inheritance.Gate372ChargedModuliDim != Gate372ChargedFlavorModuliDim || !a.Inheritance.NoEmpiricalInputsImported {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
}

func TestGate410ExtensionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Extensions.Executed || a.Extensions.CandidatesAudited < 6 || a.Extensions.NativeNontrivialBundles != 0 || a.Extensions.NativeConnections != 0 || a.Extensions.NativeNoncommutingPairs != 0 || a.Extensions.ConditionalNontrivialBundles < 1 || a.Extensions.RequiresNewAxiomCandidates < 2 {
		t.Fatalf("bad extensions: %s", FormatExtensionAudit(a.Extensions))
	}
}

func TestGate410FamilyBundleKOModularPrimitive(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FamilyBundle.TrivialMultiplicity || !a.FamilyBundle.ContainsU3Freedom || a.FamilyBundle.U3FreedomSelectedByGeometry || a.FamilyBundle.NativeFamilyConnection || a.FamilyBundle.NativeFamilyCurvature || a.FamilyBundle.ReplacesTensorC3 {
		t.Fatalf("bad family: %s", FormatFamilyBundle(a.FamilyBundle))
	}
	if !a.KOTwist.KODimensionSignsAudited || !a.KOTwist.ChangesJGammaCommutation || a.KOTwist.ChangesMultiplicity || a.KOTwist.ProducesThreeFamilies || a.KOTwist.ProducesNoncommutingTextures {
		t.Fatalf("bad KO: %s", FormatKOTwist(a.KOTwist))
	}
	if !a.ModularKMS.NontracialStateHasCapacity || a.ModularKMS.NativeHamiltonianFound || a.ModularKMS.NativeDensityMatrixFound || !a.ModularKMS.RequiresExternalHamiltonian {
		t.Fatalf("bad KMS: %s", FormatModularKMS(a.ModularKMS))
	}
	if !a.PrimitiveIdeals.ExistingIdealsAreWrongDomain || a.PrimitiveIdeals.ThreeFamilyIdealDerived || a.PrimitiveIdeals.ActsOnC3GenNoncentrally || !a.PrimitiveIdeals.RequiresAlgebraEnlargement {
		t.Fatalf("bad ideals: %s", FormatPrimitiveIdeals(a.PrimitiveIdeals))
	}
}

func TestGate410NoncommutingAndModuli(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Noncommuting.NativeFamilyOperators != 0 || a.Noncommuting.NativeNoncommutingPairs != 0 || a.Noncommuting.CKMCapacityNative || a.Noncommuting.ConditionalNoncommutingPairs == 0 || !a.Noncommuting.CKMCapacityConditional {
		t.Fatalf("bad noncommuting: %s", FormatNoncommuting(a.Noncommuting))
	}
	if a.Moduli.StartDim != Gate372ChargedFlavorModuliDim || a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Moduli.NativeReductionBelow13 || !a.Moduli.FirewallPreserved {
		t.Fatalf("bad moduli: %s", FormatModuli(a.Moduli))
	}
}

func TestGate410FirewallsStatusesNext(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaAmplitudesInserted || !a.Firewall.NoExternalHamiltonianPromoted || !a.Firewall.NoManualFamilyBundlePromoted || !a.Firewall.NoNewAxiomPromoted || !a.Firewall.NoModuliReductionClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
	joined := strings.Join(Statuses(a), "\n")
	for _, needle := range []string{StatusExtensionSearchFormalized, StatusFailedNoNativeFamilyBundle, StatusFailedKOTwistOnlyChangesSigns, StatusFailedKMSRequiresExternalHamiltonian, StatusFailedExtensionRequiresNewAxiom, StatusFailedTrivialGenerationMultiplicity, StatusFirewallPreserved13Moduli} {
		if !strings.Contains(joined, needle) {
			t.Fatalf("missing %q in\n%s", needle, joined)
		}
	}
	if a.Next.Gate != 411 || !strings.Contains(a.Next.Title, "Axiom-Candidate") {
		t.Fatalf("bad next: %s", FormatNext(a.Next))
	}
}

func TestGate410TheoremAndMarkdown(t *testing.T) {
	res := FermionicRepresentationExtensionNontrivialFamilyBundleSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 410 Registry Audit", "Extension candidate table", StatusFailedNoNativeFamilyBundle, StatusFailedKMSRequiresExternalHamiltonian, StatusFirewallPreserved13Moduli, "gate=411"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
