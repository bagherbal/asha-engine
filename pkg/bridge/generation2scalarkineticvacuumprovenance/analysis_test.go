package generation2scalarkineticvacuumprovenance

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate496ScalarKineticVacuumProvenance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.HessianCandidateAccepted || !a.Inheritance.Gate495MetricProvenanceOpen || !a.Inheritance.Gate495VacuumProvenanceOpen || !a.Inheritance.Gate495DphiProvenanceOpen {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Metric.MetricClassAvailable || !a.Metric.PositiveSemidefinite || !a.Metric.GhostRiskEliminatedStructurally || a.Metric.ActiveI4UnitMetricSelected || a.Metric.NumericalZHComputed || a.Metric.PhysicalKineticScaleDerived {
		t.Fatalf("bad metric audit: %+v", a.Metric)
	}
	if !a.Vacuum.LowPairSelected || a.Vacuum.LowPairDimension != 2 || !a.Vacuum.DiagnosticVacuumIsMinimizer || a.Vacuum.ResidualPhaseFreedomDimension != 1 || a.Vacuum.FiniteVacuumOrientationDerived || a.Vacuum.CanonicalPhaseSelected {
		t.Fatalf("bad vacuum audit: %+v", a.Vacuum)
	}
	if !a.ScalarSU2.AbstractDoubletRepresentation || a.ScalarSU2.FullSU2SelectedByScalarData || !a.ScalarSU2.U1PairRotationSelected || a.ScalarSU2.CanonicalComplexStructure || a.ScalarSU2.CovariantDerivativeDerived {
		t.Fatalf("bad scalar SU2 audit: %+v", a.ScalarSU2)
	}
	if !a.Boundary.MetricClassNative || a.Boundary.ActiveI4MetricNative || !a.Boundary.VacuumPlaneNative || a.Boundary.VacuumVectorNative || a.Boundary.ResidualPhaseQuotientNative || a.Boundary.KappaU1NativeSelected || a.Boundary.WZMassMatrixDerived {
		t.Fatalf("boundary over-promoted: %+v", a.Boundary)
	}
	if a.Firewall.WeakAngleImported || a.Firewall.NativeI4MetricWritten || a.Firewall.NativeVacuumVectorWritten || a.Firewall.NativeKappaWritten || a.Firewall.NativeWZMassWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestGate496RenderAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 496 Registry Audit",
		StatusHilbertSchmidtMetricClassFound,
		StatusLowerPairVacuumPlaneSelected,
		StatusFailedI4MetricNotSelected,
		StatusFailedVacuumVectorNotSelected,
		StatusNativeRegistryWriteBlocked,
		"Gate 497",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
