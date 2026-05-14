package contactchargenorm

import "testing"

func TestGate145CenteredContactChargeNormalizationObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ContactRows != 7 || a.CenteredPositiveRows != 3 || a.CenteredNegativeRows != 4 || a.CenteredZeroRows != 0 {
		t.Fatalf("unexpected inherited centered split: %+v", a.Summary)
	}
	if a.MaxAbs <= 0 || a.FrobeniusNorm <= 0 || a.SpectralRange <= 0 {
		t.Fatalf("normalization scales should be positive: max=%g frob=%g range=%g", a.MaxAbs, a.FrobeniusNorm, a.SpectralRange)
	}
	if a.RawAudit.DistinctEigenvalues != 7 || a.MaxAbsAudit.DistinctEigenvalues != 7 || a.FrobeniusAudit.DistinctEigenvalues != 7 || a.RangeAudit.DistinctEigenvalues != 7 {
		t.Fatalf("diagnostic normalizations should preserve seven distinct rows")
	}
	if !a.BinaryHalfAudit.TwoLevel || !a.BinaryHalfAudit.UniformMagnitude || a.BinaryHalfAudit.TraceZero || a.BinaryHalfAudit.T3RSemantic {
		t.Fatalf("binary half-charge audit should be non-trace-zero and nonsemantic: %+v", a.BinaryHalfAudit)
	}
	if !a.BalancedSplitAudit.TwoLevel || !a.BalancedSplitAudit.TraceZero || a.BalancedSplitAudit.UniformMagnitude || a.BalancedSplitAudit.HyperchargeSemantic {
		t.Fatalf("balanced split should be trace-zero but nonsemantic and non-uniform magnitude: %+v", a.BalancedSplitAudit)
	}
	if a.NormalizationsAudited != 7 || a.AvailableNormalizations != 6 || a.CanonicalDiagnosticNorms != 6 || a.TraceZeroNormalizations != 5 || a.TwoLevelNormalizations != 2 || a.ChargeSemanticNormalizations != 0 {
		t.Fatalf("unexpected normalization audit counts: %+v", a.Summary)
	}
	if a.Requirements.AllSatisfied || a.Requirements.SelectedOrientation || a.Requirements.OperatorPullback || a.Requirements.LocalFieldMap || a.Requirements.GaugeRepresentationRows {
		t.Fatalf("charge-operator requirements should remain unsatisfied: %+v", a.Requirements)
	}
	if !a.BetaPermissionFirewallClosed || a.ContactBetaRowsAllowed != 0 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 {
		t.Fatalf("firewall should remain closed: %+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("hidden observed physics leaked into Gate 145")
	}
}
