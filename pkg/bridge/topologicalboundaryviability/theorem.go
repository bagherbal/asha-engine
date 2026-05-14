package topologicalboundaryviability

import "github.com/bagherbal/asha-engine/pkg/theorem"

func TopologicalBoundaryViabilityBottomUpConvergenceComparisonTheorem() theorem.Theorem {
	const id = "BRIDGE-TOPOLOGICAL-BOUNDARY-VIABILITY-BOTTOM-UP-CONVERGENCE-COMPARISON-AUDIT"
	const name = "topological boundary viability / bottom-up convergence comparison audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: []theorem.Check{{Name: "build topological boundary viability audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: []theorem.Check{
			{Name: "Gate 199 symbolic RG boundary scaffold is inherited", Passed: a.Firewall.Gate199Inherited && a.PreviousGate199.Summary.SymbolicTopDownTrajectoryBuilt && a.PreviousGate199.Summary.BottomUpViabilityAuditSeparated && !a.PreviousGate199.Firewall.ObservedInputsImported && !a.PreviousGate199.Firewall.PhysicalRGPredictionMade, Detail: a.PreviousGate199.TruthStatement},
			{Name: "empirical Z-pole ledger is quarantined as phenomenology", Passed: a.Ledger.ExplicitPhenomenologicalInput && a.Ledger.Quarantined && !a.Ledger.UsedForFiniteDerivation && !a.Ledger.UsedForBoundaryDerivation && a.Ledger.Alpha1GUTInverse > 0 && a.Ledger.Alpha2Inverse > 0 && a.Ledger.Alpha3Inverse > 0, Detail: FormatLedger(a.Ledger)},
			{Name: "pairwise one-loop UV intersections are solved in closed form", Passed: a.Triangle.PairwiseCount == 3 && a.Triangle.AllClosedFormSolved && a.Triangle.UsesObservedLedger && !a.Triangle.ThresholdCorrected && !a.Triangle.PhysicalUnificationDerived, Detail: FormatTriangle(a.Triangle)},
			{Name: "nonzero mismatch triangle is reported instead of perfect unification", Passed: !a.Triangle.SingleIntersectionFound && a.Triangle.MismatchNonzero && a.Triangle.LogSpread > 0 && a.Triangle.TriangleAreaInLogUPlane > 0, Detail: FormatTriangle(a.Triangle)},
			{Name: "topological u*=1 branch is benchmarked but not assumed", Passed: a.Benchmark.TopologicalBranchAvailableForComparison && !a.Benchmark.TopologicalBranchAssumedAsTruth && !a.Benchmark.FiniteDerivationClaim && a.Benchmark.TopologicalU == 1 && a.Benchmark.DeltaUFromTopologicalUnit > 1 && !a.Benchmark.CloseToUnitAtLooseTolerance, Detail: FormatBenchmark(a.Benchmark)},
			{Name: "threshold-corrected numerical evaluation remains blocked", Passed: a.ThresholdAudit.Gate199SymbolicThresholdTreeInherited && a.ThresholdAudit.ExactThresholdBetaRowsAvailable && !a.ThresholdAudit.EmpiricalThresholdLedgerSupplied && !a.ThresholdAudit.ThresholdOrderingKnown && !a.ThresholdAudit.NumericalThresholdCorrectedRunAllowed && !a.ThresholdAudit.NumericalThresholdCorrectedRunExecuted && a.ThresholdAudit.SharpStepSchemeStillConditional && !a.ThresholdAudit.FiniteMatchingCorrectionsAvailable && !a.ThresholdAudit.WZThresholdsAvailable && !a.ThresholdAudit.LowEnergyZPoleStrictlyInsideDomain, Detail: FormatThresholdAudit(a.ThresholdAudit)},
			{Name: "closed-form solver and tree-level continuity are separated from finite matching corrections", Passed: a.SolverAudit.UsesExactClosedFormPairwiseLogs && !a.SolverAudit.RequiresNumericalOptimization && !a.SolverAudit.NumericalOptimizationUsed && a.SolverAudit.TreeLevelContinuityEnforced && a.SolverAudit.FiniteThresholdCorrectionsAudited && !a.SolverAudit.FiniteThresholdCorrectionsDerived, Detail: FormatSolver(a.SolverAudit)},
			{Name: "finite theorem and physical prediction firewalls remain sealed", Passed: a.Firewall.EmpiricalLedgerQuarantined && !a.Firewall.ObservedInputsUsedForFiniteDerivation && a.Firewall.ObservedInputsUsedForPhenomenologyOnly && !a.Firewall.BoundaryScaleDerived && !a.Firewall.AbsoluteCouplingDerived && !a.Firewall.TopologicalUOneDerived && !a.Firewall.TopologicalUOneAssumed && !a.Firewall.EightPiSquaredImported && !a.Firewall.ThresholdCorrectedPhysicalFitClaimed && !a.Firewall.PhysicalUnificationClaimed && !a.Firewall.PhysicalGaugeCouplingsDerived && !a.Firewall.WZThresholdsDerived && !a.Firewall.FiniteMatchingCorrectionsDerived && !a.Firewall.FiniteToContinuumNormalizationDerived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3 && a.Firewall.PhenomenologyNullityBefore == 1 && a.Firewall.PhenomenologyNullityAfter == 0 && a.Firewall.PhysicalPredictionNullityBefore == a.Firewall.PhysicalPredictionNullityAfter, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records phenomenological diagnostic without finite overclaim", Passed: a.Summary.TestsAudited == 7 && a.Summary.EmpiricalLedgerQuarantined && a.Summary.PairwiseIntersectionsSolved && a.Summary.MismatchTriangleNonzero && a.Summary.TopologicalBenchmarkComputed && a.Summary.ThresholdCorrectedEvaluationBlocked && a.Summary.NoFiniteDerivationClaim, Detail: FormatSummary(a.Summary)},
		}, Notes: []string{
			a.TruthStatement,
			"Solver answer: the current fixed-region Gate 200 comparison uses exact closed-form pairwise logarithmic intersections; numerical optimization is reserved for future scans with explicit empirical threshold ledgers and matching conventions.",
			"Matching answer: tree-level continuity is enforced as inherited scaffolding; finite threshold matching corrections are audited but remain scheme-dependent and not finite-derived.",
		}}
	}}
}
