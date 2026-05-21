package generation2operatorlevelfinitesectorprojectorledgercompatibilityaudit

import (
	"strings"
	"testing"
)

func TestGate888LedgerDefinition(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Ledger.Name != FiniteSectorLedger || len(a.Ledger.Atoms) != 3 || a.Ledger.Rank != RankHRMin || !a.Ledger.CompleteOnHRMin {
		t.Fatalf("bad ledger: %s", FormatLedger(a.Ledger))
	}
	if a.Ledger.Atoms[0].Atom != AtomPiPlus3 || a.Ledger.Atoms[1].Atom != AtomPiMinus3 || a.Ledger.Atoms[2].Atom != AtomPiMinus1 {
		t.Fatalf("bad atom order: %s", FormatLedger(a.Ledger))
	}
}

func TestGate888ProjectorCompatibility(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.Idempotent || !a.Ledger.Orthogonal || !a.Ledger.StableUnderAFOrient || a.Ledger.StableUnderFullAF || !a.Ledger.EdgeCompatible {
		t.Fatalf("bad compatibility: %s", FormatLedger(a.Ledger))
	}
	if !containsAll(a.Ledger.Supports, []string{SupportProjectorsComplete, SupportProjectorsAFOrientStable, SupportProjectorsEdgeCompatible}) {
		t.Fatalf("missing supports: %s", FormatLedger(a.Ledger))
	}
}

func TestGate888TraceReadoutAndFreeze(t *testing.T) {
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
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate {
		t.Fatalf("freeze leaked: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate888DualSealAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Seals.BoundaryAlphaSealSuppliesWeights || !a.Seals.PostOrientationSealSuppliesProjectors || a.Seals.NativeAlphaFunctorCertified || a.Seals.NativeFullAFProjectorsCertified || a.Seals.OfficialR3Eligible {
		t.Fatalf("bad dual seal: %s", FormatSeals(a.Seals))
	}
	if hasPhysicalLeak(a) {
		t.Fatalf("physical leak: %s", FormatLedger(a.Ledger))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate888Theorem(t *testing.T) {
	res := Generation2OperatorLevelFiniteSectorProjectorLedgerCompatibilityAuditTheorem().Verify()
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
