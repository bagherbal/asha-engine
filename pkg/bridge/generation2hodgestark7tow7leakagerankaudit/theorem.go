package generation2hodgestark7tow7leakagerankaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HodgeStarK7ToW7LeakageRankAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 632 — Hodge-Star K7-to-W7 Leakage Rank Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate632 Hodge-star leakage rank audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate631 explicit K7-to-W7 pairing problem", Passed: a.Inherited.Verdict == StatusGate631Inherited && a.Inherited.HDimension == 70 && a.Inherited.UDimension == 56 && a.Inherited.VDimension == 14 && a.Inherited.K7Dimension == 7 && a.Inherited.SpanDimension == 63 && a.Inherited.W7Dimension == 7 && a.Inherited.IndexZeroInherited && a.Inherited.OrthogonalRepresentative && a.Inherited.PairingProblemSharpened && a.Inherited.ProjectorAlgebraFailed && a.Inherited.HodgeRankTestRequired && a.Inherited.BoundaryAssignmentMissing && a.Inherited.Gate631FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "construct typed Hodge-star matrix on Lambda4 R8", Passed: a.HodgeStar.Verdict == StatusHodgeStarTyped && a.HodgeStar.TypedOnLambda4R8 && a.HodgeStar.MapsLambda4ToLambda4 && a.HodgeStar.MatrixDimension == 70 && a.HodgeStar.StarSquaredResidual < 1e-12 && math.Abs(a.HodgeStar.Trace) < 1e-12 && a.HodgeStar.SelfDualDimension == 35 && a.HodgeStar.AntiSelfDualDimension == 35, Detail: FormatHodgeStar(a.HodgeStar)},
			{Name: "certify K7 and W7 orthonormal bases", Passed: a.Basis.Verdict == StatusK7AndW7BasesCertified && a.Basis.QKRows == 70 && a.Basis.QKCols == 7 && a.Basis.QWRows == 70 && a.Basis.QWCols == 7 && a.Basis.SpanCols == 63 && a.Basis.QKOrthonormalResidual < 1e-8 && a.Basis.QWOrthonormalResidual < 1e-8 && a.Basis.QWQKOrthogonalityResidual < 1e-8 && a.Basis.PBQKMinusQKResidual < 1e-8 && a.Basis.PGQKMinusQKResidual < 1e-8 && a.Basis.PBQWResidual < 1e-8 && a.Basis.PGQWResidual < 1e-8 && a.Basis.QWOrthogonalToUAndV && a.Basis.K7ContainedInUAndV, Detail: FormatBasis(a.Basis)},
			{Name: "compute Hodge leakage matrix and rank", Passed: a.Leakage.Verdict == StatusHodgeStarDoesNotPairK7ToW7 && a.Leakage.Rows == 7 && a.Leakage.Cols == 7 && a.Leakage.Rank == 0 && len(a.Leakage.SingularValues) == 7 && a.Leakage.FrobeniusNorm < 1e-10 && math.Abs(a.Leakage.Determinant) < 1e-14 && strings.Contains(a.Leakage.Classification, "rank-zero"), Detail: FormatLeakage(a.Leakage)},
			{Name: "verify star K7 image is contained in U+V rather than W7", Passed: a.ImageContainment.Verdict == StatusHodgeStarDoesNotPairK7ToW7 && a.ImageContainment.StarK7FrobeniusNorm > 2.6 && a.ImageContainment.PWStarK7FrobeniusNorm < 1e-10 && a.ImageContainment.PUVStarK7FrobeniusNorm > 2.6 && a.ImageContainment.StarK7ContainedInUPlusV && !a.ImageContainment.TransverseComponentDetected, Detail: FormatImageContainment(a.ImageContainment)},
			{Name: "record degenerate pairing metric", Passed: a.PairingMetric.Verdict == StatusHodgeStarDoesNotPairK7ToW7 && a.PairingMetric.Computed && !a.PairingMetric.RankFull && a.PairingMetric.Degenerate && !a.PairingMetric.ConformalOrIsometric && !a.PairingMetric.AnisotropicNondegenerate && a.PairingMetric.Trace < 1e-20, Detail: FormatPairingMetric(a.PairingMetric)},
			{Name: "block orientation promotion because determinant is zero", Passed: a.Orientation.Verdict == StatusHodgeStarDoesNotPairK7ToW7 && !a.Orientation.NonzeroDeterminant && a.Orientation.Sign == 0 && !a.Orientation.PhysicalOrientationCertified, Detail: FormatOrientation(a.Orientation)},
			{Name: "audit alternative star/projector composites", Passed: a.AlternativeComposites.Verdict == StatusAlternativeCompositesNoHigherRank && len(a.AlternativeComposites.Rows) == 5 && !a.AlternativeComposites.AnyHigherRankThanPhiStar && !a.AlternativeComposites.AnyNondegenerate, Detail: FormatAlternativeComposites(a.AlternativeComposites)},
			{Name: "preserve boundary readiness firewall", Passed: a.BoundaryReadiness.Verdict == StatusNoBoundaryStressAssignment && !a.BoundaryReadiness.HodgePairingCertified && !a.BoundaryReadiness.K7ToW7PairingFound && a.BoundaryReadiness.BoundaryPairDimension == 2 && !a.BoundaryReadiness.BoundaryAssignmentCertified && strings.Contains(a.BoundaryReadiness.MissingObject, "R^2_boundary"), Detail: FormatBoundaryReadiness(a.BoundaryReadiness)},
			{Name: "record native status after failed Hodge leakage route", Passed: a.NativeStatus.Verdict == StatusNoCanonicalK7W7PairingFound && a.NativeStatus.Lambda4Native && a.NativeStatus.PBPGProjectorsConstructed && a.NativeStatus.K7FrameConstructed && a.NativeStatus.W7FrameConstructed && a.NativeStatus.HodgeStarMatrixConstructed && a.NativeStatus.HodgeRankComputed && !a.NativeStatus.HodgePairingNondegenerate && !a.NativeStatus.CanonicalK7ToW7PairingFound && !a.NativeStatus.BoundaryStressAssignmentNative, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate632 Hodge pairing boundary", Passed: !a.Firewalls.ClaimsBoundaryStressAssignment && !a.Firewalls.ClaimsScalarRGMatching && !a.Firewalls.ClaimsHiggsMassDerivation && !a.Firewalls.ClaimsFlavorDerivation && !a.Firewalls.ClaimsCKMPMNSDerivation && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsPhysicalOrientation && !a.Firewalls.ClaimsNativeTraceWeight && a.Firewalls.Verdict == StatusGate632Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Computed rank(Q_W^T * Q_K)=0; Hodge-star K7-to-W7 leakage route fails at the certified tolerance.")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
