package contactquarticscalaryukawabundle

import (
	"strings"
	"testing"
)

func TestGate398QuarticAndScalarDimensionMatchButNoFunctor(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Quartic.Dimension != 4 || !a.Quartic.GaloisSafePrimary || !a.Quartic.BranchFreeBlock || !a.Quartic.AbstractRankOneModule || !a.Quartic.CompanionRepresentation {
		t.Fatalf("bad quartic primary audit: %s", FormatQuartic(a.Quartic))
	}
	if a.Quartic.CanonicalHphiIdentification || a.Quartic.ScalarMinimalPolynomialDerived {
		t.Fatalf("quartic block incorrectly promoted to H_phi: %s", FormatQuartic(a.Quartic))
	}
	if a.Scalar.ActiveRealDim != 4 || a.Scalar.ComplexDoubletDim != 2 || !a.Scalar.NormalFormAvailable || a.Scalar.CanonicalQuarticAction {
		t.Fatalf("bad scalar carrier audit: %s", FormatScalar(a.Scalar))
	}
}

func TestGate398CandidateClassifications(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	abstract := findFunctor(a.Functors.Candidates, "abstract quartic primary module")
	if !abstract.Native || !abstract.DimensionCompatible || !abstract.AlgebraHomomorphism || !abstract.ProjectiveModule || abstract.PhysicalCarrierAction || abstract.PromotableAsNativeFunctor {
		t.Fatalf("bad abstract module classification: %s", FormatFunctor(abstract))
	}
	dimOnly := findFunctor(a.Functors.Candidates, "dimension-only quartic to H_phi identification")
	if !dimOnly.Native || !dimOnly.DimensionCompatible || dimOnly.AlgebraHomomorphism || dimOnly.PhysicalCarrierAction || dimOnly.ScalarMinimalPolynomial || dimOnly.PromotableAsNativeFunctor {
		t.Fatalf("bad dimension-only classification: %s", FormatFunctor(dimOnly))
	}
	sealed := findFunctor(a.Functors.Candidates, "sealed companion operator on H_phi stress test")
	if !sealed.Sealed || !sealed.Circular || !sealed.ArbitraryBasisIdentification || !sealed.AlgebraHomomorphism || !sealed.ScalarMinimalPolynomial || sealed.PromotableAsNativeFunctor {
		t.Fatalf("bad sealed companion classification: %s", FormatFunctor(sealed))
	}
}

func TestGate398OneFormAndYukawaRoutesBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Target.OneFormEdgeSupportDerived || a.Target.JDoubledEdgeCount != 10 || a.Target.YukawaChannels != 8 || a.Target.ScalarFiberEntries != 16 {
		t.Fatalf("bad bundle target: %s", FormatTarget(a.Target))
	}
	edge := findFunctor(a.Functors.Candidates, "quartic primary to one-form edge module")
	if !edge.Native || edge.DimensionCompatible || edge.Rank != 10 || edge.CompatibleWithOneFormEdges || edge.PromotableAsNativeFunctor {
		t.Fatalf("edge route should be blocked: %s", FormatFunctor(edge))
	}
	yukawa := findFunctor(a.Functors.Candidates, "quartic primary weighting of Yukawa fibers")
	if !yukawa.Native || yukawa.DimensionCompatible || yukawa.Rank != 16 || yukawa.ReducesYukawaCouplings || yukawa.ReducesFlavorModuli || yukawa.PromotableAsNativeFunctor {
		t.Fatalf("yukawa route should be blocked: %s", FormatFunctor(yukawa))
	}
}

func TestGate398ImpactFirewallAndStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Functors.PromotableNativeCount != 0 || a.Functors.PhysicalScalarActions != 0 || a.Functors.YukawaReducingActions != 0 {
		t.Fatalf("unexpected promotable functor: %s", FormatFunctors(a.Functors))
	}
	if a.Impact.NativeFlavorReduction || a.Impact.BestNativeModuliDim != 13 || !a.Impact.ScalarHiggsLanePreserved {
		t.Fatalf("bad impact audit: %s", FormatImpact(a.Impact))
	}
	statuses := Statuses(a)
	for _, req := range []string{StatusFailedNoCanonicalHphiID, StatusFailedNoQuarticMinimalScalarOp, StatusFailedNoOneFormEdgeFunctor, StatusFailedNoYukawaCouplingReduction, StatusFirewallPreserved13Moduli} {
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

func TestGate398MarkdownAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, want := range []string{"Gate 398 Registry Audit", "rho_4", "FAILED_ROUTE_NO_CANONICAL_HPHI_IDENTIFICATION", "FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION", "Gate 399"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
	res := ContactQuarticPrimaryScalarYukawaBundleFunctorAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit failed:\n%s", res.Details())
	}
}
