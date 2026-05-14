package noncommutingmodularpair

import (
	"strings"
	"testing"
)

func TestGate413OperatorAndWeyl(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Operator.ExplicitAxiom || a.Operator.ShiftNativeInCurrentAsha || a.Operator.ShiftOrder != 3 || !a.Operator.ShiftOrthogonal || !a.Operator.Noncommuting || a.Operator.KShiftCommutatorNorm <= 0 || a.Operator.KXCommutatorNorm <= 0 {
		t.Fatalf("bad operator: %s", FormatOperator(a.Operator))
	}
	if !a.Weyl.RootsOfUnityFingerprint || a.Weyl.RootsFixPhysicalAngles || a.Weyl.ClockOrder != 3 || a.Weyl.ShiftOrder != 3 {
		t.Fatalf("bad Weyl: %s", FormatWeyl(a.Weyl))
	}
}

func TestGate413CompatibilityTextureAndFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Compatibility.CommutesWithGaugeCharges || !a.Compatibility.CommutesWithHypercharge || !a.Compatibility.CommutesWithSU2L || !a.Compatibility.CommutesWithBL || !a.Compatibility.RequiresFamilyConnectionAxiom {
		t.Fatalf("bad compatibility: %s", FormatCompatibility(a.Compatibility))
	}
	if a.Texture.NativeNoncommutingPairs != 0 || a.Texture.ConditionalNoncommutingPairs == 0 || !a.Texture.CKMConditional || !a.Texture.PMNSConditional || a.Texture.CKMNative || a.Texture.PMNSNative || !a.Texture.CoefficientsRemainFree || a.Texture.CoefficientsFixedTopologically {
		t.Fatalf("bad texture: %s", FormatTexture(a.Texture))
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaMatricesInserted || !a.Firewall.PairPromotedAsAxiomOnly || !a.Firewall.NoNativeDerivationClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate413ModuliStatusesAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Moduli.NativeReductionBelow13 || !a.Moduli.ConditionalCKMPMNSCapacity || !a.Moduli.CoefficientsFree || !a.Moduli.FirewallPreserved {
		t.Fatalf("bad moduli: %s", FormatModuli(a.Moduli))
	}
	joined := strings.Join(Statuses(a), "\n")
	for _, needle := range []string{StatusSecondOperatorAxiomFormalized, StatusWeylClockShiftPairAudited, StatusNoncommutingPairActivated, StatusCKMPMNSCapacityActivated, StatusFailedSecondOperatorNotNative, StatusFailedCoefficientsRemainFree, StatusFailedRootsUnityDoNotFixAngles, StatusFirewallPreserved13Moduli} {
		if !strings.Contains(joined, needle) {
			t.Fatalf("missing %q in\n%s", needle, joined)
		}
	}
	res := SecondFamilyOperatorNoncommutingModularPairAxiomSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate413Markdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 413 Registry Audit", "Complementary operator axiom", "Weyl clock/shift fingerprint", StatusNoncommutingPairActivated, StatusFailedCoefficientsRemainFree, StatusFirewallPreserved13Moduli, "gate=414"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
