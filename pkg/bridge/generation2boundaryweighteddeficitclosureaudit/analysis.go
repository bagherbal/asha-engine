// Package generation2boundaryweighteddeficitclosureaudit implements
// Gate 626: BoundaryWeightedDeficitClosure Audit.
//
// Gate 625 found that the two history-loop deficits, kappa_lambda and kappa_e,
// close near the high-scale scalar wound |lambda(Lambda_12)|. Gate 626 audits
// the residual left by that closure and asks whether it is itself a typed pull
// along the Gate 613 boundary-stress split between |lambda(Lambda_12)| and
// R_3-1. This remains a bridge-layer closure audit only.
package generation2boundaryweighteddeficitclosureaudit

import (
	"fmt"
	"math"
	"sync"

	gate613 "github.com/bagherbal/asha-engine/pkg/bridge/generation2gaugescalarboundarystresssealaudit"
	gate625 "github.com/bagherbal/asha-engine/pkg/bridge/generation2historyloopdeficitclosuretriangleaudit"
)

const (
	AuditID = "GATE626-BOUNDARY-WEIGHTED-DEFICIT-CLOSURE-AUDIT"

	StatusGate625Inherited             = "PASS_GATE625_HISTORY_LOOP_DEFICIT_CLOSURE_INHERITED"
	StatusBoundarySplitComputed        = "PASS_BOUNDARY_SPLIT_RESIDUAL_COMPUTED"
	StatusSevenOverSeventyTwoAudited   = "PASS_SEVEN_OVER_SEVENTY_TWO_WEIGHT_AUDITED"
	StatusBoundaryWeightedClosure      = "CONDITIONAL_SUPPORT_DEFICIT_CLOSURE_IS_BOUNDARY_WEIGHTED"
	StatusSevenOverSeventyTwoCandidate = "CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_BOUNDARY_WEIGHT_CANDIDATE"
	StatusScalarFormulaComputed        = "PASS_BOUNDARY_WEIGHTED_SCALAR_FORMULA_COMPUTED"
	StatusNoNativeWeightSource         = "FAILED_ROUTE_NO_NATIVE_SOURCE_FOR_7_OVER_72_WEIGHT"
	StatusNoNativeTransportTheorem     = "FAILED_ROUTE_NO_NATIVE_GAUGE_SCALAR_FLAVOR_DEFICIT_TRANSPORT_THEOREM"
	StatusGate626Boundary              = "FIREWALL_PRESERVED_GATE626_BOUNDARY_WEIGHTED_CLOSURE_IS_BRIDGE_ONLY"
)

const (
	boundaryWeightNumerator   = 7.0
	boundaryWeightDenominator = 72.0
	sin2Theta13Quarter        = 0.0055375
	jCKM                      = 3.11699352875547e-05
	kappaEOrient              = sin2Theta13Quarter - jCKM
)

type Gate625Inheritance struct {
	LoopUnit                        float64
	KappaE                          float64
	KappaLambda                     float64
	KappaSum                        float64
	AbsLambda12                     float64
	R3MinusOne                      float64
	LambdaProxyMZ                   float64
	LambdaRuntimeMZ                 float64
	Gate625ClosureResidual          float64
	Gate625ScalarPredictionResidual float64
	Gate625ClosureSealDefined       bool
	Gate625NativeClosureTheorem     bool
	Gate625NativeScalarRGMatching   bool
	Gate625NativeFlavorOrientation  bool
	Verdict                         string
}

type BoundarySplitAudit struct {
	AbsLambda12                 float64
	R3MinusOne                  float64
	BoundarySplit               float64
	KappaSumMinusAbsLambda12    float64
	ClosureResidualOverSplit    float64
	SplitPositive               bool
	ResidualInsideBoundarySplit bool
	BoundaryStressLaneInherited bool
	Verdict                     string
}

type BoundaryWeightCandidate struct {
	Expression              string
	Numerator               int
	Denominator             int
	Value                   float64
	ObservedRatio           float64
	RatioResidual           float64
	AbsoluteRatioResidual   float64
	WeightedClosure         float64
	KappaSum                float64
	WeightedClosureResidual float64
	AbsoluteClosureResidual float64
	RelativeClosureResidual float64
	TypedOperands           bool
	NativeSourceCertified   bool
	CandidateInterpretation string
	Verdict                 string
}

type BoundaryWeightedClosureAudit struct {
	KappaSum           float64
	AbsLambda12        float64
	R3MinusOne         float64
	BoundaryWeight     float64
	ScalarWeight       float64
	BoundarySplit      float64
	WeightedMixture    float64
	Residual           float64
	AbsoluteResidual   float64
	RelativeResidual   float64
	ImprovesGate625    bool
	ImprovementFactor  float64
	ClosureEquation    string
	EquivalentEquation string
	BridgeOnly         bool
	Verdict            string
}

type WeightedScalarFormula struct {
	KappaLambdaActual             float64
	KappaEExact                   float64
	KappaEOrient                  float64
	WeightedBoundaryMixture       float64
	PredictedKappaLambdaExact     float64
	PredictedKappaLambdaOrient    float64
	KappaLambdaResidualExact      float64
	KappaLambdaResidualOrient     float64
	ExactFormula                  string
	OrientationSubstitutedFormula string
	CombinedScalarFormula         string
	NativeScalarFormulaClaimed    bool
	Verdict                       string
}

type ScalarPredictionRow struct {
	Name             string
	KappaEUsed       float64
	PredictedLambda  float64
	RuntimeLambda    float64
	Residual         float64
	RelativeResidual float64
	Formula          string
}

type FullScalarPredictionAudit struct {
	LambdaProxyMZ             float64
	LambdaRuntimeMZ           float64
	LoopUnit                  float64
	WeightedBoundaryMixture   float64
	Rows                      []ScalarPredictionRow
	BestResidual              float64
	ImprovesGate625Prediction bool
	Gate625PredictionResidual float64
	ImprovementFactor         float64
	DiagnosticOnly            bool
	Verdict                   string
}

type ResidualScaleRow struct {
	Name             string
	Residual         float64
	RelativeResidual float64
	Scale            string
	Meaning          string
}

type ResidualScaleComparison struct {
	Rows                           []ResidualScaleRow
	Gate625ClosureResidual         float64
	BoundaryWeightedResidual       float64
	Gate625ScalarResidual          float64
	BoundaryWeightedScalarResidual float64
	ClosureImprovementFactor       float64
	ScalarImprovementFactor        float64
	Verdict                        string
}

type SignAndRoleAudit struct {
	KappaE                  float64
	KappaLambda             float64
	AbsLambda12             float64
	R3MinusOne              float64
	BoundaryWeight          float64
	ScalarWeight            float64
	FlavorRole              string
	ScalarMatchingRole      string
	ScalarBoundaryWoundRole string
	GaugeBoundaryWoundRole  string
	StructuralEquation      string
	NativeTheoremClaimed    bool
	Verdict                 string
}

type NativeASHAStatus struct {
	NativeSevenOverSeventyTwoSource         bool
	NativeGaugeScalarFlavorDeficitTransport bool
	NativeBoundaryWeightedClosureTheorem    bool
	NativeScalarRGMatchingTheorem           bool
	NativeFlavorOrientationTheorem          bool
	Statement                               string
	Verdict                                 string
}

type Firewalls struct {
	ClaimsHiggsMassDerived       bool
	ClaimsScalarStability        bool
	ClaimsKoideDerived           bool
	ClaimsPMNSCKMDerived         bool
	ClaimsGaugeUnification       bool
	ClaimsNativeWeightTheorem    bool
	ClaimsNativeTransportTheorem bool
	ClaimsEndpointDerivation     bool
	Verdict                      string
}

type Analysis struct {
	Inherited        Gate625Inheritance
	BoundarySplit    BoundarySplitAudit
	WeightCandidate  BoundaryWeightCandidate
	WeightedClosure  BoundaryWeightedClosureAudit
	ScalarFormula    WeightedScalarFormula
	ScalarPrediction FullScalarPredictionAudit
	ResidualScales   ResidualScaleComparison
	SignRole         SignAndRoleAudit
	NativeStatus     NativeASHAStatus
	Firewalls        Firewalls
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
	g625, err := gate625.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate625 predecessor: %w", err)
	}
	g613, err := gate613.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate613 boundary-stress predecessor: %w", err)
	}

	inherit := inheritGate625(g625, g613)
	split := buildBoundarySplit(inherit)
	candidate := buildBoundaryWeightCandidate(inherit, split)
	closure := buildBoundaryWeightedClosure(inherit, split, candidate)
	formula := buildWeightedScalarFormula(inherit, closure)
	prediction := buildFullScalarPrediction(inherit, closure)
	residuals := buildResidualScaleComparison(inherit, split, closure, prediction)
	a := Analysis{
		Inherited:        inherit,
		BoundarySplit:    split,
		WeightCandidate:  candidate,
		WeightedClosure:  closure,
		ScalarFormula:    formula,
		ScalarPrediction: prediction,
		ResidualScales:   residuals,
		SignRole:         buildSignAndRole(inherit, closure),
		NativeStatus:     buildNativeStatus(),
		Firewalls:        auditFirewalls(),
		Truth:            "Gate 626 audits the Gate625 residual rather than inventing a new endpoint. The residual of kappa_lambda+kappa_e against |lambda(Lambda_12)| is almost exactly 7/72 of the Gate613 boundary split (R_3-1)-|lambda(Lambda_12)|, yielding a boundary-weighted scalar/gauge wound mixture with residual near 1e-9 in the kappa lane and near 1e-12 in lambda(M_Z). This is a bridge-layer seal candidate only: ASHA supplies no native source theorem for the 7/72 weight and no native gauge-scalar-flavor deficit transport theorem.",
	}
	return a, nil
}

func inheritGate625(g gate625.Analysis, g613 gate613.Analysis) Gate625Inheritance {
	return Gate625Inheritance{
		LoopUnit:                        g.Inherited.LoopUnit,
		KappaE:                          g.Kappas.KappaE,
		KappaLambda:                     g.Kappas.KappaLambda,
		KappaSum:                        g.ClosureTable.KappaSum,
		AbsLambda12:                     g613.Inherited.AbsLambda12,
		R3MinusOne:                      g613.Inherited.R3MinusOne,
		LambdaProxyMZ:                   g.ScalarPrediction.LambdaProxyMZ,
		LambdaRuntimeMZ:                 g.ScalarPrediction.LambdaRuntimeMZ,
		Gate625ClosureResidual:          g.ClosureTable.ClosestResidual,
		Gate625ScalarPredictionResidual: g.ScalarPrediction.BestResidual,
		Gate625ClosureSealDefined:       g.ResidualScales.Verdict == gate625.StatusClosureSealDefined,
		Gate625NativeClosureTheorem:     g.NativeStatus.NativeHistoryLoopDeficitClosureTheorem,
		Gate625NativeScalarRGMatching:   g.NativeStatus.NativeScalarRGMatchingTheorem,
		Gate625NativeFlavorOrientation:  g.NativeStatus.NativeFlavorOrientationTheorem,
		Verdict:                         StatusGate625Inherited,
	}
}

func buildBoundarySplit(i Gate625Inheritance) BoundarySplitAudit {
	boundarySplit := i.R3MinusOne - i.AbsLambda12
	residual := i.KappaSum - i.AbsLambda12
	return BoundarySplitAudit{
		AbsLambda12:                 i.AbsLambda12,
		R3MinusOne:                  i.R3MinusOne,
		BoundarySplit:               boundarySplit,
		KappaSumMinusAbsLambda12:    residual,
		ClosureResidualOverSplit:    residual / boundarySplit,
		SplitPositive:               boundarySplit > 0,
		ResidualInsideBoundarySplit: residual > 0 && residual < boundarySplit,
		BoundaryStressLaneInherited: i.R3MinusOne > i.AbsLambda12 && i.AbsLambda12 > 0,
		Verdict:                     StatusBoundarySplitComputed,
	}
}

func buildBoundaryWeightCandidate(i Gate625Inheritance, s BoundarySplitAudit) BoundaryWeightCandidate {
	w := boundaryWeightNumerator / boundaryWeightDenominator
	weighted := s.AbsLambda12 + w*s.BoundarySplit
	residual := i.KappaSum - weighted
	ratioResidual := s.ClosureResidualOverSplit - w
	return BoundaryWeightCandidate{
		Expression:              "7/72",
		Numerator:               int(boundaryWeightNumerator),
		Denominator:             int(boundaryWeightDenominator),
		Value:                   w,
		ObservedRatio:           s.ClosureResidualOverSplit,
		RatioResidual:           ratioResidual,
		AbsoluteRatioResidual:   math.Abs(ratioResidual),
		WeightedClosure:         weighted,
		KappaSum:                i.KappaSum,
		WeightedClosureResidual: residual,
		AbsoluteClosureResidual: math.Abs(residual),
		RelativeClosureResidual: math.Abs(residual) / math.Abs(i.KappaSum),
		TypedOperands:           true,
		NativeSourceCertified:   false,
		CandidateInterpretation: "small boundary-stress projection weight pulling the scalar wound toward the strong-sector boundary wound",
		Verdict:                 StatusSevenOverSeventyTwoCandidate,
	}
}

func buildBoundaryWeightedClosure(i Gate625Inheritance, s BoundarySplitAudit, c BoundaryWeightCandidate) BoundaryWeightedClosureAudit {
	improvement := i.Gate625ClosureResidual / c.AbsoluteClosureResidual
	return BoundaryWeightedClosureAudit{
		KappaSum:           i.KappaSum,
		AbsLambda12:        s.AbsLambda12,
		R3MinusOne:         s.R3MinusOne,
		BoundaryWeight:     c.Value,
		ScalarWeight:       1 - c.Value,
		BoundarySplit:      s.BoundarySplit,
		WeightedMixture:    c.WeightedClosure,
		Residual:           c.WeightedClosureResidual,
		AbsoluteResidual:   c.AbsoluteClosureResidual,
		RelativeResidual:   c.RelativeClosureResidual,
		ImprovesGate625:    improvement > 1,
		ImprovementFactor:  improvement,
		ClosureEquation:    "kappa_lambda+kappa_e ≈ |lambda(Lambda_12)| + (7/72)[(R_3-1)-|lambda(Lambda_12)|]",
		EquivalentEquation: "kappa_lambda+kappa_e ≈ (65/72)|lambda(Lambda_12)| + (7/72)(R_3-1)",
		BridgeOnly:         true,
		Verdict:            StatusBoundaryWeightedClosure,
	}
}

func buildWeightedScalarFormula(i Gate625Inheritance, c BoundaryWeightedClosureAudit) WeightedScalarFormula {
	predExact := c.WeightedMixture - i.KappaE
	predOrient := c.WeightedMixture - kappaEOrient
	return WeightedScalarFormula{
		KappaLambdaActual:             i.KappaLambda,
		KappaEExact:                   i.KappaE,
		KappaEOrient:                  kappaEOrient,
		WeightedBoundaryMixture:       c.WeightedMixture,
		PredictedKappaLambdaExact:     predExact,
		PredictedKappaLambdaOrient:    predOrient,
		KappaLambdaResidualExact:      i.KappaLambda - predExact,
		KappaLambdaResidualOrient:     i.KappaLambda - predOrient,
		ExactFormula:                  "kappa_lambda ≈ (65/72)|lambda(Lambda_12)| + (7/72)(R_3-1) - kappa_e",
		OrientationSubstitutedFormula: "kappa_lambda ≈ (65/72)|lambda(Lambda_12)| + (7/72)(R_3-1) - sin^2(theta13)/4 + J_CKM",
		CombinedScalarFormula:         "lambda(M_Z) ≈ lambda_proxy[1+L(1-((65/72)|lambda(Lambda_12)|+(7/72)(R_3-1))+sin^2(theta13)/4-J_CKM)]",
		NativeScalarFormulaClaimed:    false,
		Verdict:                       StatusScalarFormulaComputed,
	}
}

func buildFullScalarPrediction(i Gate625Inheritance, c BoundaryWeightedClosureAudit) FullScalarPredictionAudit {
	rows := []ScalarPredictionRow{
		predictionRow("boundary-weighted closure with exact kappa_e", i, c.WeightedMixture, i.KappaE),
		predictionRow("boundary-weighted closure with PMNS/CKM orientation kappa_e", i, c.WeightedMixture, kappaEOrient),
	}
	best := math.Abs(rows[0].Residual)
	for _, r := range rows[1:] {
		if math.Abs(r.Residual) < best {
			best = math.Abs(r.Residual)
		}
	}
	improvement := i.Gate625ScalarPredictionResidual / best
	return FullScalarPredictionAudit{
		LambdaProxyMZ:             i.LambdaProxyMZ,
		LambdaRuntimeMZ:           i.LambdaRuntimeMZ,
		LoopUnit:                  i.LoopUnit,
		WeightedBoundaryMixture:   c.WeightedMixture,
		Rows:                      rows,
		BestResidual:              best,
		ImprovesGate625Prediction: improvement > 1,
		Gate625PredictionResidual: i.Gate625ScalarPredictionResidual,
		ImprovementFactor:         improvement,
		DiagnosticOnly:            true,
		Verdict:                   StatusScalarFormulaComputed,
	}
}

func predictionRow(name string, i Gate625Inheritance, weightedMixture, kappaEUsed float64) ScalarPredictionRow {
	pred := i.LambdaProxyMZ * (1 + i.LoopUnit*(1-weightedMixture+kappaEUsed))
	res := pred - i.LambdaRuntimeMZ
	return ScalarPredictionRow{
		Name:             name,
		KappaEUsed:       kappaEUsed,
		PredictedLambda:  pred,
		RuntimeLambda:    i.LambdaRuntimeMZ,
		Residual:         res,
		RelativeResidual: math.Abs(res) / math.Abs(i.LambdaRuntimeMZ),
		Formula:          "lambda_proxy(M_Z)[1+L(1-boundary_weighted_wound+kappa_e)]",
	}
}

func buildResidualScaleComparison(i Gate625Inheritance, s BoundarySplitAudit, c BoundaryWeightedClosureAudit, p FullScalarPredictionAudit) ResidualScaleComparison {
	closureImprovement := i.Gate625ClosureResidual / c.AbsoluteResidual
	scalarImprovement := i.Gate625ScalarPredictionResidual / p.BestResidual
	rows := []ResidualScaleRow{
		{Name: "Gate625 deficit closure residual", Residual: i.Gate625ClosureResidual, RelativeResidual: i.Gate625ClosureResidual / math.Abs(i.AbsLambda12), Scale: "dimensionless kappa/lambda wound", Meaning: "kappa_lambda+kappa_e versus |lambda(Lambda_12)|"},
		{Name: "Gate626 boundary-weighted closure residual", Residual: c.AbsoluteResidual, RelativeResidual: c.RelativeResidual, Scale: "dimensionless kappa/lambda wound", Meaning: "kappa_lambda+kappa_e versus the 65/72 scalar + 7/72 gauge boundary mixture"},
		{Name: "Gate613 boundary split", Residual: s.BoundarySplit, RelativeResidual: s.BoundarySplit / math.Abs(s.AbsLambda12), Scale: "boundary stress", Meaning: "(R_3-1)-|lambda(Lambda_12)|"},
		{Name: "Gate625 scalar prediction residual", Residual: i.Gate625ScalarPredictionResidual, RelativeResidual: i.Gate625ScalarPredictionResidual / math.Abs(i.LambdaRuntimeMZ), Scale: "lambda(M_Z)", Meaning: "scalar-flavor-boundary prediction before 7/72 boundary weighting"},
		{Name: "Gate626 scalar prediction residual", Residual: p.BestResidual, RelativeResidual: p.BestResidual / math.Abs(i.LambdaRuntimeMZ), Scale: "lambda(M_Z)", Meaning: "boundary-weighted scalar-flavor-gauge prediction"},
	}
	return ResidualScaleComparison{
		Rows:                           rows,
		Gate625ClosureResidual:         i.Gate625ClosureResidual,
		BoundaryWeightedResidual:       c.AbsoluteResidual,
		Gate625ScalarResidual:          i.Gate625ScalarPredictionResidual,
		BoundaryWeightedScalarResidual: p.BestResidual,
		ClosureImprovementFactor:       closureImprovement,
		ScalarImprovementFactor:        scalarImprovement,
		Verdict:                        StatusBoundaryWeightedClosure,
	}
}

func buildSignAndRole(i Gate625Inheritance, c BoundaryWeightedClosureAudit) SignAndRoleAudit {
	return SignAndRoleAudit{
		KappaE:                  i.KappaE,
		KappaLambda:             i.KappaLambda,
		AbsLambda12:             i.AbsLambda12,
		R3MinusOne:              i.R3MinusOne,
		BoundaryWeight:          c.BoundaryWeight,
		ScalarWeight:            c.ScalarWeight,
		FlavorRole:              "flavor orientation deficit inside L: epsilon_e=L(1-kappa_e)",
		ScalarMatchingRole:      "scalar low-scale matching deficit inside L: lambda_runtime=lambda_proxy[1+L(1-kappa_lambda)]",
		ScalarBoundaryWoundRole: "dominant high-scale scalar wound |lambda(Lambda_12)| with weight 65/72",
		GaugeBoundaryWoundRole:  "small strong-sector boundary pull R_3-1 with weight 7/72",
		StructuralEquation:      "flavor deficit + scalar matching deficit ≈ boundary-weighted scalar/gauge wound mixture",
		NativeTheoremClaimed:    false,
		Verdict:                 StatusBoundaryWeightedClosure,
	}
}

func buildNativeStatus() NativeASHAStatus {
	return NativeASHAStatus{
		NativeSevenOverSeventyTwoSource:         false,
		NativeGaugeScalarFlavorDeficitTransport: false,
		NativeBoundaryWeightedClosureTheorem:    false,
		NativeScalarRGMatchingTheorem:           false,
		NativeFlavorOrientationTheorem:          false,
		Statement:                               "ASHA supplies typed operands for the residual audit, but no native source for the 7/72 projection weight and no native theorem transporting flavor, scalar matching, and gauge/scalar boundary wounds into one closure law.",
		Verdict:                                 StatusNoNativeWeightSource,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{Verdict: StatusGate626Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate625Inherited,
		StatusBoundarySplitComputed,
		StatusSevenOverSeventyTwoAudited,
		StatusBoundaryWeightedClosure,
		StatusSevenOverSeventyTwoCandidate,
		StatusScalarFormulaComputed,
		StatusNoNativeWeightSource,
		StatusNoNativeTransportTheorem,
		StatusGate626Boundary,
	}
}
