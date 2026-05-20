package generation2k7boundaryprojectionweightaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) {
		return "symbolic"
	}
	if math.IsInf(x, 1) {
		return "+Inf"
	}
	if math.IsInf(x, -1) {
		return "-Inf"
	}
	return fmt.Sprintf("%.15g", x)
}

func FormatInherited(i Gate626Inheritance) string {
	return fmt.Sprintf("kappaSum=%s absLambda12=%s R3MinusOne=%s split=%s weight=%s scalarWeight=%s mixture=%s closureResidual=%s scalarResidual=%s bridgeOnly=%t nativeWeight=%t nativeTransport=%t verdict=%q", f64(i.KappaSum), f64(i.AbsLambda12), f64(i.R3MinusOne), f64(i.BoundarySplit), f64(i.BoundaryWeight), f64(i.ScalarWeight), f64(i.WeightedMixture), f64(i.WeightedClosureResidual), f64(i.ScalarPredictionResidual), i.Gate626ClosureIsBridgeOnly, i.Gate626NativeWeightSource, i.Gate626NativeTransportTheorem, i.Verdict)
}

func FormatWeight(w WeightIdentification) string {
	return fmt.Sprintf("num=%d den=%d value=%s observed=%s residual=%s complement=%d/%d=%s verdict=%q", w.Numerator, w.Denominator, f64(w.Value), f64(w.ObservedWeight), f64(w.Residual), w.ComplementNumerator, w.ComplementDenominator, f64(w.ScalarComplement), w.Verdict)
}

func FormatNumerator(n NumeratorK7Audit) string {
	return fmt.Sprintf("num=%d dimK7=%d rankPB=%d rankPG=%d ambientLambda4=%d vectorDim=%d matches=%t nativeCarrier=%t projection=%t interpretation=%q verdict=%q", n.Numerator, n.K7Dimension, n.RankPB, n.RankPG, n.AmbientExteriorDimension, n.CliffordVectorDimension, n.MatchesDimK7, n.K7NativeCarrierCertified, n.ProjectionToBoundaryFound, n.Interpretation, n.Verdict)
}

func FormatDenominatorRow(r DenominatorCandidateRow) string {
	return fmt.Sprintf("name=%q expression=%q value=%d typed=[%s] ledger=%t boundaryCarrier=%t certified=%t requiresTheorem=%t verdict=%q", r.Name, r.Expression, r.Value, strings.Join(r.TypedFactors, "; "), r.ExistingLedgerData, r.BoundaryCarrier, r.CertifiedAsDenom, r.RequiresNewTheorem, r.Verdict)
}

func FormatDenominatorRows(rows []DenominatorCandidateRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatDenominatorRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatDenominator(d Denominator72Audit) string {
	return fmt.Sprintf("target=%d anyTyped=%t certifiedBoundary=%t best=%q bestValue=%d rows=[%s] verdict=%q", d.TargetDenominator, d.AnyExistingTypedCandidate, d.CertifiedBoundaryCarrier, d.BestCandidate, d.BestCandidateValue, FormatDenominatorRows(d.Rows), d.Verdict)
}

func FormatComplement(c ComplementProjectionAudit) string {
	return fmt.Sprintf("boundaryWeight=%s scalarWeight=%s K7=%d chamber=%d complementDim=%d complementWeight=%s equals65over72=%t arithmeticOnly=%t nativeComplement=%t equation=%q verdict=%q", f64(c.BoundaryWeight), f64(c.ScalarWeight), c.K7Numerator, c.ChamberDenominator, c.ComplementDimension, f64(c.ComplementWeight), c.ComplementEquals65Over72, c.ArithmeticComplementOnly, c.NativeComplementCarrier, c.Equation, c.Verdict)
}

func FormatMidpoint(m MidpointPullRewriteAudit) string {
	return fmt.Sprintf("absLambda12=%s R3MinusOne=%s xi=%s split=%s scalarToMidpoint=%s fullWeight=%s midpointWeight=%s full=%s midpoint=%s residual=%s xiInherited=%t nativeMidpoint=%t equation=%q verdict=%q", f64(m.AbsLambda12), f64(m.R3MinusOne), f64(m.XiBoundary), f64(m.BoundarySplit), f64(m.ScalarToMidpointPull), f64(m.FullSplitWeight), f64(m.MidpointPullWeight), f64(m.WeightedFromFullSplit), f64(m.WeightedFromMidpoint), f64(m.RewriteResidual), m.XiBoundaryInherited, m.NativeMidpointProjection, m.Equation, m.Verdict)
}

func FormatProjection(p BoundaryProjectionOperatorAudit) string {
	return fmt.Sprintf("domain=%q domainDim=%d codomain=%q chamberDim=%d weight=%s dimRatio=%t operator=%t idempotent=%t trace=%t intertwiner=%t missing=%q verdict=%q", p.DomainCarrier, p.DomainDimension, p.CandidateCodomain, p.CandidateChamberDimension, f64(p.ProjectionWeight), p.WeightEqualsDimRatio, p.ProjectionOperatorExists, p.IdempotentCertified, p.TraceCertified, p.IntertwinerCertified, p.MissingObject, p.Verdict)
}

func FormatRecurrenceRow(r RecurrenceAuditRow) string {
	return fmt.Sprintf("coefficient=%q location=%q same=%t forced=%t native=%t comment=%q", r.Coefficient, r.Location, r.SameCoefficient, r.Forced, r.NativeCertified, r.Comment)
}

func FormatRecurrenceRows(rows []RecurrenceAuditRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatRecurrenceRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatRecurrence(r CoefficientRecurrenceAudit) string {
	return fmt.Sprintf("sameElsewhere=%t nativeLaw=%t rows=[%s] verdict=%q", r.SameCoefficientElsewhere, r.NativeRecurrenceLaw, FormatRecurrenceRows(r.Rows), r.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("numK7=%t denom72=%t projection=%t complement=%t transport=%t sourceTheorem=%t statement=%q verdict=%q", n.NumeratorK7Native, n.Denominator72BoundaryCarrierNative, n.K7BoundaryProjectionNative, n.ComplementProjectionNative, n.GaugeScalarFlavorTransportNative, n.SevenOverSeventyTwoSourceTheorem, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("nativeWeight=%t carrier72=%t projection=%t scalarRG=%t flavor=%t gaugeUnification=%t higgsMass=%t endpoint=%t verdict=%q", f.ClaimsNativeWeightTheorem, f.ClaimsCertified72Carrier, f.ClaimsNativeProjection, f.ClaimsScalarRGMatching, f.ClaimsFlavorOrientation, f.ClaimsGaugeUnification, f.ClaimsHiggsMassDerived, f.ClaimsEndpointDerivation, f.Verdict)
}
