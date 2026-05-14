package contactbranchsemantics

import "testing"

func TestGate153GaloisPartitionIsNotRowComplete(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Partition.PartitionExact || a.Partition.OrbitPattern != "1+1+1+4" || a.Partition.GaloisInvariantOrbits != 4 {
		t.Fatalf("Gate 153 must construct exact Galois-invariant 1+1+1+4 partition: %+v", a.Partition)
	}
	if a.Partition.PartitionRowComplete || a.IndividualQuarticRows != 0 || a.CanonicalQuarticBranches != 0 || a.RowwiseRootAssignmentProofs != 0 {
		t.Fatalf("Gate 153 must not split quartic orbit into individual rows: partition=%+v summary=%+v", a.Partition, a.Summary)
	}
	if a.PatternMismatch.GaloisPatternMatchesCurrent || a.PatternMismatch.GaloisPatternMatchesContact || a.PatternMismatch.GaloisPatternMatchesFano {
		t.Fatalf("Gate 153 must expose pattern mismatch, not solve it: %+v", a.PatternMismatch)
	}
	if !a.BetaPermissionFirewallClosed || a.ContactBetaRowsAllowed != 0 || a.ChargeSemanticRows != 0 || a.RepresentationCompleteRows != 0 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("Gate 153 must keep physics firewall closed: summary=%+v", a.Summary)
	}
}
