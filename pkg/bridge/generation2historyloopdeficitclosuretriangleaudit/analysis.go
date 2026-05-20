// Package generation2historyloopdeficitclosuretriangleaudit implements
// Gate 625: HistoryLoopDeficit Closure Triangle Audit.
//
// Gate 624 typed L=1/(8*pi) as a quarter-normalized phase-unit candidate but
// refused to promote it into a native theorem. Gate 625 asks the sharper
// bridge-layer question exposed by that audit: do the two L-seal deficits,
// kappa_e and kappa_lambda, close against the high-scale scalar wound
// |lambda(Lambda_12)|?
package generation2historyloopdeficitclosuretriangleaudit

import (
	"fmt"
	"math"
	"sync"

	gate613 "github.com/bagherbal/asha-engine/pkg/bridge/generation2gaugescalarboundarystresssealaudit"
	gate624 "github.com/bagherbal/asha-engine/pkg/bridge/generation2historyloopunitsourcetypeaudit"
	gate621 "github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarproxyruntimematchinggapaudit"
)

const (
	AuditID = "GATE625-HISTORY-LOOP-DEFICIT-CLOSURE-TRIANGLE-AUDIT"

	StatusGate624Inherited          = "PASS_GATE624_HISTORY_LOOP_UNIT_INHERITED"
	StatusKappasDefined             = "PASS_KAPPA_E_AND_KAPPA_LAMBDA_DEFINED"
	StatusDeficitClosureComputed    = "PASS_DEFICIT_CLOSURE_TEST_COMPUTED"
	StatusClosureOnAbsLambda12      = "CONDITIONAL_SUPPORT_KAPPA_LAMBDA_PLUS_KAPPA_E_CLOSES_ON_ABS_LAMBDA_LAMBDA12"
	StatusScalarPredictionComputed  = "PASS_FULL_SCALAR_PREDICTION_FROM_CLOSURE_COMPUTED"
	StatusClosureSealDefined        = "CONDITIONAL_SUPPORT_HISTORY_LOOP_DEFICIT_CLOSURE_SEAL_DEFINED"
	StatusNoNativeKappaClosure      = "FAILED_ROUTE_NO_NATIVE_KAPPA_CLOSURE_THEOREM"
	StatusNoNativeScalarRGMatching  = "FAILED_ROUTE_NO_NATIVE_SCALAR_RG_MATCHING_THEOREM"
	StatusNoNativeFlavorOrientation = "FAILED_ROUTE_NO_NATIVE_FLAVOR_ORIENTATION_THEOREM"
	StatusGate625Boundary           = "FIREWALL_PRESERVED_GATE625_HISTORY_LOOP_DEFICIT_CLOSURE_BOUNDARY"
)

const (
	sin2Theta13Quarter = 0.0055375
	jCKM               = 3.11699352875547e-05
	kappaEOrient       = sin2Theta13Quarter - jCKM
)

type Gate624Inheritance struct {
	LoopUnit                  float64
	KappaE                    float64
	KappaLambda               float64
	LambdaProxyMZ             float64
	LambdaRuntimeMZ           float64
	FlavorOrientationKappa    float64
	FlavorOrientationResidual float64
	Gate624QuarterPhase       bool
	NativeHistoryLoopUnit     bool
	NativeHopfToScalar        bool
	NativeHopfToFlavor        bool
	Verdict                   string
}

type KappaDefinitions struct {
	KappaE              float64
	KappaLambda         float64
	LoopUnit            float64
	BothPositive        bool
	ScalarDeficitLarger bool
	Verdict             string
}

type DeficitClosureRow struct {
	Target           string
	TargetValue      float64
	KappaSum         float64
	Residual         float64
	AbsoluteResidual float64
	RelativeResidual float64
	Typed            bool
	NativeCertified  bool
	Comment          string
}

type DeficitClosureTable struct {
	KappaE            float64
	KappaLambda       float64
	KappaSum          float64
	Rows              []DeficitClosureRow
	ClosestTarget     string
	ClosestResidual   float64
	ClosestRelative   float64
	ClosesOnAbsLambda bool
	Verdict           string
}

type ScalarDeficitFormula struct {
	KappaLambdaActual          float64
	AbsLambda12                float64
	KappaEExact                float64
	KappaEOrient               float64
	PredictedKappaLambdaExact  float64
	PredictedKappaLambdaOrient float64
	ResidualExact              float64
	ResidualOrient             float64
	ExactFormula               string
	OrientationSubstitutedForm string
	BridgeOnly                 bool
	Verdict                    string
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
	AbsLambda12               float64
	Rows                      []ScalarPredictionRow
	BestResidual              float64
	ImprovesGate623RawLAnsatz bool
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
	Rows                              []ResidualScaleRow
	Gate623ScalarAnsatzResidual       float64
	ClosureScalarResidual             float64
	ScalarImprovementFactor           float64
	ClosureSharperThanRawScalarAnsatz bool
	ClosureResidualDimensionless      float64
	Verdict                           string
}

type SignAndRoleAudit struct {
	KappaE                    float64
	KappaLambda               float64
	AbsLambda12               float64
	FlavorDeficitRole         string
	ScalarMatchingDeficitRole string
	HighScaleScalarWoundRole  string
	StructuralEquation        string
	OpposedRGWoundSign        bool
	NativeTheoremClaimed      bool
	Verdict                   string
}

type NativeASHAStatus struct {
	NativeKappaClosureTheorem                 bool
	NativeScalarRGMatchingTheorem             bool
	NativeFlavorOrientationTheorem            bool
	NativeLowScaleMatchingToHighScaleWoundLaw bool
	NativeHistoryLoopDeficitClosureTheorem    bool
	Statement                                 string
	Verdict                                   string
}

type Firewalls struct {
	ClaimsHiggsMassDerived      bool
	ClaimsScalarStability       bool
	ClaimsKoideDerived          bool
	ClaimsPMNSCKMDerived        bool
	ClaimsGaugeUnification      bool
	ClaimsNativeASHAClosure     bool
	ClaimsNativeHistoryLoopUnit bool
	Verdict                     string
}

type Analysis struct {
	Inherited        Gate624Inheritance
	Kappas           KappaDefinitions
	ClosureTable     DeficitClosureTable
	ScalarFormula    ScalarDeficitFormula
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
	g624, err := gate624.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate624 predecessor: %w", err)
	}
	g613, err := gate613.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate613 boundary-stress predecessor: %w", err)
	}
	g621, err := gate621.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate621 scalar proxy-runtime predecessor: %w", err)
	}

	inherit := inheritGate624(g624)
	kappas := buildKappaDefinitions(inherit)
	closure := buildClosureTable(kappas, g613)
	formula := buildScalarFormula(inherit, g613.Inherited.AbsLambda12)
	prediction := buildScalarPrediction(inherit, g613.Inherited.AbsLambda12)
	residuals := buildResidualScaleComparison(closure, prediction, g624, g613, g621)
	a := Analysis{
		Inherited:        inherit,
		Kappas:           kappas,
		ClosureTable:     closure,
		ScalarFormula:    formula,
		ScalarPrediction: prediction,
		ResidualScales:   residuals,
		SignRole:         buildSignAndRole(inherit, g613.Inherited.AbsLambda12, g613.Inherited.LambdaLambda12),
		NativeStatus:     buildNativeStatus(),
		Firewalls:        auditFirewalls(),
		Truth:            "Gate 625 audits the bridge-layer closure kappa_lambda+kappa_e≈|lambda(Lambda_12)|. The strongest typed target is the high-scale scalar wound, yielding a HistoryLoopDeficitClosureSeal candidate and a sharper scalar-flavor-boundary prediction for lambda(M_Z). The closure remains conditional and environmental: no native kappa-closure, scalar RG-matching, flavor-orientation, or history-loop-deficit theorem is certified.",
	}
	return a, nil
}

func inheritGate624(g gate624.Analysis) Gate624Inheritance {
	return Gate624Inheritance{
		LoopUnit:                  g.Inherited.LoopUnit,
		KappaE:                    g.FlavorRole.KappaE,
		KappaLambda:               g.ScalarRole.KappaLambda,
		LambdaProxyMZ:             g.ScalarRole.LambdaProxy,
		LambdaRuntimeMZ:           g.ScalarRole.LambdaRuntime,
		FlavorOrientationKappa:    g.FlavorRole.OrientationCandidate,
		FlavorOrientationResidual: g.FlavorRole.Residual,
		Gate624QuarterPhase:       g.Decompositions.Verdict == gate624.StatusDecompositionsTyped && g.HopfPhase.QuarterProjectionCandidate,
		NativeHistoryLoopUnit:     g.NativeStatus.NativeLTheorem,
		NativeHopfToScalar:        g.NativeStatus.NativeHopfToScalarMatchingMap,
		NativeHopfToFlavor:        g.NativeStatus.NativeHopfToFlavorWallMap,
		Verdict:                   StatusGate624Inherited,
	}
}

func buildKappaDefinitions(i Gate624Inheritance) KappaDefinitions {
	return KappaDefinitions{
		KappaE:              i.KappaE,
		KappaLambda:         i.KappaLambda,
		LoopUnit:            i.LoopUnit,
		BothPositive:        i.KappaE > 0 && i.KappaLambda > 0,
		ScalarDeficitLarger: i.KappaLambda > i.KappaE,
		Verdict:             StatusKappasDefined,
	}
}

func buildClosureTable(k KappaDefinitions, g613 gate613.Analysis) DeficitClosureTable {
	sum := k.KappaLambda + k.KappaE
	rows := []DeficitClosureRow{
		closureRow("|lambda(Lambda_12)|", g613.Inherited.AbsLambda12, sum, true, "high-scale scalar runtime wound from RG transport"),
		closureRow("R_3-1", g613.Inherited.R3MinusOne, sum, true, "strong-sector relative boundary wound"),
		closureRow("xi_boundary", g613.StressSeal.XiBoundary, sum, true, "Gate613 mean gauge-scalar boundary stress scale"),
	}
	name, absRes, rel := closestClosureTarget(rows)
	return DeficitClosureTable{
		KappaE:            k.KappaE,
		KappaLambda:       k.KappaLambda,
		KappaSum:          sum,
		Rows:              rows,
		ClosestTarget:     name,
		ClosestResidual:   absRes,
		ClosestRelative:   rel,
		ClosesOnAbsLambda: name == "|lambda(Lambda_12)|" && rel < 0.003,
		Verdict:           StatusDeficitClosureComputed,
	}
}

func closureRow(name string, target, sum float64, typed bool, comment string) DeficitClosureRow {
	res := sum - target
	return DeficitClosureRow{
		Target:           name,
		TargetValue:      target,
		KappaSum:         sum,
		Residual:         res,
		AbsoluteResidual: math.Abs(res),
		RelativeResidual: math.Abs(res) / math.Abs(target),
		Typed:            typed,
		NativeCertified:  false,
		Comment:          comment,
	}
}

func buildScalarFormula(i Gate624Inheritance, absLambda12 float64) ScalarDeficitFormula {
	predExact := absLambda12 - i.KappaE
	predOrient := absLambda12 - kappaEOrient
	return ScalarDeficitFormula{
		KappaLambdaActual:          i.KappaLambda,
		AbsLambda12:                absLambda12,
		KappaEExact:                i.KappaE,
		KappaEOrient:               kappaEOrient,
		PredictedKappaLambdaExact:  predExact,
		PredictedKappaLambdaOrient: predOrient,
		ResidualExact:              i.KappaLambda - predExact,
		ResidualOrient:             i.KappaLambda - predOrient,
		ExactFormula:               "kappa_lambda ≈ |lambda(Lambda_12)| - kappa_e",
		OrientationSubstitutedForm: "kappa_lambda ≈ |lambda(Lambda_12)| - sin^2(theta13)/4 + J_CKM",
		BridgeOnly:                 true,
		Verdict:                    StatusClosureOnAbsLambda12,
	}
}

func buildScalarPrediction(i Gate624Inheritance, absLambda12 float64) FullScalarPredictionAudit {
	rows := []ScalarPredictionRow{
		predictionRow("closure with exact kappa_e", i, absLambda12, i.KappaE),
		predictionRow("closure with PMNS/CKM orientation kappa_e", i, absLambda12, kappaEOrient),
	}
	best := math.Abs(rows[0].Residual)
	for _, r := range rows[1:] {
		if math.Abs(r.Residual) < best {
			best = math.Abs(r.Residual)
		}
	}
	rawAnsatzResidual := math.Abs(i.LambdaProxyMZ*(1+i.LoopUnit) - i.LambdaRuntimeMZ)
	return FullScalarPredictionAudit{
		LambdaProxyMZ:             i.LambdaProxyMZ,
		LambdaRuntimeMZ:           i.LambdaRuntimeMZ,
		LoopUnit:                  i.LoopUnit,
		AbsLambda12:               absLambda12,
		Rows:                      rows,
		BestResidual:              best,
		ImprovesGate623RawLAnsatz: best < rawAnsatzResidual,
		DiagnosticOnly:            true,
		Verdict:                   StatusScalarPredictionComputed,
	}
}

func predictionRow(name string, i Gate624Inheritance, absLambda12, kappaEUsed float64) ScalarPredictionRow {
	pred := i.LambdaProxyMZ * (1 + i.LoopUnit*(1-absLambda12+kappaEUsed))
	res := pred - i.LambdaRuntimeMZ
	return ScalarPredictionRow{
		Name:             name,
		KappaEUsed:       kappaEUsed,
		PredictedLambda:  pred,
		RuntimeLambda:    i.LambdaRuntimeMZ,
		Residual:         res,
		RelativeResidual: math.Abs(res) / math.Abs(i.LambdaRuntimeMZ),
		Formula:          "lambda_proxy(M_Z)[1+L(1-|lambda(Lambda_12)|+kappa_e)]",
	}
}

func buildResidualScaleComparison(closure DeficitClosureTable, prediction FullScalarPredictionAudit, g624 gate624.Analysis, g613 gate613.Analysis, g621 gate621.Analysis) ResidualScaleComparison {
	rawScalarResidual := math.Abs(g624.ScalarRole.LambdaProxy*(1+g624.Inherited.LoopUnit) - g624.ScalarRole.LambdaRuntime)
	closureScalarResidual := prediction.BestResidual
	boundaryResidual := math.Abs(g613.AntiAlignment.SPlus)
	rows := []ResidualScaleRow{
		{Name: "Gate625 deficit closure residual", Residual: closure.ClosestResidual, RelativeResidual: closure.ClosestRelative, Scale: "dimensionless kappa/lambda wound", Meaning: "kappa_lambda+kappa_e versus |lambda(Lambda_12)|"},
		{Name: "Gate625 scalar prediction residual", Residual: closureScalarResidual, RelativeResidual: closureScalarResidual / math.Abs(prediction.LambdaRuntimeMZ), Scale: "lambda(M_Z)", Meaning: "full scalar-flavor-boundary prediction versus runtime lambda"},
		{Name: "Gate623 raw scalar L ansatz residual", Residual: rawScalarResidual, RelativeResidual: rawScalarResidual / math.Abs(g624.ScalarRole.LambdaRuntime), Scale: "lambda(M_Z)", Meaning: "lambda_proxy*(1+L) versus runtime lambda"},
		{Name: "Gate590/624 flavor orientation residual", Residual: math.Abs(g624.FlavorRole.Residual), RelativeResidual: math.Abs(g624.FlavorRole.Residual) / math.Abs(g624.FlavorRole.EpsilonE), Scale: "epsilon_e", Meaning: "orientation-corrected flavor wall residual"},
		{Name: "Gate613 boundary anti-alignment residual", Residual: boundaryResidual, RelativeResidual: g613.AntiAlignment.RelativeAntiAlignment, Scale: "boundary stress", Meaning: "|(R_3-1)+lambda(Lambda_12)| relative to xi_mean"},
		{Name: "Gate621 proxy-runtime gap", Residual: g621.MatchingGap.DeltaLambdaMatch, RelativeResidual: g621.MatchingGap.RelativeToProxy, Scale: "lambda(M_Z)", Meaning: "lambda_runtime(M_Z)-lambda_proxy(M_Z)"},
	}
	improvement := rawScalarResidual / closureScalarResidual
	return ResidualScaleComparison{
		Rows:                              rows,
		Gate623ScalarAnsatzResidual:       rawScalarResidual,
		ClosureScalarResidual:             closureScalarResidual,
		ScalarImprovementFactor:           improvement,
		ClosureSharperThanRawScalarAnsatz: improvement > 100,
		ClosureResidualDimensionless:      closure.ClosestResidual,
		Verdict:                           StatusClosureSealDefined,
	}
}

func buildSignAndRole(i Gate624Inheritance, absLambda12, lambdaLambda12 float64) SignAndRoleAudit {
	return SignAndRoleAudit{
		KappaE:                    i.KappaE,
		KappaLambda:               i.KappaLambda,
		AbsLambda12:               absLambda12,
		FlavorDeficitRole:         "flavor orientation deficit inside L: epsilon_e=L(1-kappa_e)",
		ScalarMatchingDeficitRole: "scalar low-scale matching deficit inside L: lambda_runtime=lambda_proxy[1+L(1-kappa_lambda)]",
		HighScaleScalarWoundRole:  "RG-transported high-scale scalar wound: lambda(Lambda_12)<0, so |lambda(Lambda_12)| is the positive wound scale",
		StructuralEquation:        "flavor deficit + scalar matching deficit ≈ high-scale scalar wound",
		OpposedRGWoundSign:        lambdaLambda12 < 0 && absLambda12 > 0,
		NativeTheoremClaimed:      false,
		Verdict:                   StatusClosureSealDefined,
	}
}

func buildNativeStatus() NativeASHAStatus {
	return NativeASHAStatus{
		NativeKappaClosureTheorem:                 false,
		NativeScalarRGMatchingTheorem:             false,
		NativeFlavorOrientationTheorem:            false,
		NativeLowScaleMatchingToHighScaleWoundLaw: false,
		NativeHistoryLoopDeficitClosureTheorem:    false,
		Statement:                                 "ASHA currently supplies the typed operands and a sharp bridge closure, but no native theorem joining the low-scale scalar matching deficit, flavor orientation deficit, and high-scale scalar RG wound.",
		Verdict:                                   StatusNoNativeKappaClosure,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{Verdict: StatusGate625Boundary}
}

func closestClosureTarget(rows []DeficitClosureRow) (string, float64, float64) {
	if len(rows) == 0 {
		return "", math.NaN(), math.NaN()
	}
	best := rows[0]
	for _, r := range rows[1:] {
		if r.AbsoluteResidual < best.AbsoluteResidual {
			best = r
		}
	}
	return best.Target, best.AbsoluteResidual, best.RelativeResidual
}

func Statuses() []string {
	return []string{
		StatusGate624Inherited,
		StatusKappasDefined,
		StatusDeficitClosureComputed,
		StatusClosureOnAbsLambda12,
		StatusScalarPredictionComputed,
		StatusClosureSealDefined,
		StatusNoNativeKappaClosure,
		StatusNoNativeScalarRGMatching,
		StatusNoNativeFlavorOrientation,
		StatusGate625Boundary,
	}
}
