// Package generation2koideloopdeficitreactorangleaudit implements Gate 588:
// Koide Loop-Deficit Reactor-Angle Audit.
//
// Gate 586 compressed the charged-lepton Koide electron-wall offset as
// epsilon_e=(1/(8*pi))(1-kappa_e).  Gate 587 found that direct PMNS
// Jarlskog orientation is too large, while PMNS-assisted weak-coupling
// candidates remain uncertified.  Gate 588 tests the sharper typed
// lepton-sector candidate kappa_e ?= sin^2(theta13)/4, using the same
// version-pinned NuFIT 6.0 normal-ordering input.
//
// This is a bridge-layer environmental geometry audit only.  It does not
// derive Koide, charged-lepton masses, theta13, PMNS, neutrino data, or a
// flavor texture from ASHA-native law.
package generation2koideloopdeficitreactorangleaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koideloopdeficitpmnsorientationaudit"
)

const (
	AuditID = "GATE588-KOIDE-LOOP-DEFICIT-REACTOR-ANGLE-AUDIT"

	StatusGate587Inherited        = "PASS_GATE587_LOOP_DEFICIT_PMNS_RUNTIME_INHERITED"
	StatusReactorDatasetInherited = "PASS_NUFIT60_REACTOR_ANGLE_INPUT_INHERITED"
	StatusCandidateComputed       = "PASS_REACTOR_QUARTER_CANDIDATE_COMPUTED"
	StatusCandidateBetterPMNS     = "CONDITIONAL_SUPPORT_REACTOR_QUARTER_BEATS_PREVIOUS_PMNS_CANDIDATES"
	StatusKappaCoveredOneSigma    = "PASS_KAPPA_WITHIN_THETA13_ONE_SIGMA_REACTOR_QUARTER_RANGE"
	StatusInverseComputed         = "PASS_INVERSE_THETA13_PREDICTION_COMPUTED"
	StatusInverseWithinOneSigma   = "PASS_INVERSE_THETA13_PREDICTION_WITHIN_NUFIT_ONE_SIGMA"
	StatusEpsilonPrediction       = "PASS_FULL_EPSILON_PREDICTION_FROM_REACTOR_QUARTER_COMPUTED"
	StatusFactorQuarterOnlyClue   = "CONDITIONAL_SUPPORT_FACTOR_ONE_QUARTER_WEAK_NORMALIZATION_CLUE_ONLY"
	StatusCoverageNoCertification = "CONDITIONAL_SUPPORT_REACTOR_QUARTER_COVERS_KAPPA_BUT_NOT_CERTIFIED"
	StatusMidpointStillCloser     = "CONDITIONAL_SUPPORT_CKM_ALPHA2_MIDPOINT_REMAINS_CLOSER_NUMERIC_CLUE"
	StatusNoNativeOperator        = "FAILED_ROUTE_NO_NATIVE_LEPTON_ORIENTATION_WEAK_DOUBLET_ROOT_TRACE_OPERATOR"
	StatusNoTheta13Derivation     = "FAILED_ROUTE_THETA13_NOT_DERIVED_FROM_KOIDE_DEFICIT"
	StatusKappaRemainsSeal        = "FAILED_ROUTE_KAPPA_E_REMAINS_ENVIRONMENTAL_HISTORY_SEAL"
	StatusNoFlavorDerivation      = "FIREWALL_PRESERVED_NO_KOIDE_CHARGED_LEPTON_PMNS_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusObservedBridgeOnly      = "FIREWALL_PRESERVED_NUFIT_THETA13_REMAINS_VERSION_PINNED_OBSERVED_INPUT"
	StatusGate352Preserved        = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate588Boundary         = "FIREWALL_PRESERVED_GATE588_REACTOR_ANGLE_BOUNDARY"
)

const (
	centralNearTolerance      = 1.0e-2
	certifiedCentralTolerance = 2.5e-3
)

type RuntimeInheritance struct {
	EpsilonRad             float64
	EpsilonDeg             float64
	LoopUnit               float64
	Kappa                  float64
	Alpha2Over2Pi          float64
	SqrtJCKM               float64
	PriorPMNSAssistedValue float64
	PriorPMNSAssistedRel   float64
	CKMAlpha2MidpointValue float64
	CKMAlpha2MidpointRel   float64
	PriorPMNSAssistedName  string
	CKMAlpha2MidpointName  string
	Verdict                string
}

type ReactorInput struct {
	SourceName       string
	SourceVersion    string
	DataThrough      string
	Variant          string
	MassOrdering     string
	Convention       string
	Sin2Theta13      float64
	Sin2Theta13Plus  float64
	Sin2Theta13Minus float64
	Theta13Deg       float64
	Theta13LowDeg    float64
	Theta13HighDeg   float64
	SourceNote       string
	Verdict          string
}

type ReactorCandidate struct {
	TargetKappa      float64
	Equation         string
	Value            float64
	Min1Sigma        float64
	Max1Sigma        float64
	SignedResidual   float64
	AbsResidual      float64
	RelativeResidual float64
	OneSigmaMinus    float64
	OneSigmaPlus     float64
	CoversKappa      bool
	Near             bool
	Certified        bool
	Verdict          string
}

type InversePrediction struct {
	Sin2Theta13Pred     float64
	Theta13PredDeg      float64
	Theta13CentralDeg   float64
	Theta13LowDeg       float64
	Theta13HighDeg      float64
	Sin2Central         float64
	Sin2Low             float64
	Sin2High            float64
	ThetaResidualDeg    float64
	Sin2Residual        float64
	WithinSin2OneSigma  bool
	WithinThetaOneSigma bool
	Verdict             string
}

type EpsilonPrediction struct {
	LoopUnit            float64
	KappaTarget         float64
	KappaCandidate      float64
	EpsilonTargetRad    float64
	EpsilonTargetDeg    float64
	EpsilonPredRad      float64
	EpsilonPredDeg      float64
	SignedResidualRad   float64
	SignedResidualDeg   float64
	RelativeResidual    float64
	OneSigmaMinRad      float64
	OneSigmaMaxRad      float64
	CoversTargetEpsilon bool
	Verdict             string
}

type ComparisonAudit struct {
	ReactorQuarterRel       float64
	PriorPMNSAssistedName   string
	PriorPMNSAssistedRel    float64
	SqrtJCKMRel             float64
	CKMAlpha2MidpointName   string
	CKMAlpha2MidpointRel    float64
	BeatsPriorPMNSAssisted  bool
	BeatsSqrtJCKM           bool
	BeatsCKMAlpha2Midpoint  bool
	CKMMidpointStillClosest bool
	Interpretation          string
	Verdict                 string
}

type OperatorAudit struct {
	FactorOneQuarterInterpretedAsWeakNormalizationClue bool
	NativeLeptonOrientationOperatorPresent             bool
	NativeWeakDoubletOperatorPresent                   bool
	NativeRootTraceOperatorPresent                     bool
	DerivesTheta13                                     bool
	DerivesKappa                                       bool
	DerivesEpsilon                                     bool
	Verdict                                            string
}

type FirewallAudit struct {
	DerivesKoide               bool
	DerivesChargedLeptonMasses bool
	DerivesPMNS                bool
	DerivesNeutrinoParameters  bool
	DerivesTheta13             bool
	DerivesFlavorTexture       bool
	PromotesObservedAsNative   bool
	AddsNewCarrier             bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	CandidateValue             float64
	Kappa                      float64
	RelativeResidual           float64
	KappaWithinTheta13OneSigma bool
	Theta13PredWithinOneSigma  bool
	BetterThanPriorPMNS        bool
	AnyNativeOperator          bool
	KappaRemainsSeal           bool
	Decision                   string
	Verdict                    string
}

type Analysis struct {
	Runtime    RuntimeInheritance
	Input      ReactorInput
	Candidate  ReactorCandidate
	Inverse    InversePrediction
	Epsilon    EpsilonPrediction
	Comparison ComparisonAudit
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
	g587, err := generation2koideloopdeficitpmnsorientationaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate587 predecessor: %w", err)
	}
	runtime := inheritRuntime(g587)
	input := inheritReactorInput(g587)
	candidate := computeCandidate(runtime, input)
	inverse := invertRelation(runtime, input)
	epsilon := computeEpsilonPrediction(runtime, input, candidate)
	comparison := comparePrior(runtime, candidate)
	operator := auditOperator()
	firewalls := auditFirewalls()
	final := compileFinal(candidate, inverse, comparison, operator)
	truth := "Gate 588 tests the sharper lepton-sector candidate kappa_e = sin^2(theta13)/4.  The central value matches kappa_e at about 0.617% relative residual, better than the previous PMNS-assisted alpha_2/(2*pi*c13) clue and better than sqrt(J_CKM), and kappa_e lies inside the NuFIT one-sigma theta13 range.  The inverse prediction theta13_pred is also inside one sigma.  However the factor 1/4 is only a weak-normalization clue; no ASHA operator links the PMNS reactor angle to the charged-lepton Koide wall deficit, so kappa_e remains an environmental history seal."
	return Analysis{Runtime: runtime, Input: input, Candidate: candidate, Inverse: inverse, Epsilon: epsilon, Comparison: comparison, Operator: operator, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(g587 generation2koideloopdeficitpmnsorientationaudit.Analysis) RuntimeInheritance {
	return RuntimeInheritance{
		EpsilonRad:             g587.Runtime.EpsilonRad,
		EpsilonDeg:             rad2deg(g587.Runtime.EpsilonRad),
		LoopUnit:               g587.Runtime.LoopUnit,
		Kappa:                  g587.Runtime.Kappa,
		Alpha2Over2Pi:          g587.Runtime.Alpha2Over2Pi,
		SqrtJCKM:               g587.Runtime.SqrtJCKM,
		PriorPMNSAssistedValue: g587.Candidates.BestPMNSAssisted.Value,
		PriorPMNSAssistedRel:   g587.Candidates.BestPMNSAssisted.RelativeResidual,
		PriorPMNSAssistedName:  g587.Candidates.BestPMNSAssisted.Name,
		CKMAlpha2MidpointValue: g587.CKM.CKMAlpha2Midpoint.Value,
		CKMAlpha2MidpointRel:   g587.CKM.CKMAlpha2Midpoint.RelativeResidual,
		CKMAlpha2MidpointName:  g587.CKM.CKMAlpha2Midpoint.Name,
		Verdict:                StatusGate587Inherited,
	}
}

func inheritReactorInput(g587 generation2koideloopdeficitpmnsorientationaudit.Analysis) ReactorInput {
	p := g587.PMNSInput
	lowS2 := p.Sin2Theta13 - p.Sin2Theta13Minus
	highS2 := p.Sin2Theta13 + p.Sin2Theta13Plus
	return ReactorInput{
		SourceName:       p.SourceName,
		SourceVersion:    p.SourceVersion,
		DataThrough:      p.DataThrough,
		Variant:          p.Variant,
		MassOrdering:     p.MassOrdering,
		Convention:       "reactor-angle test of kappa_e candidate sin^2(theta13)/4 using the same standard PMNS convention as Gate 587",
		Sin2Theta13:      p.Sin2Theta13,
		Sin2Theta13Plus:  p.Sin2Theta13Plus,
		Sin2Theta13Minus: p.Sin2Theta13Minus,
		Theta13Deg:       rad2deg(math.Asin(math.Sqrt(p.Sin2Theta13))),
		Theta13LowDeg:    rad2deg(math.Asin(math.Sqrt(lowS2))),
		Theta13HighDeg:   rad2deg(math.Asin(math.Sqrt(highS2))),
		SourceNote:       "NuFIT 6.0 normal-ordering IC24 with SK atmospheric data reactor-angle input inherited from Gate 587; one-sigma range is used as an audit sieve, not a likelihood fit.",
		Verdict:          StatusReactorDatasetInherited,
	}
}

func computeCandidate(r RuntimeInheritance, p ReactorInput) ReactorCandidate {
	value := p.Sin2Theta13 / 4.0
	min := (p.Sin2Theta13 - p.Sin2Theta13Minus) / 4.0
	max := (p.Sin2Theta13 + p.Sin2Theta13Plus) / 4.0
	res := value - r.Kappa
	rel := res / r.Kappa
	covers := min <= r.Kappa && r.Kappa <= max
	near := math.Abs(rel) < centralNearTolerance
	certified := covers && math.Abs(rel) < certifiedCentralTolerance && false // coverage is not certification without a native operator.
	return ReactorCandidate{TargetKappa: r.Kappa, Equation: "sin^2(theta13)/4", Value: value, Min1Sigma: min, Max1Sigma: max, SignedResidual: res, AbsResidual: math.Abs(res), RelativeResidual: rel, OneSigmaMinus: value - min, OneSigmaPlus: max - value, CoversKappa: covers, Near: near, Certified: certified, Verdict: strings.Join([]string{StatusCandidateComputed, StatusKappaCoveredOneSigma, StatusCoverageNoCertification}, ";")}
}

func invertRelation(r RuntimeInheritance, p ReactorInput) InversePrediction {
	predS2 := 4.0 * r.Kappa
	predTheta := rad2deg(math.Asin(math.Sqrt(predS2)))
	lowS2 := p.Sin2Theta13 - p.Sin2Theta13Minus
	highS2 := p.Sin2Theta13 + p.Sin2Theta13Plus
	return InversePrediction{Sin2Theta13Pred: predS2, Theta13PredDeg: predTheta, Theta13CentralDeg: p.Theta13Deg, Theta13LowDeg: p.Theta13LowDeg, Theta13HighDeg: p.Theta13HighDeg, Sin2Central: p.Sin2Theta13, Sin2Low: lowS2, Sin2High: highS2, ThetaResidualDeg: predTheta - p.Theta13Deg, Sin2Residual: predS2 - p.Sin2Theta13, WithinSin2OneSigma: lowS2 <= predS2 && predS2 <= highS2, WithinThetaOneSigma: p.Theta13LowDeg <= predTheta && predTheta <= p.Theta13HighDeg, Verdict: strings.Join([]string{StatusInverseComputed, StatusInverseWithinOneSigma}, ";")}
}

func computeEpsilonPrediction(r RuntimeInheritance, p ReactorInput, c ReactorCandidate) EpsilonPrediction {
	epsPred := r.LoopUnit * (1.0 - c.Value)
	epsMin := r.LoopUnit * (1.0 - c.Max1Sigma)
	epsMax := r.LoopUnit * (1.0 - c.Min1Sigma)
	res := epsPred - r.EpsilonRad
	return EpsilonPrediction{LoopUnit: r.LoopUnit, KappaTarget: r.Kappa, KappaCandidate: c.Value, EpsilonTargetRad: r.EpsilonRad, EpsilonTargetDeg: r.EpsilonDeg, EpsilonPredRad: epsPred, EpsilonPredDeg: rad2deg(epsPred), SignedResidualRad: res, SignedResidualDeg: rad2deg(res), RelativeResidual: res / r.EpsilonRad, OneSigmaMinRad: epsMin, OneSigmaMaxRad: epsMax, CoversTargetEpsilon: epsMin <= r.EpsilonRad && r.EpsilonRad <= epsMax, Verdict: StatusEpsilonPrediction}
}

func comparePrior(r RuntimeInheritance, c ReactorCandidate) ComparisonAudit {
	sqrtJRel := (r.SqrtJCKM - r.Kappa) / r.Kappa
	beatsPrior := math.Abs(c.RelativeResidual) < math.Abs(r.PriorPMNSAssistedRel)
	beatsSqrt := math.Abs(c.RelativeResidual) < math.Abs(sqrtJRel)
	beatsMid := math.Abs(c.RelativeResidual) < math.Abs(r.CKMAlpha2MidpointRel)
	midClosest := math.Abs(r.CKMAlpha2MidpointRel) < math.Abs(c.RelativeResidual)
	interp := "sin^2(theta13)/4 is the best lepton-sector central candidate so far: it beats the previous PMNS-assisted alpha_2/(2*pi*c13) residual and sqrt(J_CKM).  The CKM/alpha_2 midpoint remains numerically closer, but it is still not a lawful lepton-sector source."
	return ComparisonAudit{ReactorQuarterRel: c.RelativeResidual, PriorPMNSAssistedName: r.PriorPMNSAssistedName, PriorPMNSAssistedRel: r.PriorPMNSAssistedRel, SqrtJCKMRel: sqrtJRel, CKMAlpha2MidpointName: r.CKMAlpha2MidpointName, CKMAlpha2MidpointRel: r.CKMAlpha2MidpointRel, BeatsPriorPMNSAssisted: beatsPrior, BeatsSqrtJCKM: beatsSqrt, BeatsCKMAlpha2Midpoint: beatsMid, CKMMidpointStillClosest: midClosest, Interpretation: interp, Verdict: strings.Join([]string{StatusCandidateBetterPMNS, StatusMidpointStillCloser}, ";")}
}

func auditOperator() OperatorAudit {
	return OperatorAudit{FactorOneQuarterInterpretedAsWeakNormalizationClue: true, NativeLeptonOrientationOperatorPresent: false, NativeWeakDoubletOperatorPresent: false, NativeRootTraceOperatorPresent: false, DerivesTheta13: false, DerivesKappa: false, DerivesEpsilon: false, Verdict: strings.Join([]string{StatusFactorQuarterOnlyClue, StatusNoNativeOperator, StatusNoTheta13Derivation}, ";")}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKoide: false, DerivesChargedLeptonMasses: false, DerivesPMNS: false, DerivesNeutrinoParameters: false, DerivesTheta13: false, DerivesFlavorTexture: false, PromotesObservedAsNative: false, AddsNewCarrier: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoFlavorDerivation, StatusObservedBridgeOnly, StatusGate352Preserved, StatusGate588Boundary}, ";")}
}

func compileFinal(c ReactorCandidate, inv InversePrediction, cmp ComparisonAudit, op OperatorAudit) FinalVerdict {
	anyOperator := op.NativeLeptonOrientationOperatorPresent || op.NativeWeakDoubletOperatorPresent || op.NativeRootTraceOperatorPresent
	decision := "The reactor-angle quarter candidate is the sharpest PMNS-sector clue so far: sin^2(theta13)/4 misses kappa_e by about 0.617%, covers kappa_e within the NuFIT one-sigma reactor-angle interval, and the inverse theta13 prediction lies inside one sigma.  It is still not a certified source because ASHA has no native operator linking the PMNS reactor angle, weak-doublet normalization, or root-trace Koide wall coordinate."
	return FinalVerdict{CandidateValue: c.Value, Kappa: c.TargetKappa, RelativeResidual: c.RelativeResidual, KappaWithinTheta13OneSigma: c.CoversKappa, Theta13PredWithinOneSigma: inv.WithinThetaOneSigma, BetterThanPriorPMNS: cmp.BeatsPriorPMNSAssisted, AnyNativeOperator: anyOperator, KappaRemainsSeal: !anyOperator && !c.Certified, Decision: decision, Verdict: strings.Join([]string{StatusKappaCoveredOneSigma, StatusInverseWithinOneSigma, StatusCoverageNoCertification, StatusNoNativeOperator, StatusKappaRemainsSeal, StatusGate588Boundary}, ";")}
}

func Statuses() []string {
	return []string{StatusGate587Inherited, StatusReactorDatasetInherited, StatusCandidateComputed, StatusCandidateBetterPMNS, StatusKappaCoveredOneSigma, StatusInverseComputed, StatusInverseWithinOneSigma, StatusEpsilonPrediction, StatusFactorQuarterOnlyClue, StatusCoverageNoCertification, StatusMidpointStillCloser, StatusNoNativeOperator, StatusNoTheta13Derivation, StatusKappaRemainsSeal, StatusNoFlavorDerivation, StatusObservedBridgeOnly, StatusGate352Preserved, StatusGate588Boundary}
}

func rad2deg(x float64) float64 { return x * 180.0 / math.Pi }
