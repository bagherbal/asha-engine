package mixededgelaplaciansieve

import (
	"strings"
	"testing"
)

func TestGate400InheritsObstructionsAndArena(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate398NoCanonicalHphiID || !a.Inheritance.Gate399QuaternionicDisjoint || !a.Inheritance.Gate385OneFormEdgeSupportDerived || a.Inheritance.Gate385JDoubledEdgeCount != 10 || a.Inheritance.Gate37HphiRealDim != 4 || a.Inheritance.Gate372ChargedModuliDim != 13 {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Arena.Formalized || !a.Arena.OneFormEdgeMeasureDerived || !a.Arena.UniformEdgeMetric || a.Arena.ExplicitDFEdgeWeightsDerived || a.Arena.PhysicalMassesInserted {
		t.Fatalf("bad arena: %s", FormatArena(a.Arena))
	}
}

func TestGate400Q4TargetIsIrreducibleQuartic(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Q4.Degree != 4 || !a.Q4.IrreducibleOverQ || !a.Q4.ContactPrimary || !a.Q4.BranchFree {
		t.Fatalf("bad q4 audit: %s", FormatQ4(a.Q4))
	}
}

func TestGate400MixedNativeCandidatesDoNotSelectQ4(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	uniform := findCandidate(a.Mixed.Candidates, "uniform one-form edge Laplacian projected to H_phi")
	if !uniform.Native || !uniform.HphiEndomorphism || !uniform.CentralOnHphi || uniform.MinimalDegree != 1 || uniform.Q4ExactMatch || uniform.PromotableAsQ4Selector {
		t.Fatalf("uniform edge candidate incorrectly classified: %s", FormatCandidate(uniform))
	}
	raw := findCandidate(a.Mixed.Candidates, "raw contact-to-scalar compression P_C Delta_E P_K")
	if !raw.Native || !raw.ContactCompressed || raw.HphiEndomorphism || raw.MinimalDegree != 0 || raw.Q4ExactMatch || raw.PromotableAsQ4Selector {
		t.Fatalf("raw compression incorrectly classified: %s", FormatCandidate(raw))
	}
	squared := findCandidate(a.Mixed.Candidates, "squared contact/edge compression scalar response")
	if !squared.Native || !squared.HphiEndomorphism || !squared.PairDegenerate || squared.MinimalDegree != 2 || squared.IrreducibleQuartic || squared.Q4ExactMatch || squared.PromotableAsQ4Selector {
		t.Fatalf("squared compression incorrectly classified: %s", FormatCandidate(squared))
	}
}

func TestGate400SealedQ4CompanionIsQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	sealed := findCandidate(a.Mixed.Candidates, "sealed q4 companion operator declared on H_phi")
	if !sealed.Sealed || !sealed.Circular || !sealed.Q4ExactMatch || sealed.Native || sealed.CompatibleWithJ || sealed.CompatibleWithFirstOrder || sealed.PromotableAsQ4Selector {
		t.Fatalf("sealed q4 companion incorrectly classified: %s", FormatCandidate(sealed))
	}
	if a.Mixed.NativeQ4MatchCount != 0 || a.Mixed.PromotableNativeCount != 0 || a.Mixed.SealedQ4MatchCount != 1 {
		t.Fatalf("bad mixed summary: %s", FormatMixed(a.Mixed))
	}
}

func TestGate400FirewallAndStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Impact.HphiQuarticIdentified || a.Impact.OneFormEdgeFunctorDerived || a.Impact.YukawaCouplingsReduced || a.Impact.ChargedModuliResult != 13 || !a.Impact.FlavorFirewallPreserved || !a.Impact.HiggsLanePreserved {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	statuses := Statuses(a)
	for _, req := range []string{StatusFailedUniformEdgeLaplacianCentral, StatusFailedContactCompressionNotEndomorphism, StatusFailedMixedOperatorMinimalPolynomialNotQ4, StatusFailedPairDegeneratePolynomialNotIrreducibleQ4, StatusFailedNoNativeQ4ScalarSelector, StatusFailedNoCanonicalHphiQuarticID, StatusFirewallPreserved13Moduli} {
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

func TestGate400MarkdownAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, want := range []string{"Gate 400 Registry Audit", "Mixed Edge Laplacian", "FAILED_ROUTE_UNIFORM_EDGE_LAPLACIAN_IS_CENTRAL_ON_HPHI", "FAILED_ROUTE_NO_NATIVE_Q4_SCALAR_SELECTOR", "Gate 401"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
	res := NonQuaternionicScalarIdentityMixedEdgeLaplacianSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit failed:\n%s", res.Details())
	}
}
