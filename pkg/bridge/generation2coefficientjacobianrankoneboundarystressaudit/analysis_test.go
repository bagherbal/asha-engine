package generation2coefficientjacobianrankoneboundarystressaudit

import (
	"math"
	"testing"
)

func TestGate616CoefficientJacobian(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Inherited.XiBoundary-0.0503471644870914) > 1e-12 {
		t.Fatalf("unexpected xi %.15g", a.Inherited.XiBoundary)
	}
	if len(a.DependencyGraph) < 8 || !hasDependency(a.DependencyGraph, "C_3") || !hasDependency(a.DependencyGraph, "lambda") || !hasDependency(a.DependencyGraph, "q_boundary") {
		t.Fatalf("bad dependency graph: %+v", a.DependencyGraph)
	}
	if a.ShadowMap.RawPairTypeSafe || !a.ShadowMap.PreferredTypeSafe || a.ShadowMap.ColorShadow <= 0 || a.ShadowMap.ScalarShadow >= 0 {
		t.Fatalf("bad shadow map: %+v", a.ShadowMap)
	}
	if !hasJacobian(a.Jacobian, "C_3", "+", "0") || !hasJacobian(a.Jacobian, "lambda", "0", "+") || !hasJacobian(a.Jacobian, "q_boundary stress", "+", "-") {
		t.Fatalf("bad Jacobian: %+v", a.Jacobian)
	}
}

func TestGate616RankAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !hasBridgeRankOne(a.RankOneCandidates) {
		t.Fatalf("expected bridge q_stress rank-one candidate: %+v", a.RankOneCandidates)
	}
	if hasNativeRankOne(a.RankOneCandidates) || a.RankClassification.NativeRankOneFound {
		t.Fatalf("native rank-one source should not be found: %+v", a.RankClassification)
	}
	if !a.RankClassification.BridgeRankOneDefinable || !a.RankClassification.RankTwoIndependentSlots {
		t.Fatalf("rank classification should be B/C hybrid: %+v", a.RankClassification)
	}
	if a.AntiAlignment.CanForceAntiAlignment || a.AntiAlignment.Native || a.AntiAlignment.ResidualOverXi > 0.03 {
		t.Fatalf("bad anti-alignment audit: %+v", a.AntiAlignment)
	}
	if a.CanonicalNormalization.KPhiKnown || a.CanonicalNormalization.CanonicalScalarLedgerKnown || a.CanonicalNormalization.CanAuditLambdaBeforeAfterK {
		t.Fatalf("canonical normalization should remain incomplete: %+v", a.CanonicalNormalization)
	}
	if a.NativeStatus.SectorSplitF0 || a.NativeStatus.NativeQStress || a.NativeStatus.C3LambdaRelation || a.NativeStatus.ScalarNormalization || a.NativeStatus.ThresholdMatching || a.NativeStatus.NativeXi {
		t.Fatalf("native status violated: %+v", a.NativeStatus)
	}
	if a.Firewalls.ClaimsXiNative || a.Firewalls.ClaimsLambdaZero || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsHiggsStability || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsThresholdExistence || a.Firewalls.ClaimsNativeCorrection {
		t.Fatalf("firewall violated: %+v", a.Firewalls)
	}
}
