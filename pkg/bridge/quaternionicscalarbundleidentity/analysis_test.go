package quaternionicscalarbundleidentity

import (
	"strings"
	"testing"
)

func TestGate399QuaternionicModuleButNoGlobalSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Module.RealDimension != 4 || a.Module.ComplexDoubletDimension != 2 || !a.Module.LocalHExtracted || !a.Module.MoritaWeakHAction || !a.Module.PairComplexAvailable || !a.Module.AbstractQuaternionicTripleAvailable {
		t.Fatalf("bad quaternionic module audit: %s", FormatModule(a.Module))
	}
	if a.Module.GlobalHUnsealed || a.Module.CanonicalComplexDerived || a.Module.QuaternionicTripleSelectedByScalar || a.Module.FullScalarSU2Recovered {
		t.Fatalf("quaternionic scalar module incorrectly unsealed: %s", FormatModule(a.Module))
	}
}

func TestGate399QuaternionicFingerprintsAreQuadraticNotQ4(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"left H unit I on H_phi", "left H unit J pair-rotation on H_phi", "left H unit K on H_phi"} {
		fp := findEndomorphism(a.Endomorphisms.Candidates, name)
		if !fp.Native || !fp.QuaternionicAction || !fp.SquaresToMinusIdentity || fp.MinimalDegree != 2 || !fp.CharPolyIsSquareOfQuadratic || fp.Q4ExactMatch || fp.PromotableAsQ4Selector {
			t.Fatalf("bad H fingerprint for %s: %s", name, FormatEndomorphism(fp))
		}
	}
	j := findEndomorphism(a.Endomorphisms.Candidates, "left H unit J pair-rotation on H_phi")
	if !j.CommutesWithScalarResponse || j.ScalarCommutatorNorm > 1e-9 {
		t.Fatalf("expected pair-compatible J to commute with scalar response: %s", FormatEndomorphism(j))
	}
}

func TestGate399GenericQuaternionicElementCannotSelectQ4(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	generic := findEndomorphism(a.Endomorphisms.Candidates, "generic single quaternion element")
	if !generic.Native || !generic.QuaternionicAction || generic.MinimalDegree != 2 || !generic.CharPolyIsSquareOfQuadratic || generic.Q4ExactMatch || generic.PromotableAsQ4Selector {
		t.Fatalf("generic quaternion element incorrectly promoted: %s", FormatEndomorphism(generic))
	}
}

func TestGate399SealedCompanionIsQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	sealed := findEndomorphism(a.Endomorphisms.Candidates, "sealed q4 companion operator placed on H_phi")
	if !sealed.Sealed || !sealed.Circular || !sealed.Q4ExactMatch || sealed.QuaternionicAction || sealed.CompatibleWithJ || sealed.CompatibleWithFirstOrder || sealed.PromotableAsQ4Selector {
		t.Fatalf("sealed companion incorrectly classified: %s", FormatEndomorphism(sealed))
	}
}

func TestGate399IdentityFirewallAndStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Endomorphisms.PromotableNativeCount != 0 || a.Identity.HphiQuarticIdentified || a.Identity.YukawaCouplingsReduced || a.Identity.ChargedModuliResult != 13 || !a.Identity.FlavorFirewallPreserved {
		t.Fatalf("bad identity/firewall audit: %s :: %s", FormatEndomorphisms(a.Endomorphisms), FormatIdentity(a.Identity))
	}
	statuses := Statuses(a)
	for _, req := range []string{StatusFailedHActionPolynomialDisjoint, StatusFailedHActionMinPolyQuadratic, StatusFailedNoQ4ScalarEndomorphism, StatusFailedNoCanonicalHphiID, StatusFirewallPreserved13Moduli} {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
}

func TestGate399MarkdownAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, want := range []string{"Gate 399 Registry Audit", "Quaternionic", "q4", "FAILED_ROUTE_QUATERNIONIC_ACTION_POLYNOMIAL_DISJOINT_FROM_Q4", "FAILED_ROUTE_H_ACTION_MINIMAL_POLYNOMIAL_QUADRATIC_NOT_QUARTIC", "Gate 400"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
	res := QuaternionicScalarBundleIdentitySieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit failed:\n%s", res.Details())
	}
}
