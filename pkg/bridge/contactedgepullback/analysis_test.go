package contactedgepullback

import (
	"strings"
	"testing"
)

func TestGate405InheritanceAndArena(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate404QuotientNoQ4 || !a.Inheritance.Gate404NeedsPullback || !a.Inheritance.Gate385OneFormEdges || a.Inheritance.Gate385JDoubledEdgeCount != JDoubledEdgeCount || a.Inheritance.Gate372ChargedModuliDim != Gate372ChargedModuliDim {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Arena.Formalized || a.Arena.NativeFunctorKnown || a.Arena.ContactEdgeActionDerived || a.Arena.UsesObservedMasses || a.Arena.UsesYukawaAmplitudes || a.Arena.UsesManualRootPlacement || a.Arena.StructuralEdgeDim != StructuralEdgeCount || a.Arena.JDoubledEdgeDim != JDoubledEdgeCount {
		t.Fatalf("bad arena: %s", FormatArena(a.Arena))
	}
}

func TestGate405Q4Target(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Q4.Degree != Q4Degree || a.Q4.Dimension != ContactPrimaryDim || !a.Q4.IrreducibleOverQ || len(a.Q4.MonicCoefficients) != 5 || !strings.Contains(a.Q4.NeededMap, "pullback") {
		t.Fatalf("bad q4 target: %s", FormatQ4(a.Q4))
	}
}

func TestGate405NoNativeContactToEdgeMap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	native := findCandidate(a.Sieve.Candidates, "native contact projector to one-form edge ledger")
	if native.Native || native.Typed || native.PullbackConstructed || native.PreservesQ4Polynomial || native.DFIntertwiner || native.NaturalitySquareFormed || native.Verdict != StatusFailedNoNativeContactToEdgeMap {
		t.Fatalf("bad native candidate: %s", FormatCandidate(native))
	}
}

func TestGate405YukawaRestrictionIsWrongDirection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	y := findCandidate(a.Sieve.Candidates, "reverse of canonical Yukawa edge restriction")
	if !y.Native || !y.Typed || !y.EdgeDerived || !y.Circular || y.ContactDerived || y.PullbackConstructed || y.PreservesQ4Polynomial || y.Verdict != StatusFailedYukawaRestrictionWrongDirection {
		t.Fatalf("bad Yukawa reverse candidate: %s", FormatCandidate(y))
	}
}

func TestGate405SealedQ4PlacementsAreQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	ext := findCandidate(a.Sieve.Candidates, "sealed q4 extension to five structural edge slots")
	if !ext.Sealed || !ext.Circular || !ext.PullbackConstructed || !ext.PreservesQ4Polynomial || ext.Native || ext.Canonical || ext.JCompatible || ext.FirstOrderCompatible || ext.DFIntertwiner || ext.PromotableAsQ4EdgeWeight {
		t.Fatalf("bad q4 E5 extension: %s", FormatCandidate(ext))
	}
	j := findCandidate(a.Sieve.Candidates, "sealed J-doubled q4 pullback")
	if !j.Sealed || !j.Circular || !j.PullbackConstructed || !j.PreservesQ4Polynomial || !j.JCompatible || j.Native || j.Canonical || j.FirstOrderCompatible || j.DFIntertwiner || j.PromotableAsQ4EdgeWeight {
		t.Fatalf("bad q4 J duplicate: %s", FormatCandidate(j))
	}
}

func TestGate405NoDFIntertwinerAndImpact(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	it := findCandidate(a.Sieve.Candidates, "contact q4 as edge weight/intertwiner with native D_F edge graph")
	if !it.ContactDerived || !it.EdgeDerived || it.DFIntertwiner || it.NaturalitySquareFormed || it.CommutatorZero || it.PromotableAsQ4EdgeWeight {
		t.Fatalf("bad intertwiner candidate: %s", FormatCandidate(it))
	}
	if a.Sieve.NativePullbackCount != 0 || a.Sieve.NativeQ4PreservingCount != 0 || a.Sieve.NativeDFIntertwinerCount != 0 || a.Sieve.CanonicalNaturalTransformCount != 0 || a.Sieve.SealedOrManualCount == 0 {
		t.Fatalf("bad sieve: %s", FormatSieve(a.Sieve))
	}
	if a.Impact.ContactPullbackAchieved || a.Impact.Q4OnEdgeSpacePreserved || a.Impact.CanonicalNaturalTransformation || a.Impact.HphiQuarticIdentified || a.Impact.YukawaCouplingsReduced || a.Impact.ChargedModuliResult != Gate372ChargedModuliDim || !a.Impact.FlavorFirewallPreserved {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
}

func TestGate405FirewallAndNext(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaAmplitudesInserted || !a.Firewall.NoManualQ4HphiID || !a.Firewall.NoManualRootPlacementPromoted || !a.Firewall.NoArbitraryEdgeBasisPromoted || !a.Firewall.NoCompanionOperatorPromoted || !a.Firewall.NoFlavorModuliReductionClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
	if a.Next.Gate != 406 || !strings.Contains(a.Next.Title, "Contact-Eigenoperator") {
		t.Fatalf("bad next: %s", FormatNext(a.Next))
	}
}

func TestGate405Theorem(t *testing.T) {
	res := ContactToEdgeNaturalTransformationPullbackSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate405RenderMarkdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 405 Registry Audit", StatusFailedNoNativeContactToEdgeMap, StatusFailedNoNaturalTransformation, StatusFirewallPreserved13Moduli, "gate=406"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
