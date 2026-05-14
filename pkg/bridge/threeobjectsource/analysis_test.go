package threeobjectsource

import (
	"strings"
	"testing"
)

func TestGate396FindsNativeThreeObjectSourcesButNotGeneration(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Sources.NativeExactlyThreeSourceCount < 2 {
		t.Fatalf("expected native three-object sources, got %+v", a.Sources)
	}
	if a.Sources.PromotableGenerationSourceCount != 0 || a.Sources.NativeGenerationSourceCount != 0 {
		t.Fatalf("three-object sources were incorrectly promoted: %+v", a.Sources)
	}
	contact := findSource(a.Sources.Candidates, "contact rational singleton idempotent blocks")
	if !contact.Native || !contact.ExactlyThreeObjects || !contact.ContactSemantics || contact.GenerationSemantics || contact.PromotableAsGenerationSource {
		t.Fatalf("bad contact singleton classification: %+v", contact)
	}
	color := findSource(a.Sources.Candidates, "Fock spatial color triplet")
	if !color.Native || !color.ExactlyThreeObjects || !color.ColorSemantics || color.GenerationSemantics || color.PromotableAsGenerationSource {
		t.Fatalf("bad color triplet classification: %+v", color)
	}
}

func TestGate396RejectsSelectorAndSealedRoutes(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	fano := findSource(a.Sources.Candidates, "octonionic Fano line triples")
	if !fano.RequiresSelector || fano.FamilyCount != 7 || fano.PromotableAsGenerationSource {
		t.Fatalf("Fano triples should require selector: %+v", fano)
	}
	tau := findSource(a.Sources.Candidates, "modular tau_eta three-slot scalar trace")
	n := findSource(a.Sources.Candidates, "Schrodinger/Fock information number ladder")
	if !tau.Sealed || !tau.CircularIfPromoted || !tau.NoncentralOnGenerationSpace || tau.PromotableAsGenerationSource {
		t.Fatalf("tau_eta should remain sealed/circular: %+v", tau)
	}
	if !n.Sealed || !n.CircularIfPromoted || !n.NoncentralOnGenerationSpace || n.PromotableAsGenerationSource {
		t.Fatalf("N should remain sealed/circular: %+v", n)
	}
}

func TestGate396NoNativeCKMOrModuliReduction(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Operators.NativeEligibleOperators != 0 || a.Operators.NativeNoncommutingPairs != 0 || a.Operators.CKMCapacityNative {
		t.Fatalf("unexpected native CKM capacity: %+v", a.Operators)
	}
	if a.Operators.SealedNoncommutingPairs == 0 || a.Operators.MaxSealedCommutatorNorm <= eps {
		t.Fatalf("expected sealed noncommuting stress-test capacity: %+v", a.Operators)
	}
	if a.Moduli.NativeReductionBelow13 || a.Moduli.BestNativeDim != 13 {
		t.Fatalf("native moduli firewall was incorrectly reduced: %+v", a.Moduli)
	}
}

func TestGate396Statuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{
		StatusContactSingletonsFound,
		StatusFailedContactNoFlavorFunctor,
		StatusFailedColorTripletIsColor,
		StatusFailedFanoSelectorMissing,
		StatusFailedTauOrNStillCircular,
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

func TestGate396MarkdownAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, want := range []string{"Gate 396 Registry Audit", "contact rational singleton", "FIREWALL_PRESERVED_13_MODULI", "Gate 397"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
	res := EndogenousThreeObjectSourceBeyondSpinorChiralityTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit failed:\n%s", res.Details())
	}
}
