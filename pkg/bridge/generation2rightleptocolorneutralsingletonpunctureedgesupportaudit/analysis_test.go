package generation2rightleptocolorneutralsingletonpunctureedgesupportaudit

import (
	"strings"
	"testing"
)

func TestGate842FourCellLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Cells.FourCellLedgerCertified || !a.Cells.Gate841Inherited || !a.Cells.ActiveMinusPunctureForm {
		t.Fatalf("ledger not certified at support-anatomy level: %s", FormatCells(a.Cells))
	}
	if a.Cells.FullRank != 8 || a.Cells.ActiveRank != 7 || a.Cells.PunctureRank != 1 || len(a.Cells.Cells) != 4 {
		t.Fatalf("bad four-cell ranks: %s", FormatCells(a.Cells))
	}
	wantRanks := []int{3, 1, 3, 1}
	for i, want := range wantRanks {
		if a.Cells.Cells[i].Rank != want {
			t.Fatalf("cell %d rank got %d want %d", i, a.Cells.Cells[i].Rank, want)
		}
	}
	if !a.Cells.Cells[1].Puncture || !a.Cells.Cells[1].Leptonic || !a.Cells.Cells[1].Colorless {
		t.Fatalf("puncture cell not isolated: %+v", a.Cells.Cells[1])
	}
}

func TestGate842BMinusLPattern(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.BMinusL.EPlusP3Trace != 1 || a.BMinusL.EPlusP1Trace != -1 || a.BMinusL.EMinusP3Trace != 1 || a.BMinusL.EMinusP1Trace != -1 {
		t.Fatalf("bad cell B-L traces: %s", FormatBMinusL(a.BMinusL))
	}
	if a.BMinusL.ActiveTrace != 1 || a.BMinusL.PunctureTrace != -1 || a.BMinusL.FullTrace != 0 || !a.BMinusL.ActivePlusPunctureCancel || !a.BMinusL.FullNeutral || !a.BMinusL.CompensatingSingletonPattern {
		t.Fatalf("bad aggregate B-L pattern: %s", FormatBMinusL(a.BMinusL))
	}
}

func TestGate842EdgeAndOrientationBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Orientation.UnorderedPairCertified || !a.Orientation.EPlusColorWithEPlusLeptonPuncture || !a.Orientation.EMinusFullWQuartetCandidate {
		t.Fatalf("orientation candidate not set: %s", FormatOrientation(a.Orientation))
	}
	if a.Orientation.OrderedPhysicalOrientationCertified || a.Orientation.DominantColorOrientationCertified || a.Orientation.RestQuartetOrientationCertified {
		t.Fatalf("orientation over-certified: %s", FormatOrientation(a.Orientation))
	}
	if a.Edge.DFEdgeGraphAvailable || a.Edge.NullEdgeCertified || a.Edge.MinimalAbsenceCertified || a.Edge.SterilePunctureCertified || a.Edge.PhysicalAssignmentCertified {
		t.Fatalf("edge/sterile route over-certified: %s", FormatEdge(a.Edge))
	}
	if !containsAll(a.Edge.Failures, []string{FailureNoDFEdgeGraph, FailureNoNullEdgeTheorem, FailureNoMinimalAbsenceTheorem, FailureNoRightNeutrinoTheorem}) {
		t.Fatalf("missing edge failures: %s", strings.Join(a.Edge.Failures, ","))
	}
}

func TestGate842AggregatePlacementAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Placement.FiniteBodyLocationCandidate || a.Placement.TopRank != 3 || a.Placement.RestRank != 4 || a.Placement.TotalRank != 7 {
		t.Fatalf("bad placement candidate: %s", FormatPlacement(a.Placement))
	}
	if a.Placement.OrientedByNullEdgeCertified || a.Placement.CompressionMapCertified || a.Placement.TraceCompressionShadowCertified || a.Placement.AlphaDerivedByCompression || a.Placement.TraceMagnitudeReadoutCertified || a.Placement.R3 || a.Placement.R4 {
		t.Fatalf("placement over-certified: %s", FormatPlacement(a.Placement))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoDFEdgeGraph || !a.Firewalls.NoNullEdgeTheorem || !a.Firewalls.NoSterilePunctureTheorem || !a.Firewalls.PunctureNotPhysicalParticle || !a.Firewalls.NoRightNeutrinoTheorem || !a.Firewalls.NoTypedCompressionMap || !a.Firewalls.AlphaSealed || !a.Firewalls.NoTraceMagnitudeReadout || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate842 {
		t.Fatalf("firewalls invalid: %+v", a.Firewalls)
	}
}

func TestGate842Theorem(t *testing.T) {
	res := Generation2RightLeptoColorNeutralSingletonPunctureEdgeSupportAuditTheorem().Verify()
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
