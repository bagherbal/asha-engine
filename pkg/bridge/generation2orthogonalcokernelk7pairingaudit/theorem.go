package generation2orthogonalcokernelk7pairingaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2OrthogonalCokernelK7PairingAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 orthogonal cokernel representative and K7 defect pairing audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate631 orthogonal cokernel and K7 pairing audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate630 square index-zero operator", Passed: a.Inherited.Verdict == StatusGate630Inherited && a.Inherited.HDimension == 70 && a.Inherited.UDimension == 56 && a.Inherited.VDimension == 14 && a.Inherited.DirectSumDimension == 70 && a.Inherited.K7Dimension == 7 && a.Inherited.SpanDimension == 63 && a.Inherited.CokernelDimension == 7 && a.Inherited.Index == 0 && a.Inherited.Gate630IndexZero && a.Inherited.Gate630PairingMissing && a.Inherited.Gate630BoundaryMissing && a.Inherited.Gate630FirewallPreserved && math.Abs(a.Inherited.BoundaryWeight-7.0/72.0) < 1e-15, Detail: FormatInherited(a.Inherited)},
			{Name: "define orthogonal cokernel representative W7", Passed: a.OrthogonalW7.Verdict == StatusOrthogonalCokernelRepresentativeDefined && a.OrthogonalW7.AmbientDimension == 70 && a.OrthogonalW7.SpanDimension == 63 && a.OrthogonalW7.WDimension == 7 && a.OrthogonalW7.WOrthogonalToU && a.OrthogonalW7.WOrthogonalToV && a.OrthogonalW7.DirectSumCertified && a.OrthogonalW7.RepresentsCokernel && a.OrthogonalW7.MetricDependent && a.OrthogonalW7.NativeComplementCertified && strings.Contains(a.OrthogonalW7.WDefinition, "perp"), Detail: FormatOrthogonalW7(a.OrthogonalW7)},
			{Name: "write exact defect sequence", Passed: a.ExactSequence.Verdict == StatusExactDefectSequenceWritten && a.ExactSequence.ExactAtK7 && a.ExactSequence.ExactAtDirectSum && a.ExactSequence.ExactAtH && a.ExactSequence.ExactAtW7 && a.ExactSequence.DimensionAlternatingSum == 0 && a.ExactSequence.ExactByRankNullity && strings.Contains(a.ExactSequence.Sequence, "K_7"), Detail: FormatExactSequence(a.ExactSequence)},
			{Name: "sharpen candidate K7 to W7 pairing problem", Passed: a.CandidatePairings.Verdict == StatusNoCanonicalK7ToW7Pairing && len(a.CandidatePairings.Candidates) == 5 && !a.CandidatePairings.CanonicalPairingFound && !a.CandidatePairings.NondegeneratePairingFound && a.CandidatePairings.PairingProblemSharpened && strings.Contains(a.CandidatePairings.MissingObject, "W_7"), Detail: FormatCandidatePairingTable(a.CandidatePairings)},
			{Name: "audit Hodge-star candidate without certifying rank", Passed: a.HodgeStar.Verdict == StatusHodgeStarRequiresExplicitRankTest && a.HodgeStar.HodgeStarTypedOnLambda4 && a.HodgeStar.MapsLambda4ToLambda4 && a.HodgeStar.RequiresOrientationChoice && !a.HodgeStar.RankTestImplemented && !a.HodgeStar.NondegenerateCertified && strings.Contains(a.HodgeStar.Formula, "P_W"), Detail: FormatHodgeStar(a.HodgeStar)},
			{Name: "reject simple projector algebra pairings", Passed: a.ProjectorAlgebra.Verdict == StatusProjectorAlgebraFails && len(a.ProjectorAlgebra.Rows) == 5 && a.ProjectorAlgebra.K7FixedByPB && a.ProjectorAlgebra.K7FixedByPG && a.ProjectorAlgebra.PWKillsUPlusV && !a.ProjectorAlgebra.AnyPairingCertified, Detail: FormatProjectorAlgebra(a.ProjectorAlgebra)},
			{Name: "audit eta pairing without assuming typed Lambda4 eta", Passed: a.Eta.Verdict == StatusNoCanonicalK7ToW7Pairing && !a.Eta.TypedEtaOnLambda4Available && !a.Eta.RankTestImplemented && !a.Eta.PairingCertified && !a.Eta.CompatibilityCertified && strings.Contains(a.Eta.Reason, "no typed eta"), Detail: FormatEta(a.Eta)},
			{Name: "record determinant-line relation without pointwise isomorphism", Passed: a.DeterminantLine.Verdict == StatusK7W7PairingProblemSharpened && a.DeterminantLine.CanonicalLineRelation && !a.DeterminantLine.PointwiseIsomorphism && a.DeterminantLine.OrientationDependent && a.DeterminantLine.CanSupportVolumeBookkeeping && !a.DeterminantLine.CanSupportNormalizedTraceByItself && strings.Contains(a.DeterminantLine.DeterminantRelation, "det(K_7)"), Detail: FormatDeterminantLine(a.DeterminantLine)},
			{Name: "block boundary-stress assignment after pairing audit", Passed: a.BoundaryReadiness.Verdict == StatusNoBoundaryStressAssignment && !a.BoundaryReadiness.K7ToW7PairingCertified && a.BoundaryReadiness.DeterminantLineRelationAvailable && a.BoundaryReadiness.BoundaryPairDimension == 2 && a.BoundaryReadiness.StillRequiresW7ToBoundary && a.BoundaryReadiness.StillRequiresDefectTraceToBoundary && !a.BoundaryReadiness.BoundaryAssignmentCertified, Detail: FormatBoundaryReadiness(a.BoundaryReadiness)},
			{Name: "record native status and missing canonical K7-W7 pairing", Passed: a.NativeStatus.Lambda4Native && a.NativeStatus.AmbientMetricAdmitted && a.NativeStatus.OrthogonalRepresentativeTyped && a.NativeStatus.K7Native && a.NativeStatus.W7DimensionTyped && a.NativeStatus.ExactDefectSequenceTyped && !a.NativeStatus.HodgeStarRankCertified && !a.NativeStatus.ProjectorPairingCertified && !a.NativeStatus.EtaPairingCertified && a.NativeStatus.DeterminantLineRelationTyped && !a.NativeStatus.CanonicalK7ToW7Pairing && !a.NativeStatus.BoundaryStressAssignmentNative, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate631 K7-cokernel pairing firewall", Passed: !a.Firewalls.ClaimsCanonicalK7W7Pairing && !a.Firewalls.ClaimsBoundaryStressAssignment && !a.Firewalls.ClaimsScalarRGMatching && !a.Firewalls.ClaimsHiggsMassDerivation && !a.Firewalls.ClaimsFlavorDerivation && !a.Firewalls.ClaimsCKMPMNSDerivation && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsBoundaryPairNative && !a.Firewalls.ClaimsNativeTraceTheorem, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Missing K7-W7 pairing: "+a.CandidatePairings.MissingObject)
		notes = append(notes, "Missing boundary assignment: "+a.BoundaryReadiness.MissingObject)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
