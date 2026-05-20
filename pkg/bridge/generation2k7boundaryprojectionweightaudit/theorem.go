package generation2k7boundaryprojectionweightaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2K7BoundaryProjectionWeightAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 K7 boundary projection weight audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate627 K7 boundary projection weight audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate626 boundary-weighted closure as bridge-only", Passed: a.Inherited.Verdict == StatusGate626Inherited && a.Inherited.Gate626ClosureIsBridgeOnly && !a.Inherited.Gate626NativeWeightSource && !a.Inherited.Gate626NativeTransportTheorem && math.Abs(a.Inherited.BoundaryWeight-7.0/72.0) < 1e-15, Detail: FormatInherited(a.Inherited)},
			{Name: "identify exact 7/72 weight and 65/72 complement", Passed: a.Weight.Verdict == StatusWeightIdentified && a.Weight.Numerator == 7 && a.Weight.Denominator == 72 && math.Abs(a.Weight.Value-7.0/72.0) < 1e-15 && a.Weight.ComplementNumerator == 65, Detail: FormatWeight(a.Weight)},
			{Name: "audit numerator as dim K7 candidate", Passed: a.Numerator.Verdict == StatusNumeratorK7Candidate && a.Numerator.MatchesDimK7 && a.Numerator.K7NativeCarrierCertified && !a.Numerator.ProjectionToBoundaryFound, Detail: FormatNumerator(a.Numerator)},
			{Name: "audit denominator 72 candidates without certified boundary carrier", Passed: a.Denominator.Verdict == StatusDenominator72Candidate && a.Denominator.TargetDenominator == 72 && a.Denominator.AnyExistingTypedCandidate && !a.Denominator.CertifiedBoundaryCarrier && len(a.Denominator.Rows) >= 4, Detail: FormatDenominator(a.Denominator)},
			{Name: "audit 65/72 complement as arithmetic only", Passed: a.Complement.Verdict == StatusComplementAudited && a.Complement.ComplementEquals65Over72 && a.Complement.ArithmeticComplementOnly && !a.Complement.NativeComplementCarrier, Detail: FormatComplement(a.Complement)},
			{Name: "rewrite boundary pull through xi midpoint as 7/36", Passed: a.Midpoint.Verdict == StatusMidpointRewriteAudited && a.Midpoint.XiBoundaryInherited && !a.Midpoint.NativeMidpointProjection && math.Abs(a.Midpoint.MidpointPullWeight-7.0/36.0) < 1e-15 && math.Abs(a.Midpoint.RewriteResidual) < 1e-15, Detail: FormatMidpoint(a.Midpoint)},
			{Name: "record missing Pi_K7_to_boundary projection", Passed: a.Projection.Verdict == StatusProjectionMissing && a.Projection.WeightEqualsDimRatio && !a.Projection.ProjectionOperatorExists && !a.Projection.IdempotentCertified && !a.Projection.TraceCertified && !a.Projection.IntertwinerCertified, Detail: FormatProjection(a.Projection)},
			{Name: "avoid coefficient recurrence overclaim", Passed: a.Recurrence.Verdict == StatusNoNativeWeightTheorem && !a.Recurrence.SameCoefficientElsewhere && !a.Recurrence.NativeRecurrenceLaw, Detail: FormatRecurrence(a.Recurrence)},
			{Name: "record missing native 72-carrier and source theorem", Passed: a.NativeStatus.NumeratorK7Native && !a.NativeStatus.Denominator72BoundaryCarrierNative && !a.NativeStatus.K7BoundaryProjectionNative && !a.NativeStatus.ComplementProjectionNative && !a.NativeStatus.GaugeScalarFlavorTransportNative && !a.NativeStatus.SevenOverSeventyTwoSourceTheorem, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate627 firewalls", Passed: !a.Firewalls.ClaimsNativeWeightTheorem && !a.Firewalls.ClaimsCertified72Carrier && !a.Firewalls.ClaimsNativeProjection && !a.Firewalls.ClaimsScalarRGMatching && !a.Firewalls.ClaimsFlavorOrientation && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsHiggsMassDerived && !a.Firewalls.ClaimsEndpointDerivation, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Missing projection: "+strings.TrimSpace(FormatProjection(a.Projection)))
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
