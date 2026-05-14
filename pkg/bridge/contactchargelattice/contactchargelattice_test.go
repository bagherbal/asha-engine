package contactchargelattice

import "testing"

func TestGate146ContactChargeLatticeObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ContactRows != 7 || a.CenteredPositiveRows != 3 || a.CenteredNegativeRows != 4 || a.CenteredZeroRows != 0 {
		t.Fatalf("unexpected inherited centered split: %+v", a.Summary)
	}
	if a.HalfIntegerAudit.ExactEmbedding || a.SixthIntegerAudit.ExactEmbedding {
		t.Fatalf("raw centered spectrum should not embed in half/sixth integer charge lattices")
	}
	if !a.SeventhBalancedAudit.ExactEmbedding || !a.SeventhBalancedAudit.AppliesToBalancedSummary || a.SeventhBalancedAudit.AppliesToRawSpectrum || a.SeventhBalancedAudit.ChargeOperatorSemantic {
		t.Fatalf("seventh-balanced audit should be summary-only and nonsemantic: %+v", a.SeventhBalancedAudit)
	}
	if !a.BoundedRationalAudit.RequiresDenominatorFit || a.BoundedRationalAudit.ExactEmbedding || a.BoundedRationalAudit.ExactRows >= a.ContactRows {
		t.Fatalf("bounded rational approximation should not exactly embed all rows: %+v", a.BoundedRationalAudit)
	}
	if !a.FreeScaledAudit.RequiresScaleChoice || !a.FreeScaledAudit.RequiresDenominatorFit || !a.FreeScaledAudit.ApproximateEmbedding || a.FreeScaledAudit.ExactEmbedding {
		t.Fatalf("free scaled candidate should remain scale-dependent only: %+v", a.FreeScaledAudit)
	}
	if !a.ObservedFitAudit.RequiresObservedInput || a.ObservedFitAudit.Available {
		t.Fatalf("observed fit should be forbidden: %+v", a.ObservedFitAudit)
	}
	if a.LatticeCandidatesAudited != 6 || a.AvailableCandidates != 5 || a.RawExactEmbeddings != 0 || a.RawApproxEmbeddings != 1 || a.BalancedExactEmbeddings != 1 || a.ChargeSemanticEmbeddings != 0 || a.ScaleDependentCandidates != 2 || a.ObservedFitCandidates != 1 {
		t.Fatalf("unexpected lattice audit counts: %+v", a.Summary)
	}
	if a.Requirements.AllSatisfied || a.Requirements.FiniteSelectedLattice || a.Requirements.RawSpectrumEmbedded || a.Requirements.PhysicalChargeSemantics || a.Requirements.OperatorPullback || a.Requirements.LocalFieldMap {
		t.Fatalf("charge lattice requirements should remain unsatisfied: %+v", a.Requirements)
	}
	if !a.BetaPermissionFirewallClosed || a.ContactBetaRowsAllowed != 0 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 {
		t.Fatalf("firewall should remain closed: %+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("hidden observed physics leaked into Gate 146")
	}
}
