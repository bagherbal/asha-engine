package generation2minimalrightneutralabsencesealedgeskeletonaudit

import (
	"strings"
	"testing"
)

func TestGate843RectangleAndBMinusL(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Rectangle.FullRank != 8 || a.Rectangle.ActiveRank != 7 || a.Rectangle.PunctureRank != 1 || !a.Rectangle.ActiveIsFullMinusPuncture {
		t.Fatalf("bad rectangle: %s", FormatRectangle(a.Rectangle))
	}
	wantRanks := []int{3, 1, 3, 1}
	for i, want := range wantRanks {
		if a.Rectangle.Cells[i].Rank != want {
			t.Fatalf("cell %d rank got %d want %d", i, a.Rectangle.Cells[i].Rank, want)
		}
	}
	if a.Rectangle.BMinusLActive != 1 || a.Rectangle.BMinusLPuncture != -1 || a.Rectangle.BMinusLFull != 0 {
		t.Fatalf("bad B-L compensation: %s", FormatRectangle(a.Rectangle))
	}
}

func TestGate843BranchComparison(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Branches.MinimalBranchAdmittedAsSeal || a.Branches.MinimalBranchNative || a.Branches.MinimalRank != 7 {
		t.Fatalf("minimal branch invalid: %s", FormatBranches(a.Branches))
	}
	if !a.Branches.ExtendedBranchAvailable || a.Branches.ExtendedRank != 8 || a.Branches.ExtendedBranchMatchesR2Support || !a.Branches.ExtendedBranchNeedsExtraProjectionOrExclusion || !a.Branches.R2PlusPlusPrefersMinimalBranch {
		t.Fatalf("extended branch invalid: %s", FormatBranches(a.Branches))
	}
}

func TestGate843OrientationPlacementAndEdgeFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Orientation.DominantLocationSealed || !a.Orientation.RestLocationSealed || a.Orientation.NativeOrientationTheorem || a.Orientation.PhysicalParticleAssignment {
		t.Fatalf("orientation over/under-certified: %s", FormatOrientation(a.Orientation))
	}
	if a.Edge.DFEdgeGraphAvailable || a.Edge.ExplicitDFMatrixAvailable || a.Edge.NullEdgeCertified || a.Edge.MinimalAbsenceEdgeCertified || !a.Edge.AbsentNullEdgeCandidateOnly || a.Edge.PhysicalRightNeutrinoTheorem {
		t.Fatalf("edge route over-certified: %s", FormatEdge(a.Edge))
	}
	if !a.Placement.FiniteBodyLocationAtSealLevel || !a.Placement.TraceCompressionShadowAtSealLevel || a.Placement.NativeCompressionTheorem || a.Placement.AlphaDerived || a.Placement.TraceMagnitudeReadout || a.Placement.R3 || a.Placement.R4 {
		t.Fatalf("placement invalid: %s", FormatPlacement(a.Placement))
	}
}

func TestGate843FirewallsAndLedgerFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.OfficialFrozen || !a.Ledger.R2PlusPlus || a.Ledger.R3 || a.Ledger.R4 || a.Ledger.AlphaNative {
		t.Fatalf("ledger invalid: %s", FormatLedger(a.Ledger))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.AbsenceSealNotNative || !a.Firewalls.NoDFEdgeGraph || !a.Firewalls.NoNullEdgeTheorem || !a.Firewalls.NoPhysicalParticleAssignment || !a.Firewalls.NoRightNeutrinoTheorem || !a.Firewalls.CompressionSealNotNativeMap || !a.Firewalls.NoAlphaDerivation || !a.Firewalls.NoTraceMagnitudeReadout || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate843 {
		t.Fatalf("firewalls invalid: %+v", a.Firewalls)
	}
}

func TestGate843Theorem(t *testing.T) {
	res := Generation2MinimalRightNeutralAbsenceSealAndEdgeSkeletonAuditTheorem().Verify()
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
