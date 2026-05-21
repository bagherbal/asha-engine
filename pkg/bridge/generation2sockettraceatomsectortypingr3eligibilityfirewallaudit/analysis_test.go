package generation2sockettraceatomsectortypingr3eligibilityfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate885SocketAtoms(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Ledger.Atoms) != 3 {
		t.Fatalf("expected three socket trace atoms: %s", FormatLedger(a.Ledger))
	}
	if a.Ledger.Atoms[0].Atom != AtomPiPlus3 || a.Ledger.Atoms[1].Atom != AtomPiMinus3 || a.Ledger.Atoms[2].Atom != AtomPiMinus1 {
		t.Fatalf("bad atom order: %s", FormatAtoms(a.Ledger.Atoms))
	}
	if a.Ledger.ActiveRank != RankHRMin || !a.Ledger.Orthogonal || !a.Ledger.CompleteOnHRMin {
		t.Fatalf("bad active ledger completeness: %s", FormatLedger(a.Ledger))
	}
}

func TestGate885PostOrientationTyping(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.StableInOrientLayer || a.Ledger.StableInFullAF {
		t.Fatalf("bad stabilizer/full A_F stability classification: %s", FormatLedger(a.Ledger))
	}
	for _, atom := range a.Ledger.Atoms {
		if !atom.StableInOrientLayer || atom.StableInFullAF || atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue {
			t.Fatalf("bad atom typing firewall: %s", FormatAtom(atom))
		}
	}
	if !containsAll(a.Ledger.Supports, []string{SupportSocketAtomsTypedInOrientLayer, SupportSocketAtomsEdgeSupportAtoms}) {
		t.Fatalf("missing supports: %s", FormatLedger(a.Ledger))
	}
}

func TestGate885ReadoutStillReconstructsOperatorLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !near(a.Ledger.TraceTotal, 3+3*AlphaB) {
		t.Fatalf("trace mismatch: %s", FormatLedger(a.Ledger))
	}
	if !near(a.Ledger.SquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) {
		t.Fatalf("square trace mismatch: %s", FormatLedger(a.Ledger))
	}
	if !near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) || !near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic) {
		t.Fatalf("operator ledger mismatch: %s", FormatLedger(a.Ledger))
	}
}

func TestGate885ClassificationAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.R3CandidateUnderSeal || a.Ledger.NativeR3 || a.Ledger.PhysicalSectors || a.Ledger.Classification != Classification {
		t.Fatalf("bad classification: %s", FormatLedger(a.Ledger))
	}
	if !containsAll(a.Ledger.Failures, []string{FailureAlphaStillSealed, FailureSocketAtomsNotNativeR3Sectors, FailureSocketAtomsNotPhysical, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}) {
		t.Fatalf("missing ledger failures: %s", FormatLedger(a.Ledger))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate885OfficialFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate {
		t.Fatalf("official freeze leaked: %s", FormatFreeze(a.Freeze))
	}
	if near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) || near(a.Freeze.OperatorCYukawa, a.Freeze.OfficialCYukawa) {
		t.Fatalf("operator and official ledgers collapsed: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate885Theorem(t *testing.T) {
	res := Generation2SocketTraceAtomSectorTypingR3EligibilityFirewallAuditTheorem().Verify()
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
