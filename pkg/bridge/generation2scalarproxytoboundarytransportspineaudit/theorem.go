package generation2scalarproxytoboundarytransportspineaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ScalarProxyToBoundaryTransportSpineAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 658 — Scalar Proxy-to-Boundary Transport Spine Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate658 scalar transport spine audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate657 transport pivot", Passed: a.Inherited.TransportPivotInherited && a.Inherited.FanoBoundaryClosed && a.Inherited.ActiveBridgeVectorBuilt && a.Inherited.PrimaryWasRGTransport && a.Inherited.ScalarMatchingActive && a.Inherited.BoundaryStressActive && a.Inherited.HistoryLoopActive && a.Inherited.K7BoundaryBlocked && a.Inherited.NoFanoBoundaryMap && a.Inherited.NoSevenTraceTheorem && a.Inherited.FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "construct scalar proxy lane", Passed: a.Proxy.CloseToOneEighth && a.Proxy.TreeProxyOnly && a.Proxy.CannotReplaceRuntime && math.Abs(a.Proxy.LambdaProxyMZ-lambdaProxyMZ) < 1e-15 && math.Abs(a.Proxy.BA2Ratio-0.33307493962706664) < 1e-15, Detail: FormatProxy(a.Proxy)},
			{Name: "compute low-scale HistoryLoopUnit matching form", Passed: a.Matching.LoopSized && math.Abs(a.Matching.RelativeMatch-rhoLambdaMatch) < 1e-15 && math.Abs(a.Matching.HistoryLoopUnit-historyLoopL) < 1e-15 && math.Abs(a.Matching.KappaLambda-0.0443230430960771) < 5e-13 && math.Abs(a.Matching.ReconstructionResidual) < 2e-13 && a.Matching.RawLAnsatzResidual > 0, Detail: FormatMatching(a.Matching)},
			{Name: "record RG transport lane to Lambda_12", Passed: a.RG.StartScale == "M_Z" && a.RG.BoundaryScale == "Lambda_12" && a.RG.RuntimeTurnsNegative && a.RG.UsesCurrentV1RG && !a.RG.ClaimsThresholdLaw && math.Abs(a.RG.AbsLambdaBoundary-0.0497009420776833) < 1e-15, Detail: FormatRG(a.RG)},
			{Name: "inherit boundary stress comparison", Passed: math.Abs(a.Boundary.MeanStressRecomputed-a.Boundary.XiBoundary) < 1e-15 && a.Boundary.BoundarySplit > 0 && math.Abs(a.Boundary.AbsLambdaResidualToXi+a.Boundary.R3ResidualToXi) < 1e-15 && a.Boundary.XiPreferredOverHalfTrace, Detail: FormatBoundary(a.Boundary)},
			{Name: "separate residual and correction slots", Passed: len(a.Residuals.Slots) == 10 && a.Residuals.MatchSlotSeparated && a.Residuals.RGSlotSeparated && a.Residuals.BoundarySlotSeparated && a.Residuals.ThresholdSlotsOpen, Detail: FormatResiduals(a.Residuals)},
			{Name: "audit missing sources without random constants", Passed: !a.Sources.KappaLambdaSourceCertified && !a.Sources.XiBoundarySourceCertified && !a.Sources.HistoryLoopSourceCertified && !a.Sources.ProxyRuntimeTheorem && !a.Sources.RGThresholdTheorem && !a.Sources.BoundaryStressTheorem && !a.Sources.SearchedRandomConstants && a.Sources.TypedQuantitiesOnly, Detail: FormatSources(a.Sources)},
			{Name: "classify scalar proxy-to-boundary spine as active bridge", Passed: a.Spine.Active && a.Spine.BridgeLayerOnly && a.Spine.MergesScalarBoundary && len(a.Spine.Touches) == 7 && strings.Contains(a.Spine.NextPressurePoint, "low-scale loop matching"), Detail: FormatSpine(a.Spine)},
			{Name: "preserve scalar transport firewalls", Passed: !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsScalarStability && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsThresholdExistence && !a.Firewalls.ClaimsNativeScalarTheorem && !a.Firewalls.ClaimsBoundaryStressDerived && !a.Firewalls.ClaimsPhysicalSpacetime && !a.Firewalls.ClaimsFlavorTheorem && a.Firewalls.Verdict == StatusGate658Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
