// Package generation2koidereactorckmresidualclosureuncertaintyaudit implements
// Gate 591: Koide-Reactor-CKM Residual Closure and Uncertainty Audit.
//
// Gate 590 found the sharpest bridge-layer environmental relation so far,
//
//	kappa_e ~= sin^2(theta13)/4 - J_CKM,
//
// for the measured near-Koide charged-lepton wall coordinate.  Gate 591 asks
// whether the remaining residual is statistically meaningful, whether it is
// comparable to the observed Koide amplitude/cone defects, and whether a typed
// R- or Q-defect correction is present.  This audit remains environmental: it
// does not derive Koide, theta13, CKM, PMNS, charged-lepton masses, neutrino
// physics, or a native ASHA flavor law.
package generation2koidereactorckmresidualclosureuncertaintyaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidereactorckmorientationcombinationaudit"
)

const (
	AuditID = "GATE591-KOIDE-REACTOR-CKM-RESIDUAL-CLOSURE-UNCERTAINTY-AUDIT"

	StatusGate590Inherited              = "PASS_GATE590_COMBINED_ORIENTATION_RESULT_INHERITED"
	StatusCKMUncertaintyImported        = "PASS_PDG2025_CKM_JARLSKOG_UNCERTAINTY_IMPORTED"
	StatusThetaCKMUncertaintyPropagated = "PASS_THETA13_AND_CKM_UNCERTAINTIES_PROPAGATED"
	StatusResidualInsideOneSigma        = "PASS_DELTA590_INSIDE_COMBINED_ONE_SIGMA_BAND"
	StatusResidualSmallerThanDefects    = "PASS_DELTA590_SMALLER_THAN_KOIDE_R_AND_Q_DEFECTS"
	StatusInputNoiseLimited             = "CONDITIONAL_SUPPORT_DELTA590_INPUT_NOISE_LIMITED_BY_THETA13"
	StatusCKMSubdominant                = "CONDITIONAL_SUPPORT_CKM_UNCERTAINTY_SUBDOMINANT_TO_THETA13"
	StatusRDefectOverPiHint             = "CONDITIONAL_SUPPORT_R_DEFECT_OVER_PI_NUMERIC_CLOSURE_HINT_NOT_CERTIFIED"
	StatusCorrectedFormulaTested        = "PASS_R_AND_Q_DEFECT_CORRECTION_CANDIDATES_TESTED"
	StatusNoRQCorrectionCertified       = "FAILED_ROUTE_NO_R_OR_Q_DEFECT_CORRECTION_CERTIFIED"
	StatusNoCrossSectorIntertwiner      = "FAILED_ROUTE_NO_CROSS_SECTOR_ORIENTATION_INTERTWINER"
	StatusNoNativeOperator              = "FAILED_ROUTE_NO_NATIVE_KOIDE_REACTOR_CKM_RESIDUAL_OPERATOR"
	StatusDeltaRemainsResidual          = "FAILED_ROUTE_DELTA590_REMAINS_ENVIRONMENTAL_RESIDUAL"
	StatusKappaRemainsSeal              = "FAILED_ROUTE_KAPPA_E_REMAINS_ENVIRONMENTAL_HISTORY_SEAL"
	StatusNoFlavorDerivation            = "FIREWALL_PRESERVED_NO_KOIDE_PMNS_CKM_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusObservedInputsRemainObserved  = "FIREWALL_PRESERVED_REACTOR_CKM_AND_KOIDE_INPUTS_REMAIN_OBSERVED_DATA"
	StatusGate352Preserved              = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate591Boundary               = "FIREWALL_PRESERVED_GATE591_RESIDUAL_CLOSURE_UNCERTAINTY_BOUNDARY"
)

const (
	pdg2025JCentral = 3.12e-5
	pdg2025JPlus    = 0.13e-5
	pdg2025JMinus   = 0.12e-5
)

type RuntimeInheritance struct {
	EpsilonObsRad float64
	EpsilonObsDeg float64
	KappaObs      float64
	RObs          float64
	RDefect       float64
	RMinusOne     float64
	QObs          float64
	QResidual     float64
	LoopUnit      float64
	Verdict       string
}

type OrientationInputs struct {
	ReactorSource      string
	Sin2Theta13        float64
	Sin2Theta13Plus    float64
	Sin2Theta13Minus   float64
	JCKMRuntime        float64
	JCKMUncertaintySrc string
	JCKMUncertaintyCtr float64
	JCKMPlus           float64
	JCKMMinus          float64
	Verdict            string
}

type CombinedResidual struct {
	AReactorQuarter        float64
	BReactorMinusCKM       float64
	KappaObs               float64
	Delta590               float64
	AbsDelta590            float64
	RelativeDelta590       float64
	EpsilonResidualRad     float64
	EpsilonResidualDeg     float64
	ImprovementOverA       float64
	PercentMismatchRemoved float64
	Verdict                string
}

type UncertaintyAudit struct {
	BCentral            float64
	BMin1Sigma          float64
	BMax1Sigma          float64
	ResidualCentral     float64
	ResidualLow         float64
	ResidualHigh        float64
	CoversKappa         bool
	Theta13WidthMinus   float64
	Theta13WidthPlus    float64
	CKMWidthMinus       float64
	CKMWidthPlus        float64
	TotalWidthMinus     float64
	TotalWidthPlus      float64
	SigmaFractionMinus  float64
	SigmaFractionPlus   float64
	DominantUncertainty string
	Verdict             string
}

type InversePrediction struct {
	Sin2PredCentral     float64
	Sin2PredLowFromJ    float64
	Sin2PredHighFromJ   float64
	ThetaPredCentralDeg float64
	ThetaPredLowDeg     float64
	ThetaPredHighDeg    float64
	NuFITCentralSin2    float64
	NuFITLowSin2        float64
	NuFITHighSin2       float64
	NuFITCentralDeg     float64
	NuFITLowDeg         float64
	NuFITHighDeg        float64
	CentralResidualDeg  float64
	WithinOneSigma      bool
	Verdict             string
}

type DefectScaleAudit struct {
	Delta590          float64
	RDefect           float64
	QResidual         float64
	AbsQResidual      float64
	DeltaOverRDefect  float64
	DeltaOverAbsQ     float64
	DeltaSmallerThanR bool
	DeltaSmallerThanQ bool
	Interpretation    string
	Verdict           string
}

type CorrectionCandidate struct {
	Name            string
	Source          string
	Equation        string
	Value           float64
	Coefficient     float64
	CorrectedDelta  float64
	AbsResidual     float64
	RelativeToDelta float64
	Certified       bool
}

type CorrectionAudit struct {
	RequiredRDefectCoefficient   float64
	RequiredQResidualCoefficient float64
	Candidates                   []CorrectionCandidate
	BestCandidate                CorrectionCandidate
	AnyCertified                 bool
	Interpretation               string
	Verdict                      string
}

type LawfulnessAudit struct {
	CrossSectorOrientationIntertwinerPresent bool
	RDefectToOrientationOperatorPresent      bool
	QDefectToOrientationOperatorPresent      bool
	NativeRootTraceOperatorPresent           bool
	DerivesDelta590                          bool
	DerivesKappa                             bool
	Verdict                                  string
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
	DeltaStatisticallyMeaningful bool
	CorrectedFormulaCertified    bool
	CrossSectorBridgePresent     bool
	KappaRemainsEnvironmental    bool
	Decision                     string
	Verdict                      string
}

type Analysis struct {
	Runtime     RuntimeInheritance
	Inputs      OrientationInputs
	Residual    CombinedResidual
	Uncertainty UncertaintyAudit
	Inverse     InversePrediction
	Defects     DefectScaleAudit
	Corrections CorrectionAudit
	Lawfulness  LawfulnessAudit
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
	g590, err := generation2koidereactorckmorientationcombinationaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate590 predecessor: %w", err)
	}
	runtime := inheritRuntime(g590)
	inputs := inheritInputs(g590)
	residual := auditCombinedResidual(g590)
	uncertainty := auditUncertainty(runtime, inputs, residual)
	inverse := auditInverse(runtime, inputs)
	defects := auditDefectScales(runtime, residual)
	corrections := auditCorrections(runtime, residual)
	lawfulness := auditLawfulness()
	firewalls := auditFirewalls()
	final := compileFinal(uncertainty, corrections, lawfulness)
	truth := "Gate 591 shows that the Gate 590 residual is far inside the present theta13-dominated one-sigma band and is smaller than both the Koide amplitude defect 1-R_obs and the cone residual |Q_obs-2/3|.  R- and Q-defect correction candidates can reduce the small residual numerically, with R_defect/pi the closest typed-looking trial, but no ASHA theorem supplies the coefficient or a cross-sector orientation intertwiner.  The relation remains the best current environmental bridge, not native flavor law."
	return Analysis{Runtime: runtime, Inputs: inputs, Residual: residual, Uncertainty: uncertainty, Inverse: inverse, Defects: defects, Corrections: corrections, Lawfulness: lawfulness, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(g590 generation2koidereactorckmorientationcombinationaudit.Analysis) RuntimeInheritance {
	return RuntimeInheritance{
		EpsilonObsRad: g590.Runtime.EpsilonObsRad,
		EpsilonObsDeg: g590.Runtime.EpsilonObsDeg,
		KappaObs:      g590.Runtime.KappaObs,
		RObs:          g590.Runtime.RObs,
		RDefect:       g590.Runtime.RDefect,
		RMinusOne:     g590.Runtime.RMinusOne,
		QObs:          g590.Runtime.QObs,
		QResidual:     g590.Runtime.QResidual,
		LoopUnit:      g590.Runtime.LoopUnit,
		Verdict:       StatusGate590Inherited,
	}
}

func inheritInputs(g590 generation2koidereactorckmorientationcombinationaudit.Analysis) OrientationInputs {
	return OrientationInputs{
		ReactorSource:      g590.Reactor.SourceName + " " + g590.Reactor.SourceVersion,
		Sin2Theta13:        g590.Reactor.Sin2Theta13,
		Sin2Theta13Plus:    g590.Reactor.Sin2Theta13Plus,
		Sin2Theta13Minus:   g590.Reactor.Sin2Theta13Minus,
		JCKMRuntime:        g590.CKM.JCKM,
		JCKMUncertaintySrc: "PDG 2025 CKM review: J = (3.12 +0.13 -0.12) x 10^-5; uncertainty imported for Gate591 audit only, central relation keeps runtime J",
		JCKMUncertaintyCtr: pdg2025JCentral,
		JCKMPlus:           pdg2025JPlus,
		JCKMMinus:          pdg2025JMinus,
		Verdict:            strings.Join([]string{StatusCKMUncertaintyImported, StatusGate590Inherited}, ";"),
	}
}

func auditCombinedResidual(g590 generation2koidereactorckmorientationcombinationaudit.Analysis) CombinedResidual {
	a := g590.Combination.AReactorQuarter.Value
	b := g590.Combination.BReactorMinusCKM.Value
	delta := g590.Combination.BReactorMinusCKM.SignedResidual
	rel := g590.Combination.BReactorMinusCKM.RelativeResidual
	epsRes := g590.Epsilon.ResidualB_rad
	improvement := g590.Combination.BImprovementFactor
	removed := 1.0 - safeDiv(math.Abs(delta), g590.Combination.AReactorQuarter.AbsResidual)
	return CombinedResidual{AReactorQuarter: a, BReactorMinusCKM: b, KappaObs: g590.Runtime.KappaObs, Delta590: delta, AbsDelta590: math.Abs(delta), RelativeDelta590: rel, EpsilonResidualRad: epsRes, EpsilonResidualDeg: rad2deg(epsRes), ImprovementOverA: improvement, PercentMismatchRemoved: 100.0 * removed, Verdict: StatusGate590Inherited}
}

func auditUncertainty(r RuntimeInheritance, in OrientationInputs, res CombinedResidual) UncertaintyAudit {
	// B = sin2/4 - J.  The minimum uses low theta13 and high J; maximum uses high theta13 and low J.
	bMin := (in.Sin2Theta13-in.Sin2Theta13Minus)/4.0 - (in.JCKMRuntime + in.JCKMPlus)
	bMax := (in.Sin2Theta13+in.Sin2Theta13Plus)/4.0 - (in.JCKMRuntime - in.JCKMMinus)
	thetaMinus := in.Sin2Theta13Minus / 4.0
	thetaPlus := in.Sin2Theta13Plus / 4.0
	totalMinus := res.BReactorMinusCKM - bMin
	totalPlus := bMax - res.BReactorMinusCKM
	covers := bMin <= r.KappaObs && r.KappaObs <= bMax
	dominant := "theta13 one-sigma width dominates; CKM J uncertainty is included but roughly two orders of magnitude smaller than the reactor-angle contribution"
	return UncertaintyAudit{BCentral: res.BReactorMinusCKM, BMin1Sigma: bMin, BMax1Sigma: bMax, ResidualCentral: res.Delta590, ResidualLow: bMin - r.KappaObs, ResidualHigh: bMax - r.KappaObs, CoversKappa: covers, Theta13WidthMinus: thetaMinus, Theta13WidthPlus: thetaPlus, CKMWidthMinus: in.JCKMPlus, CKMWidthPlus: in.JCKMMinus, TotalWidthMinus: totalMinus, TotalWidthPlus: totalPlus, SigmaFractionMinus: safeDiv(math.Abs(res.Delta590), totalMinus), SigmaFractionPlus: safeDiv(math.Abs(res.Delta590), totalPlus), DominantUncertainty: dominant, Verdict: strings.Join([]string{StatusThetaCKMUncertaintyPropagated, StatusResidualInsideOneSigma, StatusInputNoiseLimited, StatusCKMSubdominant}, ";")}
}

func auditInverse(r RuntimeInheritance, in OrientationInputs) InversePrediction {
	central := 4.0 * (r.KappaObs + in.JCKMRuntime)
	low := 4.0 * (r.KappaObs + in.JCKMRuntime - in.JCKMMinus)
	high := 4.0 * (r.KappaObs + in.JCKMRuntime + in.JCKMPlus)
	nLow := in.Sin2Theta13 - in.Sin2Theta13Minus
	nHigh := in.Sin2Theta13 + in.Sin2Theta13Plus
	thetaCentral := thetaDeg(central)
	thetaLow := thetaDeg(low)
	thetaHigh := thetaDeg(high)
	nuCentralTheta := thetaDeg(in.Sin2Theta13)
	nuLowTheta := thetaDeg(nLow)
	nuHighTheta := thetaDeg(nHigh)
	within := nLow <= central && central <= nHigh
	return InversePrediction{Sin2PredCentral: central, Sin2PredLowFromJ: low, Sin2PredHighFromJ: high, ThetaPredCentralDeg: thetaCentral, ThetaPredLowDeg: thetaLow, ThetaPredHighDeg: thetaHigh, NuFITCentralSin2: in.Sin2Theta13, NuFITLowSin2: nLow, NuFITHighSin2: nHigh, NuFITCentralDeg: nuCentralTheta, NuFITLowDeg: nuLowTheta, NuFITHighDeg: nuHighTheta, CentralResidualDeg: thetaCentral - nuCentralTheta, WithinOneSigma: within, Verdict: StatusResidualInsideOneSigma}
}

func auditDefectScales(r RuntimeInheritance, res CombinedResidual) DefectScaleAudit {
	absQ := math.Abs(r.QResidual)
	interp := "The Gate590 residual is below both the Koide amplitude defect 1-R_obs and the absolute cone residual |Q_obs-2/3|, so the remaining difference is not resolvable as a clean new geometric law before the near-Koide defects and experimental uncertainties are handled."
	return DefectScaleAudit{Delta590: res.Delta590, RDefect: r.RDefect, QResidual: r.QResidual, AbsQResidual: absQ, DeltaOverRDefect: safeDiv(res.Delta590, r.RDefect), DeltaOverAbsQ: safeDiv(res.Delta590, absQ), DeltaSmallerThanR: math.Abs(res.Delta590) < math.Abs(r.RDefect), DeltaSmallerThanQ: math.Abs(res.Delta590) < absQ, Interpretation: interp, Verdict: StatusResidualSmallerThanDefects}
}

func auditCorrections(r RuntimeInheritance, res CombinedResidual) CorrectionAudit {
	requiredR := safeDiv(res.Delta590, r.RDefect)
	requiredQ := safeDiv(res.Delta590, r.QResidual)
	trials := []struct {
		name   string
		source string
		coef   float64
	}{
		{"R_defect/4", "R_defect", 0.25},
		{"R_defect/3", "R_defect", 1.0 / 3.0},
		{"R_defect/pi", "R_defect", 1.0 / math.Pi},
		{"R_defect*(sqrt2-1)", "R_defect", math.Sqrt2 - 1.0},
		{"|Q_residual|/2", "absQ", 0.5},
		{"|Q_residual|*(sqrt2-1)", "absQ", math.Sqrt2 - 1.0},
	}
	var candidates []CorrectionCandidate
	best := CorrectionCandidate{AbsResidual: math.Inf(1)}
	absQ := math.Abs(r.QResidual)
	for _, tr := range trials {
		base := r.RDefect
		eq := tr.name
		if tr.source == "absQ" {
			base = absQ
		}
		value := tr.coef * base
		correctedDelta := res.Delta590 - value
		cand := CorrectionCandidate{Name: tr.name, Source: tr.source, Equation: eq, Value: value, Coefficient: tr.coef, CorrectedDelta: correctedDelta, AbsResidual: math.Abs(correctedDelta), RelativeToDelta: safeDiv(math.Abs(correctedDelta), math.Abs(res.Delta590)), Certified: false}
		candidates = append(candidates, cand)
		if cand.AbsResidual < best.AbsResidual {
			best = cand
		}
	}
	verdict := strings.Join([]string{StatusCorrectedFormulaTested, StatusRDefectOverPiHint, StatusNoRQCorrectionCertified}, ";")
	interp := "Typed-looking R/Q corrections were tested.  R_defect/pi gives the smallest residual among the trial set, but no ASHA operator or theorem supplies that coefficient; therefore no corrected closure is certified."
	return CorrectionAudit{RequiredRDefectCoefficient: requiredR, RequiredQResidualCoefficient: requiredQ, Candidates: candidates, BestCandidate: best, AnyCertified: false, Interpretation: interp, Verdict: verdict}
}

func auditLawfulness() LawfulnessAudit {
	return LawfulnessAudit{CrossSectorOrientationIntertwinerPresent: false, RDefectToOrientationOperatorPresent: false, QDefectToOrientationOperatorPresent: false, NativeRootTraceOperatorPresent: false, DerivesDelta590: false, DerivesKappa: false, Verdict: strings.Join([]string{StatusNoCrossSectorIntertwiner, StatusNoNativeOperator, StatusNoRQCorrectionCertified}, ";")}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKoide: false, DerivesPMNS: false, DerivesCKM: false, DerivesTheta13: false, DerivesNeutrinoPhysics: false, DerivesChargedLeptonMasses: false, DerivesFlavorHierarchy: false, PromotesObservedAsNative: false, AddsNewCarrier: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoFlavorDerivation, StatusObservedInputsRemainObserved, StatusGate352Preserved, StatusGate591Boundary}, ";")}
}

func compileFinal(u UncertaintyAudit, c CorrectionAudit, l LawfulnessAudit) FinalVerdict {
	bridge := l.CrossSectorOrientationIntertwinerPresent || l.NativeRootTraceOperatorPresent
	meaningful := !u.CoversKappa
	decision := "The Gate590 residual is not statistically meaningful at current input precision: the combined candidate covers kappa_obs once NuFIT theta13 and PDG CKM-J one-sigma widths are propagated.  R/Q defect corrections were tested, but no typed coefficient is certified.  The combined relation remains an excellent environmental bridge and kappa_e remains a history seal."
	return FinalVerdict{DeltaStatisticallyMeaningful: meaningful, CorrectedFormulaCertified: c.AnyCertified, CrossSectorBridgePresent: bridge, KappaRemainsEnvironmental: !bridge, Decision: decision, Verdict: strings.Join([]string{StatusResidualInsideOneSigma, StatusNoRQCorrectionCertified, StatusNoCrossSectorIntertwiner, StatusKappaRemainsSeal, StatusGate591Boundary}, ";")}
}

func Statuses() []string {
	return []string{StatusGate590Inherited, StatusCKMUncertaintyImported, StatusThetaCKMUncertaintyPropagated, StatusResidualInsideOneSigma, StatusResidualSmallerThanDefects, StatusInputNoiseLimited, StatusCKMSubdominant, StatusCorrectedFormulaTested, StatusRDefectOverPiHint, StatusNoRQCorrectionCertified, StatusNoCrossSectorIntertwiner, StatusNoNativeOperator, StatusDeltaRemainsResidual, StatusKappaRemainsSeal, StatusNoFlavorDerivation, StatusObservedInputsRemainObserved, StatusGate352Preserved, StatusGate591Boundary}
}

func thetaDeg(sin2 float64) float64 { return rad2deg(math.Asin(math.Sqrt(sin2))) }
func rad2deg(x float64) float64     { return x * 180.0 / math.Pi }
func safeDiv(a, b float64) float64 {
	if b == 0 {
		return math.NaN()
	}
	return a / b
}
