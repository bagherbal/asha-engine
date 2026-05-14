package scalarmoritaspectralbridge

import (
	"math"
	"testing"
)

func TestGate275BranchConstraintInherited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Inheritance.ScalarMoritaSolved || !a.Inheritance.TwoBranchXYConstrained || a.Inheritance.UniqueXYLocked || a.Inheritance.InheritedBranchCount != 2 {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
}

func TestScalarMoritaBridgeFormalizedNotA2A4(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Bridge.LambdaNumerator != 1197 || a.Bridge.LambdaDenominator != 4624 || a.Bridge.KappaC != 1 || a.Bridge.KappaQ != 3 {
		t.Fatalf("unexpected bridge constants: %s", FormatBridge(a.Bridge))
	}
	if !a.Bridge.CrossTowerBridge || !a.Bridge.ScaleFreeShapeOnly || a.Bridge.EquivalentToA2A4 {
		t.Fatalf("bridge overpromoted: %s", FormatBridge(a.Bridge))
	}
}

func TestBranchMomentDiagnosticsRemainScaleDependent(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !BranchResidualsOK(a.Branches) {
		t.Fatalf("branches do not reproduce shape: %s", FormatBranches(a.Branches))
	}
	if math.Abs(a.Branches[0].R-1.645470463011191) > 1e-12 || math.Abs(a.Branches[1].R-0.6720513182085573) > 1e-12 {
		t.Fatalf("unexpected branch roots: %s", FormatBranches(a.Branches))
	}
	for _, b := range a.Branches {
		if !b.D4OverD2DependsOnScale || b.A2A4CandidateClaimed {
			t.Fatalf("branch overpromoted: %s", FormatBranch(b))
		}
	}
}

func TestNoNativeBranchSelectorYet(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Selector.UpperBranchAllowed || !a.Selector.LowerBranchAllowed || a.Selector.UniqueBranch || a.Selector.FiniteCoreSelector == true {
		t.Fatalf("selector should leave ambiguity: %s", FormatSelector(a.Selector))
	}
}

func TestHeatKernelProjectionStillBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.HeatKernel.CanMapRawTracesToA2A4 || a.HeatKernel.CutoffMomentsSpecified || a.HeatKernel.SubtractionSchemeDerived || a.HeatKernel.GaugeKineticProjection || a.HeatKernel.ScalarFluctuationMap {
		t.Fatalf("heat kernel over-derived: %s", FormatHeatKernel(a.HeatKernel))
	}
	if a.HiggsRatio.InvariantA2A4Computed || a.HiggsRatio.HiggsMassRatioComputed {
		t.Fatalf("Higgs ratio over-claimed: %s", FormatHiggs(a.HiggsRatio))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := ScalarMoritaSpectralShapeBridgeBranchSelectorHeatKernelNormalizationAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
