package contactquarticcompression

import "testing"

func TestGate154QuarticBlockCompressionKeepsBetaFirewallClosed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.BlockInvariant.ExactOverQ || !a.BlockInvariant.GaloisInvariant || a.BlockInvariant.SumRootsNumerator != 71 || a.BlockInvariant.SumRootsDenominator != 30 || a.BlockInvariant.ProductNumerator != 271 || a.BlockInvariant.ProductDenominator != 3240 {
		t.Fatalf("Gate 154 must preserve exact quartic block invariants: %+v", a.BlockInvariant)
	}
	if !a.Compression.BranchFree || a.Compression.RowComplete || a.Compression.CompressedPattern != "1+1+1+[4]" || a.Compression.RowLevelQuarticSemantics != 0 {
		t.Fatalf("Gate 154 must compress branch-free but not row-completely: %+v", a.Compression)
	}
	if a.BlockBeta.ThresholdBetaRowsAllowed != 0 || a.BlockBeta.BlockCanContributeAsMultiplet || a.RepresentationCompleteRows != 0 || a.ContactBetaRowsAllowed != 0 {
		t.Fatalf("Gate 154 must not open beta rows from the quartic block: blockBeta=%+v summary=%+v", a.BlockBeta, a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("Gate 154 must keep physics firewall closed: %+v", a.Summary)
	}
}
