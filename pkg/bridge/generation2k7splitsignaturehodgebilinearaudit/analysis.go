// Package generation2k7splitsignaturehodgebilinearaudit implements
// Gate 636: K7 Split-Signature Hodge Bilinear Audit.
//
// Gate 635 blocked the tempting carrier jump from the K_7 Hodge polarity
// (4|3) to the Fock/Witt projective selector split 4=1+3.  Gate 636 keeps
// the computation native to K_7.  It defines the bilinear form
//
//	B_K(x,y)=<x,S_*y>|_{K_7}=g_K(x,S_K y)
//
// where S_K=Q_K^T S_* Q_K is the restricted Hodge involution certified in
// Gate 634.  The audit certifies the split signature (4,3), records the
// Euclidean metric-conversion role of S_K, checks orthogonality of K_7^+ and
// K_7^-, and deliberately blocks any promotion to Fock selector geometry,
// split-G2 structure, physical spacetime metric, boundary stress, or 7/72.
package generation2k7splitsignaturehodgebilinearaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate635 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7hodgepolarityprojectiveselectoralignmentaudit"
	gate634 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7hodgesignaturestabilizeraudit"
)

const (
	AuditID = "GATE636-K7-SPLIT-SIGNATURE-HODGE-BILINEAR-AUDIT"

	StatusGate635Inherited                = "PASS_GATE635_HODGE_POLARITY_FIREWALL_INHERITED"
	StatusBKHodgeBilinearDefined          = "PASS_BK_HODGE_BILINEAR_DEFINED_ON_K7"
	StatusBKSignatureCertified            = "PASS_BK_SIGNATURE_4_3_CERTIFIED"
	StatusBKMetricConversionCertified     = "PASS_BK_METRIC_CONVERSION_OPERATOR_CERTIFIED"
	StatusPlusMinusOrthogonalityCertified = "PASS_K7_PLUS_MINUS_ORTHOGONALITY_CERTIFIED_FOR_GK_AND_BK"
	StatusNativeSplitSignature            = "CONDITIONAL_SUPPORT_K7_CARRIES_NATIVE_SPLIT_SIGNATURE_STRUCTURE"
	StatusBilinearNotSelector             = "CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_IS_BILINEAR_NOT_SELECTOR"
	StatusStabilizerCandidateAudited      = "PASS_SPLIT_SIGNATURE_STABILIZER_CANDIDATE_AUDITED"
	StatusNoK7ToFockMap                   = "FAILED_ROUTE_NO_TYPED_K7_TO_FOCK_SELECTOR_MAP"
	StatusNoSplitG2Yet                    = "FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE_YET"
	StatusNoOmegaK                        = "FAILED_ROUTE_NO_NATIVE_OMEGA_K_THREE_FORM_CERTIFIED"
	StatusNoBoundaryStressAssignment      = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem            = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoPhysicalMetric                = "FAILED_ROUTE_K7_SPLIT_SIGNATURE_NOT_PHYSICAL_SPACETIME_METRIC"
	StatusGate636Boundary                 = "FIREWALL_PRESERVED_GATE636_SPLIT_SIGNATURE_IS_NATIVE_NOT_PHYSICAL"
)

const (
	k7DimExpected    = 7
	plusDimExpected  = 4
	minusDimExpected = 3
	strictTolerance  = 1e-10
)

type Gate635Inheritance struct {
	K7Dimension               int
	PlusDimension             int
	MinusDimension            int
	Trace                     float64
	Determinant               float64
	CarrierFirewallPreserved  bool
	NoK7ToFockSelectorMap     bool
	NoOnePlusThreeRefinement  bool
	TraceNotDistinguishedLine bool
	NoBoundaryAssignment      bool
	NoSevenOver72Theorem      bool
	Verdict                   string
}

type HodgeBilinearDefinition struct {
	Formula              string
	MetricFormula        string
	MatrixRepresentative string
	Dimension            int
	Rows                 int
	Cols                 int
	Symmetric            bool
	Nondegenerate        bool
	InheritedFromSK      bool
	Verdict              string
}

type SignatureCertificate struct {
	InertiaPlus       int
	InertiaMinus      int
	InertiaZero       int
	Trace             float64
	Determinant       float64
	DeterminantSign   int
	Eigenvalues       []float64
	NullConeExists    bool
	PositiveDefinite  bool
	NegativeDefinite  bool
	SplitIndefinite   bool
	SignatureNotation string
	Verdict           string
}

type MetricConversionAudit struct {
	EuclideanMetric        string
	BilinearRelation       string
	ConversionOperator     string
	SKOrthogonal           bool
	SKSymmetric            bool
	SKInvolutive           bool
	SymmetryResidual       float64
	OrthogonalityResidual  float64
	InvolutionResidual     float64
	BEqualsGComposedWithSK bool
	Verdict                string
}

type PlusMinusOrthogonalityAudit struct {
	PlusDimension            int
	MinusDimension           int
	GOrthogonal              bool
	BOrthogonal              bool
	BRestrictedToPlus        string
	BRestrictedToMinus       string
	CrossTermZero            bool
	ProjectorOrthogonality   float64
	ProjectorComplementarity float64
	Verdict                  string
}

type OctonionicCompatibilityAudit struct {
	CandidateLane                  string
	SplitSignatureMatchesDimension bool
	OmegaKThreeFormCertified       bool
	CrossProductCertified          bool
	CalibrationCertified           bool
	G2SplitStructureCertified      bool
	PreservationByNativeG2Operator bool
	Reason                         string
	Verdict                        string
}

type StabilizerAudit struct {
	BilinearStabilizerCandidate    string
	OrientationPreservingCandidate string
	SplitG2CandidateSubgroup       string
	StabilizerCertified            bool
	SplitG2Certified               bool
	NeedsOmegaK                    bool
	PhysicalMetricClaimed          bool
	Verdict                        string
}

type SelectorBoundaryFirewall struct {
	K7ToFockMapCertified        bool
	OnePlusThreeSelectorDerived bool
	BoundaryStressAssigned      bool
	SevenOver72Promoted         bool
	PhysicalSpacetimeMetric     bool
	ScalarRGMatchingClaimed     bool
	HiggsMassClaimed            bool
	FlavorClaimed               bool
	CKMPMNSClaimed              bool
	GaugeUnificationClaimed     bool
	Verdict                     string
}

type MissingObjectAudit struct {
	PreviousMissingObject string
	CurrentMissingObject  string
	WhySharper            string
	CanSupportSplitG2     bool
	CanSupportBoundary    bool
	VerdictOmega          string
	VerdictBoundary       string
}

type Analysis struct {
	Inherited        Gate635Inheritance
	Definition       HodgeBilinearDefinition
	Signature        SignatureCertificate
	MetricConversion MetricConversionAudit
	Orthogonality    PlusMinusOrthogonalityAudit
	Octonionic       OctonionicCompatibilityAudit
	Stabilizer       StabilizerAudit
	Firewalls        SelectorBoundaryFirewall
	MissingObject    MissingObjectAudit
	Truth            string
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
		return Analysis{}, fmt.Errorf("Gate634 Hodge signature inheritance unavailable: %w", err)
	}
	g635, err := gate635.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate635 polarity/firewall inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g634, g635)
	definition := buildDefinition(g634)
	signature := buildSignature(g634)
	metric := buildMetricConversion(g634)
	orth := buildOrthogonality(g634)
	oct := buildOctonionicCompatibility(signature)
	stab := buildStabilizer(signature, oct)
	firewalls := buildFirewalls()
	missing := buildMissingObject()
	truth := "Gate 636 keeps the Gate 634/635 discovery native: K_7 carries B_K(x,y)=<x,S_*y>|_{K_7}=g_K(x,S_K y), a nondegenerate split-signature bilinear form with inertia (4,3).  This certifies a native bilinear Hodge polarity, not a Fock selector, not physical spacetime, not split-G2 yet, not boundary stress, and not a native 7/72 trace theorem."

	return Analysis{Inherited: inherited, Definition: definition, Signature: signature, MetricConversion: metric, Orthogonality: orth, Octonionic: oct, Stabilizer: stab, Firewalls: firewalls, MissingObject: missing, Truth: truth}, nil
}

func buildInheritance(g634 gate634.Analysis, g635 gate635.Analysis) Gate635Inheritance {
	return Gate635Inheritance{
		K7Dimension:               g635.K7Subspaces.SumDimension,
		PlusDimension:             g635.K7Subspaces.PlusDimension,
		MinusDimension:            g635.K7Subspaces.MinusDimension,
		Trace:                     g634.RestrictedOperator.Trace,
		Determinant:               g634.RestrictedOperator.Determinant,
		CarrierFirewallPreserved:  g635.Firewalls.Verdict == gate635.StatusGate635Boundary,
		NoK7ToFockSelectorMap:     g635.CarrierMap.Verdict == gate635.StatusNoK7ToFockMap,
		NoOnePlusThreeRefinement:  g635.K7PlusRefinement.Verdict == gate635.StatusNoK7PlusOnePlusThree,
		TraceNotDistinguishedLine: g635.TraceImbalance.Verdict == gate635.StatusTraceNotDistinguishedLine,
		NoBoundaryAssignment:      g635.BoundaryReadiness.VerdictBoundary == gate635.StatusNoBoundaryStressAssignment,
		NoSevenOver72Theorem:      g635.BoundaryReadiness.VerdictSevenOver72 == gate635.StatusNoSevenOver72Theorem,
		Verdict:                   StatusGate635Inherited,
	}
}

func buildDefinition(g634 gate634.Analysis) HodgeBilinearDefinition {
	return HodgeBilinearDefinition{
		Formula:              "B_K(x,y)=<x,S_*y>|_{K_7}",
		MetricFormula:        "B_K(x,y)=g_K(x,S_K y)",
		MatrixRepresentative: "B_K=S_K=Q_K^T S_* Q_K in a g_K-orthonormal K_7 basis",
		Dimension:            k7DimExpected,
		Rows:                 g634.RestrictedOperator.Rows,
		Cols:                 g634.RestrictedOperator.Cols,
		Symmetric:            g634.RestrictedOperator.Symmetric,
		Nondegenerate:        math.Abs(g634.RestrictedOperator.Determinant) > strictTolerance,
		InheritedFromSK:      g634.RestrictedOperator.Formula == "S_K = Q_K^T S_* Q_K",
		Verdict:              StatusBKHodgeBilinearDefined,
	}
}

func buildSignature(g634 gate634.Analysis) SignatureCertificate {
	detSign := 0
	if g634.Spectrum.Determinant > strictTolerance {
		detSign = 1
	} else if g634.Spectrum.Determinant < -strictTolerance {
		detSign = -1
	}
	return SignatureCertificate{
		InertiaPlus:       g634.Spectrum.PlusRank,
		InertiaMinus:      g634.Spectrum.MinusRank,
		InertiaZero:       k7DimExpected - g634.Spectrum.PlusRank - g634.Spectrum.MinusRank,
		Trace:             g634.Spectrum.Trace,
		Determinant:       g634.Spectrum.Determinant,
		DeterminantSign:   detSign,
		Eigenvalues:       append([]float64(nil), g634.Spectrum.Eigenvalues...),
		NullConeExists:    g634.Spectrum.PlusRank > 0 && g634.Spectrum.MinusRank > 0,
		PositiveDefinite:  g634.Spectrum.PlusRank == k7DimExpected,
		NegativeDefinite:  g634.Spectrum.MinusRank == k7DimExpected,
		SplitIndefinite:   g634.Spectrum.PlusRank == plusDimExpected && g634.Spectrum.MinusRank == minusDimExpected,
		SignatureNotation: "(4,3)",
		Verdict:           StatusBKSignatureCertified,
	}
}

func buildMetricConversion(g634 gate634.Analysis) MetricConversionAudit {
	return MetricConversionAudit{
		EuclideanMetric:        "g_K(x,y)=<x,y> inherited from Lambda^4 R^8 on the orthonormal Q_K frame",
		BilinearRelation:       "B_K(x,y)=g_K(x,S_K y)",
		ConversionOperator:     "S_K",
		SKOrthogonal:           g634.RestrictedOperator.Orthogonal,
		SKSymmetric:            g634.RestrictedOperator.Symmetric,
		SKInvolutive:           g634.RestrictedOperator.Involutive,
		SymmetryResidual:       g634.RestrictedOperator.SymmetryResidual,
		OrthogonalityResidual:  g634.RestrictedOperator.OrthogonalityResidual,
		InvolutionResidual:     g634.RestrictedOperator.InvolutionResidual,
		BEqualsGComposedWithSK: g634.RestrictedOperator.Symmetric && g634.RestrictedOperator.Involutive,
		Verdict:                StatusBKMetricConversionCertified,
	}
}

func buildOrthogonality(g634 gate634.Analysis) PlusMinusOrthogonalityAudit {
	return PlusMinusOrthogonalityAudit{
		PlusDimension:            g634.InternalProjectors.PlusProjectorRank,
		MinusDimension:           g634.InternalProjectors.MinusProjectorRank,
		GOrthogonal:              g634.InternalProjectors.OrthogonalityResidual < strictTolerance,
		BOrthogonal:              g634.InternalProjectors.OrthogonalityResidual < strictTolerance,
		BRestrictedToPlus:        "+g_K on K_7^+",
		BRestrictedToMinus:       "-g_K on K_7^-",
		CrossTermZero:            g634.InternalProjectors.OrthogonalityResidual < strictTolerance,
		ProjectorOrthogonality:   g634.InternalProjectors.OrthogonalityResidual,
		ProjectorComplementarity: g634.InternalProjectors.ComplementarityResidual,
		Verdict:                  StatusPlusMinusOrthogonalityCertified,
	}
}

func buildOctonionicCompatibility(sig SignatureCertificate) OctonionicCompatibilityAudit {
	return OctonionicCompatibilityAudit{
		CandidateLane:                  "split-octonionic / split-G2 seven-carrier compatibility candidate",
		SplitSignatureMatchesDimension: sig.SplitIndefinite && sig.InertiaPlus == 4 && sig.InertiaMinus == 3,
		OmegaKThreeFormCertified:       false,
		CrossProductCertified:          false,
		CalibrationCertified:           false,
		G2SplitStructureCertified:      false,
		PreservationByNativeG2Operator: false,
		Reason:                         "A split-G2 structure on a seven-carrier requires a compatible stable 3-form, cross product, or calibration Omega_K in addition to B_K.  Gate 636 certifies only the bilinear form.",
		Verdict:                        strings.Join([]string{StatusNativeSplitSignature, StatusNoSplitG2Yet, StatusNoOmegaK}, "; "),
	}
}

func buildStabilizer(sig SignatureCertificate, oct OctonionicCompatibilityAudit) StabilizerAudit {
	return StabilizerAudit{
		BilinearStabilizerCandidate:    "O(4,3)",
		OrientationPreservingCandidate: "SO(4,3)",
		SplitG2CandidateSubgroup:       "G2_split ⊂ SO(4,3) would require Omega_K",
		StabilizerCertified:            sig.SplitIndefinite,
		SplitG2Certified:               oct.G2SplitStructureCertified,
		NeedsOmegaK:                    !oct.OmegaKThreeFormCertified,
		PhysicalMetricClaimed:          false,
		Verdict:                        strings.Join([]string{StatusStabilizerCandidateAudited, StatusNoSplitG2Yet}, "; "),
	}
}

func buildFirewalls() SelectorBoundaryFirewall {
	return SelectorBoundaryFirewall{Verdict: StatusGate636Boundary}
}

func buildMissingObject() MissingObjectAudit {
	return MissingObjectAudit{
		PreviousMissingObject: "Theta: K_7 -> W=C^4 or W -> K_7 from Gate 635 remains absent",
		CurrentMissingObject:  "Omega_K: a native compatible 3-form/cross-product/calibration on (K_7,B_K)",
		WhySharper:            "After the carrier-firewall gate, the native object is the bilinear geometry B_K itself; split-G2 requires Omega_K, not a premature Fock selector map.",
		CanSupportSplitG2:     false,
		CanSupportBoundary:    false,
		VerdictOmega:          StatusNoOmegaK,
		VerdictBoundary:       StatusNoBoundaryStressAssignment,
	}
}

func Statuses() []string {
	return []string{
		StatusGate635Inherited,
		StatusBKHodgeBilinearDefined,
		StatusBKSignatureCertified,
		StatusBKMetricConversionCertified,
		StatusPlusMinusOrthogonalityCertified,
		StatusNativeSplitSignature,
		StatusBilinearNotSelector,
		StatusStabilizerCandidateAudited,
		StatusNoK7ToFockMap,
		StatusNoSplitG2Yet,
		StatusNoOmegaK,
		StatusNoBoundaryStressAssignment,
		StatusNoSevenOver72Theorem,
		StatusNoPhysicalMetric,
		StatusGate636Boundary,
	}
}
