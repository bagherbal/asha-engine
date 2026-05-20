// Package generation2k7kernelcokernelindexzeroaudit implements
// Gate 630: K7 Kernel-Cokernel Index-Zero Audit.
//
// Gate 629 exposed two seven-dimensional pressure objects inside the native
// Lambda^4 R^8 chamber: the intersection K_7=Im(P_B)∩Im(P_G) and the quotient
// gap Lambda^4 R^8/(Im(P_B)+Im(P_G)). Gate 630 sharpens this from a loose
// equality of dimensions into a square finite operator audit. With
// U=Im(P_B), V=Im(P_G), and dim(U⊕V)=56+14=70=dim Lambda^4 R^8, define
//
//	A: U⊕V -> Lambda^4 R^8,      A(u,v)=u+v.
//
// Then ker(A)≅K_7, im(A)=U+V has dimension 63, coker(A) has dimension 7,
// and index(A)=dim ker(A)-dim coker(A)=0. The gate asks whether ASHA supplies
// a canonical ker-coker pairing or a native boundary-stress assignment from
// this balanced K7 defect. It does not certify such a pairing, scalar theorem,
// boundary theorem, gauge unification theorem, or endpoint derivation.
package generation2k7kernelcokernelindexzeroaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/asha"
	gate629 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7intersectioncokerneldualityaudit"
)

const (
	AuditID = "GATE630-K7-KERNEL-COKERNEL-INDEX-ZERO-AUDIT"

	StatusGate629Inherited                 = "PASS_GATE629_INTERSECTION_COKERNEL_DUAL_CANDIDATE_INHERITED"
	StatusAdditionMapDefined               = "PASS_ADDITION_MAP_A_DEFINED"
	StatusKernelAIsK7                      = "PASS_KERNEL_A_IS_K7"
	StatusCokernelADim7                    = "PASS_COKERNEL_A_HAS_DIMENSION_7"
	StatusIndexZeroComputed                = "PASS_INDEX_ZERO_BOOLEAN_OCTONIONIC_DEFECT_COMPUTED"
	StatusK7BlockCompressionComputed       = "PASS_K7_BLOCK_COMPRESSION_COMPUTED"
	StatusK7DefectBlockCandidate           = "CONDITIONAL_SUPPORT_K7_DEFECT_BLOCK_SOURCE_FOR_7_OVER_72"
	StatusNoCanonicalKerCokerPairing       = "FAILED_ROUTE_NO_CANONICAL_KERNEL_COKERNEL_PAIRING_YET"
	StatusNoNativeBoundaryStressFromDefect = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_STRESS_ASSIGNMENT_FROM_INDEX_ZERO_DEFECT"
	StatusGate630Boundary                  = "FIREWALL_PRESERVED_GATE630_DEFECT_PAIRING_IS_CANDIDATE_ONLY"
)

const (
	rankPBExpected        = 56
	rankPGExpected        = 14
	k7DimExpected         = 7
	lambda4DimExpected    = 70
	booleanOctonionicSpan = 63
	boundaryPairDim       = 2
	augmentedChamberDim   = 72
	expectedFredholmIndex = 0
)

type Gate629Inheritance struct {
	UDimension                       int
	VDimension                       int
	DirectSumDimension               int
	Lambda4Dimension                 int
	IntersectionDimension            int
	SpanDimension                    int
	CokernelDimension                int
	BoundaryPairDimension            int
	AugmentedChamberDimension        int
	BoundaryWeight                   float64
	WeightedClosureResidual          float64
	Gate629DualCandidate             bool
	Gate629IsomorphismMissing        bool
	Gate629BoundaryAssignmentMissing bool
	Gate629FirewallPreserved         bool
	Verdict                          string
}

type AdditionMapAudit struct {
	MapName            string
	Domain             string
	Codomain           string
	Formula            string
	DomainDimension    int
	CodomainDimension  int
	SquareOperator     bool
	KernelExpression   string
	KernelDimension    int
	KernelIsK7         bool
	ImageExpression    string
	ImageDimension     int
	ImageIsSpan        bool
	CokernelExpression string
	CokernelDimension  int
	CokernelMatchesK7  bool
	RankDefect         int
	Index              int
	IndexZero          bool
	Verdict            string
}

type KernelCokernelDefectAudit struct {
	KernelCarrier       string
	KernelDimension     int
	CokernelCarrier     string
	CokernelDimension   int
	DefectsBalanced     bool
	Index               int
	FredholmAnalogyOnly bool
	CandidateDefectPair bool
	MissingPairing      string
	Interpretation      string
	Verdict             string
}

type K7BlockCompressionAudit struct {
	K7BlockDimension         int
	PBBlocks                 int
	PGBlocks                 int
	SpanBlocks               int
	Lambda4Blocks            int
	BoundaryCoordinates      int
	AugmentedExpression      string
	BoundaryWeightExpression string
	BoundaryWeight           float64
	CompressionExact         bool
	DefectBlockCandidate     bool
	Interpretation           string
	Verdict                  string
}

type PairingCandidateAudit struct {
	Candidates                []PairingCandidate
	CanonicalPairingFound     bool
	MetricPairingCertified    bool
	HodgeStarPairingCertified bool
	EtaPairingCertified       bool
	ProjectorPairingCertified bool
	MissingObject             string
	Verdict                   string
}

type PairingCandidate struct {
	Name              string
	SourceLane        string
	CouldTouchDefects bool
	Certified         bool
	FailureReason     string
}

type BoundaryStressAssignmentAudit struct {
	DefectBlockCanSupplySeven bool
	BoundaryPairDimension     int
	BoundaryWeight            float64
	BoundaryStressLine        string
	AssignmentCertified       bool
	NativeTransportTheorem    bool
	MissingObject             string
	Verdict                   string
}

type NativeASHAStatus struct {
	Lambda4Native                  bool
	UImageRankNative               bool
	VImageRankNative               bool
	K7IntersectionNative           bool
	AdditionMapTyped               bool
	KernelDimensionTyped           bool
	CokernelDimensionTyped         bool
	IndexZeroTyped                 bool
	CanonicalKernelCokernelPairing bool
	BoundaryStressAssignmentNative bool
	K7DefectBoundaryTraceTheorem   bool
	Statement                      string
	Verdict                        string
}

type Firewalls struct {
	ClaimsCanonicalPairing         bool
	ClaimsBoundaryStressAssignment bool
	ClaimsK7DefectTraceTheorem     bool
	ClaimsBoundaryPairNative       bool
	ClaimsScalarRGMatching         bool
	ClaimsFlavorOrientation        bool
	ClaimsGaugeUnification         bool
	ClaimsHiggsMassDerived         bool
	ClaimsEndpointDerivation       bool
	Verdict                        string
}

type Analysis struct {
	Inherited          Gate629Inheritance
	AdditionMap        AdditionMapAudit
	Defect             KernelCokernelDefectAudit
	BlockCompression   K7BlockCompressionAudit
	Pairing            PairingCandidateAudit
	BoundaryAssignment BoundaryStressAssignmentAudit
	NativeStatus       NativeASHAStatus
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
	g629, err := gate629.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate629 predecessor: %w", err)
	}
	engine := asha.New()
	geo := engine.Geometry

	inherited := inheritGate629(g629, geo)
	additionMap := auditAdditionMap(inherited)
	defect := auditKernelCokernelDefect(additionMap)
	blocks := auditK7BlockCompression(inherited, additionMap, defect)
	pairing := auditPairingCandidates(defect)
	boundary := auditBoundaryStressAssignment(inherited, blocks, pairing)

	return Analysis{
		Inherited:          inherited,
		AdditionMap:        additionMap,
		Defect:             defect,
		BlockCompression:   blocks,
		Pairing:            pairing,
		BoundaryAssignment: boundary,
		NativeStatus:       auditNativeStatus(inherited, additionMap, pairing, boundary),
		Firewalls:          auditFirewalls(),
		Truth:              "Gate 630 sharpens Gate 629 from two equal seven-dimensional objects into a square finite operator A:Im(P_B)⊕Im(P_G)->Lambda^4 R^8 with A(u,v)=u+v. The operator is rank-defective in a balanced way: ker(A)≅K_7 has dimension 7, coker(A)=Lambda^4 R^8/(Im(P_B)+Im(P_G)) has dimension 7, and index(A)=0. This compresses the native chamber as 56=8 K7-blocks, 14=2 K7-blocks, 63=9 K7-blocks, 70=10 K7-blocks, so 7/72 can be read as one K7 defect block over ten finite K7 blocks plus two boundary coordinates. The result remains bridge-only: no canonical ker-coker pairing and no native boundary-stress assignment from the index-zero defect are certified.",
	}, nil
}

func inheritGate629(g gate629.Analysis, geo asha.Geometry) Gate629Inheritance {
	lambda4 := 0
	if len(geo.GradeDimensions) > 4 {
		lambda4 = geo.GradeDimensions[4]
	}
	return Gate629Inheritance{
		UDimension:                       geo.RankPB,
		VDimension:                       geo.RankPG,
		DirectSumDimension:               geo.RankPB + geo.RankPG,
		Lambda4Dimension:                 lambda4,
		IntersectionDimension:            geo.DimK,
		SpanDimension:                    g.Span.SpanDimension,
		CokernelDimension:                g.Span.CokernelDimension,
		BoundaryPairDimension:            g.ChamberSplit.BoundaryPairDimension,
		AugmentedChamberDimension:        g.ChamberSplit.AugmentedChamberDimension,
		BoundaryWeight:                   g.WeightedMixture.BoundaryWeight,
		WeightedClosureResidual:          g.WeightedMixture.Residual,
		Gate629DualCandidate:             g.Duality.DualityCandidate,
		Gate629IsomorphismMissing:        !g.Duality.CanonicalIsomorphismFound,
		Gate629BoundaryAssignmentMissing: !g.BoundaryPullAssignment.AssignmentCertified,
		Gate629FirewallPreserved:         g.Firewalls.Verdict == gate629.StatusGate629Boundary,
		Verdict:                          StatusGate629Inherited,
	}
}

func auditAdditionMap(i Gate629Inheritance) AdditionMapAudit {
	imageDim := i.SpanDimension
	kernelDim := i.DirectSumDimension - imageDim
	cokernelDim := i.Lambda4Dimension - imageDim
	index := kernelDim - cokernelDim
	return AdditionMapAudit{
		MapName:            "A",
		Domain:             "U⊕V = Im(P_B)⊕Im(P_G)",
		Codomain:           "Lambda^4 R^8",
		Formula:            "A(u,v)=u+v",
		DomainDimension:    i.DirectSumDimension,
		CodomainDimension:  i.Lambda4Dimension,
		SquareOperator:     i.DirectSumDimension == i.Lambda4Dimension && i.DirectSumDimension == lambda4DimExpected,
		KernelExpression:   "ker(A)={(k,-k):k∈U∩V}≅K_7",
		KernelDimension:    kernelDim,
		KernelIsK7:         kernelDim == i.IntersectionDimension && kernelDim == k7DimExpected,
		ImageExpression:    "im(A)=U+V",
		ImageDimension:     imageDim,
		ImageIsSpan:        imageDim == booleanOctonionicSpan,
		CokernelExpression: "coker(A)=Lambda^4 R^8/(U+V)",
		CokernelDimension:  cokernelDim,
		CokernelMatchesK7:  cokernelDim == i.IntersectionDimension && cokernelDim == k7DimExpected,
		RankDefect:         i.CodomainDefect(imageDim),
		Index:              index,
		IndexZero:          index == expectedFredholmIndex,
		Verdict:            StatusAdditionMapDefined,
	}
}

func (i Gate629Inheritance) CodomainDefect(imageDim int) int { return i.Lambda4Dimension - imageDim }

func auditKernelCokernelDefect(a AdditionMapAudit) KernelCokernelDefectAudit {
	balanced := a.KernelDimension == a.CokernelDimension && a.IndexZero
	return KernelCokernelDefectAudit{
		KernelCarrier:       a.KernelExpression,
		KernelDimension:     a.KernelDimension,
		CokernelCarrier:     a.CokernelExpression,
		CokernelDimension:   a.CokernelDimension,
		DefectsBalanced:     balanced,
		Index:               a.Index,
		FredholmAnalogyOnly: true,
		CandidateDefectPair: balanced,
		MissingPairing:      "Phi: ker(A) -> coker(A)",
		Interpretation:      "A is a square rank-defective finite addition operator with a balanced K7 kernel defect and Lambda4 cokernel defect; the zero index is typed arithmetic, not a certified Fredholm theorem over boundary stress.",
		Verdict:             StatusIndexZeroComputed,
	}
}

func auditK7BlockCompression(i Gate629Inheritance, a AdditionMapAudit, d KernelCokernelDefectAudit) K7BlockCompressionAudit {
	block := i.IntersectionDimension
	pbBlocks := i.UDimension / block
	pgBlocks := i.VDimension / block
	spanBlocks := i.SpanDimension / block
	lambda4Blocks := i.Lambda4Dimension / block
	boundaryWeight := float64(block) / float64(i.AugmentedChamberDimension)
	return K7BlockCompressionAudit{
		K7BlockDimension:         block,
		PBBlocks:                 pbBlocks,
		PGBlocks:                 pgBlocks,
		SpanBlocks:               spanBlocks,
		Lambda4Blocks:            lambda4Blocks,
		BoundaryCoordinates:      i.BoundaryPairDimension,
		AugmentedExpression:      "72 = 10*K7 + 2 = 70 + 2",
		BoundaryWeightExpression: "7/72 = one K7 defect block / (10 finite K7 blocks + 2 boundary coordinates)",
		BoundaryWeight:           boundaryWeight,
		CompressionExact:         block == k7DimExpected && i.UDimension == 8*block && i.VDimension == 2*block && i.SpanDimension == 9*block && i.Lambda4Dimension == 10*block && i.AugmentedChamberDimension == 10*block+i.BoundaryPairDimension && a.KernelIsK7 && d.DefectsBalanced,
		DefectBlockCandidate:     d.CandidateDefectPair && math.Abs(boundaryWeight-i.BoundaryWeight) < 1e-15,
		Interpretation:           "The finite chamber compresses into K7 blocks: rank(P_B)=8K7, rank(P_G)=2K7, dim(U+V)=9K7, dim(Lambda4)=10K7. Gate626's 7/72 can therefore be read as one balanced K7 defect block over the augmented 10K7+2 bridge chamber.",
		Verdict:                  StatusK7DefectBlockCandidate,
	}
}

func auditPairingCandidates(d KernelCokernelDefectAudit) PairingCandidateAudit {
	candidates := []PairingCandidate{
		{Name: "orthogonal complement / metric pairing", SourceLane: "Euclidean Lambda^4 R^8 linear algebra", CouldTouchDefects: true, Certified: false, FailureReason: "a metric can choose complements, but Gate630 has no canonical map from ker(A) to the quotient coker(A) compatible with P_B/P_G and boundary orientation"},
		{Name: "Hodge-star pairing", SourceLane: "Lambda^4 self-dual/anti-self-dual structure candidate", CouldTouchDefects: true, Certified: false, FailureReason: "no theorem identifies K_7 with the Lambda4 quotient through Hodge star"},
		{Name: "eta-signed pairing", SourceLane: "ASHA eta/topological trace lane", CouldTouchDefects: true, Certified: false, FailureReason: "eta traces certify cancellations elsewhere but do not provide the missing ker-coker isomorphism"},
		{Name: "projector algebra pairing", SourceLane: "P_B/P_G rank and overlap algebra", CouldTouchDefects: true, Certified: false, FailureReason: "the rank ledger gives kernel and cokernel dimensions only; it does not construct Phi:ker(A)->coker(A)"},
	}
	return PairingCandidateAudit{
		Candidates:                candidates,
		CanonicalPairingFound:     false,
		MetricPairingCertified:    false,
		HodgeStarPairingCertified: false,
		EtaPairingCertified:       false,
		ProjectorPairingCertified: false,
		MissingObject:             d.MissingPairing,
		Verdict:                   StatusNoCanonicalKerCokerPairing,
	}
}

func auditBoundaryStressAssignment(i Gate629Inheritance, b K7BlockCompressionAudit, p PairingCandidateAudit) BoundaryStressAssignmentAudit {
	return BoundaryStressAssignmentAudit{
		DefectBlockCanSupplySeven: b.DefectBlockCandidate,
		BoundaryPairDimension:     i.BoundaryPairDimension,
		BoundaryWeight:            b.BoundaryWeight,
		BoundaryStressLine:        "(R_3-1)-|lambda(Lambda_12)|",
		AssignmentCertified:       false,
		NativeTransportTheorem:    false,
		MissingObject:             "typed map from balanced K7 index-zero defect to the R^2_boundary stress pair with normalized trace 7/72",
		Verdict:                   StatusNoNativeBoundaryStressFromDefect,
	}
}

func auditNativeStatus(i Gate629Inheritance, a AdditionMapAudit, p PairingCandidateAudit, b BoundaryStressAssignmentAudit) NativeASHAStatus {
	return NativeASHAStatus{
		Lambda4Native:                  i.Lambda4Dimension == lambda4DimExpected,
		UImageRankNative:               i.UDimension == rankPBExpected,
		VImageRankNative:               i.VDimension == rankPGExpected,
		K7IntersectionNative:           i.IntersectionDimension == k7DimExpected,
		AdditionMapTyped:               a.SquareOperator,
		KernelDimensionTyped:           a.KernelIsK7,
		CokernelDimensionTyped:         a.CokernelMatchesK7,
		IndexZeroTyped:                 a.IndexZero,
		CanonicalKernelCokernelPairing: p.CanonicalPairingFound,
		BoundaryStressAssignmentNative: b.AssignmentCertified,
		K7DefectBoundaryTraceTheorem:   false,
		Statement:                      "ASHA now has a typed square addition map and a balanced 7/7 index-zero defect, but it still lacks a canonical ker-coker pairing and a boundary-stress assignment theorem.",
		Verdict:                        StatusGate630Boundary,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{
		ClaimsCanonicalPairing:         false,
		ClaimsBoundaryStressAssignment: false,
		ClaimsK7DefectTraceTheorem:     false,
		ClaimsBoundaryPairNative:       false,
		ClaimsScalarRGMatching:         false,
		ClaimsFlavorOrientation:        false,
		ClaimsGaugeUnification:         false,
		ClaimsHiggsMassDerived:         false,
		ClaimsEndpointDerivation:       false,
		Verdict:                        StatusGate630Boundary,
	}
}

func Statuses() []string {
	return []string{
		StatusGate629Inherited,
		StatusAdditionMapDefined,
		StatusKernelAIsK7,
		StatusCokernelADim7,
		StatusIndexZeroComputed,
		StatusK7BlockCompressionComputed,
		StatusK7DefectBlockCandidate,
		StatusNoCanonicalKerCokerPairing,
		StatusNoNativeBoundaryStressFromDefect,
		StatusGate630Boundary,
	}
}
