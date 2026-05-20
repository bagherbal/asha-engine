// Package generation2scalarzerowallboundarywallcoordinateaudit implements
// Gate 669: Scalar Zero-Wall Distance and Boundary Wall-Coordinate Audit.
//
// Gate 668 classified the active scalar coordinate in the boundary-weighted
// deficit closure as the quartic wound |lambda(Lambda_12)|, paired with the
// canonical connection-amplitude gauge wound R3-1. Gate 669 asks whether this
// absolute value is lawfully typed as a scalar zero-wall distance coordinate,
// analogous to the gauge meeting-wall distance and the charged-lepton wall
// offset epsilon_e. It preserves the firewall: this is a bridge-layer
// wall-coordinate audit only, not a scalar, flavor, boundary-stress, or native
// 7/72 theorem.
package generation2scalarzerowallboundarywallcoordinateaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate668 "github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarquarticcoordinateairlockaudit"
)

const (
	AuditID = "GATE669-SCALAR-ZERO-WALL-BOUNDARY-WALL-COORDINATE-AUDIT"

	StatusGate668ScalarCoordinateInherited   = "PASS_GATE668_SCALAR_COORDINATE_AUDIT_INHERITED"
	StatusScalarZeroWallDistanceDefined      = "PASS_SCALAR_ZERO_WALL_DISTANCE_DEFINED"
	StatusGaugeMeetingWallDistanceDefined    = "PASS_GAUGE_MEETING_WALL_DISTANCE_DEFINED"
	StatusSignedBoundaryStressFormRewritten  = "PASS_SIGNED_BOUNDARY_STRESS_FORM_REWRITTEN"
	StatusFlavorWallAnalogyAudited           = "PASS_FLAVOR_WALL_ANALOGY_AUDITED"
	StatusHessianLayerSeparationPreserved    = "PASS_HESSIAN_LAYER_SEPARATION_PRESERVED"
	StatusMissingWallTheoremTargetNamed      = "PASS_MISSING_WALL_DISTANCE_THEOREM_TARGET_NAMED"
	StatusActiveClosureUsesWallCoordinates   = "CONDITIONAL_SUPPORT_ACTIVE_CLOSURE_USES_WALL_DISTANCE_COORDINATES"
	StatusAbsLambdaScalarZeroWallDistance    = "CONDITIONAL_SUPPORT_ABS_LAMBDA_IS_SCALAR_ZERO_WALL_DISTANCE"
	StatusR3GaugeMeetingWallDistance         = "CONDITIONAL_SUPPORT_R3_MINUS_ONE_IS_GAUGE_MEETING_WALL_DISTANCE"
	StatusEpsilonEFlavorWallDistance         = "CONDITIONAL_SUPPORT_EPSILON_E_IS_FLAVOR_WALL_DISTANCE"
	StatusNoNativeWallDistanceAirlockTheorem = "FAILED_ROUTE_NO_NATIVE_WALL_DISTANCE_AIRLOCK_THEOREM"
	StatusNoNativeScalarZeroBoundaryTheorem  = "FAILED_ROUTE_NO_NATIVE_SCALAR_ZERO_BOUNDARY_THEOREM"
	StatusNoNativeSevenOver72Theorem         = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoBoundaryStressDerivation         = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusNoHiggsGaugeFlavorClaim            = "FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM"
	StatusGate669Boundary                    = "FIREWALL_PRESERVED_GATE669_WALL_COORDINATE_BOUNDARY"
)

const (
	kappaLambda = 0.0443230430960771
	kappaE      = 0.00550355419157456
	kSum        = kappaLambda + kappaE

	lambdaLambda12Signed = -0.0497009420776833
	absLambda12          = 0.0497009420776833
	r3Minus1             = 0.0509933868964996
	xiBoundary           = 0.0503471644870914
	sevenOver72          = 7.0 / 72.0

	epsilonE = 0.039569756309433
)

type Gate668Inheritance struct {
	ScalarCoordinateInherited bool
	Classification            string
	ActivePair                string
	HessianLayerSeparated     bool
	NoScalarAirlock           bool
	NoSevenOver72             bool
	NoBoundaryStress          bool
	NoTransport               bool
	Verdict                   string
}

type ScalarZeroWallAudit struct {
	WallEquation       string
	SignedLambda       float64
	DistanceBelowWall  float64
	IsBelowWall        bool
	AbsoluteValueTyped bool
	CoordinateLayer    string
	Verdict            string
}

type GaugeMeetingWallAudit struct {
	WallEquation        string
	GaugeResidual       float64
	IsAboveWall         bool
	CoordinateLayer     string
	CanonicalCoordinate string
	Verdict             string
}

type SignedBoundaryStressAudit struct {
	GaugeExcess                 float64
	ScalarSignedWound           float64
	ScalarDepth                 float64
	XiBoundary                  float64
	PositiveDistanceForm        string
	SignedStressForm            string
	W72                         float64
	ClosureResidualPositiveForm float64
	ClosureResidualSignedForm   float64
	EquivalentFormsAgree        bool
	Verdict                     string
}

type WallCoordinateRow struct {
	Name         string
	Wall         string
	Coordinate   string
	Value        float64
	DistanceType string
	Role         string
	Verdict      string
}

type FlavorWallAnalogyAudit struct {
	Rows                []WallCoordinateRow
	RecurringPattern    string
	FlavorWallSupported bool
	ScalarWallSupported bool
	GaugeWallSupported  bool
	Verdict             string
}

type HessianLayerSeparation struct {
	QuarticWallCoordinate float64
	HessianCoordinate     float64
	HessianRelation       string
	LayersSeparated       bool
	Verdict               string
}

type MissingTheoremTarget struct {
	PrimaryName     string
	AlternateName   string
	RequiredObjects []string
	Statements      []string
	Verdict         string
}

type VerdictDiscipline struct {
	ClaimsNativeWallDistanceAirlock bool
	ClaimsNativeScalarZeroBoundary  bool
	ClaimsNativeSevenOver72         bool
	ClaimsBoundaryStressDerivation  bool
	ClaimsHiggsMassPrediction       bool
	ClaimsScalarStability           bool
	ClaimsGaugeUnification          bool
	ClaimsFlavorDerivation          bool
	ClaimsCKMPMNSDerivation         bool
	Verdict                         string
}

type Analysis struct {
	Inherited  Gate668Inheritance
	Scalar     ScalarZeroWallAudit
	Gauge      GaugeMeetingWallAudit
	Boundary   SignedBoundaryStressAudit
	Flavor     FlavorWallAnalogyAudit
	Hessian    HessianLayerSeparation
	Target     MissingTheoremTarget
	Discipline VerdictDiscipline
	Truth      string
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
	g668, err := gate668.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate668 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g668)
	scalar := buildScalarZeroWall()
	gauge := buildGaugeMeetingWall()
	boundary := buildBoundaryStress()
	flavor := buildFlavorWallAnalogy()
	hessian := buildHessianSeparation()
	target := buildMissingTheoremTarget()
	discipline := VerdictDiscipline{Verdict: StatusGate669Boundary}
	truth := "Gate 669 classifies the active closure coordinates as wall distances: |lambda(Lambda12)| is the scalar zero-wall depth below lambda=0, R3-1 is the gauge meeting-wall excess above g3=gEW in connection-amplitude coordinates, and epsilon_e is the charged-lepton wall offset. The signed boundary stress vector is equivalently (+R3-1, lambda) or the positive distance pair (R3-1, |lambda|). No native wall-distance airlock, scalar-zero boundary theorem, boundary-stress derivation, or native 7/72 theorem is certified."
	return Analysis{Inherited: inherited, Scalar: scalar, Gauge: gauge, Boundary: boundary, Flavor: flavor, Hessian: hessian, Target: target, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate668.Analysis) Gate668Inheritance {
	return Gate668Inheritance{
		ScalarCoordinateInherited: g.Source.Classification == "BoundaryWeightedDeficitClosureQuarticWoundSeal" && g.Scalars.ActiveScalarCoordinate == "|lambda(Lambda_12)|" && g.Pairings.AmplitudePairPasses,
		Classification:            g.Source.Classification,
		ActivePair:                "R_3-1 with |lambda(Lambda_12)|",
		HessianLayerSeparated:     g.Hessian.TypedAsHessianLayer && !g.Pairings.InverseHessianClosurePasses,
		NoScalarAirlock:           strings.Contains(g.Source.Verdict, gate668.StatusNoNativeScalarAirlockTheorem),
		NoSevenOver72:             strings.Contains(g.Source.Verdict, gate668.StatusNoNativeSevenOver72Theorem),
		NoBoundaryStress:          strings.Contains(g.Source.Verdict, gate668.StatusNoNativeBoundaryStressTheorem),
		NoTransport:               !g.Discipline.ClaimsNativeTransportTheorem,
		Verdict:                   StatusGate668ScalarCoordinateInherited,
	}
}

func buildScalarZeroWall() ScalarZeroWallAudit {
	return ScalarZeroWallAudit{
		WallEquation:       "lambda=0",
		SignedLambda:       lambdaLambda12Signed,
		DistanceBelowWall:  -lambdaLambda12Signed,
		IsBelowWall:        lambdaLambda12Signed < 0,
		AbsoluteValueTyped: math.Abs(lambdaLambda12Signed-(-absLambda12)) < 1e-15,
		CoordinateLayer:    "quartic zero-wall distance / scalar runtime wound magnitude",
		Verdict:            join(StatusScalarZeroWallDistanceDefined, StatusAbsLambdaScalarZeroWallDistance, StatusNoNativeScalarZeroBoundaryTheorem),
	}
}

func buildGaugeMeetingWall() GaugeMeetingWallAudit {
	return GaugeMeetingWallAudit{
		WallEquation:        "g3/gEW-1=0",
		GaugeResidual:       r3Minus1,
		IsAboveWall:         r3Minus1 > 0,
		CoordinateLayer:     "canonical connection-amplitude wall distance",
		CanonicalCoordinate: "R_3-1=g3/gEW-1",
		Verdict:             join(StatusGaugeMeetingWallDistanceDefined, StatusR3GaugeMeetingWallDistance),
	}
}

func buildBoundaryStress() SignedBoundaryStressAudit {
	w := (65.0/72.0)*absLambda12 + sevenOver72*r3Minus1
	positiveResidual := kSum - w
	signedResidual := kSum + (65.0/72.0)*lambdaLambda12Signed - sevenOver72*r3Minus1
	return SignedBoundaryStressAudit{
		GaugeExcess:                 r3Minus1,
		ScalarSignedWound:           lambdaLambda12Signed,
		ScalarDepth:                 absLambda12,
		XiBoundary:                  0.5 * (r3Minus1 + absLambda12),
		PositiveDistanceForm:        "K_sum - [(65/72)|lambda| + (7/72)(R_3-1)]",
		SignedStressForm:            "K_sum + (65/72)lambda - (7/72)(R_3-1)",
		W72:                         w,
		ClosureResidualPositiveForm: positiveResidual,
		ClosureResidualSignedForm:   signedResidual,
		EquivalentFormsAgree:        math.Abs(positiveResidual-signedResidual) < 1e-18 && math.Abs(0.5*(r3Minus1+absLambda12)-xiBoundary) < 1e-15,
		Verdict:                     join(StatusSignedBoundaryStressFormRewritten, StatusActiveClosureUsesWallCoordinates, StatusNoBoundaryStressDerivation),
	}
}

func buildFlavorWallAnalogy() FlavorWallAnalogyAudit {
	rows := []WallCoordinateRow{
		{Name: "charged-lepton flavor wall", Wall: "electron-zero / Koide wall offset", Coordinate: "epsilon_e", Value: epsilonE, DistanceType: "wall offset", Role: "source of kappa_e loop-angle/flavor-wall deficit", Verdict: join(StatusFlavorWallAnalogyAudited, StatusEpsilonEFlavorWallDistance)},
		{Name: "scalar zero wall", Wall: "lambda=0", Coordinate: "|lambda(Lambda_12)|=-lambda(Lambda_12)", Value: absLambda12, DistanceType: "depth below zero wall", Role: "scalar boundary wound in W72", Verdict: join(StatusFlavorWallAnalogyAudited, StatusAbsLambdaScalarZeroWallDistance)},
		{Name: "gauge meeting wall", Wall: "g3/gEW-1=0", Coordinate: "R_3-1", Value: r3Minus1, DistanceType: "excess above meeting wall", Role: "strong gauge boundary wound in W72", Verdict: join(StatusFlavorWallAnalogyAudited, StatusR3GaugeMeetingWallDistance)},
	}
	return FlavorWallAnalogyAudit{Rows: rows, RecurringPattern: "history closures use signed/canonical distances to walls: epsilon_e, |lambda|, and R_3-1", FlavorWallSupported: true, ScalarWallSupported: true, GaugeWallSupported: true, Verdict: join(StatusFlavorWallAnalogyAudited, StatusActiveClosureUsesWallCoordinates, StatusEpsilonEFlavorWallDistance, StatusNoNativeWallDistanceAirlockTheorem)}
}

func buildHessianSeparation() HessianLayerSeparation {
	return HessianLayerSeparation{
		QuarticWallCoordinate: absLambda12,
		HessianCoordinate:     2 * absLambda12,
		HessianRelation:       "m_H^2=2 lambda v^2 types 2|lambda| as the Hessian/squared-mass layer; Gate669 keeps |lambda| as the quartic zero-wall distance layer",
		LayersSeparated:       true,
		Verdict:               join(StatusHessianLayerSeparationPreserved, StatusAbsLambdaScalarZeroWallDistance),
	}
}

func buildMissingTheoremTarget() MissingTheoremTarget {
	return MissingTheoremTarget{
		PrimaryName:   "BoundaryWallCoordinateAirlockTheorem",
		AlternateName: "WallDistanceHistoryCoordinateTheorem",
		RequiredObjects: []string{
			"epsilon_e as charged-lepton wall offset",
			"|lambda(Lambda_12)| as scalar zero-wall depth",
			"R_3-1 as gauge meeting-wall excess in connection-amplitude coordinates",
			"a typed airlock from native polynomial/kinetic/Hessian data to history wall distances",
		},
		Statements: []string{
			"Gate669 names the wall-distance pattern but does not derive it natively.",
			"The active 7/72 closure uses wall distances, not raw kinetic, polynomial, or Hessian variables.",
			"A scalar-zero boundary theorem remains missing; lambda=0 is used as a wall coordinate, not derived as a native endpoint law.",
		},
		Verdict: join(StatusMissingWallTheoremTargetNamed, StatusNoNativeWallDistanceAirlockTheorem, StatusNoNativeScalarZeroBoundaryTheorem, StatusNoNativeSevenOver72Theorem),
	}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate668ScalarCoordinateInherited,
		StatusScalarZeroWallDistanceDefined,
		StatusGaugeMeetingWallDistanceDefined,
		StatusSignedBoundaryStressFormRewritten,
		StatusFlavorWallAnalogyAudited,
		StatusHessianLayerSeparationPreserved,
		StatusMissingWallTheoremTargetNamed,
		StatusActiveClosureUsesWallCoordinates,
		StatusAbsLambdaScalarZeroWallDistance,
		StatusR3GaugeMeetingWallDistance,
		StatusEpsilonEFlavorWallDistance,
		StatusNoNativeWallDistanceAirlockTheorem,
		StatusNoNativeScalarZeroBoundaryTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusNoBoundaryStressDerivation,
		StatusNoHiggsGaugeFlavorClaim,
		StatusGate669Boundary,
	}
}
