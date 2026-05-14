package contacteigenoperatorreconstruction

import (
	"strings"
	"testing"
)

func TestGate406Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate148Q4CandidateRows || !a.Inheritance.Gate279CompanionConstructed || !a.Inheritance.Gate279IrreducibleOverQ || !a.Inheritance.Gate279NoNontrivialIdempotentQ || !a.Inheritance.Gate405NoContactEdgePullback || a.Inheritance.Gate372ChargedModuliDim != Gate372ChargedModuliDim || !a.Inheritance.NoEmpiricalInputsImported {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
}

func TestGate406ContactQ4InternalReconstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	q := a.ContactQ4
	if q.Degree != Q4Degree || q.Dimension != ContactPrimaryDim || !q.ReconstructedInternally || !q.CharacteristicMatchesQ4 || !q.MinimalMatchesQ4 || !q.IrreducibleOverQ || !q.CompanionCyclic || q.UsesHphiBasis || q.UsesEdgeBasis || q.UsesObservedInput || len(q.MonicCoefficients) != 5 {
		t.Fatalf("bad q4: %s", FormatContactQ4(q))
	}
}

func TestGate406ContactAlgebraNoNativeSplit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.ContactAlgebra
	if c.CentralizerDimensionOverQ != 4 || !c.CentralizerIsField || c.NontrivialIdempotentsOverQ != 0 || c.TwoByTwoBlockSplitOverQ || c.IndividualRootProjectorsOverQ || !c.ResolventIrreducibleOverQ || c.ResolventRootSelectedNatively || c.NativeRootSectorSemantics {
		t.Fatalf("bad contact algebra: %s", FormatContactAlgebra(c))
	}
}

func TestGate406Routes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	internal := findRoute(a.Classification.Routes, "internal contact companion eigenoperator")
	if !internal.Native || !internal.ContactInternal || !internal.PreservesQ4Internally || internal.PromotableToScalarBundle || internal.PromotableToYukawaTexture || internal.Verdict != StatusContactQ4Reconstructed {
		t.Fatalf("bad internal route: %s", FormatRoute(internal))
	}
	hphi := findRoute(a.Classification.Routes, "H_phi scalar identity selector")
	if hphi.Native || hphi.HphiSelector || !hphi.RequiresManualBasis || hphi.PromotableToScalarBundle || hphi.Verdict != StatusFailedQ4NotHphiSelector {
		t.Fatalf("bad hphi route: %s", FormatRoute(hphi))
	}
	edge := findRoute(a.Classification.Routes, "one-form edge pullback / edge-weight selector")
	if edge.Native || edge.EdgeSelector || !edge.RequiresManualBasis || edge.PromotableToScalarBundle || edge.Verdict != StatusFailedNoContactEdgePullback {
		t.Fatalf("bad edge route: %s", FormatRoute(edge))
	}
}

func TestGate406ImpactAndFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Impact.Q4InternalContactInvariant || a.Impact.Q4ScalarBundleIdentifier || a.Impact.Q4EdgeWeightOrPullback || a.Impact.ContactProjectorOrSplitDerived || a.Impact.YukawaCouplingsReduced || a.Impact.ChargedModuliResult != Gate372ChargedModuliDim || !a.Impact.FlavorFirewallPreserved || !a.Impact.ScalarHphiLanePreserved || !a.Impact.ContactLanePreserved {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaAmplitudesInserted || !a.Firewall.NoManualQ4HphiID || !a.Firewall.NoManualRootOrderingPromoted || !a.Firewall.NoResolventRootPromoted || !a.Firewall.NoArbitraryBasisPromoted || !a.Firewall.NoCompanionOperatorCrossSector || !a.Firewall.NoFlavorModuliReductionClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate406StatusesAndNext(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(Statuses(a), "\n")
	for _, needle := range []string{StatusContactQ4Reconstructed, StatusContactOnlyClassification, StatusFailedQ4NotHphiSelector, StatusFailedNoContactEdgePullback, StatusFirewallPreserved13Moduli, StatusNextHphiNativeSelectorRequired} {
		if !strings.Contains(joined, needle) {
			t.Fatalf("missing status %q in\n%s", needle, joined)
		}
	}
	if a.Next.Gate != 407 || !strings.Contains(a.Next.Title, "Hphi-Native") {
		t.Fatalf("bad next: %s", FormatNext(a.Next))
	}
}

func TestGate406Theorem(t *testing.T) {
	res := ContactEigenoperatorInternalReconstructionQ4ContactOnlyTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate406RenderMarkdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 406 Registry Audit", StatusContactQ4Reconstructed, StatusFailedQ4NotHphiSelector, StatusFirewallPreserved13Moduli, "gate=407"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
