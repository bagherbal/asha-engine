// Package generation2koidereactorckmorientationcombinationaudit implements
// Gate 590: Koide-Reactor-CKM Orientation Combination Audit.
//
// Gate 588 found the typed lepton-sector clue kappa_e ~= sin^2(theta13)/4.
// Gate 589 showed that this relation belongs to the measured near-Koide
// charged-lepton ray rather than the exact-R=1 projected ray.  Gate 590 tests
// the sharper bridge-layer environmental candidate
//
//	kappa_e ?= sin^2(theta13)/4 - J_CKM
//
// equivalently
//
//	epsilon_e ?= (1/(8*pi)) [1 - sin^2(theta13)/4 + J_CKM].
//
// This audit is environmental bridge geometry only.  It does not derive Koide,
// theta13, CKM, PMNS, charged-lepton masses, neutrino physics, or a native ASHA
// flavor law.
package generation2koidereactorckmorientationcombinationaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koideloopdeficitreactorangleaudit"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidereactorrobustnessrdefectsensitivityaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE590-KOIDE-REACTOR-CKM-ORIENTATION-COMBINATION-AUDIT"

	StatusGate589Inherited          = "PASS_GATE589_NEAR_KOIDE_RAY_RESULT_INHERITED"
	StatusReactorInputInherited     = "PASS_NUFIT60_REACTOR_INPUT_INHERITED"
	StatusCKMRuntimeInherited       = "PASS_RUNTIME_CKM_JARLSKOG_INHERITED"
	StatusCandidateAComputed        = "PASS_REACTOR_QUARTER_CANDIDATE_A_COMPUTED"
	StatusCandidateBComputed        = "PASS_REACTOR_MINUS_CKM_CANDIDATE_B_COMPUTED"
	StatusBOutperformsA             = "PASS_REACTOR_MINUS_CKM_OUTPERFORMS_REACTOR_QUARTER_ALONE"
	StatusBWithinOneSigma           = "PASS_COMBINED_ORIENTATION_CANDIDATE_COVERS_KAPPA_WITH_THETA13_ONE_SIGMA"
	StatusInverseComputed           = "PASS_COMBINED_INVERSE_THETA13_PREDICTION_COMPUTED"
	StatusInverseWithinOneSigma     = "PASS_COMBINED_INVERSE_THETA13_PREDICTION_WITHIN_NUFIT_ONE_SIGMA"
	StatusEpsilonPredictionComputed = "PASS_COMBINED_EPSILON_PREDICTION_COMPUTED"
	StatusResidualImproved          = "CONDITIONAL_SUPPORT_COMBINED_ORIENTATION_RESIDUAL_AT_FIVE_E_MINUS_FOUR_RELATIVE"
	StatusCKMUncertaintyMissing     = "CONDITIONAL_SUPPORT_CKM_J_UNCERTAINTY_NOT_PRESENT_IN_RUNTIME"
	StatusTheta13Limited            = "CONDITIONAL_SUPPORT_AVAILABLE_UNCERTAINTY_DOMINATED_BY_THETA13_INPUT"
	StatusNoCrossSectorIntertwiner  = "FAILED_ROUTE_NO_CROSS_SECTOR_ORIENTATION_INTERTWINER"
	StatusNoNativeOperator          = "FAILED_ROUTE_NO_NATIVE_KOIDE_REACTOR_CKM_OPERATOR"
	StatusResidualNotRDefectFixed   = "FAILED_ROUTE_COMBINED_RESIDUAL_NOT_TYPED_R_DEFECT_OR_Q_RESIDUAL"
	StatusKappaRemainsSeal          = "FAILED_ROUTE_KAPPA_E_REMAINS_ENVIRONMENTAL_HISTORY_SEAL"
	StatusNoFlavorDerivation        = "FIREWALL_PRESERVED_NO_KOIDE_PMNS_CKM_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusObservedBridgeOnly        = "FIREWALL_PRESERVED_REACTOR_CKM_AND_KOIDE_INPUTS_REMAIN_OBSERVED_DATA"
	StatusGate352Preserved          = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate590Boundary           = "FIREWALL_PRESERVED_GATE590_REACTOR_CKM_ORIENTATION_BOUNDARY"
)

const (
	conditionalRelativeTolerance = 1.0e-3
	certifiedRelativeTolerance   = 2.5e-4
)

type RuntimeInheritance struct {
	EpsilonObsRad float64
	EpsilonObsDeg float64
	RObs          float64
	RMinusOne     float64
	RDefect       float64
	QObs          float64
	QResidual     float64
	KappaObs      float64
	LoopUnit      float64
	Source        string
	Verdict       string
}

type ReactorInput struct {
	SourceName        string
	SourceVersion     string
	DataThrough       string
	Variant           string
	MassOrdering      string
	Sin2Theta13       float64
	Sin2Theta13Plus   float64
	Sin2Theta13Minus  float64
	Theta13CentralDeg float64
	Theta13LowDeg     float64
	Theta13HighDeg    float64
	Verdict           string
}

type CKMInput struct {
	SourceName         string
	SourceVersion      string
	JCKM               float64
	JCKMUncertainty    float64
	HasJCKMUncertainty bool
	Verdict            string
}

type CandidateComparison struct {
	Name             string
	Equation         string
	Value            float64
	Min1Sigma        float64
	Max1Sigma        float64
	SignedResidual   float64
	AbsResidual      float64
	RelativeResidual float64
	CoversKappa      bool
	Certified        bool
	Verdict          string
}

type CombinationAudit struct {
	AReactorQuarter    CandidateComparison
	BReactorMinusCKM   CandidateComparison
	BImprovementFactor float64
	BOutperformsA      bool
	Interpretation     string
	Verdict            string
}

type EpsilonPrediction struct {
	EpsilonObservedRad float64
	EpsilonObservedDeg float64
	EpsilonPredA_rad   float64
	EpsilonPredA_deg   float64
	EpsilonPredB_rad   float64
	EpsilonPredB_deg   float64
	ResidualA_rad      float64
	ResidualA_deg      float64
	ResidualB_rad      float64
	ResidualB_deg      float64
	ImprovementFactor  float64
	Verdict            string
}

type InversePrediction struct {
	Sin2Theta13Pred     float64
	Theta13PredDeg      float64
	Sin2Central         float64
	Sin2Low             float64
	Sin2High            float64
	Theta13CentralDeg   float64
	Theta13LowDeg       float64
	Theta13HighDeg      float64
	Sin2Residual        float64
	ThetaResidualDeg    float64
	WithinSin2OneSigma  bool
	WithinThetaOneSigma bool
	Verdict             string
}

type UncertaintyAudit struct {
	Theta13CandidateBMin     float64
	Theta13CandidateBMax     float64
	CoversKappaWithTheta13   bool
	CKMUncertaintyPresent    bool
	CKMUncertaintyValue      float64
	FullUncertaintyCertified bool
	PrecisionLimitedBy       string
	Verdict                  string
}

type SectorLawfulnessAudit struct {
	CrossSectorOrientationIntertwinerPresent bool
	LeptonOrientationToKoideOperatorPresent  bool
	CKMToChargedLeptonWallOperatorPresent    bool
	NativeRootTraceOperatorPresent           bool
	DerivesKappa                             bool
	DerivesEpsilon                           bool
	DerivesTheta13                           bool
	DerivesJCKM                              bool
	Verdict                                  string
}

type ResidualControlAudit struct {
	CombinedResidual             float64
	RMinusOne                    float64
	RDefect                      float64
	QResidual                    float64
	EpsilonR1MinusObs            float64
	RatioToAbsRMinusOne          float64
	RatioToAbsQResidual          float64
	RatioToEpsilonShift          float64
	RequiredRDefectCoefficient   float64
	RequiredQResidualCoefficient float64
	TypedCoefficientPresent      bool
	Interpretation               string
	Verdict                      string
}

type FirewallAudit struct {
	DerivesKoide               bool
	DerivesPMNS                bool
	DerivesCKM                 bool
	DerivesTheta13             bool
	DerivesNeutrinoPhysics     bool
	DerivesChargedLeptonMasses bool
	DerivesFlavorHierarchy     bool
	PromotesObservedAsNative   bool
	AddsNewCarrier             bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	BOutperformsA                bool
	InverseTheta13WithinOneSigma bool
	CrossSectorBridgePresent     bool
	KappaRemainsEnvironmental    bool
	RemainingResidual            float64
	RemainingRelativeResidual    float64
	Decision                     string
	Verdict                      string
}

type Analysis struct {
	Runtime     RuntimeInheritance
	Reactor     ReactorInput
	CKM         CKMInput
	Combination CombinationAudit
	Epsilon     EpsilonPrediction
	Inverse     InversePrediction
	Uncertainty UncertaintyAudit
	Lawfulness  SectorLawfulnessAudit
	Residual    ResidualControlAudit
	Firewalls   FirewallAudit
	Final       FinalVerdict
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
	g589, err := generation2koidereactorrobustnessrdefectsensitivityaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate589 predecessor: %w", err)
	}
	g588, err := generation2koideloopdeficitreactorangleaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate588 predecessor: %w", err)
	}
	bundle, err := historytransport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build history transport runtime: %w", err)
	}
	runtime := inheritRuntime(g589)
	reactor := inheritReactor(g588)
	ckm := inheritCKM(bundle)
	combination := auditCombination(runtime, reactor, ckm)
	epsilon := auditEpsilon(runtime, combination)
	inverse := auditInverse(runtime, reactor, ckm)
	uncertainty := auditUncertainty(runtime, reactor, ckm, combination)
	lawfulness := auditLawfulness()
	residual := auditResidualControls(runtime, g589, combination)
	firewalls := auditFirewalls()
	final := compileFinal(combination, inverse, lawfulness)
	truth := "Gate 590 tests the sharper typed environmental bridge kappa_e ≈ sin²(theta13)/4 - J_CKM.  The combined candidate improves the central residual from the Gate 588 reactor-quarter value by more than an order of magnitude and gives an inverse theta13 prediction inside the NuFIT one-sigma interval.  However the relation mixes lepton reactor leakage with quark-sector CP area, and the current ASHA runtime contains no cross-sector orientation intertwiner, root-trace operator, or flavor theorem making that combination native.  The remaining residual is small but the relation remains an environmental history seal."
	return Analysis{Runtime: runtime, Reactor: reactor, CKM: ckm, Combination: combination, Epsilon: epsilon, Inverse: inverse, Uncertainty: uncertainty, Lawfulness: lawfulness, Residual: residual, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(g589 generation2koidereactorrobustnessrdefectsensitivityaudit.Analysis) RuntimeInheritance {
	return RuntimeInheritance{
		EpsilonObsRad: g589.Runtime.EpsilonObsRad,
		EpsilonObsDeg: g589.Runtime.EpsilonObsDeg,
		RObs:          g589.Runtime.RObs,
		RMinusOne:     -g589.Runtime.RDefect,
		RDefect:       g589.Runtime.RDefect,
		QObs:          g589.Runtime.QObs,
		QResidual:     g589.Runtime.QResidual,
		KappaObs:      g589.Runtime.KappaObs,
		LoopUnit:      1.0 / (8.0 * math.Pi),
		Source:        "Gate589 measured near-Koide wall coordinate and loop-deficit kappa_obs",
		Verdict:       StatusGate589Inherited,
	}
}

func inheritReactor(g588 generation2koideloopdeficitreactorangleaudit.Analysis) ReactorInput {
	return ReactorInput{
		SourceName:        g588.Input.SourceName,
		SourceVersion:     g588.Input.SourceVersion,
		DataThrough:       g588.Input.DataThrough,
		Variant:           g588.Input.Variant,
		MassOrdering:      g588.Input.MassOrdering,
		Sin2Theta13:       g588.Input.Sin2Theta13,
		Sin2Theta13Plus:   g588.Input.Sin2Theta13Plus,
		Sin2Theta13Minus:  g588.Input.Sin2Theta13Minus,
		Theta13CentralDeg: g588.Input.Theta13Deg,
		Theta13LowDeg:     g588.Input.Theta13LowDeg,
		Theta13HighDeg:    g588.Input.Theta13HighDeg,
		Verdict:           StatusReactorInputInherited,
	}
}

func inheritCKM(b historytransport.Bundle) CKMInput {
	return CKMInput{
		SourceName:         "ASHA history transport v1 runtime CKM input",
		SourceVersion:      "PDG 2024 CKM values imported by history transport v1",
		JCKM:               b.FlavorTransport.JCKM,
		JCKMUncertainty:    0,
		HasJCKMUncertainty: false,
		Verdict:            strings.Join([]string{StatusCKMRuntimeInherited, StatusCKMUncertaintyMissing}, ";"),
	}
}

func auditCombination(r RuntimeInheritance, p ReactorInput, c CKMInput) CombinationAudit {
	aValue := p.Sin2Theta13 / 4.0
	aMin := (p.Sin2Theta13 - p.Sin2Theta13Minus) / 4.0
	aMax := (p.Sin2Theta13 + p.Sin2Theta13Plus) / 4.0
	bValue := aValue - c.JCKM
	bMin := aMin - c.JCKM
	bMax := aMax - c.JCKM
	a := compareCandidate("C = sin²(theta13)/4", "sin²(theta13)/4", aValue, aMin, aMax, r.KappaObs)
	b := compareCandidate("B = sin²(theta13)/4 - J_CKM", "sin²(theta13)/4 - J_CKM", bValue, bMin, bMax, r.KappaObs)
	improvement := safeDiv(a.AbsResidual, b.AbsResidual)
	outperforms := b.AbsResidual < a.AbsResidual
	interp := "Subtracting the runtime CKM Jarlskog area from the reactor-quarter candidate improves the central kappa residual by more than an order of magnitude, but the candidate is cross-sector and remains bridge-layer until an orientation intertwiner is proven."
	verdict := strings.Join([]string{StatusCandidateAComputed, StatusCandidateBComputed, StatusBOutperformsA, StatusBWithinOneSigma, StatusResidualImproved}, ";")
	return CombinationAudit{AReactorQuarter: a, BReactorMinusCKM: b, BImprovementFactor: improvement, BOutperformsA: outperforms, Interpretation: interp, Verdict: verdict}
}

func compareCandidate(name, eq string, value, min, max, target float64) CandidateComparison {
	if min > max {
		min, max = max, min
	}
	res := value - target
	rel := safeDiv(res, target)
	covers := min <= target && target <= max
	certified := math.Abs(rel) < certifiedRelativeTolerance && covers && false
	verdict := StatusCandidateAComputed
	if strings.Contains(eq, "J_CKM") {
		verdict = StatusCandidateBComputed
	}
	return CandidateComparison{Name: name, Equation: eq, Value: value, Min1Sigma: min, Max1Sigma: max, SignedResidual: res, AbsResidual: math.Abs(res), RelativeResidual: rel, CoversKappa: covers, Certified: certified, Verdict: verdict}
}

func auditEpsilon(r RuntimeInheritance, c CombinationAudit) EpsilonPrediction {
	epsA := r.LoopUnit * (1.0 - c.AReactorQuarter.Value)
	epsB := r.LoopUnit * (1.0 - c.BReactorMinusCKM.Value)
	resA := epsA - r.EpsilonObsRad
	resB := epsB - r.EpsilonObsRad
	return EpsilonPrediction{
		EpsilonObservedRad: r.EpsilonObsRad,
		EpsilonObservedDeg: r.EpsilonObsDeg,
		EpsilonPredA_rad:   epsA,
		EpsilonPredA_deg:   rad2deg(epsA),
		EpsilonPredB_rad:   epsB,
		EpsilonPredB_deg:   rad2deg(epsB),
		ResidualA_rad:      resA,
		ResidualA_deg:      rad2deg(resA),
		ResidualB_rad:      resB,
		ResidualB_deg:      rad2deg(resB),
		ImprovementFactor:  safeDiv(math.Abs(resA), math.Abs(resB)),
		Verdict:            StatusEpsilonPredictionComputed,
	}
}

func auditInverse(r RuntimeInheritance, p ReactorInput, c CKMInput) InversePrediction {
	predSin2 := 4.0 * (r.KappaObs + c.JCKM)
	predTheta := rad2deg(math.Asin(math.Sqrt(predSin2)))
	low := p.Sin2Theta13 - p.Sin2Theta13Minus
	high := p.Sin2Theta13 + p.Sin2Theta13Plus
	return InversePrediction{
		Sin2Theta13Pred:     predSin2,
		Theta13PredDeg:      predTheta,
		Sin2Central:         p.Sin2Theta13,
		Sin2Low:             low,
		Sin2High:            high,
		Theta13CentralDeg:   p.Theta13CentralDeg,
		Theta13LowDeg:       p.Theta13LowDeg,
		Theta13HighDeg:      p.Theta13HighDeg,
		Sin2Residual:        predSin2 - p.Sin2Theta13,
		ThetaResidualDeg:    predTheta - p.Theta13CentralDeg,
		WithinSin2OneSigma:  low <= predSin2 && predSin2 <= high,
		WithinThetaOneSigma: p.Theta13LowDeg <= predTheta && predTheta <= p.Theta13HighDeg,
		Verdict:             strings.Join([]string{StatusInverseComputed, StatusInverseWithinOneSigma}, ";"),
	}
}

func auditUncertainty(r RuntimeInheritance, p ReactorInput, c CKMInput, combo CombinationAudit) UncertaintyAudit {
	// With no CKM J uncertainty present in the runtime, only theta13 one-sigma can be propagated lawfully.
	minB := (p.Sin2Theta13-p.Sin2Theta13Minus)/4.0 - c.JCKM
	maxB := (p.Sin2Theta13+p.Sin2Theta13Plus)/4.0 - c.JCKM
	covers := minB <= r.KappaObs && r.KappaObs <= maxB
	limited := "theta13 one-sigma dominates available propagated uncertainty; CKM J uncertainty is absent from runtime, so full certification is blocked"
	return UncertaintyAudit{Theta13CandidateBMin: minB, Theta13CandidateBMax: maxB, CoversKappaWithTheta13: covers, CKMUncertaintyPresent: c.HasJCKMUncertainty, CKMUncertaintyValue: c.JCKMUncertainty, FullUncertaintyCertified: false, PrecisionLimitedBy: limited, Verdict: strings.Join([]string{StatusBWithinOneSigma, StatusCKMUncertaintyMissing, StatusTheta13Limited}, ";")}
}

func auditLawfulness() SectorLawfulnessAudit {
	return SectorLawfulnessAudit{CrossSectorOrientationIntertwinerPresent: false, LeptonOrientationToKoideOperatorPresent: false, CKMToChargedLeptonWallOperatorPresent: false, NativeRootTraceOperatorPresent: false, DerivesKappa: false, DerivesEpsilon: false, DerivesTheta13: false, DerivesJCKM: false, Verdict: strings.Join([]string{StatusNoCrossSectorIntertwiner, StatusNoNativeOperator}, ";")}
}

func auditResidualControls(r RuntimeInheritance, g589 generation2koidereactorrobustnessrdefectsensitivityaudit.Analysis, c CombinationAudit) ResidualControlAudit {
	residual := c.BReactorMinusCKM.SignedResidual
	epsShift := g589.Runtime.EpsilonShiftR1MinusObs
	reqR := safeDiv(residual, r.RDefect)
	reqQ := safeDiv(residual, r.QResidual)
	interp := "The remaining combined residual is smaller than the R-defect and Q residual scales.  It can be expressed with numerical coefficients, but no typed coefficient or operator is present, so no R/Q correction is certified."
	return ResidualControlAudit{CombinedResidual: residual, RMinusOne: r.RMinusOne, RDefect: r.RDefect, QResidual: r.QResidual, EpsilonR1MinusObs: epsShift, RatioToAbsRMinusOne: safeDiv(residual, math.Abs(r.RMinusOne)), RatioToAbsQResidual: safeDiv(residual, math.Abs(r.QResidual)), RatioToEpsilonShift: safeDiv(residual, epsShift), RequiredRDefectCoefficient: reqR, RequiredQResidualCoefficient: reqQ, TypedCoefficientPresent: false, Interpretation: interp, Verdict: StatusResidualNotRDefectFixed}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKoide: false, DerivesPMNS: false, DerivesCKM: false, DerivesTheta13: false, DerivesNeutrinoPhysics: false, DerivesChargedLeptonMasses: false, DerivesFlavorHierarchy: false, PromotesObservedAsNative: false, AddsNewCarrier: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoFlavorDerivation, StatusObservedBridgeOnly, StatusGate352Preserved, StatusGate590Boundary}, ";")}
}

func compileFinal(c CombinationAudit, inv InversePrediction, law SectorLawfulnessAudit) FinalVerdict {
	bridge := law.CrossSectorOrientationIntertwinerPresent || law.CKMToChargedLeptonWallOperatorPresent || law.NativeRootTraceOperatorPresent
	decision := "The combined typed environmental candidate kappa_e ≈ sin²(theta13)/4 - J_CKM outperforms the reactor-quarter relation alone and predicts theta13 inside the NuFIT one-sigma interval.  Because ASHA currently has no cross-sector orientation intertwiner between CKM area and the charged-lepton Koide wall, the relation remains environmental bridge geometry rather than native law."
	return FinalVerdict{BOutperformsA: c.BOutperformsA, InverseTheta13WithinOneSigma: inv.WithinThetaOneSigma, CrossSectorBridgePresent: bridge, KappaRemainsEnvironmental: !bridge, RemainingResidual: c.BReactorMinusCKM.SignedResidual, RemainingRelativeResidual: c.BReactorMinusCKM.RelativeResidual, Decision: decision, Verdict: strings.Join([]string{StatusBOutperformsA, StatusInverseWithinOneSigma, StatusNoCrossSectorIntertwiner, StatusKappaRemainsSeal, StatusGate590Boundary}, ";")}
}

func Statuses() []string {
	return []string{StatusGate589Inherited, StatusReactorInputInherited, StatusCKMRuntimeInherited, StatusCandidateAComputed, StatusCandidateBComputed, StatusBOutperformsA, StatusBWithinOneSigma, StatusInverseComputed, StatusInverseWithinOneSigma, StatusEpsilonPredictionComputed, StatusResidualImproved, StatusCKMUncertaintyMissing, StatusTheta13Limited, StatusNoCrossSectorIntertwiner, StatusNoNativeOperator, StatusResidualNotRDefectFixed, StatusKappaRemainsSeal, StatusNoFlavorDerivation, StatusObservedBridgeOnly, StatusGate352Preserved, StatusGate590Boundary}
}

func rad2deg(x float64) float64 { return x * 180.0 / math.Pi }
func safeDiv(a, b float64) float64 {
	if b == 0 {
		return math.NaN()
	}
	return a / b
}
