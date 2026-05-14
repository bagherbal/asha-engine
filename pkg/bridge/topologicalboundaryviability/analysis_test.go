package topologicalboundaryviability

import "testing"

func TestEmpiricalLedgerIsQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	l := a.Ledger
	if !l.ExplicitPhenomenologicalInput || !l.Quarantined || l.UsedForFiniteDerivation || l.UsedForBoundaryDerivation {
		t.Fatalf("ledger leaked: %s", FormatLedger(l))
	}
	if l.Alpha1GUTInverse <= 0 || l.Alpha2Inverse <= 0 || l.Alpha3Inverse <= 0 || l.ScaleGeV <= 0 {
		t.Fatalf("bad ledger values: %s", FormatLedger(l))
	}
}

func TestPairwiseIntersectionsAndMismatchTriangle(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	tr := a.Triangle
	if tr.PairwiseCount != 3 || !tr.AllClosedFormSolved || !tr.UsesObservedLedger || tr.ThresholdCorrected || tr.PhysicalUnificationDerived {
		t.Fatalf("bad triangle construction: %s", FormatTriangle(tr))
	}
	if tr.SingleIntersectionFound || !tr.MismatchNonzero || tr.LogSpread <= 0 || tr.ScaleRatioMaxOverMin <= 1 || tr.TriangleAreaInLogUPlane <= 0 {
		t.Fatalf("expected nonzero mismatch triangle: %s", FormatTriangle(tr))
	}
}

func TestTopologicalUnitBenchmarkIsComparisonOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	b := a.Benchmark
	if !b.TopologicalBranchAvailableForComparison || b.TopologicalBranchAssumedAsTruth || b.FiniteDerivationClaim || b.TopologicalU != 1 {
		t.Fatalf("bad topological benchmark status: %s", FormatBenchmark(b))
	}
	if b.DeltaUFromTopologicalUnit <= 1 || b.CloseToUnitAtLooseTolerance {
		t.Fatalf("expected default comparison to be far from u*=1: %s", FormatBenchmark(b))
	}
}

func TestThresholdCorrectedEvaluationBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	x := a.ThresholdAudit
	if !x.Gate199SymbolicThresholdTreeInherited || !x.ExactThresholdBetaRowsAvailable || x.EmpiricalThresholdLedgerSupplied || x.ThresholdOrderingKnown || x.NumericalThresholdCorrectedRunAllowed || x.NumericalThresholdCorrectedRunExecuted || !x.SharpStepSchemeStillConditional || x.FiniteMatchingCorrectionsAvailable || x.WZThresholdsAvailable || x.LowEnergyZPoleStrictlyInsideDomain {
		t.Fatalf("threshold firewall leaked: %s", FormatThresholdAudit(x))
	}
}

func TestSolverAndFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.SolverAudit
	if !s.UsesExactClosedFormPairwiseLogs || s.RequiresNumericalOptimization || s.NumericalOptimizationUsed || !s.TreeLevelContinuityEnforced || !s.FiniteThresholdCorrectionsAudited || s.FiniteThresholdCorrectionsDerived {
		t.Fatalf("bad solver audit: %s", FormatSolver(s))
	}
	f := a.Firewall
	if !f.Gate199Inherited || !f.EmpiricalLedgerQuarantined || f.ObservedInputsUsedForFiniteDerivation || !f.ObservedInputsUsedForPhenomenologyOnly || f.BoundaryScaleDerived || f.AbsoluteCouplingDerived || f.TopologicalUOneDerived || f.TopologicalUOneAssumed || f.EightPiSquaredImported || f.ThresholdCorrectedPhysicalFitClaimed || f.PhysicalUnificationClaimed || f.PhysicalGaugeCouplingsDerived || f.WZThresholdsDerived || f.FiniteMatchingCorrectionsDerived || f.FiniteToContinuumNormalizationDerived {
		t.Fatalf("firewall leaked: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != 3 || f.StrictNullityAfter != 3 || f.PhenomenologyNullityBefore != 1 || f.PhenomenologyNullityAfter != 0 || f.PhysicalPredictionNullityBefore != f.PhysicalPredictionNullityAfter {
		t.Fatalf("bad nullity ledger: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := TopologicalBoundaryViabilityBottomUpConvergenceComparisonTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
