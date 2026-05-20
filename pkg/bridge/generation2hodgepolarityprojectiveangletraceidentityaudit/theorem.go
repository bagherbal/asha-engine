package generation2hodgepolarityprojectiveangletraceidentityaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HodgePolarityProjectiveAngleTraceIdentityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 642 — HodgePolarity ProjectiveAngle TraceIdentity Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate642 hodge-polarity projective-angle trace-identity audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate641 projective angle and firewalls", Passed: a.Inherited.Verdict == StatusGate641AngleInherited && a.Inherited.ComplementIdentified && a.Inherited.ProjectiveAngleAudited && a.Inherited.ThirteenSourcesAudited && !a.Inherited.TraceIdentityCertifiedByGate641 && !a.Inherited.SplitG2CertifiedByGate641 && !a.Inherited.BoundaryStressByGate641 && !a.Inherited.SevenOver72TheoremByGate641 && !a.Inherited.ScalarFlavorByGate641 && !a.Inherited.PhysicalAngleByGate641 && a.Inherited.Gate641FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "compute normalized Frobenius contraction pair 169:217 and 48:217", Passed: a.RawContractions.IntegerRatioVerified && len(a.RawContractions.Contractions) >= 3 && a.RawContractions.MaxCosSquaredDelta < traceIdentityTolerance && a.RawContractions.MaxSinSquaredDelta < traceIdentityTolerance && !a.RawContractions.NativeTraceIdentityFound && strings.Contains(a.RawContractions.Verdict, StatusRawContractionsComputed), Detail: FormatRawContractionsDetails(a.RawContractions)},
			{Name: "compute Hodge-sector block decomposition from p=4 and q=3", Passed: a.SectorBlocks.BlockSkeletonMatches && a.SectorBlocks.AlignmentAmplitude == alignmentRoot && a.SectorBlocks.FailureAmplitudeSquared == failureNumerator && a.SectorBlocks.Denominator == angleDenominator && !a.SectorBlocks.NativeTraceIdentity && strings.Contains(a.SectorBlocks.Verdict, StatusHodgeSectorBlocksComputed) && strings.Contains(a.SectorBlocks.Verdict, StatusHodgePolarityBlockSkeleton), Detail: FormatSectorBlocksDetails(a.SectorBlocks)},
			{Name: "audit projective pair and tangent skeleton", Passed: a.ProjectivePair.PairMatches && !a.ProjectivePair.DerivedFromNativeTraceIdentity && a.ProjectivePair.PythagoreanIntegerResidual == 0 && a.ProjectivePair.TanSquaredNumerator == failureNumerator && a.ProjectivePair.TanSquaredDenominator == alignmentNumerator && a.ProjectivePair.Verdict == StatusHodgePolarityBlockSkeleton, Detail: FormatProjectivePair(a.ProjectivePair)},
			{Name: "search trace/projector identity without certifying theorem", Passed: !a.TraceIdentity.NativeTraceIdentityFound && a.TraceIdentity.BestCandidateResidual < traceIdentityTolerance && strings.Contains(a.TraceIdentity.Verdict, StatusTraceIdentityCandidates) && strings.Contains(a.TraceIdentity.Verdict, StatusNoNativeTraceIdentity), Detail: FormatTraceDetails(a.TraceIdentity)},
			{Name: "classify projective angle as internal obstruction-only skeleton", Passed: a.Classification.ProjectiveAngleInherited && a.Classification.RawContractionsComputed && a.Classification.HodgeSectorBlocksComputed && a.Classification.BlockSkeletonSupported && !a.Classification.NativeTraceIdentityCertified && a.Classification.ObstructionOnly && a.Classification.Verdict == StatusHodgePolarityBlockSkeleton, Detail: FormatClassification(a.Classification)},
			{Name: "preserve trace, split-G2, boundary, 7/72, scalar/flavor, physical, Higgs/flavor/gauge firewalls", Passed: !a.Firewalls.ClaimsNativeTraceIdentity && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalAngle && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsFlavor && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate642Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Computed result: cos(theta)≈"+f64(a.Inherited.CosTheta)+", sin(theta)≈"+f64(a.Inherited.SinTheta)+", with cos^2≈"+f64(a.Inherited.CosSquared)+"=169/217 and sin^2≈"+f64(a.Inherited.SinSquared)+"=48/217. Hodge-polarity skeleton: p=4, q=3, p^2-q=13, p^2*q=48, (p^2-q)^2+p^2*q=217. No native trace identity is certified.")
		// Keep math imported in theorem.go for compile-time guard on the exact angle relation.
		if math.Abs(a.Inherited.CosTheta-(float64(alignmentRoot)/math.Sqrt(float64(angleDenominator)))) > traceIdentityTolerance {
			notes = append(notes, StatusTraceIdentityArtifact)
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
