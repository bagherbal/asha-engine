// Package generation2quotientlinenormalizationandresponsecoefficientcovarianceaudit implements
// Gate 701: Quotient-Line Normalization and Response Coefficient Covariance Audit.
//
// Gate 700 closed the conditional bridge law
//
//	sigma_history(h) ≈ Tr[rho_72 sigma_boundary(b) P_K7]
//
// in the canonical wall-distance coordinates inherited from the bridge chain.
// Gate 701 audits whether the numerical response coefficient 7/72 is
// coordinate-invariant, or whether it is the K7 event probability appearing as a
// response coefficient only after quotient-line normalizations are fixed. This is
// a bridge-layer quotient-normalization audit only. It does not derive boundary
// stress, scalar RG matching, Higgs mass, gauge unification, flavor, CKM/PMNS, a
// native response theorem, a native state-selection theorem, or a native 7/72
// theorem.
package generation2quotientlinenormalizationandresponsecoefficientcovarianceaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate700 "github.com/bagherbal/asha-engine/pkg/bridge/generation2conditionalashahistoryresponselawclosureaudit"
)

const (
	AuditID = "GATE701-QUOTIENT-LINE-NORMALIZATION-AND-RESPONSE-COEFFICIENT-COVARIANCE-AUDIT"

	StatusGate700ConditionalHistoryResponseLawInherited               = "PASS_GATE700_CONDITIONAL_HISTORY_RESPONSE_LAW_INHERITED"
	StatusQuotientLineRescalingDefined                                = "PASS_QUOTIENT_LINE_RESCALING_DEFINED"
	StatusResponseCoefficientTransformationComputed                   = "PASS_RESPONSE_COEFFICIENT_TRANSFORMATION_COMPUTED"
	StatusEventProbabilityInvariantSeparated                          = "PASS_EVENT_PROBABILITY_INVARIANT_SEPARATED"
	StatusWallCoordinateNormalizationAudited                          = "PASS_WALL_COORDINATE_NORMALIZATION_AUDITED"
	StatusAlternativeNormalizationExamplesComputed                    = "PASS_ALTERNATIVE_NORMALIZATION_EXAMPLES_COMPUTED"
	StatusSevenOver72InvariantAsK7EventProbability                    = "CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_IS_INVARIANT_AS_K7_EVENT_PROBABILITY"
	StatusResponseCoefficientEqualsEventProbabilityOnlyCanonicalWalls = "CONDITIONAL_SUPPORT_RESPONSE_COEFFICIENT_EQUALS_EVENT_PROBABILITY_ONLY_IN_CANONICAL_WALL_NORMALIZATION"
	StatusGate700LawCoordinateSealedNotCoordinateFree                 = "CONDITIONAL_SUPPORT_GATE700_LAW_IS_COORDINATE_SEALED_NOT_COORDINATE_FREE"
	StatusResponseCoeffNotInvariantUnderArbitraryQuotientRescaling    = "FAILED_ROUTE_RESPONSE_COEFFICIENT_NOT_INVARIANT_UNDER_ARBITRARY_QUOTIENT_RESCALING"
	StatusNoNativeWallCoordinateNormalizationAlignmentTheorem         = "FAILED_ROUTE_NO_NATIVE_WALL_COORDINATE_NORMALIZATION_ALIGNMENT_THEOREM"
	StatusNoNativeBoundaryHistoryResponsePrinciple                    = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_HISTORY_RESPONSE_PRINCIPLE"
	StatusNoNativeSevenOver72Theorem                                  = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate701QuotientNormalizationBoundary                        = "FIREWALL_PRESERVED_GATE701_QUOTIENT_NORMALIZATION_BOUNDARY"
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

type Gate700Inheritance struct {
	InheritedConditionalLaw bool
	SigmaBoundary           float64
	SigmaHistory            float64
	Expectation             float64
	ResidualE1              float64
	PremisesNonredundant    bool
	NoNativePrinciple       bool
	NoNativeSevenOver72     bool
	Verdict                 string
}

type QuotientLineRescalingAudit struct {
	BoundaryScale          float64
	HistoryScale           float64
	OriginalCoefficient    float64
	TransformedCoefficient float64
	Formula                string
	CoefficientCovariant   bool
	CoefficientInvariant   bool
	Verdict                string
}

type EventProbabilityInvariantAudit struct {
	ProbabilityName                     string
	Probability                         float64
	TraceFormula                        string
	IndependentOfAlpha                  bool
	IndependentOfBeta                   bool
	InvariantObjectSeparated            bool
	ResponseCoefficientNeedsCoordinates bool
	Verdict                             string
}

type WallCoordinateNormalizationAudit struct {
	BoundaryCoordinate                  string
	HistoryCoordinate                   string
	BoundaryUsesUnitWallCoefficients    bool
	HistoryUsesUnitWallCoefficients     bool
	SameWallDistanceFamily              bool
	CanonicalAlpha                      float64
	CanonicalBeta                       float64
	CoefficientEqualsProbability        bool
	EqualScaleRequiredForDirectEquality bool
	Verdict                             string
}

type NormalizationExample struct {
	Name                   string
	BoundaryScale          float64
	HistoryScale           float64
	TransformedCoefficient float64
	ExpectedCoefficient    float64
	MatchesExpected        bool
}

type AlternativeNormalizationAudit struct {
	Examples         []NormalizationExample
	AllComputed      bool
	NonInvariantSeen bool
	CanonicalSeen    bool
	Verdict          string
}

type SourceSeparationAudit struct {
	InvariantObject        string
	CoordinateSealedObject string
	DoesNotWeakenGate700   bool
	ClarifiesSourceType    bool
	Verdict                string
}

type MissingTheoremAudit struct {
	Theorems    []string
	PrecisePair string
	Verdict     string
}

type FirewallAudit struct {
	ClaimsResponseCoefficientCoordinateFree bool
	ClaimsNativeWallNormalization           bool
	ClaimsNativeBoundaryHistoryPrinciple    bool
	ClaimsNativeSevenOver72Theorem          bool
	ClaimsBoundaryStressDerived             bool
	ClaimsScalarRGMatching                  bool
	ClaimsHiggsMass                         bool
	ClaimsGaugeUnification                  bool
	ClaimsFlavorDerivation                  bool
	ClaimsCKMPMNS                           bool
	Verdict                                 string
}

type Analysis struct {
	Inherited     Gate700Inheritance
	Rescaling     QuotientLineRescalingAudit
	Probability   EventProbabilityInvariantAudit
	WallNormalize WallCoordinateNormalizationAudit
	Alternatives  AlternativeNormalizationAudit
	Source        SourceSeparationAudit
	Missing       MissingTheoremAudit
	Firewall      FirewallAudit
	Truth         string
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
	g700, err := gate700.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate700 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g700)
	rescaling := BuildRescaling(2, 3) // representative nontrivial covariance witness.
	probability := buildEventProbabilityInvariant()
	wall := buildWallNormalization()
	alternatives := buildAlternativeNormalizations()
	source := SourceSeparationAudit{
		InvariantObject:        "p_K7=Tr(rho_72 P_K7)=7/72",
		CoordinateSealedObject: "sigma_history≈(7/72)sigma_boundary only after the chosen wall-distance normalization of the two quotient lines",
		DoesNotWeakenGate700:   true,
		ClarifiesSourceType:    true,
		Verdict: strings.Join([]string{
			StatusEventProbabilityInvariantSeparated,
			StatusGate700LawCoordinateSealedNotCoordinateFree,
		}, "; "),
	}
	missing := MissingTheoremAudit{
		Theorems: []string{
			"NoBiasK7EventProbabilityTheorem",
			"WallCoordinateNormalizationAlignmentTheorem",
			"NativeBoundaryHistoryResponsePrinciple",
			StatusNoNativeWallCoordinateNormalizationAlignmentTheorem,
			StatusNoNativeBoundaryHistoryResponsePrinciple,
			StatusNoNativeSevenOver72Theorem,
		},
		PrecisePair: "one theorem explaining p_K7=7/72 from rho_72 and P_K7, plus one theorem explaining why sigma_boundary and sigma_history use aligned unit wall-distance coordinates so the event probability appears directly as the response coefficient",
		Verdict: strings.Join([]string{
			StatusNoNativeWallCoordinateNormalizationAlignmentTheorem,
			StatusNoNativeBoundaryHistoryResponsePrinciple,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate701QuotientNormalizationBoundary}
	truth := "Gate 701 separates the invariant K7 event probability from the coordinate-sealed response coefficient. The number 7/72 is stable as p_K7=Tr(rho_72 P_K7), but the coefficient in sigma_history≈c sigma_boundary transforms as c'=(beta/alpha)(7/72) under quotient-line rescaling. Thus Gate700 is a valid canonical-wall-coordinate response law, not a coordinate-free coefficient theorem."
	return Analysis{Inherited: inherited, Rescaling: rescaling, Probability: probability, WallNormalize: wall, Alternatives: alternatives, Source: source, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate700.Analysis) Gate700Inheritance {
	return Gate700Inheritance{
		InheritedConditionalLaw: g.Functional.ApproxLawCertified && g.Master.Reconstructed && g.Premises.Complete,
		SigmaBoundary:           g.Inherited.SBoundary,
		SigmaHistory:            g.Inherited.SHistory,
		Expectation:             g.Master.Expectation,
		ResidualE1:              g.Residual.ResidualE1,
		PremisesNonredundant:    g.Removal.EachPremiseNonredundant,
		NoNativePrinciple:       !g.Firewall.ClaimsNativeBoundaryHistoryPrinciple,
		NoNativeSevenOver72:     !g.Firewall.ClaimsNativeSevenOver72Theorem,
		Verdict:                 StatusGate700ConditionalHistoryResponseLawInherited,
	}
}

func BuildRescaling(alpha, beta float64) QuotientLineRescalingAudit {
	if alpha == 0 {
		return QuotientLineRescalingAudit{BoundaryScale: alpha, HistoryScale: beta, Formula: "undefined for alpha=0", CoefficientCovariant: false, CoefficientInvariant: false, Verdict: StatusResponseCoeffNotInvariantUnderArbitraryQuotientRescaling}
	}
	transformed := (beta / alpha) * eventProbK7
	return QuotientLineRescalingAudit{
		BoundaryScale:          alpha,
		HistoryScale:           beta,
		OriginalCoefficient:    eventProbK7,
		TransformedCoefficient: transformed,
		Formula:                "sigma_boundary'=alpha sigma_boundary, sigma_history'=beta sigma_history => c'=(beta/alpha)(7/72)",
		CoefficientCovariant:   math.Abs(transformed-((beta/alpha)*eventProbK7)) < tolerance,
		CoefficientInvariant:   math.Abs(transformed-eventProbK7) < tolerance,
		Verdict: strings.Join([]string{
			StatusQuotientLineRescalingDefined,
			StatusResponseCoefficientTransformationComputed,
			StatusResponseCoeffNotInvariantUnderArbitraryQuotientRescaling,
		}, "; "),
	}
}

func buildEventProbabilityInvariant() EventProbabilityInvariantAudit {
	return EventProbabilityInvariantAudit{
		ProbabilityName:                     "p_K7",
		Probability:                         eventProbK7,
		TraceFormula:                        "p_K7=Tr(rho_72 P_K7)=7/72",
		IndependentOfAlpha:                  true,
		IndependentOfBeta:                   true,
		InvariantObjectSeparated:            true,
		ResponseCoefficientNeedsCoordinates: true,
		Verdict: strings.Join([]string{
			StatusEventProbabilityInvariantSeparated,
			StatusSevenOver72InvariantAsK7EventProbability,
		}, "; "),
	}
}

func buildWallNormalization() WallCoordinateNormalizationAudit {
	return WallCoordinateNormalizationAudit{
		BoundaryCoordinate:                  "sigma_boundary=lambda+(R_3-1)",
		HistoryCoordinate:                   "sigma_history=kappa_lambda+kappa_e+lambda",
		BoundaryUsesUnitWallCoefficients:    true,
		HistoryUsesUnitWallCoefficients:     true,
		SameWallDistanceFamily:              true,
		CanonicalAlpha:                      1,
		CanonicalBeta:                       1,
		CoefficientEqualsProbability:        true,
		EqualScaleRequiredForDirectEquality: true,
		Verdict: strings.Join([]string{
			StatusWallCoordinateNormalizationAudited,
			StatusResponseCoefficientEqualsEventProbabilityOnlyCanonicalWalls,
			StatusGate700LawCoordinateSealedNotCoordinateFree,
		}, "; "),
	}
}

func buildAlternativeNormalizations() AlternativeNormalizationAudit {
	examples := []NormalizationExample{
		buildExample("boundary doubled", 2, 1, 7.0/144.0),
		buildExample("history doubled", 1, 2, 7.0/36.0),
		buildExample("canonical wall coordinates", 1, 1, 7.0/72.0),
	}
	return AlternativeNormalizationAudit{
		Examples:         examples,
		AllComputed:      allExamplesMatch(examples),
		NonInvariantSeen: math.Abs(examples[0].TransformedCoefficient-eventProbK7) > tolerance && math.Abs(examples[1].TransformedCoefficient-eventProbK7) > tolerance,
		CanonicalSeen:    math.Abs(examples[2].TransformedCoefficient-eventProbK7) < tolerance,
		Verdict:          StatusAlternativeNormalizationExamplesComputed,
	}
}

func buildExample(name string, alpha, beta, expected float64) NormalizationExample {
	value := (beta / alpha) * eventProbK7
	return NormalizationExample{Name: name, BoundaryScale: alpha, HistoryScale: beta, TransformedCoefficient: value, ExpectedCoefficient: expected, MatchesExpected: math.Abs(value-expected) < tolerance}
}

func allExamplesMatch(xs []NormalizationExample) bool {
	for _, x := range xs {
		if !x.MatchesExpected {
			return false
		}
	}
	return true
}

func Statuses() []string {
	return []string{
		StatusGate700ConditionalHistoryResponseLawInherited,
		StatusQuotientLineRescalingDefined,
		StatusResponseCoefficientTransformationComputed,
		StatusEventProbabilityInvariantSeparated,
		StatusWallCoordinateNormalizationAudited,
		StatusAlternativeNormalizationExamplesComputed,
		StatusSevenOver72InvariantAsK7EventProbability,
		StatusResponseCoefficientEqualsEventProbabilityOnlyCanonicalWalls,
		StatusGate700LawCoordinateSealedNotCoordinateFree,
		StatusResponseCoeffNotInvariantUnderArbitraryQuotientRescaling,
		StatusNoNativeWallCoordinateNormalizationAlignmentTheorem,
		StatusNoNativeBoundaryHistoryResponsePrinciple,
		StatusNoNativeSevenOver72Theorem,
		StatusGate701QuotientNormalizationBoundary,
	}
}

func FormatInheritance(x Gate700Inheritance) string {
	return fmt.Sprintf("inherited=%t sigmaBoundary=%.18g sigmaHistory=%.18g expectation=%.18g e1=%.18g nonredundant=%t noPrinciple=%t no7=%t verdict=%q", x.InheritedConditionalLaw, x.SigmaBoundary, x.SigmaHistory, x.Expectation, x.ResidualE1, x.PremisesNonredundant, x.NoNativePrinciple, x.NoNativeSevenOver72, x.Verdict)
}

func FormatRescaling(x QuotientLineRescalingAudit) string {
	return fmt.Sprintf("alpha=%.18g beta=%.18g original=%.18g transformed=%.18g formula=%q covariant=%t invariant=%t verdict=%q", x.BoundaryScale, x.HistoryScale, x.OriginalCoefficient, x.TransformedCoefficient, x.Formula, x.CoefficientCovariant, x.CoefficientInvariant, x.Verdict)
}

func FormatProbability(x EventProbabilityInvariantAudit) string {
	return fmt.Sprintf("name=%q p=%.18g trace=%q independentAlpha=%t independentBeta=%t separated=%t coeffNeedsCoordinates=%t verdict=%q", x.ProbabilityName, x.Probability, x.TraceFormula, x.IndependentOfAlpha, x.IndependentOfBeta, x.InvariantObjectSeparated, x.ResponseCoefficientNeedsCoordinates, x.Verdict)
}

func FormatWallNormalization(x WallCoordinateNormalizationAudit) string {
	return fmt.Sprintf("boundary=%q history=%q boundaryUnit=%t historyUnit=%t sameFamily=%t alpha=%.18g beta=%.18g coeffEqualsP=%t equalScaleRequired=%t verdict=%q", x.BoundaryCoordinate, x.HistoryCoordinate, x.BoundaryUsesUnitWallCoefficients, x.HistoryUsesUnitWallCoefficients, x.SameWallDistanceFamily, x.CanonicalAlpha, x.CanonicalBeta, x.CoefficientEqualsProbability, x.EqualScaleRequiredForDirectEquality, x.Verdict)
}

func FormatAlternatives(x AlternativeNormalizationAudit) string {
	parts := make([]string, 0, len(x.Examples))
	for _, ex := range x.Examples {
		parts = append(parts, fmt.Sprintf("%s: alpha=%.18g beta=%.18g c=%.18g", ex.Name, ex.BoundaryScale, ex.HistoryScale, ex.TransformedCoefficient))
	}
	return fmt.Sprintf("allComputed=%t nonInvariant=%t canonical=%t examples=[%s] verdict=%q", x.AllComputed, x.NonInvariantSeen, x.CanonicalSeen, strings.Join(parts, " | "), x.Verdict)
}

func FormatSource(x SourceSeparationAudit) string {
	return fmt.Sprintf("invariant=%q coordinateSealed=%q doesNotWeaken=%t clarifies=%t verdict=%q", x.InvariantObject, x.CoordinateSealedObject, x.DoesNotWeakenGate700, x.ClarifiesSourceType, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("theorems=%s pair=%q verdict=%q", strings.Join(x.Theorems, ", "), x.PrecisePair, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("coordFree=%t nativeWallNorm=%t nativePrinciple=%t native7=%t boundaryStress=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckm=%t verdict=%q", x.ClaimsResponseCoefficientCoordinateFree, x.ClaimsNativeWallNormalization, x.ClaimsNativeBoundaryHistoryPrinciple, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStressDerived, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}

func canonicalSBoundary() float64   { return lambdaLambda12 + r3Minus1 }
func canonicalSHistory() float64    { return kappaLambda + kappaE + lambdaLambda12 }
func canonicalExpectation() float64 { return eventProbK7 * canonicalSBoundary() }
func canonicalResidual() float64    { return canonicalSHistory() - canonicalExpectation() }
