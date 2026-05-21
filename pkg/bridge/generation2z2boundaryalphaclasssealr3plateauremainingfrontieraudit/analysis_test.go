package generation2z2boundaryalphaclasssealr3plateauremainingfrontieraudit

import (
	"strings"
	"testing"
)

func TestGate910R3ReadyUnderSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Ready
	if r.PhaseSignBlocksTraceLedger || r.BoundaryAlphaRepresentative || !r.BoundaryAlphaClassLevel || !r.TraceRowsZ2Invariant || !r.FiniteSectorZ2Ledger || !r.PositiveReadoutOnZ2Class {
		t.Fatalf("bad R3-ready sealed structure: %s", FormatReady(r))
	}
	if !r.OperatorNEffReconstructed || !r.OperatorCYukawaReconstructed || !r.OperatorCHiggsReconstructed || !near(r.Alpha, AlphaB) {
		t.Fatalf("bad operator reconstruction: %s", FormatReady(r))
	}
	if !near(r.RowWeightRest, AlphaB*(1-AlphaB)) || !near(r.RowWeightPuncture, 3*AlphaB*AlphaB) {
		t.Fatalf("bad row weights: %s", FormatReady(r))
	}
}

func TestGate910NativeBlockersAndLaterFrontier(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	b := a.Blockers
	if !b.NativeZ2BoundaryAlphaFunctorMissing || !b.NativeReducedB2FunctionalMissing || !b.FullAFDescentMissing || b.CoreBlockerCount != 3 {
		t.Fatalf("bad native blocker classification: %s", FormatBlockers(b))
	}
	if b.PhaseSignStillBlocker || b.RepresentativeAlphaStillBlocker || b.IndividualYukawaStillR3Blocker {
		t.Fatalf("loop or R4 leakage into R3 blockers: %s", FormatBlockers(b))
	}
	l := a.Later
	if !l.GenerationCarrierR4OrLater || !l.FlavorOrientationR4OrLater || !l.IndividualYukawaR4OrLater || !l.PhysicalAssignmentR4OrLater || !l.CKMPMNSR4OrLater || !l.ObservedMassSpectrumR4OrLater || l.CanEnterR4FromGate910 {
		t.Fatalf("bad later-frontier classification: %s", FormatLater(l))
	}
}

func TestGate910OfficialFreezeAndPlateau(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Freeze
	if !f.OperatorDiagnosticsOnly || !f.OfficialLedgersFrozen || f.CanUpdateOfficialNEff || f.CanUpdateCYukawaCHiggs {
		t.Fatalf("bad official freeze: %s", FormatFreeze(f))
	}
	if near(f.OperatorNEff, f.OfficialNEff) || near(f.OperatorCYukawa, f.OfficialCYukawa) || near(f.OperatorCHiggs, f.OfficialCHiggs) {
		t.Fatalf("diagnostics collapsed into official ledger: %s", FormatFreeze(f))
	}
	p := a.Plateau
	if !p.R3ReadyUnderSeal || p.NativeR3 || p.LoopBackToPhase || p.LoopBackToRepAlpha || !p.NextRailA_Z2BoundaryAlphaFirst || !p.NextRailB_FullAFSecond {
		t.Fatalf("bad plateau classification: %s", FormatPlateau(p))
	}
}

func TestGate910Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNotNativeR3, FailureNoNativeZ2BoundaryAlphaFunctor, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator} {
		if !containsAll(a.FirewallsList(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate910Theorem(t *testing.T) {
	res := Generation2Z2BoundaryAlphaClassSealR3PlateauRemainingFrontierAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing final marker %s in notes: %s", want, joined)
		}
	}
}
