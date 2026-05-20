package generation2k7overlambda4boundarypairprojectionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2K7OverLambda4BoundaryPairProjectionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 K7 over Lambda4 boundary-pair projection audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate628 K7 over Lambda4 boundary-pair projection audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate627 K7 numerator and missing 72/projection theorem", Passed: a.Inherited.Verdict == StatusGate627Inherited && a.Inherited.K7Dimension == 7 && a.Inherited.Lambda4Dimension == 70 && a.Inherited.Gate627Denominator == 72 && a.Inherited.Gate627NumeratorIsNative && !a.Inherited.Gate627Certified72Carrier && !a.Inherited.Gate627ProjectionExists && !a.Inherited.Gate627SourceTheorem && a.Inherited.Gate627FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "identify 72 as 70+2 augmented bridge chamber candidate", Passed: a.Chamber.Verdict == Status72Lambda4BoundaryPair && a.Chamber.Lambda4Dimension == 70 && a.Chamber.BoundaryPairDimension == 2 && a.Chamber.AugmentedChamberDimension == 72 && a.Chamber.EqualsTargetDenominator && a.Chamber.UsesNativeLambda4Carrier && a.Chamber.UsesBridgeBoundaryPair && !a.Chamber.DirectSumCertifiedNative && a.Chamber.BetterThan8Times9, Detail: FormatChamber(a.Chamber)},
			{Name: "rank denominator candidates with 70+2 strongest but non-native", Passed: a.DenominatorComparison.Verdict == Status70Plus2Identified && a.DenominatorComparison.BestIs70Plus2 && a.DenominatorComparison.BestValue == 72 && a.DenominatorComparison.AnyBridgeCandidate && !a.DenominatorComparison.AnyNativeDenominator && len(a.DenominatorComparison.Rows) >= 4 && a.DenominatorComparison.Rows[0].Expression == "70 + 2", Detail: FormatDenominatorComparison(a.DenominatorComparison)},
			{Name: "inherit Gate613/Gate626 boundary stress pair as bridge coordinates", Passed: a.BoundaryPair.Verdict == StatusBoundaryPairInherited && a.BoundaryPair.PairDimension == 2 && a.BoundaryPair.PairIsGate613Boundary && a.BoundaryPair.PairInheritedFromGate626 && !a.BoundaryPair.PairNativeFiniteObject && a.BoundaryPair.BridgeCoordinateOnly && a.BoundaryPair.BoundarySplit > 0, Detail: FormatBoundaryPair(a.BoundaryPair)},
			{Name: "audit K7 as numerator inside Lambda4 R8", Passed: a.K7Embedding.Verdict == StatusK7InsideLambda4Audited && a.K7Embedding.K7Dimension == 7 && a.K7Embedding.Lambda4Dimension == 70 && a.K7Embedding.K7FitsInsideLambda4 && a.K7Embedding.NativeCarrierCertified && !a.K7Embedding.ProjectionToBoundaryFound, Detail: FormatK7Embedding(a.K7Embedding)},
			{Name: "audit 65 as non-K7 Lambda4 complement plus boundary pair", Passed: a.Complement.Verdict == Status65ComplementCandidate && a.Complement.NonK7Lambda4ComplementDimension == 63 && a.Complement.BoundaryPairDimension == 2 && a.Complement.AugmentedComplementDimension == 65 && a.Complement.Equals65Over72 && a.Complement.HasStructuredComplementReading && !a.Complement.NativeComplementProjection, Detail: FormatComplement(a.Complement)},
			{Name: "compute 7/72 as augmented chamber trace fraction without certifying projector", Passed: a.ProjectionTrace.Verdict == StatusProjectionTraceCandidate && a.ProjectionTrace.DomainDimension == 72 && a.ProjectionTrace.K7TraceDimension == 7 && a.ProjectionTrace.TraceFractionMatches && math.Abs(a.ProjectionTrace.TraceFraction-7.0/72.0) < 1e-15 && !a.ProjectionTrace.ProjectionOperatorExists && !a.ProjectionTrace.TraceFunctionalCertified && !a.ProjectionTrace.IntertwinerCertified, Detail: FormatProjectionTrace(a.ProjectionTrace)},
			{Name: "carry Gate626 weighted closure with chamber ratio", Passed: a.WeightedClosure.Verdict == StatusProjectionTraceCandidate && a.WeightedClosure.ChamberRatioMatchesGate626 && math.Abs(a.WeightedClosure.BoundaryWeight-7.0/72.0) < 1e-15 && math.Abs(a.WeightedClosure.WeightedClosureResidual-a.Inherited.Gate626WeightedResidual) < 1e-15, Detail: FormatWeightedClosure(a.WeightedClosure)},
			{Name: "record native status and missing product airlock/projector", Passed: a.NativeStatus.Lambda4Native && a.NativeStatus.K7Native && !a.NativeStatus.BoundaryPairNativeFinite && !a.NativeStatus.AugmentedChamberNative && !a.NativeStatus.ProductAirlockNative && !a.NativeStatus.K7BoundaryPullProjectorNative && !a.NativeStatus.TraceFractionTheoremNative && !a.NativeStatus.GaugeScalarFlavorTransportNative, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate628 augmented-chamber firewalls", Passed: !a.Firewalls.ClaimsNativeAugmentedChamber && !a.Firewalls.ClaimsNativeBoundaryPair && !a.Firewalls.ClaimsNativeProjection && !a.Firewalls.ClaimsNativeTraceTheorem && !a.Firewalls.ClaimsScalarRGMatching && !a.Firewalls.ClaimsFlavorOrientation && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsHiggsMassDerived && !a.Firewalls.ClaimsEndpointDerivation, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Missing projector: "+strings.TrimSpace(FormatProjectionTrace(a.ProjectionTrace)))
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
