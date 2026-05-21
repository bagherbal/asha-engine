package generation2socketsectorvsfinitesectorledgerboundaryaudit

import (
	"strings"
	"testing"
)

func TestGate886SocketSectorAtomsInherited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Ledger.Atoms) != 3 || a.Ledger.ActiveRank != RankHRMin || !a.Ledger.CompleteOnHRMin {
		t.Fatalf("bad socket ledger: %s", FormatLedger(a.Ledger))
	}
	if a.Ledger.Atoms[0].Atom != AtomPiPlus3 || a.Ledger.Atoms[1].Atom != AtomPiMinus3 || a.Ledger.Atoms[2].Atom != AtomPiMinus1 {
		t.Fatalf("bad atom order: %s", FormatAtoms(a.Ledger.Atoms))
	}
}

func TestGate886SocketVsFiniteSectorBoundary(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.SocketSectorTyped || a.Ledger.FiniteSectorCertified || !a.Ledger.StableInOrientLayer || a.Ledger.StableInFullAF {
		t.Fatalf("bad socket/finite sector boundary: %s", FormatLedger(a.Ledger))
	}
	if a.Lift.LiftCertified || a.Lift.TargetCertified || a.Lift.NativeR3 || a.Lift.MapName != MissingLiftMap || a.Lift.Sigma != MissingSigmaMap {
		t.Fatalf("bad finite-sector lift status: %s", FormatLift(a.Lift))
	}
	if !containsAll(a.Lift.Failures, []string{FailureNoSocketToFiniteSectorMap, FailureNoFullUnbrokenAFSectorLedger, FailurePostOrientNotNativeFinite}) {
		t.Fatalf("missing lift failures: %s", FormatLift(a.Lift))
	}
}

func TestGate886NoPhysicalGenerationFlavorLeak(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Ledger.PhysicalSectors || a.Lift.PhysicalAssignment || a.Lift.GenerationCarrierPresent || a.Lift.FlavorOrientationPresent {
		t.Fatalf("physical/generation/flavor leak: %s %s", FormatLedger(a.Ledger), FormatLift(a.Lift))
	}
	for _, atom := range a.Ledger.Atoms {
		if atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue || atom.FiniteSectorCertified {
			t.Fatalf("bad atom firewall: %s", FormatAtom(atom))
		}
	}
}

func TestGate886DiagnosticsAndFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) || !near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic) {
		t.Fatalf("diagnostic drift: %s", FormatLedger(a.Ledger))
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate {
		t.Fatalf("official freeze leaked: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate886ClassificationAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Ledger.Classification != Classification || a.Ledger.NextBranch != NextBranch || !a.Ledger.R3CandidateUnderSeal || a.Ledger.NativeR3 {
		t.Fatalf("bad classification: %s", FormatLedger(a.Ledger))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate886Theorem(t *testing.T) {
	res := Generation2SocketSectorVsFiniteSectorLedgerBoundaryAuditTheorem().Verify()
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
