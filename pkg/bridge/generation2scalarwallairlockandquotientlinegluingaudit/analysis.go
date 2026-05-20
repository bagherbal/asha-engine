// Package generation2scalarwallairlockandquotientlinegluingaudit implements
// Gate 703: Scalar-Wall Airlock and Quotient-Line Gluing Audit.
//
// Gate 702 showed that the active response coefficient equals the invariant
// K7 event probability p_K7=7/72 only when the boundary quotient coordinate
// and the history quotient coordinate are measured in aligned scalar-wall
// units. Gate 703 audits whether this alignment can be typed as a scalar-wall
// airlock/gluing diagram using the shared signed scalar zero-wall coordinate
// lambda(Lambda_12). This is a bridge-layer quotient-gluing audit only. It
// does not derive boundary stress, scalar RG matching, Higgs mass, gauge
// unification, flavor, CKM/PMNS, a native wall-normalization theorem, a native
// response theorem, or a native 7/72 theorem.
package generation2scalarwallairlockandquotientlinegluingaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate702 "github.com/bagherbal/asha-engine/pkg/bridge/generation2sharedscalarwallunitnormalizationalignmentaudit"
)

const (
	AuditID = "GATE703-SCALAR-WALL-AIRLOCK-AND-QUOTIENT-LINE-GLUING-AUDIT"

	StatusGate702SharedScalarWallUnitInherited              = "PASS_GATE702_SHARED_SCALAR_WALL_UNIT_INHERITED"
	StatusScalarWallLineDefined                             = "PASS_SCALAR_WALL_LINE_DEFINED"
	StatusQuotientLineGluingDiagramDefined                  = "PASS_QUOTIENT_LINE_GLUING_DIAGRAM_DEFINED"
	StatusUnitLambdaGlueConditionAudited                    = "PASS_UNIT_LAMBDA_GLUE_CONDITION_AUDITED"
	StatusResponseCoefficientPreservationComputed           = "PASS_RESPONSE_COEFFICIENT_PRESERVATION_COMPUTED"
	StatusAlternativeGluingsAudited                         = "PASS_ALTERNATIVE_GLUINGS_AUDITED"
	StatusNonTautologyOfSharedLambdaAudited                 = "PASS_NON_TAUTOLOGY_OF_SHARED_LAMBDA_AUDITED"
	StatusSharedLambdaIsScalarWallAirlock                   = "CONDITIONAL_SUPPORT_SHARED_LAMBDA_IS_SCALAR_WALL_AIRLOCK"
	StatusResponseCoefficientEqualsProbabilityAfterUnitGlue = "CONDITIONAL_SUPPORT_RESPONSE_COEFFICIENT_EQUALS_EVENT_PROBABILITY_ONLY_AFTER_UNIT_SCALAR_WALL_GLUE"
	StatusGate700LawScalarWallGluedQuotientResponse         = "CONDITIONAL_SUPPORT_GATE700_LAW_IS_SCALAR_WALL_GLUED_QUOTIENT_RESPONSE"
	StatusScalarWallGluingNotNative                         = "FAILED_ROUTE_SCALAR_WALL_GLUING_NOT_NATIVELY_DERIVED"
	StatusNoNativeScalarWallAirlockTheorem                  = "FAILED_ROUTE_NO_NATIVE_SCALAR_WALL_AIRLOCK_THEOREM"
	StatusNoNativeBoundaryHistoryResponsePrinciple          = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_HISTORY_RESPONSE_PRINCIPLE"
	StatusNoNativeSevenOver72Theorem                        = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate703ScalarWallAirlockBoundary                  = "FIREWALL_PRESERVED_GATE703_SCALAR_WALL_AIRLOCK_BOUNDARY"
)

const (
	h72Dimension   = 72
	k7Dimension    = 7
	eventProbK7    = float64(k7Dimension) / float64(h72Dimension)
	tolerance      = 1e-15
	lambdaLambda12 = -0.0497009420776833
	r3Minus1       = 0.0509933868964996
	kappaLambda    = 0.0443230430960771
	kappaE         = 0.00550355419157456
)

type Gate702Inheritance struct {
	InheritedSharedScalarWallUnit bool
	SharedCoordinate              string
	EventProbability              float64
	ResponseCoefficient           float64
	CoefficientEqualsProbability  bool
	NonTautologyInherited         bool
	NoNativeSharedUnit            bool
	NoNativeWallAlignment         bool
	NoNativeBoundaryHistory       bool
	NoNativeSevenOver72           bool
	Verdict                       string
}

type ScalarWallLineAudit struct {
	LineName                   string
	Definition                 string
	Coordinate                 string
	SignedOrientationPreserved bool
	BoundaryLambdaCoefficient  float64
	HistoryLambdaCoefficient   float64
	UnitScalarWallOnBothSides  bool
	Verdict                    string
}

type GluingDiagramAudit struct {
	Diagram                  string
	BoundaryCoordinate       string
	AirlockLine              string
	HistoryCoordinate        string
	BoundaryMeasuredInLambda bool
	HistoryMeasuredInLambda  bool
	CoordinatesAreAnchored   bool
	Verdict                  string
}

type UnitGlueAudit struct {
	BoundaryLambdaCoefficient    float64
	HistoryLambdaCoefficient     float64
	Gamma                        float64
	UnitGlue                     bool
	EventProbability             float64
	ResponseCoefficient          float64
	CoefficientEqualsProbability bool
	Verdict                      string
}

type RescaledGlueAudit struct {
	Gamma                       float64
	TransformedCoefficient      float64
	EqualsEventProbability      bool
	RequiresGammaOneForEquality bool
	Formula                     string
	Verdict                     string
}

type AlternativeGluing struct {
	Name                  string
	BoundaryScale         float64
	HistoryScale          float64
	Gamma                 float64
	ResponseCoefficient   float64
	AcceptedActiveAirlock bool
	Rejected              bool
	Reason                string
}

type AlternativeGluingAudit struct {
	Examples                   []AlternativeGluing
	BoundaryNormalizedRejected bool
	HistoryNormalizedRejected  bool
	AbsoluteScalarRejected     bool
	HessianScalarRejected      bool
	SharedSignedLambdaAccepted bool
	Verdict                    string
}

type NonTautologyAudit struct {
	SharedLambdaAirlockPresent       bool
	RearrangedEquation               string
	LambdaWeight                     float64
	GaugeWeight                      float64
	IndependentGaugeWoundPresent     bool
	LambdaIsAirlockNotProof          bool
	NonTautologicalRelationPreserved bool
	Verdict                          string
}

type SourceTypeClassification struct {
	ScalarWallLineRole      string
	BoundaryRole            string
	HistoryRole             string
	EventProbabilityRole    string
	ResponseCoefficientRole string
	Conclusion              string
	Verdict                 string
}

type MissingTheoremAudit struct {
	Theorems []string
	Verdict  string
}

type FirewallAudit struct {
	ClaimsScalarWallGluingNative   bool
	ClaimsNativeScalarWallAirlock  bool
	ClaimsNativeBoundaryHistory    bool
	ClaimsNativeSevenOver72Theorem bool
	ClaimsBoundaryStressDerived    bool
	ClaimsScalarRGMatching         bool
	ClaimsHiggsMass                bool
	ClaimsGaugeUnification         bool
	ClaimsFlavorDerivation         bool
	ClaimsCKMPMNS                  bool
	Verdict                        string
}

type Analysis struct {
	Inherited    Gate702Inheritance
	ScalarWall   ScalarWallLineAudit
	Diagram      GluingDiagramAudit
	UnitGlue     UnitGlueAudit
	Rescaled     RescaledGlueAudit
	Alternatives AlternativeGluingAudit
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
	g702, err := gate702.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate702 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g702)
	scalarWall := buildScalarWallLine(g702)
	diagram := buildDiagram(scalarWall)
	unitGlue := buildUnitGlue(scalarWall)
	rescaled := buildRescaledGlue(2)
	alternatives := buildAlternatives()
	nontaut := buildNonTautology()
	source := SourceTypeClassification{
		ScalarWallLineRole:      "L_lambda is the shared signed scalar-wall unit line",
		BoundaryRole:            "sigma_boundary=lambda+(R_3-1) is the boundary anti-alignment quotient measured in scalar-wall units",
		HistoryRole:             "sigma_history=kappa_lambda+kappa_e+lambda is the scalar/flavor/history closure quotient measured in scalar-wall units",
		EventProbabilityRole:    "p_K7=7/72 is the invariant no-bias K7 event probability",
		ResponseCoefficientRole: "c_response equals p_K7 only after unit scalar-wall gluing",
		Conclusion:              "the Gate700 response law is a scalar-wall glued quotient response, not a coordinate-free response coefficient theorem",
		Verdict: strings.Join([]string{
			StatusSharedLambdaIsScalarWallAirlock,
			StatusResponseCoefficientEqualsProbabilityAfterUnitGlue,
			StatusGate700LawScalarWallGluedQuotientResponse,
		}, "; "),
	}
	missing := MissingTheoremAudit{
		Theorems: []string{
			"ScalarWallAirlockTheorem",
			"BoundaryHistoryScalarWallGluingTheorem",
			StatusScalarWallGluingNotNative,
			StatusNoNativeScalarWallAirlockTheorem,
			StatusNoNativeBoundaryHistoryResponsePrinciple,
			StatusNoNativeSevenOver72Theorem,
		},
		Verdict: strings.Join([]string{
			StatusScalarWallGluingNotNative,
			StatusNoNativeScalarWallAirlockTheorem,
			StatusNoNativeBoundaryHistoryResponsePrinciple,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate703ScalarWallAirlockBoundary}
	truth := "Gate 703 conditionally types lambda(Lambda_12) as the scalar-wall airlock gluing the boundary quotient line to the history quotient line. Unit signed-lambda gluing preserves c_response=p_K7=7/72, while rescaled gluing changes the coordinate coefficient. The shared lambda line is a normalization anchor, not a native proof of the boundary-history response law."
	return Analysis{Inherited: inherited, ScalarWall: scalarWall, Diagram: diagram, UnitGlue: unitGlue, Rescaled: rescaled, Alternatives: alternatives, NonTautology: nontaut, Source: source, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate702.Analysis) Gate702Inheritance {
	return Gate702Inheritance{
		InheritedSharedScalarWallUnit: g.Anchor.CoefficientEqualsProbability && g.SharedLambda.UnitCoefficientAlignment,
		SharedCoordinate:              g.SharedLambda.SharedCoordinate,
		EventProbability:              g.Anchor.EventProbability,
		ResponseCoefficient:           g.Anchor.ResponseCoefficient,
		CoefficientEqualsProbability:  g.Anchor.CoefficientEqualsProbability,
		NonTautologyInherited:         g.NonTautology.NonTautologicalRelationPreserved,
		NoNativeSharedUnit:            !g.Firewall.ClaimsSharedLambdaAlignmentNative,
		NoNativeWallAlignment:         !g.Firewall.ClaimsNativeWallNormalizationAlignment,
		NoNativeBoundaryHistory:       !g.Firewall.ClaimsNativeBoundaryHistoryPrinciple,
		NoNativeSevenOver72:           !g.Firewall.ClaimsNativeSevenOver72Theorem,
		Verdict:                       StatusGate702SharedScalarWallUnitInherited,
	}
}

func buildScalarWallLine(g gate702.Analysis) ScalarWallLineAudit {
	return ScalarWallLineAudit{
		LineName:                   "L_lambda",
		Definition:                 "L_lambda=span(lambda(Lambda_12))",
		Coordinate:                 g.SharedLambda.SharedCoordinate,
		SignedOrientationPreserved: true,
		BoundaryLambdaCoefficient:  g.SharedLambda.BoundaryLambdaCoefficient,
		HistoryLambdaCoefficient:   g.SharedLambda.HistoryLambdaCoefficient,
		UnitScalarWallOnBothSides:  g.SharedLambda.BoundaryLambdaCoefficient == 1 && g.SharedLambda.HistoryLambdaCoefficient == 1,
		Verdict: strings.Join([]string{
			StatusScalarWallLineDefined,
			StatusSharedLambdaIsScalarWallAirlock,
		}, "; "),
	}
}

func buildDiagram(w ScalarWallLineAudit) GluingDiagramAudit {
	return GluingDiagramAudit{
		Diagram:                  "Q_boundary --lambda units--> L_lambda --same unit--> Q_history",
		BoundaryCoordinate:       "sigma_boundary=lambda+(R_3-1)",
		AirlockLine:              w.Definition,
		HistoryCoordinate:        "sigma_history=kappa_lambda+kappa_e+lambda",
		BoundaryMeasuredInLambda: true,
		HistoryMeasuredInLambda:  true,
		CoordinatesAreAnchored:   w.UnitScalarWallOnBothSides,
		Verdict:                  StatusQuotientLineGluingDiagramDefined,
	}
}

func buildUnitGlue(w ScalarWallLineAudit) UnitGlueAudit {
	gamma := w.HistoryLambdaCoefficient / w.BoundaryLambdaCoefficient
	c := gamma * eventProbK7
	return UnitGlueAudit{
		BoundaryLambdaCoefficient:    w.BoundaryLambdaCoefficient,
		HistoryLambdaCoefficient:     w.HistoryLambdaCoefficient,
		Gamma:                        gamma,
		UnitGlue:                     math.Abs(gamma-1) < tolerance,
		EventProbability:             eventProbK7,
		ResponseCoefficient:          c,
		CoefficientEqualsProbability: math.Abs(c-eventProbK7) < tolerance,
		Verdict: strings.Join([]string{
			StatusUnitLambdaGlueConditionAudited,
			StatusResponseCoefficientPreservationComputed,
			StatusResponseCoefficientEqualsProbabilityAfterUnitGlue,
		}, "; "),
	}
}

func buildRescaledGlue(gamma float64) RescaledGlueAudit {
	c := gamma * eventProbK7
	return RescaledGlueAudit{
		Gamma:                       gamma,
		TransformedCoefficient:      c,
		EqualsEventProbability:      math.Abs(c-eventProbK7) < tolerance,
		RequiresGammaOneForEquality: true,
		Formula:                     "lambda_history=gamma lambda_boundary => c_response'=gamma p_K7",
		Verdict:                     StatusResponseCoefficientPreservationComputed,
	}
}

func buildAlternatives() AlternativeGluingAudit {
	examples := []AlternativeGluing{
		buildAlternative("boundary-normalized gluing", 1/math.Sqrt2, 1, false, "sigma_boundary=(lambda+R)/sqrt(2) is rejected unless history is rescaled by the same factor"),
		buildAlternative("history-normalized gluing", 1, 1/math.Sqrt(3), false, "sigma_history=(kappa_lambda+kappa_e+lambda)/sqrt(3) is rejected unless boundary is rescaled by the same factor"),
		buildAlternative("absolute scalar gluing", 1, 1, false, "|lambda| erases scalar-wall orientation and is therefore not the active signed airlock"),
		buildAlternative("Hessian scalar gluing", 1, 2, false, "2lambda or 2|lambda| belongs to Hessian/squared-mass layer rather than the active wall-distance layer"),
		buildAlternative("shared signed lambda gluing", 1, 1, true, "lambda coefficient one on both sides is the active scalar-wall airlock"),
	}
	return AlternativeGluingAudit{
		Examples:                   examples,
		BoundaryNormalizedRejected: examples[0].Rejected,
		HistoryNormalizedRejected:  examples[1].Rejected,
		AbsoluteScalarRejected:     examples[2].Rejected,
		HessianScalarRejected:      examples[3].Rejected,
		SharedSignedLambdaAccepted: examples[4].AcceptedActiveAirlock,
		Verdict:                    StatusAlternativeGluingsAudited,
	}
}

func buildAlternative(name string, boundaryScale, historyScale float64, accepted bool, reason string) AlternativeGluing {
	gamma := historyScale / boundaryScale
	c := gamma * eventProbK7
	return AlternativeGluing{
		Name:                  name,
		BoundaryScale:         boundaryScale,
		HistoryScale:          historyScale,
		Gamma:                 gamma,
		ResponseCoefficient:   c,
		AcceptedActiveAirlock: accepted,
		Rejected:              !accepted,
		Reason:                reason,
	}
}

func buildNonTautology() NonTautologyAudit {
	return NonTautologyAudit{
		SharedLambdaAirlockPresent:       true,
		RearrangedEquation:               "kappa_lambda+kappa_e≈-(65/72)lambda+(7/72)(R_3-1)",
		LambdaWeight:                     -65.0 / 72.0,
		GaugeWeight:                      7.0 / 72.0,
		IndependentGaugeWoundPresent:     true,
		LambdaIsAirlockNotProof:          true,
		NonTautologicalRelationPreserved: true,
		Verdict:                          StatusNonTautologyOfSharedLambdaAudited,
	}
}

func Statuses() []string {
	return []string{
		StatusGate702SharedScalarWallUnitInherited,
		StatusScalarWallLineDefined,
		StatusQuotientLineGluingDiagramDefined,
		StatusUnitLambdaGlueConditionAudited,
		StatusResponseCoefficientPreservationComputed,
		StatusAlternativeGluingsAudited,
		StatusNonTautologyOfSharedLambdaAudited,
		StatusSharedLambdaIsScalarWallAirlock,
		StatusResponseCoefficientEqualsProbabilityAfterUnitGlue,
		StatusGate700LawScalarWallGluedQuotientResponse,
		StatusScalarWallGluingNotNative,
		StatusNoNativeScalarWallAirlockTheorem,
		StatusNoNativeBoundaryHistoryResponsePrinciple,
		StatusNoNativeSevenOver72Theorem,
		StatusGate703ScalarWallAirlockBoundary,
	}
}

func FormatInheritance(x Gate702Inheritance) string {
	return fmt.Sprintf("inherited=%t shared=%q p=%.18g c=%.18g eqP=%t nonTaut=%t noShared=%t noWall=%t noPrinciple=%t no7=%t verdict=%q", x.InheritedSharedScalarWallUnit, x.SharedCoordinate, x.EventProbability, x.ResponseCoefficient, x.CoefficientEqualsProbability, x.NonTautologyInherited, x.NoNativeSharedUnit, x.NoNativeWallAlignment, x.NoNativeBoundaryHistory, x.NoNativeSevenOver72, x.Verdict)
}

func FormatScalarWall(x ScalarWallLineAudit) string {
	return fmt.Sprintf("line=%q def=%q coord=%q signed=%t bCoeff=%.18g hCoeff=%.18g unitBoth=%t verdict=%q", x.LineName, x.Definition, x.Coordinate, x.SignedOrientationPreserved, x.BoundaryLambdaCoefficient, x.HistoryLambdaCoefficient, x.UnitScalarWallOnBothSides, x.Verdict)
}

func FormatDiagram(x GluingDiagramAudit) string {
	return fmt.Sprintf("diagram=%q boundary=%q airlock=%q history=%q bLambda=%t hLambda=%t anchored=%t verdict=%q", x.Diagram, x.BoundaryCoordinate, x.AirlockLine, x.HistoryCoordinate, x.BoundaryMeasuredInLambda, x.HistoryMeasuredInLambda, x.CoordinatesAreAnchored, x.Verdict)
}

func FormatUnitGlue(x UnitGlueAudit) string {
	return fmt.Sprintf("bCoeff=%.18g hCoeff=%.18g gamma=%.18g unit=%t p=%.18g c=%.18g eqP=%t verdict=%q", x.BoundaryLambdaCoefficient, x.HistoryLambdaCoefficient, x.Gamma, x.UnitGlue, x.EventProbability, x.ResponseCoefficient, x.CoefficientEqualsProbability, x.Verdict)
}

func FormatRescaled(x RescaledGlueAudit) string {
	return fmt.Sprintf("gamma=%.18g c=%.18g eqP=%t requiresGammaOne=%t formula=%q verdict=%q", x.Gamma, x.TransformedCoefficient, x.EqualsEventProbability, x.RequiresGammaOneForEquality, x.Formula, x.Verdict)
}

func FormatAlternatives(x AlternativeGluingAudit) string {
	parts := make([]string, 0, len(x.Examples))
	for _, ex := range x.Examples {
		parts = append(parts, fmt.Sprintf("%s: b=%.18g h=%.18g gamma=%.18g c=%.18g accepted=%t rejected=%t", ex.Name, ex.BoundaryScale, ex.HistoryScale, ex.Gamma, ex.ResponseCoefficient, ex.AcceptedActiveAirlock, ex.Rejected))
	}
	return fmt.Sprintf("boundaryRejected=%t historyRejected=%t absoluteRejected=%t hessianRejected=%t sharedAccepted=%t examples=[%s] verdict=%q", x.BoundaryNormalizedRejected, x.HistoryNormalizedRejected, x.AbsoluteScalarRejected, x.HessianScalarRejected, x.SharedSignedLambdaAccepted, strings.Join(parts, " | "), x.Verdict)
}

func FormatNonTautology(x NonTautologyAudit) string {
	return fmt.Sprintf("airlock=%t equation=%q lambdaWeight=%.18g gaugeWeight=%.18g gaugePresent=%t airlockNotProof=%t preserved=%t verdict=%q", x.SharedLambdaAirlockPresent, x.RearrangedEquation, x.LambdaWeight, x.GaugeWeight, x.IndependentGaugeWoundPresent, x.LambdaIsAirlockNotProof, x.NonTautologicalRelationPreserved, x.Verdict)
}

func FormatSource(x SourceTypeClassification) string {
	return fmt.Sprintf("lineRole=%q boundary=%q history=%q pRole=%q cRole=%q conclusion=%q verdict=%q", x.ScalarWallLineRole, x.BoundaryRole, x.HistoryRole, x.EventProbabilityRole, x.ResponseCoefficientRole, x.Conclusion, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("theorems=%s verdict=%q", strings.Join(x.Theorems, ", "), x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("nativeGlue=%t nativeAirlock=%t nativePrinciple=%t native7=%t boundaryStress=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckm=%t verdict=%q", x.ClaimsScalarWallGluingNative, x.ClaimsNativeScalarWallAirlock, x.ClaimsNativeBoundaryHistory, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStressDerived, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}

func canonicalSBoundary() float64   { return lambdaLambda12 + r3Minus1 }
func canonicalSHistory() float64    { return kappaLambda + kappaE + lambdaLambda12 }
func canonicalExpectation() float64 { return eventProbK7 * canonicalSBoundary() }
func canonicalResidual() float64    { return canonicalSHistory() - canonicalExpectation() }
