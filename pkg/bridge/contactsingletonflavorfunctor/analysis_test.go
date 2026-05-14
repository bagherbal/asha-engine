package contactsingletonflavorfunctor

import (
	"strings"
	"testing"
)

func TestGate397ContactSingletonsAreNativeDomainOnly(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Singletons.NativeDomainAlgebra || a.Singletons.Dimension != 3 || a.Singletons.ExactOrthogonalIdempotents != 3 {
		t.Fatalf("bad singleton algebra: %+v", a.Singletons)
	}
	if a.Singletons.ActsOnFiniteDiracTarget || a.Singletons.NativeGenerationSemantics {
		t.Fatalf("singletons were incorrectly promoted to flavor: %+v", a.Singletons)
	}
	for _, b := range a.Singletons.Blocks {
		if !b.ProjectorExact || !b.ProjectorNative || b.RowSemantic || b.GenerationSemantic {
			t.Fatalf("bad singleton block classification: %+v", b)
		}
	}
}

func TestGate397FiniteDiracTargetUniformAndNoNativeFunctor(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Target.OneFormEdgeSupportDerived || !a.Target.EdgePatternUniform || a.Target.EdgeGenerationRank != 1 {
		t.Fatalf("bad finite-Dirac target audit: %+v", a.Target)
	}
	domain := findFunctor(a.Functors.Candidates, "contact-domain singleton algebra")
	if !domain.Native || !domain.DerivedFromContactIdempotents || domain.CompatibleWithAF || domain.CompatibleWithJ || domain.PromotableAsNativeFunctor {
		t.Fatalf("domain action incorrectly classified: %+v", domain)
	}
	edge := findFunctor(a.Functors.Candidates, "finite-Dirac edge uniform broadcast")
	if !edge.Native || !edge.CompatibleWithAF || !edge.CompatibleWithJ || !edge.CompatibleWithFirstOrder || !edge.CentralOnGeneration || edge.NoncentralOnGeneration || edge.PromotableAsNativeFunctor {
		t.Fatalf("edge broadcast incorrectly classified: %+v", edge)
	}
	if a.Functors.PromotableNativeCount != 0 || a.Functors.NativeNoncentralCount != 0 {
		t.Fatalf("unexpected promotable/native noncentral functor: %+v", a.Functors)
	}
}

func TestGate397SealedCapacityQuarantined(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	diag := findFunctor(a.Functors.Candidates, "sealed singleton-to-generation diagonal assignment")
	cycle := findFunctor(a.Functors.Candidates, "sealed singleton cyclic branch action")
	if !diag.Sealed || !diag.Circular || !diag.NoncentralOnGeneration || !diag.DiagonalOnly || diag.MixingCapacity || diag.AssignmentChoices != 6 {
		t.Fatalf("bad sealed diagonal classification: %+v", diag)
	}
	if !cycle.Sealed || !cycle.Circular || !cycle.NoncentralOnGeneration || !cycle.MixingCapacity || cycle.AssignmentChoices != 6 {
		t.Fatalf("bad sealed cycle classification: %+v", cycle)
	}
	if a.Operators.NativeNoncommutingPairs != 0 || a.Operators.CKMCapacityNative {
		t.Fatalf("unexpected native CKM capacity: %+v", a.Operators)
	}
	if a.Operators.SealedNoncommutingPairs == 0 || a.Operators.MaxSealedCommutatorNorm <= eps {
		t.Fatalf("expected sealed noncommuting stress capacity: %+v", a.Operators)
	}
}

func TestGate397ModuliFirewallAndStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Moduli.NativeReductionBelow13 || a.Moduli.BestNativeDim != 13 || a.Moduli.StartingChargedDim != 13 {
		t.Fatalf("moduli firewall changed incorrectly: %+v", a.Moduli)
	}
	statuses := Statuses(a)
	required := []string{
		StatusFailedDomainIdempotentsOnly,
		StatusFailedNoFiniteDiracActionFunctor,
		StatusFailedEdgeUniformBroadcast,
		StatusFailedAssignmentCircular,
		StatusFailedDiagonalOnlyNoCKM,
		StatusFailedNoNativeNoncommutingPair,
		StatusFirewallPreserved13Moduli,
	}
	for _, req := range required {
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

func TestGate397MarkdownAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, want := range []string{"Gate 397 Registry Audit", "rho: Q^3_contact", "FAILED_ROUTE_NO_FINITE_DIRAC_ACTION_FUNCTOR", "FIREWALL_PRESERVED_13_MODULI", "Gate 398"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
	res := ContactSingletonFiniteDiracFlavorFunctorSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit failed:\n%s", res.Details())
	}
}
