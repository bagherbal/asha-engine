package generation2scalaredgestability

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate491ScalarEdgeStability(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate490AnomalyLedgerStable || !a.Inheritance.NoObservedFlavorDataImported {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Support.HiggsIsFiniteOneForm || !a.Support.EdgeMeasureSelected || a.Support.NodeMeasureAdmissible || a.Support.JDoubledEdgeCount != 10 {
		t.Fatalf("bad one-form support: %+v", a.Support)
	}
	if !a.Kinetic.PositiveSemidefinite || a.Kinetic.NegativeTermsPermitted || a.Kinetic.ImaginaryKineticPermitted || !a.Kinetic.GhostRiskEliminated {
		t.Fatalf("bad kinetic positivity: %+v", a.Kinetic)
	}
	if a.Kinetic.NumericalZHComputed || a.Kinetic.StrictPositiveProvedNumerically || !a.Kinetic.YukawaAmplitudesSealed {
		t.Fatalf("numeric ZH should remain sealed: %+v", a.Kinetic)
	}
	if !a.Goldstone.CountResonance || a.Goldstone.GaugeEatingTheoremDerived || a.Goldstone.CovariantDerivativeDerived {
		t.Fatalf("bad Goldstone boundary: %+v", a.Goldstone)
	}
	if a.Boundary.FullHessianDerived || a.Boundary.HiggsMassDerived || a.Boundary.ContinuumScalarMatchingComplete {
		t.Fatalf("scalar boundary overclaimed: %+v", a.Boundary)
	}
	if a.Firewall.ObservedHiggsMassImported || a.Firewall.NativeQuarticMassWritten || a.Firewall.NativeFlavorModuliChanged {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestGate491RenderAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 491 Registry Audit",
		StatusScalarKineticTracePositive,
		StatusFailedFullScalarHessianNotDerived,
		"Goldstone",
		"Gate 492",
		"positive-semidefinite",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
