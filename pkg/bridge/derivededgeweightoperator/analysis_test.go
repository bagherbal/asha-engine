package derivededgeweightoperator

import (
	"strings"
	"testing"
)

func TestGate401InheritanceAndArena(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate400UniformCentral || !a.Inheritance.Gate400NoNativeQ4Selector || !a.Inheritance.Gate385OneFormEdges || a.Inheritance.Gate385JDoubledEdgeCount != 10 || a.Inheritance.Gate372ChargedModuliDim != 13 {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Arena.Formalized || a.Arena.StructuralEdgeCount != 5 || a.Arena.JDoubledEdgeCount != 10 || a.Arena.HphiDimension != 4 || !a.Arena.NativeElectroweakWeights || !a.Arena.NativeBMinusLWeights || !a.Arena.NativeT3Weights || a.Arena.ExplicitYukawaAmplitudesUsed || a.Arena.ObservedMassesUsed {
		t.Fatalf("bad arena: %s", FormatArena(a.Arena))
	}
}

func TestGate401Q4Target(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Q4.Degree != 4 || !a.Q4.IrreducibleOverQ || len(a.Q4.MonicCoefficients) != 5 {
		t.Fatalf("bad q4: %s", FormatQ4(a.Q4))
	}
}

func TestGate401CanonicalCompressionsDoNotSelectQ4(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	uniform := findCandidate(a.Sieve.Candidates, "uniform J-doubled edge measure")
	if !uniform.CentralOnHphi || uniform.MinimalDegree != 1 || uniform.Q4ExactMatch || uniform.PromotableAsQ4Selector {
		t.Fatalf("uniform incorrectly classified: %s", FormatCandidate(uniform))
	}
	t3 := findCandidate(a.Sieve.Candidates, "scalar branch T3/hypercharge weight")
	if !t3.PairDegenerate || t3.MinimalDegree != 2 || t3.Q4ExactMatch || t3.PromotableAsQ4Selector {
		t.Fatalf("t3 incorrectly classified: %s", FormatCandidate(t3))
	}
	branchY := findCandidate(a.Sieve.Candidates, "branch-averaged right-hypercharge edge Laplacian")
	if !branchY.PairDegenerate || branchY.MinimalDegree != 2 || branchY.Q4ExactMatch || branchY.PromotableAsQ4Selector {
		t.Fatalf("branch hypercharge incorrectly classified: %s", FormatCandidate(branchY))
	}
}

func TestGate401EdgeResolvedQuarticCapacityIsNotNativeQ4(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	edgeY := findCandidate(a.Sieve.Candidates, "edge-resolved right-hypercharge four-channel stress test")
	if !edgeY.NativeWeights || !edgeY.EdgeResolved || edgeY.CanonicalCompressionToHphi || edgeY.MinimalDegree != 4 || edgeY.CharacteristicResidualToQ4 <= 0.1 || edgeY.Q4ExactMatch || edgeY.PromotableAsQ4Selector {
		t.Fatalf("edge hypercharge incorrectly classified: %s", FormatCandidate(edgeY))
	}
	edgeY2 := findCandidate(a.Sieve.Candidates, "edge-resolved squared-hypercharge stress test")
	if !edgeY2.NativeWeights || !edgeY2.EdgeResolved || edgeY2.CanonicalCompressionToHphi || edgeY2.MinimalDegree != 4 || edgeY2.Q4ExactMatch || edgeY2.PromotableAsQ4Selector {
		t.Fatalf("edge squared hypercharge incorrectly classified: %s", FormatCandidate(edgeY2))
	}
	if a.Sieve.NativeQuarticCapacityCount == 0 || a.Sieve.CanonicalHphiQ4MatchCount != 0 {
		t.Fatalf("bad sieve summary: %s", FormatSieve(a.Sieve))
	}
}

func TestGate401SealedQ4CompanionIsQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	sealed := findCandidate(a.Sieve.Candidates, "sealed q4-weighted edge companion")
	if !sealed.Sealed || !sealed.Circular || !sealed.Q4ExactMatch || sealed.NativeWeights || sealed.CanonicalCompressionToHphi || sealed.PromotableAsQ4Selector {
		t.Fatalf("sealed q4 companion incorrectly classified: %s", FormatCandidate(sealed))
	}
}

func TestGate401FirewallStatusesAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Impact.HphiQuarticIdentified || a.Impact.CanonicalWeightedLaplacianFound || a.Impact.YukawaCouplingsReduced || a.Impact.ChargedModuliResult != 13 || !a.Impact.FlavorFirewallPreserved || !a.Impact.HiggsLanePreserved {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	statuses := Statuses(a)
	for _, req := range []string{StatusAnisotropicEdgeWeightsFound, StatusEdgeResolvedQuarticCapacity, StatusFailedHyperchargePolynomialDisjointQ4, StatusFailedNoNativeQ4WeightedLaplacian, StatusFailedNoCanonicalHphiQuarticID, StatusFirewallPreserved13Moduli} {
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
	md := RenderMarkdown(a)
	for _, want := range []string{"Gate 401 Registry Audit", "Hypercharge Laplacian", "FAILED_ROUTE_HYPERCHARGE_EDGE_POLYNOMIAL_DISJOINT_FROM_Q4", "Gate 402"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
	res := DerivedEdgeWeightOperatorHyperchargeLaplacianSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
