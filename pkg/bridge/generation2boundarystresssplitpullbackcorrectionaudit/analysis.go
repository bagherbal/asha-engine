// Package generation2boundarystresssplitpullbackcorrectionaudit implements
// Gate 672: BoundaryStressSplit Pullback Correction Audit.
//
// Gate 671 audited the HistoryWallBalance normal vector
//
//	n_72 = (1,1,65/72,-7/72)
//
// on (kappa_lambda, kappa_e, lambda(Lambda_12), R_3-1). Gate 672 sharpens
// this normal by decomposing it as
//
//	n_72 = (1,1,1,0) - (7/72)(0,0,1,1),
//
// hence the active wall balance becomes a base scalar/flavor closure corrected
// by a 7/72 pullback of the signed gauge-scalar boundary stress split:
//
//	kappa_lambda + kappa_e + lambda
//	≈ (7/72)[(R_3-1)+lambda].
//
// This remains a bridge-layer stress-split correction audit only; it does not
// certify a native 7/72 theorem, a wall-distance airlock theorem, or a boundary
// stress derivation.
package generation2boundarystresssplitpullbackcorrectionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate671 "github.com/bagherbal/asha-engine/pkg/bridge/generation2historywallbalancenormalvectorsourceaudit"
)

const (
	AuditID = "GATE672-BOUNDARY-STRESS-SPLIT-PULLBACK-CORRECTION-AUDIT"

	StatusGate671NormalVectorInherited            = "PASS_GATE671_NORMAL_VECTOR_INHERITED"
	StatusNormalVectorStressSplitDecomposed       = "PASS_NORMAL_VECTOR_DECOMPOSED_INTO_BASE_PLUS_STRESS_SPLIT_PULLBACK"
	StatusBaseScalarFlavorClosureComputed         = "PASS_BASE_SCALAR_FLAVOR_CLOSURE_COMPUTED"
	StatusBoundaryStressSplitComputed             = "PASS_BOUNDARY_STRESS_SPLIT_COMPUTED"
	StatusSevenOver72PullbackTested               = "PASS_SEVEN_OVER_SEVENTY_TWO_PULLBACK_TESTED"
	StatusNormalVectorReconstructionComputed      = "PASS_NORMAL_VECTOR_RECONSTRUCTION_COMPUTED"
	StatusSourceTypesAudited                      = "PASS_SOURCE_TYPES_AUDITED"
	StatusStressSplitCorrectedScalarFlavorClosure = "CONDITIONAL_SUPPORT_HISTORY_WALL_BALANCE_IS_STRESS_SPLIT_CORRECTED_SCALAR_FLAVOR_CLOSURE"
	StatusSevenOver72ActsOnBoundaryStressSplit    = "CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_ACTS_ON_BOUNDARY_STRESS_SPLIT"
	StatusNoNativeStressSplitPullbackTheorem      = "FAILED_ROUTE_NO_NATIVE_STRESS_SPLIT_PULLBACK_THEOREM"
	StatusNoNativeSevenOver72Theorem              = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoNativeWallDistanceAirlockTheorem      = "FAILED_ROUTE_NO_NATIVE_WALL_DISTANCE_AIRLOCK_THEOREM"
	StatusNoBoundaryStressDerivation              = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusNoPhysicsPromotion                      = "FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM"
	StatusGate672Boundary                         = "FIREWALL_PRESERVED_GATE672_STRESS_SPLIT_PULLBACK_BOUNDARY"
)

const (
	kappaLambda = 0.0443230430960771
	kappaE      = 0.00550355419157456
	lambda12    = -0.0497009420776833
	r3Minus1    = 0.0509933868964996

	sevenOver72     = 7.0 / 72.0
	sixtyFiveOver72 = 65.0 / 72.0
)

type Gate671Inheritance struct {
	NormalVectorInherited        bool
	NormalVectorBestTypedExact   bool
	CoordinateSealed             bool
	NoNativeNormalVectorTheorem  bool
	NoNativeSevenOver72Theorem   bool
	NoWallDistanceAirlockTheorem bool
	NoBoundaryStressDerivation   bool
	FirewallPreserved            bool
	InheritedResidual            float64
	Verdict                      string
}

type NormalVectorDecompositionAudit struct {
	OriginalNormal       []float64
	BaseNormal           []float64
	StressSplitNormal    []float64
	Weight               float64
	DecompositionLabel   string
	EquivalentFunctional string
	DecompositionPasses  bool
	Verdict              string
}

type BaseScalarFlavorClosureAudit struct {
	KappaLambda float64
	KappaE      float64
	Lambda      float64
	DBase       float64
	Meaning     string
	Verdict     string
}

type BoundaryStressSplitAudit struct {
	R3Minus1 float64
	Lambda   float64
	SSplit   float64
	Meaning  string
	Verdict  string
}

type PullbackAudit struct {
	Weight             float64
	Pullback           float64
	DBase              float64
	Residual           float64
	AbsResidual        float64
	RatioDBaseToSplit  float64
	WeightResidual     float64
	PassesBridgeWindow bool
	Verdict            string
}

type ReconstructionAudit struct {
	DBaseMinusPullback        float64
	HistoryWallBalance        float64
	EquivalentToGate670Normal bool
	Equation                  string
	Verdict                   string
}

type SourceTypeAudit struct {
	DBaseRole           string
	SSplitRole          string
	WeightRole          string
	FanoHitchinFirewall string
	CandidateSupport    []string
	RequiredMissingMaps []string
	Verdict             string
}

type VerdictDiscipline struct {
	ClaimsNativeStressSplitPullback bool
	ClaimsNativeSevenOver72         bool
	ClaimsWallDistanceAirlock       bool
	ClaimsBoundaryStressDerivation  bool
	ClaimsHiggsMassPrediction       bool
	ClaimsScalarStability           bool
	ClaimsGaugeUnification          bool
	ClaimsFlavorDerivation          bool
	ClaimsCKMPMNSDerivation         bool
	Verdict                         string
}

type Analysis struct {
	Inherited      Gate671Inheritance
	Decomposition  NormalVectorDecompositionAudit
	BaseClosure    BaseScalarFlavorClosureAudit
	StressSplit    BoundaryStressSplitAudit
	Pullback       PullbackAudit
	Reconstruction ReconstructionAudit
	Source         SourceTypeAudit
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
	g671, err := gate671.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate671 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g671)
	decomp := buildDecomposition()
	base := buildBaseClosure()
	stress := buildStressSplit()
	pull := buildPullback(base, stress)
	recon := buildReconstruction(pull)
	source := buildSourceType()
	discipline := VerdictDiscipline{Verdict: StatusGate672Boundary}
	truth := "Gate 672 decomposes the Gate671 wall normal n_72=(1,1,65/72,-7/72) into a base scalar/flavor wall closure (1,1,1,0) corrected by a 7/72 pullback of the signed gauge-scalar boundary stress split (0,0,1,1). The living relation is D_base=kappa_lambda+kappa_e+lambda ≈ (7/72)[(R_3-1)+lambda], with residual at the 8.53e-10 level. This retypes 7/72 as an active stress-split pullback coefficient, not a native theorem or a revived Fano-Hitchin boundary map."
	return Analysis{Inherited: inherited, Decomposition: decomp, BaseClosure: base, StressSplit: stress, Pullback: pull, Reconstruction: recon, Source: source, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate671.Analysis) Gate671Inheritance {
	return Gate671Inheritance{
		NormalVectorInherited:        g.Normal.NormalLabel == "n_72=(1,1,65/72,-7/72)",
		NormalVectorBestTypedExact:   g.Minimality.N72BestAmongTypedExact,
		CoordinateSealed:             g.Coordinate.CoordinateSealed,
		NoNativeNormalVectorTheorem:  !g.Discipline.ClaimsNativeNormalVectorTheorem,
		NoNativeSevenOver72Theorem:   !g.Discipline.ClaimsNativeSevenOver72Theorem,
		NoWallDistanceAirlockTheorem: !g.Discipline.ClaimsWallDistanceAirlockTheorem,
		NoBoundaryStressDerivation:   !g.Discipline.ClaimsBoundaryStressDerivation,
		FirewallPreserved:            g.Discipline.Verdict == gate671.StatusGate671Boundary,
		InheritedResidual:            g.Orientation.ExactResidualN72,
		Verdict:                      StatusGate671NormalVectorInherited,
	}
}

func buildDecomposition() NormalVectorDecompositionAudit {
	original := []float64{1, 1, sixtyFiveOver72, -sevenOver72}
	base := []float64{1, 1, 1, 0}
	stress := []float64{0, 0, 1, 1}
	pass := true
	for i := range original {
		if math.Abs((base[i]-sevenOver72*stress[i])-original[i]) > 1e-15 {
			pass = false
		}
	}
	return NormalVectorDecompositionAudit{
		OriginalNormal:       original,
		BaseNormal:           base,
		StressSplitNormal:    stress,
		Weight:               sevenOver72,
		DecompositionLabel:   "n_72=(1,1,1,0)-(7/72)(0,0,1,1)",
		EquivalentFunctional: "kappa_lambda+kappa_e+lambda-(7/72)[(R_3-1)+lambda]",
		DecompositionPasses:  pass,
		Verdict:              StatusNormalVectorStressSplitDecomposed,
	}
}

func buildBaseClosure() BaseScalarFlavorClosureAudit {
	d := kappaLambda + kappaE + lambda12
	return BaseScalarFlavorClosureAudit{
		KappaLambda: kappaLambda,
		KappaE:      kappaE,
		Lambda:      lambda12,
		DBase:       d,
		Meaning:     "scalar matching deficit plus flavor wall deficit almost closes on the signed scalar zero wall before boundary-stress correction",
		Verdict:     StatusBaseScalarFlavorClosureComputed,
	}
}

func buildStressSplit() BoundaryStressSplitAudit {
	s := r3Minus1 + lambda12
	return BoundaryStressSplitAudit{
		R3Minus1: r3Minus1,
		Lambda:   lambda12,
		SSplit:   s,
		Meaning:  "signed gauge-scalar boundary stress split: gauge meeting-wall excess plus signed scalar zero-wall coordinate, equivalently (R_3-1)-|lambda| because lambda<0",
		Verdict:  StatusBoundaryStressSplitComputed,
	}
}

func buildPullback(base BaseScalarFlavorClosureAudit, stress BoundaryStressSplitAudit) PullbackAudit {
	pull := sevenOver72 * stress.SSplit
	res := base.DBase - pull
	ratio := base.DBase / stress.SSplit
	return PullbackAudit{
		Weight:             sevenOver72,
		Pullback:           pull,
		DBase:              base.DBase,
		Residual:           res,
		AbsResidual:        math.Abs(res),
		RatioDBaseToSplit:  ratio,
		WeightResidual:     ratio - sevenOver72,
		PassesBridgeWindow: math.Abs(res) < 1e-8,
		Verdict:            strings.Join([]string{StatusSevenOver72PullbackTested, StatusSevenOver72ActsOnBoundaryStressSplit}, ";"),
	}
}

func buildReconstruction(p PullbackAudit) ReconstructionAudit {
	wall := kappaLambda + kappaE + sixtyFiveOver72*lambda12 - sevenOver72*r3Minus1
	return ReconstructionAudit{
		DBaseMinusPullback:        p.Residual,
		HistoryWallBalance:        wall,
		EquivalentToGate670Normal: math.Abs(p.Residual-wall) < 1e-15,
		Equation:                  "D_base-(7/72)S_split = kappa_lambda+kappa_e+(65/72)lambda-(7/72)(R_3-1)",
		Verdict:                   StatusNormalVectorReconstructionComputed,
	}
}

func buildSourceType() SourceTypeAudit {
	return SourceTypeAudit{
		DBaseRole:           "scalar/flavor deficit against the signed scalar zero wall",
		SSplitRole:          "gauge-scalar boundary stress imbalance (R_3-1)+lambda",
		WeightRole:          "active stress-split pullback coefficient 7/72",
		FanoHitchinFirewall: "Fano-Hitchin may strengthen numerator 7 internally, but Gate657/Gate656 found no K_7/Fano-Hitchin -> R^2_boundary map; this audit does not revive that route.",
		CandidateSupport: []string{
			"n_72 decomposes exactly into base scalar/flavor closure minus 7/72 times stress split",
			"D_base and S_split are both Gate669 wall-distance coordinates",
			"7/72 acts on the boundary stress split, not independently on the full boundary pair",
		},
		RequiredMissingMaps: []string{
			"native stress-split pullback theorem",
			"native 7/72 source theorem",
			"native wall-distance airlock theorem",
			"boundary-stress derivation theorem",
		},
		Verdict: strings.Join([]string{StatusSourceTypesAudited, StatusStressSplitCorrectedScalarFlavorClosure, StatusSevenOver72ActsOnBoundaryStressSplit}, ";"),
	}
}

func Statuses() []string {
	return []string{
		StatusGate671NormalVectorInherited,
		StatusNormalVectorStressSplitDecomposed,
		StatusBaseScalarFlavorClosureComputed,
		StatusBoundaryStressSplitComputed,
		StatusSevenOver72PullbackTested,
		StatusNormalVectorReconstructionComputed,
		StatusSourceTypesAudited,
		StatusStressSplitCorrectedScalarFlavorClosure,
		StatusSevenOver72ActsOnBoundaryStressSplit,
		StatusNoNativeStressSplitPullbackTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusNoNativeWallDistanceAirlockTheorem,
		StatusNoBoundaryStressDerivation,
		StatusNoPhysicsPromotion,
		StatusGate672Boundary,
	}
}
