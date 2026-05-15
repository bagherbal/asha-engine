package generation2vacuumgaugeorbitquotient

import "testing"

func TestGate497VacuumGaugeOrbitQuotient(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.LowerPairVacuumPlaneSelected || !a.Inheritance.ResidualS1PreviouslyOpen {
		t.Fatalf("Gate497 must inherit the Gate496 lower-plane/residual-S1 boundary: %+v", a.Inheritance)
	}
	if !a.ResidualPhase.BrokenNeutralMatchesPhaseTangent {
		t.Fatalf("expected Z=T3-Y_phi to sweep the residual S1 phase, residual=%g", a.ResidualPhase.BrokenNeutralMatchResidual)
	}
	if !a.ResidualPhase.PhotonStabilizesVacuum || a.ResidualPhase.PhotonImageNorm > eps {
		t.Fatalf("expected Q_em to stabilize the vacuum, norm=%g", a.ResidualPhase.PhotonImageNorm)
	}
	if !a.GaugeOrbit.GaugeOrbitRankThree || a.GaugeOrbit.OrbitImageRank != 3 {
		t.Fatalf("expected rank-three broken gauge orbit, got rank=%d", a.GaugeOrbit.OrbitImageRank)
	}
	if !a.GaugeOrbit.RadialSeparatedFromGaugeOrbit || a.GaugeOrbit.ScalarDimensionAfterQuotient != 1 {
		t.Fatalf("expected radial one-mode quotient, after=%d radial dot=%g", a.GaugeOrbit.ScalarDimensionAfterQuotient, a.GaugeOrbit.MaxRadialOrbitDot)
	}
	if !a.Representative.RepresentativeAllowedAfterQuotient || a.Representative.RepresentativeNativelySelected {
		t.Fatalf("unitary gauge representative must be bridge-allowed but not natively selected: %+v", a.Representative)
	}
	if a.Boundary.NativeResidualS1QuotientClosed || a.Boundary.NativeWZMassMatrixDerived {
		t.Fatalf("Gate497 must not promote native quotient or W/Z masses: %+v", a.Boundary)
	}
	if a.Firewall.ObservedWMassImported || a.Firewall.WeakAngleImported || a.Firewall.NativeWZMassWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
}
