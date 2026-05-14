package physicalfinitehilbertcompletion

import (
	"math"
	"testing"
)

func TestScalarMoritaQuadraticSolvesTwoPositiveBranches(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	q := a.Bridge.Quadratic
	if q.A != 3099 || q.B != -7182 || q.C != 3427 || q.Discriminant != 9100032 || !q.HasTwoPositiveRoots {
		t.Fatalf("unexpected quadratic: %s", FormatQuadratic(q))
	}
	if len(a.Bridge.Branches) != 2 {
		t.Fatalf("expected two branches: %s", FormatBridge(a.Bridge))
	}
	for _, b := range a.Bridge.Branches {
		if b.R <= 0 || b.AbsYOverX <= 0 {
			t.Fatalf("branch must be positive: %s", FormatBranch(b))
		}
		if b.ShapeResidualAbs > 1e-12 {
			t.Fatalf("branch does not reproduce shape: %s", FormatBranch(b))
		}
	}
}

func TestBranchNumericsMatchExpectedValues(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	upper := a.Bridge.Branches[0]
	lower := a.Bridge.Branches[1]
	if math.Abs(upper.R-1.645470463011191) > 1e-12 {
		t.Fatalf("upper r=%g", upper.R)
	}
	if math.Abs(lower.R-0.6720513182085573) > 1e-12 {
		t.Fatalf("lower r=%g", lower.R)
	}
	if math.Abs(upper.AbsYOverX-1.2827589263034542) > 1e-12 {
		t.Fatalf("upper |y/x|=%g", upper.AbsYOverX)
	}
	if math.Abs(lower.AbsYOverX-0.8197873615813782) > 1e-12 {
		t.Fatalf("lower |y/x|=%g", lower.AbsYOverX)
	}
}

func TestTwoBranchConstraintDoesNotBecomeHiggsPrediction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Bridge.RootsConstrainR || a.Bridge.UniqueBranchSelected || a.Bridge.AbsoluteScaleSelected || a.Bridge.A2A4Derived || a.Bridge.HiggsRatioDerived {
		t.Fatalf("bridge over-promoted: %s", FormatBridge(a.Bridge))
	}
	if !a.Firewall.ScalarMoritaBridgeMarkedConditional || !a.Firewall.CandidateMomentsNotHiggsPrediction || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
}

func TestJHyperchargeAndOppositeActionRemainIncomplete(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.J.ParticleAntiparticleTyped || a.J.AntiLinearImplemented || a.J.PhysicalHFCompleted {
		t.Fatalf("J over-derived: %s", FormatJ(a.J))
	}
	if a.Hypercharge.ChiralAssignmentDerived || a.Hypercharge.FullCPlusHPlusM3Representation || a.Hypercharge.EmpiricalAssignmentsInserted {
		t.Fatalf("hypercharge over-derived: %s", FormatHypercharge(a.Hypercharge))
	}
	if a.OppositeAction.FullOppositeActionDerived || a.OppositeAction.OrderOneReevaluatedOnFullAF || a.OppositeAction.XYRatioBranchSelectedByJ {
		t.Fatalf("opposite action over-derived: %s", FormatOpposite(a.OppositeAction))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := PhysicalFiniteHilbertSpaceChiralHyperchargeOppositeActionCompletionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
