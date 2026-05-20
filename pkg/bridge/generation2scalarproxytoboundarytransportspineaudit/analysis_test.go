package generation2scalarproxytoboundarytransportspineaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate658Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.TransportPivotInherited || !a.Inherited.FanoBoundaryClosed || !a.Inherited.ActiveBridgeVectorBuilt || !a.Inherited.PrimaryWasRGTransport || !a.Inherited.ScalarMatchingActive || !a.Inherited.BoundaryStressActive || !a.Inherited.HistoryLoopActive || !a.Inherited.K7BoundaryBlocked || !a.Inherited.NoFanoBoundaryMap || !a.Inherited.NoSevenTraceTheorem || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Proxy.CloseToOneEighth || !a.Proxy.TreeProxyOnly || !a.Proxy.CannotReplaceRuntime || math.Abs(a.Proxy.LambdaProxyMZ-lambdaProxyMZ) > 1e-15 || math.Abs(a.Proxy.BA2Ratio-0.33307493962706664) > 1e-15 {
		t.Fatalf("bad proxy lane: %+v", a.Proxy)
	}
	if !a.Matching.LoopSized || math.Abs(a.Matching.RelativeMatch-rhoLambdaMatch) > 1e-15 || math.Abs(a.Matching.KappaLambda-0.0443230430960771) > 5e-13 || math.Abs(a.Matching.ReconstructionResidual) > 2e-13 || a.Matching.RawLAnsatzResidual <= 0 {
		t.Fatalf("bad matching lane: %+v", a.Matching)
	}
	if !a.RG.RuntimeTurnsNegative || !a.RG.UsesCurrentV1RG || a.RG.ClaimsThresholdLaw || math.Abs(a.RG.AbsLambdaBoundary-0.0497009420776833) > 1e-15 {
		t.Fatalf("bad RG lane: %+v", a.RG)
	}
	if math.Abs(a.Boundary.MeanStressRecomputed-a.Boundary.XiBoundary) > 1e-15 || a.Boundary.BoundarySplit <= 0 || math.Abs(a.Boundary.AbsLambdaResidualToXi+a.Boundary.R3ResidualToXi) > 1e-15 || !a.Boundary.XiPreferredOverHalfTrace {
		t.Fatalf("bad boundary lane: %+v", a.Boundary)
	}
	if len(a.Residuals.Slots) != 10 || !a.Residuals.MatchSlotSeparated || !a.Residuals.RGSlotSeparated || !a.Residuals.BoundarySlotSeparated || !a.Residuals.ThresholdSlotsOpen {
		t.Fatalf("bad residual slots: %+v", a.Residuals)
	}
	if a.Sources.KappaLambdaSourceCertified || a.Sources.XiBoundarySourceCertified || a.Sources.HistoryLoopSourceCertified || a.Sources.ProxyRuntimeTheorem || a.Sources.RGThresholdTheorem || a.Sources.BoundaryStressTheorem || a.Sources.SearchedRandomConstants || !a.Sources.TypedQuantitiesOnly {
		t.Fatalf("bad source audit: %+v", a.Sources)
	}
	if !a.Spine.Active || !a.Spine.BridgeLayerOnly || !a.Spine.MergesScalarBoundary || len(a.Spine.Touches) != 7 || !strings.Contains(a.Spine.NextPressurePoint, "low-scale loop matching") {
		t.Fatalf("bad spine classification: %+v", a.Spine)
	}
	if a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsScalarStability || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsThresholdExistence || a.Firewalls.ClaimsNativeScalarTheorem || a.Firewalls.ClaimsBoundaryStressDerived || a.Firewalls.ClaimsPhysicalSpacetime || a.Firewalls.ClaimsFlavorTheorem || a.Firewalls.Verdict != StatusGate658Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2ScalarProxyToBoundaryTransportSpineAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate657TransportPivotInherited, StatusScalarProxyRuntimeChainConstructed, StatusHistoryLoopMatchingFormComputed, StatusKappaLambdaDefined, StatusRGTransportLaneRecorded, StatusBoundaryStressComparisonInherited, StatusResidualSlotsSeparated, StatusSourceAuditComputed, StatusScalarProxyBoundarySpineActive, StatusLowScaleLoopMatchingClueActive, StatusBoundaryStressTransportLive, StatusNoNativeProxyRuntimeTheorem, StatusNoNativeRGThresholdTheorem, StatusNoNativeBoundaryStressTheorem, StatusNoNativeKappaLambdaSource, StatusNoNativeHistoryLoopSourceFromScalar, StatusNoHiggsMassOrStabilityClaim, StatusGate658Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
