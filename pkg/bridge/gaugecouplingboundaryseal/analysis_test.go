package gaugecouplingboundaryseal

import "testing"

func TestBoundarySealsAreQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ScaleSeal.ExplicitBoundaryData || !a.ScaleSeal.Quarantined || !a.ScaleSeal.RequiredForTopDownEvaluation || a.ScaleSeal.DerivedFromFiniteAlgebra || a.ScaleSeal.DerivedFromTopologicalSeal || a.ScaleSeal.ObservedValueInserted {
		t.Fatalf("bad scale seal: %s", FormatScaleSeal(a.ScaleSeal))
	}
	if !a.CouplingSeal.ExplicitBoundaryData || !a.CouplingSeal.Quarantined || !a.CouplingSeal.Dimensionless || a.CouplingSeal.DerivedFromFiniteAlgebra || a.CouplingSeal.DerivedFromEightPiSquared || a.CouplingSeal.UnitTopologicalBranchAssumed || a.CouplingSeal.ObservedValueInserted {
		t.Fatalf("bad coupling seal: %s", FormatCouplingSeal(a.CouplingSeal))
	}
}

func TestSymbolicTrajectoryBuiltButNotEvaluated(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	tr := a.Trajectory
	if !tr.BaselineBetaVector.Equal(BetaVector{U1Y: R(41, 10), SU2L: R(-19, 6), SU3C: R(-7, 1)}) {
		t.Fatalf("bad baseline beta vector: %s", FormatTrajectory(tr))
	}
	if !tr.FermionContributionVector.Equal(BetaVector{U1Y: R(4, 1), SU2L: R(4, 1), SU3C: R(4, 1)}) || tr.ThresholdRows != 12 {
		t.Fatalf("bad threshold inventory: %s", FormatTrajectory(tr))
	}
	if !tr.PiecewiseClosedFormBuilt || !tr.TreeLevelContinuityInherited || !tr.FiniteMatchingCorrectionsSealed || tr.ThresholdOrderingKnown || tr.EvaluatedNumerically || tr.PhysicalPredictionMade || tr.UsesObservedCouplings {
		t.Fatalf("trajectory leaked physical evaluation: %s", FormatTrajectory(tr))
	}
}

func TestBottomUpIsSeparatedFromTopDownDerivation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	b := a.BottomUp
	if !b.AllowedAsPhenomenologicalAudit || !b.RequiresIRCouplingSeal || !b.InvertibilityEquationsBuilt || !b.PairwiseDifferenceEquationsBuilt || !b.CanSolveFormalPairwiseLogIntervals || !b.CanTestUEqualsOneIfInputsProvided {
		t.Fatalf("bottom-up audit not constructed: %s", FormatBottomUp(b))
	}
	if b.TopologicalUOneDerived || b.IRInputsDerived || b.ThresholdOrderingKnown || b.NumericalConvergenceDetermined || b.ReducesStrictNullity {
		t.Fatalf("bottom-up audit reduced strict nullity or derived data: %s", FormatBottomUp(b))
	}
}

func TestLowEnergyAndAbsoluteFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	d := a.Domain
	if !d.FormalFermionThresholdsAvailable || d.FermionThresholdOrderingKnown || d.GaugeBosonThresholdsAvailable || !d.WZThresholdsBlocked || d.RunToMZAllowed || d.DeepInfraredFlowDefined {
		t.Fatalf("bad domain firewall: %s", FormatDomain(d))
	}
	f := a.Firewall
	if !f.Gate198Inherited || f.BoundaryScaleDerivedStrict || f.AbsoluteCouplingDerivedStrict || f.GaugeCouplingsDerived || f.TopologicalEightPiSquaredImported || f.FiniteToContinuumScaleDerived || f.ObservedInputsImported || f.PhysicalRGPredictionMade || f.NumericalTrajectoryEvaluated || f.WZThresholdsDerived || f.ThresholdOrderingDerived || f.FiniteMatchingCorrectionsDerived || f.BottomUpAuditCanDeriveFiniteTheorem {
		t.Fatalf("firewall leaked: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != 3 || f.StrictNullityAfter != 3 || f.ConditionalBoundarySealNullityBefore != 1 || f.ConditionalBoundarySealNullityAfter != 0 || f.ConditionalSymbolicEvaluationNullityBefore != 1 || f.ConditionalSymbolicEvaluationNullityAfter != 0 || f.PhysicalPredictionNullityBefore != f.PhysicalPredictionNullityAfter {
		t.Fatalf("bad nullity ledger: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := GaugeCouplingBoundarySealSymbolicRGEvaluationFirewallTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
