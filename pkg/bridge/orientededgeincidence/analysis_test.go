package orientededgeincidence

import (
	"strings"
	"testing"
)

func TestGate403InheritanceAndArena(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate402UndirectedGraphNative || !a.Inheritance.Gate402FullGraphQuarticCapacity || !a.Inheritance.Gate402NoGraphQ4 || !a.Inheritance.Gate385OneFormEdges || a.Inheritance.Gate385JDoubledEdgeCount != 10 || a.Inheritance.Gate372ChargedModuliDim != 13 {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Arena.Formalized || !a.Arena.ChiralOrientationAvailable || a.Arena.MajoranaOrientationCanonical || a.Arena.HasCanonicalHphiQuotient || a.Arena.UsesYukawaAmplitudes || a.Arena.UsesObservedMasses {
		t.Fatalf("bad arena: %s", FormatArena(a.Arena))
	}
}

func TestGate403Q4Target(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Q4.Degree != 4 || !a.Q4.IrreducibleOverQ || len(a.Q4.MonicCoefficients) != 5 {
		t.Fatalf("bad q4: %s", FormatQ4(a.Q4))
	}
}

func TestGate403YukawaBoundaryPairDegenerate(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	yy := findCandidate(a.Sieve.Candidates, "four Yukawa oriented edge Gram d_Y^T d_Y")
	if !yy.Native || !yy.HphiEndomorphism || !yy.CanonicalQuotientToHphi || !yy.OrientationSignsCancel || !yy.PairDegenerate || yy.MinimalDegree != 2 || yy.Q4ExactMatch || yy.PromotableAsQ4Selector {
		t.Fatalf("bad Yukawa boundary: %s", FormatCandidate(yy))
	}
}

func TestGate403FullBoundaryIsFiveDimensionalNotQ4(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	full := findCandidate(a.Sieve.Candidates, "full five-edge oriented incidence Gram d_E^T d_E")
	if !full.Native || full.HphiEndomorphism || full.CanonicalQuotientToHphi || full.Dimension != 5 || full.MinimalDegree != 5 || full.Q4ExactMatch || full.PromotableAsQ4Selector || !full.OrientationSignsCancel {
		t.Fatalf("bad full boundary: %s", FormatCandidate(full))
	}
}

func TestGate403ForcedQuotientAndMajoranaTwistAreQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	quotient := findCandidate(a.Sieve.Candidates, "noncanonical four-mode quotient of full oriented incidence Gram")
	if !quotient.Sealed || !quotient.Circular || quotient.Native || quotient.Dimension != 4 || quotient.MinimalDegree != 4 || quotient.Q4ExactMatch || quotient.PromotableAsQ4Selector {
		t.Fatalf("bad forced quotient: %s", FormatCandidate(quotient))
	}
	twist := findCandidate(a.Sieve.Candidates, "J-twisted complex Majorana boundary d^†d")
	if !twist.Sealed || !twist.CompatibleWithJ || !twist.OrientationSignsCancel || twist.Q4ExactMatch || twist.PromotableAsQ4Selector {
		t.Fatalf("bad Majorana twist: %s", FormatCandidate(twist))
	}
}

func TestGate403ImpactFirewallAndNext(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Sieve.CanonicalHphiQ4MatchCount != 0 || a.Impact.HphiQuarticIdentified || !a.Impact.NativeBoundaryOperatorFound || a.Impact.CanonicalBoundaryQuotientFound || a.Impact.YukawaCouplingsReduced || a.Impact.ChargedModuliResult != 13 || !a.Impact.FlavorFirewallPreserved || !a.Impact.OrientedIncidenceLaneOpened {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaAmplitudesInserted || !a.Firewall.NoArbitraryBoundaryQuotient || !a.Firewall.NoFlavorModuliReductionClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
	if a.Next.Gate != 404 || !strings.Contains(a.Next.Title, "Hphi") {
		t.Fatalf("bad next: %s", FormatNext(a.Next))
	}
}

func TestGate403Theorem(t *testing.T) {
	res := OrientedEdgeIncidenceBoundaryOperatorSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate403RenderMarkdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 403 Registry Audit", StatusFailedNoNativeOrientedQ4Selector, StatusFirewallPreserved13Moduli, "Gate 404"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
