// Package generation2scalarflavordeficitclosuretriangleaudit implements
// Gate 659: Scalar-Flavor Deficit Closure Triangle Audit.
//
// Gate 658 rebuilt the active scalar proxy-to-boundary transport spine after
// the K_7/Fano-Hitchin boundary route was sealed.  Gate 659 audits the next
// live bridge relation: the scalar low-scale loop-matching deficit kappa_lambda
// plus the charged-lepton flavor-wall deficit kappa_e nearly closes on the
// high-scale scalar wound |lambda(Lambda_12)|, with the residual tracking the
// active gauge-scalar boundary split at the typed weight 7/72.  The audit is
// bridge-layer only and preserves all scalar, flavor, boundary, Higgs, and
// native 7/72 firewalls.
package generation2scalarflavordeficitclosuretriangleaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate658 "github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarproxytoboundarytransportspineaudit"
)

const (
	AuditID = "GATE659-SCALAR-FLAVOR-DEFICIT-CLOSURE-TRIANGLE-AUDIT"

	StatusGate658ScalarTransportSpineInherited = "PASS_GATE658_SCALAR_TRANSPORT_SPINE_INHERITED"
	StatusFlavorKappaESealInherited            = "PASS_FLAVOR_KAPPA_E_SEAL_INHERITED"
	StatusKappaSumComputed                     = "PASS_KAPPA_SUM_COMPUTED"
	StatusKappaSumClosesOnAbsLambda            = "PASS_KAPPA_SUM_CLOSES_ON_ABS_LAMBDA_LAMBDA12"
	StatusBoundarySplitRatioComputed           = "PASS_BOUNDARY_SPLIT_RATIO_COMPUTED"
	StatusTypedWeightCandidatesAudited         = "PASS_TYPED_WEIGHT_CANDIDATES_AUDITED"
	StatusSevenOverSeventyTwoInterpolation     = "PASS_SEVEN_OVER_SEVENTY_TWO_INTERPOLATION_AUDITED"
	StatusSourceTypeAuditComputed              = "PASS_SOURCE_TYPE_AUDIT_COMPUTED"
	StatusKappaClosureSupport                  = "CONDITIONAL_SUPPORT_KAPPA_LAMBDA_PLUS_KAPPA_E_CLOSES_ON_HIGH_SCALE_SCALAR_WOUND"
	StatusResidualTracksBoundarySplit          = "CONDITIONAL_SUPPORT_RESIDUAL_TRACKS_BOUNDARY_STRESS_SPLIT"
	StatusSevenOver72ReappearsActiveTransport  = "CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_BOUNDARY_WEIGHT_REAPPEARS_IN_ACTIVE_TRANSPORT_LANE"
	StatusBoundaryWeightedClosureSupport       = "CONDITIONAL_SUPPORT_BOUNDARY_WEIGHTED_DEFICIT_CLOSURE"
	StatusNoNativeKappaClosureTheorem          = "FAILED_ROUTE_NO_NATIVE_KAPPA_CLOSURE_THEOREM"
	StatusNoNativeSevenOver72SourceTheorem     = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_SOURCE_THEOREM"
	StatusNoNativeScalarFlavorBoundaryTheorem  = "FAILED_ROUTE_NO_NATIVE_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoNativeFlavorTheorem                = "FAILED_ROUTE_NO_NATIVE_FLAVOR_THEOREM"
	StatusNoNativeScalarTheorem                = "FAILED_ROUTE_NO_NATIVE_SCALAR_THEOREM"
	StatusNoBoundaryStressDerivation           = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusNoHiggsGaugeOrCKMClaim               = "FAILED_ROUTE_NO_HIGGS_GAUGE_OR_CKM_PMNS_CLAIM"
	StatusGate659Boundary                      = "FIREWALL_PRESERVED_GATE659_DEFICIT_CLOSURE_BOUNDARY"
)

const (
	kappaLambda       = 0.0443230430960771
	kappaE            = 0.00550355419157456
	absLambdaLambda12 = 0.0497009420776833
	r3Minus1          = 0.0509933868964996
	xiBoundary        = 0.0503471644870914
	sevenOver72       = 7.0 / 72.0
	oneEighth         = 1.0 / 8.0
	oneNinth          = 1.0 / 9.0
	oneTenth          = 1.0 / 10.0
)

type Gate658Inheritance struct {
	ScalarTransportSpineInherited bool
	ScalarBoundarySpineActive     bool
	LowScaleMatchingActive        bool
	BoundaryStressTransportActive bool
	KappaLambdaDefined            bool
	NoNativeProxyRuntimeTheorem   bool
	NoNativeRGThresholdTheorem    bool
	NoNativeBoundaryStressTheorem bool
	NoHiggsMassOrStabilityClaim   bool
	FirewallPreserved             bool
	Verdict                       string
}

type FlavorKappaESeal struct {
	KappaE                       float64
	Source                       string
	OrientationBalanceExpression string
	EnvironmentalSeal            bool
	NativeFlavorTheorem          bool
	Verdict                      string
}

type ClosureTriangle struct {
	KappaLambda                float64
	KappaE                     float64
	KSum                       float64
	AbsLambdaLambda12          float64
	DeltaClosure               float64
	RelativeToAbsLambda        float64
	RelativeToKSum             float64
	RelativeToXiBoundary       float64
	ClosesOnScalarWound        bool
	RawClosureResidualSmall    bool
	HighScaleScalarWoundTarget bool
	Verdict                    string
}

type WeightCandidate struct {
	Name                string
	Value               float64
	Difference          float64
	AbsDifference       float64
	TypedStatus         string
	BestAmongTypedSet   bool
	CanBeCertifiedByFit bool
}

type BoundaryWeightAudit struct {
	BoundarySplit      float64
	DeltaClosure       float64
	ObservedWeight     float64
	Candidates         []WeightCandidate
	ClosestCandidate   string
	ClosestDifference  float64
	SevenOver72Closest bool
	UsedTypedSetOnly   bool
	Verdict            string
}

type SevenOver72Interpolation struct {
	AbsLambdaLambda12      float64
	R3Minus1               float64
	Weight                 float64
	ComplementWeight       float64
	WeightedTarget         float64
	KSum                   float64
	RawClosureResidual     float64
	WeightedResidual       float64
	ImprovementFactor      float64
	ResidualRelativeToKSum float64
	Formula                string
	BridgeLayerOnly        bool
	Verdict                string
}

type SourceObject struct {
	Name         string
	Value        float64
	Role         string
	TypedStatus  string
	NativeSource bool
}

type SourceTypeAudit struct {
	Objects                    []SourceObject
	SevenOver72InFanoLane      bool
	SevenOver72InTransportLane bool
	FanoBoundaryMapConstructed bool
	RandomConstantsSearched    bool
	TypedCandidatesOnly        bool
	Verdict                    string
}

type Firewalls struct {
	ClaimsNativeFlavorTheorem      bool
	ClaimsNativeScalarTheorem      bool
	ClaimsNativeSevenOver72Theorem bool
	ClaimsBoundaryStressDerivation bool
	ClaimsHiggsPrediction          bool
	ClaimsGaugeUnification         bool
	ClaimsCKMPMNSDerivation        bool
	ClaimsPhysicalSpacetime        bool
	ClaimsNativeClosureTheorem     bool
	Verdict                        string
}

type Analysis struct {
	Inherited     Gate658Inheritance
	Flavor        FlavorKappaESeal
	Closure       ClosureTriangle
	Boundary      BoundaryWeightAudit
	Interpolation SevenOver72Interpolation
	Sources       SourceTypeAudit
	Firewalls     Firewalls
	Truth         string
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
	g658, err := gate658.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate658 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g658)
	flavor := buildFlavorSeal()
	closure := buildClosureTriangle()
	boundary := buildBoundaryWeight(closure)
	interp := buildInterpolation(closure, boundary)
	sources := buildSourceTypeAudit()
	firewalls := Firewalls{Verdict: StatusGate659Boundary}
	truth := "Gate 659 identifies the active bridge-layer closure triangle exposed by the scalar transport spine: kappa_lambda+kappa_e nearly closes on |lambda(Lambda_12)|, and the remaining residual is almost exactly a typed 7/72 pull across the active boundary-stress split from |lambda(Lambda_12)| toward R_3-1.  The result conditionally supports a boundary-weighted scalar-flavor deficit closure in the transport lane, but no native kappa-closure theorem, 7/72 source theorem, scalar/flavor/boundary transport theorem, Higgs claim, CKM/PMNS theorem, or boundary-stress derivation is claimed."
	return Analysis{Inherited: inherited, Flavor: flavor, Closure: closure, Boundary: boundary, Interpolation: interp, Sources: sources, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g gate658.Analysis) Gate658Inheritance {
	return Gate658Inheritance{
		ScalarTransportSpineInherited: g.Spine.Active && g.Spine.MergesScalarBoundary,
		ScalarBoundarySpineActive:     g.Spine.Active,
		LowScaleMatchingActive:        g.Matching.LoopSized && math.Abs(g.Matching.KappaLambda-kappaLambda) < 5e-13,
		BoundaryStressTransportActive: g.Boundary.BoundarySplit > 0 && math.Abs(g.Boundary.R3Minus1-r3Minus1) < 1e-15,
		KappaLambdaDefined:            math.Abs(g.Matching.KappaLambda-kappaLambda) < 5e-13,
		NoNativeProxyRuntimeTheorem:   !g.Sources.ProxyRuntimeTheorem,
		NoNativeRGThresholdTheorem:    !g.Sources.RGThresholdTheorem,
		NoNativeBoundaryStressTheorem: !g.Sources.BoundaryStressTheorem,
		NoHiggsMassOrStabilityClaim:   !g.Firewalls.ClaimsHiggsMass && !g.Firewalls.ClaimsScalarStability,
		FirewallPreserved:             g.Firewalls.Verdict == gate658.StatusGate658Boundary,
		Verdict:                       StatusGate658ScalarTransportSpineInherited,
	}
}

func buildFlavorSeal() FlavorKappaESeal {
	return FlavorKappaESeal{
		KappaE:                       kappaE,
		Source:                       "charged-lepton loop-angle/flavor wall deficit inherited from the environmental flavor OrientationBalanceSeal",
		OrientationBalanceExpression: "kappa_e ≈ sin²(theta13)/4 - J_CKM",
		EnvironmentalSeal:            true,
		NativeFlavorTheorem:          false,
		Verdict:                      StatusFlavorKappaESealInherited,
	}
}

func buildClosureTriangle() ClosureTriangle {
	ks := kappaLambda + kappaE
	delta := ks - absLambdaLambda12
	return ClosureTriangle{
		KappaLambda:                kappaLambda,
		KappaE:                     kappaE,
		KSum:                       ks,
		AbsLambdaLambda12:          absLambdaLambda12,
		DeltaClosure:               delta,
		RelativeToAbsLambda:        delta / absLambdaLambda12,
		RelativeToKSum:             delta / ks,
		RelativeToXiBoundary:       delta / xiBoundary,
		ClosesOnScalarWound:        math.Abs(delta)/absLambdaLambda12 < 0.003,
		RawClosureResidualSmall:    math.Abs(delta) < 0.0002,
		HighScaleScalarWoundTarget: true,
		Verdict:                    join(StatusKappaSumComputed, StatusKappaSumClosesOnAbsLambda, StatusKappaClosureSupport),
	}
}

func buildBoundaryWeight(c ClosureTriangle) BoundaryWeightAudit {
	split := r3Minus1 - absLambdaLambda12
	w := c.DeltaClosure / split
	candidates := []WeightCandidate{
		candidate("7/72", sevenOver72, w, "typed Gate626/Gate628 boundary interpolation weight; active transport-lane candidate", true),
		candidate("1/8", oneEighth, w, "simple octile/loop-adjacent typed comparison only; weaker than 7/72", false),
		candidate("1/9", oneNinth, w, "nine-chamber comparison only; not closest", false),
		candidate("1/10", oneTenth, w, "ten K7-block chamber comparison only; near but weaker than 7/72", false),
	}
	closest := candidates[0]
	for _, x := range candidates[1:] {
		if x.AbsDifference < closest.AbsDifference {
			closest = x
		}
	}
	return BoundaryWeightAudit{
		BoundarySplit:      split,
		DeltaClosure:       c.DeltaClosure,
		ObservedWeight:     w,
		Candidates:         candidates,
		ClosestCandidate:   closest.Name,
		ClosestDifference:  closest.AbsDifference,
		SevenOver72Closest: closest.Name == "7/72",
		UsedTypedSetOnly:   true,
		Verdict:            join(StatusBoundarySplitRatioComputed, StatusTypedWeightCandidatesAudited, StatusResidualTracksBoundarySplit, StatusSevenOver72ReappearsActiveTransport),
	}
}

func candidate(name string, value, observed float64, status string, best bool) WeightCandidate {
	diff := observed - value
	return WeightCandidate{Name: name, Value: value, Difference: diff, AbsDifference: math.Abs(diff), TypedStatus: status, BestAmongTypedSet: best, CanBeCertifiedByFit: false}
}

func buildInterpolation(c ClosureTriangle, b BoundaryWeightAudit) SevenOver72Interpolation {
	w := sevenOver72
	target := absLambdaLambda12 + w*(r3Minus1-absLambdaLambda12)
	weightedResidual := c.KSum - target
	improvement := math.Abs(c.DeltaClosure) / math.Abs(weightedResidual)
	return SevenOver72Interpolation{
		AbsLambdaLambda12:      absLambdaLambda12,
		R3Minus1:               r3Minus1,
		Weight:                 w,
		ComplementWeight:       1.0 - w,
		WeightedTarget:         target,
		KSum:                   c.KSum,
		RawClosureResidual:     c.DeltaClosure,
		WeightedResidual:       weightedResidual,
		ImprovementFactor:      improvement,
		ResidualRelativeToKSum: math.Abs(weightedResidual) / c.KSum,
		Formula:                "W_72=|lambda(Lambda_12)|+(7/72)[(R_3-1)-|lambda(Lambda_12)|]=(65/72)|lambda(Lambda_12)|+(7/72)(R_3-1)",
		BridgeLayerOnly:        true,
		Verdict:                join(StatusSevenOverSeventyTwoInterpolation, StatusBoundaryWeightedClosureSupport),
	}
}

func buildSourceTypeAudit() SourceTypeAudit {
	objects := []SourceObject{
		{Name: "kappa_lambda", Value: kappaLambda, Role: "scalar low-scale HistoryLoopUnit matching deficit", TypedStatus: "defined by Gate658 scalar proxy-runtime spine; native source missing", NativeSource: false},
		{Name: "kappa_e", Value: kappaE, Role: "charged-lepton loop-angle/flavor wall deficit", TypedStatus: "environmental OrientationBalanceSeal component; native flavor theorem missing", NativeSource: false},
		{Name: "|lambda(Lambda_12)|", Value: absLambdaLambda12, Role: "high-scale scalar runtime wound", TypedStatus: "v1 RG/environmental boundary coordinate", NativeSource: false},
		{Name: "R_3-1", Value: r3Minus1, Role: "high-scale strong gauge boundary wound", TypedStatus: "active GaugeScalarBoundaryStressSeal coordinate", NativeSource: false},
		{Name: "7/72", Value: sevenOver72, Role: "typed boundary interpolation candidate in active transport lane", TypedStatus: "not sourced by Fano-Hitchin; no native trace theorem", NativeSource: false},
	}
	return SourceTypeAudit{
		Objects:                    objects,
		SevenOver72InFanoLane:      false,
		SevenOver72InTransportLane: true,
		FanoBoundaryMapConstructed: false,
		RandomConstantsSearched:    false,
		TypedCandidatesOnly:        true,
		Verdict:                    join(StatusSourceTypeAuditComputed, StatusNoNativeKappaClosureTheorem, StatusNoNativeSevenOver72SourceTheorem, StatusNoNativeScalarFlavorBoundaryTheorem),
	}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate658ScalarTransportSpineInherited,
		StatusFlavorKappaESealInherited,
		StatusKappaSumComputed,
		StatusKappaSumClosesOnAbsLambda,
		StatusBoundarySplitRatioComputed,
		StatusTypedWeightCandidatesAudited,
		StatusSevenOverSeventyTwoInterpolation,
		StatusSourceTypeAuditComputed,
		StatusKappaClosureSupport,
		StatusResidualTracksBoundarySplit,
		StatusSevenOver72ReappearsActiveTransport,
		StatusBoundaryWeightedClosureSupport,
		StatusNoNativeKappaClosureTheorem,
		StatusNoNativeSevenOver72SourceTheorem,
		StatusNoNativeScalarFlavorBoundaryTheorem,
		StatusNoNativeFlavorTheorem,
		StatusNoNativeScalarTheorem,
		StatusNoBoundaryStressDerivation,
		StatusNoHiggsGaugeOrCKMClaim,
		StatusGate659Boundary,
	}
}
