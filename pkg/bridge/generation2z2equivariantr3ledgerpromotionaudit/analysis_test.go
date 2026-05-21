package generation2z2equivariantr3ledgerpromotionaudit

import (
	"strings"
	"testing"
)

func TestGate908OrientationClassAirlock(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.QPhiSignGaugeCandidate || a.Inherited.NativeZ2Certified {
		t.Fatalf("bad inherited audit: %s", FormatInherited(a.Inherited))
	}
	if !a.Airlock.RanksInvariant || !a.Airlock.AlphaRankSourceClassLevel || a.Airlock.RequiresAbsolutePhaseSign || a.Airlock.NativeZ2AirlockFunctor {
		t.Fatalf("bad airlock audit: %s", FormatAirlock(a.Airlock))
	}
	if len(a.Airlock.RepresentativeFlagRanks) != 2 || !sameInts(a.Airlock.RepresentativeFlagRanks[0], []int{3, 7}) || !sameInts(a.Airlock.RepresentativeFlagRanks[1], []int{3, 7}) {
		t.Fatalf("bad flag ranks: %s", FormatAirlock(a.Airlock))
	}
}

func TestGate908EdgeLedgerClass(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Edge.TauExchangesRepresentatives || !a.Edge.RankPatternsInvariant || !a.Edge.ImageRankInvariant || !a.Edge.LeftKernelInvariant || !a.Edge.OrientationClassLedgerExists || a.Edge.NativeOperatorTheorem {
		t.Fatalf("bad edge audit: %s", FormatEdge(a.Edge))
	}
	for _, r := range []Representative{a.Edge.RepresentativeA, a.Edge.RepresentativeB} {
		if !sameInts(r.RankPattern, []int{3, 3, 1}) || r.ImageRank != 7 || r.LeftKernelCount != 1 || !r.Positive {
			t.Fatalf("bad representative: %s", FormatRepresentative(r))
		}
	}
}

func TestGate908TraceMagnitudeClassInvariance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	tr := a.Trace
	if !tr.TraceInvariant || !tr.SquareInvariant || !tr.NEffInvariant || !tr.CYukawaInvariant || !tr.CHiggsInvariant || !tr.DescendsToZ2Class {
		t.Fatalf("bad trace audit: %s", FormatTrace(tr))
	}
	if !near(tr.TraceA, 3+3*AlphaB) || !near(tr.SquareTraceA, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) {
		t.Fatalf("trace formula drift: %s", FormatTrace(tr))
	}
	if !near(tr.OperatorNEffA, OperatorNEffDiagnostic) || !near(tr.OperatorCYukawaA, OperatorCYukawaDiagnostic) {
		t.Fatalf("operator diagnostic drift: %s", FormatTrace(tr))
	}
}

func TestGate908R3RequirementAndNonDescent(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.R3.ProjectorLedgerZ2Class || !a.R3.PositiveReadoutOnZ2Class || !a.R3.TraceReconstructionOnZ2Class || !a.R3.RequirementsRestatedWithoutSign || a.R3.NativeSourceCertified || a.R3.FullNativeR3 {
		t.Fatalf("bad R3 audit: %s", FormatR3(a.R3))
	}
	if a.NonDescending.SocketNamesDescend || a.NonDescending.PhysicalSectorLabelsDescend || a.NonDescending.GenerationLabelsDescend || a.NonDescending.FlavorLabelsDescend || a.NonDescending.IndividualYukawaValuesDescend || !a.NonDescending.AggregateTraceLedgerDescends {
		t.Fatalf("bad non-descending audit: %s", FormatNonDescending(a.NonDescending))
	}
}

func TestGate908FreezeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) || near(a.Freeze.OperatorCYukawa, a.Freeze.OfficialCYukawa) {
		t.Fatalf("bad freeze: %s", FormatFreeze(a.Freeze))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate908Theorem(t *testing.T) {
	res := Generation2Z2EquivariantR3LedgerPromotionAuditTheorem().Verify()
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
	if !strings.Contains(joined, FinalTruth) || !strings.Contains(joined, Classification) {
		t.Fatalf("missing final classification/truth in notes: %s", joined)
	}
}
