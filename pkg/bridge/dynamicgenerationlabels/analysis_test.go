package dynamicgenerationlabels

import (
	"strings"
	"testing"
)

func TestGate395SpinorSplitDoesNotGiveThreeGenerations(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Spinor.FullSpinorRealDimension != 16 {
		t.Fatalf("full spinor dim=%d", a.Spinor.FullSpinorRealDimension)
	}
	if got := a.Spinor.ChiralSplit; len(got) != 2 || got[0] != 8 || got[1] != 8 {
		t.Fatalf("bad chiral split: %v", got)
	}
	if a.Spinor.HasThreeNativeSectors || a.Spinor.GenerationLabelsDerived {
		t.Fatalf("spinor split was incorrectly promoted to three generations: %+v", a.Spinor)
	}
}

func TestGate395TrialityArenaNotNativeFlavorCarrier(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Triality.CategoryLevelTriple {
		t.Fatal("expected Spin(8) triality category triple")
	}
	if a.Triality.NativeFunctorToC3Gen || a.Triality.ExplicitThetaOnFiniteDiracFlavor {
		t.Fatalf("triality was incorrectly promoted: %+v", a.Triality)
	}
	branch := findLabel(a.Labels.Candidates, "triality representation-type triple")
	if !branch.Sealed || !branch.Circular || branch.Native {
		t.Fatalf("branch label action must remain sealed/circular: %+v", branch)
	}
}

func TestGate395NoNativeDynamicGenerationLabels(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Labels.NativeGenerationLabelCount != 0 {
		t.Fatalf("unexpected native generation labels: %+v", a.Labels)
	}
	if a.Operators.NativeNoncommutingPairs != 0 {
		t.Fatalf("unexpected native noncommuting pairs: %+v", a.Operators)
	}
	if a.Operators.CKMCapacityNative {
		t.Fatal("CKM capacity must not be native")
	}
	if a.Moduli.NativeReductionBelow13 || a.Moduli.BestNativeDim != 13 {
		t.Fatalf("moduli firewall rewritten: %+v", a.Moduli)
	}
}

func TestGate395SealedCapacityQuarantined(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Labels.SealedNoncentralCount < 2 {
		t.Fatalf("expected sealed noncentral stress-test operators: %+v", a.Labels)
	}
	if a.Operators.SealedNoncommutingPairs == 0 || a.Operators.MaxSealedCommutatorNorm <= eps {
		t.Fatalf("expected sealed noncommuting capacity: %+v", a.Operators)
	}
	if !a.Firewall.NoTrialityLabelsPromoted || !a.Firewall.NoNPromoted {
		t.Fatalf("sealed operators were promoted: %+v", a.Firewall)
	}
}

func TestGate395TheoremChecksPassAsFailedRoute(t *testing.T) {
	res := RepresentationOriginDynamicGenerationLabelsTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem checks failed:\n%s", res.Details())
	}
	if string(res.Status) != "FAILED_ROUTE" {
		t.Fatalf("expected FAILED_ROUTE status, got %s", res.Status)
	}
}

func TestRenderMarkdownContainsFirewallAndNextGate(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 395 Registry Audit", "FAILED_ROUTE_SPINOR_DECOMPOSITION_IS_TWO_SECTOR_NOT_THREE_GENERATION", "FIREWALL_PRESERVED_13_MODULI", "Gate 396"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
