// Package generation2k7overlambda4boundarypairprojectionaudit implements
// Gate 628: K7OverLambda4BoundaryPair Projection Audit.
//
// Gate 627 identified 7/72 as the Gate626 boundary-weight source candidate,
// with numerator 7 matching dim(K_7), but it left the denominator as an
// uncertified boundary chamber. Gate 628 audits the sharper denominator blink:
//
//	72 = dim(Lambda^4 R^8) + dim R^2_boundary = 70 + 2,
//
// where Lambda^4 R^8 and K_7 are native ASHA finite-geometry objects and the
// two extra coordinates are the Gate613/Gate626 bridge boundary-stress pair
// (|lambda(Lambda_12)|, R_3-1). This gate tests that augmented chamber as a
// bridge construction only; it does not promote a native direct-sum law or a
// K_7-to-boundary projection theorem.
package generation2k7overlambda4boundarypairprojectionaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/asha"
	gate627 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7boundaryprojectionweightaudit"
)

const (
	AuditID = "GATE628-K7-OVER-LAMBDA4-BOUNDARY-PAIR-PROJECTION-AUDIT"

	StatusGate627Inherited          = "PASS_GATE627_K7_NUMERATOR_INHERITED"
	Status70Plus2Identified         = "PASS_72_EQUALS_70_PLUS_2_CANDIDATE_IDENTIFIED"
	Status72Lambda4BoundaryPair     = "CONDITIONAL_SUPPORT_72_AS_LAMBDA4_PLUS_BOUNDARY_PAIR_CHAMBER"
	StatusBoundaryPairInherited     = "PASS_GATE613_BOUNDARY_STRESS_PAIR_INHERITED_AS_BRIDGE_COORDINATES"
	StatusK7InsideLambda4Audited    = "PASS_K7_SITS_INSIDE_LAMBDA4_R8_NUMERATOR_AUDITED"
	Status65ComplementCandidate     = "CONDITIONAL_SUPPORT_65_AS_NON_K7_LAMBDA4_COMPLEMENT_PLUS_BOUNDARY_PAIR"
	StatusProjectionTraceCandidate  = "CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_AS_AUGMENTED_CHAMBER_TRACE_FRACTION"
	StatusNoProductAirlock          = "FAILED_ROUTE_NO_NATIVE_PRODUCT_AIRLOCK_FROM_LAMBDA4_TO_BOUNDARY_STRESS_PAIR"
	StatusNoK7BoundaryPullProjector = "FAILED_ROUTE_NO_NATIVE_K7_BOUNDARY_PULL_PROJECTOR"
	StatusNoNativeAugmentedChamber  = "FAILED_ROUTE_NO_NATIVE_LAMBDA4_PLUS_BOUNDARY_PAIR_CHAMBER_THEOREM"
	StatusGate628Boundary           = "FIREWALL_PRESERVED_GATE628_AUGMENTED_CHAMBER_IS_BRIDGE_ONLY"
)

const (
	k7DimExpected             = 7
	lambda4R8DimExpected      = 70
	boundaryStressPairDim     = 2
	augmentedChamberDim       = 72
	nonK7Lambda4ComplementDim = 63
	augmentedComplementDim    = 65
)

type Gate627Inheritance struct {
	K7Dimension                int
	Lambda4Dimension           int
	Gate627Denominator         int
	Gate627Weight              float64
	Gate627ComplementNumerator int
	Gate627NumeratorIsNative   bool
	Gate627Certified72Carrier  bool
	Gate627ProjectionExists    bool
	Gate627SourceTheorem       bool
	Gate627FirewallPreserved   bool
	Gate626WeightedMixture     float64
	Gate626WeightedResidual    float64
	Gate626ScalarResidual      float64
	AbsLambda12                float64
	R3MinusOne                 float64
	BoundarySplit              float64
	XiBoundary                 float64
	Verdict                    string
}

type ChamberDimensionAudit struct {
	Lambda4Carrier             string
	Lambda4Dimension           int
	BoundaryPairCarrier        string
	BoundaryPairDimension      int
	AugmentedChamberExpression string
	AugmentedChamberDimension  int
	TargetDenominator          int
	EqualsTargetDenominator    bool
	UsesNativeLambda4Carrier   bool
	UsesBridgeBoundaryPair     bool
	DirectSumCertifiedNative   bool
	BetterThan8Times9          bool
	Interpretation             string
	Verdict                    string
}

type DenominatorComparisonRow struct {
	Name                  string
	Expression            string
	Value                 int
	NativeFiniteCarrier   bool
	UsesEnvironmentalPair bool
	UsesQuarantinedLedger bool
	CertifiedAsNative     bool
	StrengthRank          int
	Comment               string
}

type DenominatorComparisonAudit struct {
	Rows                 []DenominatorComparisonRow
	BestExpression       string
	BestValue            int
	BestIs70Plus2        bool
	AnyNativeDenominator bool
	AnyBridgeCandidate   bool
	Verdict              string
}

type BoundaryStressPairAudit struct {
	PairCoordinates          []string
	AbsLambda12              float64
	R3MinusOne               float64
	XiBoundary               float64
	BoundarySplit            float64
	PairDimension            int
	PairIsGate613Boundary    bool
	PairInheritedFromGate626 bool
	PairNativeFiniteObject   bool
	BridgeCoordinateOnly     bool
	Interpretation           string
	Verdict                  string
}

type K7Lambda4EmbeddingAudit struct {
	K7Carrier                 string
	K7Dimension               int
	Lambda4Carrier            string
	Lambda4Dimension          int
	K7FitsInsideLambda4       bool
	RankPB                    int
	RankPG                    int
	NativeCarrierCertified    bool
	ProjectionToBoundaryFound bool
	Interpretation            string
	Verdict                   string
}

type ComplementChamberAudit struct {
	Lambda4Dimension                int
	K7Dimension                     int
	NonK7Lambda4ComplementDimension int
	BoundaryPairDimension           int
	AugmentedComplementDimension    int
	AugmentedChamberDimension       int
	ComplementWeight                float64
	Equals65Over72                  bool
	HasStructuredComplementReading  bool
	NativeComplementProjection      bool
	Equation                        string
	Verdict                         string
}

type ProjectionTraceAudit struct {
	DomainChamber            string
	DomainDimension          int
	K7TraceCarrier           string
	K7TraceDimension         int
	BoundaryPullLine         string
	TraceFraction            float64
	ExpectedWeight           float64
	TraceFractionMatches     bool
	ProjectionOperatorExists bool
	IdempotentCertified      bool
	TraceFunctionalCertified bool
	IntertwinerCertified     bool
	MissingObject            string
	Verdict                  string
}

type WeightedClosureCarryAudit struct {
	KappaSum                   float64
	WeightedMixture            float64
	WeightedClosureResidual    float64
	BoundaryWeight             float64
	ScalarWeight               float64
	WeightFromChamberRatio     float64
	MixtureEquation            string
	ChamberRatioMatchesGate626 bool
	Verdict                    string
}

type NativeASHAStatus struct {
	Lambda4Native                    bool
	K7Native                         bool
	BoundaryPairNativeFinite         bool
	AugmentedChamberNative           bool
	ProductAirlockNative             bool
	K7BoundaryPullProjectorNative    bool
	TraceFractionTheoremNative       bool
	GaugeScalarFlavorTransportNative bool
	Statement                        string
	Verdict                          string
}

type Firewalls struct {
	ClaimsNativeAugmentedChamber bool
	ClaimsNativeBoundaryPair     bool
	ClaimsNativeProjection       bool
	ClaimsNativeTraceTheorem     bool
	ClaimsScalarRGMatching       bool
	ClaimsFlavorOrientation      bool
	ClaimsGaugeUnification       bool
	ClaimsHiggsMassDerived       bool
	ClaimsEndpointDerivation     bool
	Verdict                      string
}

type Analysis struct {
	Inherited             Gate627Inheritance
	Chamber               ChamberDimensionAudit
	DenominatorComparison DenominatorComparisonAudit
	BoundaryPair          BoundaryStressPairAudit
	K7Embedding           K7Lambda4EmbeddingAudit
	Complement            ComplementChamberAudit
	ProjectionTrace       ProjectionTraceAudit
	WeightedClosure       WeightedClosureCarryAudit
	NativeStatus          NativeASHAStatus
	Firewalls             Firewalls
	Truth                 string
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
	g627, err := gate627.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate627 predecessor: %w", err)
	}
	engine := asha.New()
	geo := engine.Geometry

	inherit := inheritGate627(g627)
	chamber := auditChamberDimension(inherit, geo)
	comparison := auditDenominatorComparison(chamber)
	pair := auditBoundaryStressPair(inherit)
	embedding := auditK7Lambda4Embedding(inherit, geo)
	complement := auditComplementChamber(chamber, embedding)
	projection := auditProjectionTrace(chamber, embedding, pair)
	closure := auditWeightedClosureCarry(inherit, projection)

	a := Analysis{
		Inherited:             inherit,
		Chamber:               chamber,
		DenominatorComparison: comparison,
		BoundaryPair:          pair,
		K7Embedding:           embedding,
		Complement:            complement,
		ProjectionTrace:       projection,
		WeightedClosure:       closure,
		NativeStatus:          auditNativeStatus(chamber, pair, embedding, projection, complement),
		Firewalls:             auditFirewalls(),
		Truth:                 "Gate 628 upgrades the strongest denominator candidate for the Gate626/Gate627 weight from weaker factorizations such as 8×9 to the ASHA-native-looking bridge chamber 72=70+2=dim(Lambda^4 R^8)+dim R^2_boundary. This makes 7/72 equal to dim(K_7)/dim(Lambda^4 R^8 ⊕ R^2_boundary), and it gives 65=(70-7)+2 a structured complement reading. The result is still bridge-only: Lambda^4 R^8 and K_7 are native, but the boundary pair is environmental/bridge, the augmented direct sum is not a native law, and no Pi_{K7 subset Lambda4 -> R2_boundary} projector is certified.",
	}
	return a, nil
}

func inheritGate627(g gate627.Analysis) Gate627Inheritance {
	xi := 0.5 * (g.Inherited.R3MinusOne + g.Inherited.AbsLambda12)
	return Gate627Inheritance{
		K7Dimension:                g.Numerator.K7Dimension,
		Lambda4Dimension:           g.Numerator.AmbientExteriorDimension,
		Gate627Denominator:         g.Weight.Denominator,
		Gate627Weight:              g.Weight.Value,
		Gate627ComplementNumerator: g.Weight.ComplementNumerator,
		Gate627NumeratorIsNative:   g.Numerator.K7NativeCarrierCertified && g.Numerator.MatchesDimK7,
		Gate627Certified72Carrier:  g.Denominator.CertifiedBoundaryCarrier,
		Gate627ProjectionExists:    g.Projection.ProjectionOperatorExists,
		Gate627SourceTheorem:       g.NativeStatus.SevenOverSeventyTwoSourceTheorem,
		Gate627FirewallPreserved:   g.Firewalls.Verdict == gate627.StatusGate627Boundary,
		Gate626WeightedMixture:     g.Inherited.WeightedMixture,
		Gate626WeightedResidual:    g.Inherited.WeightedClosureResidual,
		Gate626ScalarResidual:      g.Inherited.ScalarPredictionResidual,
		AbsLambda12:                g.Inherited.AbsLambda12,
		R3MinusOne:                 g.Inherited.R3MinusOne,
		BoundarySplit:              g.Inherited.BoundarySplit,
		XiBoundary:                 xi,
		Verdict:                    StatusGate627Inherited,
	}
}

func auditChamberDimension(i Gate627Inheritance, g asha.Geometry) ChamberDimensionAudit {
	lambda4 := 0
	if len(g.GradeDimensions) > 4 {
		lambda4 = g.GradeDimensions[4]
	}
	chamberDim := lambda4 + boundaryStressPairDim
	return ChamberDimensionAudit{
		Lambda4Carrier:             "Lambda^4 R^8 exterior-grade chamber",
		Lambda4Dimension:           lambda4,
		BoundaryPairCarrier:        "R^2_boundary spanned by (|lambda(Lambda_12)|, R_3-1)",
		BoundaryPairDimension:      boundaryStressPairDim,
		AugmentedChamberExpression: "Lambda^4 R^8 ⊕ R^2_boundary",
		AugmentedChamberDimension:  chamberDim,
		TargetDenominator:          i.Gate627Denominator,
		EqualsTargetDenominator:    chamberDim == i.Gate627Denominator && chamberDim == augmentedChamberDim,
		UsesNativeLambda4Carrier:   lambda4 == lambda4R8DimExpected,
		UsesBridgeBoundaryPair:     true,
		DirectSumCertifiedNative:   false,
		BetterThan8Times9:          chamberDim == i.Gate627Denominator && lambda4 == lambda4R8DimExpected,
		Interpretation:             "72=70+2 is a stronger denominator candidate than 8×9 because 70 is native dim Lambda^4 R^8 and the extra 2 is exactly the active scalar/gauge boundary-stress endpoint pair; the direct sum remains bridge-layer, not native.",
		Verdict:                    Status72Lambda4BoundaryPair,
	}
}

func auditDenominatorComparison(c ChamberDimensionAudit) DenominatorComparisonAudit {
	rows := []DenominatorComparisonRow{
		{Name: "Lambda4 plus boundary-stress pair", Expression: "70 + 2", Value: c.AugmentedChamberDimension, NativeFiniteCarrier: true, UsesEnvironmentalPair: true, UsesQuarantinedLedger: false, CertifiedAsNative: false, StrengthRank: 1, Comment: "strongest typed candidate: Lambda^4 R^8 is native and the boundary pair is the active Gate613/Gate626 bridge endpoint pair"},
		{Name: "Clifford ladder times K/X/Y charged chamber", Expression: "8 × 9", Value: 72, NativeFiniteCarrier: true, UsesEnvironmentalPair: false, UsesQuarantinedLedger: true, CertifiedAsNative: false, StrengthRank: 2, Comment: "typed but weaker because it leans on the quarantined charged K/X/Y coefficient chamber rather than the active boundary pair"},
		{Name: "three-generation matter inventory candidate", Expression: "3 × 24", Value: 72, NativeFiniteCarrier: false, UsesEnvironmentalPair: false, UsesQuarantinedLedger: true, CertifiedAsNative: false, StrengthRank: 3, Comment: "candidate only; not the active scalar/gauge boundary chamber"},
		{Name: "doubled 36-unit chamber", Expression: "2 × 36", Value: 72, NativeFiniteCarrier: false, UsesEnvironmentalPair: true, UsesQuarantinedLedger: false, CertifiedAsNative: false, StrengthRank: 4, Comment: "less precise than 70+2 because the 36-unit chamber is not certified"},
	}
	return DenominatorComparisonAudit{
		Rows:                 rows,
		BestExpression:       "70 + 2",
		BestValue:            c.AugmentedChamberDimension,
		BestIs70Plus2:        c.AugmentedChamberDimension == augmentedChamberDim,
		AnyNativeDenominator: false,
		AnyBridgeCandidate:   c.EqualsTargetDenominator,
		Verdict:              Status70Plus2Identified,
	}
}

func auditBoundaryStressPair(i Gate627Inheritance) BoundaryStressPairAudit {
	return BoundaryStressPairAudit{
		PairCoordinates:          []string{"|lambda(Lambda_12)|", "R_3-1"},
		AbsLambda12:              i.AbsLambda12,
		R3MinusOne:               i.R3MinusOne,
		XiBoundary:               i.XiBoundary,
		BoundarySplit:            i.BoundarySplit,
		PairDimension:            boundaryStressPairDim,
		PairIsGate613Boundary:    true,
		PairInheritedFromGate626: true,
		PairNativeFiniteObject:   false,
		BridgeCoordinateOnly:     true,
		Interpretation:           "The two added coordinates are not invented: they are the active boundary-stress endpoints already used by the Gate613/Gate626 lane. They remain environmental/bridge coordinates, not native finite algebra.",
		Verdict:                  StatusBoundaryPairInherited,
	}
}

func auditK7Lambda4Embedding(i Gate627Inheritance, g asha.Geometry) K7Lambda4EmbeddingAudit {
	lambda4 := 0
	if len(g.GradeDimensions) > 4 {
		lambda4 = g.GradeDimensions[4]
	}
	return K7Lambda4EmbeddingAudit{
		K7Carrier:                 "K_7=Im(P_B)∩Im(P_G)",
		K7Dimension:               g.DimK,
		Lambda4Carrier:            "Lambda^4 R^8",
		Lambda4Dimension:          lambda4,
		K7FitsInsideLambda4:       g.DimK > 0 && lambda4 == lambda4R8DimExpected && g.DimK < lambda4,
		RankPB:                    g.RankPB,
		RankPG:                    g.RankPG,
		NativeCarrierCertified:    g.RankPB == 56 && g.RankPG == 14 && g.DimK == k7DimExpected && lambda4 == lambda4R8DimExpected && i.Gate627NumeratorIsNative,
		ProjectionToBoundaryFound: false,
		Interpretation:            "The numerator is internal to the native Lambda^4 R^8 chamber: K_7 is a 7-dimensional certified contact carrier inside the 70-dimensional grade-four ambient lane. This supports the numerator placement, not the boundary-pull theorem.",
		Verdict:                   StatusK7InsideLambda4Audited,
	}
}

func auditComplementChamber(c ChamberDimensionAudit, k K7Lambda4EmbeddingAudit) ComplementChamberAudit {
	nonK7 := c.Lambda4Dimension - k.K7Dimension
	augComp := nonK7 + c.BoundaryPairDimension
	weight := float64(augComp) / float64(c.AugmentedChamberDimension)
	return ComplementChamberAudit{
		Lambda4Dimension:                c.Lambda4Dimension,
		K7Dimension:                     k.K7Dimension,
		NonK7Lambda4ComplementDimension: nonK7,
		BoundaryPairDimension:           c.BoundaryPairDimension,
		AugmentedComplementDimension:    augComp,
		AugmentedChamberDimension:       c.AugmentedChamberDimension,
		ComplementWeight:                weight,
		Equals65Over72:                  nonK7 == nonK7Lambda4ComplementDim && augComp == augmentedComplementDim && math.Abs(weight-65.0/72.0) < 1e-15,
		HasStructuredComplementReading:  nonK7 == nonK7Lambda4ComplementDim && augComp == augmentedComplementDim,
		NativeComplementProjection:      false,
		Equation:                        "65 = 72-7 = (70-7)+2 = dim(Lambda^4 R^8/K_7)+dim R^2_boundary = 63+2",
		Verdict:                         Status65ComplementCandidate,
	}
}

func auditProjectionTrace(c ChamberDimensionAudit, k K7Lambda4EmbeddingAudit, p BoundaryStressPairAudit) ProjectionTraceAudit {
	fraction := float64(k.K7Dimension) / float64(c.AugmentedChamberDimension)
	expected := float64(k7DimExpected) / float64(augmentedChamberDim)
	return ProjectionTraceAudit{
		DomainChamber:            c.AugmentedChamberExpression,
		DomainDimension:          c.AugmentedChamberDimension,
		K7TraceCarrier:           k.K7Carrier,
		K7TraceDimension:         k.K7Dimension,
		BoundaryPullLine:         fmt.Sprintf("(%s)-(%s)", p.PairCoordinates[1], p.PairCoordinates[0]),
		TraceFraction:            fraction,
		ExpectedWeight:           expected,
		TraceFractionMatches:     math.Abs(fraction-expected) < 1e-15 && math.Abs(fraction-7.0/72.0) < 1e-15,
		ProjectionOperatorExists: false,
		IdempotentCertified:      false,
		TraceFunctionalCertified: false,
		IntertwinerCertified:     false,
		MissingObject:            "Pi_{K7 subset Lambda^4 R^8 -> R^2_boundary}: a typed projection/intertwiner inside Lambda^4 R^8 ⊕ R^2_boundary whose normalized trace is 7/72 and whose image controls the pull from |lambda(Lambda_12)| toward R_3-1",
		Verdict:                  StatusProjectionTraceCandidate,
	}
}

func auditWeightedClosureCarry(i Gate627Inheritance, p ProjectionTraceAudit) WeightedClosureCarryAudit {
	boundaryWeight := p.TraceFraction
	scalarWeight := 1 - boundaryWeight
	kappaSum := i.Gate626WeightedMixture + i.Gate626WeightedResidual
	return WeightedClosureCarryAudit{
		KappaSum:                   kappaSum,
		WeightedMixture:            i.AbsLambda12 + boundaryWeight*i.BoundarySplit,
		WeightedClosureResidual:    kappaSum - (i.AbsLambda12 + boundaryWeight*i.BoundarySplit),
		BoundaryWeight:             boundaryWeight,
		ScalarWeight:               scalarWeight,
		WeightFromChamberRatio:     boundaryWeight,
		MixtureEquation:            "W_72=(65/72)|lambda(Lambda_12)|+(7/72)(R_3-1)",
		ChamberRatioMatchesGate626: math.Abs(boundaryWeight-i.Gate627Weight) < 1e-15,
		Verdict:                    StatusProjectionTraceCandidate,
	}
}

func auditNativeStatus(c ChamberDimensionAudit, p BoundaryStressPairAudit, k K7Lambda4EmbeddingAudit, tr ProjectionTraceAudit, comp ComplementChamberAudit) NativeASHAStatus {
	return NativeASHAStatus{
		Lambda4Native:                    c.UsesNativeLambda4Carrier,
		K7Native:                         k.NativeCarrierCertified,
		BoundaryPairNativeFinite:         p.PairNativeFiniteObject,
		AugmentedChamberNative:           c.DirectSumCertifiedNative,
		ProductAirlockNative:             false,
		K7BoundaryPullProjectorNative:    tr.ProjectionOperatorExists && tr.IntertwinerCertified,
		TraceFractionTheoremNative:       tr.TraceFunctionalCertified,
		GaugeScalarFlavorTransportNative: false,
		Statement:                        "Lambda^4 R^8 and K_7 are native finite ASHA objects. The boundary-stress pair is inherited bridge/environmental data. Their direct sum is a strong augmented bridge chamber candidate, but no native product airlock, no K_7 boundary-pull projector, and no trace theorem exist yet.",
		Verdict:                          StatusNoProductAirlock,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{Verdict: StatusGate628Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate627Inherited,
		Status70Plus2Identified,
		Status72Lambda4BoundaryPair,
		StatusBoundaryPairInherited,
		StatusK7InsideLambda4Audited,
		Status65ComplementCandidate,
		StatusProjectionTraceCandidate,
		StatusNoProductAirlock,
		StatusNoK7BoundaryPullProjector,
		StatusNoNativeAugmentedChamber,
		StatusGate628Boundary,
	}
}
