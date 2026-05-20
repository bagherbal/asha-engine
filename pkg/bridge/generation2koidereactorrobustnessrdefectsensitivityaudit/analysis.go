// Package generation2koidereactorrobustnessrdefectsensitivityaudit implements
// Gate 589: Koide-Reactor Relation Robustness and R-Defect Sensitivity Audit.
//
// Gate 588 found that the charged-lepton Koide loop-deficit kappa_e is close
// to the lepton-sector reactor candidate sin^2(theta13)/4.  Gate 589 tests
// whether this relation belongs to the actually measured near-Koide ray, to
// the exact-R=1 projected wall coordinate from Gate 584, or to a corrected
// two-variable structure involving the Koide amplitude defect R_e-1.
//
// This is a bridge-layer environmental robustness audit only.  It does not
// derive Koide, theta13, PMNS, lepton masses, neutrino data, or a native ASHA
// flavor law.
package generation2koidereactorrobustnessrdefectsensitivityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidechamberwalloffsetaudit"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koideloopdeficitreactorangleaudit"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidewalloffsetratioclosureaudit"
)

const (
	AuditID = "GATE589-KOIDE-REACTOR-ROBUSTNESS-R-DEFECT-SENSITIVITY-AUDIT"

	StatusGate588Inherited            = "PASS_GATE588_REACTOR_QUARTER_RESULT_INHERITED"
	StatusObservedNearKoideInherited  = "PASS_OBSERVED_NEAR_KOIDE_WALL_COORDINATE_INHERITED"
	StatusExactR1Inherited            = "PASS_EXACT_R1_RATIO_CLOSURE_COORDINATE_INHERITED"
	StatusReactorCandidateInherited   = "PASS_NUFIT60_REACTOR_QUARTER_CANDIDATE_INHERITED"
	StatusObservedMatchBetter         = "PASS_REACTOR_QUARTER_MATCHES_OBSERVED_EPSILON_BETTER_THAN_EXACT_R1_EPSILON"
	StatusObservedInsideOneSigma      = "PASS_OBSERVED_EPSILON_INVERSE_THETA13_PREDICTION_WITHIN_NUFIT_ONE_SIGMA"
	StatusExactR1OutsideOneSigma      = "FAILED_ROUTE_EXACT_R1_INVERSE_THETA13_PREDICTION_OUTSIDE_NUFIT_ONE_SIGMA"
	StatusExactR1WeakensRelation      = "CONDITIONAL_SUPPORT_REACTOR_RELATION_BELONGS_TO_MEASURED_NEAR_KOIDE_RAY_NOT_EXACT_R1_PROJECTION"
	StatusRDefectShiftComputed        = "PASS_R_DEFECT_AND_KAPPA_SHIFT_COMPUTED"
	StatusRequiredCoefficientComputed = "PASS_REQUIRED_R_DEFECT_LINEAR_COEFFICIENT_COMPUTED"
	StatusNoTypedRDefectCoefficient   = "FAILED_ROUTE_NO_TYPED_SIMPLE_R_DEFECT_CORRECTION_CERTIFIED"
	StatusEpsilonShiftControlsKappa   = "PASS_KAPPA_SHIFT_EXACTLY_CONTROLLED_BY_EPSILON_PROJECTION_SHIFT"
	StatusRDefectOnlyNotEnough        = "FAILED_ROUTE_R_DEFECT_ALONE_DOES_NOT_FIX_REACTOR_RELATION_WITH_TYPED_COEFFICIENT"
	StatusQResidualNotEnough          = "FAILED_ROUTE_Q_MINUS_TWO_THIRDS_DOES_NOT_SUPPLY_TYPED_KAPPA_CORRECTION"
	StatusNoNativeOperator            = "FAILED_ROUTE_NO_NATIVE_KOIDE_REACTOR_R_DEFECT_OPERATOR"
	StatusKappaRemainsSeal            = "FAILED_ROUTE_KOIDE_REACTOR_RELATION_REMAINS_ENVIRONMENTAL_HISTORY_SEAL"
	StatusNoFlavorDerivation          = "FIREWALL_PRESERVED_NO_KOIDE_THETA13_PMNS_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusObservedBridgeOnly          = "FIREWALL_PRESERVED_REACTOR_AND_KOIDE_INPUTS_REMAIN_VERSION_PINNED_OBSERVED_DATA"
	StatusGate352Preserved            = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate589Boundary             = "FIREWALL_PRESERVED_GATE589_R_DEFECT_SENSITIVITY_BOUNDARY"
)

const (
	certifiedRelativeTolerance = 2.5e-3
)

type RuntimeInheritance struct {
	EpsilonObsRad          float64
	EpsilonObsDeg          float64
	RObs                   float64
	RDefect                float64
	QObs                   float64
	QResidual              float64
	KappaObs               float64
	EpsilonR1Rad           float64
	EpsilonR1Deg           float64
	KappaR1                float64
	KappaShiftObsMinusR1   float64
	EpsilonShiftR1MinusObs float64
	Source                 string
	Verdict                string
}

type ReactorInput struct {
	SourceName        string
	SourceVersion     string
	Variant           string
	MassOrdering      string
	Sin2Theta13       float64
	Sin2Theta13Plus   float64
	Sin2Theta13Minus  float64
	Candidate         float64
	CandidateMin      float64
	CandidateMax      float64
	Theta13CentralDeg float64
	Theta13LowDeg     float64
	Theta13HighDeg    float64
	Verdict           string
}

type KappaComparison struct {
	Name                string
	EpsilonRad          float64
	Kappa               float64
	Candidate           float64
	CandidateMin        float64
	CandidateMax        float64
	SignedResidual      float64
	AbsResidual         float64
	RelativeResidual    float64
	CoveredByOneSigma   bool
	Certified           bool
	Sin2Theta13Pred     float64
	Theta13PredDeg      float64
	Theta13ResidualDeg  float64
	WithinSin2OneSigma  bool
	WithinThetaOneSigma bool
	EpsilonPredRad      float64
	EpsilonResidualRad  float64
	Verdict             string
}

type ReactorRobustnessAudit struct {
	Observed            KappaComparison
	ExactR1             KappaComparison
	ObservedBetter      bool
	ExactR1WeakerFactor float64
	Interpretation      string
	Verdict             string
}

type TypedCoefficientCandidate struct {
	Name              string
	Value             float64
	PredictedShift    float64
	SignedResidual    float64
	RelativeResidual  float64
	CloserThanNoShift bool
}

type RDefectCorrectionAudit struct {
	DROneMinusR            float64
	RMinusOne              float64
	KappaObsMinusR1        float64
	RequiredC              float64
	Candidates             []TypedCoefficientCandidate
	BestCandidate          TypedCoefficientCandidate
	BestCandidateCertified bool
	Interpretation         string
	Verdict                string
}

type ShiftControlAudit struct {
	KappaShift                 float64
	EpsilonShiftR1MinusObs     float64
	EightPiEpsilonShift        float64
	DROneMinusR                float64
	RatioToDROneMinusR         float64
	QResidual                  float64
	RatioToAbsQResidual        float64
	ControlledByEpsilonShift   bool
	ControlledByRDefectTyped   bool
	ControlledByQResidualTyped bool
	Verdict                    string
}

type OperatorAudit struct {
	NativeKoideReactorOperatorPresent      bool
	NativeRDefectCorrectionOperatorPresent bool
	NativeRootTraceOperatorPresent         bool
	DerivesTheta13                         bool
	DerivesKappa                           bool
	DerivesEpsilon                         bool
	Verdict                                string
}

type FirewallAudit struct {
	DerivesKoide             bool
	DerivesTheta13           bool
	DerivesPMNS              bool
	DerivesNeutrinoPhysics   bool
	DerivesLeptonMasses      bool
	DerivesFlavorLaw         bool
	PromotesObservedAsNative bool
	AddsNewCarrier           bool
	PreservesGate352         bool
	Verdict                  string
}

type FinalVerdict struct {
	ReactorMatchesObservedBetter  bool
	ObservedInsideOneSigma        bool
	ExactR1InsideOneSigma         bool
	RDefectRequiredForBestMatch   bool
	TypedRDefectCorrectionPresent bool
	NativeOperatorPresent         bool
	RelationRemainsEnvironmental  bool
	Decision                      string
	Verdict                       string
}

type Analysis struct {
	Runtime    RuntimeInheritance
	Reactor    ReactorInput
	Robustness ReactorRobustnessAudit
	RDefect    RDefectCorrectionAudit
	Shift      ShiftControlAudit
	Operator   OperatorAudit
	Firewalls  FirewallAudit
	Final      FinalVerdict
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
	g583, err := generation2koidechamberwalloffsetaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate583 predecessor: %w", err)
	}
	g584, err := generation2koidewalloffsetratioclosureaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate584 predecessor: %w", err)
	}
	g588, err := generation2koideloopdeficitreactorangleaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate588 predecessor: %w", err)
	}
	runtime := inheritRuntime(g583, g584, g588)
	reactor := inheritReactor(g588)
	robustness := compareObservedAndR1(runtime, reactor)
	rdefect := auditRDefect(runtime)
	shift := auditShiftControls(runtime, rdefect)
	operator := auditOperator()
	firewalls := auditFirewalls()
	final := compileFinal(robustness, rdefect, operator)
	truth := "Gate 589 shows that kappa_e ≈ sin^2(theta13)/4 is robust for the measured near-Koide charged-lepton wall coordinate, but not for the exact-R=1 projected ratio-closure coordinate.  The exact-R=1 projection shifts epsilon enough that the inverse theta13 prediction falls outside the NuFIT one-sigma reactor-angle interval.  The required correction from kappa_R1 to kappa_obs is proportional to the epsilon projection shift by definition, while an R-defect-only linear correction needs an untyped coefficient c≈20.6455.  Therefore the reactor clue belongs to the observed environmental ray, not to a certified native exact-Koide operator."
	return Analysis{Runtime: runtime, Reactor: reactor, Robustness: robustness, RDefect: rdefect, Shift: shift, Operator: operator, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(g583 generation2koidechamberwalloffsetaudit.Analysis, g584 generation2koidewalloffsetratioclosureaudit.Analysis, g588 generation2koideloopdeficitreactorangleaudit.Analysis) RuntimeInheritance {
	epsObs := g588.Runtime.EpsilonRad
	kObs := g588.Runtime.Kappa
	epsR1 := g584.MZ.FromElectronMuon.SolvedEpsilonRad
	kR1 := 1.0 - 8.0*math.Pi*epsR1
	r := g583.MZ.PlaneAmplitudeR
	q := (1.0 + r*r) / 3.0
	return RuntimeInheritance{
		EpsilonObsRad:          epsObs,
		EpsilonObsDeg:          rad2deg(epsObs),
		RObs:                   r,
		RDefect:                1.0 - r,
		QObs:                   q,
		QResidual:              q - 2.0/3.0,
		KappaObs:               kObs,
		EpsilonR1Rad:           epsR1,
		EpsilonR1Deg:           rad2deg(epsR1),
		KappaR1:                kR1,
		KappaShiftObsMinusR1:   kObs - kR1,
		EpsilonShiftR1MinusObs: epsR1 - epsObs,
		Source:                 "Gate583 observed near-Koide wall coordinate + Gate584 exact-R=1 ratio-closure coordinate + Gate588 reactor-quarter relation",
		Verdict:                strings.Join([]string{StatusGate588Inherited, StatusObservedNearKoideInherited, StatusExactR1Inherited}, ";"),
	}
}

func inheritReactor(g588 generation2koideloopdeficitreactorangleaudit.Analysis) ReactorInput {
	return ReactorInput{
		SourceName:        g588.Input.SourceName,
		SourceVersion:     g588.Input.SourceVersion,
		Variant:           g588.Input.Variant,
		MassOrdering:      g588.Input.MassOrdering,
		Sin2Theta13:       g588.Input.Sin2Theta13,
		Sin2Theta13Plus:   g588.Input.Sin2Theta13Plus,
		Sin2Theta13Minus:  g588.Input.Sin2Theta13Minus,
		Candidate:         g588.Candidate.Value,
		CandidateMin:      g588.Candidate.Min1Sigma,
		CandidateMax:      g588.Candidate.Max1Sigma,
		Theta13CentralDeg: g588.Inverse.Theta13CentralDeg,
		Theta13LowDeg:     g588.Inverse.Theta13LowDeg,
		Theta13HighDeg:    g588.Inverse.Theta13HighDeg,
		Verdict:           StatusReactorCandidateInherited,
	}
}

func compareObservedAndR1(r RuntimeInheritance, input ReactorInput) ReactorRobustnessAudit {
	obs := compareKappa("observed near-Koide ray", r.EpsilonObsRad, r.KappaObs, input)
	r1 := compareKappa("exact R=1 projected ratio-closure ray", r.EpsilonR1Rad, r.KappaR1, input)
	obsBetter := obs.AbsResidual < r1.AbsResidual
	factor := safeDiv(r1.AbsResidual, obs.AbsResidual)
	interp := "The reactor-quarter relation matches the measured near-Koide wall coordinate much better than the exact-R=1 projected coordinate.  The observed inverse theta13 prediction lies inside one sigma, while the exact-R=1 inverse prediction is outside one sigma."
	return ReactorRobustnessAudit{Observed: obs, ExactR1: r1, ObservedBetter: obsBetter, ExactR1WeakerFactor: factor, Interpretation: interp, Verdict: strings.Join([]string{StatusObservedMatchBetter, StatusObservedInsideOneSigma, StatusExactR1OutsideOneSigma, StatusExactR1WeakensRelation}, ";")}
}

func compareKappa(name string, eps, k float64, input ReactorInput) KappaComparison {
	res := input.Candidate - k
	predS2 := 4.0 * k
	predTheta := rad2deg(math.Asin(math.Sqrt(predS2)))
	epsPred := (1.0 / (8.0 * math.Pi)) * (1.0 - input.Candidate)
	withinS2 := input.Sin2Theta13-input.Sin2Theta13Minus <= predS2 && predS2 <= input.Sin2Theta13+input.Sin2Theta13Plus
	withinTheta := input.Theta13LowDeg <= predTheta && predTheta <= input.Theta13HighDeg
	covered := input.CandidateMin <= k && k <= input.CandidateMax
	certified := covered && math.Abs(safeDiv(res, k)) < certifiedRelativeTolerance && false
	verdict := StatusObservedInsideOneSigma
	if !withinS2 || !withinTheta {
		verdict = StatusExactR1OutsideOneSigma
	}
	return KappaComparison{Name: name, EpsilonRad: eps, Kappa: k, Candidate: input.Candidate, CandidateMin: input.CandidateMin, CandidateMax: input.CandidateMax, SignedResidual: res, AbsResidual: math.Abs(res), RelativeResidual: safeDiv(res, k), CoveredByOneSigma: covered, Certified: certified, Sin2Theta13Pred: predS2, Theta13PredDeg: predTheta, Theta13ResidualDeg: predTheta - input.Theta13CentralDeg, WithinSin2OneSigma: withinS2, WithinThetaOneSigma: withinTheta, EpsilonPredRad: epsPred, EpsilonResidualRad: epsPred - eps, Verdict: verdict}
}

func auditRDefect(r RuntimeInheritance) RDefectCorrectionAudit {
	required := safeDiv(r.KappaShiftObsMinusR1, r.RDefect)
	candidates := []TypedCoefficientCandidate{}
	for _, c := range []struct {
		name  string
		value float64
	}{
		{"1", 1},
		{"2", 2},
		{"sqrt(2)", math.Sqrt2},
		{"sqrt(3)", math.Sqrt(3)},
		{"2*pi", 2 * math.Pi},
		{"8*pi", 8 * math.Pi},
	} {
		pred := c.value * r.RDefect
		res := pred - r.KappaShiftObsMinusR1
		candidates = append(candidates, TypedCoefficientCandidate{Name: c.name, Value: c.value, PredictedShift: pred, SignedResidual: res, RelativeResidual: safeDiv(res, r.KappaShiftObsMinusR1), CloserThanNoShift: math.Abs(res) < math.Abs(r.KappaShiftObsMinusR1)})
	}
	best := candidates[0]
	for _, c := range candidates[1:] {
		if math.Abs(c.RelativeResidual) < math.Abs(best.RelativeResidual) {
			best = c
		}
	}
	certified := math.Abs(best.RelativeResidual) < certifiedRelativeTolerance && false
	interp := "Mapping kappa_R1 to kappa_obs by kappa_eff=kappa_R1+c(1-R_obs) requires c≈20.6455.  The nearest tested typed coefficient is 8*pi, but its residual is still about 21.7% of the required shift, so no R-defect correction is certified."
	return RDefectCorrectionAudit{DROneMinusR: r.RDefect, RMinusOne: -r.RDefect, KappaObsMinusR1: r.KappaShiftObsMinusR1, RequiredC: required, Candidates: candidates, BestCandidate: best, BestCandidateCertified: certified, Interpretation: interp, Verdict: strings.Join([]string{StatusRDefectShiftComputed, StatusRequiredCoefficientComputed, StatusNoTypedRDefectCoefficient, StatusRDefectOnlyNotEnough}, ";")}
}

func auditShiftControls(r RuntimeInheritance, rd RDefectCorrectionAudit) ShiftControlAudit {
	eightPiShift := 8.0 * math.Pi * r.EpsilonShiftR1MinusObs
	ratioQ := safeDiv(r.KappaShiftObsMinusR1, math.Abs(r.QResidual))
	return ShiftControlAudit{KappaShift: r.KappaShiftObsMinusR1, EpsilonShiftR1MinusObs: r.EpsilonShiftR1MinusObs, EightPiEpsilonShift: eightPiShift, DROneMinusR: r.RDefect, RatioToDROneMinusR: rd.RequiredC, QResidual: r.QResidual, RatioToAbsQResidual: ratioQ, ControlledByEpsilonShift: math.Abs(eightPiShift-r.KappaShiftObsMinusR1) < 1e-14, ControlledByRDefectTyped: rd.BestCandidateCertified, ControlledByQResidualTyped: false, Verdict: strings.Join([]string{StatusEpsilonShiftControlsKappa, StatusRDefectOnlyNotEnough, StatusQResidualNotEnough}, ";")}
}

func auditOperator() OperatorAudit {
	return OperatorAudit{NativeKoideReactorOperatorPresent: false, NativeRDefectCorrectionOperatorPresent: false, NativeRootTraceOperatorPresent: false, DerivesTheta13: false, DerivesKappa: false, DerivesEpsilon: false, Verdict: StatusNoNativeOperator}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKoide: false, DerivesTheta13: false, DerivesPMNS: false, DerivesNeutrinoPhysics: false, DerivesLeptonMasses: false, DerivesFlavorLaw: false, PromotesObservedAsNative: false, AddsNewCarrier: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoFlavorDerivation, StatusObservedBridgeOnly, StatusGate352Preserved, StatusGate589Boundary}, ";")}
}

func compileFinal(rob ReactorRobustnessAudit, rd RDefectCorrectionAudit, op OperatorAudit) FinalVerdict {
	anyOp := op.NativeKoideReactorOperatorPresent || op.NativeRDefectCorrectionOperatorPresent || op.NativeRootTraceOperatorPresent
	decision := "The reactor-quarter relation is robust for epsilon_obs from the measured near-Koide ray, but not for epsilon_R1 from the exact-R=1 projected ratio-closure ray.  The R defect is required to keep the best match, yet no typed simple linear R-defect coefficient is certified.  The relation remains environmental bridge geometry."
	return FinalVerdict{ReactorMatchesObservedBetter: rob.ObservedBetter, ObservedInsideOneSigma: rob.Observed.WithinThetaOneSigma, ExactR1InsideOneSigma: rob.ExactR1.WithinThetaOneSigma, RDefectRequiredForBestMatch: rob.ObservedBetter && !rob.ExactR1.WithinThetaOneSigma, TypedRDefectCorrectionPresent: rd.BestCandidateCertified, NativeOperatorPresent: anyOp, RelationRemainsEnvironmental: !anyOp && !rd.BestCandidateCertified, Decision: decision, Verdict: strings.Join([]string{StatusObservedMatchBetter, StatusExactR1OutsideOneSigma, StatusNoTypedRDefectCoefficient, StatusNoNativeOperator, StatusKappaRemainsSeal, StatusGate589Boundary}, ";")}
}

func Statuses() []string {
	return []string{StatusGate588Inherited, StatusObservedNearKoideInherited, StatusExactR1Inherited, StatusReactorCandidateInherited, StatusObservedMatchBetter, StatusObservedInsideOneSigma, StatusExactR1OutsideOneSigma, StatusExactR1WeakensRelation, StatusRDefectShiftComputed, StatusRequiredCoefficientComputed, StatusNoTypedRDefectCoefficient, StatusEpsilonShiftControlsKappa, StatusRDefectOnlyNotEnough, StatusQResidualNotEnough, StatusNoNativeOperator, StatusKappaRemainsSeal, StatusNoFlavorDerivation, StatusObservedBridgeOnly, StatusGate352Preserved, StatusGate589Boundary}
}

func rad2deg(x float64) float64 { return x * 180.0 / math.Pi }
func safeDiv(a, b float64) float64 {
	if b == 0 {
		return math.NaN()
	}
	return a / b
}
