package hphinativescalaralgebra

import (
	"strings"
	"testing"
)

func TestGate407Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate399QuaternionicModuleAudited || !a.Inheritance.Gate400PairDegenerateResponse || !a.Inheritance.Gate404CanonicalEdgeQuotient || !a.Inheritance.Gate406Q4ContactOnly || !a.Inheritance.Gate406Q4NotHphiSelector || a.Inheritance.Gate372ChargedModuliDim != Gate372ChargedModuliDim || !a.Inheritance.NoEmpiricalInputsImported {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
}

func TestGate407GeneratorLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.Executed || a.Ledger.HphiDimension != HphiRealDim || a.Ledger.NativeHphiEndomorphismCount < 9 || a.Ledger.QuaternionicGeneratorCount != 6 || a.Ledger.PairDegenerateGeneratorCount < 3 || a.Ledger.EdgeQuotientGeneratorCount != 1 || !a.Ledger.NoQ4Imported || !a.Ledger.NoObservedInputs {
		t.Fatalf("bad ledger: %s", FormatLedger(a.Ledger))
	}
}

func TestGate407ClosureRanks(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	obs := findClosure(a.Closures, "pair-compatible observable scalar subalgebra")
	if obs.Dimension != 4 || !obs.Commutative || obs.ContainsNoncommutingPairs || !obs.PairDegeneracyClosed || obs.GenericNondegenerateCapacity || obs.CanonicalSelectorSelected {
		t.Fatalf("bad observable closure: %s", FormatClosure(obs))
	}
	left := findClosure(a.Closures, "left quaternionic action plus scalar response")
	if left.Dimension != 8 || left.Commutative || !left.ContainsNoncommutingPairs || left.PairDegeneracyClosed || !left.GenericNondegenerateCapacity || !left.RequiresCoefficientChoice {
		t.Fatalf("bad left closure: %s", FormatClosure(left))
	}
	full := findClosure(a.Closures, "full left/right quaternionic H_phi algebra plus scalar response")
	if full.Dimension != 16 || !full.FullEndRHphi || full.Commutative || !full.ContainsNoncommutingPairs || full.PairDegeneracyClosed || !full.GenericNondegenerateCapacity || full.CanonicalSelectorSelected || !full.RequiresCoefficientChoice {
		t.Fatalf("bad full closure: %s", FormatClosure(full))
	}
}

func TestGate407Selectors(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	sphi := findSelector(a.Selectors, "native scalar response S_phi")
	if !sphi.Native || !sphi.Canonical || !sphi.PairDegenerate || sphi.DistinctEigenvalueCapacity || sphi.MinimalDegree != 2 || sphi.ReducesYukawaCouplings || sphi.ReducesFlavorModuli || sphi.Verdict != StatusFailedPairDegenerateSelectorsNoFlavor {
		t.Fatalf("bad sphi: %s", FormatSelector(sphi))
	}
	generic := findSelector(a.Selectors, "generic full-algebra anisotropic element")
	if generic.Native || !generic.Sealed || generic.Canonical || !generic.UsesArbitraryCoefficients || generic.PairDegenerate || !generic.DistinctEigenvalueCapacity || generic.MinimalDegree != 4 || generic.ReducesYukawaCouplings || generic.ReducesFlavorModuli || generic.Verdict != StatusFailedGenericAnisotropyNeedsCoeffs {
		t.Fatalf("bad generic: %s", FormatSelector(generic))
	}
}

func TestGate407ImpactAndFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Impact.ChargedModuliStart != Gate372ChargedModuliDim || a.Impact.ChargedModuliResult != Gate372ChargedModuliDim || a.Impact.NativeSelectorDerived || !a.Impact.FullAlgebraNondegenerateCapacity || a.Impact.CanonicalFlavorTextureDerived || a.Impact.YukawaCouplingsReduced || a.Impact.CKMCapacityDerived || !a.Impact.ScalarSectorFlavorBlind || !a.Impact.FlavorFirewallPreserved {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaAmplitudesInserted || !a.Firewall.NoQ4HphiForcing || !a.Firewall.NoArbitraryCoefficientPromoted || !a.Firewall.NoGenericMatrixPromoted || !a.Firewall.NoFlavorModuliReductionClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate407StatusesAndNext(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(Statuses(a), "\n")
	for _, needle := range []string{StatusFullAlgebraCapacityFound, StatusObservableSubalgebraPairClosed, StatusFailedNoCanonicalNativeSelector, StatusFailedGenericAnisotropyNeedsCoeffs, StatusFailedNoYukawaCouplingReduction, StatusFirewallPreserved13Moduli} {
		if !strings.Contains(joined, needle) {
			t.Fatalf("missing status %q in\n%s", needle, joined)
		}
	}
	if a.Next.Gate != 408 || !strings.Contains(a.Next.Title, "Variational") {
		t.Fatalf("bad next: %s", FormatNext(a.Next))
	}
}

func TestGate407Theorem(t *testing.T) {
	res := HphiNativeScalarSelectorAlgebraPairDegeneracyClosureSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate407RenderMarkdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 407 Registry Audit", StatusFullAlgebraCapacityFound, StatusFailedNoCanonicalNativeSelector, StatusFirewallPreserved13Moduli, "gate=408"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
