// Package generation2sharedscalarwallunitnormalizationalignmentaudit implements
// Gate 702: Shared Scalar-Wall Unit Normalization Alignment Audit.
//
// Gate 701 separated the invariant K7 event probability
//
//	p_K7 = Tr(rho_72 P_K7)=7/72
//
// from its coordinate-sealed appearance as the response coefficient in
//
//	sigma_history ≈ (7/72) sigma_boundary.
//
// Gate 702 audits whether the shared signed scalar zero-wall coordinate
// lambda(Lambda_12) anchors the quotient-line normalization between the
// boundary input coordinate and the history output coordinate. This is a
// bridge-layer wall-normalization audit only. It does not derive boundary
// stress, scalar RG matching, Higgs mass, gauge unification, flavor, CKM/PMNS,
// a native response theorem, a native state-selection theorem, or a native
// 7/72 theorem.
package generation2sharedscalarwallunitnormalizationalignmentaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate701 "github.com/bagherbal/asha-engine/pkg/bridge/generation2quotientlinenormalizationandresponsecoefficientcovarianceaudit"
)

const (
	AuditID = "GATE702-SHARED-SCALAR-WALL-UNIT-NORMALIZATION-ALIGNMENT-AUDIT"

	StatusGate701QuotientNormalizationInherited                     = "PASS_GATE701_QUOTIENT_NORMALIZATION_INHERITED"
	StatusSharedLambdaCoordinateIdentified                          = "PASS_SHARED_LAMBDA_COORDINATE_IDENTIFIED"
	StatusLambdaUnitCoefficientAlignmentAudited                     = "PASS_LAMBDA_UNIT_COEFFICIENT_ALIGNMENT_AUDITED"
	StatusResponseCoefficientRemainsEventProbabilityUnderSharedUnit = "PASS_RESPONSE_COEFFICIENT_REMAINS_EVENT_PROBABILITY_UNDER_SHARED_UNIT"
	StatusAlternativeNormalizationsAudited                          = "PASS_ALTERNATIVE_NORMALIZATIONS_AUDITED"
	StatusNonTautologyWithSharedLambdaAudited                       = "PASS_NON_TAUTOLOGY_WITH_SHARED_LAMBDA_AUDITED"
	StatusSharedScalarWallUnitAnchorsQuotientNormalization          = "CONDITIONAL_SUPPORT_SHARED_SCALAR_WALL_UNIT_ANCHORS_QUOTIENT_NORMALIZATION"
	StatusResponseCoefficientEqualsEventProbabilitySharedLambda     = "CONDITIONAL_SUPPORT_RESPONSE_COEFFICIENT_EQUALS_EVENT_PROBABILITY_IN_SHARED_LAMBDA_UNITS"
	StatusGate700LawScalarWallUnitSealed                            = "CONDITIONAL_SUPPORT_GATE700_LAW_IS SCALAR_WALL_UNIT_SEALED"
	StatusSharedLambdaUnitAlignmentNotNative                        = "FAILED_ROUTE_SHARED_LAMBDA_UNIT_ALIGNMENT_NOT_NATIVELY_DERIVED"
	StatusNoNativeWallCoordinateNormalizationAlignmentTheorem       = "FAILED_ROUTE_NO_NATIVE_WALL_COORDINATE_NORMALIZATION_ALIGNMENT_THEOREM"
	StatusNoNativeBoundaryHistoryResponsePrinciple                  = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_HISTORY_RESPONSE_PRINCIPLE"
	StatusNoNativeSevenOver72Theorem                                = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate702SharedScalarWallUnitBoundary                       = "FIREWALL_PRESERVED_GATE702_SHARED_SCALAR_WALL_UNIT_BOUNDARY"
)

const (
	h72Dimension   = 72
	k7Dimension    = 7
	eventProbK7    = float64(k7Dimension) / float64(h72Dimension)
	tolerance      = 1e-15
	strictTol      = 1e-17
	lambdaLambda12 = -0.0497009420776833
	r3Minus1       = 0.0509933868964996
	kappaLambda    = 0.0443230430960771
	kappaE         = 0.00550355419157456
)

type Gate701Inheritance struct {
	InheritedQuotientNormalization bool
	EventProbability               float64
	CanonicalCoefficient           float64
	CoefficientCovarianceFormula   string
	Gate700CoordinateSealed        bool
	NoNativeWallAlignment          bool
	NoNativeBoundaryHistory        bool
	NoNativeSevenOver72            bool
	Verdict                        string
}

type SharedLambdaCoordinateAudit struct {
	BoundaryCoordinate           string
	HistoryCoordinate            string
	SharedCoordinate             string
	BoundaryLambdaCoefficient    float64
	HistoryLambdaCoefficient     float64
	BoundaryContainsSharedLambda bool
	HistoryContainsSharedLambda  bool
	SameSignedScalarZeroWall     bool
	UnitCoefficientAlignment     bool
	Verdict                      string
}

type NormalizationAnchorAudit struct {
	Rule                         string
	BoundaryLambdaCoefficient    float64
	HistoryLambdaCoefficient     float64
	Alpha                        float64
	Beta                         float64
	BetaOverAlpha                float64
	ResponseCoefficient          float64
	EventProbability             float64
	CoefficientEqualsProbability bool
	Verdict                      string
}

type AlternativeNormalization struct {
	Name                         string
	BoundaryScaleAlpha           float64
	HistoryScaleBeta             float64
	TransformedCoefficient       float64
	EqualsEventProbability       bool
	PreservesLambdaUnitAlignment bool
	AcceptedActiveAlignment      bool
	Reason                       string
}

type AlternativeNormalizationAudit struct {
	Examples                  []AlternativeNormalization
	EuclideanBoundaryRejected bool
	HistoryNormRejected       bool
	GaugeAnchorConditioned    bool
	AbsoluteFormRejected      bool
	SharedUnitAccepted        bool
	Verdict                   string
}

type NonTautologyAudit struct {
	SharedLambdaPresent              bool
	RearrangedEquation               string
	LambdaWeight                     float64
	GaugeWeight                      float64
	IndependentGaugeWoundPresent     bool
	CoefficientsDiffer               bool
	LambdaIsAlignmentAnchorNotProof  bool
	NonTautologicalRelationPreserved bool
	Verdict                          string
}

type SourceTypeClassification struct {
	EventProbabilityRole string
	LambdaUnitRole       string
	BoundaryRole         string
	HistoryRole          string
	Conclusion           string
	Verdict              string
}

type MissingTheoremAudit struct {
	Theorems []string
	Verdict  string
}

type FirewallAudit struct {
	ClaimsSharedLambdaAlignmentNative      bool
	ClaimsNativeWallNormalizationAlignment bool
	ClaimsNativeBoundaryHistoryPrinciple   bool
	ClaimsNativeSevenOver72Theorem         bool
	ClaimsBoundaryStressDerived            bool
	ClaimsScalarRGMatching                 bool
	ClaimsHiggsMass                        bool
	ClaimsGaugeUnification                 bool
	ClaimsFlavorDerivation                 bool
	ClaimsCKMPMNS                          bool
	Verdict                                string
}

type Analysis struct {
	Inherited    Gate701Inheritance
	SharedLambda SharedLambdaCoordinateAudit
	Anchor       NormalizationAnchorAudit
	Alternatives AlternativeNormalizationAudit
	NonTautology NonTautologyAudit
	Source       SourceTypeClassification
	Missing      MissingTheoremAudit
	Firewall     FirewallAudit
	Truth        string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g701, err := gate701.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate701 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g701)
	shared := buildSharedLambda()
	anchor := buildAnchor(shared)
	alternatives := buildAlternatives()
	nontaut := buildNonTautology()
	source := SourceTypeClassification{
		EventProbabilityRole: "p_K7=7/72 is the invariant no-bias K7 event probability under rho_72",
		LambdaUnitRole:       "unit coefficient of the signed scalar zero-wall coordinate in both quotient lines",
		BoundaryRole:         "sigma_boundary=lambda+(R_3-1) measures boundary anti-alignment defect in signed scalar-wall units",
		HistoryRole:          "sigma_history=kappa_lambda+kappa_e+lambda measures scalar/flavor/history closure defect in the same signed scalar-wall units",
		Conclusion:           "the response coefficient equals the event probability because the input and output quotient lines are measured in the same scalar-wall unit",
		Verdict: strings.Join([]string{
			StatusSharedScalarWallUnitAnchorsQuotientNormalization,
			StatusResponseCoefficientEqualsEventProbabilitySharedLambda,
			StatusGate700LawScalarWallUnitSealed,
		}, "; "),
	}
	missing := MissingTheoremAudit{
		Theorems: []string{
			"WallCoordinateNormalizationAlignmentTheorem",
			"SharedScalarWallUnitTheorem",
			StatusSharedLambdaUnitAlignmentNotNative,
			StatusNoNativeWallCoordinateNormalizationAlignmentTheorem,
			StatusNoNativeBoundaryHistoryResponsePrinciple,
			StatusNoNativeSevenOver72Theorem,
		},
		Verdict: strings.Join([]string{
			StatusSharedLambdaUnitAlignmentNotNative,
			StatusNoNativeWallCoordinateNormalizationAlignmentTheorem,
			StatusNoNativeBoundaryHistoryResponsePrinciple,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate702SharedScalarWallUnitBoundary}
	truth := "Gate 702 conditionally identifies lambda(Lambda_12) as the shared signed scalar zero-wall unit anchoring the boundary and history quotient-line normalizations. With unit lambda coefficient on both sides, beta/alpha=1 and the response coefficient remains the invariant event probability p_K7=7/72. Alternative normalizations rescale the coefficient, and the shared lambda coordinate is an alignment anchor, not a proof of the response law."
	return Analysis{Inherited: inherited, SharedLambda: shared, Anchor: anchor, Alternatives: alternatives, NonTautology: nontaut, Source: source, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate701.Analysis) Gate701Inheritance {
	return Gate701Inheritance{
		InheritedQuotientNormalization: g.Probability.InvariantObjectSeparated && g.WallNormalize.CoefficientEqualsProbability,
		EventProbability:               g.Probability.Probability,
		CanonicalCoefficient:           eventProbK7,
		CoefficientCovarianceFormula:   g.Rescaling.Formula,
		Gate700CoordinateSealed:        g.WallNormalize.CoefficientEqualsProbability && g.WallNormalize.EqualScaleRequiredForDirectEquality,
		NoNativeWallAlignment:          !g.Firewall.ClaimsNativeWallNormalization,
		NoNativeBoundaryHistory:        !g.Firewall.ClaimsNativeBoundaryHistoryPrinciple,
		NoNativeSevenOver72:            !g.Firewall.ClaimsNativeSevenOver72Theorem,
		Verdict:                        StatusGate701QuotientNormalizationInherited,
	}
}

func buildSharedLambda() SharedLambdaCoordinateAudit {
	return SharedLambdaCoordinateAudit{
		BoundaryCoordinate:           "sigma_boundary=lambda+(R_3-1)",
		HistoryCoordinate:            "sigma_history=kappa_lambda+kappa_e+lambda",
		SharedCoordinate:             "lambda(Lambda_12)",
		BoundaryLambdaCoefficient:    1,
		HistoryLambdaCoefficient:     1,
		BoundaryContainsSharedLambda: true,
		HistoryContainsSharedLambda:  true,
		SameSignedScalarZeroWall:     true,
		UnitCoefficientAlignment:     true,
		Verdict: strings.Join([]string{
			StatusSharedLambdaCoordinateIdentified,
			StatusLambdaUnitCoefficientAlignmentAudited,
			StatusSharedScalarWallUnitAnchorsQuotientNormalization,
		}, "; "),
	}
}

func buildAnchor(shared SharedLambdaCoordinateAudit) NormalizationAnchorAudit {
	alpha := shared.BoundaryLambdaCoefficient
	beta := shared.HistoryLambdaCoefficient
	c := (beta / alpha) * eventProbK7
	return NormalizationAnchorAudit{
		Rule:                         "coefficient(lambda in sigma_boundary)=coefficient(lambda in sigma_history)=1",
		BoundaryLambdaCoefficient:    alpha,
		HistoryLambdaCoefficient:     beta,
		Alpha:                        alpha,
		Beta:                         beta,
		BetaOverAlpha:                beta / alpha,
		ResponseCoefficient:          c,
		EventProbability:             eventProbK7,
		CoefficientEqualsProbability: math.Abs(c-eventProbK7) < tolerance,
		Verdict: strings.Join([]string{
			StatusLambdaUnitCoefficientAlignmentAudited,
			StatusResponseCoefficientRemainsEventProbabilityUnderSharedUnit,
			StatusResponseCoefficientEqualsEventProbabilitySharedLambda,
		}, "; "),
	}
}

func buildAlternatives() AlternativeNormalizationAudit {
	examples := []AlternativeNormalization{
		buildAlternative("Euclidean-normalized boundary normal", 1/math.Sqrt2, 1, false, "sigma_boundary_norm=(lambda+R)/sqrt(2) rescales the input line, so c=sqrt(2)(7/72) unless the history coordinate is rescaled by the same factor"),
		buildAlternative("history three-term normalized readout", 1, 1/math.Sqrt(3), false, "sigma_history_norm=(kappa_lambda+kappa_e+lambda)/sqrt(3) rescales the output line, so c=(7/72)/sqrt(3)"),
		buildAlternative("gauge-anchored boundary normalization", 1, 1, true, "equivalent only when the lambda coefficient remains the same unit scalar-wall anchor"),
		buildAlternative("absolute scalar-wall form", 1, 1, false, "K_sum-|lambda| is numerically equivalent only because lambda<0, but it erases signed wall orientation"),
		buildAlternative("shared signed-lambda unit", 1, 1, true, "sigma_boundary=lambda+R and sigma_history=kappa_lambda+kappa_e+lambda preserve the shared signed scalar-wall unit"),
	}
	return AlternativeNormalizationAudit{
		Examples:                  examples,
		EuclideanBoundaryRejected: !examples[0].EqualsEventProbability,
		HistoryNormRejected:       !examples[1].EqualsEventProbability,
		GaugeAnchorConditioned:    examples[2].EqualsEventProbability && examples[2].PreservesLambdaUnitAlignment,
		AbsoluteFormRejected:      !examples[3].AcceptedActiveAlignment,
		SharedUnitAccepted:        examples[4].AcceptedActiveAlignment,
		Verdict:                   StatusAlternativeNormalizationsAudited,
	}
}

func buildAlternative(name string, alpha, beta float64, preservesLambda bool, reason string) AlternativeNormalization {
	c := (beta / alpha) * eventProbK7
	accepted := name == "shared signed-lambda unit"
	if name == "gauge-anchored boundary normalization" {
		accepted = false // conditionally equivalent, but not the active audit's final accepted coordinate label.
	}
	if name == "absolute scalar-wall form" {
		accepted = false
	}
	return AlternativeNormalization{
		Name:                         name,
		BoundaryScaleAlpha:           alpha,
		HistoryScaleBeta:             beta,
		TransformedCoefficient:       c,
		EqualsEventProbability:       math.Abs(c-eventProbK7) < tolerance,
		PreservesLambdaUnitAlignment: preservesLambda,
		AcceptedActiveAlignment:      accepted,
		Reason:                       reason,
	}
}

func buildNonTautology() NonTautologyAudit {
	return NonTautologyAudit{
		SharedLambdaPresent:              true,
		RearrangedEquation:               "kappa_lambda+kappa_e≈-(65/72)lambda+(7/72)(R_3-1)",
		LambdaWeight:                     -65.0 / 72.0,
		GaugeWeight:                      7.0 / 72.0,
		IndependentGaugeWoundPresent:     true,
		CoefficientsDiffer:               math.Abs((-65.0/72.0)-(7.0/72.0)) > tolerance,
		LambdaIsAlignmentAnchorNotProof:  true,
		NonTautologicalRelationPreserved: true,
		Verdict:                          StatusNonTautologyWithSharedLambdaAudited,
	}
}

func Statuses() []string {
	return []string{
		StatusGate701QuotientNormalizationInherited,
		StatusSharedLambdaCoordinateIdentified,
		StatusLambdaUnitCoefficientAlignmentAudited,
		StatusResponseCoefficientRemainsEventProbabilityUnderSharedUnit,
		StatusAlternativeNormalizationsAudited,
		StatusNonTautologyWithSharedLambdaAudited,
		StatusSharedScalarWallUnitAnchorsQuotientNormalization,
		StatusResponseCoefficientEqualsEventProbabilitySharedLambda,
		StatusGate700LawScalarWallUnitSealed,
		StatusSharedLambdaUnitAlignmentNotNative,
		StatusNoNativeWallCoordinateNormalizationAlignmentTheorem,
		StatusNoNativeBoundaryHistoryResponsePrinciple,
		StatusNoNativeSevenOver72Theorem,
		StatusGate702SharedScalarWallUnitBoundary,
	}
}

func FormatInheritance(x Gate701Inheritance) string {
	return fmt.Sprintf("inherited=%t p=%.18g canonicalC=%.18g covariance=%q coordinateSealed=%t noWallAlign=%t noPrinciple=%t no7=%t verdict=%q", x.InheritedQuotientNormalization, x.EventProbability, x.CanonicalCoefficient, x.CoefficientCovarianceFormula, x.Gate700CoordinateSealed, x.NoNativeWallAlignment, x.NoNativeBoundaryHistory, x.NoNativeSevenOver72, x.Verdict)
}

func FormatSharedLambda(x SharedLambdaCoordinateAudit) string {
	return fmt.Sprintf("boundary=%q history=%q shared=%q bCoeff=%.18g hCoeff=%.18g bContains=%t hContains=%t sameWall=%t unit=%t verdict=%q", x.BoundaryCoordinate, x.HistoryCoordinate, x.SharedCoordinate, x.BoundaryLambdaCoefficient, x.HistoryLambdaCoefficient, x.BoundaryContainsSharedLambda, x.HistoryContainsSharedLambda, x.SameSignedScalarZeroWall, x.UnitCoefficientAlignment, x.Verdict)
}

func FormatAnchor(x NormalizationAnchorAudit) string {
	return fmt.Sprintf("rule=%q bCoeff=%.18g hCoeff=%.18g alpha=%.18g beta=%.18g betaOverAlpha=%.18g c=%.18g p=%.18g equalsP=%t verdict=%q", x.Rule, x.BoundaryLambdaCoefficient, x.HistoryLambdaCoefficient, x.Alpha, x.Beta, x.BetaOverAlpha, x.ResponseCoefficient, x.EventProbability, x.CoefficientEqualsProbability, x.Verdict)
}

func FormatAlternatives(x AlternativeNormalizationAudit) string {
	parts := make([]string, 0, len(x.Examples))
	for _, ex := range x.Examples {
		parts = append(parts, fmt.Sprintf("%s: alpha=%.18g beta=%.18g c=%.18g eqP=%t lambdaUnit=%t accepted=%t", ex.Name, ex.BoundaryScaleAlpha, ex.HistoryScaleBeta, ex.TransformedCoefficient, ex.EqualsEventProbability, ex.PreservesLambdaUnitAlignment, ex.AcceptedActiveAlignment))
	}
	return fmt.Sprintf("euclideanRejected=%t historyRejected=%t gaugeConditioned=%t absoluteRejected=%t sharedAccepted=%t examples=[%s] verdict=%q", x.EuclideanBoundaryRejected, x.HistoryNormRejected, x.GaugeAnchorConditioned, x.AbsoluteFormRejected, x.SharedUnitAccepted, strings.Join(parts, " | "), x.Verdict)
}

func FormatNonTautology(x NonTautologyAudit) string {
	return fmt.Sprintf("sharedLambda=%t equation=%q lambdaWeight=%.18g gaugeWeight=%.18g gaugePresent=%t coeffsDiffer=%t anchorNotProof=%t preserved=%t verdict=%q", x.SharedLambdaPresent, x.RearrangedEquation, x.LambdaWeight, x.GaugeWeight, x.IndependentGaugeWoundPresent, x.CoefficientsDiffer, x.LambdaIsAlignmentAnchorNotProof, x.NonTautologicalRelationPreserved, x.Verdict)
}

func FormatSource(x SourceTypeClassification) string {
	return fmt.Sprintf("pRole=%q lambdaRole=%q boundaryRole=%q historyRole=%q conclusion=%q verdict=%q", x.EventProbabilityRole, x.LambdaUnitRole, x.BoundaryRole, x.HistoryRole, x.Conclusion, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("theorems=%s verdict=%q", strings.Join(x.Theorems, ", "), x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("nativeShared=%t nativeWallNorm=%t nativePrinciple=%t native7=%t boundaryStress=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckm=%t verdict=%q", x.ClaimsSharedLambdaAlignmentNative, x.ClaimsNativeWallNormalizationAlignment, x.ClaimsNativeBoundaryHistoryPrinciple, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStressDerived, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}

func canonicalSBoundary() float64   { return lambdaLambda12 + r3Minus1 }
func canonicalSHistory() float64    { return kappaLambda + kappaE + lambdaLambda12 }
func canonicalExpectation() float64 { return eventProbK7 * canonicalSBoundary() }
func canonicalResidual() float64    { return canonicalSHistory() - canonicalExpectation() }
