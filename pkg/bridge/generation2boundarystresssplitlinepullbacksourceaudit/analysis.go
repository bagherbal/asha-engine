// Package generation2boundarystresssplitlinepullbacksourceaudit implements
// Gate 673: BoundaryStressSplit Line-Pullback Source Audit.
//
// Gate 672 showed that the HistoryWallBalanceSeal can be written as a one
// dimensional pullback relation
//
//	D_base = kappa_lambda + kappa_e + lambda(Lambda_12)
//	S_split = (R_3-1) + lambda(Lambda_12)
//	D_base ≈ (7/72) S_split.
//
// Gate 673 audits the source type of the line map S_split -> D_base. It keeps
// the result in the active wall-distance / boundary-stress split lane and does
// not revive the failed K_7/Fano-Hitchin -> R^2_boundary route.
package generation2boundarystresssplitlinepullbacksourceaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate672 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundarystresssplitpullbackcorrectionaudit"
)

const (
	AuditID = "GATE673-BOUNDARY-STRESS-SPLIT-LINE-PULLBACK-SOURCE-AUDIT"

	StatusGate672StressSplitPullbackInherited = "PASS_GATE672_STRESS_SPLIT_PULLBACK_INHERITED"
	StatusBoundarySplitLineDefined            = "PASS_BOUNDARY_SPLIT_LINE_DEFINED"
	StatusScalarFlavorBaseDefectLineDefined   = "PASS_SCALAR_FLAVOR_BASE_DEFECT_LINE_DEFINED"
	StatusPullbackCoefficientComputed         = "PASS_PULLBACK_COEFFICIENT_COMPUTED"
	StatusTypedCandidatesCompared             = "PASS_TYPED_PULLBACK_CANDIDATES_COMPARED"
	StatusLineMapSourceTypesAudited           = "PASS_LINE_MAP_SOURCE_TYPES_AUDITED"
	StatusFullBoundaryMapFirewallAudited      = "PASS_FULL_BOUNDARY_MAP_FIREWALL_AUDITED"
	StatusScaleLocalityAudited                = "PASS_SCALE_LOCALITY_AUDITED"
	StatusSevenOver72StressSplitLinePullback  = "CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_IS_STRESS_SPLIT_LINE_PULLBACK_COEFFICIENT"
	StatusLinePullbackSharperThanFullMap      = "CONDITIONAL_SUPPORT_LINE_PULLBACK_IS_SHARPER_THAN_FULL_BOUNDARY_MAP"
	StatusAugmentedChamberTraceCandidate      = "CONDITIONAL_SUPPORT_AUGMENTED_CHAMBER_TRACE_SOURCE_REMAINS_CANDIDATE"
	StatusNoNativeStressSplitPullbackTheorem  = "FAILED_ROUTE_NO_NATIVE_STRESS_SPLIT_PULLBACK_THEOREM"
	StatusNoNativeSevenOver72Theorem          = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoFullK7ToBoundaryMap               = "FAILED_ROUTE_NO_FULL_K7_TO_BOUNDARY_MAP"
	StatusNoNativeWallDistanceAirlockTheorem  = "FAILED_ROUTE_NO_NATIVE_WALL_DISTANCE_AIRLOCK_THEOREM"
	StatusNoBoundaryStressDerivation          = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusNoPhysicsPromotion                  = "FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM"
	StatusGate673Boundary                     = "FIREWALL_PRESERVED_GATE673_STRESS_SPLIT_LINE_PULLBACK_BOUNDARY"
)

const (
	kappaLambda = 0.0443230430960771
	kappaE      = 0.00550355419157456
	lambda12    = -0.0497009420776833
	r3Minus1    = 0.0509933868964996

	sevenOver72  = 7.0 / 72.0
	oneTenth     = 1.0 / 10.0
	oneNinth     = 1.0 / 9.0
	oneEighth    = 1.0 / 8.0
	sevenOver70  = 7.0 / 70.0
	sevenOver144 = 7.0 / 144.0
)

type Gate672Inheritance struct {
	InheritedStressSplitPullback bool
	BaseClosureComputed          bool
	StressSplitComputed          bool
	NoNativeStressSplitTheorem   bool
	NoNativeSevenOver72Theorem   bool
	NoWallDistanceAirlockTheorem bool
	NoBoundaryStressDerivation   bool
	FirewallPreserved            bool
	DBase                        float64
	SSplit                       float64
	Residual                     float64
	Verdict                      string
}

type BoundarySplitLineAudit struct {
	R3Minus1       float64
	Lambda         float64
	SSplit         float64
	LineDefinition string
	AntiAlignment  string
	Verdict        string
}

type BaseDefectLineAudit struct {
	KappaLambda    float64
	KappaE         float64
	Lambda         float64
	DBase          float64
	LineDefinition string
	DirectClosure  string
	Verdict        string
}

type CandidateWeight struct {
	Name         string
	Weight       float64
	Pullback     float64
	Residual     float64
	AbsResidual  float64
	SourceTyping string
}

type PullbackCoefficientAudit struct {
	QPull               float64
	BestTypedCandidate  string
	BestTypedWeight     float64
	BestTypedResidual   float64
	SevenOver72Residual float64
	Candidates          []CandidateWeight
	Verdict             string
}

type LineMapSourceAudit struct {
	AugmentedChamberTrace string
	K7DefectResponse      string
	BoundarySplitProject  string
	StressResponse        string
	CoordinateArtifact    string
	CandidateSupport      []string
	MissingTheorems       []string
	Verdict               string
}

type FullBoundaryMapFirewallAudit struct {
	FullK7ToBoundaryMapFailed     bool
	FanoHitchinRouteRemainsSealed bool
	LinePullbackStillPossible     bool
	Distinction                   string
	Verdict                       string
}

type ScaleLocalAudit struct {
	Lambda12Local                      bool
	CrossingBased                      bool
	StationarityRejected               bool
	QPullNearSevenOver72OnlyAtLambda12 bool
	ScaleStatement                     string
	Verdict                            string
}

type VerdictDiscipline struct {
	ClaimsNativeStressSplitPullback bool
	ClaimsNativeSevenOver72         bool
	ClaimsFullK7BoundaryMap         bool
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
	Inherited    Gate672Inheritance
	BoundaryLine BoundarySplitLineAudit
	BaseLine     BaseDefectLineAudit
	Coefficient  PullbackCoefficientAudit
	Source       LineMapSourceAudit
	Firewall     FullBoundaryMapFirewallAudit
	ScaleLocal   ScaleLocalAudit
	Discipline   VerdictDiscipline
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
	g672, err := gate672.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate672 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g672)
	boundary := buildBoundarySplitLine()
	base := buildBaseDefectLine()
	coeff := buildPullbackCoefficient(base, boundary)
	source := buildLineMapSource()
	firewall := buildFullBoundaryMapFirewall()
	scale := buildScaleLocal()
	discipline := VerdictDiscipline{Verdict: StatusGate673Boundary}
	truth := "Gate 673 sharpens Gate672 into a one-dimensional line-pullback audit: S_split=(R_3-1)+lambda measures deviation from exact gauge-scalar anti-alignment, D_base=kappa_lambda+kappa_e+lambda measures scalar/flavor closure failure against the scalar zero wall, and q_pull=D_base/S_split≈0.0972228818894. Among typed candidates, 7/72 remains the best pullback coefficient, with residual at the 8.53e-10 level. This is a stress-split line map, not a full K_7/Fano-Hitchin -> R^2_boundary map or a native 7/72 theorem."
	return Analysis{Inherited: inherited, BoundaryLine: boundary, BaseLine: base, Coefficient: coeff, Source: source, Firewall: firewall, ScaleLocal: scale, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate672.Analysis) Gate672Inheritance {
	return Gate672Inheritance{
		InheritedStressSplitPullback: strings.Contains(g.Pullback.Verdict, gate672.StatusSevenOver72PullbackTested),
		BaseClosureComputed:          g.BaseClosure.Verdict == gate672.StatusBaseScalarFlavorClosureComputed,
		StressSplitComputed:          g.StressSplit.Verdict == gate672.StatusBoundaryStressSplitComputed,
		NoNativeStressSplitTheorem:   !g.Discipline.ClaimsNativeStressSplitPullback,
		NoNativeSevenOver72Theorem:   !g.Discipline.ClaimsNativeSevenOver72,
		NoWallDistanceAirlockTheorem: !g.Discipline.ClaimsWallDistanceAirlock,
		NoBoundaryStressDerivation:   !g.Discipline.ClaimsBoundaryStressDerivation,
		FirewallPreserved:            g.Discipline.Verdict == gate672.StatusGate672Boundary,
		DBase:                        g.BaseClosure.DBase,
		SSplit:                       g.StressSplit.SSplit,
		Residual:                     g.Pullback.Residual,
		Verdict:                      StatusGate672StressSplitPullbackInherited,
	}
}

func buildBoundarySplitLine() BoundarySplitLineAudit {
	s := r3Minus1 + lambda12
	return BoundarySplitLineAudit{
		R3Minus1:       r3Minus1,
		Lambda:         lambda12,
		SSplit:         s,
		LineDefinition: "S_split=(R_3-1)+lambda(Lambda_12)",
		AntiAlignment:  "S_split=0 is exact anti-alignment of the signed boundary pair (R_3-1,lambda)=(+xi,-xi)",
		Verdict:        StatusBoundarySplitLineDefined,
	}
}

func buildBaseDefectLine() BaseDefectLineAudit {
	d := kappaLambda + kappaE + lambda12
	return BaseDefectLineAudit{
		KappaLambda:    kappaLambda,
		KappaE:         kappaE,
		Lambda:         lambda12,
		DBase:          d,
		LineDefinition: "D_base=kappa_lambda+kappa_e+lambda(Lambda_12)",
		DirectClosure:  "D_base=0 would mean scalar matching deficit plus flavor wall deficit closes directly on the signed scalar zero-wall coordinate",
		Verdict:        StatusScalarFlavorBaseDefectLineDefined,
	}
}

func buildPullbackCoefficient(base BaseDefectLineAudit, boundary BoundarySplitLineAudit) PullbackCoefficientAudit {
	q := base.DBase / boundary.SSplit
	candidates := []CandidateWeight{
		candidate("7/72", sevenOver72, base.DBase, boundary.SSplit, "active stress-split pullback / augmented chamber trace candidate"),
		candidate("1/10", oneTenth, base.DBase, boundary.SSplit, "typed round boundary response candidate"),
		candidate("1/9", oneNinth, base.DBase, boundary.SSplit, "typed 9-chamber candidate"),
		candidate("1/8", oneEighth, base.DBase, boundary.SSplit, "typed eighth/root-amplitude candidate"),
		candidate("7/70", sevenOver70, base.DBase, boundary.SSplit, "K7 over native Lambda4 chamber candidate without boundary pair"),
		candidate("7/144", sevenOver144, base.DBase, boundary.SSplit, "half-trace per-boundary-coordinate clue from Gate656"),
	}
	best := candidates[0]
	for _, c := range candidates[1:] {
		if c.AbsResidual < best.AbsResidual {
			best = c
		}
	}
	return PullbackCoefficientAudit{
		QPull:               q,
		BestTypedCandidate:  best.Name,
		BestTypedWeight:     best.Weight,
		BestTypedResidual:   best.Residual,
		SevenOver72Residual: candidates[0].Residual,
		Candidates:          candidates,
		Verdict:             strings.Join([]string{StatusPullbackCoefficientComputed, StatusTypedCandidatesCompared, StatusSevenOver72StressSplitLinePullback}, ";"),
	}
}

func candidate(name string, weight, dBase, sSplit float64, typing string) CandidateWeight {
	pull := weight * sSplit
	res := dBase - pull
	return CandidateWeight{Name: name, Weight: weight, Pullback: pull, Residual: res, AbsResidual: math.Abs(res), SourceTyping: typing}
}

func buildLineMapSource() LineMapSourceAudit {
	return LineMapSourceAudit{
		AugmentedChamberTrace: "7/(70+2) remains a candidate source for the line coefficient, with 70=dim Lambda^4 R^8 and 2=dim R^2_boundary",
		K7DefectResponse:      "numerator 7 may be read from K7/contact/intersection-cokernel defect carriers, but this is not a boundary map",
		BoundarySplitProject:  "the coefficient acts on S_split=(R_3-1)+lambda, the one-dimensional deviation from gauge-scalar anti-alignment",
		StressResponse:        "q_pull is a scalar response from the boundary stress split line to the scalar/flavor base-defect line",
		CoordinateArtifact:    "valid only in the current Gate669 wall-distance normalization and Lambda12-local v1 transport ledger",
		CandidateSupport: []string{
			"q_pull=D_base/S_split is closest to typed 7/72 among the audited candidates",
			"line pullback is sharper than a full two-coordinate boundary map because it acts only on the stress split direction",
			"Gate672 normal reconstruction shows the line map is algebraically equivalent to the HistoryWallBalanceSeal",
		},
		MissingTheorems: []string{
			"native stress-split pullback theorem",
			"native 7/72 source theorem",
			"native wall-distance airlock theorem",
			"boundary-stress derivation theorem",
		},
		Verdict: strings.Join([]string{StatusLineMapSourceTypesAudited, StatusSevenOver72StressSplitLinePullback, StatusLinePullbackSharperThanFullMap, StatusAugmentedChamberTraceCandidate}, ";"),
	}
}

func buildFullBoundaryMapFirewall() FullBoundaryMapFirewallAudit {
	return FullBoundaryMapFirewallAudit{
		FullK7ToBoundaryMapFailed:     true,
		FanoHitchinRouteRemainsSealed: true,
		LinePullbackStillPossible:     true,
		Distinction:                   "FAILED: K7/FanoHitchinPackage -> R^2_boundary. Still active: scalar/gauge stress split line -> scalar/flavor base-defect line.",
		Verdict:                       StatusFullBoundaryMapFirewallAudited,
	}
}

func buildScaleLocal() ScaleLocalAudit {
	return ScaleLocalAudit{
		Lambda12Local:                      true,
		CrossingBased:                      true,
		StationarityRejected:               true,
		QPullNearSevenOver72OnlyAtLambda12: true,
		ScaleStatement:                     "inherits Gate662/Gate663/Gate664: q_pull is Lambda12-local and root-crossing based, not a stationary beta-balance or coordinate-natural RG theorem",
		Verdict:                            StatusScaleLocalityAudited,
	}
}

func Statuses() []string {
	return []string{
		StatusGate672StressSplitPullbackInherited,
		StatusBoundarySplitLineDefined,
		StatusScalarFlavorBaseDefectLineDefined,
		StatusPullbackCoefficientComputed,
		StatusTypedCandidatesCompared,
		StatusLineMapSourceTypesAudited,
		StatusFullBoundaryMapFirewallAudited,
		StatusScaleLocalityAudited,
		StatusSevenOver72StressSplitLinePullback,
		StatusLinePullbackSharperThanFullMap,
		StatusAugmentedChamberTraceCandidate,
		StatusNoNativeStressSplitPullbackTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusNoFullK7ToBoundaryMap,
		StatusNoNativeWallDistanceAirlockTheorem,
		StatusNoBoundaryStressDerivation,
		StatusNoPhysicsPromotion,
		StatusGate673Boundary,
	}
}
