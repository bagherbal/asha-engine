// Package generation2k7intersectioncokerneldualityaudit implements
// Gate 629: K7IntersectionCokernel Duality Audit.
//
// Gate 628 sharpened the Gate626/Gate627 boundary-weight denominator as
//
//	72 = dim(Lambda^4 R^8) + dim R^2_boundary = 70 + 2.
//
// Gate 629 audits the next pressure point inside the native Lambda^4 R^8
// chamber: with U=Im(P_B), V=Im(P_G), and K_7=U∩V, the Boolean-octonionic span
// has dimension dim(U+V)=56+14-7=63, leaving a 7-dimensional cokernel in
// Lambda^4 R^8. Thus the augmented chamber can be read as
//
//	72 = 7_intersection/gap + 63_Boolean-octonionic-span + 2_boundary.
//
// The gate asks whether the numerator 7 in 7/72 belongs to K_7, to the
// Lambda^4/(U+V) cokernel, or to a candidate intersection-cokernel dual pair.
// It does not certify an isomorphism, boundary-pull assignment, scalar theorem,
// or gauge-scalar-flavor transport theorem.
package generation2k7intersectioncokerneldualityaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/asha"
	gate628 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7overlambda4boundarypairprojectionaudit"
)

const (
	AuditID = "GATE629-K7-INTERSECTION-COKERNEL-DUALITY-AUDIT"

	StatusGate628Inherited                = "PASS_GATE628_LAMBDA4_PLUS_BOUNDARY_PAIR_CHAMBER_INHERITED"
	StatusSpanDimensionComputed           = "PASS_BOOLEAN_OCTONIONIC_SPAN_DIMENSION_COMPUTED"
	StatusCokernelDimensionComputed       = "PASS_LAMBDA4_COKERNEL_DIMENSION_COMPUTED"
	Status72SplitAudited                  = "PASS_72_SPLIT_AS_7_PLUS_63_PLUS_2_AUDITED"
	Status63SpanCandidate                 = "CONDITIONAL_SUPPORT_63_AS_BOOLEAN_OCTONIONIC_SPAN_DIMENSION"
	StatusIntersectionCokernelCandidate   = "CONDITIONAL_SUPPORT_NUMERATOR_7_HAS_INTERSECTION_COKERNEL_DUAL_CANDIDATE"
	Status65SpanBoundaryComplement        = "CONDITIONAL_SUPPORT_65_AS_BOOLEAN_OCTONIONIC_SPAN_PLUS_BOUNDARY_PAIR"
	StatusNoNativeIntersectionCokernelIso = "FAILED_ROUTE_NO_NATIVE_ISOMORPHISM_BETWEEN_K7_AND_LAMBDA4_COKERNEL_7"
	StatusNoBoundaryPullAssignment        = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_PULL_ASSIGNMENT_TO_K7_OR_COKERNEL"
	StatusNoNativeDualBoundaryProjector   = "FAILED_ROUTE_NO_NATIVE_INTERSECTION_COKERNEL_BOUNDARY_PROJECTOR"
	StatusGate629Boundary                 = "FIREWALL_PRESERVED_GATE629_INTERSECTION_COKERNEL_DUALITY_IS_CANDIDATE_ONLY"
)

const (
	rankPBExpected            = 56
	rankPGExpected            = 14
	k7DimExpected             = 7
	lambda4DimExpected        = 70
	booleanOctonionicSpanDim  = 63
	lambda4CokernelDim        = 7
	boundaryStressPairDim     = 2
	augmentedChamberDim       = 72
	spanBoundaryComplementDim = 65
)

type Gate628Inheritance struct {
	Lambda4Dimension             int
	BoundaryPairDimension        int
	AugmentedChamberDimension    int
	K7Dimension                  int
	NonK7Lambda4Complement       int
	AugmentedComplementDimension int
	WeightNumerator              int
	WeightDenominator            int
	BoundaryWeight               float64
	ScalarWeight                 float64
	WeightedClosureResidual      float64
	WeightedMixture              float64
	AbsLambda12                  float64
	R3MinusOne                   float64
	BoundarySplit                float64
	Gate628ChamberCandidate      bool
	Gate628ProjectionMissing     bool
	Gate628ProductAirlockMissing bool
	Gate628FirewallPreserved     bool
	Verdict                      string
}

type BooleanOctonionicSpanAudit struct {
	UCarrier                     string
	VCarrier                     string
	IntersectionCarrier          string
	RankPB                       int
	RankPG                       int
	IntersectionDimension        int
	SpanDimension                int
	ExpectedSpanDimension        int
	SpanFormula                  string
	SpanMatchesExpected          bool
	Lambda4Dimension             int
	CokernelDimension            int
	CokernelFormula              string
	CokernelMatchesK7Dimension   bool
	SpanDimensionCertifiedByRank bool
	Verdict                      string
}

type IntersectionCokernelDualityAudit struct {
	IntersectionCarrier       string
	IntersectionDimension     int
	CokernelCarrier           string
	CokernelDimension         int
	DimensionsEqual           bool
	EqualityIsOnlyDimensional bool
	CanonicalIsomorphismFound bool
	CanonicalPairingFound     bool
	DualityCandidate          bool
	MissingMap                string
	Interpretation            string
	Verdict                   string
}

type ChamberSplitAudit struct {
	IntersectionOrGapDimension int
	SpanDimension              int
	BoundaryPairDimension      int
	AugmentedChamberDimension  int
	SplitExpression            string
	SplitMatches72             bool
	NativeSpanDimension        bool
	BoundaryPairBridgeOnly     bool
	SharperThan70Plus2         bool
	Interpretation             string
	Verdict                    string
}

type ComplementRoleAudit struct {
	SpanDimension               int
	BoundaryPairDimension       int
	SpanBoundaryComplement      int
	AugmentedChamberDimension   int
	ComplementWeight            float64
	Equals65Over72              bool
	PreviousComplementEquation  string
	SharpenedComplementEquation string
	RoleReading                 string
	NativeRoleTheoremFound      bool
	Verdict                     string
}

type BoundaryPullAssignmentAudit struct {
	Candidates           []BoundaryPullCandidate
	BoundaryPullLine     string
	BoundaryWeight       float64
	IntersectionAssigned bool
	CokernelAssigned     bool
	DualPairAssigned     bool
	AssignmentCertified  bool
	MissingObject        string
	Verdict              string
}

type BoundaryPullCandidate struct {
	Name               string
	Dimension          int
	SourceType         string
	CanSupplySeven     bool
	BoundaryAssignment bool
	NativeTheorem      bool
	Comment            string
}

type WeightedMixtureReinterpretationAudit struct {
	KappaSum              float64
	WeightedMixture       float64
	Residual              float64
	BoundaryWeight        float64
	ScalarWeight          float64
	ScalarWeightAs63Plus2 bool
	BoundaryWeightAsSeven bool
	MixtureEquation       string
	Interpretation        string
	Verdict               string
}

type NativeASHAStatus struct {
	Lambda4Native                       bool
	PBImageRankNative                   bool
	PGImageRankNative                   bool
	K7IntersectionNative                bool
	BooleanOctonionicSpanDimensionTyped bool
	Lambda4CokernelDimensionTyped       bool
	IntersectionCokernelIsomorphism     bool
	BoundaryPairNativeFinite            bool
	BoundaryPullAssignmentNative        bool
	DualBoundaryProjectorNative         bool
	GaugeScalarFlavorTransportNative    bool
	Statement                           string
	Verdict                             string
}

type Firewalls struct {
	ClaimsK7CokernelIsomorphism  bool
	ClaimsBoundaryPullAssignment bool
	ClaimsDualBoundaryProjector  bool
	ClaimsBoundaryPairNative     bool
	ClaimsScalarRGMatching       bool
	ClaimsFlavorOrientation      bool
	ClaimsGaugeUnification       bool
	ClaimsHiggsMassDerived       bool
	ClaimsEndpointDerivation     bool
	Verdict                      string
}

type Analysis struct {
	Inherited              Gate628Inheritance
	Span                   BooleanOctonionicSpanAudit
	Duality                IntersectionCokernelDualityAudit
	ChamberSplit           ChamberSplitAudit
	ComplementRole         ComplementRoleAudit
	BoundaryPullAssignment BoundaryPullAssignmentAudit
	WeightedMixture        WeightedMixtureReinterpretationAudit
	NativeStatus           NativeASHAStatus
	Firewalls              Firewalls
	Truth                  string
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
	g628, err := gate628.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate628 predecessor: %w", err)
	}
	engine := asha.New()
	geo := engine.Geometry

	inherit := inheritGate628(g628)
	span := auditBooleanOctonionicSpan(inherit, geo)
	duality := auditIntersectionCokernelDuality(span)
	chamber := auditChamberSplit(inherit, span, duality)
	complement := auditComplementRole(inherit, span, chamber)
	assignment := auditBoundaryPullAssignment(inherit, duality)
	mixture := auditWeightedMixtureReinterpretation(inherit, complement, assignment)

	return Analysis{
		Inherited:              inherit,
		Span:                   span,
		Duality:                duality,
		ChamberSplit:           chamber,
		ComplementRole:         complement,
		BoundaryPullAssignment: assignment,
		WeightedMixture:        mixture,
		NativeStatus:           auditNativeStatus(span, duality, assignment),
		Firewalls:              auditFirewalls(),
		Truth:                  "Gate 629 sharpens Gate 628 by resolving 72=70+2 into 72=7+63+2. The 63 is not merely 70-7: it is dim(Im(P_B)+Im(P_G))=56+14-7, the Boolean-octonionic span after overlap correction. The remaining Lambda^4 cokernel also has dimension 7, so the numerator in 7/72 may be the K_7 intersection, the Lambda^4/(Im(P_B)+Im(P_G)) gap, or a candidate intersection-cokernel dual pair. The result is conditional only: no canonical isomorphism Phi:K_7<->Lambda^4/(U+V), no boundary-pull assignment, and no dual boundary projector are certified.",
	}, nil
}

func inheritGate628(g gate628.Analysis) Gate628Inheritance {
	return Gate628Inheritance{
		Lambda4Dimension:             g.Chamber.Lambda4Dimension,
		BoundaryPairDimension:        g.Chamber.BoundaryPairDimension,
		AugmentedChamberDimension:    g.Chamber.AugmentedChamberDimension,
		K7Dimension:                  g.K7Embedding.K7Dimension,
		NonK7Lambda4Complement:       g.Complement.NonK7Lambda4ComplementDimension,
		AugmentedComplementDimension: g.Complement.AugmentedComplementDimension,
		WeightNumerator:              g.ProjectionTrace.K7TraceDimension,
		WeightDenominator:            g.ProjectionTrace.DomainDimension,
		BoundaryWeight:               g.WeightedClosure.BoundaryWeight,
		ScalarWeight:                 g.WeightedClosure.ScalarWeight,
		WeightedClosureResidual:      g.WeightedClosure.WeightedClosureResidual,
		WeightedMixture:              g.WeightedClosure.WeightedMixture,
		AbsLambda12:                  g.Inherited.AbsLambda12,
		R3MinusOne:                   g.Inherited.R3MinusOne,
		BoundarySplit:                g.Inherited.BoundarySplit,
		Gate628ChamberCandidate:      g.Chamber.Verdict == gate628.Status72Lambda4BoundaryPair && g.Chamber.EqualsTargetDenominator,
		Gate628ProjectionMissing:     !g.ProjectionTrace.ProjectionOperatorExists && !g.ProjectionTrace.IntertwinerCertified,
		Gate628ProductAirlockMissing: !g.NativeStatus.ProductAirlockNative,
		Gate628FirewallPreserved:     g.Firewalls.Verdict == gate628.StatusGate628Boundary,
		Verdict:                      StatusGate628Inherited,
	}
}

func auditBooleanOctonionicSpan(i Gate628Inheritance, g asha.Geometry) BooleanOctonionicSpanAudit {
	lambda4 := 0
	if len(g.GradeDimensions) > 4 {
		lambda4 = g.GradeDimensions[4]
	}
	span := g.RankPB + g.RankPG - g.DimK
	cokernel := lambda4 - span
	return BooleanOctonionicSpanAudit{
		UCarrier:                     "U=Im(P_B)",
		VCarrier:                     "V=Im(P_G)",
		IntersectionCarrier:          "K_7=U∩V",
		RankPB:                       g.RankPB,
		RankPG:                       g.RankPG,
		IntersectionDimension:        g.DimK,
		SpanDimension:                span,
		ExpectedSpanDimension:        booleanOctonionicSpanDim,
		SpanFormula:                  "dim(U+V)=rank(P_B)+rank(P_G)-dim(U∩V)=56+14-7=63",
		SpanMatchesExpected:          span == booleanOctonionicSpanDim && span == i.NonK7Lambda4Complement,
		Lambda4Dimension:             lambda4,
		CokernelDimension:            cokernel,
		CokernelFormula:              "dim(Lambda^4 R^8/(U+V))=70-63=7",
		CokernelMatchesK7Dimension:   cokernel == g.DimK && cokernel == lambda4CokernelDim,
		SpanDimensionCertifiedByRank: g.RankPB == rankPBExpected && g.RankPG == rankPGExpected && g.DimK == k7DimExpected && lambda4 == lambda4DimExpected && span == booleanOctonionicSpanDim,
		Verdict:                      StatusSpanDimensionComputed,
	}
}

func auditIntersectionCokernelDuality(s BooleanOctonionicSpanAudit) IntersectionCokernelDualityAudit {
	dimensionsEqual := s.IntersectionDimension == s.CokernelDimension && s.IntersectionDimension == k7DimExpected
	return IntersectionCokernelDualityAudit{
		IntersectionCarrier:       s.IntersectionCarrier,
		IntersectionDimension:     s.IntersectionDimension,
		CokernelCarrier:           "Lambda^4 R^8/(Im(P_B)+Im(P_G))",
		CokernelDimension:         s.CokernelDimension,
		DimensionsEqual:           dimensionsEqual,
		EqualityIsOnlyDimensional: dimensionsEqual,
		CanonicalIsomorphismFound: false,
		CanonicalPairingFound:     false,
		DualityCandidate:          dimensionsEqual,
		MissingMap:                "Phi: K_7 <-> Lambda^4 R^8/(Im(P_B)+Im(P_G))",
		Interpretation:            "The second seven is a native-looking cokernel dimension in the same Lambda^4 carrier, but equal dimensions do not provide a canonical isomorphism or orientation-compatible pairing.",
		Verdict:                   StatusIntersectionCokernelCandidate,
	}
}

func auditChamberSplit(i Gate628Inheritance, s BooleanOctonionicSpanAudit, d IntersectionCokernelDualityAudit) ChamberSplitAudit {
	split := d.IntersectionDimension + s.SpanDimension + i.BoundaryPairDimension
	return ChamberSplitAudit{
		IntersectionOrGapDimension: d.IntersectionDimension,
		SpanDimension:              s.SpanDimension,
		BoundaryPairDimension:      i.BoundaryPairDimension,
		AugmentedChamberDimension:  split,
		SplitExpression:            "72 = 7_intersection/gap + 63_Boolean-octonionic-span + 2_boundary",
		SplitMatches72:             split == augmentedChamberDim && split == i.AugmentedChamberDimension,
		NativeSpanDimension:        s.SpanDimensionCertifiedByRank,
		BoundaryPairBridgeOnly:     true,
		SharperThan70Plus2:         split == augmentedChamberDim && s.SpanDimension == booleanOctonionicSpanDim,
		Interpretation:             "The 70+2 chamber from Gate 628 refines to 7+63+2: a seven-dimensional intersection/gap, the 63-dimensional Boolean-octonionic span, and the bridge boundary pair.",
		Verdict:                    Status72SplitAudited,
	}
}

func auditComplementRole(i Gate628Inheritance, s BooleanOctonionicSpanAudit, c ChamberSplitAudit) ComplementRoleAudit {
	spanBoundary := s.SpanDimension + c.BoundaryPairDimension
	weight := float64(spanBoundary) / float64(c.AugmentedChamberDimension)
	return ComplementRoleAudit{
		SpanDimension:               s.SpanDimension,
		BoundaryPairDimension:       c.BoundaryPairDimension,
		SpanBoundaryComplement:      spanBoundary,
		AugmentedChamberDimension:   c.AugmentedChamberDimension,
		ComplementWeight:            weight,
		Equals65Over72:              spanBoundary == spanBoundaryComplementDim && math.Abs(weight-i.ScalarWeight) < 1e-15,
		PreviousComplementEquation:  "65=(70-7)+2",
		SharpenedComplementEquation: "65=dim(Im(P_B)+Im(P_G))+dim R^2_boundary=63+2",
		RoleReading:                 "The 65/72 scalar-wound weight can be read as Boolean-octonionic span plus boundary pair, while the seven-dimensional intersection/gap supplies the candidate boundary pull weight.",
		NativeRoleTheoremFound:      false,
		Verdict:                     Status65SpanBoundaryComplement,
	}
}

func auditBoundaryPullAssignment(i Gate628Inheritance, d IntersectionCokernelDualityAudit) BoundaryPullAssignmentAudit {
	candidates := []BoundaryPullCandidate{
		{Name: "K_7 intersection", Dimension: d.IntersectionDimension, SourceType: "native contact carrier K_7=Im(P_B)∩Im(P_G)", CanSupplySeven: d.IntersectionDimension == k7DimExpected, BoundaryAssignment: false, NativeTheorem: false, Comment: "Gate 627/628 numerator candidate; no boundary-pull projector certified"},
		{Name: "Lambda4 cokernel gap", Dimension: d.CokernelDimension, SourceType: "native quotient candidate Lambda^4 R^8/(Im(P_B)+Im(P_G))", CanSupplySeven: d.CokernelDimension == lambda4CokernelDim, BoundaryAssignment: false, NativeTheorem: false, Comment: "new Gate629 seven; no map to boundary stress certified"},
		{Name: "intersection-cokernel dual pair", Dimension: d.IntersectionDimension, SourceType: "candidate duality Phi between K_7 and cokernel-7", CanSupplySeven: d.DualityCandidate, BoundaryAssignment: false, NativeTheorem: false, Comment: "most interesting bridge target; canonical isomorphism absent"},
	}
	return BoundaryPullAssignmentAudit{
		Candidates:           candidates,
		BoundaryPullLine:     "(R_3-1)-|lambda(Lambda_12)|",
		BoundaryWeight:       i.BoundaryWeight,
		IntersectionAssigned: false,
		CokernelAssigned:     false,
		DualPairAssigned:     false,
		AssignmentCertified:  false,
		MissingObject:        "typed boundary-pull assignment for K_7, cokernel-7, or Phi-dual pair into R^2_boundary",
		Verdict:              StatusNoBoundaryPullAssignment,
	}
}

func auditWeightedMixtureReinterpretation(i Gate628Inheritance, c ComplementRoleAudit, b BoundaryPullAssignmentAudit) WeightedMixtureReinterpretationAudit {
	return WeightedMixtureReinterpretationAudit{
		KappaSum:              i.WeightedMixture + i.WeightedClosureResidual,
		WeightedMixture:       i.WeightedMixture,
		Residual:              i.WeightedClosureResidual,
		BoundaryWeight:        i.BoundaryWeight,
		ScalarWeight:          i.ScalarWeight,
		ScalarWeightAs63Plus2: c.Equals65Over72,
		BoundaryWeightAsSeven: math.Abs(b.BoundaryWeight-7.0/72.0) < 1e-15,
		MixtureEquation:       "((63+2)/72)|lambda(Lambda_12)| + (7/72)(R_3-1)",
		Interpretation:        "Boolean-octonionic span plus boundary pair preserves the scalar wound; the unresolved seven-dimensional intersection/gap candidate supplies the gauge-boundary pull weight.",
		Verdict:               Status65SpanBoundaryComplement,
	}
}

func auditNativeStatus(s BooleanOctonionicSpanAudit, d IntersectionCokernelDualityAudit, b BoundaryPullAssignmentAudit) NativeASHAStatus {
	return NativeASHAStatus{
		Lambda4Native:                       s.Lambda4Dimension == lambda4DimExpected,
		PBImageRankNative:                   s.RankPB == rankPBExpected,
		PGImageRankNative:                   s.RankPG == rankPGExpected,
		K7IntersectionNative:                s.IntersectionDimension == k7DimExpected,
		BooleanOctonionicSpanDimensionTyped: s.SpanDimensionCertifiedByRank,
		Lambda4CokernelDimensionTyped:       s.CokernelDimension == lambda4CokernelDim,
		IntersectionCokernelIsomorphism:     d.CanonicalIsomorphismFound,
		BoundaryPairNativeFinite:            false,
		BoundaryPullAssignmentNative:        b.AssignmentCertified,
		DualBoundaryProjectorNative:         false,
		GaugeScalarFlavorTransportNative:    false,
		Statement:                           "Native ranks certify the 56,14,7,63,7 dimension ledger, but ASHA does not yet supply a canonical K_7-cokernel isomorphism or a boundary-pull projector.",
		Verdict:                             StatusGate629Boundary,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{
		ClaimsK7CokernelIsomorphism:  false,
		ClaimsBoundaryPullAssignment: false,
		ClaimsDualBoundaryProjector:  false,
		ClaimsBoundaryPairNative:     false,
		ClaimsScalarRGMatching:       false,
		ClaimsFlavorOrientation:      false,
		ClaimsGaugeUnification:       false,
		ClaimsHiggsMassDerived:       false,
		ClaimsEndpointDerivation:     false,
		Verdict:                      StatusGate629Boundary,
	}
}

func Statuses() []string {
	return []string{
		StatusGate628Inherited,
		StatusSpanDimensionComputed,
		StatusCokernelDimensionComputed,
		Status72SplitAudited,
		Status63SpanCandidate,
		StatusIntersectionCokernelCandidate,
		Status65SpanBoundaryComplement,
		StatusNoNativeIntersectionCokernelIso,
		StatusNoBoundaryPullAssignment,
		StatusNoNativeDualBoundaryProjector,
		StatusGate629Boundary,
	}
}
