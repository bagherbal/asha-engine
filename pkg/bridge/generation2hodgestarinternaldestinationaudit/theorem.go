package generation2hodgestarinternaldestinationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HodgeStarInternalDestinationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 633 — Hodge-Star Internal Destination and Octonionic Residual Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate633 Hodge-star internal destination audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate632 transverse Hodge failure", Passed: a.Inherited.Verdict == StatusGate632Inherited && a.Inherited.HDimension == 70 && a.Inherited.UDimension == 56 && a.Inherited.VDimension == 14 && a.Inherited.K7Dimension == 7 && a.Inherited.SpanDimension == 63 && a.Inherited.W7Dimension == 7 && a.Inherited.LeakageRank == 0 && a.Inherited.PWStarK7FrobeniusNorm < 1e-10 && a.Inherited.PUVStarK7FrobeniusNorm > 2.6 && a.Inherited.TransverseFailureCertified && a.Inherited.StarK7InsideUPlusV && a.Inherited.NoBoundaryAssignment && a.Inherited.Gate632FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "define L7=*K7 and confirm internal containment", Passed: a.L7.Verdict == StatusHodgeCompanionL7Defined && a.L7.Rows == 70 && a.L7.Cols == 7 && a.L7.Rank == 7 && a.L7.QLOthonormalResidual < 1e-12 && a.L7.PWQLFrobeniusNorm < 1e-10 && math.Abs(a.L7.PUVProjectionFraction-1) < 1e-8 && a.L7.PUVContainmentResidual < 1e-10 && a.L7.StarTwoCycleResidual < 1e-12 && a.L7.InternalContainmentCertified, Detail: FormatL7(a.L7)},
			{Name: "test and certify K7 Hodge stability", Passed: a.K7Preservation.Verdict == StatusHodgeStarPreservesK7 && a.K7Preservation.Rank == 7 && len(a.K7Preservation.SingularValues) == 7 && math.Abs(a.K7Preservation.ProjectionFraction-1) < 1e-8 && a.K7Preservation.ContainmentResidual < 1e-10 && a.K7Preservation.StarPreservesK7 && a.K7Preservation.HodgeStable && math.Abs(math.Abs(a.K7Preservation.Determinant)-1) < 1e-8, Detail: FormatK7Preservation(a.K7Preservation)},
			{Name: "rule out T56 internal complement destination", Passed: a.T56Complement.Verdict == StatusNoObliqueInternalSevenPlane && a.T56Complement.Dimension == 56 && a.T56Complement.Rank == 0 && a.T56Complement.ProjectionFraction < 1e-20 && a.T56Complement.ContainmentResidual > 2.6 && !a.T56Complement.L7InsideT56, Detail: FormatT56(a.T56Complement)},
			{Name: "test and reject octonionic residual V0 destination", Passed: a.OctonionicResidual.Verdict == StatusDoesNotPairOctonionicResidualV0 && a.OctonionicResidual.VDimension == 14 && a.OctonionicResidual.K7Dimension == 7 && a.OctonionicResidual.V0Dimension == 7 && a.OctonionicResidual.Rank == 0 && a.OctonionicResidual.ProjectionFraction < 1e-20 && a.OctonionicResidual.ContainmentResidual > 2.6 && !a.OctonionicResidual.L7EqualsV0 && !a.OctonionicResidual.VDecomposesAsK7StarK7, Detail: FormatOctonionicResidual(a.OctonionicResidual)},
			{Name: "test and reject Boolean residual U0 destination", Passed: a.BooleanResidual.Verdict == StatusDoesNotEnterBooleanResidualU0 && a.BooleanResidual.UDimension == 56 && a.BooleanResidual.K7Dimension == 7 && a.BooleanResidual.U0Dimension == 49 && a.BooleanResidual.Rank == 0 && a.BooleanResidual.ProjectionFraction < 1e-20 && a.BooleanResidual.ContainmentResidual > 2.6 && !a.BooleanResidual.L7InsideU0, Detail: FormatBooleanResidual(a.BooleanResidual)},
			{Name: "rule out oblique internal seven-plane", Passed: a.ObliqueDecomposition.Verdict == StatusNoObliqueInternalSevenPlane && a.ObliqueDecomposition.CandidateT56Dimension == 56 && a.ObliqueDecomposition.ProjectionFraction < 1e-20 && a.ObliqueDecomposition.ContainmentResidual > 2.6 && a.ObliqueDecomposition.DirectSumCoordinateNorm < 1e-10 && !a.ObliqueDecomposition.ObliquePlaneDetected, Detail: FormatOblique(a.ObliqueDecomposition)},
			{Name: "record Hodge two-cycle classification", Passed: a.StarTwoCycle.Verdict == StatusHodgeStarPreservesK7 && a.StarTwoCycle.StarSquaredResidual < 1e-12 && a.StarTwoCycle.L7EqualsK7 && !a.StarTwoCycle.L7EqualsV0 && !a.StarTwoCycle.L7EqualsU0 && strings.Contains(a.StarTwoCycle.CarrierClassification, "Hodge-stable"), Detail: FormatStarTwoCycle(a.StarTwoCycle)},
			{Name: "preserve 7/72 and boundary firewall consequences", Passed: a.ConsequenceFor7Over72.Verdict == StatusNoBoundaryStressAssignment && a.ConsequenceFor7Over72.K7Stable && !a.ConsequenceFor7Over72.Octonionic14HodgeSplit && !a.ConsequenceFor7Over72.NewSevenPlaneDiscovered && a.ConsequenceFor7Over72.BoundaryPairDimension == 2 && !a.ConsequenceFor7Over72.TraceWeightPromoted && a.ConsequenceFor7Over72.BoundaryAssignmentMissing, Detail: FormatConsequence(a.ConsequenceFor7Over72)},
			{Name: "preserve Gate633 internal Hodge destination boundary", Passed: !a.Firewalls.ClaimsBoundaryStressAssignment && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarRGMatching && !a.Firewalls.ClaimsHiggsMassDerivation && !a.Firewalls.ClaimsFlavorDerivation && !a.Firewalls.ClaimsCKMPMNSDerivation && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate633Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Computed result: L_7=*K_7 equals K_7 at the certified tolerance; the V_0=V⊖K_7 hypothesis fails.")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
