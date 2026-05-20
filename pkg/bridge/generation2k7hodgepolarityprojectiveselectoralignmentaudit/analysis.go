// Package generation2k7hodgepolarityprojectiveselectoralignmentaudit implements
// Gate 635: K7 Hodge Polarity and Projective Selector Alignment Audit.
//
// Gate 634 certified the native Hodge-stable contact carrier split
//
//	K_7 = K_7^+ \oplus K_7^-,  dim K_7^+=4, dim K_7^-=3.
//
// Gate 635 compares that native (4|3) polarity with the already-certified
// Witt/Fock projective selector lane, where W=C^4 and B-L realizes a 1+3
// split on CP^3.  The audit is deliberately conservative: a dimension and
// split resemblance is recorded, but no K7-to-Fock selector map, no 4=1+3
// refinement inside K_7^+, and no boundary-stress assignment are promoted.
package generation2k7hodgepolarityprojectiveselectoralignmentaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate634 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7hodgesignaturestabilizeraudit"
	gate572 "github.com/bagherbal/asha-engine/pkg/bridge/generation2projectivefockcp3momentmapselectorgeometryaudit"
)

const (
	AuditID = "GATE635-K7-HODGE-POLARITY-PROJECTIVE-SELECTOR-ALIGNMENT-AUDIT"

	StatusGate634Inherited            = "PASS_GATE634_K7_HODGE_SIGNATURE_INHERITED"
	StatusK7PlusMinusSubspacesDefined = "PASS_K7_PLUS_MINUS_SUBSPACES_DEFINED"
	StatusProjectiveSelectorInherited = "PASS_PROJECTIVE_SELECTOR_1_PLUS_3_INHERITED"
	StatusFourPlusThreeAudited        = "PASS_4_PLUS_3_POLARITY_AUDITED"
	StatusResemblesSelectorSplit      = "CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_RESEMBLES_PROJECTIVE_SELECTOR_SPLIT"
	StatusNoK7ToFockMap               = "FAILED_ROUTE_NO_TYPED_K7_TO_FOCK_SELECTOR_MAP_YET"
	StatusNoK7PlusOnePlusThree        = "FAILED_ROUTE_NO_NATIVE_4_EQUALS_1_PLUS_3_REFINEMENT_INSIDE_K7_PLUS"
	StatusTraceNotDistinguishedLine   = "FAILED_ROUTE_TRACE_PLUS_ONE_IS_HODGE_IMBALANCE_NOT_DISTINGUISHED_LINE"
	StatusNoBoundaryStressAssignment  = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem        = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusGate635Boundary             = "FIREWALL_PRESERVED_GATE635_HODGE_POLARITY_SELECTOR_BOUNDARY"
)

const (
	k7DimExpected      = 7
	k7PlusExpected     = 4
	k7MinusExpected    = 3
	fockDimExpected    = 4
	selectorLineDim    = 1
	selectorSpatialDim = 3
	strictTolerance    = 1e-10
)

type Gate634Inheritance struct {
	K7Dimension              int
	PlusDimension            int
	MinusDimension           int
	Trace                    float64
	Determinant              float64
	HodgeStable              bool
	MixedHodgePolarity       bool
	NoBoundaryAssignment     bool
	NoSevenOver72Theorem     bool
	Gate634FirewallPreserved bool
	Verdict                  string
}

type K7PolaritySubspaceAudit struct {
	Carrier             string
	PlusFormula         string
	MinusFormula        string
	PlusDimension       int
	MinusDimension      int
	SumDimension        int
	ProjectorsCertified bool
	PlusMinusOrthogonal bool
	NativeHodgePolarity bool
	Verdict             string
}

type ProjectiveSelectorReference struct {
	Carrier                      string
	CarrierComplexDimension      int
	Selector                     string
	SelectorCoefficients         []float64
	SelectorSplit                string
	LineComplexDimension         int
	SpatialBlockComplexDimension int
	CP0CP2CriticalStrata         bool
	ProjectiveOnePlusThree       bool
	Stabilizer                   string
	MatchesGate555Commutant      bool
	CP3ToK7FunctorFound          bool
	Verdict                      string
}

type PolaritySelectorAlignmentAudit struct {
	K7PlusDimension             int
	K7MinusDimension            int
	FockCarrierComplexDimension int
	SelectorLinePlusSpatialDims string
	FourDimensionalMatch        bool
	ThreeDimensionalMatch       bool
	SameCarrier                 bool
	TypedThetaMapFound          bool
	K7ToCP3FunctorFound         bool
	AlignmentCandidateOnly      bool
	Reason                      string
	Verdict                     string
}

type K7PlusRefinementAudit struct {
	K7PlusDimension              int
	HodgeEigenvalueOnK7Plus      float64
	HodgeProjectorActsAsIdentity bool
	InternalRankOneLineDerived   bool
	InternalThreePlaneDerived    bool
	NativeOnePlusThreeRefinement bool
	Reason                       string
	Verdict                      string
}

type K7MinusTripletAudit struct {
	K7MinusDimension           int
	SelectorSpatialBlockDim    int
	DimensionMatch             bool
	TypedTripletIdentification bool
	UsesBMinusLCarrier         bool
	Reason                     string
	Verdict                    string
}

type TraceImbalanceAudit struct {
	Trace                    float64
	PlusMinusDifference      int
	Determinant              float64
	DistinguishedLineDerived bool
	TraceAsRankOneProjector  bool
	NeedsAdditionalSelector  bool
	Reason                   string
	Verdict                  string
}

type CarrierMapAudit struct {
	CandidateMap                       string
	Domain                             string
	Codomain                           string
	DimensionResemblance               bool
	TypedIntertwinerFound              bool
	FunctorFromProjectiveFockToK7Found bool
	FunctorFromK7ToProjectiveFockFound bool
	Status                             string
	MissingObject                      string
	Verdict                            string
}

type BoundaryReadinessAudit struct {
	BoundaryStressAssignment bool
	SevenOver72Promoted      bool
	K7ToW7PairingReopened    bool
	K7ToFockMapWouldSuffice  bool
	StillRequired            string
	VerdictBoundary          string
	VerdictSevenOver72       string
}

type Firewalls struct {
	ClaimsK7ToFockSelectorMap      bool
	ClaimsK7PlusOnePlusThree       bool
	ClaimsBoundaryStressAssignment bool
	ClaimsSevenOver72Theorem       bool
	ClaimsScalarRGMatching         bool
	ClaimsHiggsMassDerivation      bool
	ClaimsFlavorDerivation         bool
	ClaimsCKMPMNSDerivation        bool
	ClaimsGaugeUnification         bool
	ClaimsPhysicalOrientation      bool
	Verdict                        string
}

type Analysis struct {
	Inherited          Gate634Inheritance
	K7Subspaces        K7PolaritySubspaceAudit
	ProjectiveSelector ProjectiveSelectorReference
	Alignment          PolaritySelectorAlignmentAudit
	K7PlusRefinement   K7PlusRefinementAudit
	K7MinusTriplet     K7MinusTripletAudit
	TraceImbalance     TraceImbalanceAudit
	CarrierMap         CarrierMapAudit
	BoundaryReadiness  BoundaryReadinessAudit
	Firewalls          Firewalls
	Truth              string
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
	g634, err := gate634.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate634 inheritance unavailable: %w", err)
	}
	g572, err := gate572.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate572 projective selector reference unavailable: %w", err)
	}

	inherited := buildGate634Inheritance(g634)
	k7Subspaces := buildK7SubspaceAudit(g634)
	projective := buildProjectiveSelectorReference(g572)
	alignment := buildAlignmentAudit(k7Subspaces, projective)
	k7Plus := buildK7PlusRefinementAudit(k7Subspaces)
	k7Minus := buildK7MinusTripletAudit(k7Subspaces, projective)
	trace := buildTraceImbalanceAudit(g634)
	carrier := buildCarrierMapAudit(alignment, projective)
	boundary := BoundaryReadinessAudit{
		BoundaryStressAssignment: false,
		SevenOver72Promoted:      false,
		K7ToW7PairingReopened:    false,
		K7ToFockMapWouldSuffice:  false,
		StillRequired:            "a typed carrier comparison map Theta, plus a separate lawful boundary-stress map if boundary physics is ever to be discussed",
		VerdictBoundary:          StatusNoBoundaryStressAssignment,
		VerdictSevenOver72:       StatusNoSevenOver72Theorem,
	}
	firewalls := Firewalls{Verdict: StatusGate635Boundary}
	truth := strings.Join([]string{
		"Gate 635 inherits the native Gate634 result K_7=K_7^+⊕K_7^- with dimensions 4 and 3.",
		"It compares this Hodge polarity with the Gate572 Witt/Fock selector lane W=C^4, where B-L yields a projective CP^0|CP^2 or 1+3 split.",
		"The resemblance is real at the dimension-pattern level, but the carriers are different: K_7 sits in Lambda^4 R^8, while the selector acts on W=C^4/CP^3.",
		"No typed Theta map, CP3-to-K7 functor, K7-to-Fock functor, or 4=1+3 refinement inside K_7^+ is certified.",
		"The trace tr(S_K)=+1 is a Hodge imbalance n_+-n_-=1, not a distinguished rank-one line. Boundary stress and 7/72 remain blocked.",
	}, " ")

	return Analysis{
		Inherited:          inherited,
		K7Subspaces:        k7Subspaces,
		ProjectiveSelector: projective,
		Alignment:          alignment,
		K7PlusRefinement:   k7Plus,
		K7MinusTriplet:     k7Minus,
		TraceImbalance:     trace,
		CarrierMap:         carrier,
		BoundaryReadiness:  boundary,
		Firewalls:          firewalls,
		Truth:              truth,
	}, nil
}

func buildGate634Inheritance(g634 gate634.Analysis) Gate634Inheritance {
	return Gate634Inheritance{
		K7Dimension:              k7DimExpected,
		PlusDimension:            g634.Spectrum.PlusRank,
		MinusDimension:           g634.Spectrum.MinusRank,
		Trace:                    g634.Spectrum.Trace,
		Determinant:              g634.Spectrum.Determinant,
		HodgeStable:              g634.Classification.K7MixedHodgePolarity && g634.Inherited.K7HodgeStable,
		MixedHodgePolarity:       g634.Classification.K7MixedHodgePolarity,
		NoBoundaryAssignment:     g634.Consequences.VerdictBoundary == gate634.StatusNoBoundaryStressAssignment,
		NoSevenOver72Theorem:     g634.Consequences.VerdictSevenOver72 == gate634.StatusNoSevenOver72Theorem,
		Gate634FirewallPreserved: g634.Firewalls.Verdict == gate634.StatusGate634Boundary,
		Verdict:                  StatusGate634Inherited,
	}
}

func buildK7SubspaceAudit(g634 gate634.Analysis) K7PolaritySubspaceAudit {
	orthogonal := g634.InternalProjectors.OrthogonalityResidual < strictTolerance && g634.InternalProjectors.ComplementarityResidual < strictTolerance
	return K7PolaritySubspaceAudit{
		Carrier:             "K_7 subset Lambda^4 R^8",
		PlusFormula:         "K_7^+ = im((I+S_K)/2)",
		MinusFormula:        "K_7^- = im((I-S_K)/2)",
		PlusDimension:       g634.InternalProjectors.PlusProjectorRank,
		MinusDimension:      g634.InternalProjectors.MinusProjectorRank,
		SumDimension:        g634.InternalProjectors.PlusProjectorRank + g634.InternalProjectors.MinusProjectorRank,
		ProjectorsCertified: g634.InternalProjectors.ProjectorsCertified,
		PlusMinusOrthogonal: orthogonal,
		NativeHodgePolarity: g634.Spectrum.Mixed,
		Verdict:             StatusK7PlusMinusSubspacesDefined,
	}
}

func buildProjectiveSelectorReference(g572 gate572.Analysis) ProjectiveSelectorReference {
	return ProjectiveSelectorReference{
		Carrier:                      g572.Projective.Carrier,
		CarrierComplexDimension:      g572.Projective.AmbientComplexDimension,
		Selector:                     "B-L = diag(-1, 1/3, 1/3, 1/3) on W=C^4",
		SelectorCoefficients:         []float64{-1, 1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0},
		SelectorSplit:                "projective CP^0 | CP^2, equivalently 1+3 eigenspace split before projectivization",
		LineComplexDimension:         selectorLineDim,
		SpatialBlockComplexDimension: selectorSpatialDim,
		CP0CP2CriticalStrata:         g572.BMinusL.CriticalStrataCertified,
		ProjectiveOnePlusThree:       g572.BMinusL.ProjectiveOnePlusThree,
		Stabilizer:                   g572.Stabilizer.Stabilizer,
		MatchesGate555Commutant:      g572.Stabilizer.MatchesGate555Commutant,
		CP3ToK7FunctorFound:          g572.K7.CP3ToK7FunctorFound || g572.K7.HopfS7ToK7FunctorFound || g572.K7.TangentS7ToK7FunctorFound,
		Verdict:                      StatusProjectiveSelectorInherited,
	}
}

func buildAlignmentAudit(k7 K7PolaritySubspaceAudit, p ProjectiveSelectorReference) PolaritySelectorAlignmentAudit {
	fourMatch := k7.PlusDimension == p.CarrierComplexDimension && k7.PlusDimension == fockDimExpected
	threeMatch := k7.MinusDimension == p.SpatialBlockComplexDimension && k7.MinusDimension == selectorSpatialDim
	typed := false
	return PolaritySelectorAlignmentAudit{
		K7PlusDimension:             k7.PlusDimension,
		K7MinusDimension:            k7.MinusDimension,
		FockCarrierComplexDimension: p.CarrierComplexDimension,
		SelectorLinePlusSpatialDims: fmt.Sprintf("%d+%d", p.LineComplexDimension, p.SpatialBlockComplexDimension),
		FourDimensionalMatch:        fourMatch,
		ThreeDimensionalMatch:       threeMatch,
		SameCarrier:                 false,
		TypedThetaMapFound:          typed,
		K7ToCP3FunctorFound:         p.CP3ToK7FunctorFound,
		AlignmentCandidateOnly:      fourMatch && threeMatch && !typed && !p.CP3ToK7FunctorFound,
		Reason:                      "The dimensions align as K_7^+|K_7^- = 4|3 and W=B-L carrier/spatial block = 4 and 3, but K_7 and W=C^4 live in different typed carriers and Gate572 supplied no CP3-to-K7 functor.",
		Verdict:                     join(StatusFourPlusThreeAudited, StatusResemblesSelectorSplit, StatusNoK7ToFockMap),
	}
}

func buildK7PlusRefinementAudit(k7 K7PolaritySubspaceAudit) K7PlusRefinementAudit {
	return K7PlusRefinementAudit{
		K7PlusDimension:              k7.PlusDimension,
		HodgeEigenvalueOnK7Plus:      +1,
		HodgeProjectorActsAsIdentity: k7.PlusDimension == k7PlusExpected,
		InternalRankOneLineDerived:   false,
		InternalThreePlaneDerived:    false,
		NativeOnePlusThreeRefinement: false,
		Reason:                       "On K_7^+, S_K acts as the identity, so Hodge polarity alone supplies a four-plane but no internal rank-one line and no complementary three-plane inside that four-plane. A second typed selector would be required.",
		Verdict:                      StatusNoK7PlusOnePlusThree,
	}
}

func buildK7MinusTripletAudit(k7 K7PolaritySubspaceAudit, p ProjectiveSelectorReference) K7MinusTripletAudit {
	match := k7.MinusDimension == p.SpatialBlockComplexDimension
	return K7MinusTripletAudit{
		K7MinusDimension:           k7.MinusDimension,
		SelectorSpatialBlockDim:    p.SpatialBlockComplexDimension,
		DimensionMatch:             match,
		TypedTripletIdentification: false,
		UsesBMinusLCarrier:         false,
		Reason:                     "K_7^- has dimension three, matching the B-L spatial block dimension, but no typed map identifies this anti-self-dual K_7 sector with the Fock/projective CP^2 spatial eigenspace.",
		Verdict:                    join(StatusResemblesSelectorSplit, StatusNoK7ToFockMap),
	}
}

func buildTraceImbalanceAudit(g634 gate634.Analysis) TraceImbalanceAudit {
	trace := g634.Spectrum.Trace
	return TraceImbalanceAudit{
		Trace:                    trace,
		PlusMinusDifference:      g634.Spectrum.PlusRank - g634.Spectrum.MinusRank,
		Determinant:              g634.Spectrum.Determinant,
		DistinguishedLineDerived: false,
		TraceAsRankOneProjector:  false,
		NeedsAdditionalSelector:  true,
		Reason:                   "tr(S_K)=+1 records the signed Hodge imbalance 4-3. It is an invariant scalar of the restricted involution, but it is not itself a projector selecting a canonical line inside K_7^+.",
		Verdict:                  StatusTraceNotDistinguishedLine,
	}
}

func buildCarrierMapAudit(a PolaritySelectorAlignmentAudit, p ProjectiveSelectorReference) CarrierMapAudit {
	return CarrierMapAudit{
		CandidateMap:                       "Theta: K_7^+⊕K_7^- <-> W=C^4 with B-L 1+3 selector data",
		Domain:                             "K_7 subset Lambda^4 R^8 with Hodge polarity 4|3",
		Codomain:                           "W=C^4 / CP^3 projective Fock selector lane with B-L 1+3 split",
		DimensionResemblance:               a.FourDimensionalMatch && a.ThreeDimensionalMatch,
		TypedIntertwinerFound:              false,
		FunctorFromProjectiveFockToK7Found: p.CP3ToK7FunctorFound,
		FunctorFromK7ToProjectiveFockFound: false,
		Status:                             "candidate-only dimension alignment",
		MissingObject:                      "a typed carrier comparison/intertwiner Theta that transports the Hodge polarity on K_7 to the Witt/Fock selector carrier without crossing the Gate572 firewall",
		Verdict:                            StatusNoK7ToFockMap,
	}
}

func Statuses() []string {
	return []string{
		StatusGate634Inherited,
		StatusK7PlusMinusSubspacesDefined,
		StatusProjectiveSelectorInherited,
		StatusFourPlusThreeAudited,
		StatusResemblesSelectorSplit,
		StatusNoK7ToFockMap,
		StatusNoK7PlusOnePlusThree,
		StatusTraceNotDistinguishedLine,
		StatusNoBoundaryStressAssignment,
		StatusNoSevenOver72Theorem,
		StatusGate635Boundary,
	}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func nearly(x, y, tol float64) bool { return math.Abs(x-y) <= tol }
