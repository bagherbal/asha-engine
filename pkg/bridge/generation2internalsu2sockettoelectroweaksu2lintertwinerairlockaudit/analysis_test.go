package generation2internalsu2sockettoelectroweaksu2lintertwinerairlockaudit

import (
	"strings"
	"testing"
)

func TestGate716SU2AlgebraAndRepresentationCompatibility(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.SU2DoubletSocketInherited || a.Inherited.InternalCommutantDimension != 3 || !a.Inherited.InternalDoubletShapeCertified || !a.Inherited.InternalCTraceZero || !a.Inherited.InternalCTwistorInvariant || a.Inherited.InternalCPhysicalSU2L || a.Inherited.CanonicalThetaSU2 || !a.Inherited.U1PhaseSelectorDependent || a.Inherited.HyperchargeCertified || a.Inherited.TypedHiggsDoubletMap {
		t.Fatalf("bad Gate715 inheritance: %+v", a.Inherited)
	}
	if !a.Electroweak.TargetLaneIdentified || a.Electroweak.TargetAlgebra != "su(2)_L" || a.Electroweak.TargetComplexDimension != 2 || !a.Electroweak.AlreadyDerivedAsFiniteLane || a.Electroweak.ImportsObservedData || a.Electroweak.IncludesHypercharge || a.Electroweak.DerivesMassOrRuntime {
		t.Fatalf("bad electroweak target lane: %+v", a.Electroweak)
	}
	if !a.Algebra.BothCompactSU2Type || !a.Algebra.ExistsUpToBasis || a.Algebra.CanonicalPhiSelected || !strings.Contains(a.Algebra.BracketPreservingPhi, "phi_SU2") {
		t.Fatalf("bad algebra isomorphism audit: %+v", a.Algebra)
	}
	if !a.Intertwiner.RepresentationCompatible || !a.Intertwiner.ComplexLinearIsomorphism || a.Intertwiner.InternalComplexDimension != 2 || a.Intertwiner.TargetComplexDimension != 2 || !a.Intertwiner.InternalActionIrreducible || !a.Intertwiner.TargetActionDoubletShaped || a.Intertwiner.PhysicalHiggsMapCertified || !strings.Contains(a.Intertwiner.Condition, "Theta_H_SU2") {
		t.Fatalf("bad representation intertwiner audit: %+v", a.Intertwiner)
	}
}

func TestGate716FirewallsAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Noncanonical.SU2AutomorphismFreedom || !a.Noncanonical.ComplexUnitaryBasisFreedom || !a.Noncanonical.TwistorJHChoiceFreedom || !a.Noncanonical.GeneratorNormalizationFreedom || !a.Noncanonical.MovingU1PhaseFreedom || a.Noncanonical.CanonicalThetaSU2Selected {
		t.Fatalf("noncanonical firewall violated: %+v", a.Noncanonical)
	}
	if a.Hypercharge.SpanJHEqualsPhysicalU1Y || a.Hypercharge.InternalU1PhaseEqualsHypercharge || a.Hypercharge.HyperchargeAssignment || a.Hypercharge.HyperchargeNormalization || a.Hypercharge.FullPhysicalHiggsDoubletCertified || !strings.Contains(a.Hypercharge.Verdict, StatusNoU1YAssignmentOrNormalization) {
		t.Fatalf("hypercharge firewall violated: %+v", a.Hypercharge)
	}
	if !a.Twistor.SU2AlgebraSelectorIndependent || !a.Twistor.ComplexCarrierSelectorDependent || !a.Twistor.U1PhaseSelectorDependent || !a.Twistor.SU2CompatibleWithoutSolvingU1 || a.Twistor.PhysicalSelectorSolved {
		t.Fatalf("twistor dependence firewall violated: %+v", a.Twistor)
	}
	if a.Physical.CEqualsPhysicalSU2L || a.Physical.ThetaHPhysical || a.Physical.K7PlusPhysicalHiggsDoublet || a.Physical.EWRepresentationDerivedFromK7 || a.Physical.HyperchargeDerived || a.Physical.HiggsMass || a.Physical.ScalarRuntime || a.Physical.YukawaOperator || a.Physical.YukawaEigenvalues || len(a.Physical.MissingMaps) != 4 {
		t.Fatalf("physical firewall violated: %+v", a.Physical)
	}
	res := Generation2InternalSU2SocketToElectroweakSU2LIntertwinerAirlockAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
