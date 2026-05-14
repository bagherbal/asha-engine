package scalarchernweiltaudit

import (
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarorientationseal"
	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func SealedScalarBundleChernWeilCarrierHeatKernelPreflightTheorem() theorem.Theorem {
	const id = "BRIDGE-SEALED-SCALAR-BUNDLE-CHERN-WEIL-CARRIER-HEAT-KERNEL-PREFLIGHT-AUDIT"
	const name = "sealed scalar-bundle Chern-Weil carrier / heat-kernel preflight audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{{Name: "build sealed scalar-bundle Chern-Weil preflight", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 191 sealed scalar bundle is inherited explicitly", Passed: a.Summary.InheritedSealedBundle && a.Seal.Seal.ExplicitAxiom && a.Seal.Seal.Quarantined && a.Seal.Firewall.ConditionalPhysicalBundleDerived, Detail: scalarorientationSealDetail(a)},
			{Name: "finite curvature square traces are stable rational representation traces", Passed: a.Curvature.PrimitiveGaugeKineticTracesStable && a.Curvature.T1TraceStable && a.Curvature.T2TraceStable && a.Curvature.T3TraceStable && a.Curvature.YPhiTraceStable && a.Curvature.NeutralQTraceStable && a.Curvature.NeutralZTraceStable && a.Curvature.HighLowFiberTraceAvailable && a.Curvature.GaugeKineticCarrierPreflight && !a.Curvature.PhysicalGaugeCouplingsDerived, Detail: FormatCurvature(a.Curvature)},
			{Name: "eta seal is a valid grading and only the neutral split gives a nontrivial signed carrier", Passed: a.Grading.EtaDerivedFromSeal && a.Grading.EtaSquaredResidual == 0 && a.Grading.EtaTrace == 0 && a.Grading.PrimitiveDiagonalGradedTracesZero && a.Grading.ChargedGradedTracesZero && a.Grading.NontrivialSignedNeutralCarrier && !a.Grading.IntegerTopologicalChargeMapDerived && !a.Grading.ContinuumOrientationDerived, Detail: FormatGrading(a.Grading)},
			{Name: "heat-kernel a4 support is preflight only", Passed: a.HeatKernel.FiniteMatrixTraceAvailable && a.HeatKernel.SealedScalarBundleDimension == 4 && a.HeatKernel.GaugeFluctuationSquareTraceNonzero && a.HeatKernel.A4LocalAlgebraicIngredientPresent && !a.HeatKernel.DiracOperatorDerived && !a.HeatKernel.OrderOneAxiomVerified && !a.HeatKernel.DixmierTraceDerived && !a.HeatKernel.SpectralActionEvaluated && !a.HeatKernel.HeatKernelCoefficientPromoted, Detail: FormatHeatKernel(a.HeatKernel)},
			{Name: "topological-coupling firewall forbids 8pi2 import, instanton matching, thresholds, couplings, and constants", Passed: a.Firewall.UsesSpontaneousOrientationSeal && !a.Firewall.ImportsTopologicalSeal8PiSquared && !a.Firewall.EquatesFiniteTraceWithInstanton && a.Firewall.ChernWeilCarrierPreflight && !a.Firewall.CompleteChernWeilCarrierDerived && a.Firewall.HeatKernelPreflightPassed && !a.Firewall.HeatKernelMatchingDerived && !a.Firewall.ThresholdBetaRowsDerived && !a.Firewall.AbsoluteCouplingPromoted && !a.Firewall.PhysicalConstantsDerived && !a.Firewall.PhysicalGaugeCouplingsDerived && !a.Firewall.PhysicalMassesDerived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records a local carrier preflight, not a continuum normalization theorem", Passed: a.Summary.TestsAudited == 4 && a.Summary.FiniteCurvatureTracesStable && a.Summary.EtaGradingValid && a.Summary.NontrivialSignedNeutralCarrier && a.Summary.HeatKernelPreflightPassed && a.Summary.ChernWeilOnlyPreflight && a.Summary.CouplingsAndThresholdsStillSealed, Detail: FormatSummary(a.Summary)},
		}, Notes: []string{
			a.TruthStatement,
			"The nonzero signed trace is finite local carrier data; it is not an integer instanton charge and is not equated with S_top=8π².",
			"The next missing object is an integration/fundamental-class bridge, not another eta-orientation selector.",
		}}
	}}
}

func scalarorientationSealDetail(a Analysis) string {
	return "seal=" + scalarorientationseal.FormatSeal(a.Seal.Seal) + "; frame=" + scalarorientationseal.FormatSealedFrame(a.Seal.SealedFrame)
}
