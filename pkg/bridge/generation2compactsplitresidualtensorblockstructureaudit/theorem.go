package generation2compactsplitresidualtensorblockstructureaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2CompactSplitResidualTensorBlockStructureAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 643 — CompactSplit ResidualTensor BlockStructure Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate643 residual-tensor block-structure audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate642 angle skeleton and firewall", Passed: a.Inherited.Verdict == StatusGate642AngleInherited && a.Inherited.HodgePolaritySkeleton && !a.Inherited.NativeTraceIdentityCertified && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && !a.Inherited.ScalarFlavorTransport && !a.Inherited.PhysicalAngle && !a.Inherited.PhysicalMetric && a.Inherited.Gate642FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "define unit residual tensors orthogonal to B_hat", Passed: a.ResidualTensor.ResidualTensorsCertified && len(a.ResidualTensor.Routes) >= 3 && a.ResidualTensor.MaxOrthogonalityToBHat < angleTolerance && a.ResidualTensor.MaxResidualUnitNormDrift < angleTolerance, Detail: FormatResidualTensorDetails(a.ResidualTensor)},
			{Name: "compute Hodge-polarity block profiles", Passed: a.BlockSummary.RouteCount >= 3 && a.BlockSummary.HasTypedBlockStructure && !a.BlockSummary.NativeTraceIdentityFound && strings.Contains(a.BlockSummary.Verdict, StatusHodgePolarityBlocksComputed), Detail: FormatBlockSummary(a.BlockSummary)},
			{Name: "classify residual tensor without theorem promotion", Passed: a.Interpretation.AnglePairInherited && a.Interpretation.ResidualTensorDefined && a.Interpretation.BlocksComputed && a.Interpretation.TypedBlockStructure && !a.Interpretation.NativeTraceIdentityFound, Detail: FormatInterpretation(a.Interpretation)},
			{Name: "preserve split-G2, boundary, scalar/flavor, physical-geometry, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsNativeTraceIdentity && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalAngle && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsFlavor && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate643Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := []string{
			StatusGate642AngleInherited,
			StatusResidualTensorDefined,
			StatusHodgePolarityBlocksComputed,
			StatusRouteBlockProfilesComputed,
			StatusResidualBlockStructure,
			StatusSameSectorHodgeDiagonal,
			StatusNoOffSectorCarrier,
			StatusNoNativeTraceIdentity,
			StatusNoCertifiedSplitG2,
			StatusNoBoundaryStress,
			StatusNoSevenOver72Theorem,
			StatusNoScalarFlavorTransport,
			StatusNoPhysicalAngle,
			StatusNoPhysicalMetric,
			StatusNoHiggsFlavorGauge,
			StatusGate643Boundary,
			"Gate643 constructs R_hat=(G_hat-<G_hat,B_hat>B_hat)/rho for omega_1_alt, omega_2_alt, and omega_B_alt, then decomposes it into K7+ x K7+, K7- x K7-, and K7+ x K7- blocks. The repeated result is same-sector Hodge-diagonal: ||R++||^2=3/7, ||R--||^2=4/7, and 2||R+-||^2=0. Therefore the off-sector block does not carry the residual tensor, and no native trace identity derives the 169:48:217 angle pair.",
		}
		if a.BlockSummary.OffSectorDominantRoutes > 0 {
			notes = append(notes, StatusOffSectorCarrierCandidate)
		}
		if !a.BlockSummary.HasTypedBlockStructure {
			notes = append(notes, StatusNoSimpleBlockStructure)
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
