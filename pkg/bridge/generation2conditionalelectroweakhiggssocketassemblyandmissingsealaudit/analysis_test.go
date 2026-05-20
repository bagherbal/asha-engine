package generation2conditionalelectroweakhiggssocketassemblyandmissingsealaudit

import (
	"strings"
	"testing"
)

func TestGate719AssemblyAndIntertwiner(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SU2Inherited.SU2SideInherited || !a.SU2Inherited.InternalCCompatibleWithEWHiggs || !a.SU2Inherited.SU2SideStructurallyReady || a.SU2Inherited.CanonicalThetaSU2 || a.SU2Inherited.InternalCPhysicalSU2L || a.SU2Inherited.HyperchargeDerived || a.SU2Inherited.FullTypedHiggsMap || a.SU2Inherited.HiggsMassOrRuntime || a.SU2Inherited.YukawaOperatorOrEigenvalue {
		t.Fatalf("bad SU2 inheritance: %+v", a.SU2Inherited)
	}
	if !a.U1Inherited.U1SideInherited || !a.U1Inherited.PhaseLineCompatibleAfterNAndQ || !a.U1Inherited.FullU2CompatibleOnlyAfterNAndQ || a.U1Inherited.PhaseLineFixesHypercharge || a.U1Inherited.NativeTwistorSelector || a.U1Inherited.NativeThetaYNormalization || a.U1Inherited.FullTypedHiggsMap || a.U1Inherited.HiggsMassOrRuntime || a.U1Inherited.YukawaOperatorOrEigenvalue {
		t.Fatalf("bad U1 inheritance: %+v", a.U1Inherited)
	}
	if !a.Socket.Assembled || a.Socket.Dimension != 4 || a.Socket.ComplexDimension != 2 || !a.Socket.RequiresN || !a.Socket.RequiresQ || a.Socket.PhysicalEWClaimed || !strings.Contains(a.Socket.SocketSymbol, "qJ_H") {
		t.Fatalf("bad socket assembly: %+v", a.Socket)
	}
	if !a.Target.FullLaneIdentified || !a.Target.FiniteSpectralTripleLane || a.Target.TargetComplexDimension != 2 || a.Target.ImportsMassOrYukawaData || !strings.Contains(a.Target.TargetAlgebra, "su(2)_L") || !strings.Contains(a.Target.TargetAlgebra, "u(1)_Y") {
		t.Fatalf("bad target lane: %+v", a.Target)
	}
	if !a.Intertwiner.SU2Compatible || !a.Intertwiner.U1Compatible || !a.Intertwiner.CarrierCompatible || !a.Intertwiner.RequiresN || !a.Intertwiner.RequiresQ || !a.Intertwiner.RepresentationCompatible || a.Intertwiner.PhysicalIdentityClaimed || !strings.Contains(a.Intertwiner.Condition, "Theta_H") {
		t.Fatalf("bad full intertwiner: %+v", a.Intertwiner)
	}
}

func TestGate719ChoicesHyperchargeAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Choices.TwistorPointN || !a.Choices.PhaseNormalizationQ || !a.Choices.SU2BasisIntertwinerChoice || !a.Choices.ComplexBasisChoice || !a.Choices.TargetHyperchargeConvention || a.Choices.CanonicalN || a.Choices.CanonicalQ || a.Choices.CanonicalThetaH {
		t.Fatalf("noncanonical choice audit failed: %+v", a.Choices)
	}
	if !a.Hypercharge.CanMatchTargetYHConvention || !strings.Contains(a.Hypercharge.ExampleTargetConvention, "Y_H=1/2") || a.Hypercharge.QDerivedNatively || a.Hypercharge.HyperchargeDerived || a.Hypercharge.HyperchargeNormalized {
		t.Fatalf("hypercharge firewall failed: %+v", a.Hypercharge)
	}
	if a.Physical.K7PlusPhysicalHiggsDoublet || a.Physical.GIntPhysicalEWAlgebra || a.Physical.QDerivedHypercharge || a.Physical.NDerivedVacuumSelector || a.Physical.ScalarPotential || a.Physical.QuarticRuntimeLambda || a.Physical.HiggsPoleMass || a.Physical.YukawaOperator || a.Physical.FlavorHierarchy || a.Physical.CKMPMNS || len(a.Physical.MissingMaps) != 5 {
		t.Fatalf("physical firewall failed: %+v", a.Physical)
	}
	res := Generation2ConditionalElectroweakHiggsSocketAssemblyAndMissingSealAuditTheorem().Verify()
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
