package hphivariationalselector

import (
	"strings"
	"testing"
)

func TestGate408Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate407FullAlgebraCapacity || !a.Inheritance.Gate407NoCanonicalSelector || !a.Inheritance.Gate407PairDegenerateSelectedObservables || !a.Inheritance.Gate407ChargedModuliPreserved || a.Inheritance.Gate372ChargedModuliDim != Gate372ChargedModuliDim || !a.Inheritance.NoEmpiricalInputsImported {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
}

func TestGate408FunctionalLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.Executed || a.Ledger.HphiDimension != HphiRealDim || a.Ledger.NativeFunctionalCount != 4 || a.Ledger.VariationalFunctionalCount != 5 || a.Ledger.ExternalSourceCount != 1 || a.Ledger.NondegenerateNativeSelectors != 0 || !a.Ledger.NoObservedInputs || !a.Ledger.NoYukawaInputs || !a.Ledger.NoArbitrarySourcesPromoted {
		t.Fatalf("bad ledger: %s", FormatLedger(a.Ledger))
	}
}

func TestGate408NativeFunctionals(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	h := functionalByName(a.Ledger.Functionals, "spectral-action Hessian on H_phi")
	if !h.Native || !h.SelectedElementUnique || !h.SelectedElementCanonical || !h.SelectedElementPairDegenerate || h.SelectedMinimalDegree != 2 || h.ReducesYukawaCouplings || h.ReducesFlavorModuli {
		t.Fatalf("bad Hessian: %s", FormatFunctional(h))
	}
	r := functionalByName(a.Ledger.Functionals, "radial scalar potential normal form")
	if !r.Native || r.SelectedElementUnique || r.MinimizerFamilyDimension != 3 || !r.SelectedElementCentral || r.SelectedMinimalDegree != 1 {
		t.Fatalf("bad radial: %s", FormatFunctional(r))
	}
	k := functionalByName(a.Ledger.Functionals, "one-form kinetic trace / complex-compatibility penalty")
	if !k.Native || k.SelectedElementUnique || k.MinimizerFamilyDimension != 4 || !k.SelectedElementPairDegenerate || k.SelectedMinimalDegree != 2 {
		t.Fatalf("bad kinetic: %s", FormatFunctional(k))
	}
	q := functionalByName(a.Ledger.Functionals, "quaternionic-invariant trace/norm functional")
	if !q.Native || !q.InvariantUnderQuaternionic || !q.SelectedElementCentral || q.SelectedMinimalDegree != 1 {
		t.Fatalf("bad quaternionic trace: %s", FormatFunctional(q))
	}
}

func TestGate408SealedSource(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := functionalByName(a.Ledger.Functionals, "sealed generic source functional stress test")
	if s.Native || !s.UsesExternalSource || !s.NondegenerateCapacity || s.SelectedElementNative || s.SelectedElementCanonical || s.SelectedElementPairDegenerate || s.SelectedMinimalDegree != 4 {
		t.Fatalf("bad source: %s", FormatFunctional(s))
	}
}

func TestGate408OutcomeImpactFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Outcome.NativeNondegenerateSelector || !a.Outcome.OnlyCentralOrPairSelected || !a.Outcome.GenericSourceWouldSelectAnyElement || a.Outcome.GenericSourcePromoted || !a.Outcome.HphiScalarLaneFlavorBlind {
		t.Fatalf("bad outcome: %s", FormatOutcome(a.Outcome))
	}
	if a.Impact.ChargedModuliStart != Gate372ChargedModuliDim || a.Impact.ChargedModuliResult != Gate372ChargedModuliDim || a.Impact.NativeNondegenerateSelector || a.Impact.YukawaCouplingsReduced || a.Impact.CKMCapacityDerived || !a.Impact.FlavorFirewallPreserved {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaAmplitudesInserted || !a.Firewall.NoExternalSourcePromoted || !a.Firewall.NoArbitraryCoefficientPromoted || !a.Firewall.NoGenericMatrixPromoted || !a.Firewall.NoFlavorModuliReductionClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate408StatusesAndNext(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(Statuses(a), "\n")
	for _, needle := range []string{StatusFunctionalLedgerAudited, StatusFailedNoUniqueVariationalSelector, StatusFailedFunctionalsSelectCentralOrPair, StatusFailedGenericSourceRequiresExternalJ, StatusFailedNoYukawaCouplingReduction, StatusFirewallPreserved13Moduli} {
		if !strings.Contains(joined, needle) {
			t.Fatalf("missing %q in\n%s", needle, joined)
		}
	}
	if a.Next.Gate != 409 || !strings.Contains(a.Next.Title, "Yukawa") {
		t.Fatalf("bad next: %s", FormatNext(a.Next))
	}
}

func TestGate408Theorem(t *testing.T) {
	res := HphiVariationalFunctionalCanonicalCoefficientSelectorSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate408RenderMarkdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 408 Registry Audit", StatusFailedNoUniqueVariationalSelector, StatusFailedGenericSourceRequiresExternalJ, StatusFirewallPreserved13Moduli, "gate=409"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
