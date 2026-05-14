package spectralgraphedgeadjacency

import (
	"strings"
	"testing"
)

func TestGate402InheritanceAndArena(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate400NoNativeQ4Selector || !a.Inheritance.Gate401AnisotropicWeightsFound || !a.Inheritance.Gate401NoNativeWeightedLaplacian || !a.Inheritance.Gate385OneFormEdges || a.Inheritance.Gate385JDoubledEdgeCount != 10 || !a.Inheritance.Gate297FirstOrderEdgeGraph || !a.Inheritance.Gate298InnerFluctuationFields || a.Inheritance.Gate372ChargedModuliDim != 13 {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Arena.Formalized || a.Arena.StructuralEdgeCount != 5 || a.Arena.YukawaEdgeCount != 4 || a.Arena.JDoubledEdgeCount != 10 || !a.Arena.HasCanonicalEndpointIncidence || a.Arena.HasCanonicalHphiQuotient || a.Arena.UsesGaugeChargeWeights || a.Arena.UsesYukawaAmplitudes || a.Arena.UsesObservedMasses {
		t.Fatalf("bad arena: %s", FormatArena(a.Arena))
	}
}

func TestGate402Q4Target(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Q4.Degree != 4 || !a.Q4.IrreducibleOverQ || len(a.Q4.MonicCoefficients) != 5 {
		t.Fatalf("bad q4: %s", FormatQ4(a.Q4))
	}
}

func TestGate402FourYukawaGraphIsPairDegenerate(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	adj := findCandidate(a.Sieve.Candidates, "four Yukawa-edge adjacency graph K2 disjoint union K2")
	if !adj.HphiEndomorphism || !adj.CanonicalQuotientToHphi || !adj.PairDegenerate || adj.MinimalDegree != 2 || adj.Q4ExactMatch || adj.PromotableAsQ4Selector {
		t.Fatalf("adjacency incorrectly classified: %s", FormatCandidate(adj))
	}
	lap := findCandidate(a.Sieve.Candidates, "four Yukawa-edge graph Laplacian K2 disjoint union K2")
	if !lap.HphiEndomorphism || !lap.CanonicalQuotientToHphi || !lap.PairDegenerate || lap.MinimalDegree != 2 || lap.Q4ExactMatch || lap.PromotableAsQ4Selector {
		t.Fatalf("laplacian incorrectly classified: %s", FormatCandidate(lap))
	}
}

func TestGate402FullEdgeGraphHasQuarticCapacityButNotQ4(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	full := findCandidate(a.Sieve.Candidates, "full five-edge structural Laplacian P3 disjoint union K2")
	if !full.Native || full.HphiEndomorphism || full.CanonicalQuotientToHphi || full.Dimension != 5 || full.MinimalDegree != 4 || !full.IrreducibleQuarticCapacity || full.MinimalResidualToQ4 <= 1 || full.Q4ExactMatch || full.PromotableAsQ4Selector {
		t.Fatalf("full graph incorrectly classified: %s", FormatCandidate(full))
	}
	pos := findCandidate(a.Sieve.Candidates, "positive-spectrum quotient of full five-edge Laplacian")
	if pos.Dimension != 3 || pos.HphiEndomorphism || pos.Q4ExactMatch || pos.PromotableAsQ4Selector {
		t.Fatalf("positive quotient incorrectly classified: %s", FormatCandidate(pos))
	}
}

func TestGate402JDoubleDoesNotCreateScalarSelector(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	d := findCandidate(a.Sieve.Candidates, "J-doubled structural edge graph")
	if !d.Native || !d.CompatibleWithJ || d.Dimension != 10 || d.HphiEndomorphism || d.Q4ExactMatch || d.PromotableAsQ4Selector {
		t.Fatalf("J doubled graph incorrectly classified: %s", FormatCandidate(d))
	}
}

func TestGate402SealedCompanionQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	sealed := findCandidate(a.Sieve.Candidates, "sealed q4 edge-graph companion quotient")
	if !sealed.Sealed || !sealed.Circular || !sealed.Q4ExactMatch || sealed.Native || sealed.CanonicalQuotientToHphi || sealed.PromotableAsQ4Selector {
		t.Fatalf("sealed companion incorrectly classified: %s", FormatCandidate(sealed))
	}
}

func TestGate402ImpactFirewallAndNext(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Sieve.CanonicalHphiQ4MatchCount != 0 || a.Impact.HphiQuarticIdentified || !a.Impact.NativeEdgeAdjacencyFound || a.Impact.CanonicalGraphQuotientFound || a.Impact.YukawaCouplingsReduced || a.Impact.ChargedModuliResult != 13 || !a.Impact.FlavorFirewallPreserved || !a.Impact.EdgeGraphLaneOpenedButUnsealed {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaAmplitudesInserted || !a.Firewall.NoArbitraryGraphQuotient || !a.Firewall.NoFlavorModuliReductionClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
	if a.Next.Gate != 403 || !strings.Contains(a.Next.Title, "Oriented") {
		t.Fatalf("bad next: %s", FormatNext(a.Next))
	}
}

func TestGate402Theorem(t *testing.T) {
	res := SpectralGraphEdgeAdjacencyOperatorSearchTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate402RenderMarkdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 402 Registry Audit", StatusFailedNoNativeQ4EdgeAdjacency, StatusFirewallPreserved13Moduli, "Gate 403"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
