package noncartanflavorvacuum

import "testing"

func TestGate260NonCartanFlavorVacuumAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Inheritance.TauEtaRetrieved || !a.Inheritance.ConditionalU12WeakPlaneSelected || a.Inheritance.CartanNeutral3PlaneDerived {
		t.Fatalf("unexpected Gate 259 inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.NonCartan.PauliBasisRetrieved || !a.NonCartan.LieAlgebraClosed || !a.NonCartan.ChangesGaugeDirection || a.NonCartan.ChangesChargeSpectrum {
		t.Fatalf("unexpected non-Cartan audit: %s", FormatNonCartan(a.NonCartan))
	}
	if !a.GaugeOrbit.KernelDimensionGaugeInvariant || !a.GaugeOrbit.AllDirectionsMatchCartanSpectrum || a.GaugeOrbit.NonCartanCanEnlargeKernel {
		t.Fatalf("unexpected gauge orbit audit: %s", FormatGaugeOrbit(a.GaugeOrbit))
	}
	for _, d := range a.GaugeOrbit.DirectionsAudited {
		if !d.SameAsCartanSpectrum || d.CanIncreaseBeyondCartanKernel || d.YPlusHalfZeroMultiplicity != 1 || d.YMinusHalfZeroMultiplicity != 1 {
			t.Fatalf("bad direction audit: %s", FormatDirection(d))
		}
	}
	if a.EightVClosure.Neutral3PlaneAvailable || !a.Summary.EightVRouteClosed || a.Summary.EightVNeutral3PlaneDerived {
		t.Fatalf("8_v route should remain closed: %s", FormatEightVClosure(a.EightVClosure))
	}
	if a.Generation.Dimension != 3 || a.Generation.SignedDistinctEigenvalueCount != 3 || !a.Generation.NativeGenerationBreakingCapacity || !a.Generation.OperatorSpaceNotVector8V || !a.Generation.Bypasses8VNeutralKernel {
		t.Fatalf("direct generation carrier not opened correctly: %s", FormatGeneration(a.Generation))
	}
	if !a.YukawaSource.TauEtaSourceMapCandidate || !a.YukawaSource.CanBreakGenerationDegeneracy || a.YukawaSource.YukawaTextureDerived || !a.YukawaSource.RequiresFiniteYukawaAction {
		t.Fatalf("unexpected Yukawa source audit: %s", FormatYukawaSource(a.YukawaSource))
	}
	if !a.Firewall.Gate259NoGoPreserved || !a.Firewall.DoesNotTreatWpmAsChargeOperator || !a.Firewall.DoesNotPromoteGaugeRotationToNewSpectrum || !a.Firewall.DoesNotForceKernelDimThree || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}
