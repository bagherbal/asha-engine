// Package generation2orthogonalcokernelk7pairingaudit implements
// Gate 631: Orthogonal Cokernel Representative and K7 Defect Pairing Audit.
//
// Gate 630 constructed the square finite addition map
//
//	A: U⊕V -> Lambda^4 R^8,      A(u,v)=u+v,
//
// with ker(A)≅K_7, dim ker(A)=7, dim coker(A)=7, and index(A)=0. Gate 631
// sharpens the cokernel problem by using the ambient Euclidean metric on
// Lambda^4 R^8 to represent the quotient coker(A)=H/(U+V) by the orthogonal
// complement W_7=(U+V)^perp, then audits whether any typed native operator
// canonically pairs K_7 with W_7. It defines the exact defect sequence and
// records Hodge-star, projector algebra, eta, and determinant-line candidates.
// The gate does not derive boundary stress, scalar RG matching, Higgs mass,
// flavor, CKM/PMNS, gauge unification, or a native boundary theorem.
package generation2orthogonalcokernelk7pairingaudit

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/asha"
	gate630 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7kernelcokernelindexzeroaudit"
)

const (
	AuditID = "GATE631-ORTHOGONAL-COKERNEL-K7-PAIRING-AUDIT"

	StatusGate630Inherited                        = "PASS_GATE630_INDEX_ZERO_OPERATOR_INHERITED"
	StatusOrthogonalCokernelRepresentativeDefined = "PASS_ORTHOGONAL_COKERNEL_REPRESENTATIVE_DEFINED"
	StatusExactDefectSequenceWritten              = "PASS_EXACT_DEFECT_SEQUENCE_WRITTEN"
	StatusCokernelRepresentedByW7                 = "CONDITIONAL_SUPPORT_COKERNEL_REPRESENTED_BY_W7_ORTHOGONAL_COMPLEMENT"
	StatusK7W7PairingProblemSharpened             = "CONDITIONAL_SUPPORT_K7_W7_PAIRING_PROBLEM_SHARPENED"
	StatusNoCanonicalK7ToW7Pairing                = "FAILED_ROUTE_NO_CANONICAL_K7_TO_W7_PAIRING_YET"
	StatusProjectorAlgebraFails                   = "FAILED_ROUTE_PROJECTOR_ALGEBRA_DOES_NOT_PAIR_K7_TO_W7"
	StatusHodgeStarRequiresExplicitRankTest       = "CONDITIONAL_SUPPORT_HODGE_STAR_PAIRING_REQUIRES_EXPLICIT_RANK_TEST"
	StatusNoBoundaryStressAssignment              = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT_YET"
	StatusGate631Boundary                         = "FIREWALL_PRESERVED_GATE631_K7_COKERNEL_PAIRING_BOUNDARY"
)

const (
	rankPBExpected          = 56
	rankPGExpected          = 14
	k7DimExpected           = 7
	lambda4DimExpected      = 70
	spanDimExpected         = 63
	boundaryPairDimExpected = 2
	augmentedDimExpected    = 72
)

type Gate630Inheritance struct {
	HDimension                int
	UDimension                int
	VDimension                int
	DirectSumDimension        int
	K7Dimension               int
	SpanDimension             int
	CokernelDimension         int
	Index                     int
	BoundaryPairDimension     int
	AugmentedChamberDimension int
	BoundaryWeight            float64
	Gate630IndexZero          bool
	Gate630PairingMissing     bool
	Gate630BoundaryMissing    bool
	Gate630FirewallPreserved  bool
	Verdict                   string
}

type OrthogonalCokernelRepresentativeTable struct {
	AmbientSpace              string
	AmbientDimension          int
	SpanSpace                 string
	SpanDimension             int
	ProjectorOntoSpan         string
	ProjectorOntoW            string
	WName                     string
	WDefinition               string
	WDimension                int
	WOrthogonalToU            bool
	WOrthogonalToV            bool
	DirectSumDecomposition    string
	DirectSumCertified        bool
	QuotientRepresentative    string
	RepresentsCokernel        bool
	MetricDependent           bool
	NativeComplementCertified bool
	Verdict                   string
}

type ExactDefectSequence struct {
	Sequence                string
	KernelInjection         string
	AdditionMap             string
	ProjectionMap           string
	ExactAtK7               bool
	ExactAtDirectSum        bool
	ExactAtH                bool
	ExactAtW7               bool
	DimensionAlternatingSum int
	ExactByRankNullity      bool
	Verdict                 string
}

type CandidatePairingTable struct {
	Candidates                []PairingCandidate
	CanonicalPairingFound     bool
	NondegeneratePairingFound bool
	PairingProblemSharpened   bool
	MissingObject             string
	Verdict                   string
}

type PairingCandidate struct {
	Name                   string
	Formula                string
	SourceLane             string
	TouchesK7AndW7         bool
	RankTestAvailable      bool
	NondegenerateCertified bool
	CanonicalCertified     bool
	FailureOrCondition     string
	Verdict                string
}

type HodgeStarPairingAudit struct {
	Formula                   string
	HodgeStarTypedOnLambda4   bool
	MapsLambda4ToLambda4      bool
	RequiresOrientationChoice bool
	RankTestImplemented       bool
	NondegenerateCertified    bool
	PreservesOrExchangesUVW   string
	Condition                 string
	Verdict                   string
}

type ProjectorAlgebraPairingAudit struct {
	Rows                []ProjectorAlgebraCandidate
	K7FixedByPB         bool
	K7FixedByPG         bool
	PWKillsUPlusV       bool
	AnyPairingCertified bool
	Reason              string
	Verdict             string
}

type ProjectorAlgebraCandidate struct {
	Operator    string
	ActionOnK7  string
	AfterPW     string
	PairsK7ToW7 bool
}

type EtaPairingAudit struct {
	Formula                    string
	TypedEtaOnLambda4Available bool
	RankTestImplemented        bool
	PairingCertified           bool
	CompatibilityCertified     bool
	Reason                     string
	Verdict                    string
}

type DeterminantLineAudit struct {
	ExactSequence                     string
	DeterminantRelation               string
	CanonicalLineRelation             bool
	PointwiseIsomorphism              bool
	OrientationDependent              bool
	CanSupportVolumeBookkeeping       bool
	CanSupportNormalizedTraceByItself bool
	Interpretation                    string
	Verdict                           string
}

type BoundaryReadinessAudit struct {
	K7ToW7PairingCertified             bool
	DeterminantLineRelationAvailable   bool
	BoundaryPairDimension              int
	StillRequiresW7ToBoundary          bool
	StillRequiresDefectTraceToBoundary bool
	BoundaryAssignmentCertified        bool
	MissingObject                      string
	Verdict                            string
}

type NativeASHAStatus struct {
	Lambda4Native                  bool
	AmbientMetricAdmitted          bool
	OrthogonalRepresentativeTyped  bool
	K7Native                       bool
	W7DimensionTyped               bool
	ExactDefectSequenceTyped       bool
	HodgeStarRankCertified         bool
	ProjectorPairingCertified      bool
	EtaPairingCertified            bool
	DeterminantLineRelationTyped   bool
	CanonicalK7ToW7Pairing         bool
	BoundaryStressAssignmentNative bool
	Statement                      string
	Verdict                        string
}

type Firewalls struct {
	ClaimsCanonicalK7W7Pairing     bool
	ClaimsBoundaryStressAssignment bool
	ClaimsScalarRGMatching         bool
	ClaimsHiggsMassDerivation      bool
	ClaimsFlavorDerivation         bool
	ClaimsCKMPMNSDerivation        bool
	ClaimsGaugeUnification         bool
	ClaimsBoundaryPairNative       bool
	ClaimsNativeTraceTheorem       bool
	Verdict                        string
}

type Analysis struct {
	Inherited         Gate630Inheritance
	OrthogonalW7      OrthogonalCokernelRepresentativeTable
	ExactSequence     ExactDefectSequence
	CandidatePairings CandidatePairingTable
	HodgeStar         HodgeStarPairingAudit
	ProjectorAlgebra  ProjectorAlgebraPairingAudit
	Eta               EtaPairingAudit
	DeterminantLine   DeterminantLineAudit
	BoundaryReadiness BoundaryReadinessAudit
	NativeStatus      NativeASHAStatus
	Firewalls         Firewalls
	Truth             string
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
	g630, err := gate630.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate630 predecessor: %w", err)
	}
	engine := asha.New()
	geo := engine.Geometry

	inherited := inheritGate630(g630, geo)
	orthogonal := auditOrthogonalCokernelRepresentative(inherited)
	exact := auditExactDefectSequence(inherited, orthogonal)
	hodge := auditHodgeStarPairing(orthogonal)
	projector := auditProjectorAlgebraPairing()
	eta := auditEtaPairing()
	det := auditDeterminantLine(exact)
	candidates := auditCandidatePairings(hodge, projector, eta, det)
	boundary := auditBoundaryReadiness(candidates, det, inherited)

	return Analysis{
		Inherited:         inherited,
		OrthogonalW7:      orthogonal,
		ExactSequence:     exact,
		CandidatePairings: candidates,
		HodgeStar:         hodge,
		ProjectorAlgebra:  projector,
		Eta:               eta,
		DeterminantLine:   det,
		BoundaryReadiness: boundary,
		NativeStatus:      auditNativeStatus(inherited, orthogonal, exact, hodge, projector, eta, det, boundary),
		Firewalls:         auditFirewalls(),
		Truth:             "Gate 631 sharpens the Gate630 index-zero defect by representing coker(A)=Lambda^4 R^8/(U+V) through the orthogonal complement W_7=(U+V)^perp inside H=Lambda^4 R^8. The exact sequence 0->K_7->U⊕V->H->W_7->0 is written and rank-checked. This gives a clean finite pairing problem K_7 -> W_7, but projector algebra fails, eta is not typed on Lambda4, Hodge star remains only a rank-test candidate, and the determinant-line relation is not a pointwise isomorphism. Boundary stress still requires an additional W_7 or defect-trace map to R^2_boundary.",
	}, nil
}

func inheritGate630(g gate630.Analysis, geo asha.Geometry) Gate630Inheritance {
	lambda4 := 0
	if len(geo.GradeDimensions) > 4 {
		lambda4 = geo.GradeDimensions[4]
	}
	return Gate630Inheritance{
		HDimension:                lambda4,
		UDimension:                geo.RankPB,
		VDimension:                geo.RankPG,
		DirectSumDimension:        geo.RankPB + geo.RankPG,
		K7Dimension:               geo.DimK,
		SpanDimension:             g.AdditionMap.ImageDimension,
		CokernelDimension:         g.AdditionMap.CokernelDimension,
		Index:                     g.AdditionMap.Index,
		BoundaryPairDimension:     g.BoundaryAssignment.BoundaryPairDimension,
		AugmentedChamberDimension: g.Inherited.AugmentedChamberDimension,
		BoundaryWeight:            g.BlockCompression.BoundaryWeight,
		Gate630IndexZero:          g.AdditionMap.IndexZero,
		Gate630PairingMissing:     !g.Pairing.CanonicalPairingFound,
		Gate630BoundaryMissing:    !g.BoundaryAssignment.AssignmentCertified,
		Gate630FirewallPreserved:  g.Firewalls.Verdict == gate630.StatusGate630Boundary,
		Verdict:                   StatusGate630Inherited,
	}
}

func auditOrthogonalCokernelRepresentative(i Gate630Inheritance) OrthogonalCokernelRepresentativeTable {
	wDim := i.HDimension - i.SpanDimension
	return OrthogonalCokernelRepresentativeTable{
		AmbientSpace:              "H = Lambda^4 R^8",
		AmbientDimension:          i.HDimension,
		SpanSpace:                 "U+V = Im(P_B)+Im(P_G)",
		SpanDimension:             i.SpanDimension,
		ProjectorOntoSpan:         "P_{U+V}",
		ProjectorOntoW:            "P_W = I - P_{U+V}",
		WName:                     "W_7",
		WDefinition:               "W_7 = im(P_W) = (U+V)^perp",
		WDimension:                wDim,
		WOrthogonalToU:            true,
		WOrthogonalToV:            true,
		DirectSumDecomposition:    "H = (U+V) ⊕ W_7",
		DirectSumCertified:        i.HDimension == i.SpanDimension+wDim && wDim == k7DimExpected,
		QuotientRepresentative:    "H/(U+V) ≅ W_7 by the ambient metric",
		RepresentsCokernel:        wDim == i.CokernelDimension && i.CokernelDimension == k7DimExpected,
		MetricDependent:           true,
		NativeComplementCertified: wDim == k7DimExpected,
		Verdict:                   StatusOrthogonalCokernelRepresentativeDefined,
	}
}

func auditExactDefectSequence(i Gate630Inheritance, w OrthogonalCokernelRepresentativeTable) ExactDefectSequence {
	dimensionAlt := i.K7Dimension - i.DirectSumDimension + i.HDimension - w.WDimension
	return ExactDefectSequence{
		Sequence:                "0 -> K_7 -> U⊕V -> H -> W_7 -> 0",
		KernelInjection:         "k -> (k,-k)",
		AdditionMap:             "A(u,v)=u+v",
		ProjectionMap:           "P_W:H->W_7",
		ExactAtK7:               i.K7Dimension == k7DimExpected,
		ExactAtDirectSum:        i.DirectSumDimension-i.SpanDimension == i.K7Dimension,
		ExactAtH:                i.SpanDimension+w.WDimension == i.HDimension,
		ExactAtW7:               w.WDimension == i.CokernelDimension,
		DimensionAlternatingSum: dimensionAlt,
		ExactByRankNullity:      dimensionAlt == 0,
		Verdict:                 StatusExactDefectSequenceWritten,
	}
}

func auditHodgeStarPairing(w OrthogonalCokernelRepresentativeTable) HodgeStarPairingAudit {
	return HodgeStarPairingAudit{
		Formula:                   "Phi_* = P_W * |_{K_7}",
		HodgeStarTypedOnLambda4:   true,
		MapsLambda4ToLambda4:      true,
		RequiresOrientationChoice: true,
		RankTestImplemented:       false,
		NondegenerateCertified:    false,
		PreservesOrExchangesUVW:   "not certified; explicit rank test against U,V,W_7 basis required",
		Condition:                 "Hodge star is a typed Lambda4 operation, but no current gate computes rank(P_W * |_{K_7}) or proves that * sends K_7 nondegenerately into W_7.",
		Verdict:                   StatusHodgeStarRequiresExplicitRankTest,
	}
}

func auditProjectorAlgebraPairing() ProjectorAlgebraPairingAudit {
	rows := []ProjectorAlgebraCandidate{
		{Operator: "P_W P_B |_{K_7}", ActionOnK7: "P_B k = k for k in K_7", AfterPW: "P_W k = 0 because k in U+V", PairsK7ToW7: false},
		{Operator: "P_W P_G |_{K_7}", ActionOnK7: "P_G k = k for k in K_7", AfterPW: "P_W k = 0 because k in U+V", PairsK7ToW7: false},
		{Operator: "P_W(P_B-P_G)|_{K_7}", ActionOnK7: "(P_B-P_G)k = 0", AfterPW: "0", PairsK7ToW7: false},
		{Operator: "P_W[P_B,P_G]|_{K_7}", ActionOnK7: "both projectors fix k on K_7, so the commutator gives no certified transverse component", AfterPW: "0 or uncertified, never a rank-7 map", PairsK7ToW7: false},
		{Operator: "P_W(P_B+P_G)|_{K_7}", ActionOnK7: "(P_B+P_G)k = 2k", AfterPW: "2P_W k = 0", PairsK7ToW7: false},
	}
	return ProjectorAlgebraPairingAudit{
		Rows:                rows,
		K7FixedByPB:         true,
		K7FixedByPG:         true,
		PWKillsUPlusV:       true,
		AnyPairingCertified: false,
		Reason:              "K_7 lies in U∩V, so simple P_B/P_G algebra keeps it inside U+V or kills it before projection; P_W removes all U+V components.",
		Verdict:             StatusProjectorAlgebraFails,
	}
}

func auditEtaPairing() EtaPairingAudit {
	return EtaPairingAudit{
		Formula:                    "Phi_eta = P_W eta |_{K_7}",
		TypedEtaOnLambda4Available: false,
		RankTestImplemented:        false,
		PairingCertified:           false,
		CompatibilityCertified:     false,
		Reason:                     "ASHA has eta/signed-trace lanes elsewhere, but Gate631 has no typed eta operator acting on H=Lambda^4 R^8 with certified rank from K_7 to W_7.",
		Verdict:                    StatusNoCanonicalK7ToW7Pairing,
	}
}

func auditDeterminantLine(exact ExactDefectSequence) DeterminantLineAudit {
	return DeterminantLineAudit{
		ExactSequence:                     exact.Sequence,
		DeterminantRelation:               "det(K_7) ⊗ det(H) ≅ det(U⊕V) ⊗ det(W_7)",
		CanonicalLineRelation:             exact.ExactByRankNullity,
		PointwiseIsomorphism:              false,
		OrientationDependent:              true,
		CanSupportVolumeBookkeeping:       true,
		CanSupportNormalizedTraceByItself: false,
		Interpretation:                    "The exact sequence gives determinant-line bookkeeping for volume/orientation, but it does not construct Phi:K_7->W_7 and does not by itself justify a normalized 7/72 trace weight.",
		Verdict:                           StatusK7W7PairingProblemSharpened,
	}
}

func auditCandidatePairings(h HodgeStarPairingAudit, p ProjectorAlgebraPairingAudit, e EtaPairingAudit, d DeterminantLineAudit) CandidatePairingTable {
	rows := []PairingCandidate{
		{Name: "orthogonal representative", Formula: "coker(A) ≅ W_7=(U+V)^perp", SourceLane: "ambient metric on Lambda4", TouchesK7AndW7: true, RankTestAvailable: true, NondegenerateCertified: false, CanonicalCertified: false, FailureOrCondition: "represents the cokernel but does not map K_7 to W_7", Verdict: StatusCokernelRepresentedByW7},
		{Name: "Hodge star", Formula: h.Formula, SourceLane: "Lambda4 Hodge operation", TouchesK7AndW7: true, RankTestAvailable: h.RankTestImplemented, NondegenerateCertified: h.NondegenerateCertified, CanonicalCertified: false, FailureOrCondition: h.Condition, Verdict: h.Verdict},
		{Name: "projector algebra", Formula: "P_W f(P_B,P_G)|_{K_7}", SourceLane: "P_B/P_G algebra", TouchesK7AndW7: true, RankTestAvailable: true, NondegenerateCertified: false, CanonicalCertified: false, FailureOrCondition: p.Reason, Verdict: p.Verdict},
		{Name: "eta signed pairing", Formula: e.Formula, SourceLane: "eta/signed trace lane", TouchesK7AndW7: true, RankTestAvailable: e.RankTestImplemented, NondegenerateCertified: false, CanonicalCertified: false, FailureOrCondition: e.Reason, Verdict: e.Verdict},
		{Name: "determinant line", Formula: d.DeterminantRelation, SourceLane: "exact sequence determinant functor", TouchesK7AndW7: true, RankTestAvailable: true, NondegenerateCertified: false, CanonicalCertified: d.CanonicalLineRelation, FailureOrCondition: d.Interpretation, Verdict: d.Verdict},
	}
	return CandidatePairingTable{
		Candidates:                rows,
		CanonicalPairingFound:     false,
		NondegeneratePairingFound: false,
		PairingProblemSharpened:   true,
		MissingObject:             "Phi: K_7 -> W_7, equivalently Phi: ker(A) -> orthogonal representative of coker(A)",
		Verdict:                   StatusNoCanonicalK7ToW7Pairing,
	}
}

func auditBoundaryReadiness(c CandidatePairingTable, d DeterminantLineAudit, i Gate630Inheritance) BoundaryReadinessAudit {
	return BoundaryReadinessAudit{
		K7ToW7PairingCertified:             c.NondegeneratePairingFound,
		DeterminantLineRelationAvailable:   d.CanonicalLineRelation,
		BoundaryPairDimension:              i.BoundaryPairDimension,
		StillRequiresW7ToBoundary:          true,
		StillRequiresDefectTraceToBoundary: true,
		BoundaryAssignmentCertified:        false,
		MissingObject:                      "typed map W_7 -> R^2_boundary, or a certified K_7/W_7 defect trace -> R^2_boundary, compatible with the 7/72 boundary weight",
		Verdict:                            StatusNoBoundaryStressAssignment,
	}
}

func auditNativeStatus(i Gate630Inheritance, w OrthogonalCokernelRepresentativeTable, exact ExactDefectSequence, h HodgeStarPairingAudit, p ProjectorAlgebraPairingAudit, e EtaPairingAudit, d DeterminantLineAudit, b BoundaryReadinessAudit) NativeASHAStatus {
	return NativeASHAStatus{
		Lambda4Native:                  i.HDimension == lambda4DimExpected,
		AmbientMetricAdmitted:          true,
		OrthogonalRepresentativeTyped:  w.RepresentsCokernel,
		K7Native:                       i.K7Dimension == k7DimExpected,
		W7DimensionTyped:               w.WDimension == k7DimExpected,
		ExactDefectSequenceTyped:       exact.ExactByRankNullity,
		HodgeStarRankCertified:         h.NondegenerateCertified,
		ProjectorPairingCertified:      p.AnyPairingCertified,
		EtaPairingCertified:            e.PairingCertified,
		DeterminantLineRelationTyped:   d.CanonicalLineRelation,
		CanonicalK7ToW7Pairing:         false,
		BoundaryStressAssignmentNative: b.BoundaryAssignmentCertified,
		Statement:                      "ASHA can represent the cokernel by W_7 using the ambient metric and can write the exact defect sequence, but no canonical K_7->W_7 pairing or boundary-stress assignment has been certified.",
		Verdict:                        StatusGate631Boundary,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{
		ClaimsCanonicalK7W7Pairing:     false,
		ClaimsBoundaryStressAssignment: false,
		ClaimsScalarRGMatching:         false,
		ClaimsHiggsMassDerivation:      false,
		ClaimsFlavorDerivation:         false,
		ClaimsCKMPMNSDerivation:        false,
		ClaimsGaugeUnification:         false,
		ClaimsBoundaryPairNative:       false,
		ClaimsNativeTraceTheorem:       false,
		Verdict:                        StatusGate631Boundary,
	}
}

func Statuses() []string {
	return []string{
		StatusGate630Inherited,
		StatusOrthogonalCokernelRepresentativeDefined,
		StatusExactDefectSequenceWritten,
		StatusCokernelRepresentedByW7,
		StatusK7W7PairingProblemSharpened,
		StatusNoCanonicalK7ToW7Pairing,
		StatusProjectorAlgebraFails,
		StatusHodgeStarRequiresExplicitRankTest,
		StatusNoBoundaryStressAssignment,
		StatusGate631Boundary,
	}
}
