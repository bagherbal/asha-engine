package generation2k7hodgepolarityprojectiveselectoralignmentaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2K7HodgePolarityProjectiveSelectorAlignmentAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 635 — K7 Hodge Polarity and Projective Selector Alignment Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate635 K7 Hodge polarity/projective selector alignment audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate634 K7 Hodge signature", Passed: a.Inherited.Verdict == StatusGate634Inherited && a.Inherited.K7Dimension == 7 && a.Inherited.PlusDimension == 4 && a.Inherited.MinusDimension == 3 && math.Abs(a.Inherited.Trace-1) < 1e-10 && math.Abs(a.Inherited.Determinant+1) < 1e-8 && a.Inherited.HodgeStable && a.Inherited.MixedHodgePolarity && a.Inherited.NoBoundaryAssignment && a.Inherited.NoSevenOver72Theorem && a.Inherited.Gate634FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "define K7 plus/minus Hodge subspaces", Passed: a.K7Subspaces.Verdict == StatusK7PlusMinusSubspacesDefined && a.K7Subspaces.PlusDimension == 4 && a.K7Subspaces.MinusDimension == 3 && a.K7Subspaces.SumDimension == 7 && a.K7Subspaces.ProjectorsCertified && a.K7Subspaces.PlusMinusOrthogonal && a.K7Subspaces.NativeHodgePolarity, Detail: FormatK7Subspaces(a.K7Subspaces)},
			{Name: "inherit projective B-L selector 1+3 reference", Passed: a.ProjectiveSelector.Verdict == StatusProjectiveSelectorInherited && a.ProjectiveSelector.CarrierComplexDimension == 4 && a.ProjectiveSelector.LineComplexDimension == 1 && a.ProjectiveSelector.SpatialBlockComplexDimension == 3 && a.ProjectiveSelector.CP0CP2CriticalStrata && a.ProjectiveSelector.ProjectiveOnePlusThree && a.ProjectiveSelector.MatchesGate555Commutant && !a.ProjectiveSelector.CP3ToK7FunctorFound, Detail: FormatProjectiveSelector(a.ProjectiveSelector)},
			{Name: "audit 4|3 polarity against 4=1+3 selector pattern", Passed: strings.Contains(a.Alignment.Verdict, StatusFourPlusThreeAudited) && strings.Contains(a.Alignment.Verdict, StatusResemblesSelectorSplit) && strings.Contains(a.Alignment.Verdict, StatusNoK7ToFockMap) && a.Alignment.FourDimensionalMatch && a.Alignment.ThreeDimensionalMatch && !a.Alignment.SameCarrier && !a.Alignment.TypedThetaMapFound && !a.Alignment.K7ToCP3FunctorFound && a.Alignment.AlignmentCandidateOnly, Detail: FormatAlignment(a.Alignment)},
			{Name: "block native 4=1+3 refinement inside K7 plus", Passed: a.K7PlusRefinement.Verdict == StatusNoK7PlusOnePlusThree && a.K7PlusRefinement.K7PlusDimension == 4 && a.K7PlusRefinement.HodgeProjectorActsAsIdentity && !a.K7PlusRefinement.InternalRankOneLineDerived && !a.K7PlusRefinement.InternalThreePlaneDerived && !a.K7PlusRefinement.NativeOnePlusThreeRefinement, Detail: FormatK7PlusRefinement(a.K7PlusRefinement)},
			{Name: "keep K7 minus triplet as dimension-only resemblance", Passed: strings.Contains(a.K7MinusTriplet.Verdict, StatusResemblesSelectorSplit) && strings.Contains(a.K7MinusTriplet.Verdict, StatusNoK7ToFockMap) && a.K7MinusTriplet.K7MinusDimension == 3 && a.K7MinusTriplet.SelectorSpatialBlockDim == 3 && a.K7MinusTriplet.DimensionMatch && !a.K7MinusTriplet.TypedTripletIdentification && !a.K7MinusTriplet.UsesBMinusLCarrier, Detail: FormatK7MinusTriplet(a.K7MinusTriplet)},
			{Name: "classify trace plus one as imbalance not distinguished line", Passed: a.TraceImbalance.Verdict == StatusTraceNotDistinguishedLine && math.Abs(a.TraceImbalance.Trace-1) < 1e-10 && a.TraceImbalance.PlusMinusDifference == 1 && math.Abs(a.TraceImbalance.Determinant+1) < 1e-8 && !a.TraceImbalance.DistinguishedLineDerived && !a.TraceImbalance.TraceAsRankOneProjector && a.TraceImbalance.NeedsAdditionalSelector, Detail: FormatTraceImbalance(a.TraceImbalance)},
			{Name: "preserve missing carrier comparison map", Passed: a.CarrierMap.Verdict == StatusNoK7ToFockMap && a.CarrierMap.DimensionResemblance && !a.CarrierMap.TypedIntertwinerFound && !a.CarrierMap.FunctorFromProjectiveFockToK7Found && !a.CarrierMap.FunctorFromK7ToProjectiveFockFound && strings.Contains(a.CarrierMap.MissingObject, "Theta"), Detail: FormatCarrierMap(a.CarrierMap)},
			{Name: "preserve boundary and 7/72 firewalls", Passed: !a.BoundaryReadiness.BoundaryStressAssignment && !a.BoundaryReadiness.SevenOver72Promoted && !a.BoundaryReadiness.K7ToW7PairingReopened && !a.BoundaryReadiness.K7ToFockMapWouldSuffice && a.BoundaryReadiness.VerdictBoundary == StatusNoBoundaryStressAssignment && a.BoundaryReadiness.VerdictSevenOver72 == StatusNoSevenOver72Theorem, Detail: FormatBoundaryReadiness(a.BoundaryReadiness)},
			{Name: "preserve Gate635 Hodge-polarity selector boundary", Passed: !a.Firewalls.ClaimsK7ToFockSelectorMap && !a.Firewalls.ClaimsK7PlusOnePlusThree && !a.Firewalls.ClaimsBoundaryStressAssignment && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarRGMatching && !a.Firewalls.ClaimsHiggsMassDerivation && !a.Firewalls.ClaimsFlavorDerivation && !a.Firewalls.ClaimsCKMPMNSDerivation && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsPhysicalOrientation && a.Firewalls.Verdict == StatusGate635Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Computed posture: K_7 has native Hodge polarity 4|3; Gate572 has projective Fock 1+3 selector geometry; the resemblance is candidate-only because no typed Theta:K_7->W or W->K_7 carrier map is certified.")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
