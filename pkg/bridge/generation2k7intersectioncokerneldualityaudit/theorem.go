package generation2k7intersectioncokerneldualityaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2K7IntersectionCokernelDualityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 K7 intersection-cokernel duality audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate629 K7 intersection-cokernel duality audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate628 Lambda4 plus boundary-pair chamber and missing projector", Passed: a.Inherited.Verdict == StatusGate628Inherited && a.Inherited.Lambda4Dimension == 70 && a.Inherited.BoundaryPairDimension == 2 && a.Inherited.AugmentedChamberDimension == 72 && a.Inherited.K7Dimension == 7 && a.Inherited.NonK7Lambda4Complement == 63 && a.Inherited.AugmentedComplementDimension == 65 && a.Inherited.WeightNumerator == 7 && a.Inherited.WeightDenominator == 72 && a.Inherited.Gate628ChamberCandidate && a.Inherited.Gate628ProjectionMissing && a.Inherited.Gate628ProductAirlockMissing && a.Inherited.Gate628FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "compute Boolean-octonionic span and Lambda4 cokernel dimensions", Passed: a.Span.Verdict == StatusSpanDimensionComputed && a.Span.RankPB == 56 && a.Span.RankPG == 14 && a.Span.IntersectionDimension == 7 && a.Span.SpanDimension == 63 && a.Span.SpanMatchesExpected && a.Span.Lambda4Dimension == 70 && a.Span.CokernelDimension == 7 && a.Span.CokernelMatchesK7Dimension && a.Span.SpanDimensionCertifiedByRank, Detail: FormatSpan(a.Span)},
			{Name: "audit intersection-cokernel seven as candidate dual pair only", Passed: a.Duality.Verdict == StatusIntersectionCokernelCandidate && a.Duality.IntersectionDimension == 7 && a.Duality.CokernelDimension == 7 && a.Duality.DimensionsEqual && a.Duality.EqualityIsOnlyDimensional && a.Duality.DualityCandidate && !a.Duality.CanonicalIsomorphismFound && !a.Duality.CanonicalPairingFound && strings.Contains(a.Duality.MissingMap, "Phi"), Detail: FormatDuality(a.Duality)},
			{Name: "audit 72 split as 7 plus 63 plus 2", Passed: a.ChamberSplit.Verdict == Status72SplitAudited && a.ChamberSplit.IntersectionOrGapDimension == 7 && a.ChamberSplit.SpanDimension == 63 && a.ChamberSplit.BoundaryPairDimension == 2 && a.ChamberSplit.AugmentedChamberDimension == 72 && a.ChamberSplit.SplitMatches72 && a.ChamberSplit.NativeSpanDimension && a.ChamberSplit.BoundaryPairBridgeOnly && a.ChamberSplit.SharperThan70Plus2, Detail: FormatChamberSplit(a.ChamberSplit)},
			{Name: "sharpen 65 complement as Boolean-octonionic span plus boundary pair", Passed: a.ComplementRole.Verdict == Status65SpanBoundaryComplement && a.ComplementRole.SpanDimension == 63 && a.ComplementRole.BoundaryPairDimension == 2 && a.ComplementRole.SpanBoundaryComplement == 65 && a.ComplementRole.AugmentedChamberDimension == 72 && a.ComplementRole.Equals65Over72 && math.Abs(a.ComplementRole.ComplementWeight-65.0/72.0) < 1e-15 && !a.ComplementRole.NativeRoleTheoremFound, Detail: FormatComplementRole(a.ComplementRole)},
			{Name: "block boundary-pull assignment to K7, cokernel, or dual pair", Passed: a.BoundaryPullAssignment.Verdict == StatusNoBoundaryPullAssignment && len(a.BoundaryPullAssignment.Candidates) == 3 && a.BoundaryPullAssignment.Candidates[0].CanSupplySeven && a.BoundaryPullAssignment.Candidates[1].CanSupplySeven && a.BoundaryPullAssignment.Candidates[2].CanSupplySeven && !a.BoundaryPullAssignment.IntersectionAssigned && !a.BoundaryPullAssignment.CokernelAssigned && !a.BoundaryPullAssignment.DualPairAssigned && !a.BoundaryPullAssignment.AssignmentCertified, Detail: FormatBoundaryPullAssignment(a.BoundaryPullAssignment)},
			{Name: "reinterpret Gate626 weighted mixture as (63+2)/72 scalar plus 7/72 boundary", Passed: a.WeightedMixture.Verdict == Status65SpanBoundaryComplement && a.WeightedMixture.ScalarWeightAs63Plus2 && a.WeightedMixture.BoundaryWeightAsSeven && math.Abs(a.WeightedMixture.BoundaryWeight-7.0/72.0) < 1e-15 && math.Abs(a.WeightedMixture.ScalarWeight-65.0/72.0) < 1e-15 && math.Abs(a.WeightedMixture.Residual-a.Inherited.WeightedClosureResidual) < 1e-15, Detail: FormatWeightedMixture(a.WeightedMixture)},
			{Name: "record native status and missing Phi/projector theorem", Passed: a.NativeStatus.Lambda4Native && a.NativeStatus.PBImageRankNative && a.NativeStatus.PGImageRankNative && a.NativeStatus.K7IntersectionNative && a.NativeStatus.BooleanOctonionicSpanDimensionTyped && a.NativeStatus.Lambda4CokernelDimensionTyped && !a.NativeStatus.IntersectionCokernelIsomorphism && !a.NativeStatus.BoundaryPairNativeFinite && !a.NativeStatus.BoundaryPullAssignmentNative && !a.NativeStatus.DualBoundaryProjectorNative && !a.NativeStatus.GaugeScalarFlavorTransportNative, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate629 intersection-cokernel firewall", Passed: !a.Firewalls.ClaimsK7CokernelIsomorphism && !a.Firewalls.ClaimsBoundaryPullAssignment && !a.Firewalls.ClaimsDualBoundaryProjector && !a.Firewalls.ClaimsBoundaryPairNative && !a.Firewalls.ClaimsScalarRGMatching && !a.Firewalls.ClaimsFlavorOrientation && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsHiggsMassDerived && !a.Firewalls.ClaimsEndpointDerivation, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Missing duality map: "+a.Duality.MissingMap)
		notes = append(notes, "Missing boundary assignment: "+a.BoundaryPullAssignment.MissingObject)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
