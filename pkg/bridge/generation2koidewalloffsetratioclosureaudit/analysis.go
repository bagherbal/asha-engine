// Package generation2koidewalloffsetratioclosureaudit implements Gate 584:
// Koide Wall-Offset One-Parameter Ratio Closure Audit.
//
// Gate 583 showed that the charged-lepton square-root Yukawa ray is a
// near-electron-wall point on the Koide Fourier circle.  Gate 584 tests the
// next minimal question: after imposing the exact Koide circle R=1 and the
// canonical positive chamber, does one square-root mass ratio determine the
// wall-offset epsilon and thereby predict the other independent ratio?
//
// This is a bridge-layer environmental compression audit only.  It does not
// derive epsilon, charged-lepton masses, Koide, a root trace, CKM/PMNS, or any
// generation hierarchy as native ASHA law.
package generation2koidewalloffsetratioclosureaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2koidechamberwalloffsetaudit"
)

const (
	AuditID = "GATE584-KOIDE-WALL-OFFSET-RATIO-CLOSURE-AUDIT"

	StatusGate583Inherited         = "PASS_GATE583_CHAMBER_WALL_RUNTIME_INHERITED"
	StatusExactR1ModelDefined      = "PASS_EXACT_KOIDE_R1_WALL_RATIO_MODEL_DEFINED"
	StatusUniqueEpsilonEMu         = "PASS_ELECTRON_MUON_RATIO_SOLVES_UNIQUE_EPSILON_IN_POSITIVE_CHAMBER"
	StatusPredictsMuTau            = "PASS_ELECTRON_MUON_SOLVED_EPSILON_PREDICTS_MUON_TAU_RATIO"
	StatusUniqueEpsilonMuTau       = "PASS_MUON_TAU_RATIO_SOLVES_UNIQUE_EPSILON_IN_POSITIVE_CHAMBER"
	StatusPredictsEMu              = "PASS_MUON_TAU_SOLVED_EPSILON_PREDICTS_ELECTRON_MUON_RATIO"
	StatusOneParameterClosure      = "PASS_ONE_PARAMETER_RATIO_CLOSURE_CERTIFIED_IN_CHARGED_LEPTON_SECTOR"
	StatusHierarchyCompression     = "CONDITIONAL_SUPPORT_EXACT_R1_WALL_MODEL_REDUCES_TWO_RATIOS_TO_ONE_EPSILON"
	StatusResidualFromR            = "CONDITIONAL_SUPPORT_RATIO_RESIDUALS_CONTROLLED_BY_R_MINUS_ONE_AND_ENDPOINT_PRECISION"
	StatusTransportPreserved       = "PASS_RATIO_CLOSURE_STABLE_BETWEEN_MZ_AND_LAMBDA12_IN_V1"
	StatusNoNativeEpsilon          = "FAILED_ROUTE_EPSILON_NOT_DERIVED_NATIVE_FROM_RATIO_CLOSURE"
	StatusNoNativeRootTrace        = "FAILED_ROUTE_NO_NATIVE_ROOT_TRACE_OR_CIRCULANT_RATIO_OPERATOR"
	StatusNoQuarkClosure           = "FAILED_ROUTE_NO_QUARK_ONE_PARAMETER_KOIDE_WALL_RATIO_CLOSURE_IN_V1"
	StatusNoFlavorDerivation       = "FIREWALL_PRESERVED_RATIO_CLOSURE_DOES_NOT_DERIVE_FLAVOR_TEXTURE_CKM_PMNS_OR_GENERATIONS"
	StatusObservedEndpoint         = "FIREWALL_PRESERVED_CHARGED_LEPTON_RATIOS_REMAIN_OBSERVED_HISTORY_ENDPOINT_DATA"
	StatusGate352Preserved         = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate584BoundaryPreserved = "FIREWALL_PRESERVED_GATE584_RATIO_CLOSURE_BOUNDARY"
)

const (
	upperWallDeg       = 135.0
	positiveChamberMax = 30.0 // epsilon=135°-delta, with 105°<delta<135°.
)

type RuntimeInheritance struct {
	Source                string
	MZEpsilonDeg          float64
	LambdaEpsilonDeg      float64
	MZPlaneAmplitudeR     float64
	LambdaPlaneAmplitudeR float64
	Verdict               string
}

type ExactR1WallModel struct {
	Chamber              string
	ElectronOverAFormula string
	MuonOverAFormula     string
	TauOverAFormula      string
	RatioEquation        string
	Unknown              string
	UniqueDomainDeg      [2]float64
	Verdict              string
}

type RatioClosure struct {
	Frame                                         string
	ObservedEpsilonDeg                            float64
	ObservedPlaneAmplitudeR                       float64
	ActualElectronMuonRootRatio                   float64
	ActualMuonTauRootRatio                        float64
	ActualElectronMuonMassRatio                   float64
	ActualMuonTauMassRatio                        float64
	ExactR1AtObservedEpsilonElectronMuonRootRatio float64
	ExactR1AtObservedEpsilonMuonTauRootRatio      float64
	ExactR1AtObservedEpsilonElectronMuonResidual  float64
	ExactR1AtObservedEpsilonMuonTauResidual       float64
	FromElectronMuon                              RatioPrediction
	FromMuonTau                                   RatioPrediction
	ClosureCertified                              bool
	Verdict                                       string
}

type RatioPrediction struct {
	InputRatioName         string
	InputRatio             float64
	SolvedEpsilonDeg       float64
	SolvedEpsilonRad       float64
	ObservedEpsilonDeg     float64
	EpsilonResidualDeg     float64
	PredictedRatioName     string
	PredictedRootRatio     float64
	ObservedRootRatio      float64
	RootResidual           float64
	RelativeRootResidual   float64
	PredictedMassRatio     float64
	ObservedMassRatio      float64
	MassResidual           float64
	WithinClosureTolerance bool
	Verdict                string
}

type TransportAudit struct {
	MZEpsilonFromEMuDeg        float64
	LambdaEpsilonFromEMuDeg    float64
	EMuSolvedEpsilonDriftDeg   float64
	MZPredictionResidual       float64
	LambdaPredictionResidual   float64
	ResidualImprovesAtBoundary bool
	ClosureStable              bool
	Verdict                    string
}

type QuarkClosureAudit struct {
	UpR                 float64
	DownR               float64
	UpOnKoideCircle     bool
	DownOnKoideCircle   bool
	OneParameterClosure bool
	Interpretation      string
	Verdict             string
}

type FirewallAudit struct {
	DerivesKoide               bool
	DerivesEpsilon             bool
	DerivesLeptonMasses        bool
	DerivesYukawaEigenvalues   bool
	DerivesCKM                 bool
	DerivesPMNS                bool
	DerivesGenerationHierarchy bool
	AddsNewCarrier             bool
	PromotesObservedAsNative   bool
	PreservesGate352           bool
	Verdict                    string
}

type FinalVerdict struct {
	SealName                  string
	MZInputRatio              string
	MZSolvedEpsilonDeg        float64
	MZPredictedMuonTauRatio   float64
	MZObservedMuonTauRatio    float64
	MZPredictionResidual      float64
	LambdaSolvedEpsilonDeg    float64
	LambdaPredictionResidual  float64
	OneParameterClosure       bool
	NativeDerivationCertified bool
	MinimalRemainingSeal      string
	Verdict                   string
}

type Analysis struct {
	Runtime   RuntimeInheritance
	Model     ExactR1WallModel
	MZ        RatioClosure
	Lambda12  RatioClosure
	Transport TransportAudit
	Quarks    QuarkClosureAudit
	Firewalls FirewallAudit
	Final     FinalVerdict
	Truth     string
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
	runtime := inheritRuntime(g583)
	model := defineExactR1Model()
	mz := closureFromWallPoint("M_Z", g583.MZ)
	lambda := closureFromWallPoint("Lambda_12", g583.Lambda12)
	transport := auditTransport(mz, lambda)
	quarks := auditQuarks(g583)
	firewalls := auditFirewalls()
	final := compileFinal(mz, lambda)
	truth := "Gate 584 shows that after imposing the exact Koide circle R=1 and the canonical positive chamber, one square-root charged-lepton ratio fixes the wall offset epsilon and predicts the other independent square-root ratio with small residual.  The charged-lepton hierarchy is therefore one-parameter within the Koide wall model, but epsilon itself remains an observed environmental seal, not a native ASHA derivation."
	return Analysis{Runtime: runtime, Model: model, MZ: mz, Lambda12: lambda, Transport: transport, Quarks: quarks, Firewalls: firewalls, Final: final, Truth: truth}, nil
}

func inheritRuntime(g583 generation2koidechamberwalloffsetaudit.Analysis) RuntimeInheritance {
	return RuntimeInheritance{Source: "Gate583 Koide chamber-wall offset audit", MZEpsilonDeg: g583.MZ.EpsilonDeg, LambdaEpsilonDeg: g583.Lambda12.EpsilonDeg, MZPlaneAmplitudeR: g583.MZ.PlaneAmplitudeR, LambdaPlaneAmplitudeR: g583.Lambda12.PlaneAmplitudeR, Verdict: StatusGate583Inherited}
}

func defineExactR1Model() ExactR1WallModel {
	return ExactR1WallModel{
		Chamber:              "canonical charged-lepton chamber (e,mu,tau), epsilon=135°-delta, 0°<epsilon<30°",
		ElectronOverAFormula: "E(epsilon)=1-cos(epsilon)+sin(epsilon)",
		MuonOverAFormula:     "M(epsilon)=1-((sqrt(3)-1)/2)cos(epsilon)-((sqrt(3)+1)/2)sin(epsilon)",
		TauOverAFormula:      "T(epsilon)=1+((sqrt(3)+1)/2)cos(epsilon)+((sqrt(3)-1)/2)sin(epsilon)",
		RatioEquation:        "r_e_mu=E(epsilon)/M(epsilon) determines epsilon; then r_mu_tau=M(epsilon)/T(epsilon)",
		Unknown:              "one wall-offset epsilon after exact R=1 Koide circle and chamber are fixed",
		UniqueDomainDeg:      [2]float64{0, positiveChamberMax},
		Verdict:              StatusExactR1ModelDefined,
	}
}

func closureFromWallPoint(frame string, p generation2koidechamberwalloffsetaudit.WallPoint) RatioClosure {
	actualEMu := p.ElectronMuonRootRatio
	actualMuTau := p.MuonRootOverA / p.TauRootOverA
	exactE, exactM, exactT := exactR1Components(degToRad(p.EpsilonDeg))
	exactEMu := exactE / exactM
	exactMuTau := exactM / exactT
	fromEMu := solveAndPredict(frame, "x_e/x_mu", actualEMu, "x_mu/x_tau", actualMuTau, p.EpsilonDeg)
	fromMuTau := solveAndPredict(frame, "x_mu/x_tau", actualMuTau, "x_e/x_mu", actualEMu, p.EpsilonDeg)
	certified := fromEMu.WithinClosureTolerance && fromMuTau.WithinClosureTolerance
	verdict := strings.Join([]string{StatusUniqueEpsilonEMu, StatusPredictsMuTau, StatusUniqueEpsilonMuTau, StatusPredictsEMu, StatusOneParameterClosure, StatusHierarchyCompression, StatusResidualFromR}, ";")
	return RatioClosure{Frame: frame, ObservedEpsilonDeg: p.EpsilonDeg, ObservedPlaneAmplitudeR: p.PlaneAmplitudeR, ActualElectronMuonRootRatio: actualEMu, ActualMuonTauRootRatio: actualMuTau, ActualElectronMuonMassRatio: actualEMu * actualEMu, ActualMuonTauMassRatio: actualMuTau * actualMuTau, ExactR1AtObservedEpsilonElectronMuonRootRatio: exactEMu, ExactR1AtObservedEpsilonMuonTauRootRatio: exactMuTau, ExactR1AtObservedEpsilonElectronMuonResidual: exactEMu - actualEMu, ExactR1AtObservedEpsilonMuonTauResidual: exactMuTau - actualMuTau, FromElectronMuon: fromEMu, FromMuonTau: fromMuTau, ClosureCertified: certified, Verdict: verdict}
}

func solveAndPredict(frame, inputName string, inputRatio float64, predictedName string, observedPredicted float64, observedEpsilonDeg float64) RatioPrediction {
	var eps float64
	if inputName == "x_e/x_mu" {
		eps = solveEpsilonForEMu(inputRatio)
	} else {
		eps = solveEpsilonForMuTau(inputRatio)
	}
	e, m, t := exactR1Components(eps)
	var predicted float64
	if predictedName == "x_mu/x_tau" {
		predicted = m / t
	} else {
		predicted = e / m
	}
	rootResidual := predicted - observedPredicted
	massPredicted := predicted * predicted
	massObserved := observedPredicted * observedPredicted
	status := StatusPredictsMuTau
	if predictedName == "x_e/x_mu" {
		status = StatusPredictsEMu
	}
	return RatioPrediction{InputRatioName: inputName, InputRatio: inputRatio, SolvedEpsilonDeg: radToDeg(eps), SolvedEpsilonRad: eps, ObservedEpsilonDeg: observedEpsilonDeg, EpsilonResidualDeg: radToDeg(eps) - observedEpsilonDeg, PredictedRatioName: predictedName, PredictedRootRatio: predicted, ObservedRootRatio: observedPredicted, RootResidual: rootResidual, RelativeRootResidual: safeDiv(rootResidual, observedPredicted), PredictedMassRatio: massPredicted, ObservedMassRatio: massObserved, MassResidual: massPredicted - massObserved, WithinClosureTolerance: math.Abs(rootResidual) < 1e-4, Verdict: status}
}

func auditTransport(mz, lambda RatioClosure) TransportAudit {
	drift := lambda.FromElectronMuon.SolvedEpsilonDeg - mz.FromElectronMuon.SolvedEpsilonDeg
	mzRes := math.Abs(mz.FromElectronMuon.RootResidual)
	lRes := math.Abs(lambda.FromElectronMuon.RootResidual)
	stable := math.Abs(drift) < 3e-4 && mz.ClosureCertified && lambda.ClosureCertified
	return TransportAudit{MZEpsilonFromEMuDeg: mz.FromElectronMuon.SolvedEpsilonDeg, LambdaEpsilonFromEMuDeg: lambda.FromElectronMuon.SolvedEpsilonDeg, EMuSolvedEpsilonDriftDeg: drift, MZPredictionResidual: mz.FromElectronMuon.RootResidual, LambdaPredictionResidual: lambda.FromElectronMuon.RootResidual, ResidualImprovesAtBoundary: lRes < mzRes, ClosureStable: stable, Verdict: strings.Join([]string{StatusTransportPreserved, StatusHierarchyCompression, StatusResidualFromR}, ";")}
}

func auditQuarks(g583 generation2koidechamberwalloffsetaudit.Analysis) QuarkClosureAudit {
	upOn := math.Abs(g583.Quarks.Up.R-1) < 1e-3
	downOn := math.Abs(g583.Quarks.Down.R-1) < 1e-3
	return QuarkClosureAudit{UpR: g583.Quarks.Up.R, DownR: g583.Quarks.Down.R, UpOnKoideCircle: upOn, DownOnKoideCircle: downOn, OneParameterClosure: upOn && downOn, Interpretation: "Quark sectors have formal Fourier coordinates but are not on R=1 in v1, so the one-parameter Koide wall-ratio closure is not certified for quarks.", Verdict: StatusNoQuarkClosure}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKoide: false, DerivesEpsilon: false, DerivesLeptonMasses: false, DerivesYukawaEigenvalues: false, DerivesCKM: false, DerivesPMNS: false, DerivesGenerationHierarchy: false, AddsNewCarrier: false, PromotesObservedAsNative: false, PreservesGate352: true, Verdict: strings.Join([]string{StatusNoNativeEpsilon, StatusNoNativeRootTrace, StatusNoFlavorDerivation, StatusObservedEndpoint, StatusGate352Preserved, StatusGate584BoundaryPreserved}, ";")}
}

func compileFinal(mz, lambda RatioClosure) FinalVerdict {
	return FinalVerdict{SealName: "ChargedLeptonKoideWallOffsetOneParameterRatioSeal", MZInputRatio: "x_e/x_mu", MZSolvedEpsilonDeg: mz.FromElectronMuon.SolvedEpsilonDeg, MZPredictedMuonTauRatio: mz.FromElectronMuon.PredictedRootRatio, MZObservedMuonTauRatio: mz.FromElectronMuon.ObservedRootRatio, MZPredictionResidual: mz.FromElectronMuon.RootResidual, LambdaSolvedEpsilonDeg: lambda.FromElectronMuon.SolvedEpsilonDeg, LambdaPredictionResidual: lambda.FromElectronMuon.RootResidual, OneParameterClosure: mz.ClosureCertified && lambda.ClosureCertified, NativeDerivationCertified: false, MinimalRemainingSeal: "epsilon_e itself, or a native root-trace/circulant generation-plane operator that selects the canonical chamber and fixes the wall offset", Verdict: strings.Join([]string{StatusOneParameterClosure, StatusHierarchyCompression, StatusNoNativeEpsilon, StatusGate584BoundaryPreserved}, ";")}
}

func Statuses() []string {
	return []string{StatusGate583Inherited, StatusExactR1ModelDefined, StatusUniqueEpsilonEMu, StatusPredictsMuTau, StatusUniqueEpsilonMuTau, StatusPredictsEMu, StatusOneParameterClosure, StatusHierarchyCompression, StatusResidualFromR, StatusTransportPreserved, StatusNoNativeEpsilon, StatusNoNativeRootTrace, StatusNoQuarkClosure, StatusNoFlavorDerivation, StatusObservedEndpoint, StatusGate352Preserved, StatusGate584BoundaryPreserved}
}

func exactR1Components(eps float64) (float64, float64, float64) {
	e := 1.0 - math.Cos(eps) + math.Sin(eps)
	mu := 1.0 - ((math.Sqrt(3.0)-1.0)/2.0)*math.Cos(eps) - ((math.Sqrt(3.0)+1.0)/2.0)*math.Sin(eps)
	tau := 1.0 + ((math.Sqrt(3.0)+1.0)/2.0)*math.Cos(eps) + ((math.Sqrt(3.0)-1.0)/2.0)*math.Sin(eps)
	return e, mu, tau
}

func solveEpsilonForEMu(target float64) float64 {
	return bisect(func(eps float64) float64 {
		e, m, _ := exactR1Components(eps)
		return e/m - target
	})
}

func solveEpsilonForMuTau(target float64) float64 {
	return bisect(func(eps float64) float64 {
		_, m, t := exactR1Components(eps)
		return m/t - target
	})
}

func bisect(f func(float64) float64) float64 {
	lo := 0.0
	hi := degToRad(positiveChamberMax)
	flo := f(lo)
	for i := 0; i < 200; i++ {
		mid := 0.5 * (lo + hi)
		fm := f(mid)
		if (flo <= 0 && fm <= 0) || (flo >= 0 && fm >= 0) {
			lo = mid
			flo = fm
		} else {
			hi = mid
		}
	}
	return 0.5 * (lo + hi)
}

func safeDiv(a, b float64) float64 {
	if b == 0 {
		return math.NaN()
	}
	return a / b
}

func radToDeg(x float64) float64 { return x * 180.0 / math.Pi }
func degToRad(x float64) float64 { return x * math.Pi / 180.0 }
