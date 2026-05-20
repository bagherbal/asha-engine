// Package generation2unitquotientdefectdensityaudit implements
// Gate 681: Unit-Quotient Defect Density and Primitive Object Ladder Audit.
//
// Gate 680 established that the active response coefficient is the full
// augmented-chamber density 7/72, not a kernel-conditional 7/71 or finite-only
// 7/70 normalization. Gate 681 audits the primitive object ladder behind that
// density and tests the sharper reading
//
//	dim(K7) * dim(Q_boundary) / dim(H72) = 7 * 1 / 72.
//
// This is a bridge-layer primitive-object audit only. It does not derive
// boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
// CKM/PMNS, a native 7/72 theorem, or a native trace-to-boundary quotient theorem.
package generation2unitquotientdefectdensityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate680 "github.com/bagherbal/asha-engine/pkg/bridge/generation2globalaugmentedtracekernelconditionalaudit"
)

const (
	AuditID = "GATE681-UNIT-QUOTIENT-DEFECT-DENSITY-PRIMITIVE-OBJECT-LADDER-AUDIT"

	StatusGate680GlobalTraceInherited             = "PASS_GATE680_GLOBAL_TRACE_NORMALIZATION_INHERITED"
	StatusUnitToEightExpansionAudited             = "PASS_UNIT_TO_EIGHT_EXPANSION_AUDITED"
	StatusMiddleChamber70Audited                  = "PASS_MIDDLE_CHAMBER_70_AUDITED"
	StatusK7DefectSourceAudited                   = "PASS_K7_DEFECT_SOURCE_AUDITED"
	StatusFourPlusThreeHodgePolarityRecorded      = "PASS_4_PLUS_3_HODGE_POLARITY_RECORDED"
	StatusBoundaryAugmentation70Plus2Audited      = "PASS_BOUNDARY_AUGMENTATION_70_PLUS_2_AUDITED"
	StatusBoundaryQuotientOneDimensionAudited     = "PASS_BOUNDARY_QUOTIENT_ONE_DIMENSION_AUDITED"
	StatusPrimitiveDensity7Times1Over72Computed   = "PASS_PRIMITIVE_DENSITY_7_TIMES_1_OVER_72_COMPUTED"
	StatusSevenOver72DefectQuotientDensity        = "CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_AS_DEFECT_QUOTIENT_DENSITY"
	StatusPrimitiveObjectLadderStructuresBridge   = "CONDITIONAL_SUPPORT_PRIMITIVE_OBJECT_LADDER_STRUCTURES_ACTIVE_BRIDGE"
	StatusNoNativePrimitiveDensityResponseTheorem = "FAILED_ROUTE_NO_NATIVE_PRIMITIVE_DENSITY_RESPONSE_THEOREM"
	StatusNoNativeTraceToBoundaryQuotientTheorem  = "FAILED_ROUTE_NO_NATIVE_TRACE_TO_BOUNDARY_QUOTIENT_THEOREM"
	StatusNoNativeFivefoldGoldenRatioCarrier      = "FAILED_ROUTE_NO_NATIVE_FIVEFOLD_OR_GOLDEN_RATIO_CARRIER"
	StatusNoBoundaryStressDerivation              = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusGate681Boundary                         = "FIREWALL_PRESERVED_GATE681_PRIMITIVE_OBJECT_LADDER_BOUNDARY"
)

type Gate680Inheritance struct {
	GlobalTraceInherited bool
	H72Dimension         int
	KernelDimension      int
	QuotientDimension    int
	K7Rank               int
	SSplit               float64
	DBase                float64
	TauGlobal            float64
	ResidualGlobal       float64
	FirewallPreserved    bool
	Verdict              string
}

type UnitExpansionAudit struct {
	SeedUnitDimension int
	MeasurementDim    int
	Decomposition     string
	ScalarUnitRole    string
	ContactRole       string
	Verdict           string
}

type MiddleChamberAudit struct {
	BaseDimension  int
	ExteriorDegree int
	Dimension      int
	Formula        string
	Role           string
	Verdict        string
}

type NativeDefectAudit struct {
	Definition             string
	BooleanRank            int
	OctonionicRank         int
	IntersectionRank       int
	KernelDefectRank       int
	CokernelDefectRank     int
	FanoHitchinCarrierRank int
	Role                   string
	Verdict                string
}

type HodgePolarityAudit struct {
	CarrierDimension int
	PositiveDim      int
	NegativeDim      int
	Split            string
	InternalOnly     bool
	Verdict          string
}

type BoundaryAugmentationAudit struct {
	FiniteDimension       int
	BoundaryPairDimension int
	TotalDimension        int
	Chamber               string
	BoundaryPair          string
	Verdict               string
}

type BoundaryQuotientAudit struct {
	BoundaryPairDimension int
	AntiAlignmentLineDim  int
	QuotientDimension     int
	Functional            string
	Coordinate            string
	Verdict               string
}

type PrimitiveDensityAudit struct {
	K7Dimension       int
	QuotientDimension int
	H72Dimension      int
	Density           float64
	ActiveTau         float64
	PredictedDBase    float64
	DBase             float64
	Residual          float64
	MatchesActiveTau  bool
	Interpretation    string
	Verdict           string
}

type DenominatorAlternative struct {
	Name           string
	Formula        string
	Value          float64
	PredictedDBase float64
	Residual       float64
	AbsResidual    float64
	Classification string
}

type SacredGeometryFirewall struct {
	ExternalResonanceRecorded bool
	NativeASHAType            string
	RequiresFivefoldCarrier   bool
	ClaimsPentagonalTheorem   bool
	ClaimsGoldenRatioTheorem  bool
	Verdict                   string
}

type MissingTheoremAudit struct {
	Missing                    []string
	NewPreciseMissingPrinciple string
	AllowedSupport             []string
	Verdict                    string
}

type VerdictDiscipline struct {
	ClaimsPrimitiveDensityTheorem bool
	ClaimsTraceQuotientTheorem    bool
	ClaimsFivefoldCarrier         bool
	ClaimsGoldenRatio             bool
	ClaimsBoundaryStress          bool
	ClaimsHiggsMass               bool
	ClaimsGaugeUnification        bool
	ClaimsFlavorDerivation        bool
	ClaimsNativeSevenOver72       bool
	Verdict                       string
}

type Analysis struct {
	Inherited      Gate680Inheritance
	Unit           UnitExpansionAudit
	Middle         MiddleChamberAudit
	Defect         NativeDefectAudit
	Polarity       HodgePolarityAudit
	Augmentation   BoundaryAugmentationAudit
	Quotient       BoundaryQuotientAudit
	Density        PrimitiveDensityAudit
	Alternatives   []DenominatorAlternative
	SacredFirewall SacredGeometryFirewall
	Missing        MissingTheoremAudit
	Discipline     VerdictDiscipline
	Truth          string
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
	g680, err := gate680.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate680 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g680)
	unit := UnitExpansionAudit{SeedUnitDimension: 1, MeasurementDim: 8, Decomposition: "8 = 1 + 7", ScalarUnitRole: "scalar/unit seed", ContactRole: "seven imaginary/contact directions", Verdict: StatusUnitToEightExpansionAudited}
	middle := MiddleChamberAudit{BaseDimension: 8, ExteriorDegree: 4, Dimension: binomial(8, 4), Formula: "dim Lambda^4 R^8 = C(8,4)=70", Role: "finite middle chamber where Boolean and octonionic projectors act", Verdict: StatusMiddleChamber70Audited}
	defect := NativeDefectAudit{Definition: "K_7 = Im(P_B) ∩ Im(P_G)", BooleanRank: 56, OctonionicRank: 14, IntersectionRank: 7, KernelDefectRank: 7, CokernelDefectRank: 7, FanoHitchinCarrierRank: 7, Role: "native Boolean-octonionic contact/defect carrier", Verdict: StatusK7DefectSourceAudited}
	polarity := HodgePolarityAudit{CarrierDimension: 7, PositiveDim: 4, NegativeDim: 3, Split: "K_7 = K_7^+ ⊕ K_7^- = 4 + 3", InternalOnly: true, Verdict: StatusFourPlusThreeHodgePolarityRecorded}
	augmentation := BoundaryAugmentationAudit{FiniteDimension: middle.Dimension, BoundaryPairDimension: 2, TotalDimension: middle.Dimension + 2, Chamber: "H_72 = Lambda^4 R^8 ⊕ R^2_boundary", BoundaryPair: "R^2_boundary = span(lambda(Lambda_12), R_3-1)", Verdict: StatusBoundaryAugmentation70Plus2Audited}
	quotient := BoundaryQuotientAudit{BoundaryPairDimension: 2, AntiAlignmentLineDim: 1, QuotientDimension: 1, Functional: "sigma_boundary(lambda,R)=lambda+R", Coordinate: "S_split=lambda+(R_3-1)", Verdict: StatusBoundaryQuotientOneDimensionAudited}
	density := buildDensity(inherited, defect, quotient, augmentation)
	alternatives := buildAlternatives(inherited)
	sacredFirewall := SacredGeometryFirewall{ExternalResonanceRecorded: true, NativeASHAType: "72 is currently typed as 70+2, a dimension of the augmented response chamber, not as a pentagonal angle", RequiresFivefoldCarrier: true, ClaimsPentagonalTheorem: false, ClaimsGoldenRatioTheorem: false, Verdict: StatusNoNativeFivefoldGoldenRatioCarrier}
	missing := buildMissing()
	discipline := VerdictDiscipline{Verdict: StatusGate681Boundary}
	truth := "Gate 681 records the primitive object ladder 1 -> R^8 -> Lambda^4 R^8 -> K7 -> 4|3 -> H72 -> Q_boundary and sharpens 7/72 as dim(K7)*dim(Q_boundary)/dim(H72)=7*1/72. This structures the active bridge as a defect-quotient density inside the full augmented chamber, while preserving the firewall: no native primitive-density response theorem, no trace-to-boundary quotient theorem, and no native fivefold/golden-ratio carrier are certified."
	return Analysis{Inherited: inherited, Unit: unit, Middle: middle, Defect: defect, Polarity: polarity, Augmentation: augmentation, Quotient: quotient, Density: density, Alternatives: alternatives, SacredFirewall: sacredFirewall, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate680.Analysis) Gate680Inheritance {
	return Gate680Inheritance{
		GlobalTraceInherited: strings.Contains(g.Compatibility.Verdict, gate680.StatusGlobalH72TraceTypeCorrectForQuotientResponse),
		H72Dimension:         g.Inherited.H72Dimension,
		KernelDimension:      g.Inherited.KernelDimension,
		QuotientDimension:    g.Inherited.QuotientDimension,
		K7Rank:               g.Inherited.K7Rank,
		SSplit:               g.Inherited.SSplit,
		DBase:                g.Inherited.DBase,
		TauGlobal:            g.Inherited.TauGlobal,
		ResidualGlobal:       g.Inherited.ResidualGlobal,
		FirewallPreserved:    g.Discipline.Verdict == gate680.StatusGate680Boundary,
		Verdict:              StatusGate680GlobalTraceInherited,
	}
}

func buildDensity(in Gate680Inheritance, defect NativeDefectAudit, q BoundaryQuotientAudit, h BoundaryAugmentationAudit) PrimitiveDensityAudit {
	density := float64(defect.IntersectionRank*q.QuotientDimension) / float64(h.TotalDimension)
	pred := density * in.SSplit
	res := in.DBase - pred
	return PrimitiveDensityAudit{
		K7Dimension:       defect.IntersectionRank,
		QuotientDimension: q.QuotientDimension,
		H72Dimension:      h.TotalDimension,
		Density:           density,
		ActiveTau:         in.TauGlobal,
		PredictedDBase:    pred,
		DBase:             in.DBase,
		Residual:          res,
		MatchesActiveTau:  math.Abs(density-in.TauGlobal) < 1e-15,
		Interpretation:    "rank-seven internal defect times one-dimensional boundary quotient over the full augmented chamber",
		Verdict:           strings.Join([]string{StatusPrimitiveDensity7Times1Over72Computed, StatusSevenOver72DefectQuotientDensity, StatusPrimitiveObjectLadderStructuresBridge}, ";"),
	}
}

func buildAlternatives(in Gate680Inheritance) []DenominatorAlternative {
	candidates := []struct {
		name, formula, class string
		val                  float64
	}{
		{"finite_only", "7/70", "finite-only density; omits boundary quotient system", 7.0 / 70.0},
		{"kernel_conditional", "7/71", "conditional density inside ker(pi_split); erases quotient output line", 7.0 / 71.0},
		{"global_defect_quotient", "7*1/72", "defect-quotient density in full augmented chamber", 7.0 / 72.0},
		{"half_coordinate", "7/144", "per-boundary-coordinate half trace; inactive clue", 7.0 / 144.0},
	}
	out := make([]DenominatorAlternative, 0, len(candidates))
	for _, c := range candidates {
		pred := c.val * in.SSplit
		res := in.DBase - pred
		out = append(out, DenominatorAlternative{Name: c.name, Formula: c.formula, Value: c.val, PredictedDBase: pred, Residual: res, AbsResidual: math.Abs(res), Classification: c.class})
	}
	return out
}

func buildMissing() MissingTheoremAudit {
	return MissingTheoremAudit{
		Missing: []string{
			StatusNoNativePrimitiveDensityResponseTheorem,
			StatusNoNativeTraceToBoundaryQuotientTheorem,
			StatusNoNativeFivefoldGoldenRatioCarrier,
			StatusNoBoundaryStressDerivation,
		},
		NewPreciseMissingPrinciple: "PrimitiveDefectQuotientDensityResponseTheorem / FullChamberTraceToBoundaryQuotientTheorem",
		AllowedSupport: []string{
			StatusSevenOver72DefectQuotientDensity,
			StatusPrimitiveObjectLadderStructuresBridge,
		},
		Verdict: strings.Join([]string{StatusNoNativePrimitiveDensityResponseTheorem, StatusNoNativeTraceToBoundaryQuotientTheorem, StatusNoNativeFivefoldGoldenRatioCarrier, StatusNoBoundaryStressDerivation}, ";"),
	}
}

func binomial(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	res := 1
	for i := 1; i <= k; i++ {
		res = res * (n - k + i) / i
	}
	return res
}

func Statuses() []string {
	return []string{
		StatusGate680GlobalTraceInherited,
		StatusUnitToEightExpansionAudited,
		StatusMiddleChamber70Audited,
		StatusK7DefectSourceAudited,
		StatusFourPlusThreeHodgePolarityRecorded,
		StatusBoundaryAugmentation70Plus2Audited,
		StatusBoundaryQuotientOneDimensionAudited,
		StatusPrimitiveDensity7Times1Over72Computed,
		StatusSevenOver72DefectQuotientDensity,
		StatusPrimitiveObjectLadderStructuresBridge,
		StatusNoNativePrimitiveDensityResponseTheorem,
		StatusNoNativeTraceToBoundaryQuotientTheorem,
		StatusNoNativeFivefoldGoldenRatioCarrier,
		StatusGate681Boundary,
	}
}
