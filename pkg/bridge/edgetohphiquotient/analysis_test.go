package edgetohphiquotient

import (
	"strings"
	"testing"
)

func TestGate404InheritanceAndArena(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate403NeedsQuotient || !a.Inheritance.Gate403OrientationNoQ4 || !a.Inheritance.Gate385OneFormEdges || a.Inheritance.Gate385JDoubledEdgeCount != 10 || a.Inheritance.Gate372ChargedModuliDim != 13 {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Arena.Formalized || a.Arena.HasCanonicalFullEdgeQuotient || !a.Arena.HasCanonicalYukawaRestriction || !a.Arena.HasCanonicalBranchMap || !a.Arena.HasCanonicalJEvenMap || a.Arena.UsesObservedMasses || a.Arena.UsesYukawaAmplitudes || a.Arena.UsesManualQ4Placement {
		t.Fatalf("bad arena: %s", FormatArena(a.Arena))
	}
}

func TestGate404Q4Target(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Q4.Degree != 4 || !a.Q4.IrreducibleOverQ || len(a.Q4.MonicCoefficients) != 5 {
		t.Fatalf("bad q4: %s", FormatQ4(a.Q4))
	}
}

func TestGate404CanonicalYukawaQuotientIsPairDegenerate(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	qy := findCandidate(a.Sieve.Candidates, "canonical Higgs/Yukawa edge restriction Q_Y: E_5 -> E_Y ~= H_phi")
	if !qy.Native || !qy.CanonicalQuotient || !qy.HphiEndomorphism || qy.Rank != 4 || qy.KernelDimension != 1 || !qy.PairDegenerate || qy.MinimalDegree != 2 || qy.Q4ExactMatch || qy.PromotableAsQ4Selector {
		t.Fatalf("bad Q_Y: %s", FormatCandidate(qy))
	}
}

func TestGate404BranchJAndContactQuotientsAreQuadratic(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	branch := findCandidate(a.Sieve.Candidates, "scalar branch quotient Q_branch: E_Y -> Phi_+ ⊕ Phi_-")
	if !branch.Native || !branch.CanonicalQuotient || branch.Rank != 2 || branch.MinimalDegree != 2 || !branch.PairDegenerate || branch.Q4ExactMatch || branch.PromotableAsQ4Selector {
		t.Fatalf("bad branch quotient: %s", FormatCandidate(branch))
	}
	j := findCandidate(a.Sieve.Candidates, "J-even/J-odd quotient from ten J-doubled edge slots")
	if !j.Native || !j.CanonicalQuotient || !j.JCompatible || j.MinimalDegree != 2 || !j.PairDegenerate || j.Q4ExactMatch || j.PromotableAsQ4Selector {
		t.Fatalf("bad J quotient: %s", FormatCandidate(j))
	}
	contact := findCandidate(a.Sieve.Candidates, "contact/scalar response quotient Q_contact from active contact sector")
	if !contact.Native || !contact.ContactDerived || !contact.HphiEndomorphism || contact.MinimalDegree != 2 || !contact.PairDegenerate || contact.Q4ExactMatch || contact.PromotableAsQ4Selector {
		t.Fatalf("bad contact quotient: %s", FormatCandidate(contact))
	}
}

func TestGate404ManualQuarticRoutesAreQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	full := findCandidate(a.Sieve.Candidates, "full five-edge spectral quotient by chosen edge mode")
	if !full.Sealed || !full.Circular || full.Native || full.CanonicalQuotient || full.MinimalDegree != 4 || !full.IrreducibleQuarticCapacity || full.PromotableAsQ4Selector {
		t.Fatalf("bad full quotient: %s", FormatCandidate(full))
	}
	sealed := findCandidate(a.Sieve.Candidates, "sealed q4 edge-to-Hphi companion quotient")
	if !sealed.Sealed || !sealed.Circular || sealed.Native || sealed.CanonicalQuotient || !sealed.Q4ExactMatch || sealed.PromotableAsQ4Selector {
		t.Fatalf("bad sealed q4 quotient: %s", FormatCandidate(sealed))
	}
}

func TestGate404ImpactFirewallAndNext(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Sieve.CanonicalHphiQ4MatchCount != 0 || a.Sieve.NativeQuarticCapacityCount != 0 || !a.Impact.CanonicalQuotientFound || !a.Impact.CanonicalYukawaQuotientFound || a.Impact.HphiQuarticIdentified || a.Impact.NativeIntertwinerQ4Found || a.Impact.YukawaCouplingsReduced || a.Impact.ChargedModuliResult != 13 || !a.Impact.FlavorFirewallPreserved {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaAmplitudesInserted || !a.Firewall.NoArbitraryFullEdgeQuotientPromoted || !a.Firewall.NoFlavorModuliReductionClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
	if a.Next.Gate != 405 || !strings.Contains(a.Next.Title, "Contact-to-Edge") {
		t.Fatalf("bad next: %s", FormatNext(a.Next))
	}
}

func TestGate404Theorem(t *testing.T) {
	res := CanonicalEdgeToHphiQuotientContactEdgeIntertwinerSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate404RenderMarkdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 404 Registry Audit", StatusFailedNoNativeIntertwinerQ4, StatusFirewallPreserved13Moduli, "Gate 405"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
