package generation2historyloophessianradialprojectoralignmentanddualrolefirewallaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HistoryLoopHessianRadialProjectorAlignmentAndDualRoleFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 767 — HistoryLoop-Hessian Radial Projector Alignment and Dual-Role Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate767HistoryLoopHessianAlignmentBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 766 Hessian/tree-proxy separation", Passed: a.Gate766.Inherited && strings.Contains(a.Gate766.HessianFormula, "P_rad") && strings.Contains(a.Gate766.TreeProxyFormula, "lambda_runtime_eff") && a.Gate766.HistoryLoopRoleSeparated && !a.Gate766.NativeHistoryLoopHessianTheorem && !a.Gate766.NativePotentialTheorem && !a.Gate766.TreeProxyPoleMassTheorem, Detail: FormatGate766(a.Gate766)},
			{Name: "record HistoryLoop radial trace role", Passed: a.HistoryLoop.Carrier == "K7+ ~= R^4" && a.HistoryLoop.State == "rho_plus=I_K7+/4" && a.HistoryLoop.ProjectorRank == radialRank && near(a.HistoryLoop.TraceWeight, 0.25, 1e-15) && near(a.HistoryLoop.LHopf, 1/(8*math.Pi), 1e-15) && a.HistoryLoop.DependsOnRankOnly && !a.HistoryLoop.IdentifiesHessianSupport, Detail: FormatHistoryLoop(a.HistoryLoop)},
			{Name: "record Hessian radial support role", Passed: strings.Contains(a.Hessian.Carrier, "R^4") && strings.Contains(a.Hessian.Potential, "lambda/4") && a.Hessian.Projector == "P_hessian=u_rad u_rad^T" && strings.Contains(a.Hessian.HessianFormula, "P_hessian") && a.Hessian.RadialEigenvalueFormula == "2 lambda v^2" && near(a.Hessian.RadialEigenvalueGeV2, 2*lambdaRuntimeEff*vevGate741GeV*vevGate741GeV, 1e-9) && len(a.Hessian.AngularEigenvalues) == angularDim && a.Hessian.HessianRank == radialRank && !a.Hessian.NativePotentialTheorem, Detail: FormatHessian(a.Hessian)},
			{Name: "audit projector space and rank compatibility", Passed: a.Compatibility.SameAmbientCarrier == "K7+ ~= R^4" && a.Compatibility.HistoryLoopProjectorRank == radialRank && a.Compatibility.HessianProjectorRank == radialRank && a.Compatibility.BothOrthogonalProjectors && a.Compatibility.BothRealRankOne && a.Compatibility.SpaceRankCompatible && !a.Compatibility.CompatibilitySufficient, Detail: FormatCompatibility(a.Compatibility)},
			{Name: "define conditional alignment seal", Passed: a.Alignment.SealName == "HistoryLoopHessianRadialAlignmentSeal" && strings.Contains(a.Alignment.Premise, "HistoryLoop radial event") && a.Alignment.Identification == "P_history = P_hessian = P_rad" && a.Alignment.LawfulReuse && !a.Alignment.NativeAlignmentTheorem && a.Alignment.BridgeConditional && strings.Contains(a.Alignment.Verdict, StatusSamePRadLawfulAfterAlignment), Detail: FormatAlignment(a.Alignment)},
			{Name: "reject notation collision only under explicit alignment premises", Passed: a.Collision.CollisionIfUnaligned && a.Collision.RejectedUnderPremises && strings.Contains(a.Collision.Reason, "same supplied unit radial vector") && a.Collision.RequiresExplicitSeal && !a.Collision.NativeSemanticIdentity, Detail: FormatCollision(a.Collision)},
			{Name: "audit rank-invariance limitation", Passed: strings.Contains(a.RankLimit.TraceFormula, "rank(P)/4") && a.RankLimit.AnyRankOneGivesSameTrace && !a.RankLimit.NumericalLHopfProvesAlignment && a.RankLimit.HessianLaneNeededForTypedSupport && strings.Contains(a.RankLimit.Limitation, "does not select") && strings.Contains(a.RankLimit.Verdict, StatusRankInvarianceDoesNotIdentifyHessianPRad), Detail: FormatRankLimit(a.RankLimit)},
			{Name: "record dual-role scalar pipeline", Passed: strings.Contains(a.Pipeline.HistoryLoopLane, "L_Hopf") && strings.Contains(a.Pipeline.HessianLane, "H_V") && strings.Contains(a.Pipeline.SharedProjectorRole, "bridge-aligned") && strings.Contains(a.Pipeline.ScalarBridgeRole, "lambda_runtime_eff") && strings.Contains(a.Pipeline.TreeProxyRole, "m_H_tree_proxy") && a.Pipeline.PipelineCoherent && !a.Pipeline.NativePipelineTheorem, Detail: FormatPipeline(a.Pipeline)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.NativeHistoryLoopHessianAlignment && !a.Firewalls.RankTraceIdentifiesHessianPRad && !a.Firewalls.NativePotentialTheorem && !a.Firewalls.NativeHistoryLoopUnitTheorem && !a.Firewalls.IndependentScalarRuntimeTheorem && !a.Firewalls.TreeProxyPoleMassTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewalls.Verdict, StatusGate767HistoryLoopHessianAlignmentBoundary), Detail: FormatFirewalls(a.Firewalls)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func near(x, y, tol float64) bool { return math.Abs(x-y) <= tol }
