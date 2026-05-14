package topologicalactionvariationalprinciple

import "testing"

func TestGate287TopologicalActionVariationalAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	if !a.Inheritance.NCGCalculusFormalized || !a.Inheritance.InnerFluctuationBuilt || a.Inheritance.NontrivialSaddleDerived || a.Inheritance.FourOverPiGenerated {
		t.Fatalf("bad Gate286 inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Constraint.TreatedAsDerivedFiniteDatum || a.Constraint.TreatedAsCompleteDynamics || !a.Constraint.RequiresCutoffMoments {
		t.Fatalf("bad S_top constraint: %s", FormatConstraint(a.Constraint))
	}
	if !a.MomentModel.HasTwoBranches || len(a.MomentModel.Branches) != 2 {
		t.Fatalf("expected inherited two-branch r model: %s", FormatMomentModel(a.MomentModel))
	}
	for _, b := range a.MomentModel.Branches {
		if b.ShapeResidualAbs > 1e-12 {
			t.Fatalf("branch does not satisfy shape constraint: %+v", b)
		}
	}
	if a.Variation.UniqueBranchSelected || a.Variation.SelectsUpperBranch || a.Variation.SelectsLowerBranch || a.Variation.BranchesAreShapeExtrema {
		t.Fatalf("variation should not select branch: %s", FormatVariation(a.Variation))
	}
	if !a.Variation.ArbitrarySignedMomentsCanFitAnyR || a.Variation.PositiveCutoffMomentsSelectPositiveR {
		t.Fatalf("unexpected stationarity verdict: %s", FormatVariation(a.Variation))
	}
	if !a.Rank.Underdetermined || a.Rank.CutoffMomentRatiosExtracted || a.Rank.AbsoluteScaleExtracted || a.Rank.JExtractedAsSymmetry {
		t.Fatalf("rank audit should remain underdetermined: %s", FormatRank(a.Rank))
	}
	if a.J.PhysicalJDerived || a.J.KOAxiomsVerified || a.J.OppositeActionConstructed {
		t.Fatalf("J should not be derived: %s", FormatJ(a.J))
	}
	if !a.Cutoff.InfiniteCutoffSolutions || a.Cutoff.F0F2F4RatiosExtracted || a.Cutoff.HeatKernelSubtractionDerived {
		t.Fatalf("cutoff extraction should fail: %s", FormatCutoff(a.Cutoff))
	}
	if !a.FourPi.STopCanEncodeFourOverPi || a.FourPi.ProducesInstantonLaw || a.FourPi.BGapAsInverseCouplingDerived {
		t.Fatalf("4/pi test should remain blocked: %s", FormatFourPi(a.FourPi))
	}
	if !a.Firewalls.DoesNotTreatSTopAsFullAction || !a.Firewalls.DoesNotInventCutoffMoments || !a.Firewalls.DoesNotInventPhysicalJ || !a.Firewalls.DoesNotClaimBGapInstanton || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
	if a.Summary.BranchSelected || a.Summary.PhysicalJDerived || a.Summary.FourPiInstantonDerived || a.Summary.HiggsPredictionDerived || a.Summary.IntermediateSealGranted {
		t.Fatalf("summary overclaimed: %s", FormatSummary(a.Summary))
	}
}

func TestGate287TheoremPassesChecks(t *testing.T) {
	res := TopologicalActionVariationalPrincipleBoundarySelectorAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem checks failed:\n%s", res.Details())
	}
	if res.Status != "BRIDGE_REQUIRED" {
		t.Fatalf("Gate 287 should remain BridgeRequired, got %s", res.Status)
	}
}
