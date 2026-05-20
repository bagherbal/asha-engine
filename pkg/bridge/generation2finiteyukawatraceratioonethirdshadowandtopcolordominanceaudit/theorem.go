package generation2finiteyukawatraceratioonethirdshadowandtopcolordominanceaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2FiniteYukawaTraceRatioOneThirdShadowAndTopColorDominanceAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 754 — Finite Yukawa Trace Ratio One-Third Shadow and Top-Color Dominance Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate754 finite Yukawa trace-ratio audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate753 scalar proxy coefficient typing", Passed: a.Gate753.Inherited && strings.Contains(a.Gate753.ScalarProxyFormula, "b/a^2") && math.Abs(a.Gate753.GaugeNormalizationCoefficient-threeEighths) < 1e-15 && a.Gate753.BA2OneThirdTheoremBlocked && a.Gate753.ScalarProxyDerivationBlocked && a.Gate753.ProxyToRuntimeMatchingBlocked && a.Gate753.HiggsMassOrPoleMassTheoremBlocked, Detail: FormatGate753(a.Gate753)},
			{Name: "compute sealed b/a^2 trace-ratio input", Passed: a.Inputs.NativeTraceShapes && a.Inputs.SealedYukawaLedger && a.Inputs.Dimensionless && a.Inputs.NonNegative && !a.Inputs.NativeRatioTheorem && math.Abs(a.Inputs.BOverA2Computed-bOverA2MZSeed) < 1e-15 && math.Abs(a.Inputs.RatioSeedResidual) < 1e-15, Detail: FormatInputs(a.Inputs)},
			{Name: "derive exact one-third in top-color dominance limit", Passed: a.TopLimit.SingleDominantChannel && a.TopLimit.ColoredChannel && a.TopLimit.ExactLimitDerived && !a.TopLimit.NativeYukawaTheorem && math.Abs(a.TopLimit.ATopAtUnitY-3) < 1e-15 && math.Abs(a.TopLimit.BTopAtUnitY-3) < 1e-15 && math.Abs(a.TopLimit.RatioAtUnitY-oneThird) < 1e-15 && strings.Contains(a.TopLimit.RatioFormula, "1/3"), Detail: FormatTopLimit(a.TopLimit)},
			{Name: "compute one-third deviation and block source assignment", Passed: a.Deviation.MeasuresNonTopDominanceCorrection && !a.Deviation.DeltaSourceAssigned && !a.Deviation.NativeDeltaDecomposition && len(a.Deviation.ExpectedCorrectionCandidates) == 5 && math.Abs(a.Deviation.DeltaRatio+0.0002583937062663466) < 1e-15 && math.Abs(a.Deviation.RelativeToOneThird+0.0007751811187990398) < 1e-15, Detail: FormatDeviation(a.Deviation)},
			{Name: "compute one-eighth scalar proxy shadow", Passed: a.OneEighth.CloseToOneEighth && !a.OneEighth.ScalarPotentialTheorem && math.Abs(a.OneEighth.IdealProxy-oneEighth) < 1e-15 && math.Abs(a.OneEighth.ActualProxyComputed-threeEighths*a.Inputs.BOverA2Computed) < 1e-15 && math.Abs(a.OneEighth.ActualProxyComputed-lambdaProxyMZ) < 1e-13 && math.Abs(a.OneEighth.ProxyMinusOneEighth-a.OneEighth.CoefficientTimesTraceDeviation) < 1e-16 && math.Abs(a.OneEighth.ShadowIdentityResidual) < 1e-16, Detail: FormatOneEighth(a.OneEighth)},
			{Name: "enforce source-layer firewall", Passed: a.SourceLayers.AllLayersSeparated && !a.SourceLayers.ClaimsThreeEighthsScalarLaw && !a.SourceLayers.ClaimsOneThirdYukawaLaw && !a.SourceLayers.ClaimsOneEighthScalarLaw && strings.Contains(a.SourceLayers.ThreeEighthsLayer, "gauge") && strings.Contains(a.SourceLayers.OneThirdLayer, "top-color") && strings.Contains(a.SourceLayers.OneEighthLayer, "proxy"), Detail: FormatSourceLayers(a.SourceLayers)},
			{Name: "enforce Yukawa theorem firewall", Passed: !a.Yukawa.ClaimsYuDerived && !a.Yukawa.ClaimsYdDerived && !a.Yukawa.ClaimsYeDerived && !a.Yukawa.ClaimsYnuDerived && !a.Yukawa.ClaimsTopYukawaDerived && !a.Yukawa.ClaimsYukawaHierarchyDerived && !a.Yukawa.ClaimsCKMPMNSDerived && !a.Yukawa.ClaimsGenerationCarrier && a.Yukawa.SealedLedgerExplicit, Detail: FormatYukawaFirewall(a.Yukawa)},
			{Name: "enforce runtime and mass firewall", Passed: !a.Runtime.ClaimsRuntimeLambdaTheorem && !a.Runtime.ClaimsIndependentScalarRuntime && !a.Runtime.ClaimsProxyToRuntimeMatching && !a.Runtime.ClaimsHiggsMassTheorem && !a.Runtime.ClaimsPoleMassTheorem && a.Runtime.RequiresHistoryLoopTransport && a.Runtime.RequiresBoundaryHistoryResponse && a.Runtime.RequiresKappaEReduction && a.Runtime.RequiresScalarRuntimeBridge, Detail: FormatRuntimeFirewall(a.Runtime)},
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
