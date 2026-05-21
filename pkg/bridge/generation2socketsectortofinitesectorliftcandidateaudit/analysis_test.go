package generation2socketsectortofinitesectorliftcandidateaudit

import (
	"strings"
	"testing"
)

func TestGate887InheritsSocketDomain(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Domain.Atoms) != 3 || a.Domain.ActiveRank != RankHRMin || !a.Domain.CompleteOnHRMin || !a.Domain.Positive {
		t.Fatalf("bad domain: %s", FormatDomain(a.Domain))
	}
	if a.Domain.Atoms[0].Atom != AtomPiPlus3 || a.Domain.Atoms[1].Atom != AtomPiMinus3 || a.Domain.Atoms[2].Atom != AtomPiMinus1 {
		t.Fatalf("bad atom ordering: %s", FormatAtoms(a.Domain.Atoms))
	}
}

func TestGate887LiftRoutes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Lift.Routes) != 3 || a.Lift.BestRoute != RouteStabilizerSectorLift || !a.Lift.LiftCandidate || a.Lift.LiftCertified || a.Lift.NativeR3 {
		t.Fatalf("bad lift audit: %s", FormatLift(a.Lift))
	}
	if !routeOK(a.Lift.Routes, RouteStabilizerSectorLift, true, false, true) {
		t.Fatalf("bad stabilizer route: %s", FormatLift(a.Lift))
	}
	if !routeOK(a.Lift.Routes, RouteFullAFLift, false, false, false) {
		t.Fatalf("bad full AF route: %s", FormatLift(a.Lift))
	}
	if !routeOK(a.Lift.Routes, RouteEdgeSupportLift, true, false, true) {
		t.Fatalf("bad edge-support route: %s", FormatLift(a.Lift))
	}
}

func TestGate887FullAFAndPhysicalFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.Lift.Failures, []string{FailureNoFullAFLift, FailureNoNativeSocketToFiniteSectorMap, FailurePostOrientationNotNativeR3}) {
		t.Fatalf("missing lift failures: %s", FormatLift(a.Lift))
	}
	if hasPhysicalLeak(a) {
		t.Fatalf("physical/generation/flavor leak: %s %s", FormatDomain(a.Domain), FormatLift(a.Lift))
	}
	for _, atom := range a.Domain.Atoms {
		if atom.FiniteSectorCertified || atom.FullAFCertified || atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue {
			t.Fatalf("atom leaked: %s", FormatAtom(atom))
		}
	}
}

func TestGate887DiagnosticsAndFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !near(a.Domain.OperatorNEff, OperatorNEffDiagnostic) || !near(a.Domain.OperatorCYukawa, OperatorCYukawaDiagnostic) {
		t.Fatalf("diagnostic drift: %s", FormatDomain(a.Domain))
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate {
		t.Fatalf("official freeze leaked: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate887Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate887Theorem(t *testing.T) {
	res := Generation2SocketSectorToFiniteSectorLiftCandidateAuditTheorem().Verify()
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
