package generation2electroweakcurvatureaction

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate493ElectroweakCurvatureAction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate492DphiTemplateFound || !a.Inheritance.FullCurvatureNextGateRequested {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Curvature.Closed || !a.Curvature.FullFieldStrengthTyped || a.Curvature.AdjointRank != 3 || !a.Curvature.AbelianNullDirectionFound {
		t.Fatalf("bad curvature audit: %+v", a.Curvature)
	}
	if a.Curvature.SecondVariationComputed || a.Curvature.U1KineticSelected || a.Curvature.NativeCurvatureAction {
		t.Fatalf("curvature action should not be promoted: %+v", a.Curvature)
	}
	if !a.Quadratic.FullQuadraticFamilyTyped || !a.Quadratic.PositiveQuadraticFamilyExists || !a.Quadratic.AbelianCompletionTyped || !a.Quadratic.Diag114ReachableInFamily || a.Quadratic.Diag114Kappa != 6 {
		t.Fatalf("bad quadratic family: %+v", a.Quadratic)
	}
	if a.Quadratic.AbelianCoefficientSelected || a.Quadratic.GaugeKineticHessianSelected || a.Quadratic.PhysicalCouplingsOrMasses {
		t.Fatalf("quadratic family should not select physical data: %+v", a.Quadratic)
	}
	if !a.Gauge.BrokenDiag114CandidateFound || !a.Gauge.CandidatePositive || !a.Gauge.WhitenedExact || a.Gauge.SelectedByFiniteAction {
		t.Fatalf("bad gauge Hessian candidate: %+v", a.Gauge)
	}
	if !a.Coupled.CoupledActionSocketTyped || a.Coupled.PhysicalWZMassMatrixDerived || a.Coupled.WeakMixingAngleDerived {
		t.Fatalf("bad coupled socket: %+v", a.Coupled)
	}
	if a.Boundary.NativeElectroweakMassTheorem || a.Boundary.PhysicalWZMassMatrixDerived || a.Boundary.WeakMixingAngleDerived {
		t.Fatalf("boundary over-promoted electroweak theorem: %+v", a.Boundary)
	}
	if a.Firewall.ObservedWMassImported || a.Firewall.NativeWZMassWritten || a.Firewall.NativeWeakAngleWritten || a.Firewall.NativeGaugeCouplingWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestGate493RenderAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 493 Registry Audit",
		StatusFullEWConnectionClosed,
		StatusEWQuadraticFamilyTyped,
		StatusFailedAbelianCoefficientNotSelected,
		StatusFailedGaugeHessianNotSelected,
		"kappa_U1 = 6",
		"Gate 494",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
