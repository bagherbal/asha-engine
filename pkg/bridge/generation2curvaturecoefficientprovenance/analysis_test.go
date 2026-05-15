package generation2curvaturecoefficientprovenance

import (
	"strings"
	"testing"
)

func TestGate510Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate509GravitySocketInherited || !a.Inheritance.Gate509NormalizationBlocked || !a.Inheritance.HeatKernelConventionDeclared {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Endomorphism.CurvatureEndomorphismAudited || !nearly(a.Endomorphism.CombinedA2RPart, -1.0/12.0, 1e-15) || a.Endomorphism.SignConventionClosed || a.Endomorphism.PhysicalMetricDynamicsDerived {
		t.Fatalf("bad endomorphism audit: %+v", a.Endomorphism)
	}
	if !a.A2.Gate377RawCoefficientMatched || !a.A2.DimensionlessTraceWeightNative || a.A2.IncludesCutoffMoment || a.A2.PhysicalCoefficientNative {
		t.Fatalf("bad a2 audit: %+v", a.A2)
	}
	if !nearly(a.A2.A2WeightMagnitudeBefore4Pi, 8, 1e-15) {
		t.Fatalf("bad finite trace weight %.18g", a.A2.A2WeightMagnitudeBefore4Pi)
	}
	if a.Convention.UniqueTraceConventionSelected || a.Convention.CanPromoteEitherToNewtonNative {
		t.Fatalf("convention overpromoted: %+v", a.Convention)
	}
	if !a.Cutoff.RequiresF2LambdaSquaredProduct || a.Cutoff.F2MomentSeparatedFromLambda || a.Cutoff.CutoffLambdaSelected || a.Cutoff.NewtonConstantDerived || !a.Cutoff.GravityNormalizationBridgeOnly {
		t.Fatalf("cutoff firewall failed: %+v", a.Cutoff)
	}
	if a.Firewall.NewtonConstantImported || a.Firewall.PlanckMassImported || a.Firewall.NativeGravityNormalizationWritten || a.Firewall.ElectroweakScaleImported || a.Firewall.FlavorDataImported {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 511 {
		t.Fatalf("unexpected next gate %+v", a.Next)
	}
}

func TestGate510Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{
		"# Gate 510 Registry Audit",
		StatusLichnerowiczCurvatureTermAudited,
		StatusA2TraceWeightComputed,
		StatusFailedNewtonConstantStillNotDerived,
		StatusFirewallNativeGravityNormalizationStop,
		"Tr_F(1)/12 = 96/12 = 8",
		"Gate 511",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate510Theorem(t *testing.T) {
	result := Generation2CurvatureCoefficientProvenanceHeatKernelTraceConventionAuditTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
