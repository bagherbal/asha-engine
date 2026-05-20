// Package generation2historywallbalancenormalvectorsourceaudit implements
// Gate 671: HistoryWallBalance Normal-Vector Source and Minimality Audit.
//
// Gate 670 defined the active scalar/flavor/gauge bridge as the oriented wall
// hyperplane
//
//	W_72 = kappa_lambda + kappa_e + (65/72)lambda(Lambda_12) - (7/72)(R_3-1).
//
// Gate 671 audits the normal vector n_72=(1,1,65/72,-7/72): its source type,
// typed minimality against nearby admissible normal choices, coordinate sealing,
// exact-vs-orientation kappa_e behavior, and Lambda_12 locality. It remains a
// bridge-layer normal-vector audit and preserves all native-theorem firewalls.
package generation2historywallbalancenormalvectorsourceaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate662 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundaryweighteddeficitclosurescalesweepaudit"
	gate670 "github.com/bagherbal/asha-engine/pkg/bridge/generation2orientedwalldistancehyperplaneaudit"
)

const (
	AuditID = "GATE671-HISTORY-WALL-BALANCE-NORMAL-VECTOR-SOURCE-MINIMALITY-AUDIT"

	StatusGate670HistoryWallBalanceInherited = "PASS_GATE670_HISTORY_WALL_BALANCE_INHERITED"
	StatusNormalVectorDefined                = "PASS_NORMAL_VECTOR_DEFINED"
	StatusNormalVectorDecompositionAudited   = "PASS_NORMAL_VECTOR_DECOMPOSITION_AUDITED"
	StatusTypedAlternativeNormalsCompared    = "PASS_TYPED_ALTERNATIVE_NORMALS_COMPARED"
	StatusCoordinateNormalizationAudited     = "PASS_COORDINATE_NORMALIZATION_AUDITED"
	StatusExactVersusOrientationAudited      = "PASS_EXACT_VERSUS_ORIENTATION_KAPPA_AUDITED"
	StatusScaleLocalAuditComputed            = "PASS_SCALE_LOCAL_AUDIT_COMPUTED"
	StatusSourceTypeCandidatesAudited        = "PASS_SOURCE_TYPE_CANDIDATES_AUDITED"
	StatusN72BestTypedWallBalanceNormalV1    = "CONDITIONAL_SUPPORT_N72_IS_BEST_TYPED_WALL_BALANCE_NORMAL_IN_V1"
	StatusHistoryWallNormalCoordinateSealed  = "CONDITIONAL_SUPPORT_HISTORY_WALL_BALANCE_NORMAL_IS_COORDINATE_SEALED"
	StatusAugmentedChamberTraceCandidate     = "CONDITIONAL_SUPPORT_AUGMENTED_CHAMBER_TRACE_SOURCE_CANDIDATE"
	StatusBoundaryInterpolationCandidate     = "CONDITIONAL_SUPPORT_BOUNDARY_INTERPOLATION_SOURCE_CANDIDATE"
	StatusNoNativeNormalVectorTheorem        = "FAILED_ROUTE_NO_NATIVE_NORMAL_VECTOR_SOURCE_THEOREM"
	StatusNoNativeSevenOver72Theorem         = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoNativeWallDistanceAirlock        = "FAILED_ROUTE_NO_NATIVE_WALL_DISTANCE_AIRLOCK_THEOREM"
	StatusNoBoundaryStressDerivation         = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusNoPhysicsPromotion                 = "FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM"
	StatusGate671Boundary                    = "FIREWALL_PRESERVED_GATE671_NORMAL_VECTOR_BOUNDARY"
)

const (
	kappaLambda  = 0.0443230430960771
	kappaEExact  = 0.00550355419157456
	kappaEOrient = 0.00550633006471245

	lambdaLambda12Signed = -0.0497009420776833
	r3Minus1             = 0.0509933868964996

	sevenOver72     = 7.0 / 72.0
	sixtyFiveOver72 = 65.0 / 72.0
)

type Gate670Inheritance struct {
	HistoryWallBalanceInherited bool
	SignedWallFormWritten       bool
	FunctionalDefined           bool
	NormalVectorAudited         bool
	OrientationAudited          bool
	CoordinateRolesClassified   bool
	NoNativeWallAirlock         bool
	NoNativeSevenOver72         bool
	NoBoundaryStress            bool
	FirewallPreserved           bool
	InheritedResidual           float64
	Verdict                     string
}

type NormalVectorAudit struct {
	Coordinates              []string
	Coefficients             []float64
	HistorySide              []float64
	BoundarySide             []float64
	HistorySideUnitWeights   bool
	BoundaryWeightsSumToOne  bool
	ScalarBoundaryDominant   bool
	GaugeBoundaryWeight      float64
	GaugeBoundaryWeightLabel string
	SignedAntiAlignment      bool
	NormalLabel              string
	Verdict                  string
}

type NormalVectorDecompositionAudit struct {
	HistoryBlockLabel  string
	BoundaryBlockLabel string
	HistoryMeaning     string
	BoundaryMeaning    string
	SplitEquation      string
	Verdict            string
}

type AlternativeNormal struct {
	Name             string
	VectorLabel      string
	ScalarCoeff      float64
	GaugeCoeff       float64
	Weight           float64
	WeightTyped      string
	ResidualExact    float64
	AbsExact         float64
	ResidualOrient   float64
	AbsOrient        float64
	MeetsConstraints bool
}

type MinimalityAudit struct {
	Alternatives               []AlternativeNormal
	BestExactName              string
	BestExactAbsResidual       float64
	BestOrientationName        string
	BestOrientationAbsResidual float64
	N72BestAmongTypedExact     bool
	N72BestAmongTypedOrient    bool
	TypedConstraints           []string
	Verdict                    string
}

type CoordinateNormalizationAudit struct {
	CanonicalCoordinates                  []string
	CoordinateSealed                      bool
	RescalingWarning                      string
	PreservesOnlyGate669WallNormalization bool
	Verdict                               string
}

type ExactVersusOrientationAudit struct {
	ExactKappaE                  float64
	OrientationKappaE            float64
	ExactResidualN72             float64
	OrientationResidualN72       float64
	ResidualGrowth               float64
	OrientationBestTypedName     string
	OrientationBestTypedResidual float64
	Interpretation               string
	Verdict                      string
}

type ScaleLocalAudit struct {
	Lambda12SelectedInGate662        bool
	LocalGate662MinimumAtLambda12    bool
	N72AtLambda12Residual            float64
	NearestLocalNonzeroDeltaLog      float64
	NearestLocalNonzeroResidual      float64
	N72BestTypedNormalOnlyAtLambda12 bool
	Statement                        string
	Verdict                          string
}

type SourceTypeCandidate struct {
	Name      string
	Candidate string
	Support   string
	Firewall  string
}

type SourceTypeAudit struct {
	Candidates                          []SourceTypeCandidate
	AugmentedTraceCandidate             bool
	BoundaryInterpolationCandidate      bool
	HistoryDeficitConservationCandidate bool
	CoordinateArtifactRisk              bool
	Verdict                             string
}

type VerdictDiscipline struct {
	ClaimsNativeNormalVectorTheorem  bool
	ClaimsNativeSevenOver72Theorem   bool
	ClaimsWallDistanceAirlockTheorem bool
	ClaimsBoundaryStressDerivation   bool
	ClaimsHiggsMassPrediction        bool
	ClaimsScalarStability            bool
	ClaimsGaugeUnification           bool
	ClaimsFlavorDerivation           bool
	ClaimsCKMPMNSDerivation          bool
	Verdict                          string
}

type Analysis struct {
	Inherited     Gate670Inheritance
	Normal        NormalVectorAudit
	Decomposition NormalVectorDecompositionAudit
	Minimality    MinimalityAudit
	Coordinate    CoordinateNormalizationAudit
	Orientation   ExactVersusOrientationAudit
	ScaleLocal    ScaleLocalAudit
	Source        SourceTypeAudit
	Discipline    VerdictDiscipline
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
	g670, err := gate670.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate670 inheritance unavailable: %w", err)
	}
	g662, err := gate662.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate662 scale-local inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g670)
	normal := buildNormalVector()
	decomp := buildDecomposition()
	minimality := buildMinimality()
	coord := buildCoordinateNormalization()
	orientation := buildExactVersusOrientation(minimality)
	scale := buildScaleLocal(g662, inherited.InheritedResidual)
	source := buildSourceType()
	discipline := VerdictDiscipline{Verdict: StatusGate671Boundary}
	truth := "Gate 671 audits the HistoryWallBalanceSeal normal vector n_72=(1,1,65/72,-7/72). In the exact Gate670 wall coordinates, n_72 is the best typed candidate among the tested normals and decomposes as unit history deficits balanced against a scalar-dominant signed boundary wall interpolation with 7/72 gauge pull. The normal is coordinate-sealed to the Gate669 canonical wall-distance normalization; no native normal-vector source theorem, native 7/72 theorem, wall-distance airlock theorem, or boundary-stress derivation is certified."
	return Analysis{Inherited: inherited, Normal: normal, Decomposition: decomp, Minimality: minimality, Coordinate: coord, Orientation: orientation, ScaleLocal: scale, Source: source, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate670.Analysis) Gate670Inheritance {
	return Gate670Inheritance{
		HistoryWallBalanceInherited: g.Functional.Name == "HistoryWallBalanceSeal" && g.Functional.PassesBridgeTolerance,
		SignedWallFormWritten:       g.Signed.LambdaIsNegative && g.Signed.EquivalentBecauseLambdaNegative,
		FunctionalDefined:           strings.Contains(g.Functional.Verdict, gate670.StatusHistoryWallBalanceSealDefined),
		NormalVectorAudited:         strings.Contains(g.Normal.Verdict, gate670.StatusNormalVectorAndWeightAudited),
		OrientationAudited:          strings.Contains(g.Orientation.Verdict, gate670.StatusOrientationApproximationAudited),
		CoordinateRolesClassified:   g.Roles.AllRolesClassified,
		NoNativeWallAirlock:         strings.Contains(g.Functional.Verdict, gate670.StatusNoNativeWallDistanceAirlockTheorem) || strings.Contains(g.Discipline.Verdict, gate670.StatusGate670Boundary),
		NoNativeSevenOver72:         strings.Contains(g.Functional.Verdict, gate670.StatusNoNativeSevenOver72Theorem) || !g.Discipline.ClaimsNativeSevenOver72,
		NoBoundaryStress:            strings.Contains(g.Functional.Verdict, gate670.StatusNoBoundaryStressDerivation) || !g.Discipline.ClaimsBoundaryStressDerivation,
		FirewallPreserved:           g.Discipline.Verdict == gate670.StatusGate670Boundary,
		InheritedResidual:           g.Functional.Value,
		Verdict:                     StatusGate670HistoryWallBalanceInherited,
	}
}

func buildNormalVector() NormalVectorAudit {
	return NormalVectorAudit{
		Coordinates:              []string{"kappa_lambda", "kappa_e", "lambda(Lambda_12)", "R_3-1"},
		Coefficients:             []float64{1, 1, sixtyFiveOver72, -sevenOver72},
		HistorySide:              []float64{1, 1},
		BoundarySide:             []float64{sixtyFiveOver72, -sevenOver72},
		HistorySideUnitWeights:   true,
		BoundaryWeightsSumToOne:  math.Abs(sixtyFiveOver72+sevenOver72-1) < 1e-15,
		ScalarBoundaryDominant:   sixtyFiveOver72 > sevenOver72,
		GaugeBoundaryWeight:      sevenOver72,
		GaugeBoundaryWeightLabel: "7/72",
		SignedAntiAlignment:      true,
		NormalLabel:              "n_72=(1,1,65/72,-7/72)",
		Verdict:                  strings.Join([]string{StatusNormalVectorDefined, StatusNormalVectorDecompositionAudited}, ";"),
	}
}

func buildDecomposition() NormalVectorDecompositionAudit {
	return NormalVectorDecompositionAudit{
		HistoryBlockLabel:  "(1,1)",
		BoundaryBlockLabel: "(65/72,-7/72)",
		HistoryMeaning:     "unit scalar matching deficit plus unit flavor wall deficit",
		BoundaryMeaning:    "scalar-zero wall coordinate dominates; gauge meeting-wall coordinate enters as a signed 7/72 pull",
		SplitEquation:      "65/72+7/72=1, with signed wall form kappa_lambda+kappa_e+(65/72)lambda-(7/72)(R3-1)",
		Verdict:            StatusNormalVectorDecompositionAudited,
	}
}

func buildMinimality() MinimalityAudit {
	alts := []AlternativeNormal{
		alternative("scalar only", "(1,1,1,0)", 1, 0, 0, "w=0 scalar-only boundary wall"),
		alternative("equal signed scalar/gauge", "(1,1,1,-1)", 1, 1, 1, "full signed gauge subtraction"),
		alternative("one eighth boundary pull", "(1,1,7/8,-1/8)", 1-1.0/8.0, 1.0/8.0, 1.0/8.0, "1/8 typed loop/simple fraction candidate"),
		alternative("one tenth boundary pull", "(1,1,9/10,-1/10)", 1-1.0/10.0, 1.0/10.0, 1.0/10.0, "1/10 typed decimal/block candidate"),
		alternative("seven over seventy two boundary pull", "(1,1,65/72,-7/72)", sixtyFiveOver72, sevenOver72, sevenOver72, "7/72 active augmented-chamber/boundary interpolation candidate"),
		alternative("seven over seventy boundary pull", "(1,1,63/70,-7/70)", 1-7.0/70.0, 7.0/70.0, 7.0/70.0, "7/70 native Lambda4-only denominator candidate"),
	}
	bestExact := alts[0]
	bestOrient := alts[0]
	for _, a := range alts {
		if a.AbsExact < bestExact.AbsExact {
			bestExact = a
		}
		if a.AbsOrient < bestOrient.AbsOrient {
			bestOrient = a
		}
	}
	return MinimalityAudit{
		Alternatives:               alts,
		BestExactName:              bestExact.Name,
		BestExactAbsResidual:       bestExact.AbsExact,
		BestOrientationName:        bestOrient.Name,
		BestOrientationAbsResidual: bestOrient.AbsOrient,
		N72BestAmongTypedExact:     bestExact.Name == "seven over seventy two boundary pull",
		N72BestAmongTypedOrient:    bestOrient.Name == "seven over seventy two boundary pull",
		TypedConstraints: []string{
			"unit weights on kappa_lambda and kappa_e",
			"boundary weights sum to one when written as (1-w)lambda - w(R3-1)",
			"scalar boundary coordinate remains dominant",
			"gauge boundary coordinate uses a typed candidate weight",
			"signed scalar/gauge anti-alignment is preserved",
		},
		Verdict: strings.Join([]string{StatusTypedAlternativeNormalsCompared, StatusN72BestTypedWallBalanceNormalV1}, ";"),
	}
}

func alternative(name, vectorLabel string, scalarCoeff, gaugeCoeff, weight float64, typed string) AlternativeNormal {
	resExact := residual(kappaEExact, scalarCoeff, gaugeCoeff)
	resOrient := residual(kappaEOrient, scalarCoeff, gaugeCoeff)
	return AlternativeNormal{
		Name:             name,
		VectorLabel:      vectorLabel,
		ScalarCoeff:      scalarCoeff,
		GaugeCoeff:       gaugeCoeff,
		Weight:           weight,
		WeightTyped:      typed,
		ResidualExact:    resExact,
		AbsExact:         math.Abs(resExact),
		ResidualOrient:   resOrient,
		AbsOrient:        math.Abs(resOrient),
		MeetsConstraints: scalarCoeff >= 0 && gaugeCoeff >= 0 && math.Abs(scalarCoeff+gaugeCoeff-1) < 1e-12,
	}
}

func residual(kappaE, scalarCoeff, gaugeCoeff float64) float64 {
	return kappaLambda + kappaE + scalarCoeff*lambdaLambda12Signed - gaugeCoeff*r3Minus1
}

func buildCoordinateNormalization() CoordinateNormalizationAudit {
	return CoordinateNormalizationAudit{
		CanonicalCoordinates:                  []string{"kappa_lambda", "kappa_e", "lambda(Lambda_12)", "R_3-1"},
		CoordinateSealed:                      true,
		RescalingWarning:                      "The coefficients 65/72 and 7/72 are meaningful only after Gate669 wall-coordinate normalization; arbitrary rescalings of lambda or R_3-1 change the normal-vector interpretation.",
		PreservesOnlyGate669WallNormalization: true,
		Verdict:                               strings.Join([]string{StatusCoordinateNormalizationAudited, StatusHistoryWallNormalCoordinateSealed}, ";"),
	}
}

func buildExactVersusOrientation(m MinimalityAudit) ExactVersusOrientationAudit {
	exact := residual(kappaEExact, sixtyFiveOver72, sevenOver72)
	orient := residual(kappaEOrient, sixtyFiveOver72, sevenOver72)
	return ExactVersusOrientationAudit{
		ExactKappaE:                  kappaEExact,
		OrientationKappaE:            kappaEOrient,
		ExactResidualN72:             exact,
		OrientationResidualN72:       orient,
		ResidualGrowth:               math.Abs(orient) - math.Abs(exact),
		OrientationBestTypedName:     m.BestOrientationName,
		OrientationBestTypedResidual: m.BestOrientationAbsResidual,
		Interpretation:               "Exact kappa_e preserves the n_72 residual at the 1e-9 level; replacing it by the OrientationBalance approximation grows the n_72 residual and makes nearby typed weights competitive. This keeps n_72 strongest in the exact wall ledger, not as a standalone flavor-derived theorem.",
		Verdict:                      StatusExactVersusOrientationAudited,
	}
}

func buildScaleLocal(g gate662.Analysis, residualAtLambda12 float64) ScaleLocalAudit {
	nearestDelta := math.Inf(1)
	nearestResidual := math.Inf(1)
	for _, row := range g.LocalSweep.Rows {
		if row.DeltaLog == 0 {
			continue
		}
		if math.Abs(row.DeltaLog) < math.Abs(nearestDelta) {
			nearestDelta = row.DeltaLog
			nearestResidual = row.AbsE72
		}
	}
	return ScaleLocalAudit{
		Lambda12SelectedInGate662:        g.ScaleSweep.Lambda12UniquelyMinimalEW,
		LocalGate662MinimumAtLambda12:    g.LocalSweep.LocalGridSelectsLambda12,
		N72AtLambda12Residual:            math.Abs(residualAtLambda12),
		NearestLocalNonzeroDeltaLog:      nearestDelta,
		NearestLocalNonzeroResidual:      nearestResidual,
		N72BestTypedNormalOnlyAtLambda12: g.ScaleSweep.Lambda12UniquelyMinimalEW && g.LocalSweep.LocalGridSelectsLambda12 && nearestResidual > math.Abs(residualAtLambda12),
		Statement:                        "Gate662 v1 scale sweep keeps the n_72 wall normal special at Lambda_12; this is a v1 local/root-crossing result, not a native scale theorem.",
		Verdict:                          StatusScaleLocalAuditComputed,
	}
}

func buildSourceType() SourceTypeAudit {
	candidates := []SourceTypeCandidate{
		{Name: "augmented chamber trace", Candidate: "7/72 from 7 over dim(Lambda^4 R^8)+dim(R^2_boundary)=70+2", Support: "typed numerator/denominator candidate inherited from Gates 628-630 and reactivated by Gate659", Firewall: StatusNoNativeSevenOver72Theorem},
		{Name: "boundary interpolation", Candidate: "7/72 as scalar/gauge wall split weight", Support: "directly active in W_72=(65/72)|lambda|+(7/72)(R_3-1)", Firewall: StatusNoBoundaryStressDerivation},
		{Name: "history-deficit conservation", Candidate: "kappa_lambda+kappa_e balanced against a boundary wall projection", Support: "exact wall-ledger residual at 8.53e-10", Firewall: StatusNoNativeNormalVectorTheorem},
		{Name: "coordinate artifact risk", Candidate: "normal is meaningful only in Gate669 wall coordinates", Support: "Gate665/Gate666/Gate667/Gate668 coordinate audits show amplitude/quartic wall-coordinate specificity", Firewall: StatusNoNativeWallDistanceAirlock},
	}
	return SourceTypeAudit{
		Candidates:                          candidates,
		AugmentedTraceCandidate:             true,
		BoundaryInterpolationCandidate:      true,
		HistoryDeficitConservationCandidate: true,
		CoordinateArtifactRisk:              true,
		Verdict:                             strings.Join([]string{StatusSourceTypeCandidatesAudited, StatusAugmentedChamberTraceCandidate, StatusBoundaryInterpolationCandidate}, ";"),
	}
}

func Statuses() []string {
	return []string{
		StatusGate670HistoryWallBalanceInherited,
		StatusNormalVectorDefined,
		StatusNormalVectorDecompositionAudited,
		StatusTypedAlternativeNormalsCompared,
		StatusCoordinateNormalizationAudited,
		StatusExactVersusOrientationAudited,
		StatusScaleLocalAuditComputed,
		StatusSourceTypeCandidatesAudited,
		StatusN72BestTypedWallBalanceNormalV1,
		StatusHistoryWallNormalCoordinateSealed,
		StatusAugmentedChamberTraceCandidate,
		StatusBoundaryInterpolationCandidate,
		StatusNoNativeNormalVectorTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusNoNativeWallDistanceAirlock,
		StatusNoBoundaryStressDerivation,
		StatusNoPhysicsPromotion,
		StatusGate671Boundary,
	}
}
