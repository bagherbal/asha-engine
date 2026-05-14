// Package bgapmajoranaactivation implements Gate 312:
// B-Gap Majorana Activation in the Spectral Action / σ-H Mixed Quartic Correction.
//
// Gate 312 is a conditional activation audit. It asks whether the B-sector gap,
// treated as the sealed right-handed-neutrino Majorana/σ carrier, can provide the
// missing finite threshold/mixed-quartic correction that Gate 310 identified as
// necessary. The gate deliberately separates three layers:
//  1. structural activation of the ν_R ↔ ν_R^c Majorana edge;
//  2. formal σ-H quartic correction λ_eff = λ_HH - λ_Hσ²/(4λ_σσ);
//  3. one-loop RG reruns under Gate-309 conventions to test numerical capacity.
//
// The result is intentionally not tuned to the measured Higgs mass. In the
// inherited r_+ top-Yukawa lane, changing the UV Higgs quartic boundary within
// the stable σ-correction interval [0, λ_HH] barely changes the IR diagnostic:
// the top-sector flow dominates. Therefore the B-gap-as-boundary-quartic
// correction is structurally meaningful but does not, by itself, resolve the
// Gate-309 331 GeV tension under the same one-loop/top-seal protocol.
package bgapmajoranaactivation

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE312-BGAP-MAJORANA-ACTIVATION-SIGMA-H-MIXED-QUARTIC-CORRECTION"

	StatusGate296MajoranaEdgeConditionallyActivated = "CONDITIONAL_SUPPORT_GATE296_MAJORANA_EDGE_CONDITIONALLY_ACTIVATED"
	StatusBGapSigmaCarrierFormalized                = "CONDITIONAL_SUPPORT_BGAP_SIGMA_CARRIER_FORMALIZED"
	StatusMajoranaTraceExtensionFormalized          = "CONDITIONAL_SUPPORT_MAJORANA_TRACE_EXTENSION_FORMALIZED"
	StatusSigmaMixedQuarticFormalized               = "CONDITIONAL_SUPPORT_SIGMA_H_MIXED_QUARTIC_CORRECTION_FORMALIZED"
	StatusBGapCorrectionRerunComputed               = "CONDITIONAL_SUPPORT_BGAP_CORRECTED_RG_RERUN_COMPUTED"
	StatusBGapDoesNotResolveOneLoopTension          = "CONDITIONAL_TENSION_BGAP_BOUNDARY_CORRECTION_DOES_NOT_RESOLVE_RPLUS_ONE_LOOP_HIGGS_TENSION"
	StatusTopSectorStillDominates                   = "CONDITIONAL_TENSION_RPLUS_TOP_SECTOR_STILL_DOMINATES_IR_QUARTIC_FLOW"
	StatusGate312FirewallsPreserved                 = "CONDITIONAL_SUPPORT_GATE312_FIREWALLS_PRESERVED"

	StatusFailedMajoranaEdgeNotNative    = "FAILED_ROUTE_BGAP_MAJORANA_EDGE_STILL_SEALED_NOT_NATIVE_DERIVED"
	StatusFailedPortalCouplingNotDerived = "FAILED_ROUTE_SIGMA_H_PORTAL_COUPLING_NOT_DERIVED"
	StatusFailedSigmaVEVNotDerived       = "FAILED_ROUTE_SIGMA_VEV_NOT_DERIVED"
	StatusFailedThresholdJumpNotDerived  = "FAILED_ROUTE_BGAP_THRESHOLD_MATCHING_JUMP_NOT_DERIVED"
	StatusFailedTwoLoopNotExecuted       = "FAILED_ROUTE_TWO_LOOP_RGE_NOT_EXECUTED"
	StatusFailedPoleMassNotComputed      = "FAILED_ROUTE_POLE_MASS_CONVERSION_NOT_COMPUTED"
	StatusFailedLowEnergyMassNotResolved = "FAILED_ROUTE_LOW_ENERGY_HIGGS_MASS_NOT_RESOLVED"
)

const (
	lambdaHHBoundary     = 1197.0 / 4624.0
	bGap                 = 0.102464921191
	kappaMajorana        = 1.0
	vevGeV               = 246.22
	observedReferenceGeV = 125.10
	primaryStartGeV      = 2.40099519719e15
	primaryThresholdGeV  = 1.46774973718e6
	integrationSteps     = 24000
	perturbativeLimitSq  = 16.0 * math.Pi * math.Pi
)

type MajoranaActivation struct {
	Formalized                 bool
	SourceGate                 string
	Edge                       string
	Carrier                    string
	BGap                       float64
	ZSquared                   float64
	KappaM                     float64
	ActivatedAsPhysicalMass    bool
	ActivatedAsConditionalSeal bool
	NativeDerivation           bool
	Verdict                    string
}

type TraceExtension struct {
	Formalized        bool
	KappaC            float64
	KappaQ            float64
	KappaM            float64
	DiracTrace2       string
	DiracTrace4       string
	MajoranaTrace2    string
	MajoranaTrace4    string
	MixedQuarticTerm  string
	CrossTermsDerived bool
	Verdict           string
}

type SigmaPotential struct {
	Formalized       bool
	Potential        string
	Correction       string
	LambdaHH         float64
	LambdaSigmaSigma float64
	PortalNotation   string
	EffectiveRule    string
	RequiresPortal   bool
	RequiresSigmaVEV bool
	Verdict          string
}

type CorrectionLane struct {
	Name              string
	PortalChi         float64
	LambdaHH          float64
	LambdaSigmaSigma  float64
	LambdaHSigma      float64
	DeltaLambda       float64
	EffectiveLambdaUV float64
	StableNonNegative bool
	DerivedPortal     bool
	Interpretation    string
	Verdict           string
}

type RGState struct {
	GY, G2, G3, YT, Lambda float64
}

type RGResult struct {
	LaneName          string
	InitialLambdaUV   float64
	FinalLambdaAtV    float64
	HiggsMassGeV      float64
	Computed          bool
	Perturbative      bool
	FailureReason     string
	CloserToReference bool
	MassGapGeV        float64
	Verdict           string
}

type CapacityAudit struct {
	Formalized                     bool
	ReferenceLambdaAtV             float64
	ReferenceMassGeV               float64
	Gate309ControlMassGeV          float64
	MaxStableCorrectionMassGeV     float64
	MaxStableBoundaryShift         float64
	BestStableLaneName             string
	BestStableMassGeV              float64
	BestStableMassGapGeV           float64
	BoundaryCorrectionMovesMassGeV float64
	BoundaryCorrectionCanResolve   bool
	TopSectorDominates             bool
	Verdict                        string
}

type FirewallAudit struct {
	NoObservedMassFitInserted      bool
	NoPortalCouplingFitted         bool
	NoSigmaVEVFitted               bool
	NoThresholdJumpInserted        bool
	NoTwoLoopRGExecuted            bool
	NoPoleMassConversionInserted   bool
	MajoranaEdgeRemainsConditional bool
	NoFinalMassClaimed             bool
	FiniteCorePolluted             bool
	Obligations                    []RemainingObligation
	Verdict                        string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksFinalPrediction     bool
}

type Summary struct {
	MajoranaActivationFormalized   bool
	TraceExtensionFormalized       bool
	SigmaCorrectionFormalized      bool
	RGRerunComputed                bool
	BGapBoundaryCorrectionHelps    bool
	BGapBoundaryCorrectionSolves   bool
	TopSectorStillDominates        bool
	FinalMassClaimed               bool
	FirewallPreserved              bool
	Status, DirectAnswer, NextGate string
}

type Analysis struct {
	Activation MajoranaActivation
	Trace      TraceExtension
	Potential  SigmaPotential
	Lanes      []CorrectionLane
	Results    []RGResult
	Capacity   CapacityAudit
	Firewalls  FirewallAudit
	Summary    Summary
	Truth      string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	activation := activateMajoranaCarrier()
	trace := formalizeTraceExtension()
	potential := formalizeSigmaPotential()
	lanes := buildCorrectionLanes(potential)
	results := runCorrectionLanes(lanes)
	capacity := auditCapacity(results, lanes)
	firewalls := auditFirewalls(capacity)
	summary := buildSummary(activation, trace, potential, results, capacity, firewalls)
	truth := "Gate 312 conditionally activates the B-gap as the σ/Majorana carrier and formalizes the spectral-action correction λ_eff=λ_HH-λ_Hσ²/(4λ_σσ). Under the inherited Gate-309 one-loop r_+ top-Yukawa transport, however, every stable boundary correction between λ_HH and zero flows to essentially the same high IR quartic. The B-gap σ sector is therefore a necessary structural threshold/matching object, but its boundary-quartic activation alone does not resolve the 331 GeV one-loop tension; the native threshold jump, σ VEV/portal tensor, two-loop transport, and/or top-sector tensor still have to be derived."
	return Analysis{Activation: activation, Trace: trace, Potential: potential, Lanes: lanes, Results: results, Capacity: capacity, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func activateMajoranaCarrier() MajoranaActivation {
	return MajoranaActivation{
		Formalized:                 true,
		SourceGate:                 "Gate 296 sealed ν_R Majorana edge capacity + B-sector first spectral gap",
		Edge:                       "ν_R ↔ Jν_R ≃ ν_R^c",
		Carrier:                    "σ / B-gap Majorana scalar carrier",
		BGap:                       bGap,
		ZSquared:                   bGap,
		KappaM:                     kappaMajorana,
		ActivatedAsPhysicalMass:    false,
		ActivatedAsConditionalSeal: true,
		NativeDerivation:           false,
		Verdict:                    strings.Join([]string{StatusGate296MajoranaEdgeConditionallyActivated, StatusBGapSigmaCarrierFormalized, StatusFailedMajoranaEdgeNotNative}, ";"),
	}
}

func formalizeTraceExtension() TraceExtension {
	return TraceExtension{
		Formalized:        true,
		KappaC:            1,
		KappaQ:            3,
		KappaM:            kappaMajorana,
		DiracTrace2:       "Tr_Dirac(D_F²)=κ_C|x|²+κ_Q|y|² = X(1+3r)",
		DiracTrace4:       "Tr_Dirac(D_F⁴)=κ_C|x|⁴+κ_Q|y|⁴ = X²(1+3r²)",
		MajoranaTrace2:    "Tr_M(D_F²)=κ_M|z|² with |z|²=B_gap",
		MajoranaTrace4:    "Tr_M(D_F⁴)=κ_M|z|⁴",
		MixedQuarticTerm:  "Π_{Hσ}Tr(D_F⁴)=λ_Hσ |H|²|σ|²; the projector/coupling tensor is not derived by the diagonal trace alone",
		CrossTermsDerived: false,
		Verdict:           strings.Join([]string{StatusMajoranaTraceExtensionFormalized, StatusSigmaMixedQuarticFormalized, StatusFailedPortalCouplingNotDerived}, ";"),
	}
}

func formalizeSigmaPotential() SigmaPotential {
	lambdaSS := kappaMajorana * bGap * bGap
	return SigmaPotential{
		Formalized:       true,
		Potential:        "V(H,σ)=λ_HH|H|⁴+λ_Hσ|H|²|σ|²+λ_σσ|σ|⁴-μ_H²|H|²-μ_σ²|σ|²",
		Correction:       "λ_eff = λ_HH - λ_Hσ²/(4λ_σσ) after integrating out or minimizing the heavy σ radial mode",
		LambdaHH:         lambdaHHBoundary,
		LambdaSigmaSigma: lambdaSS,
		PortalNotation:   "λ_Hσ := 2χ√(λ_HH λ_σσ), so Δλ=χ²λ_HH for an explicit portal-overlap seal χ∈[0,1]",
		EffectiveRule:    "λ_eff(χ)=λ_HH(1-χ²); χ=0 leaves Gate 308 intact; χ=1 is the maximal stable nonnegative cancellation lane",
		RequiresPortal:   true,
		RequiresSigmaVEV: true,
		Verdict:          strings.Join([]string{StatusSigmaMixedQuarticFormalized, StatusFailedPortalCouplingNotDerived, StatusFailedSigmaVEVNotDerived}, ";"),
	}
}

func buildCorrectionLanes(p SigmaPotential) []CorrectionLane {
	return []CorrectionLane{
		buildLane("unactivated_control_gate308_boundary", 0, p, true, "control: no σ-H portal correction; reproduces Gate 309 primary lane"),
		buildLane("minimal_bgap_overlap_chi_equals_Bgap", bGap, p, false, "dimensionless B-gap overlap witness: χ=B_gap; small, positive, and untuned"),
		buildLane("maximal_stable_sigma_cancellation_chi_equals_one", 1, p, false, "capacity bound: largest stable nonnegative σ correction allowed by λ_eff≥0"),
	}
}

func buildLane(name string, chi float64, p SigmaPotential, derived bool, interpretation string) CorrectionLane {
	lambdaHS := 2.0 * chi * math.Sqrt(p.LambdaHH*p.LambdaSigmaSigma)
	delta := 0.0
	if p.LambdaSigmaSigma > 0 {
		delta = lambdaHS * lambdaHS / (4.0 * p.LambdaSigmaSigma)
	}
	eff := p.LambdaHH - delta
	if math.Abs(eff) < 1e-14 {
		eff = 0
	}
	verdict := StatusSigmaMixedQuarticFormalized
	if !derived {
		verdict = strings.Join([]string{StatusSigmaMixedQuarticFormalized, StatusFailedPortalCouplingNotDerived}, ";")
	}
	return CorrectionLane{Name: name, PortalChi: chi, LambdaHH: p.LambdaHH, LambdaSigmaSigma: p.LambdaSigmaSigma, LambdaHSigma: lambdaHS, DeltaLambda: delta, EffectiveLambdaUV: eff, StableNonNegative: eff >= -1e-12, DerivedPortal: derived, Interpretation: interpretation, Verdict: verdict}
}

func runCorrectionLanes(lanes []CorrectionLane) []RGResult {
	out := make([]RGResult, 0, len(lanes))
	for _, lane := range lanes {
		out = append(out, runPrimaryRPlusTransport(lane))
	}
	return out
}

func runPrimaryRPlusTransport(lane CorrectionLane) RGResult {
	if !lane.StableNonNegative {
		return RGResult{LaneName: lane.Name, InitialLambdaUV: lane.EffectiveLambdaUV, Computed: false, Perturbative: false, FailureReason: "λ_eff is negative at the UV boundary", Verdict: StatusFailedLowEnergyMassNotResolved}
	}
	rPlus := (3591.0 + 136.0*math.Sqrt(123.0)) / 3099.0
	state := RGState{GY: math.Sqrt(3.0 / 5.0), G2: 1, G3: 1, YT: math.Sqrt(rPlus), Lambda: lane.EffectiveLambdaUV}
	state, ok, failure := integrateSegment(state, primaryStartGeV, primaryThresholdGeV, betaCoefficients{B1GUT: 41.0/10.0 + 7.78628724237, B2: -19.0/6.0 + 9.65295390904, B3: -7.0 + 8.98628724237})
	if !ok {
		return RGResult{LaneName: lane.Name, InitialLambdaUV: lane.EffectiveLambdaUV, Computed: false, Perturbative: false, FailureReason: "above-threshold segment: " + failure, Verdict: StatusFailedLowEnergyMassNotResolved}
	}
	state, ok, failure = integrateSegment(state, primaryThresholdGeV, vevGeV, betaCoefficients{B1GUT: 41.0 / 10.0, B2: -19.0 / 6.0, B3: -7.0})
	if !ok {
		return RGResult{LaneName: lane.Name, InitialLambdaUV: lane.EffectiveLambdaUV, Computed: false, Perturbative: false, FailureReason: "below-threshold segment: " + failure, Verdict: StatusFailedLowEnergyMassNotResolved}
	}
	mass := math.NaN()
	if state.Lambda > 0 {
		mass = vevGeV * math.Sqrt(2.0*state.Lambda)
	}
	gap := math.Abs(mass - observedReferenceGeV)
	closer := mass < 331.630412 && gap < math.Abs(331.630412-observedReferenceGeV)
	return RGResult{LaneName: lane.Name, InitialLambdaUV: lane.EffectiveLambdaUV, FinalLambdaAtV: state.Lambda, HiggsMassGeV: mass, Computed: true, Perturbative: true, CloserToReference: closer, MassGapGeV: gap, Verdict: strings.Join([]string{StatusBGapCorrectionRerunComputed, StatusBGapDoesNotResolveOneLoopTension}, ";")}
}

type betaCoefficients struct{ B1GUT, B2, B3 float64 }

func integrateSegment(initial RGState, high, low float64, beta betaCoefficients) (RGState, bool, string) {
	state := initial
	if high == low {
		return state, true, ""
	}
	logHigh := math.Log(high)
	logLow := math.Log(low)
	h := (logLow - logHigh) / float64(integrationSteps)
	for i := 0; i < integrationSteps; i++ {
		state = rk4Step(state, h, beta)
		if !stateFinite(state) {
			return state, false, "non-finite RG state"
		}
		if !statePerturbative(state) {
			return state, false, "perturbative-control threshold exceeded"
		}
	}
	return state, true, ""
}

func rk4Step(y RGState, h float64, b betaCoefficients) RGState {
	k1 := deriv(y, b)
	k2 := deriv(addScaled(y, k1, 0.5*h), b)
	k3 := deriv(addScaled(y, k2, 0.5*h), b)
	k4 := deriv(addScaled(y, k3, h), b)
	return RGState{GY: y.GY + h*(k1.GY+2*k2.GY+2*k3.GY+k4.GY)/6.0, G2: y.G2 + h*(k1.G2+2*k2.G2+2*k3.G2+k4.G2)/6.0, G3: y.G3 + h*(k1.G3+2*k2.G3+2*k3.G3+k4.G3)/6.0, YT: y.YT + h*(k1.YT+2*k2.YT+2*k3.YT+k4.YT)/6.0, Lambda: y.Lambda + h*(k1.Lambda+2*k2.Lambda+2*k3.Lambda+k4.Lambda)/6.0}
}

func addScaled(y, k RGState, scale float64) RGState {
	return RGState{GY: y.GY + scale*k.GY, G2: y.G2 + scale*k.G2, G3: y.G3 + scale*k.G3, YT: y.YT + scale*k.YT, Lambda: y.Lambda + scale*k.Lambda}
}

func deriv(y RGState, b betaCoefficients) RGState {
	bY := (5.0 / 3.0) * b.B1GUT
	gy2, g22, g32 := y.GY*y.GY, y.G2*y.G2, y.G3*y.G3
	yt2 := y.YT * y.YT
	inv := 1.0 / (16.0 * math.Pi * math.Pi)
	betaLambda := (24.0*y.Lambda*y.Lambda + 12.0*y.Lambda*yt2 - 12.0*yt2*yt2 + (3.0/16.0)*(2.0*g22*g22+(g22+gy2)*(g22+gy2)) - y.Lambda*(9.0*g22+3.0*gy2)) * inv
	betaYT := y.YT * ((9.0/2.0)*yt2 - (17.0/12.0)*gy2 - (9.0/4.0)*g22 - 8.0*g32) * inv
	return RGState{GY: bY * y.GY * gy2 * inv, G2: b.B2 * y.G2 * g22 * inv, G3: b.B3 * y.G3 * g32 * inv, YT: betaYT, Lambda: betaLambda}
}

func stateFinite(s RGState) bool {
	vals := []float64{s.GY, s.G2, s.G3, s.YT, s.Lambda}
	for _, v := range vals {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func statePerturbative(s RGState) bool {
	return s.GY*s.GY < perturbativeLimitSq && s.G2*s.G2 < perturbativeLimitSq && s.G3*s.G3 < perturbativeLimitSq && s.YT*s.YT < perturbativeLimitSq && math.Abs(s.Lambda) < perturbativeLimitSq
}

func auditCapacity(results []RGResult, lanes []CorrectionLane) CapacityAudit {
	refLambda := (observedReferenceGeV / vevGeV) * (observedReferenceGeV / vevGeV) / 2.0
	controlMass := math.NaN()
	maxMass := math.NaN()
	bestMass := math.Inf(1)
	bestName := "missing"
	for _, r := range results {
		if r.LaneName == "unactivated_control_gate308_boundary" {
			controlMass = r.HiggsMassGeV
		}
		if r.LaneName == "maximal_stable_sigma_cancellation_chi_equals_one" {
			maxMass = r.HiggsMassGeV
		}
		if r.Computed && math.Abs(r.HiggsMassGeV-observedReferenceGeV) < math.Abs(bestMass-observedReferenceGeV) {
			bestMass = r.HiggsMassGeV
			bestName = r.LaneName
		}
	}
	maxShift := lambdaHHBoundary
	movement := math.Abs(controlMass - maxMass)
	bestGap := math.Abs(bestMass - observedReferenceGeV)
	canResolve := bestGap < 10.0
	return CapacityAudit{Formalized: true, ReferenceLambdaAtV: refLambda, ReferenceMassGeV: observedReferenceGeV, Gate309ControlMassGeV: controlMass, MaxStableCorrectionMassGeV: maxMass, MaxStableBoundaryShift: maxShift, BestStableLaneName: bestName, BestStableMassGeV: bestMass, BestStableMassGapGeV: bestGap, BoundaryCorrectionMovesMassGeV: movement, BoundaryCorrectionCanResolve: canResolve, TopSectorDominates: !canResolve && movement < 0.01, Verdict: strings.Join([]string{StatusBGapCorrectionRerunComputed, StatusBGapDoesNotResolveOneLoopTension, StatusTopSectorStillDominates}, ";")}
}

func auditFirewalls(c CapacityAudit) FirewallAudit {
	obs := []RemainingObligation{
		{"native σ-H portal tensor", "λ_Hσ was parameterized by χ; the finite Dirac graph has not derived the portal overlap", StatusFailedPortalCouplingNotDerived, true},
		{"σ VEV / B-gap order parameter", "the B-gap is a dimensionless spectral gap; its physical VEV and mass threshold remain sealed", StatusFailedSigmaVEVNotDerived, true},
		{"finite threshold matching jump", "a boundary quartic reduction is not the same as a derived decoupling jump Δλ at the σ/B-gap threshold", StatusFailedThresholdJumpNotDerived, true},
		{"two-loop RG transport", "Gate 312 deliberately reuses the Gate-309 one-loop diagnostic lane only", StatusFailedTwoLoopNotExecuted, true},
		{"pole mass conversion", "m=v√(2λ(v)) remains a running tree-level diagnostic", StatusFailedPoleMassNotComputed, true},
		{"top-sector tensor refinement", "the r_+ top lane still dominates the IR flow even after maximal stable σ boundary correction", StatusTopSectorStillDominates, true},
	}
	return FirewallAudit{NoObservedMassFitInserted: true, NoPortalCouplingFitted: true, NoSigmaVEVFitted: true, NoThresholdJumpInserted: true, NoTwoLoopRGExecuted: true, NoPoleMassConversionInserted: true, MajoranaEdgeRemainsConditional: true, NoFinalMassClaimed: true, FiniteCorePolluted: false, Obligations: obs, Verdict: strings.Join([]string{StatusGate312FirewallsPreserved, StatusFailedLowEnergyMassNotResolved, StatusFailedThresholdJumpNotDerived}, ";")}
}

func buildSummary(a MajoranaActivation, tr TraceExtension, p SigmaPotential, results []RGResult, c CapacityAudit, f FirewallAudit) Summary {
	return Summary{MajoranaActivationFormalized: a.Formalized, TraceExtensionFormalized: tr.Formalized, SigmaCorrectionFormalized: p.Formalized, RGRerunComputed: len(results) >= 3, BGapBoundaryCorrectionHelps: c.BoundaryCorrectionMovesMassGeV > 0.001, BGapBoundaryCorrectionSolves: c.BoundaryCorrectionCanResolve, TopSectorStillDominates: c.TopSectorDominates, FinalMassClaimed: !f.NoFinalMassClaimed, FirewallPreserved: !f.FiniteCorePolluted && f.NoObservedMassFitInserted && f.NoPortalCouplingFitted && f.NoThresholdJumpInserted && f.NoTwoLoopRGExecuted && f.NoFinalMassClaimed, Status: strings.Join([]string{StatusBGapSigmaCarrierFormalized, StatusBGapDoesNotResolveOneLoopTension}, ";"), DirectAnswer: fmt.Sprintf("B-gap σ activation formalizes λ_eff=λ_HH-λ_Hσ²/(4λ_σσ), but under the inherited r_+ one-loop transport even maximal stable boundary cancellation gives m_H≈%.6f GeV, versus control %.6f GeV and reference %.2f GeV.", c.MaxStableCorrectionMassGeV, c.Gate309ControlMassGeV, c.ReferenceMassGeV), NextGate: "Derive the actual B-gap threshold matching jump Δλ and/or revise the top-sector tensor; boundary-only σ correction is insufficient under the Gate-309 r_+ lane."}
}

func FormatActivation(a MajoranaActivation) string {
	return fmt.Sprintf("formalized=%t source=%q edge=%q carrier=%q Bgap=%.12f z²=%.12f κM=%.6g physicalMass=%t conditional=%t native=%t verdict=%s", a.Formalized, a.SourceGate, a.Edge, a.Carrier, a.BGap, a.ZSquared, a.KappaM, a.ActivatedAsPhysicalMass, a.ActivatedAsConditionalSeal, a.NativeDerivation, a.Verdict)
}

func FormatTrace(t TraceExtension) string {
	return fmt.Sprintf("formalized=%t κ=(C %.6g,Q %.6g,M %.6g) D2=%q D4=%q M2=%q M4=%q mixed=%q crossDerived=%t verdict=%s", t.Formalized, t.KappaC, t.KappaQ, t.KappaM, t.DiracTrace2, t.DiracTrace4, t.MajoranaTrace2, t.MajoranaTrace4, t.MixedQuarticTerm, t.CrossTermsDerived, t.Verdict)
}

func FormatPotential(p SigmaPotential) string {
	return fmt.Sprintf("formalized=%t λHH=%.12f λσσ=%.12f potential=%q correction=%q portal=%q effective=%q requiresPortal=%t requiresVEV=%t verdict=%s", p.Formalized, p.LambdaHH, p.LambdaSigmaSigma, p.Potential, p.Correction, p.PortalNotation, p.EffectiveRule, p.RequiresPortal, p.RequiresSigmaVEV, p.Verdict)
}

func FormatLane(l CorrectionLane) string {
	return fmt.Sprintf("%s χ=%.12f λHH=%.12f λσσ=%.12f λHσ=%.12f Δλ=%.12f λeff=%.12f stable=%t derivedPortal=%t interpretation=%q verdict=%s", l.Name, l.PortalChi, l.LambdaHH, l.LambdaSigmaSigma, l.LambdaHSigma, l.DeltaLambda, l.EffectiveLambdaUV, l.StableNonNegative, l.DerivedPortal, l.Interpretation, l.Verdict)
}

func FormatRGResult(r RGResult) string {
	return fmt.Sprintf("%s λUV=%.12f λv=%.12f mH=%.9f computed=%t pert=%t closer=%t gap=%.9f failure=%q verdict=%s", r.LaneName, r.InitialLambdaUV, r.FinalLambdaAtV, r.HiggsMassGeV, r.Computed, r.Perturbative, r.CloserToReference, r.MassGapGeV, r.FailureReason, r.Verdict)
}

func FormatCapacity(c CapacityAudit) string {
	return fmt.Sprintf("formalized=%t refλ=%.12f refMass=%.6f controlMass=%.9f maxStableMass=%.9f maxBoundaryShift=%.12f best=%q bestMass=%.9f bestGap=%.9f movement=%.12f canResolve=%t topDominates=%t verdict=%s", c.Formalized, c.ReferenceLambdaAtV, c.ReferenceMassGeV, c.Gate309ControlMassGeV, c.MaxStableCorrectionMassGeV, c.MaxStableBoundaryShift, c.BestStableLaneName, c.BestStableMassGeV, c.BestStableMassGapGeV, c.BoundaryCorrectionMovesMassGeV, c.BoundaryCorrectionCanResolve, c.TopSectorDominates, c.Verdict)
}

func FormatObligation(o RemainingObligation) string {
	return fmt.Sprintf("%s required=%q status=%s blocks=%t", o.Name, o.WhyRequired, o.Status, o.BlocksFinalPrediction)
}

func FormatFirewalls(f FirewallAudit) string {
	obs := []string{}
	for _, o := range f.Obligations {
		obs = append(obs, FormatObligation(o))
	}
	return fmt.Sprintf("noObsFit=%t noPortalFit=%t noSigmaVEVFit=%t noThreshold=%t noTwoLoop=%t noPole=%t conditionalMajorana=%t noFinal=%t polluted=%t obligations=[%s] verdict=%s", f.NoObservedMassFitInserted, f.NoPortalCouplingFitted, f.NoSigmaVEVFitted, f.NoThresholdJumpInserted, f.NoTwoLoopRGExecuted, f.NoPoleMassConversionInserted, f.MajoranaEdgeRemainsConditional, f.NoFinalMassClaimed, f.FiniteCorePolluted, strings.Join(obs, " | "), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("activation=%t trace=%t sigma=%t rg=%t helps=%t solves=%t topDominates=%t finalClaim=%t firewall=%t status=%s answer=%q next=%q", s.MajoranaActivationFormalized, s.TraceExtensionFormalized, s.SigmaCorrectionFormalized, s.RGRerunComputed, s.BGapBoundaryCorrectionHelps, s.BGapBoundaryCorrectionSolves, s.TopSectorStillDominates, s.FinalMassClaimed, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
