package generation2k7kernelcokernelindexzeroaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2K7KernelCokernelIndexZeroAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 K7 kernel-cokernel index-zero audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate630 K7 kernel-cokernel index-zero audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate629 intersection-cokernel dual candidate and missing maps", Passed: a.Inherited.Verdict == StatusGate629Inherited && a.Inherited.UDimension == 56 && a.Inherited.VDimension == 14 && a.Inherited.DirectSumDimension == 70 && a.Inherited.Lambda4Dimension == 70 && a.Inherited.IntersectionDimension == 7 && a.Inherited.SpanDimension == 63 && a.Inherited.CokernelDimension == 7 && a.Inherited.BoundaryPairDimension == 2 && a.Inherited.AugmentedChamberDimension == 72 && a.Inherited.Gate629DualCandidate && a.Inherited.Gate629IsomorphismMissing && a.Inherited.Gate629BoundaryAssignmentMissing && a.Inherited.Gate629FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "define square addition map A from U direct-sum V to Lambda4", Passed: a.AdditionMap.Verdict == StatusAdditionMapDefined && a.AdditionMap.MapName == "A" && a.AdditionMap.DomainDimension == 70 && a.AdditionMap.CodomainDimension == 70 && a.AdditionMap.SquareOperator && strings.Contains(a.AdditionMap.Formula, "u+v"), Detail: FormatAdditionMap(a.AdditionMap)},
			{Name: "compute kernel of A as K7", Passed: a.AdditionMap.KernelIsK7 && a.AdditionMap.KernelDimension == 7 && strings.Contains(a.AdditionMap.KernelExpression, "K_7"), Detail: FormatAdditionMap(a.AdditionMap)},
			{Name: "compute image and cokernel of A", Passed: a.AdditionMap.ImageIsSpan && a.AdditionMap.ImageDimension == 63 && a.AdditionMap.CokernelMatchesK7 && a.AdditionMap.CokernelDimension == 7 && a.AdditionMap.RankDefect == 7, Detail: FormatAdditionMap(a.AdditionMap)},
			{Name: "compute index-zero Boolean-octonionic defect", Passed: a.Defect.Verdict == StatusIndexZeroComputed && a.Defect.KernelDimension == 7 && a.Defect.CokernelDimension == 7 && a.Defect.DefectsBalanced && a.Defect.Index == 0 && a.Defect.CandidateDefectPair && a.Defect.FredholmAnalogyOnly && strings.Contains(a.Defect.MissingPairing, "ker(A)"), Detail: FormatDefect(a.Defect)},
			{Name: "compress finite chamber into K7 blocks", Passed: a.BlockCompression.Verdict == StatusK7DefectBlockCandidate && a.BlockCompression.K7BlockDimension == 7 && a.BlockCompression.PBBlocks == 8 && a.BlockCompression.PGBlocks == 2 && a.BlockCompression.SpanBlocks == 9 && a.BlockCompression.Lambda4Blocks == 10 && a.BlockCompression.BoundaryCoordinates == 2 && a.BlockCompression.CompressionExact && a.BlockCompression.DefectBlockCandidate && math.Abs(a.BlockCompression.BoundaryWeight-7.0/72.0) < 1e-15, Detail: FormatBlockCompression(a.BlockCompression)},
			{Name: "block canonical kernel-cokernel pairing candidates", Passed: a.Pairing.Verdict == StatusNoCanonicalKerCokerPairing && len(a.Pairing.Candidates) == 4 && !a.Pairing.CanonicalPairingFound && !a.Pairing.MetricPairingCertified && !a.Pairing.HodgeStarPairingCertified && !a.Pairing.EtaPairingCertified && !a.Pairing.ProjectorPairingCertified && strings.Contains(a.Pairing.MissingObject, "ker(A)"), Detail: FormatPairing(a.Pairing)},
			{Name: "block boundary-stress assignment from balanced defect", Passed: a.BoundaryAssignment.Verdict == StatusNoNativeBoundaryStressFromDefect && a.BoundaryAssignment.DefectBlockCanSupplySeven && a.BoundaryAssignment.BoundaryPairDimension == 2 && math.Abs(a.BoundaryAssignment.BoundaryWeight-7.0/72.0) < 1e-15 && !a.BoundaryAssignment.AssignmentCertified && !a.BoundaryAssignment.NativeTransportTheorem, Detail: FormatBoundaryAssignment(a.BoundaryAssignment)},
			{Name: "record native status and missing trace theorem", Passed: a.NativeStatus.Lambda4Native && a.NativeStatus.UImageRankNative && a.NativeStatus.VImageRankNative && a.NativeStatus.K7IntersectionNative && a.NativeStatus.AdditionMapTyped && a.NativeStatus.KernelDimensionTyped && a.NativeStatus.CokernelDimensionTyped && a.NativeStatus.IndexZeroTyped && !a.NativeStatus.CanonicalKernelCokernelPairing && !a.NativeStatus.BoundaryStressAssignmentNative && !a.NativeStatus.K7DefectBoundaryTraceTheorem, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate630 defect-pairing firewall", Passed: !a.Firewalls.ClaimsCanonicalPairing && !a.Firewalls.ClaimsBoundaryStressAssignment && !a.Firewalls.ClaimsK7DefectTraceTheorem && !a.Firewalls.ClaimsBoundaryPairNative && !a.Firewalls.ClaimsScalarRGMatching && !a.Firewalls.ClaimsFlavorOrientation && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsHiggsMassDerived && !a.Firewalls.ClaimsEndpointDerivation, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Missing ker-coker pairing: "+a.Pairing.MissingObject)
		notes = append(notes, "Missing boundary assignment: "+a.BoundaryAssignment.MissingObject)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
