package generation2phaseorientedneutralpunctureairlockr3sealclassificationaudit

import (
	"strings"
	"testing"
)

func TestGate900MatureSealedChainComplete(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	m := a.MatureChain
	if !m.NeutralAirlockFamily || !m.RightPhaseSealRequired || !m.BoundaryAlphaReconstructed || !m.HiggsOrientationPunctureKernel || !m.ProjectorLedger || !m.PositiveReadoutRows || !m.OperatorNEffReconstructed {
		t.Fatalf("bad mature chain: %s", FormatMatureChain(m))
	}
	if !containsAll(m.Supports, []string{SupportR3SealedCandidateComplete, SupportNeutralAirlockUnifiesWounds, SupportNativeBlockersReduced}) {
		t.Fatalf("missing supports: %s", FormatMatureChain(m))
	}
}

func TestGate900OrderedRepresentativeRequiresPhaseSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Representative
	if r.Puncture != PunctureOrderedPlus || r.PhaseOrder != RightPhaseOrder || !r.HRMinComplete || r.SelectedNatively {
		t.Fatalf("bad representative: %s", FormatRepresentative(r))
	}
	if !containsAll(r.Failures, []string{FailureNoNativeSelectionSigmaPlus, FailureNoNativeRightPhaseOrientation}) {
		t.Fatalf("missing failures: %s", FormatRepresentative(r))
	}
}

func TestGate900ReadoutAndFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Readout.Positive || !a.Readout.ReproducesNEff || !near(a.Readout.OperatorNEff, OperatorNEffDiagnostic) || !near(a.Readout.OperatorCYukawa, OperatorCYukawaDiagnostic) {
		t.Fatalf("readout mismatch: %s", FormatReadout(a.Readout))
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("freeze leak: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate900PromotionChecklistBlocksNativeR3(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.Promotion
	if p.AllowedNativeR3 || p.NativeAlphaSource || p.NativePhaseOrientation || p.FullAFDescent || p.NativeSectorLedger || p.PhysicalInterpretation {
		t.Fatalf("promotion leaked: %s", FormatPromotion(p))
	}
	if !containsAll(p.Failures, []string{FailureAlphaStillSealed, FailureNoNativeNeutralPunctureAirlock, FailureNoNativeRightPhaseOrientation, FailureNoNativeBoundaryIncidenceFlag, FailureNoNativeDescentFullAF, FailureNoR4NativeYukawaTheorem}) {
		t.Fatalf("missing promotion failures: %s", FormatPromotion(p))
	}
}

func TestGate900Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate900Theorem(t *testing.T) {
	res := Generation2PhaseOrientedNeutralPunctureAirlockR3SealClassificationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
