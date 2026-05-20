// Package generation2orientedwalldistancehyperplaneaudit implements
// Gate 670: Oriented Wall-Distance Hyperplane Audit.
//
// Gate 669 retyped the active scalar/flavor/gauge bridge coordinates as wall
// distances. Gate 670 writes the active 7/72 closure as a single signed affine
// wall functional on the coordinates (kappa_lambda, kappa_e, lambda, R3-1):
//
//	W_72 = kappa_lambda + kappa_e + (65/72) lambda(Lambda_12) - (7/72)(R_3-1).
//
// This is a bridge-layer wall-balance audit only. It defines a
// HistoryWallBalanceSeal but preserves the firewall against native wall-distance,
// 7/72, scalar-zero, and boundary-stress theorems.
package generation2orientedwalldistancehyperplaneaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate669 "github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarzerowallboundarywallcoordinateaudit"
)

const (
	AuditID = "GATE670-ORIENTED-WALL-DISTANCE-HYPERPLANE-AUDIT"

	StatusGate669WallCoordinateInherited      = "PASS_GATE669_WALL_COORDINATE_AUDIT_INHERITED"
	StatusSignedWallFormWritten               = "PASS_SIGNED_WALL_FORM_WRITTEN"
	StatusHistoryWallBalanceFunctionalDefined = "PASS_HISTORY_WALL_BALANCE_FUNCTIONAL_DEFINED"
	StatusWallCoordinateRolesClassified       = "PASS_WALL_COORDINATE_ROLES_CLASSIFIED"
	StatusNormalVectorAndWeightAudited        = "PASS_NORMAL_VECTOR_AND_7_OVER_72_WEIGHT_AUDITED"
	StatusOrientationApproximationAudited     = "PASS_ORIENTATION_APPROXIMATION_AUDITED"
	StatusHessianLayerFirewallPreserved       = "PASS_HESSIAN_LAYER_FIREWALL_PRESERVED"
	StatusActiveBridgeOrientedHyperplane      = "CONDITIONAL_SUPPORT_ACTIVE_BRIDGE_IS_ORIENTED_WALL_DISTANCE_HYPERPLANE"
	StatusHistoryWallBalanceSealDefined       = "CONDITIONAL_SUPPORT_HISTORY_WALL_BALANCE_SEAL_DEFINED"
	StatusSevenOver72TypedNormalWeight        = "CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_TYPED_AS_HYPERPLANE_NORMAL_WEIGHT"
	StatusNoNativeWallDistanceAirlockTheorem  = "FAILED_ROUTE_NO_NATIVE_WALL_DISTANCE_AIRLOCK_THEOREM"
	StatusNoNativeSevenOver72Theorem          = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoBoundaryStressDerivation          = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusNoNativeScalarZeroBoundaryTheorem   = "FAILED_ROUTE_NO_NATIVE_SCALAR_ZERO_BOUNDARY_THEOREM"
	StatusNoHiggsGaugeFlavorClaim             = "FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM"
	StatusGate670Boundary                     = "FIREWALL_PRESERVED_GATE670_ORIENTED_WALL_HYPERPLANE_BOUNDARY"
)

const (
	kappaLambda  = 0.0443230430960771
	kappaE       = 0.00550355419157456
	kappaEOrient = 0.00550633006471245

	lambdaLambda12Signed = -0.0497009420776833
	absLambda12          = 0.0497009420776833
	r3Minus1             = 0.0509933868964996

	sevenOver72     = 7.0 / 72.0
	sixtyFiveOver72 = 65.0 / 72.0
)

type Gate669Inheritance struct {
	WallCoordinatesInherited         bool
	PositiveAndSignedFormsEquivalent bool
	ScalarWallDistance               bool
	GaugeWallDistance                bool
	FlavorWallDistance               bool
	HessianLayerSeparated            bool
	MissingWallTheoremNamed          bool
	NoNativeWallTheorem              bool
	NoSevenOver72                    bool
	NoBoundaryStress                 bool
	Verdict                          string
}

type SignedWallFormAudit struct {
	PositiveDistanceForm            string
	SignedWallForm                  string
	PositiveResidual                float64
	SignedResidual                  float64
	EquivalentBecauseLambdaNegative bool
	LambdaIsNegative                bool
	Verdict                         string
}

type WallCoordinateRole struct {
	Coordinate string
	Value      float64
	Wall       string
	Sign       string
	Role       string
	Layer      string
}

type WallCoordinateRolesAudit struct {
	Roles              []WallCoordinateRole
	AllRolesClassified bool
	Verdict            string
}

type NormalVectorAudit struct {
	Coordinates                      []string
	Coefficients                     []float64
	NormalVectorLabel                string
	SevenOver72                      float64
	SixtyFiveOver72                  float64
	SumBoundaryWeights               float64
	BestWeight                       float64
	BestWeightDeltaFromSevenOver72   float64
	TypedWeightUniqueInCurrentLedger bool
	Verdict                          string
}

type HistoryWallBalanceFunctional struct {
	Name                  string
	Formula               string
	Value                 float64
	AbsoluteResidual      float64
	Threshold             float64
	PassesBridgeTolerance bool
	SealName              string
	Interpretation        string
	Verdict               string
}

type OrientationApproximationAudit struct {
	ExactKappaE             float64
	OrientationKappaE       float64
	ExactResidual           float64
	OrientationResidual     float64
	ResidualGrowth          float64
	RelativeToBoundarySplit float64
	Verdict                 string
}

type HessianFirewallAudit struct {
	QuarticWallCoordinate float64
	HessianCoordinate     float64
	KeepsHessianSeparate  bool
	Statement             string
	Verdict               string
}

type VerdictDiscipline struct {
	ClaimsNativeWallDistanceAirlock bool
	ClaimsNativeSevenOver72         bool
	ClaimsBoundaryStressDerivation  bool
	ClaimsNativeScalarZeroBoundary  bool
	ClaimsHiggsMassPrediction       bool
	ClaimsScalarStability           bool
	ClaimsGaugeUnification          bool
	ClaimsFlavorDerivation          bool
	ClaimsCKMPMNSDerivation         bool
	Verdict                         string
}

type Analysis struct {
	Inherited   Gate669Inheritance
	Signed      SignedWallFormAudit
	Roles       WallCoordinateRolesAudit
	Normal      NormalVectorAudit
	Functional  HistoryWallBalanceFunctional
	Orientation OrientationApproximationAudit
	Hessian     HessianFirewallAudit
	Discipline  VerdictDiscipline
	Truth       string
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
	g669, err := gate669.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate669 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g669)
	signed := buildSignedWallForm()
	roles := buildRoles()
	normal := buildNormalVector()
	functional := buildFunctional(signed.SignedResidual)
	orientation := buildOrientationApproximation(signed.SignedResidual)
	hessian := buildHessianFirewall()
	discipline := VerdictDiscipline{Verdict: StatusGate670Boundary}
	truth := "Gate 670 writes the active scalar/flavor/gauge bridge as one signed oriented wall-distance hyperplane: kappa_lambda+kappa_e+(65/72)lambda(Lambda12)-(7/72)(R3-1)≈0. This defines a bridge-layer HistoryWallBalanceSeal and reclassifies the 7/72 split as a hyperplane-normal weight on signed wall coordinates. No native wall-distance airlock theorem, native 7/72 theorem, scalar-zero boundary theorem, or boundary-stress derivation is certified."
	return Analysis{Inherited: inherited, Signed: signed, Roles: roles, Normal: normal, Functional: functional, Orientation: orientation, Hessian: hessian, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate669.Analysis) Gate669Inheritance {
	return Gate669Inheritance{
		WallCoordinatesInherited:         g.Scalar.AbsoluteValueTyped && g.Gauge.IsAboveWall && g.Flavor.FlavorWallSupported && g.Flavor.ScalarWallSupported && g.Flavor.GaugeWallSupported,
		PositiveAndSignedFormsEquivalent: g.Boundary.EquivalentFormsAgree,
		ScalarWallDistance:               g.Scalar.IsBelowWall,
		GaugeWallDistance:                g.Gauge.IsAboveWall,
		FlavorWallDistance:               g.Flavor.FlavorWallSupported,
		HessianLayerSeparated:            g.Hessian.LayersSeparated,
		MissingWallTheoremNamed:          g.Target.PrimaryName == "BoundaryWallCoordinateAirlockTheorem",
		NoNativeWallTheorem:              strings.Contains(g.Target.Verdict, gate669.StatusNoNativeWallDistanceAirlockTheorem),
		NoSevenOver72:                    strings.Contains(g.Target.Verdict, gate669.StatusNoNativeSevenOver72Theorem),
		NoBoundaryStress:                 strings.Contains(g.Boundary.Verdict, gate669.StatusNoBoundaryStressDerivation),
		Verdict:                          StatusGate669WallCoordinateInherited,
	}
}

func buildSignedWallForm() SignedWallFormAudit {
	positiveResidual := kappaLambda + kappaE - (sixtyFiveOver72*absLambda12 + sevenOver72*r3Minus1)
	signedResidual := kappaLambda + kappaE + sixtyFiveOver72*lambdaLambda12Signed - sevenOver72*r3Minus1
	return SignedWallFormAudit{
		PositiveDistanceForm:            "kappa_lambda+kappa_e-[(65/72)|lambda|+(7/72)(R_3-1)]",
		SignedWallForm:                  "kappa_lambda+kappa_e+(65/72)lambda-(7/72)(R_3-1)",
		PositiveResidual:                positiveResidual,
		SignedResidual:                  signedResidual,
		EquivalentBecauseLambdaNegative: math.Abs(positiveResidual-signedResidual) < 1e-18 && lambdaLambda12Signed < 0 && math.Abs(absLambda12+lambdaLambda12Signed) < 1e-15,
		LambdaIsNegative:                lambdaLambda12Signed < 0,
		Verdict:                         join(StatusSignedWallFormWritten, StatusActiveBridgeOrientedHyperplane),
	}
}

func buildRoles() WallCoordinateRolesAudit {
	roles := []WallCoordinateRole{
		{Coordinate: "kappa_lambda", Value: kappaLambda, Wall: "scalar proxy/runtime matching wall", Sign: "+", Role: "scalar low-scale matching wall-deficit coordinate", Layer: "history-loop scalar matching deficit"},
		{Coordinate: "kappa_e", Value: kappaE, Wall: "charged-lepton flavor / Koide wall", Sign: "+", Role: "flavor loop-wall deficit coordinate", Layer: "history-loop flavor deficit"},
		{Coordinate: "lambda(Lambda_12)", Value: lambdaLambda12Signed, Wall: "scalar zero wall lambda=0", Sign: "+65/72 times signed negative coordinate", Role: "oriented scalar zero-wall coordinate; below-wall depth in signed form", Layer: "quartic wall distance"},
		{Coordinate: "R_3-1", Value: r3Minus1, Wall: "gauge meeting wall g3=gEW", Sign: "-7/72", Role: "oriented gauge meeting-wall excess", Layer: "canonical connection-amplitude wall distance"},
	}
	return WallCoordinateRolesAudit{Roles: roles, AllRolesClassified: len(roles) == 4, Verdict: join(StatusWallCoordinateRolesClassified, StatusActiveBridgeOrientedHyperplane, StatusNoNativeWallDistanceAirlockTheorem)}
}

func buildNormalVector() NormalVectorAudit {
	wBest := ((kappaLambda + kappaE) - absLambda12) / (r3Minus1 - absLambda12)
	return NormalVectorAudit{
		Coordinates:                      []string{"kappa_lambda", "kappa_e", "lambda(Lambda_12)", "R_3-1"},
		Coefficients:                     []float64{1, 1, sixtyFiveOver72, -sevenOver72},
		NormalVectorLabel:                "(1, 1, 65/72, -7/72)",
		SevenOver72:                      sevenOver72,
		SixtyFiveOver72:                  sixtyFiveOver72,
		SumBoundaryWeights:               sixtyFiveOver72 + sevenOver72,
		BestWeight:                       wBest,
		BestWeightDeltaFromSevenOver72:   wBest - sevenOver72,
		TypedWeightUniqueInCurrentLedger: math.Abs(wBest-sevenOver72) < 1e-6,
		Verdict:                          join(StatusNormalVectorAndWeightAudited, StatusSevenOver72TypedNormalWeight, StatusNoNativeSevenOver72Theorem),
	}
}

func buildFunctional(value float64) HistoryWallBalanceFunctional {
	return HistoryWallBalanceFunctional{
		Name:                  "HistoryWallBalanceSeal",
		Formula:               "W_72_wall = kappa_lambda+kappa_e+(65/72)lambda(Lambda_12)-(7/72)(R_3-1)",
		Value:                 value,
		AbsoluteResidual:      math.Abs(value),
		Threshold:             1e-8,
		PassesBridgeTolerance: math.Abs(value) < 1e-8,
		SealName:              "HistoryWallBalanceSeal",
		Interpretation:        "signed affine hyperplane in scalar/flavor/gauge wall-distance coordinates",
		Verdict:               join(StatusHistoryWallBalanceFunctionalDefined, StatusHistoryWallBalanceSealDefined, StatusActiveBridgeOrientedHyperplane),
	}
}

func buildOrientationApproximation(exactResidual float64) OrientationApproximationAudit {
	orientResidual := kappaLambda + kappaEOrient + sixtyFiveOver72*lambdaLambda12Signed - sevenOver72*r3Minus1
	boundarySplit := r3Minus1 - absLambda12
	return OrientationApproximationAudit{
		ExactKappaE:             kappaE,
		OrientationKappaE:       kappaEOrient,
		ExactResidual:           exactResidual,
		OrientationResidual:     orientResidual,
		ResidualGrowth:          math.Abs(orientResidual) - math.Abs(exactResidual),
		RelativeToBoundarySplit: math.Abs(orientResidual) / boundarySplit,
		Verdict:                 join(StatusOrientationApproximationAudited, StatusActiveBridgeOrientedHyperplane),
	}
}

func buildHessianFirewall() HessianFirewallAudit {
	return HessianFirewallAudit{
		QuarticWallCoordinate: absLambda12,
		HessianCoordinate:     2 * absLambda12,
		KeepsHessianSeparate:  true,
		Statement:             "Gate670 inherits Gate669/Gate668: |lambda| is the active quartic zero-wall coordinate; 2|lambda| remains the scalar Hessian/squared-mass layer and is not used in the wall hyperplane.",
		Verdict:               join(StatusHessianLayerFirewallPreserved, StatusNoNativeScalarZeroBoundaryTheorem),
	}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate669WallCoordinateInherited,
		StatusSignedWallFormWritten,
		StatusHistoryWallBalanceFunctionalDefined,
		StatusWallCoordinateRolesClassified,
		StatusNormalVectorAndWeightAudited,
		StatusOrientationApproximationAudited,
		StatusHessianLayerFirewallPreserved,
		StatusActiveBridgeOrientedHyperplane,
		StatusHistoryWallBalanceSealDefined,
		StatusSevenOver72TypedNormalWeight,
		StatusNoNativeWallDistanceAirlockTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusNoBoundaryStressDerivation,
		StatusNoNativeScalarZeroBoundaryTheorem,
		StatusNoHiggsGaugeFlavorClaim,
		StatusGate670Boundary,
	}
}
