package generation2k7hodgesignaturestabilizeraudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2K7HodgeSignatureStabilizerAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 634 — K7 Hodge-Signature Stabilizer Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate634 K7 Hodge-signature stabilizer audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate633 K7 Hodge stability", Passed: a.Inherited.Verdict == StatusGate633Inherited && a.Inherited.HDimension == 70 && a.Inherited.K7Dimension == 7 && a.Inherited.StarPreservesK7 && a.Inherited.K7HodgeStable && a.Inherited.NoNewCompanionSevenPlane && a.Inherited.NoK7ToW7Pairing && a.Inherited.NoBoundaryAssignment && a.Inherited.Gate633FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "define restricted Hodge operator S_K", Passed: a.RestrictedOperator.Rows == 7 && a.RestrictedOperator.Cols == 7 && a.RestrictedOperator.Formula == "S_K = Q_K^T S_* Q_K" && a.RestrictedOperator.Verdict == StatusSKOrthogonalSymmetricInvolutive, Detail: FormatRestrictedOperator(a.RestrictedOperator)},
			{Name: "certify S_K symmetric orthogonal involution", Passed: a.RestrictedOperator.Symmetric && a.RestrictedOperator.Orthogonal && a.RestrictedOperator.Involutive && a.RestrictedOperator.SymmetryResidual < 1e-12 && a.RestrictedOperator.OrthogonalityResidual < 1e-12 && a.RestrictedOperator.InvolutionResidual < 1e-12 && math.Abs(a.RestrictedOperator.Trace-1) < 1e-10 && math.Abs(a.RestrictedOperator.Determinant+1) < 1e-8, Detail: FormatRestrictedOperator(a.RestrictedOperator)},
			{Name: "compute K7 Hodge spectrum", Passed: a.Spectrum.Verdict == StatusSpectrumComputed && len(a.Spectrum.Eigenvalues) == 7 && a.Spectrum.PlusRank == 4 && a.Spectrum.MinusRank == 3 && math.Abs(a.Spectrum.Trace-1) < 1e-10 && math.Abs(a.Spectrum.Determinant+1) < 1e-8 && a.Spectrum.Mixed && !a.Spectrum.FullySelfDual && !a.Spectrum.FullyAntiSelfDual, Detail: FormatSpectrum(a.Spectrum)},
			{Name: "certify internal self/anti-self projectors", Passed: a.InternalProjectors.Verdict == StatusInternalProjectorsCertified && a.InternalProjectors.PlusProjectorRank == 4 && a.InternalProjectors.MinusProjectorRank == 3 && a.InternalProjectors.ProjectorsCertified && a.InternalProjectors.PlusProjectorIdempotence < 1e-12 && a.InternalProjectors.MinusProjectorIdempotence < 1e-12 && a.InternalProjectors.ComplementarityResidual < 1e-12 && a.InternalProjectors.OrthogonalityResidual < 1e-12, Detail: FormatInternalProjectors(a.InternalProjectors)},
			{Name: "compute ambient self/anti-self-dual projection weights", Passed: a.AmbientProjection.Verdict == StatusAmbientProjectionComputed && a.AmbientProjection.AmbientSelfDualRank == 35 && a.AmbientProjection.AmbientAntiSelfDualRank == 35 && a.AmbientProjection.AmbientHodgeStarSquaredResidual < 1e-12 && math.Abs(a.AmbientProjection.AmbientTrace) < 1e-12 && math.Abs(a.AmbientProjection.K7SelfDualFrobeniusSquared-4) < 1e-8 && math.Abs(a.AmbientProjection.K7AntiSelfDualFrobeniusSquared-3) < 1e-8 && math.Abs(a.AmbientProjection.K7SelfDualFraction-4.0/7.0) < 1e-8 && math.Abs(a.AmbientProjection.K7AntiSelfDualFraction-3.0/7.0) < 1e-8, Detail: FormatAmbientProjection(a.AmbientProjection)},
			{Name: "classify K7 as mixed Hodge-stable carrier", Passed: a.Classification.Verdict == StatusMixedHodgeSignature && !a.Classification.K7FullySelfDual && !a.Classification.K7FullyAntiSelfDual && a.Classification.K7MixedHodgePolarity && a.Classification.PlusDimension == 4 && a.Classification.MinusDimension == 3 && strings.Contains(a.Classification.Statement, "4+3"), Detail: FormatClassification(a.Classification)},
			{Name: "preserve prior-route and 7/72 firewalls", Passed: !a.Consequences.K7ToW7PairingReopened && !a.Consequences.OctonionicResidualReopened && !a.Consequences.BoundaryAssignmentPromoted && !a.Consequences.SevenOver72Promoted && a.Consequences.VerdictBoundary == StatusNoBoundaryStressAssignment && a.Consequences.VerdictSevenOver72 == StatusNoSevenOver72Theorem, Detail: FormatConsequences(a.Consequences)},
			{Name: "preserve Gate634 K7 Hodge-signature boundary", Passed: !a.Firewalls.ClaimsBoundaryStressAssignment && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarRGMatching && !a.Firewalls.ClaimsHiggsMassDerivation && !a.Firewalls.ClaimsFlavorDerivation && !a.Firewalls.ClaimsCKMPMNSDerivation && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsPhysicalOrientation && a.Firewalls.Verdict == StatusGate634Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Computed result: Spec(S_*|_{K_7}) = {+1,+1,+1,+1,-1,-1,-1}; K_7 has Hodge signature (4,3), trace +1, determinant -1.")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
