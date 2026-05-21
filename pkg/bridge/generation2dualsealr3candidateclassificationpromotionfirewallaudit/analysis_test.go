package generation2dualsealr3candidateclassificationpromotionfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate889LedgerInherited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.R3SealedCandidate || a.Ledger.NativeR3 || a.Ledger.Rank != RankHRMin || len(a.Ledger.Atoms) != 3 {
		t.Fatalf("bad ledger classification: %s", FormatLedger(a.Ledger))
	}
	if a.Ledger.Name != OrientedLedger || a.Ledger.Atoms[0].Atom != AtomPiPlus3 || a.Ledger.Atoms[1].Atom != AtomPiMinus3 || a.Ledger.Atoms[2].Atom != AtomPiMinus1 {
		t.Fatalf("bad atom order: %s", FormatLedger(a.Ledger))
	}
}

func TestGate889DiagnosticsAndFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !near(a.Ledger.TraceTotal, 3+3*AlphaB) || !near(a.Ledger.SquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) {
		t.Fatalf("trace drift: %s", FormatLedger(a.Ledger))
	}
	if !near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) || !near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic) {
		t.Fatalf("diagnostic drift: %s", FormatLedger(a.Ledger))
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("freeze leak: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate889DualSealClassification(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Seals.BoundaryAlphaSealSuppliesWeights || !a.Seals.PostOrientationSealSuppliesProjectors || !a.Seals.ProjectorLedgerCompleteUnderSeals || !a.Seals.TraceReadoutCompleteUnderSeals {
		t.Fatalf("missing seal support: %s", FormatSeals(a.Seals))
	}
	if a.Seals.NativeAlphaFunctorCertified || a.Seals.NativeFullAFProjectorsCertified || a.Seals.NativeR3 || a.Seals.R4NativeYukawa {
		t.Fatalf("seal promoted incorrectly: %s", FormatSeals(a.Seals))
	}
}

func TestGate889PromotionRequirementsAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Requirements.Satisfied || !a.Requirements.NeedsNativeIncidenceFunctor || !a.Requirements.NeedsFullUnbrokenAFProjectors || !a.Requirements.NeedsSealFreeTraceMagnitudeReadout {
		t.Fatalf("bad promotion requirements: %s", FormatRequirements(a.Requirements))
	}
	if hasPhysicalLeak(a) {
		t.Fatalf("physical leak: %s", FormatLedger(a.Ledger))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate889Theorem(t *testing.T) {
	res := Generation2DualSealR3CandidateClassificationPromotionFirewallAuditTheorem().Verify()
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
