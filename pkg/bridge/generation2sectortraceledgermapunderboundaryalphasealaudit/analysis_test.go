package generation2sectortraceledgermapunderboundaryalphasealaudit

import (
	"strings"
	"testing"
)

func TestGate883SocketTraceAtoms(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Ledger.Atoms) != 3 {
		t.Fatalf("expected three active socket trace atoms: %s", FormatLedger(a.Ledger))
	}
	if a.Ledger.Atoms[0].Name != AtomPiPlus3 || a.Ledger.Atoms[1].Name != AtomPiMinus3 || a.Ledger.Atoms[2].Name != AtomPiMinus1 {
		t.Fatalf("bad atom order: %s", FormatAtoms(a.Ledger.Atoms))
	}
	if a.Ledger.ActiveRank != 7 || !a.Ledger.CompleteOnHRMin || !a.Ledger.Orthogonal {
		t.Fatalf("bad active ledger decomposition: %s", FormatLedger(a.Ledger))
	}
}

func TestGate883PositiveWeightsAndTraceReadout(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, atom := range a.Ledger.Atoms {
		if atom.Weight <= 0 {
			t.Fatalf("non-positive atom: %s", FormatAtom(atom))
		}
	}
	if !near(a.Ledger.TraceTotal, 3+3*AlphaB) {
		t.Fatalf("trace mismatch: %s", FormatLedger(a.Ledger))
	}
	if !near(a.Ledger.SquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) {
		t.Fatalf("square trace mismatch: %s", FormatLedger(a.Ledger))
	}
	if !near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) || !near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic) {
		t.Fatalf("operator readout mismatch: %s", FormatLedger(a.Ledger))
	}
}

func TestGate883NotNativeR3(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Ledger.NativeR3 || a.Ledger.OfficialUpdatesAllowed || a.Ledger.Classification != Classification {
		t.Fatalf("bad classification: %s", FormatLedger(a.Ledger))
	}
	if !containsAll(a.Ledger.Failures, []string{FailureAlphaStillSealed, FailureSocketProjectorsNotPhysical, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}) {
		t.Fatalf("missing firewalled failures: %s", FormatLedger(a.Ledger))
	}
}

func TestGate883OfficialFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate {
		t.Fatalf("official freeze leaked: %s", FormatFreeze(a.Freeze))
	}
	if near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) || near(a.Freeze.OperatorCYukawa, a.Freeze.OfficialCYukawa) {
		t.Fatalf("operator and official ledger collapsed: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate883Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate883Theorem(t *testing.T) {
	res := Generation2SectorTraceLedgerMapUnderBoundaryAlphaSealAuditTheorem().Verify()
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
