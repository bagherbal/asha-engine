package generation2historyloophessianradialprojectoralignmentanddualrolefirewallaudit

import (
	"math"
	"strings"
	"testing"
)

func closeTo(x, y float64) bool { return math.Abs(x-y) <= 1e-12 }

func TestGate767InheritsDualRadialRoles(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate766.Inherited || !strings.Contains(a.Gate766.HessianFormula, "P_rad") || !strings.Contains(a.Gate766.TreeProxyFormula, "lambda_runtime_eff") || !a.Gate766.HistoryLoopRoleSeparated || a.Gate766.NativeHistoryLoopHessianTheorem || a.Gate766.NativePotentialTheorem || a.Gate766.TreeProxyPoleMassTheorem {
		t.Fatalf("bad Gate766 inheritance: %+v", a.Gate766)
	}
	if a.HistoryLoop.Carrier != "K7+ ~= R^4" || a.HistoryLoop.ProjectorRank != radialRank || !closeTo(a.HistoryLoop.TraceWeight, 0.25) || !closeTo(a.HistoryLoop.LHopf, 1/(8*math.Pi)) || !a.HistoryLoop.DependsOnRankOnly || a.HistoryLoop.IdentifiesHessianSupport {
		t.Fatalf("bad HistoryLoop role: %+v", a.HistoryLoop)
	}
	if !strings.Contains(a.Hessian.Potential, "lambda/4") || a.Hessian.Projector != "P_hessian=u_rad u_rad^T" || !strings.Contains(a.Hessian.HessianFormula, "2 lambda v^2") || a.Hessian.RadialEigenvalueFormula != "2 lambda v^2" || !closeTo(a.Hessian.RadialEigenvalueGeV2, 2*lambdaRuntimeEff*vevGate741GeV*vevGate741GeV) || len(a.Hessian.AngularEigenvalues) != angularDim || a.Hessian.HessianRank != radialRank || a.Hessian.NativePotentialTheorem {
		t.Fatalf("bad Hessian role: %+v", a.Hessian)
	}
}

func TestGate767CompatibilityAlignmentAndCollision(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Compatibility.SameAmbientCarrier != "K7+ ~= R^4" || a.Compatibility.HistoryLoopProjectorRank != radialRank || a.Compatibility.HessianProjectorRank != radialRank || !a.Compatibility.BothOrthogonalProjectors || !a.Compatibility.BothRealRankOne || !a.Compatibility.SpaceRankCompatible || a.Compatibility.CompatibilitySufficient {
		t.Fatalf("bad compatibility audit: %+v", a.Compatibility)
	}
	if a.Alignment.SealName != "HistoryLoopHessianRadialAlignmentSeal" || a.Alignment.Identification != "P_history = P_hessian = P_rad" || !a.Alignment.LawfulReuse || a.Alignment.NativeAlignmentTheorem || !a.Alignment.BridgeConditional {
		t.Fatalf("bad alignment seal: %+v", a.Alignment)
	}
	if !a.Collision.CollisionIfUnaligned || !a.Collision.RejectedUnderPremises || !strings.Contains(a.Collision.Reason, "same supplied unit radial vector") || !a.Collision.RequiresExplicitSeal || a.Collision.NativeSemanticIdentity {
		t.Fatalf("bad collision audit: %+v", a.Collision)
	}
}

func TestGate767RankInvarianceLimitationAndPipeline(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.RankLimit.TraceFormula, "rank(P)/4") || !a.RankLimit.AnyRankOneGivesSameTrace || a.RankLimit.NumericalLHopfProvesAlignment || !a.RankLimit.HessianLaneNeededForTypedSupport || !strings.Contains(a.RankLimit.Limitation, "does not select") {
		t.Fatalf("bad rank-invariance limitation: %+v", a.RankLimit)
	}
	if !strings.Contains(a.Pipeline.HistoryLoopLane, "L_Hopf") || !strings.Contains(a.Pipeline.HessianLane, "H_V") || !strings.Contains(a.Pipeline.SharedProjectorRole, "bridge-aligned") || !strings.Contains(a.Pipeline.ScalarBridgeRole, "lambda_runtime_eff") || !strings.Contains(a.Pipeline.TreeProxyRole, "m_H_tree_proxy") || !a.Pipeline.PipelineCoherent || a.Pipeline.NativePipelineTheorem {
		t.Fatalf("bad dual-role pipeline: %+v", a.Pipeline)
	}
}

func TestGate767FirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Audited || a.Firewalls.NativeHistoryLoopHessianAlignment || a.Firewalls.RankTraceIdentifiesHessianPRad || a.Firewalls.NativePotentialTheorem || a.Firewalls.NativeHistoryLoopUnitTheorem || a.Firewalls.IndependentScalarRuntimeTheorem || a.Firewalls.TreeProxyPoleMassTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2HistoryLoopHessianRadialProjectorAlignmentAndDualRoleFirewallAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
