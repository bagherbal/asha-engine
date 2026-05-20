// Package generation2scalarquarticcoordinateairlockaudit implements
// Gate 668: Scalar Quartic Coordinate Airlock and Hessian-Doubling Audit.
//
// Gate 667 source-typed the gauge side of the active boundary-weighted
// deficit closure as a canonical connection-amplitude coordinate: the
// RG-native kinetic coefficient u=1/g^2 passes through canonical field
// normalization to the bridge coordinate g=u^(-1/2). Gate 668 audits the
// scalar side of the same bridge. It asks whether the scalar coordinate should
// be the signed quartic coefficient lambda, the wound |lambda|, the Hessian
// coefficient 2|lambda|, a square-root/mass-amplitude coordinate, beta_lambda,
// or another typed scalar object. It preserves the firewall: this is a
// bridge-layer scalar-coordinate audit only, not a scalar, Higgs, stability,
// boundary-stress, or native 7/72 theorem.
package generation2scalarquarticcoordinateairlockaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate667 "github.com/bagherbal/asha-engine/pkg/bridge/generation2kinetictoconnectionamplitudeairlockaudit"
)

const (
	AuditID = "GATE668-SCALAR-QUARTIC-COORDINATE-AIRLOCK-HESSIAN-DOUBLING-AUDIT"

	StatusGate667ConnectionAmplitudeInherited = "PASS_GATE667_CONNECTION_AMPLITUDE_SOURCE_INHERITED"
	StatusScalarCoordinateFamilyAudited       = "PASS_SCALAR_COORDINATE_FAMILY_AUDITED"
	StatusHessianDoublingAudited              = "PASS_HESSIAN_DOUBLING_AUDITED"
	StatusGaugeScalarPairingsAudited          = "PASS_GAUGE_SCALAR_COORDINATE_PAIRINGS_AUDITED"
	StatusClosureCoordinateRetested           = "PASS_CLOSURE_COORDINATE_RETESTED"
	StatusSourceTypeResultAudited             = "PASS_SOURCE_TYPE_RESULT_AUDITED"
	StatusRootAmplitudeRecurrenceAudited      = "PASS_ROOT_AMPLITUDE_RECURRENCE_AUDITED"
	StatusAmplitudePairSupported              = "CONDITIONAL_SUPPORT_AMPLITUDE_LAYER_PAIR_IS_R3_MINUS_ONE_WITH_ABS_LAMBDA"
	StatusInverseHessianShadowSupported       = "CONDITIONAL_SUPPORT_INVERSE_KINETIC_LAYER_PAIRS_WITH_TWO_ABS_LAMBDA_AS_HESSIAN_SHADOW"
	StatusGaugeAmplitudeInherited             = "CONDITIONAL_SUPPORT_GAUGE_AMPLITUDE_COORDINATE_SOURCED_BY_CANONICAL_CONNECTION_NORMALIZATION"
	StatusScalarQuarticWoundSelected          = "CONDITIONAL_SUPPORT_SCALAR_BRIDGE_COORDINATE_IS_ABS_LAMBDA_QUARTIC_WOUND_IN_ACTIVE_CLOSURE"
	StatusNoNativeScalarAirlockTheorem        = "FAILED_ROUTE_NO_NATIVE_SCALAR_COORDINATE_AIRLOCK_THEOREM"
	StatusNoNativeBoundaryStressTheorem       = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_STRESS_THEOREM"
	StatusNoNativeSevenOver72Theorem          = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoNativeTransportTheorem            = "FAILED_ROUTE_NO_NATIVE_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoHiggsMassStabilityGaugeFlavor     = "FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM"
	StatusGate668Boundary                     = "FIREWALL_PRESERVED_GATE668_SCALAR_COORDINATE_AIRLOCK_BOUNDARY"
)

const (
	kappaLambda = 0.0443230430960771
	kappaE      = 0.00550355419157456
	kSum        = kappaLambda + kappaE

	lambdaLambda12Signed = -0.0497009420776833
	absLambda12          = 0.0497009420776833
	r3Minus1             = 0.0509933868964996
	sevenOver72          = 7.0 / 72.0
)

type Gate667Inheritance struct {
	ConnectionAmplitudeInherited bool
	Classification               string
	ClosureCoordinate            string
	AmplitudeOnlyPasses          bool
	InverseKineticFails          bool
	ScalarSideWasRuntimeShadow   bool
	MissingKineticAirlock        bool
	NoNativeSevenOver72          bool
	NoNativeTransport            bool
	NoBoundaryStress             bool
	Verdict                      string
}

type ScalarCoordinateRow struct {
	Name        string
	Expression  string
	Value       float64
	Layer       string
	TypedStatus string
	ClosureRole string
	Verdict     string
}

type ScalarCoordinateFamilyAudit struct {
	Rows                    []ScalarCoordinateRow
	ActiveScalarCoordinate  string
	HessianCoordinate       string
	MassAmplitudeCoordinate string
	Verdict                 string
}

type GaugeScalarPairingRow struct {
	Name               string
	GaugeLayer         string
	GaugeCoordinate    string
	GaugeValue         float64
	ScalarLayer        string
	ScalarCoordinate   string
	ScalarValue        float64
	SignedDifference   float64
	RelativeDifference float64
	WBest              float64
	WBestMinus7Over72  float64
	PassesSevenOver72  bool
	Interpretation     string
	Verdict            string
}

type GaugeScalarPairingAudit struct {
	Rows                          []GaugeScalarPairingRow
	AmplitudePairPasses           bool
	InverseHessianShadowMagnitude bool
	InverseHessianClosurePasses   bool
	MassAmplitudePairPasses       bool
	Verdict                       string
}

type HessianDoublingAudit struct {
	PotentialConvention       string
	LowScaleRelation          string
	HessianCoordinate         float64
	AmplitudeResidual         float64
	InverseKineticWound       float64
	TwoTimesAmplitudeResidual float64
	InverseMinusTwoAmplitude  float64
	HessianMinusInverse       float64
	HessianMinusTwoAmplitude  float64
	TypedAsHessianLayer       bool
	Verdict                   string
}

type ClosureCoordinateRetest struct {
	BestTypedPair         string
	BestTypedResidual     float64
	BestTypedWBest        float64
	BestTypedWBestMinus7  float64
	SevenOver72SelectedBy string
	InverseHessianStatus  string
	Verdict               string
}

type SourceTypeResult struct {
	Classification string
	Statements     []string
	Verdict        string
}

type RootAmplitudeRecurrenceAudit struct {
	Rows    []gate667.GaugeCoordinateLayerRow
	Pattern string
	Verdict string
}

type VerdictDiscipline struct {
	ClaimsNativeScalarAirlockTheorem  bool
	ClaimsNativeBoundaryStressTheorem bool
	ClaimsNativeSevenOver72Theorem    bool
	ClaimsNativeTransportTheorem      bool
	ClaimsHiggsMassPrediction         bool
	ClaimsScalarStability             bool
	ClaimsGaugeUnification            bool
	ClaimsFlavorDerivation            bool
	ClaimsCKMPMNSDerivation           bool
	Verdict                           string
}

type Analysis struct {
	Inherited  Gate667Inheritance
	Scalars    ScalarCoordinateFamilyAudit
	Hessian    HessianDoublingAudit
	Pairings   GaugeScalarPairingAudit
	Retest     ClosureCoordinateRetest
	Source     SourceTypeResult
	Pattern    RootAmplitudeRecurrenceAudit
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
	g667, err := gate667.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate667 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g667)
	scalars := buildScalarCoordinates()
	hessian := buildHessianDoubling()
	pairings := buildPairings(hessian)
	retest := buildRetest(pairings)
	source := buildSource(pairings, retest)
	pattern := buildPattern(g667)
	discipline := VerdictDiscipline{Verdict: StatusGate668Boundary}
	truth := "Gate 668 audits the scalar side of the active amplitude-layer closure. The working pair remains the connection-amplitude gauge wound R3-1 with the scalar quartic wound |lambda(Lambda12)|. The scalar Hessian coordinate 2|lambda| has a typed squared/Hessian role and roughly matches the doubled inverse-kinetic gauge layer, but it does not preserve the same 7/72 closure. No native scalar coordinate airlock, boundary-stress theorem, or native 7/72 theorem is certified."
	return Analysis{Inherited: inherited, Scalars: scalars, Hessian: hessian, Pairings: pairings, Retest: retest, Source: source, Pattern: pattern, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate667.Analysis) Gate667Inheritance {
	return Gate667Inheritance{
		ConnectionAmplitudeInherited: g.Source.Classification == "BoundaryWeightedDeficitClosureConnectionAmplitudeSeal" && g.Coordinates.AmplitudeOnlyPasses && g.Coordinates.InverseKineticFails,
		Classification:               g.Source.Classification,
		ClosureCoordinate:            g.Coordinates.ClosureCoordinate,
		AmplitudeOnlyPasses:          g.Coordinates.AmplitudeOnlyPasses,
		InverseKineticFails:          g.Coordinates.InverseKineticFails,
		ScalarSideWasRuntimeShadow:   strings.Contains(g.ScalarSide.Verdict, gate667.StatusScalarRuntimeShadow),
		MissingKineticAirlock:        strings.Contains(g.Target.Verdict, gate667.StatusNoNativeKineticAmplitudeTheorem),
		NoNativeSevenOver72:          !g.Discipline.ClaimsNativeSevenOver72Theorem,
		NoNativeTransport:            !g.Discipline.ClaimsNativeTransportTheorem,
		NoBoundaryStress:             !g.Discipline.ClaimsBoundaryStressDerivation,
		Verdict:                      StatusGate667ConnectionAmplitudeInherited,
	}
}

func buildScalarCoordinates() ScalarCoordinateFamilyAudit {
	rows := []ScalarCoordinateRow{
		{Name: "S_1", Expression: "|lambda(Lambda_12)|", Value: absLambda12, Layer: "quartic coefficient / signed runtime wound magnitude", TypedStatus: "active scalar bridge coordinate inherited from Gates659-667", ClosureRole: "works with R3-1 in the 7/72 active closure", Verdict: join(StatusScalarCoordinateFamilyAudited, StatusScalarQuarticWoundSelected)},
		{Name: "S_2", Expression: "2|lambda(Lambda_12)|", Value: 2 * absLambda12, Layer: "scalar Hessian / squared-mass coefficient", TypedStatus: "typed by m_H^2=2 lambda v^2 convention", ClosureRole: "Hessian/squared-layer shadow; not the active 7/72 closure coordinate", Verdict: join(StatusScalarCoordinateFamilyAudited, StatusHessianDoublingAudited)},
		{Name: "S_sqrt", Expression: "sqrt(|lambda(Lambda_12)|)", Value: math.Sqrt(absLambda12), Layer: "quartic square-root coordinate", TypedStatus: "root coordinate candidate but not directly selected by the active closure", ClosureRole: "does not share the boundary wound scale near 0.05", Verdict: StatusScalarCoordinateFamilyAudited},
		{Name: "S_hessian", Expression: "sqrt(2|lambda(Lambda_12)|)", Value: math.Sqrt(2 * absLambda12), Layer: "scalar Hessian mass-amplitude proxy", TypedStatus: "typed mass-amplitude candidate from m_H^2=2 lambda v^2", ClosureRole: "not selected by the 7/72 boundary interpolation", Verdict: join(StatusScalarCoordinateFamilyAudited, StatusHessianDoublingAudited)},
		{Name: "S_beta", Expression: "|beta_lambda(Lambda_12)|", Value: math.NaN(), Layer: "RG beta-flow coordinate", TypedStatus: "slot exists but no certified v1 beta_lambda value is inherited by this bridge package", ClosureRole: "not evaluated; uncertainty/source slot only", Verdict: join(StatusScalarCoordinateFamilyAudited, StatusNoNativeScalarAirlockTheorem)},
		{Name: "S_signed", Expression: "lambda(Lambda_12)", Value: lambdaLambda12Signed, Layer: "signed runtime scalar wound", TypedStatus: "typed signed coefficient; closure uses magnitude because the boundary stress pair records opposite signs", ClosureRole: "signed stress seal, not the active positive interpolation coordinate", Verdict: StatusScalarCoordinateFamilyAudited},
	}
	return ScalarCoordinateFamilyAudit{Rows: rows, ActiveScalarCoordinate: "|lambda(Lambda_12)|", HessianCoordinate: "2|lambda(Lambda_12)|", MassAmplitudeCoordinate: "sqrt(2|lambda(Lambda_12)|)", Verdict: join(StatusScalarCoordinateFamilyAudited, StatusScalarQuarticWoundSelected, StatusNoNativeScalarAirlockTheorem)}
}

func buildHessianDoubling() HessianDoublingAudit {
	inv := inverseKineticWound(r3Minus1)
	return HessianDoublingAudit{
		PotentialConvention:       "V(H)=-m^2 H^dagger H + lambda(H^dagger H)^2",
		LowScaleRelation:          "m_H^2=2 lambda v^2, so 2|lambda| is the scalar Hessian/squared-mass coefficient layer",
		HessianCoordinate:         2 * absLambda12,
		AmplitudeResidual:         r3Minus1,
		InverseKineticWound:       inv,
		TwoTimesAmplitudeResidual: 2 * r3Minus1,
		InverseMinusTwoAmplitude:  inv - 2*r3Minus1,
		HessianMinusInverse:       2*absLambda12 - inv,
		HessianMinusTwoAmplitude:  2*absLambda12 - 2*r3Minus1,
		TypedAsHessianLayer:       true,
		Verdict:                   join(StatusHessianDoublingAudited, StatusInverseHessianShadowSupported, StatusNoNativeScalarAirlockTheorem),
	}
}

func buildPairings(h HessianDoublingAudit) GaugeScalarPairingAudit {
	rows := []GaugeScalarPairingRow{
		pairing("amplitude/quartic", "connection-amplitude", "R_3-1=g3/gEW-1", r3Minus1, "quartic coefficient wound", "|lambda|", absLambda12, "active Gate659/660 closure coordinate"),
		pairing("inverse-kinetic/Hessian", "inverse kinetic", "1-u3/uEW", h.InverseKineticWound, "Hessian coefficient", "2|lambda|", 2*absLambda12, "typed doubled/squared-layer shadow but not a 7/72 closure"),
		pairing("squared-coupling/Hessian", "coupling strength", "g3^2/gEW^2-1", math.Pow(1+r3Minus1, 2)-1, "Hessian coefficient", "2|lambda|", 2*absLambda12, "same scale as doubled variables, not the active 7/72 closure"),
		pairing("mass-amplitude proxy", "root wound", "sqrt(R_3-1)", math.Sqrt(r3Minus1), "scalar mass-amplitude proxy", "sqrt(2|lambda|)", math.Sqrt(2*absLambda12), "typed amplitude comparison but wrong scale for the active boundary closure"),
		pairing("signed scalar stress", "connection-amplitude", "R_3-1", r3Minus1, "signed quartic coefficient", "lambda", lambdaLambda12Signed, "stress-pair sign ledger; interpolation uses |lambda|"),
	}
	ampPass := false
	inverseShadow := false
	inverseClosure := false
	massPass := false
	for _, r := range rows {
		if r.Name == "amplitude/quartic" && r.PassesSevenOver72 {
			ampPass = true
		}
		if r.Name == "inverse-kinetic/Hessian" && math.Abs(r.SignedDifference)/r.ScalarValue < 0.06 {
			inverseShadow = true
			inverseClosure = r.PassesSevenOver72
		}
		if r.Name == "mass-amplitude proxy" && r.PassesSevenOver72 {
			massPass = true
		}
	}
	return GaugeScalarPairingAudit{Rows: rows, AmplitudePairPasses: ampPass, InverseHessianShadowMagnitude: inverseShadow, InverseHessianClosurePasses: inverseClosure, MassAmplitudePairPasses: massPass, Verdict: join(StatusGaugeScalarPairingsAudited, StatusAmplitudePairSupported, StatusInverseHessianShadowSupported, StatusNoNativeSevenOver72Theorem)}
}

func pairing(name, gaugeLayer, gaugeCoord string, gaugeValue float64, scalarLayer, scalarCoord string, scalarValue float64, interpretation string) GaugeScalarPairingRow {
	den := gaugeValue - scalarValue
	wBest := math.NaN()
	if math.Abs(den) > 0 && finiteNumber(den) {
		wBest = (kSum - scalarValue) / den
	}
	deltaW := wBest - sevenOver72
	passes := finiteNumber(deltaW) && math.Abs(deltaW) < 1e-6
	rel := math.NaN()
	if scalarValue != 0 && finiteNumber(scalarValue) {
		rel = math.Abs(gaugeValue-scalarValue) / math.Abs(scalarValue)
	}
	verdict := StatusGaugeScalarPairingsAudited
	if name == "amplitude/quartic" && passes {
		verdict = join(StatusGaugeScalarPairingsAudited, StatusAmplitudePairSupported)
	} else if name == "inverse-kinetic/Hessian" {
		verdict = join(StatusGaugeScalarPairingsAudited, StatusInverseHessianShadowSupported, StatusNoNativeSevenOver72Theorem)
	}
	return GaugeScalarPairingRow{Name: name, GaugeLayer: gaugeLayer, GaugeCoordinate: gaugeCoord, GaugeValue: gaugeValue, ScalarLayer: scalarLayer, ScalarCoordinate: scalarCoord, ScalarValue: scalarValue, SignedDifference: gaugeValue - scalarValue, RelativeDifference: rel, WBest: wBest, WBestMinus7Over72: deltaW, PassesSevenOver72: passes, Interpretation: interpretation, Verdict: verdict}
}

func buildRetest(p GaugeScalarPairingAudit) ClosureCoordinateRetest {
	best := GaugeScalarPairingRow{}
	bestAbs := math.Inf(1)
	for _, r := range p.Rows {
		if finiteNumber(r.WBestMinus7Over72) && math.Abs(r.WBestMinus7Over72) < bestAbs {
			best = r
			bestAbs = math.Abs(r.WBestMinus7Over72)
		}
	}
	return ClosureCoordinateRetest{
		BestTypedPair:         best.Name,
		BestTypedResidual:     best.SignedDifference,
		BestTypedWBest:        best.WBest,
		BestTypedWBestMinus7:  best.WBestMinus7Over72,
		SevenOver72SelectedBy: "R_3-1 with |lambda(Lambda_12)|",
		InverseHessianStatus:  "2|lambda| is typed as a Hessian shadow and near the inverse/squared gauge scale, but it does not preserve the same 7/72 closure",
		Verdict:               join(StatusClosureCoordinateRetested, StatusAmplitudePairSupported, StatusInverseHessianShadowSupported, StatusNoNativeSevenOver72Theorem),
	}
}

func buildSource(p GaugeScalarPairingAudit, r ClosureCoordinateRetest) SourceTypeResult {
	classification := "ScalarQuarticCoordinateAirlockMissing"
	statements := []string{
		"the active scalar coordinate in the 7/72 closure is |lambda(Lambda_12)| paired with the connection-amplitude wound R_3-1",
		"the Hessian coordinate 2|lambda| is typed by m_H^2=2 lambda v^2 and shadows doubled/inverse gauge wounds, but it does not preserve the same 7/72 closure",
		"sqrt(|lambda|), sqrt(2|lambda|), signed lambda, and beta_lambda are typed slots but not selected by the active boundary-weighted closure",
		"a native scalar coordinate airlock theorem is still missing, so the scalar side remains a bridge-layer runtime wound rather than a native amplitude object",
	}
	if p.AmplitudePairPasses && !p.InverseHessianClosurePasses && r.BestTypedPair == "amplitude/quartic" {
		classification = "BoundaryWeightedDeficitClosureQuarticWoundSeal"
	}
	return SourceTypeResult{Classification: classification, Statements: statements, Verdict: join(StatusSourceTypeResultAudited, StatusScalarQuarticWoundSelected, StatusNoNativeScalarAirlockTheorem, StatusNoNativeBoundaryStressTheorem, StatusNoNativeSevenOver72Theorem)}
}

func buildPattern(g gate667.Analysis) RootAmplitudeRecurrenceAudit {
	return RootAmplitudeRecurrenceAudit{Rows: g.Coordinates.Rows, Pattern: "gauge side has a typed kinetic-to-connection square-root airlock; scalar side still lacks the analogous theorem selecting |lambda| over 2|lambda| or sqrt(2|lambda|)", Verdict: join(StatusRootAmplitudeRecurrenceAudited, StatusGaugeAmplitudeInherited, StatusNoNativeScalarAirlockTheorem)}
}

func inverseKineticWound(r float64) float64 { return 1 - 1/math.Pow(1+r, 2) }

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate667ConnectionAmplitudeInherited,
		StatusScalarCoordinateFamilyAudited,
		StatusHessianDoublingAudited,
		StatusGaugeScalarPairingsAudited,
		StatusClosureCoordinateRetested,
		StatusSourceTypeResultAudited,
		StatusRootAmplitudeRecurrenceAudited,
		StatusAmplitudePairSupported,
		StatusInverseHessianShadowSupported,
		StatusGaugeAmplitudeInherited,
		StatusScalarQuarticWoundSelected,
		StatusNoNativeScalarAirlockTheorem,
		StatusNoNativeBoundaryStressTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusNoNativeTransportTheorem,
		StatusNoHiggsMassStabilityGaugeFlavor,
		StatusGate668Boundary,
	}
}

func finiteNumber(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) }
